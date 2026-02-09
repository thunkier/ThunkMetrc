package csharp

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
	return "CSharpWrapper"
}

func (g *WrapperGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-cs", "ThunkMetrc.Wrapper")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	// .csproj
	if err := generateWrapperCsproj(outDir, ctx.Version, ctx.Deps); err != nil {
		return err
	}

	// RateLimiter
	if err := generateRateLimiter(outDir); err != nil {
		return err
	}

	// Factory
	if err := generateFactory(outDir); err != nil {
		return err
	}

	// Wrapper Code & Model Classes
	if err := generateWrapperCode(outDir, ctx.Services, ctx.ResponseModels); err != nil {
		return err
	}

	// Extensions (High-Level Utils)
	if err := generateExtensions(outDir); err != nil {
		return err
	}

	// README.md
	if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/README.md.tmpl", filepath.Join(outDir, "README.md"), map[string]interface{}{
		"Version": ctx.Version,
	}, nil); err != nil {
		return err
	}

	return nil
}

func generateWrapperCsproj(dir, version string, deps map[string]string) error {
	dependencies := make([]map[string]string, 0)
	prefix := "csharp.wrapper."
	for k, v := range deps {
		if strings.HasPrefix(k, prefix) {
			dependencies = append(dependencies, map[string]string{
				"Name":    strings.TrimPrefix(k, prefix),
				"Version": v,
			})
		}
	}

	data := map[string]interface{}{
		"Version":     version,
		"PackageId":   "ThunkMetrc.Wrapper",
		"Description": "Type-safe wrapper for ThunkMetrc C# client.",
		"Tags":        "ThunkMetrc;Cannabis;Compliance;API;SDK;Wrapper",
		"ProjectReferences": []string{
			"..\\ThunkMetrc.Client\\ThunkMetrc.Client.csproj",
		},
		"Dependencies": dependencies,
	}
	return app.RenderTemplate(app.TemplatesFS, "templates/csproj.xml.tmpl", filepath.Join(dir, "ThunkMetrc.Wrapper.csproj"), data, nil)
}

func generateRateLimiter(dir string) error {
	return app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/MetrcRateLimiter.cs.tmpl", filepath.Join(dir, "MetrcRateLimiter.cs"), nil, nil)
}

func generateFactory(dir string) error {
	return app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/MetrcFactory.cs.tmpl", filepath.Join(dir, "MetrcFactory.cs"), nil, nil)
}

func generateExtensions(dir string) error {
	return app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/Extensions.cs.tmpl", filepath.Join(dir, "MetrcExtensions.cs"), nil, nil)
}

func generateWrapperCode(dir string, services []models.Service, responseModels map[string]*jsonschema.ModelIR) error {
	modelsFuncs := app.GetBaseTemplateFuncs()
	modelsFuncs["ResolveType"] = resolveCSharpType

	modelsFuncs["SafeCSharpField"] = safeCSharpField

	reqModels := app.CollectRequestModels(services, csharpTypeMapper, brunoKindMapper, func(svcName, opName string) string {
		safeGroup := app.ToPascalCase(app.CleanName(svcName))
		safeOpName := app.ToPascalCase(app.CleanName(opName))
		return safeGroup + safeOpName + "Request"
	})

	var modelList []*jsonschema.ModelIR
	existingModels := make(map[string]bool)
	for _, m := range responseModels {
		modelList = append(modelList, m)
		existingModels[m.Name] = true
	}

	for _, m := range reqModels {
		if !existingModels[m.Name] {
			modelList = append(modelList, m)
			existingModels[m.Name] = true
		}
	}

	svcDir := filepath.Join(dir, "Services")
	if err := os.MkdirAll(svcDir, 0755); err != nil {
		return err
	}

	wrapperFuncs := app.GetBaseTemplateFuncs()
	wrapperFuncs["GetPathParams"] = app.GetPathParams
	wrapperFuncs["ResolveType"] = resolveCSharpType
	wrapperFuncs["IsVoid"] = func(method string) bool { return method != "GET" }
	wrapperFuncs["WrapperSignatureArgs"] = wrapperSignatureArgs
	wrapperFuncs["ClientCallArgs"] = clientCallArgs
	wrapperFuncs["HasPagination"] = hasPagination
	wrapperFuncs["PaginationSignature"] = paginationSignatureArgs
	wrapperFuncs["PaginationClientCall"] = paginationClientCallArgs
	wrapperFuncs["WrapperParamsDocs"] = wrapperParamsDocs
	wrapperFuncs["PaginationParamsDocs"] = paginationParamsDocs
	ResolveResponseType := func(op models.Operation) string {
		if op.ResponseType != "" {
			if _, exists := responseModels[op.ResponseType]; exists {
				if op.IsPaginated {
					return fmt.Sprintf("PaginatedResponse<%s>", op.ResponseType)
				}
				if op.IsArray {
					return fmt.Sprintf("List<%s>", op.ResponseType)
				}
				return op.ResponseType
			}
		}
		return "object"
	}
	wrapperFuncs["ResolveResponseType"] = ResolveResponseType

	wrapperFuncs["GetIteratorItemType"] = func(op models.Operation) string {
		if op.ResponseType != "" {
			if _, exists := responseModels[op.ResponseType]; exists {
				return op.ResponseType
			}
		}
		return "JsonElement"
	}
	wrapperFuncs["FormatMethodName"] = app.FormatMethodName

	modelsDir := filepath.Join(dir, "Models")
	if err := os.MkdirAll(modelsDir, 0755); err != nil {
		return err
	}

	if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/PaginatedResponse.cs.tmpl", filepath.Join(modelsDir, "PaginatedResponse.cs"), nil, nil); err != nil {
		return err
	}

	for _, m := range modelList {
		if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/Model.cs.tmpl", filepath.Join(modelsDir, m.Name+".cs"), map[string]interface{}{
			"Model": m,
		}, modelsFuncs); err != nil {
			return err
		}
	}

	for _, svc := range services {
		if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/Service.cs.tmpl", filepath.Join(svcDir, fmt.Sprintf("%sService.cs", app.ToPascalCase(app.CleanName(svc.Name)))), map[string]interface{}{
			"Service": svc,
		}, wrapperFuncs); err != nil {
			return err
		}
	}

	return app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/MetrcWrapper.cs.tmpl", filepath.Join(dir, "MetrcWrapper.cs"), map[string]interface{}{
		"Services": services,
	}, wrapperFuncs)
}

func csharpTypeMapper(s *schema.Schema, parentName, propName string) string {
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

func safeCSharpField(name string) string {
	return app.SafeFieldName(app.ToPascalCase(name), "csharp")
}

func resolveCSharpType(s *schema.Schema, parentName, propName string) string {
	if s == nil {
		return "object?"
	}
	switch s.Type {
	case "string":
		return "string?"
	case "integer", "int":
		return "int"
	case "number", "float":
		return "double"
	case "boolean", "bool":
		return "bool"
	case "array":
		innerType := "object"
		if s.Items != nil {
			innerType = resolveCSharpType(s.Items, parentName, propName)
		}
		return fmt.Sprintf("List<%s>?", innerType)
	case "object":
		return parentName + app.ToPascalCase(propName) + "?"
	default:
		return "object?"
	}
}

func wrapperSignatureArgs(op models.Operation, bodyType string) string {
	var args []string
	for _, p := range op.PathParams {
		args = append(args, fmt.Sprintf("string %s", p))
	}
	if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
		if bodyType == "" {
			bodyType = "object?"
		}
		args = append(args, fmt.Sprintf("%s body = null", bodyType))
	}
	for _, p := range op.QueryParams {
		name := strings.TrimSuffix(p.Name, "optional")
		args = append(args, fmt.Sprintf("string? %s = null", name))
	}
	if len(args) > 0 {
		return strings.Join(args, ", ") + ", "
	}
	return ""
}

func clientCallArgs(op models.Operation) string {
	var args []string
	for _, p := range op.PathParams {
		args = append(args, p)
	}
	if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
		args = append(args, "body")
	}
	for _, p := range op.QueryParams {
		name := strings.TrimSuffix(p.Name, "optional")
		args = append(args, name)
	}
	return strings.Join(args, ", ")
}

func hasPagination(op models.Operation) bool {
	for _, p := range op.QueryParams {
		if p.Name == "pageNumber" || p.Name == "page" {
			return true
		}
	}
	return false
}

func paginationSignatureArgs(op models.Operation, bodyType string) string {
	var args []string
	for _, p := range op.PathParams {
		args = append(args, fmt.Sprintf("string %s", p))
	}
	if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
		if bodyType == "" {
			bodyType = "object?"
		}
		args = append(args, fmt.Sprintf("%s body = null", bodyType))
	}
	for _, p := range op.QueryParams {
		name := strings.TrimSuffix(p.Name, "optional")
		if name == "pageNumber" || name == "page" {
			continue
		}
		args = append(args, fmt.Sprintf("string? %s = null", name))
	}
	if len(args) > 0 {
		return strings.Join(args, ", ") + ", "
	}
	return ""
}

func paginationClientCallArgs(op models.Operation) string {
	var args []string
	for _, p := range op.PathParams {
		args = append(args, p)
	}
	if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
		args = append(args, "body")
	}
	for _, p := range op.QueryParams {
		name := strings.TrimSuffix(p.Name, "optional")
		if name == "pageNumber" {
			args = append(args, "page.ToString()")
		} else {
			args = append(args, name)
		}
	}
	return strings.Join(args, ", ")
}

func wrapperParamsDocs(op models.Operation) string {
	paramDocs := app.ParseParamDocs(op.Description)
	var lines []string
	for _, p := range op.PathParams {
		desc := paramDocs[p]
		if desc == "" {
			desc = p
		}
		lines = append(lines, fmt.Sprintf("/// <param name=\"%s\">%s</param>", p, desc))
	}
	if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
		lines = append(lines, "/// <param name=\"body\">Request body.</param>")
	}
	for _, p := range op.QueryParams {
		name := strings.TrimSuffix(p.Name, "optional")
		desc := paramDocs[name]
		if desc == "" {
			desc = "Optional query parameter."
		}
		lines = append(lines, fmt.Sprintf("/// <param name=\"%s\">%s</param>", name, desc))
	}
	return strings.Join(lines, "\n        ")
}

func paginationParamsDocs(op models.Operation) string {
	paramDocs := app.ParseParamDocs(op.Description)
	var lines []string
	for _, p := range op.PathParams {
		desc := paramDocs[p]
		if desc == "" {
			desc = p
		}
		lines = append(lines, fmt.Sprintf("/// <param name=\"%s\">%s</param>", p, desc))
	}
	if op.Method == "POST" || op.Method == "PUT" || op.Method == "PATCH" {
		lines = append(lines, "/// <param name=\"body\">Request body.</param>")
	}
	for _, p := range op.QueryParams {
		name := strings.TrimSuffix(p.Name, "optional")
		if name == "pageNumber" || name == "page" {
			continue
		}
		desc := paramDocs[name]
		if desc == "" {
			desc = "Optional query parameter."
		}
		lines = append(lines, fmt.Sprintf("/// <param name=\"%s\">%s</param>", name, desc))
	}
	return strings.Join(lines, "\n        ")
}

