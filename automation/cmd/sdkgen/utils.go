package main

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func cleanName(s string) string {
	s = regexp.MustCompile(`(?i)\((GET|POST|PUT|DELETE|PATCH)\)`).ReplaceAllString(s, "")

	// Split camelCase
	s = regexp.MustCompile(`([a-z0-9])([A-Z])`).ReplaceAllString(s, "${1} ${2}")

	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "/", " ")
	s = strings.ReplaceAll(s, "_", " ")

	s = regexp.MustCompile(`\s+\d+$`).ReplaceAllString(s, "")

	s = strings.TrimSpace(s)
	s = strings.Join(strings.Fields(s), " ")

	return s
}

func ToPascalCase(s string) string {
	if strings.Contains(s, " ") {
		parts := strings.Fields(s)
		for i, p := range parts {
			parts[i] = casesTitle(p)
		}
		return strings.Join(parts, "")
	}
	return casesTitle(s)
}

func ToCamelCase(s string) string {
	pascal := ToPascalCase(s)
	if pascal == "" {
		return ""
	}
	r := []rune(pascal)
	r[0] = []rune(strings.ToLower(string(r[0])))[0]
	return string(r)
}

func ToSnakeCase(s string) string {
	if strings.Contains(s, " ") {
		return strings.ToLower(strings.Join(strings.Fields(s), "_"))
	}

	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func casesTitle(s string) string {
	c := cases.Title(language.English)
	return c.String(s)
}

// GetPathParams extracts {param} from URL
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

func cleanDocs(docs string) string {
	// Remove "Example Response" section and everything after
	lower := strings.ToLower(docs)
	idx := strings.Index(lower, "example response")
	if idx != -1 {
		docs = docs[:idx]
	}
	return strings.TrimSpace(docs)
}
