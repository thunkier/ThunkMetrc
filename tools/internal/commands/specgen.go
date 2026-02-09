package commands

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/thunkier/thunkmetrc/tools/internal/gen/specgen"
)

func RunSpecGen(args []string) {
	fs := flag.NewFlagSet("specgen", flag.ExitOnError)
	skipScrape := fs.Bool("skip-scrape", false, "Skip downloading specs and generate files from existing local data")
	states := fs.String("states", "", "Comma-separated list of states/regions to scrape (overrides config)")
	timeoutSeconds := fs.Int("timeout-seconds", 0, "HTTP timeout in seconds (0 = use config/default)")
	maxRetries := fs.Int("max-retries", -1, "Max retries for HTTP fetches (-1 = use config/default)")
	retryBackoffMs := fs.Int("retry-backoff-ms", 0, "Retry backoff in milliseconds (0 = use config/default)")
	concurrency := fs.Int("concurrency", 0, "Concurrent state fetch workers (0 = use config/default)")
	userAgent := fs.String("user-agent", "", "HTTP User-Agent for scraping (empty = use config/default)")

	if err := fs.Parse(args); err != nil {
		fs.Usage()
		os.Exit(1)
	}

	fmt.Println("Starting Metrc Spec Generator...")
	if *skipScrape {
		fmt.Println("Skipping scrape phase...")
	}

	opts := specgen.Options{
		SkipScrape:      *skipScrape,
		StatesCSV:       *states,
		TimeoutSeconds:  *timeoutSeconds,
		MaxRetries:      *maxRetries,
		RetryBackoffMs:  *retryBackoffMs,
		Concurrency:     *concurrency,
		UserAgent:       *userAgent,
	}

	if err := specgen.Run(opts); err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Spec generation complete.")
}
