package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// Run executes the scraping process for Metrc documentation.
func Run() error {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 5)
	var mu sync.Mutex
	validStates := []string{}

	outDir := "../specs/metrc-html"
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	brunoDir := "../specs/metrc-bruno"
	os.RemoveAll(brunoDir)
	if err := os.MkdirAll(brunoDir, 0755); err != nil {
		return err
	}

	createBrunoConfig(brunoDir)

	// Load Config
	config := loadConfig("config.json")

	for _, state := range MetrcStates {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()

			time.Sleep(500 * time.Millisecond)

			url := fmt.Sprintf("https://api-%s.metrc.com/Documentation/PrintableList", s)
			fmt.Printf("Checking %s...\n", s)

			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", s, err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == 200 {
				fmt.Printf("Found docs for %s\n", s)
				mu.Lock()
				validStates = append(validStates, s)
				mu.Unlock()

				content, _ := io.ReadAll(resp.Body)
				os.WriteFile(filepath.Join(outDir, s+".html"), content, 0644)

				if s == "or" {
					fmt.Printf("Generating Specs from %s...\n", s)
					ops, err := ParseDocumentation(strings.NewReader(string(content)), config.Scraper)
					if err == nil {
						GenerateBrunoFiles(ops, brunoDir, config.PermissionOverrides)
					}
				}
			}
		}(state)
	}

	wg.Wait()
	fmt.Printf("Found documentation for %d states: %v\n", len(validStates), validStates)

	return nil
}

func createBrunoConfig(dir string) {
	config := `{
  "version": "1",
  "name": "Metrc API",
  "type": "collection",
  "ignore": [".git"]
}`
	os.WriteFile(filepath.Join(dir, "bruno.json"), []byte(config), 0644)
}

// GenerateBrunoFiles creates .bru files from parsed operations.
func GenerateBrunoFiles(ops []ParsedOperation, outDir string, permOverrides map[string]string) {
	for _, op := range ops {
		groupDir := filepath.Join(outDir, cleanFileName(op.Group))
		os.MkdirAll(groupDir, 0755)

		fileName := fmt.Sprintf("%s.bru", cleanFileName(op.Name))
		path := filepath.Join(groupDir, fileName)

		f, err := os.Create(path)
		if err != nil {
			continue
		}

		writeMeta(f, op)
		writeRequest(f, op)
		writeParams(f, op)
		writeAuth(f)
		writeBody(f, op)
		writeDocs(f, op, permOverrides)

		f.Close()
	}
}

type fullConfig struct {
	PermissionOverrides map[string]string `json:"PermissionOverrides"`
	Scraper             ScraperConfig     `json:"Scraper"`
}

func loadConfig(configFile string) fullConfig {
	var cfg fullConfig
	paths := []string{configFile, "../" + configFile, "../../" + configFile}
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			data, err := os.ReadFile(p)
			if err == nil {
				_ = json.Unmarshal(data, &cfg)
				fmt.Printf("Loaded configuration from %s\n", p)
				return cfg
			}
		}
	}
	return cfg
}

func writeMeta(w io.Writer, op ParsedOperation) {
	fmt.Fprintf(w, "meta {\n  name: %s\n  type: http\n  seq: 1\n}\n\n", op.Name)
}

func writeRequest(w io.Writer, op ParsedOperation) {
	fmt.Fprintf(w, "%s {\n  url: {{baseUrl}}%s\n  body: json\n  auth: basic\n}\n\n", strings.ToLower(op.Method), op.URL)
}

func writeParams(w io.Writer, op ParsedOperation) {
	if len(op.Params) > 0 {
		fmt.Fprintln(w, "params:query {")
		for _, p := range op.Params {
			fmt.Fprintf(w, "  %s: \n", p.Name)
		}
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "")
	}
}

func writeAuth(w io.Writer) {
	fmt.Fprintln(w, "auth:basic {\n  username: {{username}}\n  password: {{password}}\n}")
	fmt.Fprintln(w, "")
}

func writeBody(w io.Writer, op ParsedOperation) {
	if op.Body != "" {
		fmt.Fprintln(w, "body:json {")
		fmt.Fprintln(w, op.Body)
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "")
	}
}

func writeDocs(w io.Writer, op ParsedOperation, overrides map[string]string) {
	fmt.Fprintln(w, "docs {")
	fmt.Fprintln(w, "  "+op.Description)
	if len(op.Permissions) > 0 {
		fmt.Fprintln(w, "\n  Permissions Required:")
		for _, p := range op.Permissions {
			permName := p
			if val, ok := overrides[p]; ok {
				permName = val
			}
			fmt.Fprintf(w, "  - %s\n", permName)
		}
	}
	if op.Response != "" {
		fmt.Fprintln(w, "\n  Example Response:")
		fmt.Fprintln(w, "  "+strings.ReplaceAll(op.Response, "\n", "\n  "))
	}
	fmt.Fprintln(w, "}")
}

func cleanFileName(s string) string {
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, ":", "")
	return strings.TrimSpace(s)
}
