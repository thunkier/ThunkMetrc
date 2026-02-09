package common

import (
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
	"github.com/thunkier/thunkmetrc/tools/pkg/jsonschema"
)

type GeneratorContext struct {
	OutputDir      string
	ResponsesDir   string
	Version        string
	Services       []models.Service
	ResponseModels map[string]*jsonschema.ModelIR
	Deps           map[string]string
}

type SDKGenerator interface {
	Name() string
	Generate(ctx *GeneratorContext) error
}
