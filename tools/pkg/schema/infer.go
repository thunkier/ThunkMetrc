package schema

import (
	"sort"
	"strings"
	"unicode"
)

type SchemaType string

const (
	TypeString  SchemaType = "string"
	TypeInteger SchemaType = "integer"
	TypeNumber  SchemaType = "number"
	TypeBoolean SchemaType = "boolean"
	TypeObject  SchemaType = "object"
	TypeArray   SchemaType = "array"
	TypeNull    SchemaType = "null"
)

type Schema struct {
	Type        SchemaType
	Properties  map[string]*Schema
	Items       *Schema
	Nullable    bool
	Format      string
	Description string
}

func InferFromJSON(data interface{}) *Schema {
	if data == nil {
		return &Schema{Type: TypeNull, Nullable: true}
	}

	switch v := data.(type) {
	case bool:
		return &Schema{Type: TypeBoolean}
	case float64:
		if v == float64(int(v)) {
			return &Schema{Type: TypeInteger}
		}
		return &Schema{Type: TypeNumber}
	case string:
		s := &Schema{Type: TypeString}
		if looksLikeDateTime(v) {
			s.Format = "date-time"
		} else if looksLikeDate(v) {
			s.Format = "date"
		}
		return s
	case []interface{}:
		s := &Schema{Type: TypeArray}
		if len(v) > 0 {
			s.Items = InferFromJSON(v[0])
		}
		return s
	case map[string]interface{}:
		return inferObject(v)
	default:
		return &Schema{Type: TypeString}
	}
}

func inferObject(obj map[string]interface{}) *Schema {
	s := &Schema{
		Type:       TypeObject,
		Properties: make(map[string]*Schema),
	}

	for k, v := range obj {
		prop := InferFromJSON(v)
		if v == nil {
			prop = inferNullableType(k)
		} else {
			if prop.Type == TypeInteger && shouldBeNumber(k) {
				prop.Type = TypeNumber
			}
		}
		s.Properties[k] = prop
	}

	return s
}

func shouldBeNumber(fieldName string) bool {
	upper := toExported(fieldName)
	return strings.HasSuffix(upper, "Weight") ||
		strings.HasSuffix(upper, "Volume") ||
		strings.HasSuffix(upper, "Content") ||
		strings.HasSuffix(upper, "Percent") ||
		strings.HasSuffix(upper, "Dose") ||
		strings.HasSuffix(upper, "Quantity")
}

func inferNullableType(fieldName string) *Schema {
	upper := toExported(fieldName)

	if strings.HasSuffix(upper, "Id") ||
		strings.HasSuffix(upper, "Count") ||
		strings.HasSuffix(upper, "Hours") ||
		strings.HasSuffix(upper, "Days") ||
		strings.HasSuffix(upper, "Allowed") {
		return &Schema{Type: TypeInteger, Nullable: true}
	}

	if strings.HasSuffix(upper, "Weight") ||
		strings.HasSuffix(upper, "Volume") ||
		strings.HasSuffix(upper, "Content") ||
		strings.HasSuffix(upper, "Percent") ||
		strings.HasSuffix(upper, "Dose") ||
		strings.HasSuffix(upper, "Quantity") {
		return &Schema{Type: TypeNumber, Nullable: true}
	}

	if strings.HasPrefix(upper, "Allow") ||
		strings.HasPrefix(upper, "Can") ||
		strings.HasPrefix(upper, "Is") ||
		strings.HasPrefix(upper, "Has") ||
		strings.HasPrefix(upper, "Enable") ||
		strings.HasPrefix(upper, "Require") {
		return &Schema{Type: TypeBoolean, Nullable: true}
	}

	return &Schema{Type: TypeString, Nullable: true}
}

func looksLikeDateTime(s string) bool {

	if len(s) >= 19 && s[4] == '-' && s[7] == '-' && (s[10] == 'T' || s[10] == ' ') {
		return true
	}
	return false
}

func looksLikeDate(s string) bool {

	if len(s) == 10 && s[4] == '-' && s[7] == '-' {
		return true
	}
	return false
}

func toExported(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func (s *Schema) SortedPropertyKeys() []string {
	if s.Properties == nil {
		return nil
	}
	keys := make([]string, 0, len(s.Properties))
	for k := range s.Properties {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (s *Schema) IsPaginated() bool {
	if s.Type != TypeObject || s.Properties == nil {
		return false
	}
	_, hasData := s.Properties["Data"]
	_, hasTotal := s.Properties["Total"]
	_, hasPage := s.Properties["Page"]
	return hasData && (hasTotal || hasPage)
}

func (s *Schema) UnwrapPaginated() *Schema {
	if !s.IsPaginated() {
		return s
	}
	if data, ok := s.Properties["Data"]; ok {
		if data.Type == TypeArray && data.Items != nil {
			return data.Items
		}
	}
	return s
}

func (s *Schema) UnwrapArray() *Schema {
	if s.Type == TypeArray && s.Items != nil {
		return s.Items
	}
	return s
}

func (s *Schema) IsWriteResponse() bool {
	if s.Type != TypeObject || s.Properties == nil {
		return false
	}

	ids, hasIds := s.Properties["Ids"]
	if !hasIds {
		return false
	}

	if ids.Type != TypeArray || ids.Items == nil || ids.Items.Type != TypeInteger {
		return false
	}

	if len(s.Properties) > 2 {
		return false
	}

	if len(s.Properties) == 2 {
		_, hasWarnings := s.Properties["Warnings"]
		return hasWarnings
	}

	return true
}

func (s *Schema) GetPaginationFields() []string {
	fields := []string{"Data", "CurrentPage", "Page", "PageSize", "Total", "TotalPages", "TotalRecords", "RecordsOnPage"}
	var present []string
	for _, f := range fields {
		if _, ok := s.Properties[f]; ok {
			present = append(present, f)
		}
	}
	return present
}

func (s *Schema) Signature() string {
	if s.Type != TypeObject || s.Properties == nil {
		return string(s.Type)
	}
	keys := s.SortedPropertyKeys()
	return strings.Join(keys, ",")
}

func (s *Schema) PropertyCount() int {
	if s.Properties == nil {
		return 0
	}
	return len(s.Properties)
}

func (s *Schema) IsEmpty() bool {
	switch s.Type {
	case TypeObject:
		if len(s.Properties) == 0 {
			return true
		}
		if s.IsPaginated() {
			data, hasData := s.Properties["Data"]
			if hasData && data != nil && data.Type == TypeArray {
				if data.Items == nil || data.Items.IsEmpty() {
					return true
				}
			}
		}
		return false
	case TypeArray:
		return s.Items == nil || s.Items.IsEmpty()
	case TypeString, TypeInteger, TypeNumber, TypeBoolean, TypeNull:
		return true
	default:
		return false
	}
}

func (s *Schema) Clone() *Schema {
	if s == nil {
		return nil
	}
	newSchema := &Schema{
		Type:        s.Type,
		Nullable:    s.Nullable,
		Format:      s.Format,
		Description: s.Description,
	}
	if s.Properties != nil {
		newSchema.Properties = make(map[string]*Schema, len(s.Properties))
		for k, v := range s.Properties {
			newSchema.Properties[k] = v.Clone()
		}
	}
	if s.Items != nil {
		newSchema.Items = s.Items.Clone()
	}
	return newSchema
}
