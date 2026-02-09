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

//go:embed templates/main/ServiceExtensions.kt.tmpl
var mainTemplatesFS embed.FS

type MainGenerator struct{}

func (g *MainGenerator) Name() string {
	return "Kotlin Main"
}

func (g *MainGenerator) Generate(ctx *app.GeneratorContext) error {
	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-kotlin", "thunkmetrc")
	pkgDir := filepath.Join(outDir, "src", "main", "kotlin", "io", "github", "thunkier", "thunkmetrc")

	if err := os.MkdirAll(pkgDir, 0755); err != nil {
		return err
	}

	funcs := app.GetBaseTemplateFuncs()

	formatMethodName := func(opName, group string) string {
		return app.ToCamelCase(app.FormatMethodName(opName, group))
	}
	funcs["FormatMethodName"] = formatMethodName

	funcs["SafeKotlinField"] = func(name string) string {
		return app.SafeFieldName(app.ToCamelCase(name), "kotlin")
	}

	funcs["GetReturnType"] = func(op models.Operation) string {
		if op.ResponseType == "" {
			return "Any"
		}
		if _, ok := ctx.ResponseModels[op.ResponseType]; !ok {
			return "Any"
		}
		typeName := op.ResponseType
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

	// String helpers
	funcs["TrimSuffix"] = func(suffix, s string) string { return strings.TrimSuffix(s, suffix) }
	funcs["TrimPrefix"] = func(prefix, s string) string { return strings.TrimPrefix(s, prefix) }
	funcs["Contains"] = strings.Contains
	funcs["join"] = strings.Join
	funcs["list"] = func(values ...string) []string { return values }
	funcs["append"] = func(l []string, v string) []string { return append(l, v) }

	tmpl, err := template.New("ServiceExtensions.kt.tmpl").Funcs(funcs).ParseFS(mainTemplatesFS, "templates/main/ServiceExtensions.kt.tmpl")
	if err != nil {
		return err
	}

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

		svcName := app.ToPascalCase(app.CleanName(svc.Name))
		if svcName == "" {
			continue
		}

		f, err := os.Create(filepath.Join(pkgDir, fmt.Sprintf("%sExtensions.kt", svcName)))
		if err != nil {
			return err
		}

		if err := tmpl.Execute(f, map[string]interface{}{
			"Service": svc,
			"Name":    svcName,
		}); err != nil {
			f.Close()
			return err
		}
		f.Close()
	}

	return nil
}
