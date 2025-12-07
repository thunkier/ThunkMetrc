package runner

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/thunkmetrc/automation/internal/runner/common"
	"github.com/thunkmetrc/automation/pkg/bruno"
	"github.com/thunkmetrc/automation/pkg/metrc/client"
)

func Orchestrate() {
	dryRunPtr := flag.Bool("dry-run", true, "Enable dry run mode (no real execution)")
	stepPtr := flag.Bool("step", false, "Enable manual stepping (process one facility at a time with user confirmation)")
	flag.Parse()

	config := common.LoadConfig()
	config.DryRun = *dryRunPtr
	config.ManualStep = *stepPtr

	fmt.Printf("Starting Automation using Specs in: %s (DryRun: %v)\n", config.SpecsDir, config.DryRun)
	fmt.Println("DEBUG: Running Refactored Orchestrator")
	if !config.DryRun {
		fmt.Printf("Saving responses to: %s\n", config.ResponsesDir)
	}

	logFile := "automation_run.log"
	if config.DryRun {
		logFile = "automation_dryrun.log"
	}

	logManager, err := common.NewGlobalLogManager(logFile)
	if err != nil {
		fmt.Printf("Failed to create log file: %v\n", err)
		return
	}
	defer logManager.Close()
	fmt.Printf("Logging detailed output to '%s'...\n", logFile)

	config.Logger = common.NewFacilityLogger(logManager, "Global")

	if true {
		facilities, err := FetchFacilities(config)
		if err != nil {
			fmt.Printf("Error fetching facilities: %v\n", err)
		}

		if len(facilities) == 0 && config.DryRun {
		}

		fmt.Printf("Found %d facilities:\n", len(facilities))
		for _, f := range facilities {
			fmt.Printf(" - %s (%s)\n", f.Name, f.License.Number)
		}

		collection, err := bruno.ParseCollection(config.SpecsDir)
		if err != nil {
			fmt.Printf("Failed to parse collection: %v\n", err)
			return
		}
		fmt.Printf("Found %d requests.\n", len(collection.Requests))

		groups := make(map[string][]bruno.Request)
		for _, req := range collection.Requests {
			groups[req.Group] = append(groups[req.Group], req)
		}

		fmt.Println("--- Starting Brute Force Execution ---")

		var completedGroups sync.Map
		var wg sync.WaitGroup

		for _, f := range facilities {
			fmt.Printf("Queuing Facility: %s (%s)\n", f.Name, f.License.Number)

			facilityConfig := config
			fCopy := f
			facilityConfig.Facility = &fCopy
			facilityConfig.FacilityLicenseNumber = f.License.Number

			logger := common.NewFacilityLogger(logManager, fmt.Sprintf("%s (%s)", f.Name, f.License.Number))
			facilityConfig.Logger = logger

			processFacility := func(cfg common.Config, facName string, grps map[string][]bruno.Request, l *common.FacilityLogger) {
				defer l.Flush()

				for groupName, requests := range grps {
					if _, done := completedGroups.Load(groupName); done {
						continue
					}

					shouldRun := ShouldRunGroup(groupName, cfg.Facility.FacilityType)
					fmt.Printf("DEBUG: Group '%s' ShouldRun: %v\n", groupName, shouldRun)
					if !shouldRun {
						l.Log("  [SKIP] Group '%s' not supported by facility permissions.", groupName)
						continue
					}

					localRequests := make([]bruno.Request, len(requests))
					for i, r := range requests {
						localRequests[i] = r.Clone()
					}

					dataFound := ExecuteGroupWorkflow(cfg, groupName, localRequests)

					if dataFound {
						fmt.Printf("DEBUG: Group '%s' completed with data. Skipping for future facilities.\n", groupName)
						completedGroups.Store(groupName, true)
					}
				}
				fmt.Printf("Finished Facility: %s\n", facName)
			}

			if config.ManualStep {
				fmt.Printf("\n[STEP] About to process facility: %s (%s)\nPress ENTER to continue...", f.Name, f.License.Number)
				bufio.NewReader(os.Stdin).ReadString('\n')

				if facilityConfig.ExternalConfig.SpecialEmployeeUserKey != "" {
					var permReq *bruno.Request
					if empGroup, ok := groups["Employees"]; ok {
						for _, r := range empGroup {
							if strings.Contains(r.Name, "GetPermissions V2") {
								permReq = &r
								break
							}
						}
					}
					if permReq != nil {
						clone := permReq.Clone()
						sep := "?"
						if strings.Contains(clone.URL, "?") {
							sep = "&"
						}
						clone.URL = clone.URL + sep + "employeeLicenseNumber=" + facilityConfig.ExternalConfig.SpecialEmployeeUserKey

						logger.Log("  [PRE-FLIGHT] Fetching Employee Permissions...")
						body := client.ExecuteRequest(facilityConfig, &clone, "")

						var perms []string
						if err := json.Unmarshal([]byte(body), &perms); err == nil {
							facilityConfig.EmployeePermissions = perms
							logger.Log("  [PRE-FLIGHT] User has %d permissions.", len(perms))
						} else {
							logger.Log("  [WARN] Failed to parse permissions: %v. Proceeding with empty permissions (Strict mode will likely fail requests).", err)
						}
					}
				}

				processFacility(facilityConfig, f.Name, groups, logger)
			} else {
				wg.Add(1)
				go func(cfg common.Config, facName string, grps map[string][]bruno.Request, l *common.FacilityLogger) {
					defer wg.Done()

					if cfg.ExternalConfig.SpecialEmployeeUserKey != "" {
						var permReq *bruno.Request
						if empGroup, ok := groups["Employees"]; ok {
							for _, r := range empGroup {
								if strings.Contains(r.Name, "GetPermissions V2") {
									permReq = &r
									break
								}
							}
						}
						if permReq != nil {
							clone := permReq.Clone()
							sep := "?"
							if strings.Contains(clone.URL, "?") {
								sep = "&"
							}
							clone.URL = clone.URL + sep + "employeeLicenseNumber=" + cfg.ExternalConfig.SpecialEmployeeUserKey

							l.Log("  [PRE-FLIGHT] Fetching Employee Permissions...")
							body := client.ExecuteRequest(cfg, &clone, "")
							var perms []string
							if err := json.Unmarshal([]byte(body), &perms); err == nil {
								// Update local config copy
								cfg.EmployeePermissions = perms
								l.Log("  [PRE-FLIGHT] User has %d permissions.", len(perms))
							}
						}
					}

					processFacility(cfg, facName, grps, l)
				}(facilityConfig, f.Name, groups, logger)
			}
		}

		wg.Wait()
		fmt.Println("All chains completed. Check 'automation_run.log' for details (if not empty).")

		fmt.Println("\n--- Unused Endpoints Analysis ---")

		unusedByGroup := make(map[string][]*bruno.Request)
		totalUnused := 0
		for i := range collection.Requests {
			req := &collection.Requests[i]
			if _, used := config.UsedRequests.Load(req.Name); !used {
				unusedByGroup[req.Group] = append(unusedByGroup[req.Group], req)
				totalUnused++
			}
		}

		if totalUnused > 0 {
			fmt.Printf("Total Unused Endpoints: %d / %d\n", totalUnused, len(collection.Requests))
			for grp, reqs := range unusedByGroup {
				fmt.Printf("\nGroup: %s (%d unused)\n", grp, len(reqs))
				for _, r := range reqs {
					fmt.Printf("  - %s (%s)\n", r.Name, r.Method)
				}
			}
		} else {
			fmt.Println("Amazing! 100% of endpoints were covered.")
		}
	}
}
