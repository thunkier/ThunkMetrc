package app

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

type ErrorEvent struct {
	Facility   string
	Workflow   string
	Method     string
	Endpoint   string // Normalized URL/Path
	StatusCode int
}

type ErrorAggregator struct {
	mu sync.Mutex

	facilityErrors       map[string][]ErrorEvent
	globalEndpointErrors map[string]map[string]int

	allFacilities map[string]bool
}

func NewErrorAggregator() *ErrorAggregator {
	return &ErrorAggregator{
		facilityErrors:       make(map[string][]ErrorEvent),
		globalEndpointErrors: make(map[string]map[string]int),
		allFacilities:        make(map[string]bool),
	}
}

func (ea *ErrorAggregator) RecordError(facility, workflow, method, endpoint string, statusCode int) {
	ea.mu.Lock()
	defer ea.mu.Unlock()

	if after, ok := strings.CutPrefix(endpoint, "{{baseUrl}}"); ok {
		endpoint = after
	}
	endpoint = strings.TrimPrefix(endpoint, "/")

	event := ErrorEvent{
		Facility:   facility,
		Workflow:   workflow,
		Method:     method,
		Endpoint:   endpoint,
		StatusCode: statusCode,
	}

	ea.facilityErrors[facility] = append(ea.facilityErrors[facility], event)
	ea.allFacilities[facility] = true

	key := fmt.Sprintf("%s %s", method, endpoint)
	if ea.globalEndpointErrors[key] == nil {
		ea.globalEndpointErrors[key] = make(map[string]int)
	}
	ea.globalEndpointErrors[key][facility] = statusCode
}

func (ea *ErrorAggregator) PrintReport(logger Logger) {
	ea.mu.Lock()
	defer ea.mu.Unlock()

	logger.Log("\n===== AUTOMATION ERROR REPORT =====")

	totalFacilities := len(ea.allFacilities)
	if totalFacilities == 0 {
		logger.Log("No facilities processed or no errors recorded.")
		return
	}

	logger.Log("\n[ GLOBAL FAILURES ] (Endpoints failing on ALL %d facilities)", totalFacilities)
	globalFailuresFound := false

	var keys []string
	for k := range ea.globalEndpointErrors {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		facMap := ea.globalEndpointErrors[key]
		if len(facMap) == totalFacilities {
			firstCode := 0
			allSame := true
			for _, code := range facMap {
				if firstCode == 0 {
					firstCode = code
				} else if code != firstCode {
					allSame = false
					break
				}
			}

			if allSame {
				logger.Log("  - [%d] %s", firstCode, key)
				globalFailuresFound = true
			}
		}
	}
	if !globalFailuresFound {
		logger.Log("  (None found)")
	}

	logger.Log("\n[ FACILITY FAILURES ] (Unique errors per facility)")

	var facilities []string
	for f := range ea.allFacilities {
		facilities = append(facilities, f)
	}
	sort.Strings(facilities)

	for _, fac := range facilities {
		events := ea.facilityErrors[fac]
		if len(events) == 0 {
			continue
		}

		logger.Log("  Facility: %s", fac)
		uniqueErrors := make(map[string]int)
		for _, e := range events {
			k := fmt.Sprintf("[%d] %s %s (Workflow: %s)", e.StatusCode, e.Method, e.Endpoint, e.Workflow)
			uniqueErrors[k]++
		}

		var errKeys []string
		for k := range uniqueErrors {
			errKeys = append(errKeys, k)
		}
		sort.Strings(errKeys)

		for _, k := range errKeys {
			count := uniqueErrors[k]
			if count > 1 {
				logger.Log("    - %s (x%d)", k, count)
			} else {
				logger.Log("    - %s", k)
			}
		}
	}
	logger.Log("\n===================================")
}
