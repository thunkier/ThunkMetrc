package commands

import (
	"flag"
	"log"
	"os"

	"github.com/thunkier/thunkmetrc/tools/internal/gen/schemas"
)

func RunSchemaGen(args []string) {
	fs := flag.NewFlagSet("schema-gen", flag.ExitOnError)
	responsesDir := fs.String("responses", "../specs/source/captured/responses", "Directory containing response JSON files")
	manualDir := fs.String("manual", "../specs/source/manual/responses", "Directory containing manual response JSON files")
	brunoDir := fs.String("bruno", "../specs/generated/bruno", "Directory containing Bruno files")
	outputDir := fs.String("output", "../specs/generated/schemas/public", "Output directory for JSON Schemas")
	internalMode := fs.Bool("internal", false, "Generate internal schemas from Bruno files")

	if err := fs.Parse(args); err != nil {
		fs.Usage()
		os.Exit(1)
	}

	if *internalMode && *outputDir == "../specs/generated/schemas/public" {
		*outputDir = "../specs/generated/schemas/internal"
	}

	cmd := schemas.Config{
		ResponsesDir: *responsesDir,
		ManualDir:    *manualDir,
		BrunoDir:     *brunoDir,
		OutputDir:    *outputDir,
		InternalMode: *internalMode,
	}

	if err := schemas.Run(cmd); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
