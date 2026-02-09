package schema

import (
	"sort"
	"strings"
)

type GoStruct struct {
	Name   string
	Fields []GoField
}

type GoField struct {
	Name string
	Type string
	JSON string
}

func ToGoStructs(s *Schema, rootName string) []GoStruct {
	var structs []GoStruct
	generateGoStruct(s, rootName, &structs)
	return structs
}

func generateGoStruct(s *Schema, name string, structs *[]GoStruct) {
	if s.Type != TypeObject || s.Properties == nil {
		return
	}

	var fields []GoField
	keys := s.SortedPropertyKeys()

	for _, k := range keys {
		prop := s.Properties[k]
		fieldName := toExported(k)
		typeName := goTypeFor(prop, name, fieldName, structs)

		fields = append(fields, GoField{
			Name: fieldName,
			Type: typeName,
			JSON: k,
		})
	}

	*structs = append(*structs, GoStruct{
		Name:   name,
		Fields: fields,
	})
}

func goTypeFor(s *Schema, parentName, fieldName string, structs *[]GoStruct) string {
	switch s.Type {
	case TypeBoolean:
		if s.Nullable {
			return "*bool"
		}
		return "bool"
	case TypeInteger:
		if s.Nullable {
			return "*int"
		}
		return "int"
	case TypeNumber:
		if s.Nullable {
			return "*float64"
		}
		return "float64"
	case TypeString:
		if s.Nullable {
			return "*string"
		}
		return "string"
	case TypeArray:
		if s.Items != nil {
			itemType := goTypeFor(s.Items, parentName, singularize(fieldName), structs)
			return "[]" + itemType
		}
		return "[]interface{}"
	case TypeObject:
		subName := parentName + fieldName

		if fieldName == "FacilityType" {
			subName = "FacilityType"
		}
		generateGoStruct(s, subName, structs)
		return subName
	case TypeNull:
		return "*string"
	default:
		return "interface{}"
	}
}

var singularOverrides = map[string]string{
	"CaregiversStatus":   "CaregiversStatus",
	"PatientsStatus":     "PatientsStatus",
	"PlantBatches":       "PlantBatch",
	"UnitsOfMeasure":     "UnitOfMeasure",
	"AdditivesTemplates": "AdditivesTemplate",
}

func singularize(s string) string {
	if override, ok := singularOverrides[s]; ok {
		return override
	}
	if strings.HasSuffix(s, "us") {
		return s
	}
	if strings.HasSuffix(s, "ies") {
		return strings.TrimSuffix(s, "ies") + "y"
	}
	if strings.HasSuffix(s, "ches") || strings.HasSuffix(s, "shes") || strings.HasSuffix(s, "xes") {
		return strings.TrimSuffix(s, "es")
	}
	if strings.HasSuffix(s, "s") && !strings.HasSuffix(s, "ss") {
		return strings.TrimSuffix(s, "s")
	}
	return s
}

func RenderGoCode(structs []GoStruct) string {
	var sb strings.Builder

	sortedStructs := make([]GoStruct, len(structs))
	copy(sortedStructs, structs)
	sort.Slice(sortedStructs, func(i, j int) bool {
		return sortedStructs[i].Name < sortedStructs[j].Name
	})

	for i, st := range sortedStructs {
		if i > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString("type ")
		sb.WriteString(st.Name)
		sb.WriteString(" struct {\n")

		for _, f := range st.Fields {
			sb.WriteString("\t")
			sb.WriteString(f.Name)
			sb.WriteString(" ")
			sb.WriteString(f.Type)
			sb.WriteString(" `json:\"")
			sb.WriteString(f.JSON)
			sb.WriteString("\"`\n")
		}

		sb.WriteString("}\n")
	}

	return sb.String()
}
