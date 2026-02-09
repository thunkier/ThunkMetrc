package specgen

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/thunkier/thunkmetrc/tools/internal/app"
)

type Options struct {
	SkipScrape     bool
	StatesCSV      string
	TimeoutSeconds int
	MaxRetries     int
	RetryBackoffMs int
	Concurrency    int
	UserAgent      string
}

type scrapeOptions struct {
	States       []string
	Timeout      time.Duration
	MaxRetries   int
	RetryBackoff time.Duration
	Concurrency  int
	UserAgent    string
}

func Run(opts Options) error {
	htmlDir := "../specs/source/scraped/html"
	postmanDir := "../specs/source/scraped/postman"
	brunoDir := "../specs/generated/bruno"

	cfg, err := app.LoadFileConfig("config.json")
	if err != nil {
		fmt.Printf("Warning: Failed to load config: %v. Using defaults.\n", err)
	}

	scrapeCfg := resolveScrapeOptions(opts, cfg.Scraper)
	if len(scrapeCfg.States) == 0 {
		return fmt.Errorf("no states configured for scraping")
	}

	if !opts.SkipScrape {
		if err := os.RemoveAll(htmlDir); err != nil {
			return fmt.Errorf("failed to clean html scrape dir: %w", err)
		}
		if err := os.RemoveAll(postmanDir); err != nil {
			return fmt.Errorf("failed to clean postman scrape dir: %w", err)
		}
	}

	if err := os.MkdirAll(htmlDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(postmanDir, 0755); err != nil {
		return err
	}

	if err := os.RemoveAll(brunoDir); err != nil {
		return fmt.Errorf("failed to clean bruno dir: %w", err)
	}
	if err := os.MkdirAll(brunoDir, 0755); err != nil {
		return err
	}

	createBrunoConfig(brunoDir)

	if !opts.SkipScrape {
		fetcher := NewHTTPFetcher(scrapeCfg.Timeout, scrapeCfg.MaxRetries, scrapeCfg.RetryBackoff, scrapeCfg.UserAgent)
		fmt.Printf("Phase 1: Downloading specs for %d states (concurrency=%d, timeout=%s, retries=%d)...\n", len(scrapeCfg.States), scrapeCfg.Concurrency, scrapeCfg.Timeout, scrapeCfg.MaxRetries)

		var wg sync.WaitGroup
		semaphore := make(chan struct{}, scrapeCfg.Concurrency)
		var mu sync.Mutex
		validStates := []string{}

		for _, state := range scrapeCfg.States {
			state := state
			wg.Add(1)
			go func() {
				defer wg.Done()
				semaphore <- struct{}{}
				defer func() { <-semaphore }()

				baseURL := fmt.Sprintf("https://api-%s.metrc.com/Documentation", state)
				fmt.Printf("Checking %s...\n", state)

				htmlURL := baseURL + "/PrintableList"
				if err := fetchHTMLSpec(state, htmlURL, htmlDir, fetcher); err != nil {
					fmt.Printf("  Error fetching HTML docs for %s: %v\n", state, err)
					return
				}

				mu.Lock()
				validStates = append(validStates, state)
				mu.Unlock()
				fmt.Printf("  Found docs for %s\n", state)

				postmanURL := baseURL + "/postman-json"
				if err := fetchPostmanSpec(state, postmanURL, postmanDir, fetcher); err != nil {
					fmt.Printf("  Error fetching Postman spec for %s: %v\n", state, err)
				}
			}()
		}
		wg.Wait()

		sort.Strings(validStates)
		fmt.Printf("Downloaded docs from %d states\n", len(validStates))
		if len(validStates) == 0 {
			return fmt.Errorf("no valid state documentation found from live scrape; check network access, states, or scraper settings")
		}
	} else {
		fmt.Println("Phase 1: Skipped downloading specs (using local files)")
	}

	fmt.Println("\nPhase 2: Parsing Postman specs for endpoints...")
	allOps := make(map[string]ParsedOperation)
	stateEndpoints := make(map[string]map[string]bool)
	endpointStates := make(map[string][]string)

	files, err := os.ReadDir(postmanDir)
	if err != nil {
		return fmt.Errorf("failed to read postman dir: %w", err)
	}
	if len(files) == 0 {
		return fmt.Errorf("no postman files found in %s", postmanDir)
	}

	for _, f := range files {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".json") {
			continue
		}

		state := strings.TrimSuffix(f.Name(), ".json")
		stateEndpoints[state] = make(map[string]bool)

		path := filepath.Join(postmanDir, f.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading %s: %v\n", f.Name(), err)
			continue
		}

		ops, err := ParsePostmanCollection(strings.NewReader(string(data)), cfg.Scraper)
		if err != nil {
			fmt.Printf("Error parsing %s: %v\n", f.Name(), err)
			continue
		}

		for _, op := range ops {
			key := normalizeKey(op.Method, op.URL)
			stateEndpoints[state][key] = true
			endpointStates[key] = append(endpointStates[key], state)

			if existing, exists := allOps[key]; exists {
				if op.Body != "" && existing.Body == "" {
					existing.Body = op.Body
					allOps[key] = existing
				}
			} else {
				allOps[key] = op
			}
		}

		fmt.Printf("  Parsed %s: %d endpoints\n", f.Name(), len(ops))
	}
	fmt.Printf("Total unique endpoints from Postman: %d\n", len(allOps))

	fmt.Println("\nPhase 3: Enriching with HTML documentation...")
	endpointsWithDocs := make(map[string]bool)

	htmlFiles, _ := os.ReadDir(htmlDir)
	for _, f := range htmlFiles {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".html") {
			continue
		}

		path := filepath.Join(htmlDir, f.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		htmlOps, err := ParseDocumentation(strings.NewReader(string(data)), cfg.Scraper)
		if err != nil {
			continue
		}

		enriched := 0
		for _, htmlOp := range htmlOps {
			key := normalizeKey(htmlOp.Method, htmlOp.URL)
			if existing, exists := allOps[key]; exists {
				endpointsWithDocs[key] = true
				if existing.Description == "" {
					existing.Description = htmlOp.Description
				}
				if len(existing.Permissions) == 0 {
					existing.Permissions = htmlOp.Permissions
				}
				if existing.Response == "" {
					existing.Response = htmlOp.Response
				}
				if existing.Body == "" && htmlOp.Body != "" {
					existing.Body = htmlOp.Body
				}

				if len(htmlOp.Params) > 0 {
					htmlParams := make(map[string]Param)
					for _, p := range htmlOp.Params {
						htmlParams[strings.ToLower(p.Name)] = p
					}

					for i, p := range existing.Params {
						if hp, ok := htmlParams[strings.ToLower(p.Name)]; ok {
							existing.Params[i].Description = hp.Description
							existing.Params[i].Required = hp.Required
							delete(htmlParams, strings.ToLower(p.Name))
						}
					}

					for i, p := range existing.Params {
						if strings.EqualFold(p.Name, "licenseNumber") {
							existing.Params[i].Required = true
						}
					}
				}
				allOps[key] = existing
				enriched++
			}
		}
		if enriched > 0 {
			fmt.Printf("  Enriched %d endpoints from %s\n", enriched, f.Name())
		}
	}

	var finalOps []ParsedOperation
	for _, op := range allOps {
		finalOps = append(finalOps, op)
	}

	fmt.Printf("\nPhase 4: Generating %d Bruno files...\n", len(finalOps))
	GenerateBrunoFiles(finalOps, brunoDir, cfg.PermissionOverrides)
	reportDir := "../docs/reports"
	writeCombinedReport(allOps, stateEndpoints, endpointStates, endpointsWithDocs, reportDir)

	return nil
}

func resolveScrapeOptions(opts Options, cfg app.ScraperConfig) scrapeOptions {
	states := normalizeStates(parseStates(opts.StatesCSV))
	if len(states) == 0 {
		states = normalizeStates(cfg.States)
	}
	if len(states) == 0 {
		states = append([]string{}, DefaultMetrcStates...)
	}

	timeoutSeconds := opts.TimeoutSeconds
	if timeoutSeconds <= 0 {
		timeoutSeconds = cfg.RequestTimeoutSeconds
	}
	if timeoutSeconds <= 0 {
		timeoutSeconds = 30
	}

	maxRetries := opts.MaxRetries
	if maxRetries < 0 {
		if cfg.MaxRetries > 0 {
			maxRetries = cfg.MaxRetries
		} else {
			maxRetries = 2
		}
	}

	retryBackoffMs := opts.RetryBackoffMs
	if retryBackoffMs <= 0 {
		retryBackoffMs = cfg.RetryBackoffMs
	}
	if retryBackoffMs <= 0 {
		retryBackoffMs = 500
	}

	concurrency := opts.Concurrency
	if concurrency <= 0 {
		concurrency = cfg.Concurrency
	}
	if concurrency <= 0 {
		concurrency = 5
	}

	userAgent := strings.TrimSpace(opts.UserAgent)
	if userAgent == "" {
		userAgent = strings.TrimSpace(cfg.UserAgent)
	}

	return scrapeOptions{
		States:       states,
		Timeout:      time.Duration(timeoutSeconds) * time.Second,
		MaxRetries:   maxRetries,
		RetryBackoff: time.Duration(retryBackoffMs) * time.Millisecond,
		Concurrency:  concurrency,
		UserAgent:    userAgent,
	}
}

func parseStates(statesCSV string) []string {
	if strings.TrimSpace(statesCSV) == "" {
		return nil
	}
	parts := strings.Split(statesCSV, ",")
	states := make([]string, 0, len(parts))
	for _, p := range parts {
		state := strings.TrimSpace(strings.ToLower(p))
		if state == "" {
			continue
		}
		states = append(states, state)
	}
	return states
}

func normalizeStates(input []string) []string {
	seen := make(map[string]struct{}, len(input))
	states := make([]string, 0, len(input))
	for _, state := range input {
		s := strings.TrimSpace(strings.ToLower(state))
		if s == "" {
			continue
		}
		if _, exists := seen[s]; exists {
			continue
		}
		seen[s] = struct{}{}
		states = append(states, s)
	}
	return states
}

func normalizeKey(method, url string) string {
	method = strings.ToUpper(strings.TrimSpace(method))
	url = "/" + strings.Trim(strings.TrimSpace(url), "/")
	return method + " " + url
}

var DefaultMetrcStates = []string{
	"ak", "al", "ar", "az", "ca", "co", "ct", "dc", "de", "fl",
	"ga", "gu", "hi", "ia", "id", "il", "in", "ks", "ky", "la",
	"ma", "md", "me", "mi", "mn", "mo", "ms", "mt", "nc", "nd",
	"ne", "nh", "nj", "nm", "nv", "ny", "oh", "ok", "or", "pa",
	"pr", "ri", "sc", "sd", "tn", "tx", "ut", "va", "vi", "vt",
	"wa", "wi", "wv", "wy",
}
