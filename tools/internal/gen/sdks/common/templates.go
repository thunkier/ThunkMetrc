package common

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed templates/*
var TemplatesFS embed.FS

func RenderTemplate(fs embed.FS, tmplPath, outPath string, data interface{}, funcs template.FuncMap) error {
	t, err := template.New(filepath.Base(tmplPath)).Funcs(funcs).ParseFS(fs, tmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", tmplPath, err)
	}

	f, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outPath, err)
	}
	defer f.Close()

	if err := t.Execute(f, data); err != nil {
		return fmt.Errorf("failed to execute template %s: %w", tmplPath, err)
	}
	return nil
}

func GetBaseTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"ToPascalCase":     ToPascalCase,
		"ToCamelCase":      ToCamelCase,
		"ToSnakeCase":      ToSnakeCase,
		"CleanName":        CleanName,
		"CleanDocs":        CleanDocs,
		"FormatMethodName": FormatMethodName,
		"dict": func(values ...interface{}) (map[string]interface{}, error) {
			if len(values)%2 != 0 {
				return nil, fmt.Errorf("invalid dict call")
			}
			dict := make(map[string]interface{}, len(values)/2)
			for i := 0; i < len(values); i += 2 {
				key, ok := values[i].(string)
				if !ok {
					return nil, fmt.Errorf("dict keys must be strings")
				}
				dict[key] = values[i+1]
			}
			return dict, nil
		},
		"list":           func(values ...string) []string { return values },
		"append":         func(l []string, v string) []string { return append(l, v) },
		"join":           strings.Join,
		"Split":          strings.Split,
		"TrimSuffix":     func(suffix, s string) string { return strings.TrimSuffix(s, suffix) },
		"ToLower":        strings.ToLower,
		"Contains":       strings.Contains,
		"GetPathParams":  GetPathParams,
		"ParseParamDocs": ParseParamDocs,
		"default": func(def, val interface{}) interface{} {
			str, ok := val.(string)
			if ok && str == "" {
				return def
			}
			if val == nil {
				return def
			}
			return val
		},
	}
}

type PomConfig struct {
	ArtifactId       string
	Name             string
	Description      string
	Version          string
	IsKotlin         bool
	IsWrapper        bool
	ClientArtifactId string
	Deps             map[string]string
}

func RenderPom(outPath string, config PomConfig) error {
	return RenderTemplate(TemplatesFS, "templates/pom.xml.tmpl", outPath, config, GetBaseTemplateFuncs())
}

type PackageConfig struct {
	Version     string
	Description string
	IsWrapper   bool
	Deps        map[string]string
}

func RenderPackageJSON(outPath string, config PackageConfig) error {
	return RenderTemplate(TemplatesFS, "templates/package.json.tmpl", outPath, config, GetBaseTemplateFuncs())
}

type CargoConfig struct {
	Version     string
	Description string
	IsWrapper   bool
	Deps        map[string]string
}

func RenderCargoToml(outPath string, config CargoConfig) error {
	return RenderTemplate(TemplatesFS, "templates/Cargo.toml.tmpl", outPath, config, GetBaseTemplateFuncs())
}

type PyProjectConfig struct {
	Version     string
	Description string
	IsWrapper   bool
	Deps        map[string]string
}

func RenderPyProjectToml(outPath string, config PyProjectConfig) error {
	return RenderTemplate(TemplatesFS, "templates/pyproject.toml.tmpl", outPath, config, GetBaseTemplateFuncs())
}
