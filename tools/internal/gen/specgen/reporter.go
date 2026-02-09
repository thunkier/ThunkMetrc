package specgen

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type CoverageReport struct {
	TotalEndpoints       int                 `json:"totalEndpoints"`
	TotalStates          int                 `json:"totalStates"`
	EndpointsWithoutDocs []string            `json:"endpointsWithoutDocs"`
	MissingByState       map[string][]string `json:"missingByState"`
	EndpointAvailability map[string][]string `json:"endpointAvailability"`
	EndpointLog          []EndpointLogEntry  `json:"endpointLog"`
}

type EndpointLogEntry struct {
	Group  string `json:"group"`
	Method string `json:"method"`
	URL    string `json:"url"`
	Name   string `json:"name"`
	IsV1   bool   `json:"isV1"`
}

func writeCombinedReport(
	allOps map[string]ParsedOperation,
	stateEndpoints map[string]map[string]bool,
	endpointStates map[string][]string,
	endpointsWithDocs map[string]bool,
	outDir string,
) {
	if err := os.MkdirAll(outDir, 0755); err != nil {
		fmt.Printf("Warning: could not create report directory: %v\n", err)
		return
	}

	report := CoverageReport{
		TotalEndpoints:       len(allOps),
		TotalStates:          len(stateEndpoints),
		EndpointsWithoutDocs: []string{},
		MissingByState:       make(map[string][]string),
		EndpointAvailability: make(map[string][]string),
		EndpointLog:          []EndpointLogEntry{},
	}

	var allEndpointKeys []string
	for key := range allOps {
		allEndpointKeys = append(allEndpointKeys, key)
	}

	// Populate EndpointLog and EndpointsWithoutDocs
	for _, op := range allOps {
		isV1 := isV1Endpoint(op.URL)
		report.EndpointLog = append(report.EndpointLog, EndpointLogEntry{
			Group:  op.Group,
			Method: op.Method,
			URL:    op.URL,
			Name:   op.Name,
			IsV1:   isV1,
		})

		key := op.Method + " " + op.URL
		if !endpointsWithDocs[key] {
			report.EndpointsWithoutDocs = append(report.EndpointsWithoutDocs, key)
		}
	}

	// Find missing endpoints per state
	for state, endpoints := range stateEndpoints {
		var missing []string
		for _, key := range allEndpointKeys {
			if !endpoints[key] {
				missing = append(missing, key)
			}
		}
		if len(missing) > 0 {
			report.MissingByState[state] = missing
		}
	}

	// Store endpoint availability (which states have each endpoint)
	report.EndpointAvailability = endpointStates

	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		fmt.Printf("Warning: could not marshal combined report: %v\n", err)
		return
	}

	reportPath := filepath.Join(outDir, "combined_report.json")
	if err := os.WriteFile(reportPath, data, 0644); err != nil {
		fmt.Printf("Warning: could not write combined report: %v\n", err)
		return
	}

	fmt.Printf("Wrote combined report to %s\n", reportPath)
	fmt.Printf("  - %d total endpoints\n", report.TotalEndpoints)
	fmt.Printf("  - %d endpoints without HTML docs\n", len(report.EndpointsWithoutDocs))
	fmt.Printf("  - %d states with missing endpoints\n", len(report.MissingByState))
}
