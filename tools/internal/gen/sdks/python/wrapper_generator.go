package python

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

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
	return "Python Wrapper"
}

func (g *WrapperGenerator) Generate(ctx *common.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-py", "wrapper")

	// Collect sorted model list for deterministic output
	var allModels []*jsonschema.ModelIR
	seen := make(map[string]bool)
	collect := func(m *jsonschema.ModelIR) {
		if m == nil || seen[m.Name] {
			return
		}
		seen[m.Name] = true
		// Only collect nested types if not already seen? Not recurse function exists?
		// Assuming ResponseModels are fully formed.
		allModels = append(allModels, m)
	}

	for _, m := range ctx.ResponseModels {
		collect(m)
	}

	// Collect Request Models (Deduplicated)
	reqModels := app.CollectRequestModels(ctx.Services, pythonTypeMapper, brunoKindMapper, func(_, opName string) string {
		return opName + "Request"
	})
	for _, m := range reqModels {
		if !seen[m.Name] {
			allModels = append(allModels, m)
			seen[m.Name] = true
		}
	}

	data := map[string]interface{}{
		"Version":        ctx.Version,
		"Services":       ctx.Services,
		"ResponseModels": allModels,
	}

	formatMethodName := func(opName, group string) string {
		return app.ToSnakeCase(app.FormatMethodName(opName, group))
	}

	funcs := app.GetBaseTemplateFuncs()
	funcs["ResolveType"] = resolveType
	funcs["FormatMethodName"] = formatMethodName
	// ... existing funcs ...
	// Need to duplicate GetReturnType and ToPythonType from existing file logic or refactor
	// Since we are replacing the whole function block, I will re-include them inline or rely on existing helpers if extracted.
	// The previous file content shows them defined inside Generate. I should preserve them.

	funcs["GetReturnType"] = func(op models.Operation) string {
		if op.ResponseType == "" {
			return "Any"
		}
		if _, ok := ctx.ResponseModels[op.ResponseType]; !ok {
			return "Any"
		}
		typeName := op.ResponseType
		if op.IsPaginated {
			return fmt.Sprintf("PaginatedResponse['%s']", typeName)
		}
		if op.IsArray {
			return fmt.Sprintf("List['%s']", typeName)
		}
		return fmt.Sprintf("'%s'", typeName)
	}

	// Helper to resolve Python types from ModelIR properties
	var toPythonType func(t jsonschema.TypeRef) string
	toPythonType = func(t jsonschema.TypeRef) string {
		switch t.Kind {
		case jsonschema.TypeString:
			return "str"
		case jsonschema.TypeInteger:
			return "int"
		case jsonschema.TypeNumber:
			return "float"
		case jsonschema.TypeBoolean:
			return "bool"
		case jsonschema.TypeArray:
			if t.ItemType != nil {
				return fmt.Sprintf("List[%s]", toPythonType(*t.ItemType))
			}
			return "List[Any]"
		case jsonschema.TypeObject:
			if t.Name != "" {
				return fmt.Sprintf("'%s'", t.Name)
			}
			return "Any"
		default:
			if t.Name != "" {
				return fmt.Sprintf("'%s'", t.Name)
			}
			return "Any"
		}
	}
	funcs["ToPythonType"] = toPythonType

	if err := app.RenderPyProjectToml(filepath.Join(outDir, "pyproject.toml"), app.PyProjectConfig{
		Version:     ctx.Version,
		Description: "Type-safe wrapper for ThunkMetrc Python client",
		IsWrapper:   true,
		Deps:        ctx.Deps,
	}); err != nil {
		return err
	}

	if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/README.md.tmpl", filepath.Join(outDir, "README.md"), data, funcs); err != nil {
		return err
	}

	srcDir := filepath.Join(outDir, "src", "thunkmetrc", "wrapper")
	servicesDir := filepath.Join(srcDir, "services")
	if err := os.MkdirAll(servicesDir, 0755); err != nil {
		return err
	}

	if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/__init__.py.tmpl", filepath.Join(srcDir, "__init__.py"), data, funcs); err != nil {
		return err
	}

	if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/ratelimiter.py.tmpl", filepath.Join(srcDir, "ratelimiter.py"), data, funcs); err != nil {
		return err
	}

	if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/factory.py.tmpl", filepath.Join(srcDir, "factory.py"), data, funcs); err != nil {
		return err
	}

	if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/utils.py.tmpl", filepath.Join(srcDir, "utils.py"), data, funcs); err != nil {
		return err
	}

	// models directory
	modelsDir := filepath.Join(srcDir, "models")
	if err := os.MkdirAll(modelsDir, 0755); err != nil {
		return err
	}

	// models/__init__.py
	if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/models_init.py.tmpl", filepath.Join(modelsDir, "__init__.py"), map[string]interface{}{
		"Models": allModels,
	}, funcs); err != nil {
		return err
	}

	// models/paginated_response.py
	if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/paginated_response.py.tmpl", filepath.Join(modelsDir, "paginated_response.py"), nil, funcs); err != nil {
		return err
	}

	// Individual model files
	for _, m := range allModels {
		filename := app.ToSnakeCase(m.Name) + ".py"
		if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/model.py.tmpl", filepath.Join(modelsDir, filename), m, funcs); err != nil {
			return err
		}
	}

	// wrapper.py (Facade)
	if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/wrapper.py.tmpl", filepath.Join(srcDir, "wrapper.py"), data, funcs); err != nil {
		return err
	}

	// Services
	if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/services_init.py.tmpl", filepath.Join(servicesDir, "__init__.py"), data, funcs); err != nil {
		return err
	}

	for _, svc := range ctx.Services {
		safeGroupSnake := app.ToSnakeCase(app.CleanName(svc.Name))
		if err := app.RenderTemplate(wrapperTemplatesFS, "templates/wrapper/service.py.tmpl", filepath.Join(servicesDir, fmt.Sprintf("%s_service.py", safeGroupSnake)), map[string]interface{}{
			"Service": svc,
		}, funcs); err != nil {
			return err
		}
	}

	return nil
}

func pythonTypeMapper(s *schema.Schema, parentName, propName string) string {
	if s.Type == schema.TypeObject {
		return parentName + propName
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

func resolveType(s *schema.Schema, parentName, propName string) string {
	if s == nil {
		return "Any"
	}
	switch s.Type {
	case "string":
		return "str"
	case "integer", "int":
		return "int"
	case "number", "float":
		return "float"
	case "boolean", "bool":
		return "bool"
	case "array":
		innerType := "Any"
		if s.Items != nil {
			innerType = resolveType(s.Items, parentName, propName)
		}
		return "List[" + innerType + "]"
	case "object":
		return parentName + propName
	default:
		return "Any"
	}
}
