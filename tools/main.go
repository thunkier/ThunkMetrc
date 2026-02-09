package main

import (
	"fmt"
	"os"

	"github.com/thunkier/thunkmetrc/tools/internal/commands"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "gen":
		handleGen(args)
	case "mock":
		handleMock(args)
	case "analyze":
		commands.RunAnalyze(args)

	default:
		fmt.Printf("Unknown command: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func handleGen(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: thunk gen <subcommand> [flags]")
		fmt.Println("Subcommands: specs, schemas, internal, sdks")
		os.Exit(1)
	}

	sub := args[0]
	subArgs := args[1:]

	switch sub {
	case "specs":
		commands.RunSpecGen(subArgs)
	case "schemas":
		commands.RunSchemaGen(subArgs)
	case "internal":
		commands.RunInternalGen(subArgs)
	case "sdks":
		commands.RunSDKGen(subArgs)
	default:
		fmt.Printf("Unknown gen subcommand: %s\n", sub)
		os.Exit(1)
	}
}

func handleMock(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: thunk mock <subcommand> [flags]")
		fmt.Println("Subcommands: server, gen")
		os.Exit(1)
	}

	sub := args[0]
	subArgs := args[1:]

	switch sub {
	case "server":
		commands.RunMockServer(subArgs)
	case "gen":
		commands.RunMockGen(subArgs)
	default:
		fmt.Printf("Unknown mock subcommand: %s\n", sub)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: thunk <command> <subcommand> [flags]")
	fmt.Println("Commands:")
	fmt.Println("  gen      Generate specs, schemas, models, SDKs")
	fmt.Println("  mock     Run mock server or generate handlers")
	fmt.Println("  analyze  Analyze response coverage")

}
