package specgen

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/thunkier/thunkmetrc/tools/internal/app"
)

type PostmanCollection struct {
	Info PostmanInfo   `json:"info"`
	Item []PostmanItem `json:"item"`
}

type PostmanInfo struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type PostmanItem struct {
	Name    string          `json:"name"`
	Item    []PostmanItem   `json:"item,omitempty"`
	Request *PostmanRequest `json:"request,omitempty"`
}

type PostmanRequest struct {
	Method string       `json:"method"`
	URL    PostmanURL   `json:"url"`
	Body   *PostmanBody `json:"body,omitempty"`
}

type PostmanURL struct {
	Raw   string         `json:"raw,omitempty"`
	Host  []string       `json:"host,omitempty"`
	Path  []string       `json:"path,omitempty"`
	Query []PostmanQuery `json:"query,omitempty"`
}

type PostmanQuery struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PostmanBody struct {
	Mode string `json:"mode,omitempty"`
	Raw  string `json:"raw,omitempty"`
}

func ParsePostmanCollection(r io.Reader, cfg app.ScraperConfig) ([]ParsedOperation, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read postman collection: %w", err)
	}

	var collection PostmanCollection
	if err := json.Unmarshal(data, &collection); err != nil {
		return nil, fmt.Errorf("failed to parse postman collection: %w", err)
	}

	var operations []ParsedOperation
	seen := make(map[string]bool)

	processItems(collection.Item, "", &operations, seen, cfg)

	return operations, nil
}

func processItems(items []PostmanItem, parentGroup string, ops *[]ParsedOperation, seen map[string]bool, cfg app.ScraperConfig) {
	for _, item := range items {
		if len(item.Item) > 0 {
			group := item.Name
			if strings.ToLower(group) == "v1" {
				continue
			}
			if strings.ToLower(group) == "v2" {
				processItems(item.Item, parentGroup, ops, seen, cfg)
			} else {
				newGroup := group
				if parentGroup != "" && parentGroup != "v2" {
					newGroup = parentGroup + " " + group
				}
				processItems(item.Item, newGroup, ops, seen, cfg)
			}
			continue
		}

		if item.Request == nil {
			continue
		}

		url := buildURL(item.Request.URL)
		if url == "" {
			continue
		}

		if isV1Endpoint(url) {
			continue
		}

		method := strings.ToUpper(item.Request.Method)
		key := method + " " + url
		if seen[key] {
			continue
		}
		seen[key] = true

		group := parentGroup
		if group == "" {
			group = "Default"
		}

		op := ParsedOperation{
			Group:  group,
			Name:   GenerateName(group, method, url, 0, cfg),
			Method: method,
			URL:    url,
		}

		for _, q := range item.Request.URL.Query {
			op.Params = append(op.Params, Param{
				Name:  q.Key,
				Value: q.Value,
			})
		}

		if item.Request.Body != nil && item.Request.Body.Raw != "" {
			op.Body = item.Request.Body.Raw
		}

		*ops = append(*ops, op)
	}
}

func buildURL(url PostmanURL) string {
	if url.Raw != "" {
		raw := url.Raw
		if idx := strings.Index(raw, "}}"); idx > 0 {
			raw = raw[idx+2:]
		}
		return strings.TrimPrefix(raw, "/")
	}

	var parts []string
	for _, h := range url.Host {
		if idx := strings.Index(h, "}}"); idx > 0 {
			h = h[idx+2:]
		}
		h = strings.TrimPrefix(h, "/")
		if h != "" {
			parts = append(parts, h)
		}
	}
	for _, p := range url.Path {
		if p != "" {
			p = strings.ReplaceAll(p, "{{", "{")
			p = strings.ReplaceAll(p, "}}", "}")
			parts = append(parts, p)
		}
	}

	if len(parts) == 0 {
		return ""
	}

	return "/" + strings.Join(parts, "/")
}
