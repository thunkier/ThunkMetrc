package commands

import (
	"flag"
	"os"

	"github.com/thunkier/thunkmetrc/tools/internal/gen/mock"
)

func RunMockGen(args []string) {
	fs := flag.NewFlagSet("mockgen", flag.ExitOnError)
	specsDir := fs.String("specs", "../specs/generated/bruno", "Path to Bruno specs directory")
	outputFile := fs.String("out", "../mockserver/pkg/server/handlers/generated.go", "Output file path")

	if err := fs.Parse(args); err != nil {
		fs.Usage()
		os.Exit(1)
	}

	cfg := mock.Config{
		SpecsDir:   *specsDir,
		OutputFile: *outputFile,
	}

	mock.Run(cfg)
}
