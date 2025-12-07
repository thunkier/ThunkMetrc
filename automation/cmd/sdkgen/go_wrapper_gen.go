package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/thunkmetrc/automation/pkg/bruno"
)

func generateGoWrapper(groups map[string][]bruno.Request, version string) {
	log.Println("Generating Go Wrapper...")
	outDir := filepath.Join(outputDir, "go", "wrapper")
	os.MkdirAll(outDir, 0755)

	generateGoWrapperMod(outDir)
	generateGoRateLimiter(outDir)
	generateGoWrapperCode(outDir, groups)
}

func generateGoRateLimiter(dir string) {
	content := `package wrapper

import (
	"context"
	"strings"
	"sync"
	"time"
)

type RateLimiterConfig struct {
	Enabled                     bool
	MaxGetPerSecondPerFacility  float64
	MaxGetPerSecondIntegrator   float64
	MaxConcurrentGetPerFacility int
	MaxConcurrentGetIntegrator  int
}

func DefaultConfig() RateLimiterConfig {
	return RateLimiterConfig{
		Enabled:                     false,
		MaxGetPerSecondPerFacility:  50,
		MaxGetPerSecondIntegrator:   150,
		MaxConcurrentGetPerFacility: 10,
		MaxConcurrentGetIntegrator:  30,
	}
}

type TokenBucket struct {
	rate       float64
	capacity   float64
	tokens     float64
	lastRefill time.Time
	mu         sync.Mutex
}

func NewTokenBucket(rate float64, capacity float64) *TokenBucket {
	return &TokenBucket{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

func (tb *TokenBucket) Wait(ctx context.Context) error {
	tb.mu.Lock()
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsed * tb.rate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now

	if tb.tokens >= 1.0 {
		tb.tokens -= 1.0
		tb.mu.Unlock()
		return nil
	}

	// Calculate wait time
	missing := 1.0 - tb.tokens
	waitSecs := missing / tb.rate
	wait := time.Duration(waitSecs * float64(time.Second))
	tb.mu.Unlock()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(wait):
		// Recurse to ensure we successfully take a token (race condition in simple impl, but good enough for sdk)
		// For strictness, we should re-acquire lock and check again.
		return tb.Wait(ctx)
	}
}

type FacilityLimiter struct {
	rate *TokenBucket
	sem  chan struct{}
}

type RateLimiter struct {
	config         RateLimiterConfig
	integratorRate *TokenBucket
	integratorSem  chan struct{}
	facilitiesMu   sync.Mutex
	facilities     map[string]*FacilityLimiter
}

func NewRateLimiter(config RateLimiterConfig) *RateLimiter {
	return &RateLimiter{
		config:         config,
		integratorRate: NewTokenBucket(config.MaxGetPerSecondIntegrator, config.MaxGetPerSecondIntegrator),
		integratorSem:  make(chan struct{}, config.MaxConcurrentGetIntegrator),
		facilities:     make(map[string]*FacilityLimiter),
	}
}

func (r *RateLimiter) Execute(ctx context.Context, facility string, isGet bool, op func() (interface{}, error)) (interface{}, error) {
	if !r.config.Enabled || !isGet {
		return op()
	}

	// 1. Acquire Integrator Concurrent Permit
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case r.integratorSem <- struct{}{}:
		defer func() { <-r.integratorSem }()
	}

	// 2. Acquire Facility Concurrent Permit (if facility known)
	if facility != "" {
		fl := r.getFacilityLimiter(facility)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case fl.sem <- struct{}{}:
			defer func() { <-fl.sem }()
		}
	}

	// 3. Rate Limit Wait (Global + Facility)
	if err := r.integratorRate.Wait(ctx); err != nil {
		return nil, err
	}
	if facility != "" {
		fl := r.getFacilityLimiter(facility)
		if err := fl.rate.Wait(ctx); err != nil {
			return nil, err
		}
	}

	// 4. Retry Loop for 429
	for {
		res, err := op()
		if err != nil {
			// Check for 429
			// The Client returns generic error string currently: "API Error: %d"
			errMsg := err.Error()
			if strings.Contains(errMsg, "API Error: 429") {
				// Retry
				// Ideally parse Retry-After header, but client doesn't expose it in current error return
				// We fallback to exponential backoff or default wait
				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				case <-time.After(1 * time.Second): // Default backoff
					continue
				}
			}
			return nil, err
		}
		return res, nil
	}
}

func (r *RateLimiter) getFacilityLimiter(facility string) *FacilityLimiter {
	r.facilitiesMu.Lock()
	defer r.facilitiesMu.Unlock()

	if fl, ok := r.facilities[facility]; ok {
		return fl
	}

	fl := &FacilityLimiter{
		rate: NewTokenBucket(r.config.MaxGetPerSecondPerFacility, r.config.MaxGetPerSecondPerFacility),
		sem:  make(chan struct{}, r.config.MaxConcurrentGetPerFacility),
	}
	r.facilities[facility] = fl
	return fl
}
`
	os.WriteFile(filepath.Join(dir, "ratelimiter.go"), []byte(content), 0644)
}

func generateGoWrapperMod(dir string) {
	content := `module github.com/thunkmetrc/sdks/go/wrapper

go 1.21

require (
	github.com/thunkmetrc/sdks/go/client v0.0.0
)

replace github.com/thunkmetrc/sdks/go/client => ../client
`
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte(content), 0644)
}

func generateGoWrapperCode(dir string, groups map[string][]bruno.Request) {
	sb := strings.Builder{}
	sb.WriteString("package wrapper\n\n")
	sb.WriteString("import (\n")
	sb.WriteString("\t\"context\"\n")
	sb.WriteString("\t\"github.com/thunkmetrc/sdks/go/client\"\n")
	sb.WriteString(")\n\n")

	sb.WriteString("type MetrcWrapper struct {\n")
	sb.WriteString("\tClient *client.MetrcClient\n")
	sb.WriteString("\tRateLimiter *RateLimiter\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func New(client *client.MetrcClient) *MetrcWrapper {\n")
	sb.WriteString("	return &MetrcWrapper{\n")
	sb.WriteString("		Client: client,\n")
	sb.WriteString("		RateLimiter: NewRateLimiter(DefaultConfig()),\n")
	sb.WriteString("	}\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func NewClient(baseUrl, vendorKey, userKey string) *MetrcWrapper {\n")
	sb.WriteString("    return New(client.New(baseUrl, vendorKey, userKey))\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func NewWithConfig(client *client.MetrcClient, config RateLimiterConfig) *MetrcWrapper {\n")
	sb.WriteString("	return &MetrcWrapper{\n")
	sb.WriteString("		Client: client,\n")
	sb.WriteString("		RateLimiter: NewRateLimiter(config),\n")
	sb.WriteString("	}\n")
	sb.WriteString("}\n\n")

	modelsSb := strings.Builder{}
	modelsSb.WriteString("package wrapper\n\n")

	for group, reqs := range groups {
		safeGroup := ToPascalCase(cleanName(group))
		for _, req := range reqs {
			methodName := ToPascalCase(cleanName(req.Name))
			goMethodName := safeGroup + methodName

			pathParams := GetPathParams(req.URL)

			var argsList []string
			for _, p := range pathParams {
				argsList = append(argsList, fmt.Sprintf("%s string", ToCamelCase(p)))
			}

			// Query Params
			var sortedKeys []string
			for k := range req.Params {
				sortedKeys = append(sortedKeys, k)
			}
			sort.Strings(sortedKeys)

			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				goParamName := ToCamelCase(paramName)
				argsList = append(argsList, fmt.Sprintf("%s string", goParamName))
			}

			callBodyArg := "nil"

			if req.Method == "POST" || req.Method == "PUT" {
				if req.BodySchema != nil {
					structName := fmt.Sprintf("%s%sRequest", safeGroup, methodName)
					targetSchema := req.BodySchema
					if targetSchema.Type == bruno.TypeArray {
						targetSchema = targetSchema.ItemType
						structName = fmt.Sprintf("%sItem", fmt.Sprintf("%s%sRequest", safeGroup, methodName))
						generateGoStruct(structName, targetSchema, &modelsSb)

						argsList = append(argsList, fmt.Sprintf("body []%s", structName))
					} else {
						generateGoStruct(structName, targetSchema, &modelsSb)
						argsList = append(argsList, fmt.Sprintf("body %s", structName))
					}
					callBodyArg = "body"
				} else {
					argsList = append(argsList, "body interface{}")
					callBodyArg = "body"
				}
			}

			sigArgs := strings.Join(argsList, ", ")

			var clientCallArgs []string
			for _, p := range pathParams {
				clientCallArgs = append(clientCallArgs, ToCamelCase(p))
			}
			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				goParamName := ToCamelCase(paramName)
				clientCallArgs = append(clientCallArgs, goParamName)
			}
			if req.Method == "POST" || req.Method == "PUT" {
				clientCallArgs = append(clientCallArgs, callBodyArg)
			}

			clientCallStr := strings.Join(clientCallArgs, ", ")

			// DETERMINE FACILITY from args
			facilityArg := "\"\""
			for _, p := range pathParams {
				if strings.ToLower(cleanName(p)) == "licensenumber" {
					facilityArg = ToCamelCase(p)
					break
				}
			}

			isGet := "false"
			if req.Method == "GET" {
				isGet = "true"
			}

			sb.WriteString(fmt.Sprintf("// %s %s\n", req.Method, req.Name))
			docs := cleanDocs(req.Docs)
			if docs != "" {
				for _, line := range strings.Split(docs, "\n") {
					sb.WriteString(fmt.Sprintf("// %s\n", line))
				}
			}
			sb.WriteString(fmt.Sprintf("func (w *MetrcWrapper) %s(%s) (interface{}, error) {\n", goMethodName, sigArgs))
			sb.WriteString(fmt.Sprintf("\treturn w.RateLimiter.Execute(context.TODO(), %s, %s, func() (interface{}, error) {\n", facilityArg, isGet))
			sb.WriteString(fmt.Sprintf("\t\treturn w.Client.%s(%s)\n", goMethodName, clientCallStr))
			sb.WriteString("\t})\n")
			sb.WriteString("}\n\n")
		}
	}

	os.WriteFile(filepath.Join(dir, "wrapper.go"), []byte(sb.String()), 0644)
	os.WriteFile(filepath.Join(dir, "models.go"), []byte(modelsSb.String()), 0644)
}

func generateGoStruct(name string, schema *bruno.Schema, sb *strings.Builder) {
	if schema.Type == bruno.TypeObject {
		sb.WriteString(fmt.Sprintf("type %s struct {\n", name))

		nested := make(map[string]*bruno.Schema)
		var keys []string
		for k := range schema.Properties {
			keys = append(keys, k)
		}

		for _, propName := range keys {
			propSchema := schema.Properties[propName]
			goField := ToPascalCase(cleanName(propName))
			if len(goField) > 0 && goField[0] >= '0' && goField[0] <= '9' {
				goField = "_" + goField
			}

			goType, nestedSchema, nestedName := resolveGoType(propSchema, name, goField)
			if nestedSchema != nil {
				nested[nestedName] = nestedSchema
			}

			sb.WriteString(fmt.Sprintf("\t%s %s `json:\"%s,omitempty\"`\n", goField, goType, propName))
		}
		sb.WriteString("}\n\n")

		for subName, subSchema := range nested {
			generateGoStruct(subName, subSchema, sb)
		}
	}
}

func resolveGoType(s *bruno.Schema, parentName, propName string) (string, *bruno.Schema, string) {
	switch s.Type {
	case bruno.TypeString:
		return "string", nil, ""
	case bruno.TypeInt:
		return "int", nil, ""
	case bruno.TypeFloat:
		return "float64", nil, ""
	case bruno.TypeBool:
		return "bool", nil, ""
	case bruno.TypeArray:
		innerType, innerSchema, innerName := "interface{}", (*bruno.Schema)(nil), ""
		if s.ItemType != nil {
			innerType, innerSchema, innerName = resolveGoType(s.ItemType, parentName, propName)
		}
		return fmt.Sprintf("[]%s", innerType), innerSchema, innerName
	case bruno.TypeObject:
		subStructName := fmt.Sprintf("%s_%s", parentName, propName)
		return subStructName, s, subStructName
	default:
		return "interface{}", nil, ""
	}
}
