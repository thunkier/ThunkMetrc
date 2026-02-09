package client

import (
	"log"
	"sync"
)

type Config struct {
	BaseURL               string
	APIKey                string // Vendor Key
	UserKey               string // User Key
	FacilityLicenseNumber string
	DryRun                bool
	Logger                Logger
	ErrorReporter         ErrorReporter
	UsedRequests          *sync.Map
	CurrentWorkflow       string
	BlockedEndpoints      []string
	CanUseComplianceLabel bool
	ResponsesDir          string
	SuppressConsoleLog    bool
}

type Logger interface {
	Log(format string, args ...interface{})
	Flush()
}

type ErrorReporter interface {
	RecordError(facility, workflow, method, endpoint string, statusCode int)
}

type defaultLogger struct{}

func (d defaultLogger) Log(format string, args ...interface{}) { log.Printf(format, args...) }
func (d defaultLogger) Flush()                                 {}
