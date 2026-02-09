package kotlin

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"sort"
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
	return "Kotlin Wrapper"
}

func (g *WrapperGenerator) Generate(ctx *common.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-kotlin", "wrapper")

	pkgDir := filepath.Join(outDir, "src", "main", "kotlin", "io", "github", "thunkier", "thunkmetrc", "wrapper")
	modelDir := filepath.Join(pkgDir, "models")
	svcDir := filepath.Join(pkgDir, "services")

	if err := os.MkdirAll(modelDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(svcDir, 0755); err != nil {
		return err
	}

	// 1. Collect Response Models
	var allModels []*jsonschema.ModelIR
	modelMap := make(map[string]*jsonschema.ModelIR)

	collect := func(m *jsonschema.ModelIR) {
		if m == nil {
			return
		}
		if _, exists := modelMap[m.Name]; exists {
			return
		}
		modelMap[m.Name] = m
	}
	// Deep collect helper
	var deepCollect func(*jsonschema.ModelIR)
	deepCollect = func(m *jsonschema.ModelIR) {
		if m == nil {
			return
		}
		collect(m)
		for i := range m.NestedTypes {
			deepCollect(&m.NestedTypes[i])
		}
	}

	for _, m := range ctx.ResponseModels {
		deepCollect(m)
	}

	// 2. Collect Request Models (Deduplicated)
	reqModels := app.CollectRequestModels(ctx.Services, kotlinTypeMapper, brunoKindMapper, func(svcName, opName string) string {
		safeGroup := app.ToPascalCase(app.CleanName(svcName))
		safeOpName := app.ToPascalCase(app.CleanName(opName))
		return safeGroup + safeOpName + "Request"
	})
	for _, m := range reqModels {
		deepCollect(m)
	}

	// Flatten to list
	for _, m := range modelMap {
		allModels = append(allModels, m)
	}
	// Sort
	sort.Slice(allModels, func(i, j int) bool { return allModels[i].Name < allModels[j].Name })

	data := map[string]interface{}{
		"Version":        ctx.Version,
		"Services":       ctx.Services,
		"ResponseModels": ctx.ResponseModels, // For GetReturnType lookup
	}

	// Helper Funcs
	formatMethodName := func(opName, group string) string {
		return app.ToCamelCase(app.FormatMethodName(opName, group))
	}

	// ResolveType for models
	var resolveType func(s *schema.Schema, parentName, propName string) string
	resolveType = func(s *schema.Schema, parentName, propName string) string {
		if s == nil {
			return "Any"
		}
		switch s.Type {
		case "string":
			return "String"
		case "integer", "int":
			return "Int"
		case "number", "float":
			return "Double"
		case "boolean", "bool":
			return "Boolean"
		case "array":
			inner := "Any"
			if s.Items != nil {
				inner = resolveType(s.Items, parentName, propName)
			}
			return fmt.Sprintf("List<%s>", inner)
		case "object":
			// If we are in this function, it usually means we are resolving a property type
			// inside a template, but we prefer ModelIR for complex types now.
			// This might still be used by legacy templates if any.
			// But for generating explicit classes, we rely on ModelIR type.
			// However, for request body generation inside Service calls, we might need names.
			return propName
		default:
			return "Any"
		}
	}

	// ToKotlinType for ModelIR
	var toKotlinType func(t jsonschema.TypeRef) string
	toKotlinType = func(t jsonschema.TypeRef) string {
		switch t.Kind {
		case jsonschema.TypeString:
			return "String"
		case jsonschema.TypeInteger:
			return "Int"
		case jsonschema.TypeNumber:
			return "Double"
		case jsonschema.TypeBoolean:
			return "Boolean"
		case jsonschema.TypeArray:
			if t.ItemType != nil {
				return fmt.Sprintf("List<%s>", toKotlinType(*t.ItemType))
			}
			return "List<Any>"
		case jsonschema.TypeObject:
			if t.Name != "" {
				return t.Name
			}
			return "Any"
		default:
			if t.Name != "" {
				return t.Name
			}
			return "Any"
		}
	}

	funcs := app.GetBaseTemplateFuncs()
	funcs["ResolveType"] = resolveType
	funcs["SafeKotlinField"] = safeKotlinField
	funcs["FormatMethodName"] = formatMethodName
	funcs["ToKotlinType"] = toKotlinType
	funcs["add"] = func(a, b int) int { return a + b }
	funcs["GetReturnType"] = func(op models.Operation) string {
		if op.ResponseType == "" {
			return "Any"
		}
		if _, ok := ctx.ResponseModels[op.ResponseType]; !ok {
			return "Any"
		}
		typeName := op.ResponseType
		if op.IsPaginated {
			return fmt.Sprintf("PaginatedResponse<%s>", typeName)
		}
		if op.IsArray {
			return fmt.Sprintf("List<%s>", typeName)
		}
		return typeName
	}

	// pom.xml
	if err := app.RenderPom(filepath.Join(outDir, "pom.xml"), app.PomConfig{
		ArtifactId:       "thunkmetrc-kotlin-wrapper",
		Name:             "ThunkMetrc Kotlin Wrapper",
		Description:      "Type-safe wrapper for ThunkMetrc Kotlin client with rate limiting.",
		Version:          ctx.Version,
		IsKotlin:         true,
		IsWrapper:        true,
		ClientArtifactId: "thunkmetrc-kotlin-client",
		Deps:             ctx.Deps,
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

	// MetrcRateLimiter.kt
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/MetrcRateLimiter.kt.tmpl",
		filepath.Join(pkgDir, "MetrcRateLimiter.kt"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// MetrcWrapper.kt (Facade)
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/MetrcWrapper.kt.tmpl",
		filepath.Join(pkgDir, "MetrcWrapper.kt"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// MetrcFactory.kt
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/MetrcFactory.kt.tmpl",
		filepath.Join(pkgDir, "MetrcFactory.kt"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// MetrcExtensions.kt
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/MetrcExtensions.kt.tmpl",
		filepath.Join(pkgDir, "MetrcExtensions.kt"),
		nil,
		funcs,
	); err != nil {
		return err
	}

	// PaginatedResponse.kt
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/PaginatedResponse.kt.tmpl",
		filepath.Join(modelDir, "PaginatedResponse.kt"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// Render all models
	for _, m := range allModels {
		modelData := map[string]interface{}{
			"Model": m,
		}
		if err := app.RenderTemplate(
			wrapperTemplatesFS,
			"templates/wrapper/Model.kt.tmpl",
			filepath.Join(modelDir, fmt.Sprintf("%s.kt", m.Name)),
			modelData,
			funcs,
		); err != nil {
			return err
		}
	}

	// Render Services
	svcTmpl, err := template.New("Service.kt.tmpl").Funcs(funcs).ParseFS(wrapperTemplatesFS, "templates/wrapper/Service.kt.tmpl")
	if err != nil {
		return err
	}

	for _, svc := range ctx.Services {
		name := app.ToPascalCase(app.CleanName(svc.Name))
		if name == "" {
			continue
		}

		svcData := map[string]interface{}{
			"Name":           name, // Service Name e.g. "Facilities"
			"Operations":     svc.Operations,
			"ResponseModels": ctx.ResponseModels,
		}

		// Create file
		f, err := os.Create(filepath.Join(svcDir, fmt.Sprintf("%sService.kt", name)))
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

// Helpers adapted from Java generator

func kotlinTypeMapper(s *schema.Schema, parentName, propName string) string {
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

func safeKotlinField(name string) string {
	return app.SafeFieldName(app.ToCamelCase(name), "kotlin")
}
