package common

var PrimitiveTypeMap = map[string]map[string]string{
	"string": {
		"csharp":     "string",
		"java":       "String",
		"kotlin":     "String",
		"go":         "string",
		"python":     "str",
		"rust":       "String",
		"typescript": "string",
	},
	"integer": {
		"csharp":     "long?",
		"java":       "Long",
		"kotlin":     "Long?",
		"go":         "*int64",
		"python":     "int",
		"rust":       "Option<i64>",
		"typescript": "number",
	},
	"number": {
		"csharp":     "double?",
		"java":       "Double",
		"kotlin":     "Double?",
		"go":         "*float64",
		"python":     "float",
		"rust":       "Option<f64>",
		"typescript": "number",
	},
	"boolean": {
		"csharp":     "bool?",
		"java":       "Boolean",
		"kotlin":     "Boolean?",
		"go":         "*bool",
		"python":     "bool",
		"rust":       "Option<bool>",
		"typescript": "boolean",
	},
	"null": {
		"csharp":     "object",
		"java":       "Object",
		"kotlin":     "Any?",
		"go":         "interface{}",
		"python":     "None",
		"rust":       "()",
		"typescript": "null",
	},
}

var ArrayTypeFormat = map[string]string{
	"csharp":     "List<%s>",
	"java":       "List<%s>",
	"kotlin":     "List<%s>",
	"go":         "[]%s",
	"python":     "List[%s]",
	"rust":       "Vec<%s>",
	"typescript": "%s[]",
}

var ObjectType = map[string]string{
	"csharp":     "object",
	"java":       "Object",
	"kotlin":     "Any",
	"go":         "map[string]interface{}",
	"python":     "dict",
	"rust":       "serde_json::Value",
	"typescript": "Record<string, unknown>",
}

func GetPrimitiveType(jsonType, lang string) string {
	if m, ok := PrimitiveTypeMap[jsonType]; ok {
		if t, ok := m[lang]; ok {
			return t
		}
	}
	return ObjectType[lang]
}

func GetArrayTypeFormat(lang string) string {
	if f, ok := ArrayTypeFormat[lang]; ok {
		return f
	}
	return "List<%s>"
}

func GetObjectType(lang string) string {
	if t, ok := ObjectType[lang]; ok {
		return t
	}
	return "object"
}
