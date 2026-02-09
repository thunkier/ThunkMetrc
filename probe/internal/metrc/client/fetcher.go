package client

import (
	"encoding/json"
	"fmt"

	"github.com/thunkier/thunkmetrc/tools/pkg/bruno"
)


type Fetcher struct {
	Cfg Config
}

func NewFetcher(cfg Config) *Fetcher {
	return &Fetcher{Cfg: cfg}
}


type PaginatedResponse[T any] struct {
	Data          []T `json:"Data"`
	Page          int `json:"Page"`
	PageSize      int `json:"PageSize"`
	Total         int `json:"Total"`
	TotalPages    int `json:"TotalPages"`
	TotalRecords  int `json:"TotalRecords"`
	RecordsOnPage int `json:"RecordsOnPage"`
	CurrentPage   int `json:"CurrentPage"`
}

func fetchList[T any](f *Fetcher, group, name, method, url string, body any) ([]T, error) {
	req := &bruno.Request{Name: name, Method: method, URL: url, Group: group}
	if body != nil {
		reqBody, _ := json.Marshal(body)
		req.Body = string(reqBody)
	}
	respBody := ExecuteRequest(f.Cfg, req, "")
	if len(respBody) == 0 {
		return []T{}, nil
	}
	

	var wrapper struct {
		Data []T `json:"Data"`
	}

	if err := json.Unmarshal([]byte(respBody), &wrapper); err == nil && wrapper.Data != nil {
		return wrapper.Data, nil
	}
	

	var list []T
	if err := json.Unmarshal([]byte(respBody), &list); err == nil {
		return list, nil
	}
	
	return nil, fmt.Errorf("failed to parse response for %s", name)
}

func fetchPaginated[T any](f *Fetcher, group, name, method, url string, body any) (PaginatedResponse[T], error) {
	req := &bruno.Request{Name: name, Method: method, URL: url, Group: group}
	if body != nil {
		reqBody, _ := json.Marshal(body)
		req.Body = string(reqBody)
	}
	respBody := ExecuteRequest(f.Cfg, req, "")
	var result PaginatedResponse[T]
	if len(respBody) == 0 {
		return result, nil
	}
	if err := json.Unmarshal([]byte(respBody), &result); err != nil {
		return result, fmt.Errorf("unmarshal: %w", err)
	}
	return result, nil
}

func fetchOne[T any](f *Fetcher, group, name, method, url string, body any) (T, error) {
	req := &bruno.Request{Name: name, Method: method, URL: url, Group: group}
	if body != nil {
		reqBody, _ := json.Marshal(body)
		req.Body = string(reqBody)
	}
	respBody := ExecuteRequest(f.Cfg, req, "")
	var result T
	if len(respBody) == 0 {
		return result, nil
	}
	if err := json.Unmarshal([]byte(respBody), &result); err != nil {
		return result, fmt.Errorf("unmarshal: %w", err)
	}
	return result, nil
}
