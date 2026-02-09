package python

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	app "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
)

//go:embed templates/client/*.tmpl
var clientTemplatesFS embed.FS

// Generator for Python client code.
type ClientGenerator struct{}

// Name returns the name of the generator.
func (g *ClientGenerator) Name() string {
	return "Python"
}

// Generate creates the Python client SDK.
func (g *ClientGenerator) Generate(ctx *common.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-py", "client")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(outDir, "src", "thunkmetrc", "client"), 0755); err != nil {
		return err
	}

	if err := generatePyProject(outDir, ctx.Version, ctx.Deps); err != nil {
		return err
	}

	if err := generateReadme(outDir); err != nil {
		return err
	}

	if err := generateClient(outDir, ctx.Services); err != nil {
		return err
	}

	return nil
}

// generatePyProject creates the pyproject.toml file.
func generatePyProject(dir, version string, deps map[string]string) error {
	return app.RenderPyProjectToml(filepath.Join(dir, "pyproject.toml"), app.PyProjectConfig{
		Version:     version,
		Description: "Auto-generated Python client for ThunkMetrc",
		IsWrapper:   false,
		Deps:        deps,
	})
}

// generateReadme creates the README.md file.
func generateReadme(dir string) error {
	content := `# ThunkMetrc Client

Auto-generated Python client for ThunkMetrc API.

## Installation

` + "```bash" + `
pip install thunkmetrc-client
` + "```" + `

## Usage

swag
`
	return os.WriteFile(filepath.Join(dir, "README.md"), []byte(content), 0644)
}

// generateClient creates the Python client code.
func generateClient(dir string, services []models.Service) error {
	srcDir := filepath.Join(dir, "src", "thunkmetrc", "client")
	servicesDir := filepath.Join(srcDir, "services")

	// Create directories
	if err := os.MkdirAll(servicesDir, 0755); err != nil {
		return err
	}

	// Create services/__init__.py
	os.WriteFile(filepath.Join(servicesDir, "__init__.py"), []byte(""), 0644)

	// Base Client
	if err := generateBaseClient(srcDir, services); err != nil {
		return err
	}

	// Service Clients
	for _, svc := range services {
		if err := generateServiceClient(servicesDir, svc); err != nil {
			return err
		}
	}

	// Main __init__ re-export
	os.WriteFile(filepath.Join(srcDir, "__init__.py"), []byte("from .client import MetrcClient\n"), 0644)

	return nil
}

func generateBaseClient(dir string, services []models.Service) error {
	funcs := app.GetBaseTemplateFuncs()
	data := map[string]interface{}{
		"Services": services,
	}
	return app.RenderTemplate(clientTemplatesFS, "templates/client/base_client.py.tmpl", filepath.Join(dir, "client.py"), data, funcs)
}

func generateServiceClient(dir string, svc models.Service) error {
	funcs := app.GetBaseTemplateFuncs()
	funcs["FormatParams"] = formatParams
	funcs["FormatUrl"] = formatUrl
	// Override FormatMethodName to use snake_case
	funcs["FormatMethodName"] = func(op, group string) string {
		return app.ToSnakeCase(app.FormatMethodName(op, group))
	}

	filename := fmt.Sprintf("%s.py", app.ToSnakeCase(svc.Name))
	return app.RenderTemplate(clientTemplatesFS, "templates/client/service_client.py.tmpl", filepath.Join(dir, filename), svc, funcs)
}

// formatParams formats operation parameters for Python function signatures.
func formatParams(op models.Operation) string {
	var args []string
	// Path parameters
	for _, p := range op.PathParams {
		args = append(args, fmt.Sprintf("%s: str", app.ToSnakeCase(p)))
	}

	// Body parameter
	if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
		args = append(args, "body: Any = None")
	}

	// Query Params
	for _, p := range op.QueryParams {
		name := strings.TrimSuffix(p.Name, "optional")
		args = append(args, fmt.Sprintf("%s: Optional[str] = None", app.ToSnakeCase(name)))
	}

	if len(args) == 0 {
		return "self"
	}
	return "self, " + strings.Join(args, ", ")
}

func formatUrl(op models.Operation) string {
	url := op.Path
	for _, p := range op.PathParams {
		snake := app.ToSnakeCase(p)
		url = strings.ReplaceAll(url, "{"+p+"}", "{urllib.parse.quote("+snake+")}")
	}
	return url
}

