package csharp

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	app "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
)

//go:embed templates/main/*.tmpl
var mainTemplatesFS embed.FS

type MainGenerator struct{}

func (g *MainGenerator) Name() string {
	return "CSharpMain"
}

func (g *MainGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-cs", "ThunkMetrc")

	extDir := filepath.Join(outDir, "Extensions")
	if err := os.RemoveAll(extDir); err != nil {
		return fmt.Errorf("failed to clean extensions dir: %w", err)
	}
	if err := os.MkdirAll(extDir, 0755); err != nil {
		return fmt.Errorf("failed to create extensions dir: %w", err)
	}

	funcs := app.GetBaseTemplateFuncs()
	funcs["ResolveType"] = resolveCSharpType
	funcs["GetIteratorItemType"] = func(op models.Operation) string {
		if op.ResponseType != "" {
			return op.ResponseType
		}
		return "object"
	}
	funcs["PaginationSignature"] = paginationSignatureArgs
	funcs["PaginationClientCall"] = paginationClientCallArgs
	funcs["PaginationParamsDocs"] = mainPaginationParamsDocs
	funcs["FormatMethodName"] = app.FormatMethodName
	funcs["HasPagination"] = hasPagination
	funcs["Replace"] = strings.ReplaceAll
	funcs["CleanDocs"] = app.CleanDocs
	funcs["Split"] = func(s, sep string) []string { return strings.Split(s, sep) }
	funcs["IsTimeWindowed"] = func(op models.Operation) bool {
		hasStart := false
		hasEnd := false
		for _, p := range op.QueryParams {
			if p.Name == "lastModifiedStart" {
				hasStart = true
			}
			if p.Name == "lastModifiedEnd" {
				hasEnd = true
			}
		}
		return hasStart && hasEnd
	}
	funcs["TimeWindowSignature"] = func(op models.Operation, excluded ...string) string {
		excludedMap := make(map[string]bool)
		for _, e := range excluded {
			excludedMap[e] = true
		}

		var args []string
		for _, p := range op.PathParams {
			if excludedMap[p] {
				continue
			}
			args = append(args, fmt.Sprintf("string %s", p))
		}

		if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
			args = append(args, "object? body = null")
		}
		for _, p := range op.QueryParams {
			if p.Name == "lastModifiedStart" || p.Name == "lastModifiedEnd" || p.Name == "pageNumber" || p.Name == "page" {
				continue
			}
			name := strings.TrimSuffix(p.Name, "optional")
			if excludedMap[name] {
				continue
			}

			args = append(args, fmt.Sprintf("string? %s = null", name))
		}
		if len(args) > 0 {
			return strings.Join(args, ",\n    ") + ",\n    "
		}
		return ""
	}

	funcs["TimeWindowClientCall"] = func(op models.Operation, targetVar string) string {

		var args []string
		for _, p := range op.PathParams {
			if p == "licenseNumber" && targetVar != "" {
				args = append(args, targetVar+".LicenseNumber")
			} else {
				args = append(args, p)
			}
		}
		if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
			args = append(args, "body")
		}
		for _, p := range op.QueryParams {
			name := strings.TrimSuffix(p.Name, "optional")
			if name == "lastModifiedStart" {
				args = append(args, "startStr")
			} else if name == "lastModifiedEnd" {
				args = append(args, "endStr")
			} else if name == "pageNumber" || name == "page" {
				continue
			} else if name == "licenseNumber" && targetVar != "" {
				args = append(args, targetVar+".LicenseNumber")
			} else {
				args = append(args, name)
			}
		}
		return strings.Join(args, ", ")
	}

	for _, svc := range ctx.Services {
		hasPaginated := false
		for _, op := range svc.Operations {
			if hasPagination(op) {
				hasPaginated = true
				break
			}
		}

		if hasPaginated {
			fname := fmt.Sprintf("%sServiceExtensions.cs", app.ToPascalCase(app.CleanName(svc.Name)))
			if err := app.RenderTemplate(mainTemplatesFS, "templates/main/ServiceExtensions.cs.tmpl", filepath.Join(extDir, fname), map[string]interface{}{
				"Service": svc,
			}, funcs); err != nil {
				return err
			}
		}
	}

	return nil
}

func mainPaginationParamsDocs(op models.Operation) string {
	paramDocs := app.ParseParamDocs(op.Description)
	var lines []string
	for _, p := range op.PathParams {
		desc := paramDocs[p]
		if desc == "" {
			desc = p
		}
		lines = append(lines, fmt.Sprintf("/// <param name=\"%s\">%s</param>", p, desc))
	}
	if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
		lines = append(lines, "/// <param name=\"body\">Request body.</param>")
	}
	for _, p := range op.QueryParams {
		name := strings.TrimSuffix(p.Name, "optional")
		if name == "pageNumber" || name == "page" {
			continue
		}
		desc := paramDocs[name]
		if desc == "" {
			desc = "Optional query parameter."
		}
		lines = append(lines, fmt.Sprintf("/// <param name=\"%s\">%s</param>", name, desc))
	}
	return strings.Join(lines, "\n    ")
}

