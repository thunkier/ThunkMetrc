package main

import (
	"encoding/json"
	"strings"

	"github.com/thunkmetrc/automation/pkg/bruno"
)

func InferSchema(jsonBody string) (*bruno.Schema, error) {
	var data interface{}
	if err := json.Unmarshal([]byte(jsonBody), &data); err != nil {
		return nil, err
	}
	return inferType("Root", data), nil
}

func inferType(fieldName string, data interface{}) *bruno.Schema {
	if data == nil {
		return inferNullType(fieldName)
	}

	switch v := data.(type) {
	case string:
		return &bruno.Schema{Type: bruno.TypeString}
	case bool:
		return &bruno.Schema{Type: bruno.TypeBool}
	case float64:
		if isFloatField(fieldName) {
			return &bruno.Schema{Type: bruno.TypeFloat}
		}
		if v == float64(int(v)) {
			return &bruno.Schema{Type: bruno.TypeInt}
		}
		return &bruno.Schema{Type: bruno.TypeFloat}
	case []interface{}:
		s := &bruno.Schema{Type: bruno.TypeArray}
		if len(v) > 0 {
			s.ItemType = inferType(fieldName, v[0])
		} else {
			s.ItemType = &bruno.Schema{Type: bruno.TypeUnknown}
		}
		return s
	case map[string]interface{}:
		s := &bruno.Schema{
			Type:       bruno.TypeObject,
			Properties: make(map[string]*bruno.Schema),
		}
		for k, val := range v {
			s.Properties[k] = inferType(k, val)
		}
		return s
	default:
		return &bruno.Schema{Type: bruno.TypeUnknown}
	}
}

func isFloatField(fieldName string) bool {
	return strings.HasSuffix(fieldName, "Percentage") ||
		strings.HasSuffix(fieldName, "Percent") ||
		strings.HasSuffix(fieldName, "Level") ||
		strings.HasSuffix(fieldName, "Content") ||
		strings.HasSuffix(fieldName, "Volume") ||
		strings.HasSuffix(fieldName, "Weight") ||
		strings.HasSuffix(fieldName, "Dose")
}

func inferNullType(fieldName string) *bruno.Schema {
	if strings.HasSuffix(fieldName, "Hours") ||
		strings.HasSuffix(fieldName, "Days") ||
		strings.HasSuffix(fieldName, "Allowed") ||
		strings.HasSuffix(fieldName, "Count") ||
		strings.HasSuffix(fieldName, "Id") {
		return &bruno.Schema{Type: bruno.TypeInt, IsNullable: true}
	}
	if isFloatField(fieldName) {
		return &bruno.Schema{Type: bruno.TypeFloat, IsNullable: true}
	}
	if strings.HasPrefix(fieldName, "Allow") ||
		strings.HasPrefix(fieldName, "Is") ||
		strings.HasPrefix(fieldName, "Can") ||
		strings.HasPrefix(fieldName, "Enable") ||
		strings.HasPrefix(fieldName, "Require") {
		return &bruno.Schema{Type: bruno.TypeBool, IsNullable: true}
	}
	return &bruno.Schema{Type: bruno.TypeString, IsNullable: true}
}
