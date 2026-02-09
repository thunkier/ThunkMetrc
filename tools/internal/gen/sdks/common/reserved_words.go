package common

import "strings"

var ReservedWords = map[string]map[string]bool{
	"java": {
		"package": true, "class": true, "interface": true, "extends": true,
		"implements": true, "import": true, "new": true, "default": true,
		"abstract": true, "final": true, "static": true, "public": true,
		"private": true, "protected": true, "void": true, "return": true,
		"if": true, "else": true, "for": true, "while": true, "do": true,
		"switch": true, "case": true, "break": true, "continue": true,
		"try": true, "catch": true, "finally": true, "throw": true,
		"throws": true, "synchronized": true, "volatile": true,
		"transient": true, "native": true, "strictfp": true,
		"boolean": true, "byte": true, "char": true, "short": true,
		"int": true, "long": true, "float": true, "double": true,
		"null": true, "true": true, "false": true, "this": true, "super": true,
		"enum": true, "assert": true, "const": true, "goto": true,
	},
	"kotlin": {
		"package": true, "class": true, "interface": true, "object": true,
		"fun": true, "val": true, "var": true, "when": true, "is": true,
		"as": true, "in": true, "out": true, "if": true, "else": true,
		"for": true, "while": true, "do": true, "return": true, "break": true,
		"continue": true, "throw": true, "try": true, "catch": true,
		"finally": true, "null": true, "true": true, "false": true,
		"this": true, "super": true, "typeof": true, "typealias": true,
		"import": true, "companion": true, "data": true, "sealed": true,
		"annotation": true, "enum": true, "inner": true, "open": true,
		"override": true, "private": true, "protected": true, "public": true,
		"internal": true, "abstract": true, "final": true, "suspend": true,
	},
	"rust": {
		"type": true, "impl": true, "mod": true, "use": true, "fn": true,
		"pub": true, "let": true, "mut": true, "const": true, "static": true,
		"struct": true, "enum": true, "trait": true, "where": true,
		"self": true, "Self": true, "super": true, "crate": true,
		"if": true, "else": true, "match": true, "for": true, "while": true,
		"loop": true, "break": true, "continue": true, "return": true,
		"as": true, "in": true, "ref": true, "move": true, "box": true,
		"dyn": true, "async": true, "await": true, "unsafe": true,
		"extern": true, "true": true, "false": true,
	},
	"typescript": {
		"package": true, "class": true, "interface": true, "extends": true,
		"implements": true, "import": true, "export": true, "default": true,
		"new": true, "function": true, "const": true, "let": true, "var": true,
		"if": true, "else": true, "for": true, "while": true, "do": true,
		"switch": true, "case": true, "break": true, "continue": true,
		"return": true, "try": true, "catch": true, "finally": true,
		"throw": true, "typeof": true, "instanceof": true, "in": true,
		"of": true, "async": true, "await": true, "yield": true,
		"null": true, "undefined": true, "true": true, "false": true,
		"this": true, "super": true, "type": true, "enum": true,
		"public": true, "private": true, "protected": true, "static": true,
		"readonly": true, "abstract": true, "as": true, "is": true,
		"keyof": true, "never": true, "unknown": true, "any": true,
	},
	"go": {
		"break": true, "case": true, "chan": true, "const": true,
		"continue": true, "default": true, "defer": true, "else": true,
		"fallthrough": true, "for": true, "func": true, "go": true,
		"goto": true, "if": true, "import": true, "interface": true,
		"map": true, "package": true, "range": true, "return": true,
		"select": true, "struct": true, "switch": true, "type": true,
		"var": true,
	},
	"csharp": {
		"abstract": true, "as": true, "base": true, "bool": true,
		"break": true, "byte": true, "case": true, "catch": true,
		"char": true, "checked": true, "class": true, "const": true,
		"continue": true, "decimal": true, "default": true, "delegate": true,
		"do": true, "double": true, "else": true, "enum": true,
		"event": true, "explicit": true, "extern": true, "false": true,
		"finally": true, "fixed": true, "float": true, "for": true,
		"foreach": true, "goto": true, "if": true, "implicit": true,
		"in": true, "int": true, "interface": true, "internal": true,
		"is": true, "lock": true, "long": true, "namespace": true,
		"new": true, "null": true, "object": true, "operator": true,
		"out": true, "override": true, "params": true, "private": true,
		"protected": true, "public": true, "readonly": true, "ref": true,
		"return": true, "sbyte": true, "sealed": true, "short": true,
		"sizeof": true, "stackalloc": true, "static": true, "string": true,
		"struct": true, "switch": true, "this": true, "throw": true,
		"true": true, "try": true, "typeof": true, "uint": true,
		"ulong": true, "unchecked": true, "unsafe": true, "ushort": true,
		"using": true, "virtual": true, "void": true, "volatile": true,
		"while": true,
	},
	"python": {
		"False": true, "None": true, "True": true, "and": true,
		"as": true, "assert": true, "async": true, "await": true,
		"break": true, "class": true, "continue": true, "def": true,
		"del": true, "elif": true, "else": true, "except": true,
		"finally": true, "for": true, "from": true, "global": true,
		"if": true, "import": true, "in": true, "is": true,
		"lambda": true, "nonlocal": true, "not": true, "or": true,
		"pass": true, "raise": true, "return": true, "try": true,
		"while": true, "with": true, "yield": true,
	},
}

func SafeFieldName(name, lang string) string {
	reserved, ok := ReservedWords[lang]
	if !ok {
		return name
	}

	lower := strings.ToLower(name)
	if reserved[lower] || reserved[name] {
		switch lang {
		case "java", "kotlin":
			return name + "_"
		case "rust":
			return "r#" + name
		case "typescript":
			return "'" + name + "'"
		case "go":
			return name + "_"
		case "csharp":
			return "@" + name
		case "python":
			return name + "_"
		default:
			return name + "_"
		}
	}
	return name
}
