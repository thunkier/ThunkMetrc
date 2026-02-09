package rust

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	app "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
	"github.com/thunkier/thunkmetrc/tools/pkg/jsonschema"
)

//go:embed templates/main/service_extensions.rs.tmpl
var mainTemplatesFS embed.FS

type MainGenerator struct{}

func (g *MainGenerator) Name() string {
	return "Rust Main"
}

func (g *MainGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-rs", "thunkmetrc")
	extDir := filepath.Join(outDir, "src", "extensions")

	if err := os.MkdirAll(extDir, 0755); err != nil {
		return err
	}

	funcs := app.GetBaseTemplateFuncs()
	funcs["ToSnakeCase"] = app.ToSnakeCase
	funcs["ToPascalCase"] = app.ToPascalCase
	funcs["FormatMethodName"] = func(opName, group string) string {
		return app.ToSnakeCase(app.FormatMethodName(opName, group))
	}

	// Rust Type helpers (copied/adapted from wrapper generator)
	funcs["SafeRustField"] = func(name string) string {
		snake := app.ToSnakeCase(name)
		if len(snake) > 0 && snake[0] >= '0' && snake[0] <= '9' {
			return "_" + snake
		}
		return app.SafeFieldName(snake, "rust")
	}

	funcs["GetReturnType"] = func(op models.Operation) string {
		// We only care about the inner type for extensions usually
		typeName := "serde_json::Value"
		if op.ResponseType != "" {
			// simplified lookup, assume context or strings match
			typeName = op.ResponseType
		}
		return typeName
	}

	// Helper to determine if we should generate extensions for this op
	funcs["HasPagination"] = func(op models.Operation) bool {
		hasPageParam := false
		for _, qp := range op.QueryParams {
			if strings.Contains(qp.Name, "pageNumber") {
				hasPageParam = true
				break
			}
		}
		if (op.IsPaginated || op.IsArray) && hasPageParam {
			return true
		}
		return false
	}
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

	// Resolving Rust types
	var toRustType func(t jsonschema.TypeRef) string
	toRustType = func(t jsonschema.TypeRef) string {
		switch t.Kind {
		case jsonschema.TypeString:
			return "String"
		case jsonschema.TypeInteger:
			return "i64"
		case jsonschema.TypeNumber:
			return "f64"
		case jsonschema.TypeBoolean:
			return "bool"
		case jsonschema.TypeArray:
			if t.ItemType != nil {
				return fmt.Sprintf("Vec<%s>", toRustType(*t.ItemType))
			}
			return "Vec<serde_json::Value>"
		case jsonschema.TypeObject:
			if t.Name != "" {
				return t.Name
			}
			return "serde_json::Value"
		default:
			if t.Name != "" {
				return t.Name
			}
			return "serde_json::Value"
		}
	}
	funcs["ToRustType"] = toRustType

	// String helpers
	funcs["TrimSuffix"] = func(suffix, s string) string { return strings.TrimSuffix(s, suffix) }
	funcs["TrimPrefix"] = func(prefix, s string) string { return strings.TrimPrefix(s, prefix) }
	funcs["HasPrefix"] = strings.HasPrefix
	funcs["Contains"] = strings.Contains
	funcs["join"] = strings.Join
	funcs["add"] = func(a, b int) int { return a + b }
	funcs["list"] = func(values ...string) []string { return values }
	funcs["append"] = func(l []string, v string) []string { return append(l, v) }

	tmpl, err := template.New("service_extensions.rs.tmpl").Funcs(funcs).ParseFS(mainTemplatesFS, "templates/main/service_extensions.rs.tmpl")
	if err != nil {
		return err
	}

	generatedModules := []string{}

	for _, svc := range ctx.Services {
		shouldGen := false
		for _, op := range svc.Operations {
			if funcs["HasPagination"].(func(models.Operation) bool)(op) || funcs["IsTimeWindowed"].(func(models.Operation) bool)(op) {
				shouldGen = true
				break
			}
		}

		if !shouldGen {
			continue
		}

		svcName := app.ToSnakeCase(app.CleanName(svc.Name))
		if svcName == "" {
			continue
		}

		moduleName := fmt.Sprintf("%s_extensions", svcName)
		generatedModules = append(generatedModules, moduleName)

		f, err := os.Create(filepath.Join(extDir, fmt.Sprintf("%s.rs", moduleName)))
		if err != nil {
			return err
		}

		if err := tmpl.Execute(f, map[string]interface{}{
			"Service": svc,
			"Name":    app.ToPascalCase(app.CleanName(svc.Name)), // Pascal for Struct Name
		}); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}

	// Generate mod.rs for extensions
	modPath := filepath.Join(extDir, "mod.rs")
	modFile, err := os.Create(modPath)
	if err != nil {
		return err
	}
	defer modFile.Close()

	for _, mod := range generatedModules {
		fmt.Fprintf(modFile, "pub mod %s;\n", mod)
	}

	return nil
}
