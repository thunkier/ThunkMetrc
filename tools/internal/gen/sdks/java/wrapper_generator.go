package java

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	app "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
	"github.com/thunkier/thunkmetrc/tools/pkg/jsonschema"
	"github.com/thunkier/thunkmetrc/tools/pkg/schema"
)

//go:embed templates/wrapper/*.tmpl
var wrapperTemplatesFS embed.FS

type WrapperGenerator struct{}

func (g *WrapperGenerator) Name() string {
	return "Java Wrapper"
}

func (g *WrapperGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-java", "wrapper")

	pkgDir := filepath.Join(outDir, "src", "main", "java", "io", "github", "thunkier", "thunkmetrc", "wrapper")
	modelDir := filepath.Join(pkgDir, "models")
	svcDir := filepath.Join(pkgDir, "services")

	if err := os.MkdirAll(modelDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(svcDir, 0755); err != nil {
		return err
	}

	// 1. Collect Requests Use Common
	reqModels := app.CollectRequestModels(ctx.Services, javaTypeMapper, brunoKindMapper, func(svcName, opName string) string {
		return app.ToPascalCase(app.CleanName(svcName)) + app.ToPascalCase(app.CleanName(opName)) + "Request"
	})

	// 2. recursive collect logic for flatten
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
	for _, m := range reqModels {
		collect(m)
	}

	data := map[string]interface{}{
		"Version":        ctx.Version,
		"Services":       ctx.Services,
		"ResponseModels": allModels,
	}

	formatMethodName := func(opName, group string) string {
		return app.ToCamelCase(app.FormatMethodName(opName, group))
	}

	funcs := app.GetBaseTemplateFuncs()
	funcs["SafeJavaField"] = safeJavaField

	funcs["FormatMethodName"] = formatMethodName

	// SafeJavaType logic
	safeJavaType := func(typeName string) string {
		if typeName == "Package" {
			return "io.github.thunkier.thunkmetrc.wrapper.models.Package"
		}
		return typeName
	}
	funcs["SafeJavaType"] = safeJavaType

	funcs["GetReturnType"] = func(op models.Operation) string {
		if op.ResponseType == "" {
			return "Object"
		}
		if _, ok := ctx.ResponseModels[op.ResponseType]; !ok {
			return "Object"
		}
		typeName := safeJavaType(op.ResponseType)

		if op.IsPaginated {
			return fmt.Sprintf("PaginatedResponse<%s>", typeName)
		}
		if op.IsArray {
			return fmt.Sprintf("List<%s>", typeName)
		}
		return typeName
	}

	// Helper to resolve Java types from ModelIR properties
	var toJavaType func(t jsonschema.TypeRef) string
	toJavaType = func(t jsonschema.TypeRef) string {
		switch t.Kind {
		case jsonschema.TypeString:
			return "String"
		case jsonschema.TypeInteger:
			return "Integer"
		case jsonschema.TypeNumber:
			return "Double"
		case jsonschema.TypeBoolean:
			return "Boolean"
		case jsonschema.TypeArray:
			if t.ItemType != nil {
				return fmt.Sprintf("List<%s>", toJavaType(*t.ItemType))
			}
			return "List<Object>"
		case jsonschema.TypeObject:
			if t.Name != "" {
				if t.Name == "Package" {
					return "io.github.thunkier.thunkmetrc.wrapper.models.Package"
				}
				return t.Name
			}
			return "Object"
		default:
			if t.Name != "" {
				if t.Name == "Package" {
					return "io.github.thunkier.thunkmetrc.wrapper.models.Package"
				}
				return t.Name
			}
			return "Object"
		}
	}
	funcs["ToJavaType"] = toJavaType
	funcs["HasPagination"] = func(ops []models.Operation) bool {
		for _, op := range ops {
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
		}
		return false
	}
	funcs["HasModels"] = func(ops []models.Operation) bool {
		for _, op := range ops {
			if op.IsPaginated {
				return true
			}
			for _, qp := range op.QueryParams {
				if strings.Contains(qp.Name, "pageNumber") {
					return true
				}
			}
			if op.ResponseType != "" && op.ResponseType != "Object" {
				return true
			}
			if op.BodySchema != nil {
				return true
			}
		}
		return false
	}
	funcs["HasList"] = func(ops []models.Operation) bool {
		for _, op := range ops {
			if op.IsArray && !op.IsPaginated {
				return true
			}
			if op.BodySchema != nil && op.BodySchema.Type == "array" {
				return true
			}
			if op.ResponseType != "" {
				if strings.HasPrefix(op.ResponseType, "List<") || strings.Contains(op.ResponseType, "List<") {
					return true
				}
			}
		}
		return false
	}

	// pom.xml
	if err := app.RenderPom(filepath.Join(outDir, "pom.xml"), app.PomConfig{
		ArtifactId:       "thunkmetrc-wrapper",
		Name:             "ThunkMetrc Wrapper",
		Description:      "Type-safe wrapper for ThunkMetrc Java client with rate limiting.",
		Version:          ctx.Version,
		IsKotlin:         false,
		IsWrapper:        true,
		ClientArtifactId: "thunkmetrc-client",
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

	// MetrcRateLimiter.java
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/MetrcRateLimiter.java.tmpl",
		filepath.Join(pkgDir, "MetrcRateLimiter.java"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// MetrcWrapper.java (Main Facade)
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/MetrcWrapper.java.tmpl",
		filepath.Join(pkgDir, "MetrcWrapper.java"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// MetrcFactory.java
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/MetrcFactory.java.tmpl",
		filepath.Join(pkgDir, "MetrcFactory.java"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// Metrc	// Extensions (Utility) Class
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/MetrcExtensions.java.tmpl",
		filepath.Join(pkgDir, "MetrcExtensions.java"),
		nil,
		funcs,
	); err != nil {
		return err
	}

	// Services
	for _, svc := range ctx.Services {
		svcName := app.ToPascalCase(app.CleanName(svc.Name)) // e.g., Facilities
		if svcName == "" {
			continue
		}

		svcData := map[string]interface{}{
			"Name":       svcName,
			"Operations": svc.Operations,
		}
		if err := app.RenderTemplate(
			wrapperTemplatesFS,
			"templates/wrapper/Service.java.tmpl",
			filepath.Join(svcDir, fmt.Sprintf("%sService.java", svcName)),
			svcData,
			funcs,
		); err != nil {
			return err
		}
	}

	// PaginatedResponse.java
	if err := app.RenderTemplate(
		wrapperTemplatesFS,
		"templates/wrapper/PaginatedResponse.java.tmpl",
		filepath.Join(modelDir, "PaginatedResponse.java"),
		data,
		funcs,
	); err != nil {
		return err
	}

	// Render all models (Response + Request)
	for _, m := range allModels {
		modelData := map[string]interface{}{
			"Model": m,
		}
		if err := app.RenderTemplate(
			wrapperTemplatesFS,
			"templates/wrapper/Model.java.tmpl",
			filepath.Join(modelDir, fmt.Sprintf("%s.java", m.Name)),
			modelData,
			funcs,
		); err != nil {
			return err
		}
	}

	return nil
}

func javaTypeMapper(s *schema.Schema, parentName, propName string) string {
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

func safeJavaField(name string) string {
	return app.SafeFieldName(app.ToCamelCase(name), "java")
}
