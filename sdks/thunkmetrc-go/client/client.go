package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ClientConfig holds configuration for the MetrcClient
type ClientConfig struct {
	HTTPClient *http.Client
	Logger     *slog.Logger
	UserAgent  string
}

// ClientOption is a function that configures a ClientConfig
type ClientOption func(*ClientConfig)

// WithHTTPClient sets a custom http.Client (e.g. for mTLS or custom transport)
func WithHTTPClient(c *http.Client) ClientOption {
	return func(cfg *ClientConfig) {
		cfg.HTTPClient = c
	}
}

// WithTimeout sets the timeout for the default http.Client.
// Ignored if WithHTTPClient is used.
func WithTimeout(d time.Duration) ClientOption {
	return func(cfg *ClientConfig) {
		if cfg.HTTPClient != nil {
			cfg.HTTPClient.Timeout = d
		}
	}
}

// WithLogger sets a structured logger (slog.Logger)
func WithLogger(l *slog.Logger) ClientOption {
	return func(cfg *ClientConfig) {
		cfg.Logger = l
	}
}

// WithUserAgent sets a custom User-Agent header
func WithUserAgent(ua string) ClientOption {
	return func(cfg *ClientConfig) {
		cfg.UserAgent = ua
	}
}

type MetrcClient struct {
	BaseURL   string
	VendorKey string
	UserKey   string
	Client    *http.Client
	Logger    *slog.Logger
	UserAgent string
}

// MetrcError represents a structured API error from Metrc.
type MetrcError struct {
	StatusCode       int
	Message          string   `json:"Message"`
	ValidationErrors []string `json:"ValidationErrors"`
}

func (e *MetrcError) Error() string {
	msg := fmt.Sprintf("API Error %d: %s", e.StatusCode, e.Message)
	if len(e.ValidationErrors) > 0 {
		msg += fmt.Sprintf(" (Errors: %v)", e.ValidationErrors)
	}
	return msg
}

// New creates a new MetrcClient with optional configuration.
// By default, it uses a 100s timeout and a default User-Agent.
func New(baseURL, vendorKey, userKey string, opts ...ClientOption) *MetrcClient {
	// Default Configuration
	cfg := &ClientConfig{
		HTTPClient: &http.Client{Timeout: 100 * time.Second},
		Logger:     nil, // Silent by default
		UserAgent:  "ThunkMetrc/0.1 Go",
	}

	// Apply Options
	for _, opt := range opts {
		opt(cfg)
	}

	return &MetrcClient{
		BaseURL:   strings.TrimRight(baseURL, "/"),
		VendorKey: vendorKey,
		UserKey:   userKey,
		Client:    cfg.HTTPClient,
		Logger:    cfg.Logger,
		UserAgent: cfg.UserAgent,
	}
}

// NewWithClient creates a new MetrcClient with a specific http.Client.
// Deprecated: Use New(..., WithHTTPClient(c)) instead.
func NewWithClient(baseURL, vendorKey, userKey string, httpClient *http.Client) *MetrcClient {
	return New(baseURL, vendorKey, userKey, WithHTTPClient(httpClient))
}

func (c *MetrcClient) send(ctx context.Context, method, path string, queryParams map[string]string, body interface{}) (interface{}, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil { return nil, err }
		bodyReader = bytes.NewReader(jsonBytes)
	}

	u, err := url.Parse(c.BaseURL + path)
	if err != nil { return nil, err }
	q := u.Query()
	for k, v := range queryParams {
		if v != "" {
			q.Set(k, v)
		}
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil { return nil, err }

	req.SetBasicAuth(c.VendorKey, c.UserKey)
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.Logger != nil {
		c.Logger.DebugContext(ctx, "Sending Request", "method", method, "url", u.String())
	}

	resp, err := c.Client.Do(req)
	if err != nil { 
		if c.Logger != nil {
			c.Logger.ErrorContext(ctx, "Request Failed", "error", err)
		}
		return nil, err 
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		if c.Logger != nil {
			c.Logger.WarnContext(ctx, "API Error Response", "status", resp.StatusCode)
		}

		// Try to parse structured error
		bodyBytes, err := io.ReadAll(resp.Body)
		if err == nil {
			var apiErr MetrcError
			if jsonErr := json.Unmarshal(bodyBytes, &apiErr); jsonErr == nil && (apiErr.Message != "" || len(apiErr.ValidationErrors) > 0) {
				apiErr.StatusCode = resp.StatusCode
				return nil, &apiErr
			}
			// Fallback if not valid JSON or empty fields
			return nil, fmt.Errorf("API Error %d: %s", resp.StatusCode, string(bodyBytes))
		}
		
		return nil, fmt.Errorf("API Error: %d", resp.StatusCode)
	}
	if resp.StatusCode == 204 {
		return nil, nil
	}

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
