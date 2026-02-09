package common

import (
	"regexp"
	"strings"
	"unicode"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	str = strings.ReplaceAll(str, " ", "_")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	re := regexp.MustCompile(`_+`)
	snake = re.ReplaceAllString(snake, "_")
	return strings.ToLower(snake)
}

func ToCamelCase(str string) string {
	s := ToPascalCase(str)
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func ToPascalCase(str string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	parts := reg.Split(str, -1)

	var sb strings.Builder
	for _, part := range parts {
		if part == "" {
			continue
		}
		runes := []rune(part)
		runes[0] = unicode.ToUpper(runes[0])
		sb.WriteString(string(runes))
	}
	return sb.String()
}

func CleanName(name string) string {
	return strings.TrimSpace(name)
}

func CleanDocs(docs string) string {
	if docs == "" {
		return ""
	}
	docs = strings.ReplaceAll(docs, "\r\n", "\n")
	docs = strings.ReplaceAll(docs, "\r", "\n")
	lines := strings.Split(docs, "\n")
	var cleaned []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		lower := strings.ToLower(trimmed)
		if strings.Contains(lower, "example response") || strings.Contains(lower, "parameters:") {
			break
		}

		if trimmed != "" {
			cleaned = append(cleaned, trimmed)
		}
	}
	return strings.Join(cleaned, "\n")
}

func ParseParamDocs(docs string) map[string]string {
	paramDocs := make(map[string]string)
	if docs == "" {
		return paramDocs
	}
	lines := strings.Split(docs, "\n")
	inParams := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(strings.ToLower(trimmed), "parameters:") {
			inParams = true
			continue
		}
		if !inParams {
			continue
		}
		if trimmed == "" || strings.Contains(strings.ToLower(trimmed), "example response") {
			break
		}

		if strings.HasPrefix(trimmed, "-") {
			trimmed = strings.TrimPrefix(trimmed, "-")
			trimmed = strings.TrimSpace(trimmed)
			parts := strings.SplitN(trimmed, ":", 2)
			if len(parts) == 2 {
				namePart := parts[0]
				descPart := strings.TrimSpace(parts[1])

				nameParts := strings.Fields(namePart)
				if len(nameParts) > 0 {
					paramName := nameParts[0]
					paramDocs[paramName] = descPart
				}
			}
		}
	}
	return paramDocs
}

func GetPathParams(url string) []string {
	re := regexp.MustCompile(`\{([a-zA-Z0-9_]+)\}`)
	matches := re.FindAllStringSubmatch(url, -1)
	var params []string
	for _, m := range matches {
		if len(m) > 1 {
			params = append(params, m[1])
		}
	}
	return params
}

func FormatMethodName(opName, group string) string {
	verbs := []string{"Get", "Create", "Update", "Delete", "Post", "Put", "Patch"}
	for _, verb := range verbs {
		if strings.HasPrefix(opName, verb) {
			rest := strings.TrimPrefix(opName, verb)
			if rest == "" || rest == "Active" {
				return verb + group
			}
			if strings.Contains(rest, group) {
				return opName
			}
			return verb + group + rest
		}
	}
	return opName + group
}
