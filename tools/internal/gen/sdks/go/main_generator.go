package golang

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	app "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
)

//go:embed templates/main/*.tmpl
var mainTemplatesFS embed.FS

type MainGenerator struct{}

func (g *MainGenerator) Name() string {
	return "GoMain"
}

func (g *MainGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-go", "thunkmetrc", "extensions")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	funcs := app.GetBaseTemplateFuncs()
	funcs["HasPagination"] = hasPagination
	funcs["IsTimeWindowed"] = isTimeWindowed

	// Register missing funcs used in template
	funcs["ToPascalCase"] = app.ToPascalCase
	funcs["ToCamelCase"] = app.ToCamelCase
	funcs["CleanName"] = app.CleanName
	funcs["HasPrefix"] = strings.HasPrefix
	funcs["TrimPrefix"] = func(prefix, s string) string { return strings.TrimPrefix(s, prefix) }
	funcs["TrimSuffix"] = func(suffix, s string) string { return strings.TrimSuffix(s, suffix) }
	funcs["join"] = strings.Join
	funcs["list"] = func(values ...string) []string { return values }
	funcs["append"] = func(l []string, v string) []string { return append(l, v) }

	// FormatMethodName reused logic
	funcs["FormatMethodName"] = app.FormatMethodName

	// Simplified GetBodyType reusing Wrapper conventions
	funcs["GetBodyType"] = func(group string, op models.Operation) string {
		if op.BodySchema == nil {
			return "interface{}"
		}

		opName := app.ToPascalCase(app.CleanName(op.Name))
		reqName := group + opName + "Request"

		if op.BodySchema.Type == "array" {
			return "[]*models." + reqName + "Item"
		}
		return "*models." + reqName
	}

	// GetReturnType needed to distinguish paginated types
	funcs["GetReturnType"] = func(op models.Operation) string {
		if op.ResponseType == "" {
			return "interface{}"
		}
		// ResponseModels map is available in ctx
		if _, exists := ctx.ResponseModels[op.ResponseType]; !exists {
			return "interface{}"
		}
		typeName := "*models." + op.ResponseType
		if op.IsPaginated {
			return fmt.Sprintf("*models.PaginatedResponse[%s]", typeName)
		}
		if op.IsArray {
			return fmt.Sprintf("[]%s", typeName)
		}
		return typeName
	}

	tmpl, err := template.New("service_extensions.go.tmpl").Funcs(funcs).ParseFS(mainTemplatesFS, "templates/main/service_extensions.go.tmpl")
	if err != nil {
		return err
	}

	for _, svc := range ctx.Services {
		hasPaginated := false
		hasTimeWindowed := false
		for _, op := range svc.Operations {
			if hasPagination(op) {
				hasPaginated = true
			}
			if isTimeWindowed(op) {
				hasTimeWindowed = true
			}
		}

		if hasPaginated || hasTimeWindowed {
			name := app.ToPascalCase(app.CleanName(svc.Name))
			fname := fmt.Sprintf("generated_%s_extensions.go", strings.ToLower(name))
			f, err := os.Create(filepath.Join(outDir, fname))
			if err != nil {
				return err
			}

			if err := tmpl.Execute(f, map[string]interface{}{
				"Service": svc,
				"Name":    name,
			}); err != nil {
				f.Close()
				return err
			}
			f.Close()
		}
	}

	return nil
}

func hasPagination(op models.Operation) bool {
	for _, p := range op.QueryParams {
		if strings.Contains(p.Name, "pageNumber") {
			return true
		}
	}
	return false
}

func isTimeWindowed(op models.Operation) bool {
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
