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

//go:embed templates/client/*.tmpl
var clientTemplatesFS embed.FS

type ClientGenerator struct{}

func (g *ClientGenerator) Name() string {
	return "CSharp"
}

func (g *ClientGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-cs", "ThunkMetrc.Client")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	// .csproj
	if err := generateCsproj(outDir, ctx.Version, ctx.Deps); err != nil {
		return err
	}

	// Base Client
	if err := generateBaseClient(outDir); err != nil {
		return err
	}

	// Operation Classes
	for _, svc := range ctx.Services {
		if err := generateService(outDir, svc); err != nil {
			return err
		}
	}

	return nil
}

func generateCsproj(dir, version string, deps map[string]string) error {
	dependencies := make([]map[string]string, 0)
	prefix := "csharp.client."
	for k, v := range deps {
		if strings.HasPrefix(k, prefix) {
			dependencies = append(dependencies, map[string]string{
				"Name":    strings.TrimPrefix(k, prefix),
				"Version": v,
			})
		}
	}

	data := map[string]interface{}{
		"Version":      version,
		"PackageId":    "ThunkMetrc.Client",
		"Description":  "Auto-generated C# client for ThunkMetrc API.",
		"Tags":         "ThunkMetrc;Cannabis;Compliance;API;SDK",
		"Dependencies": dependencies,
	}
	return app.RenderTemplate(app.TemplatesFS, "templates/csproj.xml.tmpl", filepath.Join(dir, "ThunkMetrc.Client.csproj"), data, nil)
}

func generateBaseClient(dir string) error {
	return app.RenderTemplate(clientTemplatesFS, "templates/client/MetrcClient.cs.tmpl", filepath.Join(dir, "MetrcClient.cs"), nil, nil)
}

func generateService(dir string, svc models.Service) error {
	safeGroup := app.ToPascalCase(app.CleanName(svc.Name))
	filename := fmt.Sprintf("MetrcClient.%s.cs", safeGroup)

	funcs := app.GetBaseTemplateFuncs()
	funcs["FormatDocs"] = formatDocs
	funcs["FormatParams"] = formatParams
	funcs["FormatUrl"] = formatUrl
	funcs["SendArgs"] = sendArgs

	return app.RenderTemplate(clientTemplatesFS, "templates/client/service.cs.tmpl", filepath.Join(dir, filename), svc, funcs)
}

func formatParams(op models.Operation) string {
	var args []string
	for _, p := range op.PathParams {
		args = append(args, fmt.Sprintf("string %s", p))
	}

	if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
		args = append(args, "object? body = null")
	}

	for _, p := range op.QueryParams {
		name := strings.TrimSuffix(p.Name, "optional")
		args = append(args, fmt.Sprintf("string? %s = null", name))
	}
	return strings.Join(args, ", ")
}

func formatUrl(op models.Operation) string {
	url := op.Path
	for _, p := range op.PathParams {
		url = strings.ReplaceAll(url, "{"+p+"}", "{Uri.EscapeDataString("+p+")}")
	}
	return url
}

func sendArgs(op models.Operation) string {
	if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
		return "body"
	}
	return "null"
}

func formatDocs(docs string) string {
	lines := strings.Split(docs, "\n")
	for i, line := range lines {
		lines[i] = "        /// " + line
	}
	return strings.Join(lines, "\n")
}

