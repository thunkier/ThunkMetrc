package typescript

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	app "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
)

//go:embed templates/main/*.tmpl
var mainTemplatesFS embed.FS

type MainGenerator struct{}

func (g *MainGenerator) Name() string {
	return "TypeScriptMain"
}

func (g *MainGenerator) Generate(ctx *app.GeneratorContext) error {
	_ = ctx

	outDir := filepath.Join(ctx.OutputDir, "thunkmetrc-ts", "thunkmetrc")
	extDir := filepath.Join(outDir, "src", "extensions")

	// Clear stale extension outputs. The current extension template set is not
	// compatible with the split client/wrapper package layout.
	if err := os.RemoveAll(extDir); err != nil {
		return fmt.Errorf("failed to clean extensions dir: %w", err)
	}
	if err := os.MkdirAll(extDir, 0755); err != nil {
		return fmt.Errorf("failed to create extensions dir: %w", err)
	}

	return nil
}

func hasPagination(op models.Operation) bool {
	// reuse logic from csharp/main_generator.go or similar
	// Check if Last params exist
	hasPage := false
	for _, p := range op.QueryParams {
		if p.Name == "pageNumber" || p.Name == "page" {
			hasPage = true
			break
		}
	}
	return hasPage
}

func isTimeWindowed(op models.Operation) bool {
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
