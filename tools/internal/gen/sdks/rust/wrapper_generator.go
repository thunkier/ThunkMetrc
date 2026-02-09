package rust

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	app "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
	"github.com/thunkier/thunkmetrc/tools/pkg/jsonschema"
	"github.com/thunkier/thunkmetrc/tools/pkg/schema"
)

//go:embed templates/wrapper/*.tmpl
var wrapperTemplatesFS embed.FS

var ErrDict = fmt.Errorf("dict keys must be strings")

type WrapperGenerator struct{}

func (g *WrapperGenerator) Name() string {
	return "Rust Wrapper"
}

func (g *WrapperGenerator) Generate(ctx *common.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-rs", "wrapper")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	srcDir := filepath.Join(outDir, "src")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		return err
	}
	modelDir := filepath.Join(srcDir, "models")
	if err := os.MkdirAll(modelDir, 0755); err != nil {
		return err
	}

	// Collect models for easier template iteration
	var allModels []*jsonschema.ModelIR
	var collect func(*jsonschema.ModelIR)
	seen := make(map[string]bool)
	collect = func(m *jsonschema.ModelIR) {
		if m == nil || seen[m.Name] {
			return
		}
		seen[m.Name] = true
		for i := range m.NestedTypes {
			collect(&m.NestedTypes[i])
		}
		allModels = append(allModels, m)
	}

	for _, m := range ctx.ResponseModels {
		collect(m)
	}

	// Request Models
	reqModels := app.CollectRequestModels(ctx.Services, rustTypeMapper, brunoKindMapper, func(_, opName string) string {
		return opName + "Request"
	})
	for _, m := range reqModels {
		collect(m)
	}

	// Sort models by name for stable mod.rs generation
	sort.Slice(allModels, func(i, j int) bool {
		return allModels[i].Name < allModels[j].Name
	})

	svcDir := filepath.Join(srcDir, "services")
	if err := os.MkdirAll(svcDir, 0755); err != nil {
		return err
	}

	// Helper function (defined below/inline)
	formatMethodName := func(opName, group string) string {
		return app.ToSnakeCase(app.FormatMethodName(opName, group))
	}

	funcs := app.GetBaseTemplateFuncs()
	funcs["SafeRustField"] = safeRustField
	funcs["FormatMethodName"] = formatMethodName
	funcs["GetReturnType"] = func(op models.Operation) string {
		typeName := "serde_json::Value"
		if op.ResponseType != "" {
			if _, ok := ctx.ResponseModels[op.ResponseType]; ok {
				typeName = op.ResponseType
			}
		}

		if op.IsPaginated {
			return fmt.Sprintf("PaginatedResponse<%s>", typeName)
		}
		if op.IsArray {
			return fmt.Sprintf("Vec<%s>", typeName)
		}
		return typeName
	}

	// Helper to resolve Rust types from ModelIR properties
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
	funcs["HasStreams"] = func(ops []models.Operation) bool {
		for _, op := range ops {
			hasPageParam := false
			for _, qp := range op.QueryParams {
				if strings.Contains(qp.Name, "pageNumber") {
					hasPageParam = true
					break
				}
			}
			if hasPageParam && (op.IsPaginated || op.IsArray) {
				return true
			}
		}
		return false
	}

	data := map[string]interface{}{
		"Version":        ctx.Version,
		"Services":       ctx.Services,
		"ResponseModels": allModels,
	}

	// Cargo.toml
	if err := app.RenderCargoToml(filepath.Join(outDir, "Cargo.toml"), app.CargoConfig{
		Version:     ctx.Version,
		Description: "Type-safe wrapper for ThunkMetrc Rust client with rate limiting",
		IsWrapper:   true,
		Deps:        ctx.Deps,
	}); err != nil {
		return err
	}

	// README.md
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/README.md.tmpl",
		filepath.Join(outDir, "README.md"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// ratelimiter.rs
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/ratelimiter.rs.tmpl",
		filepath.Join(srcDir, "ratelimiter.rs"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// lib.rs
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/lib.rs.tmpl",
		filepath.Join(srcDir, "lib.rs"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// factory.rs
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/factory.rs.tmpl",
		filepath.Join(srcDir, "factory.rs"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// utils.rs
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/utils.rs.tmpl",
		filepath.Join(srcDir, "utils.rs"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// models/paginated_response.rs
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/paginated_response.rs.tmpl",
		filepath.Join(modelDir, "paginated_response.rs"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// models/mod.rs
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/mod.rs.tmpl",
		filepath.Join(modelDir, "mod.rs"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// Services/mod.rs
	svcModPath := filepath.Join(svcDir, "mod.rs")
	svcModFile, err := os.Create(svcModPath)
	if err != nil {
		return err
	}
	for _, svc := range ctx.Services {
		name := app.ToSnakeCase(app.CleanName(svc.Name))
		if name == "" {
			continue
		}
		fmt.Fprintf(svcModFile, "pub mod %s_service;\n", name)
	}
	svcModFile.Close()

	// Services Service Files
	svcTmpl, err := template.New("service.rs.tmpl").Funcs(funcs).ParseFS(wrapperTemplatesFS, "templates/wrapper/service.rs.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse service template: %w", err)
	}

	for _, svc := range ctx.Services {
		name := app.ToSnakeCase(app.CleanName(svc.Name))
		if name == "" {
			continue
		}
		svcData := map[string]interface{}{
			"Service": svc,
		}
		f, err := os.Create(filepath.Join(svcDir, fmt.Sprintf("%s_service.rs", name)))
		if err != nil {
			return err
		}
		err = svcTmpl.Execute(f, svcData)
		f.Close()
		if err != nil {
			return err
		}
	}

	// Render all models
	for _, m := range allModels {
		modelFileName := app.ToSnakeCase(m.Name)
		modelData := map[string]interface{}{
			"Model": m,
		}
		if err := app.RenderTemplate(
			wrapperTemplatesFS,
			"templates/wrapper/model.rs.tmpl",
			filepath.Join(modelDir, fmt.Sprintf("%s.rs", modelFileName)),
			modelData,
			funcs,
		); err != nil {
			return err
		}
	}

	return nil
}

func rustTypeMapper(s *schema.Schema, parentName, propName string) string {
	if s.Type == schema.TypeObject {
		return parentName + app.ToPascalCase(propName)
	}
	return ""
}

func brunoKindMapper(t schema.SchemaType) jsonschema.TypeKind {
	switch t {
	case schema.TypeString:
		return jsonschema.TypeString
	case schema.TypeInteger:
		return jsonschema.TypeInteger
	case schema.TypeNumber:
		return jsonschema.TypeNumber
	case schema.TypeBoolean:
		return jsonschema.TypeBoolean
	case schema.TypeArray:
		return jsonschema.TypeArray
	case schema.TypeObject:
		return jsonschema.TypeObject
	default:
		return jsonschema.TypeAny
	}
}

func safeRustField(name string) string {
	snake := app.ToSnakeCase(name)
	if len(snake) > 0 && snake[0] >= '0' && snake[0] <= '9' {
		return "_" + snake
	}
	return app.SafeFieldName(snake, "rust")
}
