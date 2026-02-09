package typescript

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

type WrapperGenerator struct{}

func (g *WrapperGenerator) Name() string {
	return "TypeScript Wrapper"
}

func (g *WrapperGenerator) Generate(ctx *common.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-ts", "wrapper")
	srcDir := filepath.Join(outDir, "src")
	svcDir := filepath.Join(srcDir, "services")

	if err := os.RemoveAll(svcDir); err != nil {
		return err
	}
	if err := os.MkdirAll(svcDir, 0755); err != nil {
		return err
	}

	// 1. Collect Request Models
	reqModels := app.CollectRequestModels(ctx.Services, tsTypeMapper, brunoKindMapper, func(_, opName string) string {
		return opName + "Request"
	})

	// 2. Collect Response Models and Flatten
	var allModels []*jsonschema.ModelIR
	seen := make(map[string]bool)

	collect := func(m *jsonschema.ModelIR) {
		var visit func(m *jsonschema.ModelIR)
		visit = func(m *jsonschema.ModelIR) {
			if m == nil || seen[m.Name] {
				return
			}
			seen[m.Name] = true
			for i := range m.NestedTypes {
				visit(&m.NestedTypes[i])
			}
			allModels = append(allModels, m)
		}
		visit(m)
	}

	// Add Response Models
	for _, m := range ctx.ResponseModels {
		collect(m)
	}
	// Add Request Models
	for _, m := range reqModels {
		collect(m)
	}

	// Sort models by name
	sort.Slice(allModels, func(i, j int) bool {
		return allModels[i].Name < allModels[j].Name
	})

	// Backfill missing models
	knownModels := make(map[string]bool)
	for _, m := range allModels {
		knownModels[m.Name] = true
	}

	for _, svc := range ctx.Services {
		for _, op := range svc.Operations {
			rt := op.ResponseType
			if rt != "" && rt != "any" && rt != "string" && rt != "number" && rt != "boolean" && !knownModels[rt] {
				fmt.Printf("[WARN] Backfilling missing model: %s\n", rt)
				allModels = append(allModels, &jsonschema.ModelIR{
					Name: rt,
				})
				knownModels[rt] = true
			}
		}
	}
	// Re-sort
	sort.Slice(allModels, func(i, j int) bool {
		return allModels[i].Name < allModels[j].Name
	})

	data := map[string]interface{}{
		"Version":  ctx.Version,
		"Services": ctx.Services,
		"Models":   allModels,
	}

	funcs := app.GetBaseTemplateFuncs()
	funcs["TrimPrefix"] = func(prefix, s string) string { return strings.TrimPrefix(s, prefix) }
	funcs["ToTSType"] = toTSType
	funcs["FormatMethodName"] = func(opName, group string) string {
		return app.ToCamelCase(app.FormatMethodName(opName, group))
	}
	funcs["InferType"] = inferType
	funcs["GetReturnType"] = func(op models.Operation) string {
		if op.ResponseType == "" {
			return "any"
		}
		typeName := op.ResponseType
		if op.IsPaginated {
			return fmt.Sprintf("PaginatedResponse<%s>", typeName)
		}
		if op.IsArray {
			return fmt.Sprintf("%s[]", typeName)
		}
		return typeName
	}
	funcs["SafeTSField"] = safeTSField

	// package.json
	if err := app.RenderPackageJSON(filepath.Join(outDir, "package.json"), app.PackageConfig{
		Version:     ctx.Version,
		Description: "Type-safe wrapper for ThunkMetrc TypeScript client",
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

	// tsconfig.json
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/tsconfig.json.tmpl",
		filepath.Join(outDir, "tsconfig.json"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// RateLimiter.ts
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/RateLimiter.ts.tmpl",
		filepath.Join(srcDir, "RateLimiter.ts"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// Utils.ts
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/Utils.ts.tmpl",
		filepath.Join(srcDir, "Utils.ts"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// Factory.ts
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/Factory.ts.tmpl",
		filepath.Join(srcDir, "Factory.ts"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// models directory
	modelsDir := filepath.Join(srcDir, "models")
	if err := os.RemoveAll(modelsDir); err != nil {
		return err
	}
	if err := os.MkdirAll(modelsDir, 0755); err != nil {
		return err
	}

	// models/index.ts
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/models_index.ts.tmpl",
		filepath.Join(modelsDir, "index.ts"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// models/PaginatedResponse.ts
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/PaginatedResponse.ts.tmpl",
		filepath.Join(modelsDir, "PaginatedResponse.ts"),
		nil,
		funcs,
	); err != nil {
		return err
	}

	// Individual model files
	modelTmpl, err := template.New("model.ts.tmpl").Funcs(funcs).ParseFS(wrapperTemplatesFS, "templates/wrapper/model.ts.tmpl")
	if err != nil {
		return err
	}

	for _, m := range allModels {
		modelData := map[string]interface{}{
			"Model":   m,
			"Imports": collectModelImports(m, knownModels),
		}

		f, err := os.Create(filepath.Join(modelsDir, fmt.Sprintf("%s.ts", m.Name)))
		if err != nil {
			return err
		}
		if err := modelTmpl.Execute(f, modelData); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}

	// index.ts (Facade)
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/index.ts.tmpl",
		filepath.Join(srcDir, "index.ts"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// Services
	svcTmpl, err := template.New("Service.ts.tmpl").Funcs(funcs).ParseFS(wrapperTemplatesFS, "templates/wrapper/Service.ts.tmpl")
	if err != nil {
		return err
	}

	for _, svc := range ctx.Services {
		name := app.ToPascalCase(app.CleanName(svc.Name))
		if name == "" {
			continue
		}

		// Compute Imports
		importSet := make(map[string]bool)
		for _, op := range svc.Operations {
			// 1. Return Types
			if op.ResponseType != "" {
				baseType := op.ResponseType
				if baseType != "any" && baseType != "string" && baseType != "number" && baseType != "boolean" {
					importSet[baseType] = true
				}
			}

			// 2. Request Types
			if op.BodySchema != nil {
				opName := app.ToPascalCase(app.CleanName(op.Name))
				// svcPascal used to be used here
				reqName := opName + "Request"

				if op.BodySchema.Type == "array" {
					if op.BodySchema.Items != nil {
						importSet[reqName+"Item"] = true
					}
				} else {
					importSet[reqName] = true
				}
			}
		}

		var imports []string
		for k := range importSet {
			imports = append(imports, k)
		}
		sort.Strings(imports)

		svcData := map[string]interface{}{
			"Service": svc,
			"Imports": imports,
		}
		f, err := os.Create(filepath.Join(svcDir, fmt.Sprintf("%sService.ts", name)))
		if err != nil {
			return err
		}
		err = svcTmpl.Execute(f, svcData)
		f.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

// Helpers

func tsTypeMapper(s *schema.Schema, parentName, propName string) string {
	if s.Type == schema.TypeObject {
		return parentName + app.ToPascalCase(propName)
	}
	return ""
}

func brunoKindMapper(t schema.SchemaType) jsonschema.TypeKind {
	switch t {
	case schema.TypeString:
		return jsonschema.TypeString
	case schema.TypeInteger, schema.TypeNumber:
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

func safeTSField(name string) string {
	if strings.ContainsAny(name, "- \t\n!@#$%^&*()[]{}+=<>?,./;:'\"\\|`~") {
		return fmt.Sprintf("'%s'", name)
	}
	return app.SafeFieldName(name, "typescript")
}

func toTSType(t jsonschema.TypeRef) string {
	switch t.Kind {
	case jsonschema.TypeString:
		return "string"
	case jsonschema.TypeInteger, jsonschema.TypeNumber:
		return "number"
	case jsonschema.TypeBoolean:
		return "boolean"
	case jsonschema.TypeArray:
		if t.ItemType != nil {
			return toTSType(*t.ItemType) + "[]"
		}
		return "any[]"
	case jsonschema.TypeObject:
		if t.Name != "" {
			return t.Name
		}
		return "any"
	default:
		if t.Name != "" {
			return t.Name
		}
		return "any"
	}
}

func inferType(name string) string {
	lower := strings.ToLower(name)
	if lower == "lastmodifiedend" || lower == "lastmodifiedstart" {
		return "string"
	}
	if strings.Contains(lower, "date") || strings.Contains(lower, "time") {
		return "string"
	}
	if lower == "licensequeryset" {
		return "string"
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

var tsTypeNamePattern = regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_]*`)

func collectModelImports(model *jsonschema.ModelIR, knownModels map[string]bool) []string {
	importSet := make(map[string]bool)

	for _, prop := range model.Properties {
		typeExpr := toTSType(prop.Type)
		for _, token := range tsTypeNamePattern.FindAllString(typeExpr, -1) {
			switch token {
			case "string", "number", "boolean", "any", "null", "undefined":
				continue
			}
			if token == model.Name {
				continue
			}
			if !knownModels[token] {
				continue
			}
			importSet[token] = true
		}
	}

	imports := make([]string, 0, len(importSet))
	for name := range importSet {
		imports = append(imports, name)
	}
	sort.Strings(imports)
	return imports
}
