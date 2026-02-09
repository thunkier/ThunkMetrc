package client

import (
	"strings"
	"fmt"
)

// CreateIntegratorSetup (Sandbox)
// POST {{baseUrl}}/sandbox/v2/integrator/setup
//   userKey (optional): An existing user key to reuse for integrator setup. Provide this if you already have a user key that you want to continue using. If you don't already have a user key, or you would like a new one, do not provide a value.
func (f *Fetcher) CreateIntegratorSetup(userKey string) (map[string]any, error) {
	url := "{{baseUrl}}/sandbox/v2/integrator/setup"
	
	queryParams := make([]string, 0)
	if userKey != "" {
		queryParams = append(queryParams, fmt.Sprintf("userKey=%s", userKey))
	}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	return fetchOne[map[string]any](f, "Sandbox", "CreateIntegratorSetup", "POST", url, nil)
}
