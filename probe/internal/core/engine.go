package core

import (
	"fmt"

	"github.com/thunkier/thunkmetrc/probe/internal/app"
	"github.com/thunkier/thunkmetrc/probe/internal/scenarios"

	"github.com/thunkier/thunkmetrc/probe/internal/metrc/client"
)

func ExecuteDataWorkflows(cfg app.Config) (string, []string, *app.WorkflowResult) {
	f := client.NewFetcher(cfg.ToClientConfig())
	var passedWorkflows []string
	var scenarioName string

	combinedResult := &app.WorkflowResult{
		CreatedIDs:       []int{},
		DataGenerated:    make(map[string]int),
		SkippedEndpoints: make(map[string]string),
		EmptyEndpoints:   []string{},
		Outputs:          make(map[string]any),
	}

	var matchedScenario *scenarios.Scenario
	for _, s := range scenarios.ScenarioRegistry {
		if cfg.Facility != nil && s.Matches(*cfg.Facility) {
			matchedScenario = &s
			break
		}
	}

	if matchedScenario == nil {
		cfg.Logger.Log("  [WARN] No matching scenario found for facility.")
		if cfg.Dashboard != nil {
			cfg.Dashboard.Update(cfg.FacilityLicenseNumber, "No Scenario Found", "Error")
		}
		return "None", passedWorkflows, combinedResult
	}

	scenarioName = matchedScenario.Name
	cfg.Logger.Log("  [SCENARIO] Matches '%s': %s", scenarioName, matchedScenario.Description)
	if cfg.Dashboard != nil {
		cfg.Dashboard.UpdateScenario(cfg.FacilityLicenseNumber, scenarioName)
		cfg.Dashboard.Update(cfg.FacilityLicenseNumber, "Scenario Matched", "Running")
	}

	for _, w := range matchedScenario.Steps {
		if !w.ShouldRun(cfg) {
			continue
		}

		cfg.CurrentWorkflow = w.Name()
		cfg.Logger.Log("  [WORKFLOW] Starting '%s'...", w.Name())

		if cfg.Dashboard != nil {
			cfg.Dashboard.Update(cfg.FacilityLicenseNumber, fmt.Sprintf("Running %s", w.Name()), "Running")
		}

		result, err := w.Execute(cfg, f)
		if result == nil {
			result = &app.WorkflowResult{}
		}
		if result.DataGenerated == nil {
			result.DataGenerated = make(map[string]int)
		}
		if err != nil {
			cfg.Logger.Log("  [WORKFLOW-ERR] '%s' failed: %v", w.Name(), err)
			continue
		}

		passedWorkflows = append(passedWorkflows, w.Name())

		for k, v := range result.DataGenerated {
			combinedResult.DataGenerated[k] += v
		}

		if result.SkippedEndpoints != nil {
			for k, v := range result.SkippedEndpoints {
				combinedResult.SkippedEndpoints[k] = v
			}
		}

		if len(result.EmptyEndpoints) > 0 {
			combinedResult.EmptyEndpoints = append(combinedResult.EmptyEndpoints, result.EmptyEndpoints...)
		}

		if result.Outputs != nil {
			for k, v := range result.Outputs {
				combinedResult.Outputs[k] = v
			}
		}

		cfg.Logger.Log("  [WORKFLOW-OK] '%s' created %d entities.", w.Name(), len(result.CreatedIDs))

		if len(result.EmptyEndpoints) > 0 {
			cfg.Logger.Log("    [EMPTY] %d endpoints returned no data:", len(result.EmptyEndpoints))
			for _, ep := range result.EmptyEndpoints {
				cfg.Logger.Log("      - %s", ep)
			}
		}

		if len(result.SkippedEndpoints) > 0 {
			cfg.Logger.Log("    [SKIPPED] %d endpoints were not run:", len(result.SkippedEndpoints))
			// Group by reason for cleaner output
			skippedByReason := make(map[string][]string)
			for ep, reason := range result.SkippedEndpoints {
				skippedByReason[reason] = append(skippedByReason[reason], ep)
			}
			for reason, eps := range skippedByReason {
				cfg.Logger.Log("      Reason: %s", reason)
				for _, ep := range eps {
					cfg.Logger.Log("        - %s", ep)
				}
			}
		}

		for entity, count := range result.DataGenerated {
			cfg.Logger.Log("    - %s: %d", entity, count)
		}

		if len(result.Outputs) > 0 {
			for k, v := range result.Outputs {
				cfg.DependencyCache.Store(k, v)
				cfg.Logger.Log("    [OUTPUT] %s = %v", k, v)
			}
		}
	}

	if cfg.Dashboard != nil {
		cfg.Dashboard.Update(cfg.FacilityLicenseNumber, "Workflows Complete", "Done")
	}

	cfg.Logger.Log("  [CLEANUP] Executing %d deferred tasks...", len(cfg.Cleanup.Tasks))

	for i := len(cfg.Cleanup.Tasks) - 1; i >= 0; i-- {
		cfg.Cleanup.Tasks[i]()
	}

	cfg.Cleanup.Tasks = []func(){}

	return scenarioName, passedWorkflows, combinedResult
}
