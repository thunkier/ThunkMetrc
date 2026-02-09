package common

import (
	"sort"

	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
	"github.com/thunkier/thunkmetrc/tools/pkg/jsonschema"
	"github.com/thunkier/thunkmetrc/tools/pkg/schema"
)

type TypeMapper func(s *schema.Schema, parentName, propName string) string
type TypeKindMapper func(t schema.SchemaType) jsonschema.TypeKind

type ModelNamer func(serviceName, opName string) string

func CollectRequestModels(services []models.Service, mapper TypeMapper, kindMapper TypeKindMapper, namer ModelNamer) []*jsonschema.ModelIR {
	var reqModels []*jsonschema.ModelIR
	seen := make(map[string]bool)

	for _, svc := range services {
		for _, op := range svc.Operations {
			if op.BodySchema == nil {
				continue
			}
			modelName := namer(svc.Name, op.Name)
			if seen[modelName] {
				continue
			}
			seen[modelName] = true

			if op.BodySchema.Type == schema.TypeArray && op.BodySchema.Items != nil {
				if op.BodySchema.Items.Type == schema.TypeObject {

					itemModelName := modelName + "Item"
					ir := SchemaToModelIR(op.BodySchema.Items, itemModelName, mapper, kindMapper)
					if ir != nil {
						reqModels = append(reqModels, ir)
					}
				}
			} else {
				ir := SchemaToModelIR(op.BodySchema, modelName, mapper, kindMapper)
				if ir != nil {
					reqModels = append(reqModels, ir)
				}
			}
		}
	}
	return reqModels
}

func SchemaToModelIR(s *schema.Schema, name string, mapper TypeMapper, kindMapper TypeKindMapper) *jsonschema.ModelIR {
	if s == nil || (s.Type != schema.TypeObject && s.Type != "") {
		return nil
	}

	ir := &jsonschema.ModelIR{
		Name: name,
	}

	var propNames []string
	for pn := range s.Properties {
		propNames = append(propNames, pn)
	}
	sort.Strings(propNames)

	for _, pn := range propNames {
		ps := s.Properties[pn]
		if ps == nil {
			continue
		}

		resolvedType := mapper(ps, name, pn)
		prop := jsonschema.PropertyIR{
			Name:     pn,
			JSONName: pn,
			Type: jsonschema.TypeRef{
				Kind:       kindMapper(ps.Type),
				Name:       resolvedType,
				IsNullable: true,
			},
		}

		if ps.Type == schema.TypeObject && len(ps.Properties) > 0 {
			nestedName := name + ToPascalCase(pn)
			nested := SchemaToModelIR(ps, nestedName, mapper, kindMapper)
			if nested != nil {
				ir.NestedTypes = append(ir.NestedTypes, *nested)
				prop.Type.Name = nestedName
				prop.Type.Kind = jsonschema.TypeObject
			}
		}

		if ps.Type == schema.TypeArray && ps.Items != nil {
			if ps.Items.Type == schema.TypeObject && len(ps.Items.Properties) > 0 {
				nestedName := name + ToPascalCase(pn) + "Item"
				nested := SchemaToModelIR(ps.Items, nestedName, mapper, kindMapper)
				if nested != nil {
					ir.NestedTypes = append(ir.NestedTypes, *nested)
				}
			}
		}

		ir.Properties = append(ir.Properties, prop)
	}

	return ir
}
