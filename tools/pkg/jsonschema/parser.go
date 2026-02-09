package jsonschema

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/thunkier/thunkmetrc/tools/pkg/schema"
)

type ModelIR struct {
	Name        string
	Properties  []PropertyIR
	NestedTypes []ModelIR
}

type PropertyIR struct {
	Name       string
	JSONName   string
	Type       TypeRef
	IsRequired bool
}

type TypeRef struct {
	Kind       TypeKind
	Name       string
	ItemType   *TypeRef
	IsNullable bool
}

type TypeKind string

const (
	TypeString  TypeKind = "string"
	TypeInteger TypeKind = "integer"
	TypeNumber  TypeKind = "number"
	TypeBoolean TypeKind = "boolean"
	TypeArray   TypeKind = "array"
	TypeObject  TypeKind = "object"
	TypeAny     TypeKind = "any"
)

// JSONSchema usage replaced by schema.JSONSchema

func ParseSchemaFile(path string) (*ModelIR, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file: %w", err)
	}

	var schemaObj schema.JSONSchema
	if err := json.Unmarshal(data, &schemaObj); err != nil {
		return nil, fmt.Errorf("failed to parse JSON schema: %w", err)
	}

	modelName := schemaObj.Title
	if modelName == "" {
		modelName = "Unknown"
	}

	return parseSchema(modelName, &schemaObj), nil
}

func ParseSchemaData(data []byte, defaultName string) (*ModelIR, error) {
	var schemaObj schema.JSONSchema
	if err := json.Unmarshal(data, &schemaObj); err != nil {
		return nil, fmt.Errorf("failed to parse JSON schema: %w", err)
	}

	modelName := schemaObj.Title
	if modelName == "" {
		modelName = defaultName
	}

	return parseSchema(modelName, &schemaObj), nil
}

func parseSchema(name string, s *schema.JSONSchema) *ModelIR {
	model := &ModelIR{
		Name: name,
	}

	if s.Properties == nil {
		return model
	}

	requiredMap := make(map[string]bool)
	for _, r := range s.Required {
		requiredMap[r] = true
	}

	var propNames []string
	for propName := range s.Properties {
		propNames = append(propNames, propName)
	}
	sort.Strings(propNames)

	for _, propName := range propNames {
		propSchema := s.Properties[propName]
		typeRef := resolveType(propSchema, name+propName)

		prop := PropertyIR{
			Name:       propName,
			JSONName:   propName,
			Type:       typeRef,
			IsRequired: requiredMap[propName],
		}
		model.Properties = append(model.Properties, prop)

		if typeRef.Kind == TypeObject && propSchema.Properties != nil {
			nestedModel := parseSchema(name+propName, propSchema)
			model.NestedTypes = append(model.NestedTypes, *nestedModel)
		}

		if typeRef.Kind == TypeArray && typeRef.ItemType != nil && typeRef.ItemType.Kind == TypeObject {
			if propSchema.Items != nil && propSchema.Items.Properties != nil {
				nestedModel := parseSchema(name+propName+"Item", propSchema.Items)
				model.NestedTypes = append(model.NestedTypes, *nestedModel)
			}
		}
	}

	return model
}

func resolveType(s *schema.JSONSchema, nestedTypeName string) TypeRef {
	if s == nil {
		return TypeRef{Kind: TypeAny}
	}

	var typeStr string
	var isNullable bool

	switch t := s.Type.(type) {
	case string:
		typeStr = t
	case []interface{}:
		for _, v := range t {
			if s, ok := v.(string); ok {
				if s == "null" {
					isNullable = true
				} else {
					typeStr = s
				}
			}
		}
	}

	ref := TypeRef{IsNullable: isNullable}

	switch typeStr {
	case "string":
		ref.Kind = TypeString
	case "integer":
		ref.Kind = TypeInteger
	case "number":
		ref.Kind = TypeNumber
	case "boolean":
		ref.Kind = TypeBoolean
	case "array":
		ref.Kind = TypeArray
		if s.Items != nil {
			itemType := resolveType(s.Items, nestedTypeName+"Item")
			ref.ItemType = &itemType
		} else {
			ref.ItemType = &TypeRef{Kind: TypeAny}
		}
	case "object":
		ref.Kind = TypeObject
		ref.Name = nestedTypeName
	default:
		ref.Kind = TypeAny
	}

	return ref
}

func (t TypeRef) ToCSharp() string {
	suffix := ""
	if t.IsNullable {
		suffix = "?"
	}

	switch t.Kind {
	case TypeString:
		return "string" + suffix
	case TypeInteger:
		return "int" + suffix
	case TypeNumber:
		return "double" + suffix
	case TypeBoolean:
		return "bool" + suffix
	case TypeArray:
		if t.ItemType != nil {
			return fmt.Sprintf("List<%s>%s", t.ItemType.ToCSharp(), suffix)
		}
		return "List<object>" + suffix
	case TypeObject:
		return t.Name + suffix
	default:
		return "object" + suffix
	}
}

func (t TypeRef) ToGo() string {
	prefix := ""
	if t.IsNullable {
		prefix = "*"
	}

	switch t.Kind {
	case TypeString:
		if t.IsNullable {
			return "*string"
		}
		return "string"
	case TypeInteger:
		if t.IsNullable {
			return "*int64"
		}
		return "int64"
	case TypeNumber:
		if t.IsNullable {
			return "*float64"
		}
		return "float64"
	case TypeBoolean:
		if t.IsNullable {
			return "*bool"
		}
		return "bool"
	case TypeArray:
		if t.ItemType != nil {
			return "[]" + t.ItemType.ToGo()
		}
		return "[]interface{}"
	case TypeObject:
		return prefix + t.Name
	default:
		return "interface{}"
	}
}

func (t TypeRef) ToPython() string {
	inner := ""
	switch t.Kind {
	case TypeString:
		inner = "str"
	case TypeInteger:
		inner = "int"
	case TypeNumber:
		inner = "float"
	case TypeBoolean:
		inner = "bool"
	case TypeArray:
		if t.ItemType != nil {
			inner = "List[" + t.ItemType.ToPython() + "]"
		} else {
			inner = "List[Any]"
		}
	case TypeObject:
		inner = t.Name
	default:
		inner = "Any"
	}

	if t.IsNullable {
		return "Optional[" + inner + "]"
	}
	return inner
}

func (t TypeRef) ToTypeScript() string {
	inner := ""
	switch t.Kind {
	case TypeString:
		inner = "string"
	case TypeInteger, TypeNumber:
		inner = "number"
	case TypeBoolean:
		inner = "boolean"
	case TypeArray:
		if t.ItemType != nil {
			inner = t.ItemType.ToTypeScript() + "[]"
		} else {
			inner = "any[]"
		}
	case TypeObject:
		inner = t.Name
	default:
		inner = "any"
	}

	if t.IsNullable {
		return inner + " | null"
	}
	return inner
}

func (t TypeRef) ToJava() string {
	switch t.Kind {
	case TypeString:
		return "String"
	case TypeInteger:
		if t.IsNullable {
			return "Integer"
		}
		return "int"
	case TypeNumber:
		if t.IsNullable {
			return "Double"
		}
		return "double"
	case TypeBoolean:
		if t.IsNullable {
			return "Boolean"
		}
		return "boolean"
	case TypeArray:
		if t.ItemType != nil {
			return fmt.Sprintf("List<%s>", t.ItemType.ToJava())
		}
		return "List<Object>"
	case TypeObject:
		return t.Name
	default:
		return "Object"
	}
}

func (t TypeRef) ToKotlin() string {
	suffix := ""
	if t.IsNullable {
		suffix = "?"
	}

	switch t.Kind {
	case TypeString:
		return "String" + suffix
	case TypeInteger:
		return "Int" + suffix
	case TypeNumber:
		return "Double" + suffix
	case TypeBoolean:
		return "Boolean" + suffix
	case TypeArray:
		if t.ItemType != nil {
			inner := t.ItemType.ToKotlin()
			inner = strings.TrimSuffix(inner, "?")
			return fmt.Sprintf("List<%s>%s", inner, suffix)
		}
		return "List<Any>" + suffix
	case TypeObject:
		return t.Name + suffix
	default:
		return "Any" + suffix
	}
}

func (t TypeRef) ToRust() string {
	inner := ""
	switch t.Kind {
	case TypeString:
		inner = "String"
	case TypeInteger:
		inner = "i64"
	case TypeNumber:
		inner = "f64"
	case TypeBoolean:
		inner = "bool"
	case TypeArray:
		if t.ItemType != nil {
			inner = fmt.Sprintf("Vec<%s>", t.ItemType.ToRust())
		} else {
			inner = "Vec<serde_json::Value>"
		}
	case TypeObject:
		inner = t.Name
	default:
		inner = "serde_json::Value"
	}

	if t.IsNullable {
		return "Option<" + inner + ">"
	}
	return inner
}
