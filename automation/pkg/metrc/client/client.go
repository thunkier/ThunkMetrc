package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/thunkmetrc/automation/internal/runner/common"
	"github.com/thunkmetrc/automation/pkg/bruno"
)

var rateLimiter = time.NewTicker(time.Second / 50)

func ExecuteRequest(cfg common.Config, req *bruno.Request, id string) string {
	lowerUrl := strings.ToLower(req.URL)
	blocked := false
	for _, blockedPath := range cfg.ExternalConfig.BlockedEndpoints {
		if strings.Contains(lowerUrl, strings.ToLower(blockedPath)) {
			blocked = true
			break
		}
	}
	if blocked {
		return ""
	}

	if !cfg.DryRun {
		<-rateLimiter.C
	}

	if cfg.UsedRequests != nil {
		cfg.UsedRequests.Store(req.Name, true)
	}

	if cfg.DryRun {
		cfg.Logger.Log("  [DRY] %s %s", req.Method, req.URL)
		return GetMockResponse(req)
	}

	rawUrl := PrepareRequestURL(cfg, req, id)

	client := &http.Client{}

	var bodyReader io.Reader
	if req.Body != "" && (req.Method == "POST" || req.Method == "PUT") {
		bodyReader = strings.NewReader(req.Body)
	}

	r, _ := http.NewRequest(req.Method, rawUrl, bodyReader)

	if bodyReader != nil {
		r.Header.Set("Content-Type", "application/json")
	}

	r.SetBasicAuth(cfg.APIKey, os.Getenv("METRC_USER_API_KEY"))

	resp, err := client.Do(r)
	if err != nil {
		cfg.Logger.Log("  [ERR] %s", err)
		return ""
	}
	defer resp.Body.Close()

	bodyBytes, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		cfg.Logger.Log("  [WARN] Failed to read response body: %v", readErr)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyStr := string(bodyBytes)
		if len(bodyBytes) == 0 {
			bodyStr = "(empty body)"
		}
		cfg.Logger.Log("  [%d] %s %s | Response: %s", resp.StatusCode, req.Method, rawUrl, bodyStr)
		fmt.Printf("  [ERR %d] %s %s | %s\n", resp.StatusCode, req.Method, rawUrl, bodyStr)
	} else {
		bodyStr := string(bodyBytes)
		flatBody := strings.ReplaceAll(strings.ReplaceAll(bodyStr, "\n", " "), "\r", "")

		if len(flatBody) < 200 {
			fmt.Printf("  [%d] %s %s | Body: %s\n", resp.StatusCode, req.Method, rawUrl, flatBody)
		} else {
			fmt.Printf("  [%d] %s %s | (Body omitted)\n", resp.StatusCode, req.Method, rawUrl)
		}
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 500 {
		SaveResponse(cfg, req, bodyBytes)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 600 {
		return string(bodyBytes)
	}

	return ""
}
