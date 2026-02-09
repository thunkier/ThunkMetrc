package commands

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/thunkier/thunkmetrc/tools/internal/analysis/coverage"
	"github.com/thunkier/thunkmetrc/tools/internal/analysis/permissions"
)

func RunAnalyze(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: thunk analyze <subcommand> [flags]")
		fmt.Println("Subcommands:")
		fmt.Println("  permissions  Analyze facility permissions from responses")
		fmt.Println("  coverage     Analyze response coverage against specs")
		os.Exit(1)
	}

	sub := args[0]
	flags := args[1:]

	switch sub {
	case "permissions":
		permissions.Run()
	case "coverage":
		runCoverage(flags)
	default:

		if len(args) > 0 && args[0][0] != '-' {
			fmt.Printf("Unknown analyze subcommand: %s\n", sub)
			os.Exit(1)
		}
		runCoverage(args)
	}
}

func runCoverage(args []string) {
	fs := flag.NewFlagSet("coverage", flag.ExitOnError)
	manifestPath := fs.String("manifest", "../specs/generated/schemas/public/manifest.json", "Path to schema manifest")
	responsesDir := fs.String("responses", "../specs/source/captured/responses", "Directory containing captured response JSON files")
	outPath := fs.String("out", "../docs/API_COVERAGE.md", "Output report path")

	if err := fs.Parse(args); err != nil {
		fs.Usage()
		os.Exit(1)
	}

	if err := coverage.Run(*manifestPath, *responsesDir, *outPath); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
