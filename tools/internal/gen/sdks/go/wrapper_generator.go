package golang

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	app "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
	"github.com/thunkier/thunkmetrc/tools/pkg/jsonschema"
	"github.com/thunkier/thunkmetrc/tools/pkg/schema"
)

//go:embed templates/wrapper/*.tmpl
var wrapperTemplatesFS embed.FS

type WrapperGenerator struct{}

func (g *WrapperGenerator) Name() string {
	return "GoWrapper"
}

func (g *WrapperGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-go", "wrapper")
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	// go.mod
	if err := generateWrapperGoMod(outDir, ctx.Deps); err != nil {
		return err
	}

	// RateLimiter
	if err := generateWrapperRateLimiter(outDir); err != nil {
		return err
	}

	// Factory
	if err := generateWrapperFactory(outDir); err != nil {
		return err
	}

	// Wrapper Code & Models
	if err := generateWrapperCode(outDir, ctx.Services, ctx.ResponseModels); err != nil {
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

func generateWrapperGoMod(dir string, deps map[string]string) error {
	timeVersion := deps["go.wrapper.time"]
	if timeVersion == "" {
		timeVersion = "v0.5.0"
	}

	data := map[string]interface{}{
		"Module": "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper",
		"Require": []map[string]interface{}{
			{"Path": "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client", "Version": "v0.0.0"},
			{"Path": "golang.org/x/time", "Version": timeVersion},
		},
		"Replace": []map[string]interface{}{
			{"Old": "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client", "New": "../client"},
		},
	}
	return app.RenderTemplate(app.TemplatesFS, "templates/go.mod.tmpl", filepath.Join(dir, "go.mod"), data, nil)
}

func generateWrapperRateLimiter(dir string) error {
	tmpl, err := template.ParseFS(wrapperTemplatesFS, "templates/wrapper/ratelimiter.go.tmpl")
	if err != nil {
		return err
	}

	servicesDir := filepath.Join(dir, "services")
	if err := os.MkdirAll(servicesDir, 0755); err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(servicesDir, "ratelimiter.go"))
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, nil)
}

func generateWrapperFactory(dir string) error {
	tmpl, err := template.ParseFS(wrapperTemplatesFS, "templates/wrapper/factory.go.tmpl")
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(dir, "factory.go"))
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, nil)
}

func generateWrapperCode(dir string, services []models.Service, responseModels map[string]*jsonschema.ModelIR) error {
	goModels := collectModels(services, responseModels)

	funcs := template.FuncMap{
		"ToPascalCase":     app.ToPascalCase,
		"ToCamelCase":      app.ToCamelCase,
		"CleanName":        app.CleanName,
		"CleanDocs":        app.CleanDocs,
		"GetPathParams":    app.GetPathParams,
		"TrimSuffix":       func(suffix, s string) string { return strings.TrimSuffix(s, suffix) },
		"TrimPrefix":       func(prefix, s string) string { return strings.TrimPrefix(s, prefix) },
		"Split":            strings.Split,
		"Contains":         strings.Contains,
		"HasPrefix":        strings.HasPrefix,
		"join":             strings.Join,
		"list":             func(values ...string) []string { return values },
		"append":           func(l []string, v string) []string { return append(l, v) },
		"GetBodyType":      getBodyType,
		"FormatMethodName": app.FormatMethodName,
		"GetReturnType": func(op models.Operation) string {
			if op.ResponseType == "" {
				return "interface{}"
			}

			if _, exists := responseModels[op.ResponseType]; !exists {
				return "interface{}"
			}

			typeName := "*models." + op.ResponseType

			if op.IsPaginated {
				return fmt.Sprintf("*models.PaginatedResponse[%s]", typeName)
			}
			if op.IsArray {
				return fmt.Sprintf("[]%s", typeName)
			}
			return typeName
		},
	}

	modelsTmpl, err := template.New("model.go.tmpl").Funcs(funcs).ParseFS(wrapperTemplatesFS, "templates/wrapper/model.go.tmpl")
	if err != nil {
		return err
	}

	modelsDir := filepath.Join(dir, "models")
	if err := os.MkdirAll(modelsDir, 0755); err != nil {
		return err
	}

	for _, m := range goModels {
		f, err := os.Create(filepath.Join(modelsDir, fmt.Sprintf("%s.go", strings.ToLower(app.ToSnakeCase(m.Name)))))
		if err != nil {
			return err
		}

		if err := modelsTmpl.Execute(f, m); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}

	paginatedTmpl, err := template.New("paginated_response.go.tmpl").Funcs(funcs).ParseFS(wrapperTemplatesFS, "templates/wrapper/paginated_response.go.tmpl")
	if err != nil {
		return err
	}
	fPaginated, err := os.Create(filepath.Join(modelsDir, "paginated_response.go"))
	if err != nil {
		return err
	}
	defer fPaginated.Close()
	if err := paginatedTmpl.Execute(fPaginated, nil); err != nil {
		return err
	}

	utilsTmpl, err := template.New("utils.go.tmpl").Funcs(funcs).ParseFS(wrapperTemplatesFS, "templates/wrapper/utils.go.tmpl")
	if err != nil {
		return err
	}
	fUtils, err := os.Create(filepath.Join(modelsDir, "utils.go"))
	if err != nil {
		return err
	}
	defer fUtils.Close()
	if err := utilsTmpl.Execute(fUtils, nil); err != nil {
		return err
	}

	wrapperTmpl, err := template.New("wrapper.go.tmpl").Funcs(funcs).ParseFS(wrapperTemplatesFS, "templates/wrapper/wrapper.go.tmpl")
	if err != nil {
		return err
	}

	wrapperFile, err := os.Create(filepath.Join(dir, "wrapper.go"))
	if err != nil {
		return err
	}
	defer wrapperFile.Close()

	if err := wrapperTmpl.Execute(wrapperFile, map[string]interface{}{
		"Services":       services,
		"ResponseModels": responseModels,
	}); err != nil {
		return err
	}

	// Render Services
	servicesDir := filepath.Join(dir, "services")
	if err := os.MkdirAll(servicesDir, 0755); err != nil {
		return err
	}

	svcTmpl, err := template.New("service.go.tmpl").Funcs(funcs).ParseFS(wrapperTemplatesFS, "templates/wrapper/service.go.tmpl")
	if err != nil {
		return err
	}

	for _, svc := range services {
		name := app.ToPascalCase(app.CleanName(svc.Name))
		if name == "" {
			continue
		}

		f, err := os.Create(filepath.Join(servicesDir, fmt.Sprintf("generated_%s.go", strings.ToLower(name))))
		if err != nil {
			return err
		}

		err = svcTmpl.Execute(f, map[string]interface{}{
			"Name":           name,
			"Operations":     svc.Operations,
			"ResponseModels": responseModels,
			"Services":       services,
		})
		f.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

type GoModel struct {
	Name   string
	Fields []GoField
}

type GoField struct {
	Name string
	Type string
	Tag  string
}

var knownModels = make(map[string]bool)
var allModels []GoModel
var operationBodyTypes = make(map[string]string)

func collectModels(services []models.Service, responseModels map[string]*jsonschema.ModelIR) []GoModel {
	knownModels = make(map[string]bool)
	allModels = make([]GoModel, 0)
	operationBodyTypes = make(map[string]string)

	for _, svc := range services {
		svcName := app.ToPascalCase(app.CleanName(svc.Name))
		for _, op := range svc.Operations {
			if op.BodySchema != nil {
				opName := app.ToPascalCase(app.CleanName(op.Name))
				rootName := svcName + opName + "Request"
				typeStr := walkSchema(op.BodySchema, rootName)
				operationBodyTypes[svcName+opName] = typeStr
			}
		}
	}

	for _, model := range responseModels {
		processResponseModel(model)
	}

	sort.Slice(allModels, func(i, j int) bool {
		return allModels[i].Name < allModels[j].Name
	})

	return allModels
}

func processResponseModel(ir *jsonschema.ModelIR) {
	if ir == nil {
		return
	}

	if knownModels[ir.Name] {
		return
	}

	for i := range ir.NestedTypes {
		processResponseModel(&ir.NestedTypes[i])
	}

	if knownModels[ir.Name] {
		return
	}

	fields := make([]GoField, 0)
	for _, prop := range ir.Properties {
		goType := prop.Type.ToGo()

		fields = append(fields, GoField{
			Name: app.ToPascalCase(prop.Name),
			Type: goType,
			Tag:  fmt.Sprintf("`json:\"%s,omitempty\"`", prop.JSONName),
		})
	}

	sort.Slice(fields, func(i, j int) bool {
		return fields[i].Name < fields[j].Name
	})

	allModels = append(allModels, GoModel{
		Name:   ir.Name,
		Fields: fields,
	})
	knownModels[ir.Name] = true
}

func walkSchema(s *schema.Schema, name string) string {
	if s == nil {
		return "interface{}"
	}

	if knownModels[name] {
		return "*" + name
	}

	switch s.Type {
	case "object":
		if knownModels[name] {
			return "*" + name
		}

		fields := make([]GoField, 0)
		var keys []string
		for k := range s.Properties {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, propName := range keys {
			propSchema := s.Properties[propName]
			fieldName := app.ToPascalCase(propName)
			childName := name + fieldName
			fieldType := walkSchema(propSchema, childName)

			fields = append(fields, GoField{
				Name: fieldName,
				Type: fieldType,
				Tag:  fmt.Sprintf("`json:\"%s,omitempty\"`", propName),
			})
		}

		allModels = append(allModels, GoModel{
			Name:   name,
			Fields: fields,
		})
		knownModels[name] = true
		return "*" + name

	case "array":
		if s.Items != nil {
			itemName := name + "Item"
			if strings.HasSuffix(name, "s") {
				itemName = strings.TrimSuffix(name, "s")
			} else {
				itemName = name + "Item"
			}

			innerType := walkSchema(s.Items, itemName)
			return "[]" + innerType
		}
		return "[]interface{}"
	case "integer", "int":
		return "int"
	case "number", "float":
		return "float64"
	case "boolean", "bool":
		return "bool"
	case "string":
		return "string"
	default:
		return "interface{}"
	}
}

func getBodyType(svcName, opName string) string {
	key := app.ToPascalCase(app.CleanName(svcName)) + app.ToPascalCase(app.CleanName(opName))
	if t, ok := operationBodyTypes[key]; ok {
		if strings.HasPrefix(t, "*") {
			return "*models." + t[1:]
		}
		if strings.HasPrefix(t, "[]*") {
			return "[]*models." + t[3:]
		}
		return t
	}
	return "interface{}"
}
