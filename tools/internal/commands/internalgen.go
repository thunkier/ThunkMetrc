package commands

import (
	"flag"
	"log"
	"os"

	"github.com/thunkier/thunkmetrc/tools/internal/gen/internalgen"
)

func RunInternalGen(args []string) {
	fs := flag.NewFlagSet("internal-gen", flag.ExitOnError)
	brunoDir := fs.String("input", "../specs/generated/bruno", "Path to Bruno specs directory")
	schemasDir := fs.String("schemas", "../specs/generated/schemas/internal", "Path to internal schemas directory")
	modelsDir := fs.String("models", "../probe/pkg/metrc/models", "Path to output Go models directory")
	clientDir := fs.String("client-dir", "../probe/internal/metrc/client", "Path to output Go client directory")

	if err := fs.Parse(args); err != nil {
		fs.Usage()
		os.Exit(1)
	}

	cfg := internalgen.Config{
		BrunoDir:   *brunoDir,
		SchemasDir: *schemasDir,
		ModelsDir:  *modelsDir,
		ClientDir:  *clientDir,
	}

	if err := internalgen.Run(cfg); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
