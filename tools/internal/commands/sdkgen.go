package commands

import (
	"flag"
	"os"

	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks"
)

func RunSDKGen(args []string) {
	fs := flag.NewFlagSet("sdkgen", flag.ExitOnError)
	specsDir := fs.String("specs", "../specs/generated/bruno", "Path to Bruno specs directory")
	responsesDir := fs.String("responses", "../specs/source/captured/responses", "Path to API responses directory")
	outputDir := fs.String("out", "../sdks", "Path to output SDKs directory")

	if err := fs.Parse(args); err != nil {
		fs.Usage()
		os.Exit(1)
	}

	cfg := sdks.Config{
		SpecsDir:     *specsDir,
		ResponsesDir: *responsesDir,
		OutputDir:    *outputDir,
	}

	sdks.Run(cfg)
}
