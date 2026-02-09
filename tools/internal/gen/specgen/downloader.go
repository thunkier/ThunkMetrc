package specgen

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const defaultScraperUserAgent = "ThunkMetrc-SpecGen/1.0 (+https://github.com/thunkier/thunkmetrc)"

type HTTPFetcher struct {
	Client       *http.Client
	MaxRetries   int
	RetryBackoff time.Duration
	UserAgent    string
}

func NewHTTPFetcher(timeout time.Duration, maxRetries int, retryBackoff time.Duration, userAgent string) *HTTPFetcher {
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	if maxRetries < 0 {
		maxRetries = 0
	}
	if retryBackoff <= 0 {
		retryBackoff = 500 * time.Millisecond
	}
	if userAgent == "" {
		userAgent = defaultScraperUserAgent
	}

	return &HTTPFetcher{
		Client: &http.Client{Timeout: timeout},
		MaxRetries:   maxRetries,
		RetryBackoff: retryBackoff,
		UserAgent:    userAgent,
	}
}

func (f *HTTPFetcher) get(url string) ([]byte, int, error) {
	var lastErr error
	var statusCode int

	for attempt := 0; attempt <= f.MaxRetries; attempt++ {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, 0, err
		}
		req.Header.Set("User-Agent", f.UserAgent)
		req.Header.Set("Accept", "application/json, text/html;q=0.9, */*;q=0.8")

		resp, err := f.Client.Do(req)
		if err != nil {
			lastErr = err
			if attempt < f.MaxRetries {
				time.Sleep(f.backoffForAttempt(attempt))
				continue
			}
			return nil, 0, err
		}

		body, readErr := io.ReadAll(resp.Body)
		resp.Body.Close()

		statusCode = resp.StatusCode
		if readErr != nil {
			lastErr = readErr
			if attempt < f.MaxRetries {
				time.Sleep(f.backoffForAttempt(attempt))
				continue
			}
			return nil, statusCode, readErr
		}

		if resp.StatusCode == http.StatusOK {
			return body, statusCode, nil
		}

		lastErr = fmt.Errorf("status code %d", resp.StatusCode)
		if shouldRetryStatus(resp.StatusCode) && attempt < f.MaxRetries {
			time.Sleep(f.backoffForAttempt(attempt))
			continue
		}

		return body, statusCode, lastErr
	}

	if lastErr != nil {
		return nil, statusCode, lastErr
	}
	return nil, statusCode, fmt.Errorf("request failed for %s", url)
}

func (f *HTTPFetcher) backoffForAttempt(attempt int) time.Duration {
	multiplier := 1 << attempt
	return time.Duration(multiplier) * f.RetryBackoff
}

func shouldRetryStatus(statusCode int) bool {
	if statusCode == http.StatusTooManyRequests {
		return true
	}
	return statusCode >= 500 && statusCode <= 599
}

func fetchPostmanSpec(state, url, outDir string, fetcher *HTTPFetcher) error {
	content, status, err := fetcher.get(url)
	if err != nil {
		return fmt.Errorf("error fetching Postman spec for %s (status %d): %w", state, status, err)
	}

	outPath := filepath.Join(outDir, state+".json")
	if err := os.WriteFile(outPath, content, 0644); err != nil {
		return fmt.Errorf("error saving Postman spec for %s: %w", state, err)
	}

	fmt.Printf("  Saved Postman spec for %s\n", state)
	return nil
}

func fetchHTMLSpec(state, url, outDir string, fetcher *HTTPFetcher) error {
	content, status, err := fetcher.get(url)
	if err != nil {
		return fmt.Errorf("error fetching %s (status %d): %w", state, status, err)
	}

	if err := os.WriteFile(filepath.Join(outDir, state+".html"), content, 0644); err != nil {
		return fmt.Errorf("error writing file for %s: %w", state, err)
	}

	return nil
}
