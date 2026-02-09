package typescript

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	app "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
)

//go:embed templates/client/*.tmpl
var clientTemplatesFS embed.FS

type ClientGenerator struct{}

func (g *ClientGenerator) Name() string {
	return "TypeScript"
}

func (g *ClientGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-ts", "client")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	// package.json
	if err := generateClientPackageJson(outDir, ctx.Version, ctx.Deps); err != nil {
		return err
	}

	// tsconfig.json
	if err := generateClientTsConfig(outDir); err != nil {
		return err
	}

	// Client Code
	if err := generateClient(outDir, ctx.Services); err != nil {
		return err
	}

	return nil
}

func generateClientPackageJson(dir, version string, deps map[string]string) error {
	return app.RenderPackageJSON(filepath.Join(dir, "package.json"), app.PackageConfig{
		Version:     version,
		Description: "Auto-generated TypeScript client for ThunkMetrc",
		IsWrapper:   false,
		Deps:        deps,
	})
}

func generateClientTsConfig(dir string) error {
	// Static file
	content := `{
  "compilerOptions": {
    "target": "ES2020",
    "module": "commonjs",
    "declaration": true,
    "outDir": "./dist",
    "strict": true,
    "esModuleInterop": true
  },
  "include": ["src/**/*"]
}
`
	return os.WriteFile(filepath.Join(dir, "tsconfig.json"), []byte(content), 0644)
}

func generateClient(dir string, services []models.Service) error {
	srcDir := filepath.Join(dir, "src")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		return err
	}

	funcs := app.GetBaseTemplateFuncs()
	funcs["FormatParams"] = formatParams
	funcs["FormatUrl"] = formatUrl
	// Override FormatMethodName to use camelCase
	funcs["FormatMethodName"] = func(op, group string) string {
		return app.ToCamelCase(app.FormatMethodName(op, group))
	}

	data := map[string]interface{}{
		"Services": services,
	}
	return app.RenderTemplate(clientTemplatesFS, "templates/client/index.ts.tmpl", filepath.Join(srcDir, "index.ts"), data, funcs)
}

func formatParams(op models.Operation) string {
	var args []string
	// Path params are always strings in standard TS gen
	for _, p := range op.PathParams {
		args = append(args, fmt.Sprintf("%s: string", p))
	}

	args = append(args, "body?: TBody")

	for _, p := range op.QueryParams {
		// handle optional suffix logic if relevant, or just use p.Name
		name := strings.TrimSuffix(p.Name, "optional")
		tsType := inferTSType(name)
		args = append(args, fmt.Sprintf("%s?: %s", name, tsType))
	}
	return strings.Join(args, ", ")
}

func inferTSType(name string) string {
	lower := strings.ToLower(name)
	if lower == "lastmodifiedend" || lower == "lastmodifiedstart" {
		return "string" // Dates are passed as strings
	}
	if strings.Contains(lower, "date") || strings.Contains(lower, "time") {
		return "string"
	}
	if lower == "licensequeryset" {
		return "string" // Looks like string
	}
	if strings.HasSuffix(lower, "id") || strings.HasSuffix(lower, "count") ||
		strings.HasSuffix(lower, "limit") || strings.HasSuffix(lower, "offset") ||
		strings.HasSuffix(lower, "number") || lower == "seq" {
		if strings.Contains(lower, "license") {
			return "string"
		}
		return "number"
	}
	if strings.HasPrefix(lower, "is") || strings.HasPrefix(lower, "has") ||
		strings.HasPrefix(lower, "active") {
		return "boolean"
	}
	return "string"
}

func formatUrl(op models.Operation) string {
	url := op.Path
	for _, p := range op.PathParams {
		url = strings.ReplaceAll(url, "{"+p+"}", "${encodeURIComponent("+p+")}")
	}
	return url
}
