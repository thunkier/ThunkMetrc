package schema

import (
	"encoding/json"
	"sort"
)

type JSONSchema struct {
	Schema      string                 `json:"$schema,omitempty"`
	ID          string                 `json:"$id,omitempty"`
	Title       string                 `json:"title,omitempty"`
	Description string                 `json:"description,omitempty"`
	Type        interface{}            `json:"type,omitempty"`
	Properties  map[string]*JSONSchema `json:"properties,omitempty"`
	Items       *JSONSchema            `json:"items,omitempty"`
	Required    []string               `json:"required,omitempty"`
	Format      string                 `json:"format,omitempty"`
	Definitions map[string]*JSONSchema `json:"definitions,omitempty"`
	Ref         string                 `json:"$ref,omitempty"`
}

func ToJSONSchema(s *Schema, title string) *JSONSchema {
	js := &JSONSchema{
		Schema: "http://json-schema.org/draft-07/schema#",
		Title:  title,
	}

	convertSchemaInto(s, js)
	return js
}

func convertSchemaInto(s *Schema, js *JSONSchema) {
	if s.Nullable {
		js.Type = []string{string(s.Type), "null"}
	} else {
		js.Type = string(s.Type)
	}

	if s.Format != "" {
		js.Format = s.Format
	}

	if s.Description != "" {
		js.Description = s.Description
	}

	if s.Type == TypeObject && s.Properties != nil {
		js.Properties = make(map[string]*JSONSchema)
		var required []string

		keys := s.SortedPropertyKeys()
		for _, k := range keys {
			prop := s.Properties[k]
			propJS := &JSONSchema{}
			convertSchemaInto(prop, propJS)
			js.Properties[k] = propJS

			if !prop.Nullable {
				required = append(required, k)
			}
		}

		if len(required) > 0 {
			sort.Strings(required)
			js.Required = required
		}
	}

	if s.Type == TypeArray && s.Items != nil {
		js.Items = &JSONSchema{}
		convertSchemaInto(s.Items, js.Items)
	}
}

func ToJSONSchemaBytes(s *Schema, title string) ([]byte, error) {
	js := ToJSONSchema(s, title)
	return json.MarshalIndent(js, "", "  ")
}

func ToJSONSchemaMap(s *Schema) map[string]interface{} {
	result := make(map[string]interface{})

	if s.Nullable {
		result["type"] = []string{string(s.Type), "null"}
	} else {
		result["type"] = string(s.Type)
	}

	if s.Format != "" {
		result["format"] = s.Format
	}

	if s.Type == TypeObject && s.Properties != nil {
		props := make(map[string]interface{})
		var required []string

		for k, prop := range s.Properties {
			props[k] = ToJSONSchemaMap(prop)
			if !prop.Nullable {
				required = append(required, k)
			}
		}

		result["properties"] = props
		if len(required) > 0 {
			sort.Strings(required)
			result["required"] = required
		}
	}

	if s.Type == TypeArray && s.Items != nil {
		result["items"] = ToJSONSchemaMap(s.Items)
	}

	return result
}

func (js *JSONSchema) ToSchema() *Schema {
	s := &Schema{}

	switch v := js.Type.(type) {
	case string:
		s.Type = SchemaType(v)
	case []interface{}:
		for _, t := range v {
			if str, ok := t.(string); ok {
				if str == "null" {
					s.Nullable = true
				} else {
					s.Type = SchemaType(str)
				}
			}
		}
	}

	if js.Format != "" {
		s.Format = js.Format
	}
	if js.Description != "" {
		s.Description = js.Description
	}

	if s.Type == TypeObject && js.Properties != nil {
		s.Properties = make(map[string]*Schema)
		for k, prop := range js.Properties {
			s.Properties[k] = prop.ToSchema()
		}
	}

	if s.Type == TypeArray && js.Items != nil {
		s.Items = js.Items.ToSchema()
	}

	return s
}
