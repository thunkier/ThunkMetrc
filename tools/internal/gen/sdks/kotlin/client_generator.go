package kotlin

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

//go:embed templates/client/*.tmpl
var clientTemplatesFS embed.FS

type ClientGenerator struct{}

func (g *ClientGenerator) Name() string {
	return "Kotlin"
}

func (g *ClientGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-kotlin", "client")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	pkgDir := filepath.Join(outDir, "src", "main", "kotlin", "io", "github", "thunkier", "thunkmetrc", "client")
	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		return err
	}

	// pom.xml
	if err := generateClientPomXml(outDir, ctx.Version, ctx.Deps); err != nil {
		return err
	}

	// Client Code
	if err := generateMainClient(pkgDir, ctx.Services); err != nil {
		return err
	}
	if err := generateServiceClients(pkgDir, ctx.Services); err != nil {
		return err
	}

	if err := app.RenderTemplate(clientTemplatesFS, "templates/client/MetrcException.kt.tmpl", filepath.Join(pkgDir, "MetrcException.kt"), nil, nil); err != nil {
		return err
	}

	return nil
}

func generateClientPomXml(dir, version string, deps map[string]string) error {
	return app.RenderPom(filepath.Join(dir, "pom.xml"), app.PomConfig{
		ArtifactId:  "thunkmetrc-kotlin-client",
		Name:        "ThunkMetrc Kotlin Client",
		Description: "Auto-generated Kotlin client for ThunkMetrc API.",
		Version:     version,
		IsKotlin:    true,
		IsWrapper:   false,
		Deps:        deps,
	})
}

func getClientFuncs() template.FuncMap {
	funcs := app.GetBaseTemplateFuncs()
	funcs["FormatUrl"] = formatUrl
	// Override FormatMethodName to use camelCase
	funcs["FormatMethodName"] = func(op, group string) string {
		return app.ToCamelCase(app.FormatMethodName(op, group))
	}
	return funcs
}

func generateMainClient(dir string, services []models.Service) error {
	data := map[string]interface{}{
		"Services": services,
	}
	return app.RenderTemplate(clientTemplatesFS, "templates/client/MetrcClient.kt.tmpl", filepath.Join(dir, "MetrcClient.kt"), data, getClientFuncs())
}

func generateServiceClients(dir string, services []models.Service) error {
	funcs := getClientFuncs()
	for _, svc := range services {
		name := app.ToPascalCase(app.CleanName(svc.Name))
		if name == "" {
			continue
		}
		if err := app.RenderTemplate(clientTemplatesFS, "templates/client/ServiceClient.kt.tmpl", filepath.Join(dir, fmt.Sprintf("%sClient.kt", name)), svc, funcs); err != nil {
			return err
		}
	}
	return nil
}

func formatUrl(op models.Operation) string {
	url := op.Path
	for _, p := range op.PathParams {
		url = strings.ReplaceAll(url, "{"+p+"}", "${URLEncoder.encode("+app.ToCamelCase(p)+", StandardCharsets.UTF_8)}")
	}
	return url
}
