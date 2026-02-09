package models

import "github.com/thunkier/thunkmetrc/tools/pkg/schema"

// Parameter represents a query or path parameter.
type Parameter struct {
	Name        string
	Type        string
	Description string
	Required    bool
}

// Operation represents a single API endpoint.
type Operation struct {
	Name         string
	Description  string
	Method       string
	Path         string
	BodyType     string
	QueryParams  []Parameter
	PathParams   []string
	BodySchema   *schema.Schema
	ReturnSchema *schema.Schema
	// Manifest-derived fields
	ResponseType string // e.g., "Facility", "Package"
	IsArray      bool   // Response is an array
	IsPaginated  bool   // Response uses pagination wrapper
	RequestType  string // e.g., "CreatePackagesRequest"
}

// Service represents a group of operations.
type Service struct {
	Name       string
	Operations []Operation
}
