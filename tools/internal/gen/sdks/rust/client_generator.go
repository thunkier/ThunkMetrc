package rust

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	app "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
)

//go:embed templates/client/*.tmpl
var clientTemplatesFS embed.FS

type ClientGenerator struct{}

func (g *ClientGenerator) Name() string {
	return "Rust"
}

func (g *ClientGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-rs", "client")
	// Create src dir
	if err := os.MkdirAll(filepath.Join(outDir, "src"), 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	// Cargo.toml
	if err := generateClientCargoToml(outDir, ctx.Version, ctx.Deps); err != nil {
		return err
	}

	// Client Code
	if err := generateClient(outDir, ctx.Services); err != nil {
		return err
	}

	return nil
}

func generateClientCargoToml(dir, version string, deps map[string]string) error {
	return app.RenderCargoToml(filepath.Join(dir, "Cargo.toml"), app.CargoConfig{
		Version:     version,
		Description: "Auto-generated Rust client for ThunkMetrc",
		IsWrapper:   false,
		Deps:        deps,
	})
}

func generateClient(dir string, services []models.Service) error {
	funcs := app.GetBaseTemplateFuncs()
	funcs["FormatUrl"] = formatUrl
	// Override FormatMethodName to use snake_case
	funcs["FormatMethodName"] = func(op, group string) string {
		return app.ToSnakeCase(app.FormatMethodName(op, group))
	}

	data := map[string]interface{}{
		"Services": services,
	}

	// 1. Generate lib.rs (Exports)
	if err := app.RenderTemplate(clientTemplatesFS, "templates/client/lib.rs.tmpl", filepath.Join(dir, "src", "lib.rs"), data, funcs); err != nil {
		return err
	}

	// 2. Generate client.rs (Base Client)
	if err := app.RenderTemplate(clientTemplatesFS, "templates/client/base_client.rs.tmpl", filepath.Join(dir, "src", "client.rs"), data, funcs); err != nil {
		return err
	}

	// 3. Create services directory
	servicesDir := filepath.Join(dir, "src", "services")
	if err := os.MkdirAll(servicesDir, 0755); err != nil {
		return fmt.Errorf("failed to create services dir: %w", err)
	}

	// 4. Generate services/mod.rs
	modContent := ""
	for _, svc := range services {
		modContent += fmt.Sprintf("pub mod %s;\n", app.ToSnakeCase(app.CleanName(svc.Name)))
		modContent += fmt.Sprintf("pub use %s::%sClient;\n", app.ToSnakeCase(app.CleanName(svc.Name)), app.ToPascalCase(app.CleanName(svc.Name)))
	}
	if err := os.WriteFile(filepath.Join(servicesDir, "mod.rs"), []byte(modContent), 0644); err != nil {
		return err
	}

	// 5. Generate individual service files
	for _, svc := range services {
		fname := fmt.Sprintf("%s.rs", app.ToSnakeCase(app.CleanName(svc.Name)))
		if err := app.RenderTemplate(clientTemplatesFS, "templates/client/service_client.rs.tmpl", filepath.Join(servicesDir, fname), svc, funcs); err != nil {
			return err
		}
	}

	return nil
}

func formatUrl(op models.Operation) string {
	// Use regex to replace all {param} with {}
	re := regexp.MustCompile(`\{[^}]+\}`)
	return re.ReplaceAllString(op.Path, "{}")
}
