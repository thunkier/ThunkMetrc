package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"os"
	"encoding/json"

	"github.com/thunkier/thunkmetrc/tools/pkg/bruno"
)

func ExecuteRequest(cfg Config, req *bruno.Request, userApiKey string) string {
	logger := cfg.Logger
	if logger == nil {
		logger = defaultLogger{}
	}

	url := req.URL
	url = strings.ReplaceAll(url, "{{baseUrl}}", cfg.BaseURL)
	url = strings.ReplaceAll(url, ":licenseNumber", cfg.FacilityLicenseNumber)

	if !strings.Contains(url, "licenseNumber=") {
		if strings.Contains(url, "?") {
			url += "&licenseNumber=" + cfg.FacilityLicenseNumber
		} else {
			url += "?licenseNumber=" + cfg.FacilityLicenseNumber
		}
	}

	for _, blocked := range cfg.BlockedEndpoints {
		if strings.Contains(req.Name, blocked) {
			if !cfg.SuppressConsoleLog {
				logger.Log("[BLOCKED] %s is in blocked endpoints list", req.Name)
			}
			return ""
		}
	}

	if cfg.UsedRequests != nil {
		key := fmt.Sprintf("%s/%s", req.Group, req.Name)
		cfg.UsedRequests.Store(key, true)
	}

	// Basic Auth: VendorKey:UserKey
	vendorKey := cfg.APIKey
	userKey := cfg.UserKey

	// Allow override if provided (though mostly unused now)
	if userApiKey != "" {
		userKey = userApiKey
	}

	var bodyReader io.Reader
	if req.Body != "" {
		bodyReader = bytes.NewBufferString(req.Body)
	}

	httpReq, err := http.NewRequest(req.Method, url, bodyReader)
	if err != nil {
		if cfg.ErrorReporter != nil {
			cfg.ErrorReporter.RecordError(cfg.FacilityLicenseNumber, cfg.CurrentWorkflow, req.Method, req.URL, 0)
		}
		return ""
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth(vendorKey, userKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		if cfg.ErrorReporter != nil {
			cfg.ErrorReporter.RecordError(cfg.FacilityLicenseNumber, cfg.CurrentWorkflow, req.Method, req.URL, 0)
		}
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if cfg.ErrorReporter != nil {
			cfg.ErrorReporter.RecordError(cfg.FacilityLicenseNumber, cfg.CurrentWorkflow, req.Method, req.URL, resp.StatusCode)
		}
		return ""
	}

	if resp.StatusCode >= 400 {
		if cfg.ErrorReporter != nil {
			cfg.ErrorReporter.RecordError(cfg.FacilityLicenseNumber, cfg.CurrentWorkflow, req.Method, req.URL, resp.StatusCode)
		}
		// Log error for debugging
		if !cfg.SuppressConsoleLog {
			logger.Log("[ERROR] %s returned %d: %s", req.Name, resp.StatusCode, string(body))
		}
		return ""
	}

	if cfg.ResponsesDir != "" && req.Method == "GET" && len(body) > 0 {
		saveResponse(cfg.ResponsesDir, req.Group, req.Name, body)
	}

	return string(body)
}

func saveResponse(dir, group, name string, data []byte) {
	groupDir := filepath.Join(dir, group)
	if err := os.MkdirAll(groupDir, 0755); err != nil {
		return
	}

	var formatted bytes.Buffer
	if json.Indent(&formatted, data, "", "  ") == nil {
		data = formatted.Bytes()
	}

	filePath := filepath.Join(groupDir, name+".json")
	os.WriteFile(filePath, data, 0644)
}
