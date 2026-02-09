package app

import (
	"github.com/thunkier/thunkmetrc/probe/internal/metrc/client"
)

type DataWorkflow interface {
	Name() string
	ShouldRun(cfg Config) bool
	Execute(cfg Config, f *client.Fetcher) (*WorkflowResult, error)
}

type WorkflowResult struct {
	CreatedIDs       []int
	DataGenerated    map[string]int
	Errors           []string
	Outputs          map[string]any
	SkippedEndpoints map[string]string
	EmptyEndpoints   []string
}
