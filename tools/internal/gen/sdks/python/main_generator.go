package python

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	app "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
)

//go:embed templates/main/*.tmpl
var mainTemplatesFS embed.FS

type MainGenerator struct{}

func (g *MainGenerator) Name() string {
	return "PythonMain"
}

func (g *MainGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-py", "thunkmetrc", "src", "thunkmetrc")
	extDir := filepath.Join(outDir, "extensions")

	if err := os.MkdirAll(extDir, 0755); err != nil {
		return fmt.Errorf("failed to create extensions dir: %w", err)
	}

	// Init Extensions package
	if err := os.WriteFile(filepath.Join(extDir, "__init__.py"), []byte(""), 0644); err != nil {
		return err
	}

	// Render Metrc Extensions (Shared Logic)
	if err := app.RenderTemplate(mainTemplatesFS, "templates/main/metrc_extensions.py.tmpl", filepath.Join(extDir, "metrc_extensions.py"), nil, nil); err != nil {
		return err
	}

	funcs := app.GetBaseTemplateFuncs()
	funcs["ToSnakeCase"] = app.ToSnakeCase
	funcs["ToPascalCase"] = app.ToPascalCase
	funcs["HasPagination"] = hasPagination
	funcs["IsTimeWindowed"] = isTimeWindowed
	funcs["ResolveType"] = resolveType

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
			svcNameSnake := app.ToSnakeCase(app.CleanName(svc.Name))
			fname := fmt.Sprintf("%s_extensions.py", svcNameSnake)
			if err := app.RenderTemplate(mainTemplatesFS, "templates/main/service_extensions.py.tmpl", filepath.Join(extDir, fname), map[string]interface{}{
				"Service": svc,
			}, funcs); err != nil {
				return err
			}
		}
	}

	return nil
}

func hasPagination(op models.Operation) bool {
	// Check if page params exist
	for _, p := range op.QueryParams {
		if p.Name == "pageNumber" || p.Name == "page" {
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
