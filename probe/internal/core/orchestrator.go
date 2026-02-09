package core

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	mockserver "github.com/thunkier/metrc-mock-server/pkg/server"
	"github.com/thunkier/thunkmetrc/probe/internal/app"
	"github.com/thunkier/thunkmetrc/probe/internal/metrc/client"
	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
)

func Orchestrate() {
	dryRunPtr := flag.Bool("dry-run", true, "Enable dry run mode (no real execution)")
	stepPtr := flag.Bool("step", false, "Enable manual stepping (process one facility at a time with user confirmation)")
	mockPermissionsPtr := flag.Bool("mock-permissions", false, "Load mock employee permissions from fixtures")
	flag.Parse()

	config := app.LoadConfig()
	config.DryRun = *dryRunPtr
	config.ManualStep = *stepPtr
	config.UseMockPermissions = *mockPermissionsPtr
	if !config.UseMockPermissions && os.Getenv("METRC_USE_MOCK_PERMISSIONS") == "true" {
		config.UseMockPermissions = true
	}

	if config.DryRun {
		config.BaseURL = "http://localhost:8080"
		fmt.Println("DryRun active (flag): Using Mock Server at http://localhost:8080")

		go func() {
			srv := mockserver.New(mockserver.Config{
				Port:          8080,
				Scenario:      "default",
				RequireAuth:   false,
				Stateful:      false,
				ResponsesPath: config.ResponsesDir,
			})
			if err := srv.ListenAndServe(); err != nil {
				fmt.Printf("Mock server error: %v\n", err)
			}
		}()
		// Allow server to bind
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("Starting Automation using Specs in: %s (DryRun: %v, MockPermissions: %v)\n", config.SpecsDir, config.DryRun, config.UseMockPermissions)
	if !config.DryRun {
		fmt.Printf("Saving responses to: %s\n", config.ResponsesDir)
	}

	logFile := "automation_run.log"
	if config.DryRun {
		logFile = "automation_dryrun.log"
	}

	logManager, err := app.NewGlobalLogManager(logFile)
	if err != nil {
		fmt.Printf("Failed to create log file: %v\n", err)
		return
	}
	defer logManager.Close()
	fmt.Printf("Logging detailed output to '%s'...\n", logFile)

	config.Logger = app.NewFacilityLogger(logManager, "Global")

	fetcher := client.NewFetcher(config.ToClientConfig())
	facilities, err := fetcher.GetFacilities()
	if err != nil {
		fmt.Printf("Error fetching facilities: %v\n", err)
	} else if len(facilities) == 0 && !config.DryRun {
		fmt.Printf("Error: no facilities found\n")
	}

	dashboard := app.NewDashboard()
	config.Dashboard = dashboard
	for _, f := range facilities {
		dashboard.Register(f.License.Number, f.Name)
	}

	go dashboard.Run()
	defer dashboard.Stop()

	var wg sync.WaitGroup

	var statsMutex sync.Mutex
	globalSkipped := make(map[string]map[string]int)
	globalEmpty := make(map[string]int)
	globalExecuted := make(map[string]int)
	totalFacilitiesRun := 0

	for _, fac := range facilities {
		wg.Add(1)
		go func(fac models.Facility) {
			defer wg.Done()

			facConfig := config
			facConfig.Facility = &fac
			facConfig.FacilityLicenseNumber = fac.License.Number
			facConfig.Logger = app.NewFacilityLogger(logManager, fmt.Sprintf("%s (%s)", fac.Name, fac.License.Number))
			facConfig.Cleanup = &app.CleanupManager{}
			facConfig.DependencyCache = &sync.Map{}

			if config.UseMockPermissions {
				perms, err := loadMockEmployeePermissions(fac.License.Number)
				if err != nil {
					facConfig.Logger.Log("[WARN] Failed to load mock permissions for %s: %v", fac.License.Number, err)
				} else {
					facConfig.EmployeePermissions = perms
					facConfig.Logger.Log("[INFO] Loaded %d mock permissions", len(perms))
				}
			}

			_, _, result := ExecuteDataWorkflows(facConfig)
			facConfig.Logger.Flush()

			statsMutex.Lock()
			totalFacilitiesRun++

			// Executed
			for ep, count := range result.DataGenerated {
				if count > 0 {
					globalExecuted[ep]++
				}
			}

			// Empty
			for _, ep := range result.EmptyEndpoints {
				globalEmpty[ep]++
			}

			// Skipped
			for ep, reason := range result.SkippedEndpoints {
				if globalSkipped[ep] == nil {
					globalSkipped[ep] = make(map[string]int)
				}
				globalSkipped[ep][reason]++
			}
			statsMutex.Unlock()

		}(fac)
	}

	wg.Wait()

	if config.ErrorAggregator != nil {
		config.ErrorAggregator.PrintReport(config.Logger)
		config.Logger.Flush()
	}

	fmt.Println("\n=== Global Endpoint Statistics ===")
	fmt.Printf("Total Facilities Processed: %d\n", totalFacilitiesRun)

	allEndpoints := make(map[string]bool)
	for ep := range globalExecuted {
		allEndpoints[ep] = true
	}
	for ep := range globalEmpty {
		allEndpoints[ep] = true
	}
	for ep := range globalSkipped {
		allEndpoints[ep] = true
	}

	var neverExecuted []string
	var partiallyExecuted []string
	var alwaysEmpty []string

	for ep := range allEndpoints {
		execCount := globalExecuted[ep]
		emptyCount := globalEmpty[ep]

		runCount := execCount + emptyCount

		if runCount == 0 {
			neverExecuted = append(neverExecuted, ep)
		} else if runCount < totalFacilitiesRun {
			partiallyExecuted = append(partiallyExecuted, ep)
		}

		if execCount == 0 && emptyCount > 0 {
			alwaysEmpty = append(alwaysEmpty, ep)
		}
	}

	sort.Strings(neverExecuted)
	if len(neverExecuted) > 0 {
		fmt.Printf("\n[NEVER RUN] The following endpoints were NOT executed on ANY facility:\n")
		count := 0
		for _, ep := range neverExecuted {
			if strings.HasSuffix(strings.ToLower(ep), "byid") || strings.HasSuffix(strings.ToLower(ep), "id") {
				continue
			}

			reasons := globalSkipped[ep]
			var dominantReason string
			var maxCount int
			for r, c := range reasons {
				if c > maxCount {
					maxCount = c
					dominantReason = r
				}
			}
			fmt.Printf(" - %s (Most common reason: %s)\n", ep, dominantReason)
			count++
		}
		if count == 0 {
			fmt.Println(" (None unrelated to IDs)")
		}
	}

	sort.Strings(alwaysEmpty)
	if len(alwaysEmpty) > 0 {
		fmt.Printf("\n[ALWAYS EMPTY] The following endpoints executed but ALWAYS returned 0 items:\n")
		count := 0
		for _, ep := range alwaysEmpty {
			var sibling string
			if strings.HasPrefix(ep, "GetActive") {
				sibling = strings.Replace(ep, "GetActive", "GetInactive", 1)
			} else if strings.HasPrefix(ep, "GetInactive") {
				sibling = strings.Replace(ep, "GetInactive", "GetActive", 1)
			}

			if sibling != "" {
				if globalExecuted[sibling] > 0 {
					continue
				}
			}

			fmt.Printf(" - %s\n", ep)
			count++
		}
		if count == 0 {
			fmt.Println(" (None)")
		}
	}
}

func loadMockEmployeePermissions(licenseNumber string) ([]string, error) {
	parts := strings.Split(licenseNumber, "-")
	prefix := licenseNumber
	if len(parts) >= 3 {
		prefix = strings.Join(parts[:3], "-")
	}

	path := filepath.Join("..", "mockserver", "pkg", "server", "handlers", "fixtures", "permissions", prefix+".json")

	data, err := os.ReadFile(path)
	if err != nil {
		path = filepath.Join("mockserver", "pkg", "server", "handlers", "fixtures", "permissions", prefix+".json")
		data, err = os.ReadFile(path)
		if err != nil {
			return nil, err
		}
	}

	var perms []string
	if err := json.Unmarshal(data, &perms); err != nil {
		return nil, fmt.Errorf("failed to parse permissions json: %w", err)
	}

	return perms, nil
}
