package golang

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
	return "Go"
}

func (g *ClientGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-go", "client")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	// go.mod
	if err := generateGoMod(outDir); err != nil {
		return err
	}

	// Base Client Code
	if err := generateBaseClient(outDir); err != nil {
		return err
	}

	// Service Clients
	for _, svc := range ctx.Services {
		if err := generateServiceClient(outDir, svc); err != nil {
			return err
		}
	}

	return nil
}

func generateGoMod(dir string) error {
	data := map[string]interface{}{
		"Module": "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client",
	}
	return app.RenderTemplate(app.TemplatesFS, "templates/go.mod.tmpl", filepath.Join(dir, "go.mod"), data, nil)
}

func generateBaseClient(dir string) error {
	return app.RenderTemplate(clientTemplatesFS, "templates/client/base_client.go.tmpl", filepath.Join(dir, "client.go"), nil, nil)
}

func generateServiceClient(dir string, svc models.Service) error {
	funcs := app.GetBaseTemplateFuncs()
	funcs["FormatUrl"] = formatUrl
	funcs["FirstSentence"] = func(s string) string {
		lines := strings.Split(s, "\n")
		var validLines []string
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "Permissions Required") || strings.HasPrefix(trimmed, "-") {
				continue
			}
			if trimmed != "" {
				validLines = append(validLines, trimmed)
			}
		}

		if len(validLines) == 0 {
			return "performs the operation."
		}

		s = strings.Join(validLines, " ")
		parts := strings.Split(s, ".")
		if len(parts) > 0 && parts[0] != "" {
			return parts[0] + "."
		}
		return s + "."
	}
	funcs["NeedsUrlImport"] = func(svc models.Service) bool {
		for _, op := range svc.Operations {
			if len(op.PathParams) > 0 {
				return true
			}
		}
		return false
	}

	safeGroup := app.ToSnakeCase(app.CleanName(svc.Name))
	filename := fmt.Sprintf("client_%s.go", safeGroup)

	return app.RenderTemplate(clientTemplatesFS, "templates/client/service_client.go.tmpl", filepath.Join(dir, filename), svc, funcs)
}

func formatUrl(op models.Operation) string {
	url := op.Path
	for _, p := range op.PathParams {
		url = strings.ReplaceAll(url, "{"+p+"}", "\"+url.QueryEscape("+app.ToCamelCase(p)+")+\"")
	}
	return url
}
