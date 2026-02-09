package sdks

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/common"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/csharp"
	golang "github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/go"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/java"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/kotlin"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/manifests"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/models"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/python"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/rust"
	"github.com/thunkier/thunkmetrc/tools/internal/gen/sdks/typescript"
	"github.com/thunkier/thunkmetrc/tools/pkg/bruno"
	"github.com/thunkier/thunkmetrc/tools/pkg/jsonschema"
	"github.com/thunkier/thunkmetrc/tools/pkg/schema"
	"gopkg.in/yaml.v3"
)

var (
	specsDir     string
	responsesDir string
	outputDir    string
)

type Config struct {
	SpecsDir     string
	ResponsesDir string
	OutputDir    string
}

func Run(cfg Config) {
	specsDir = cfg.SpecsDir
	responsesDir = cfg.ResponsesDir
	outputDir = cfg.OutputDir

	run()
}

func run() {
	log.Println("Starting SDK Generation...")
	log.Printf("Specs: %s\n", specsDir)
	log.Printf("Output: %s\n", outputDir)

	collection, err := bruno.ParseCollection(specsDir)
	if err != nil {
		log.Fatalf("Failed to parse specs: %v", err)
	}
	log.Printf("Found %d requests in collection.\n", len(collection.Requests))

	manifestPath := filepath.Join(specsDir, "../schemas/public/manifest.json")
	var manifest map[string]*schema.ManifestEntry
	if data, err := os.ReadFile(manifestPath); err == nil {
		if err := json.Unmarshal(data, &manifest); err != nil {
			log.Printf("[WARN] Failed to parse manifest: %v", err)
		} else {
			log.Printf("Loaded manifest with %d entries.\n", len(manifest))
		}
	} else {
		log.Printf("[WARN] Failed to load manifest from %s: %v", manifestPath, err)
	}

	groups := make(map[string][]bruno.Request)
	for i, req := range collection.Requests {
		req.URL = strings.ReplaceAll(req.URL, "{{baseUrl}}", "")
		req.URL = strings.ReplaceAll(req.URL, "{{url}}", "")

		if req.BodyType == "json" && req.Body != "" {
			var data interface{}
			if err := json.Unmarshal([]byte(req.Body), &data); err == nil {
				req.BodySchema = schema.InferFromJSON(data)
			} else {
				log.Printf("[WARN] Failed to parse body JSON for %s: %v", req.Name, err)
			}
		}

		bruFileName := strings.TrimSuffix(filepath.Base(req.FilePath), ".bru")
		manifestKey := fmt.Sprintf("%s/%s", req.Group, bruFileName)

		if entry, ok := manifest[manifestKey]; ok && entry.SchemaFile != "" {

			schemaPath := filepath.Join(filepath.Dir(manifestPath), entry.SchemaFile)

			if data, err := os.ReadFile(schemaPath); err == nil {
				var jsonData interface{}
				if err := json.Unmarshal(data, &jsonData); err == nil {
					req.ReturnSchema = schema.InferFromJSON(jsonData)
				} else {
					log.Printf("[WARN] Failed to parse response JSON for %s: %v", req.Name, err)
				}
			} else {
				log.Printf("[WARN] Failed to read schema file %s: %v", schemaPath, err)
			}
		}

		// Update collection with modified req (BodySchema/ReturnSchema)
		collection.Requests[i] = req

		if req.Group != "" {
			groups[req.Group] = append(groups[req.Group], req)
		}
	}
	log.Printf("Grouped REQUESTS into %d categories.\n", len(groups))

	for groupName, reqs := range groups {
		sort.Slice(reqs, func(i, j int) bool {
			return reqs[i].Name < reqs[j].Name
		})

		seen := make(map[string]int)
		for i, req := range reqs {
			baseName := common.CleanName(req.Name)
			checkName := strings.ToLower(baseName)

			if count, exists := seen[checkName]; exists {
				seen[checkName] = count + 1
				reqs[i].Name = fmt.Sprintf("%s_%d", baseName, count+1)
			} else {
				seen[checkName] = 1
			}
		}
		groups[groupName] = reqs
	}

	versionsPath := "../versions.yaml"
	versionsBytes, err := os.ReadFile(versionsPath)
	if err != nil {
		log.Fatalf("Failed to read versions.yaml: %v", err)
	}

	var verCfg struct {
		SDKVersion   string                            `yaml:"sdk_version"`
		Dependencies map[string]map[string]interface{} `yaml:"dependencies"`
	}
	if err := yaml.Unmarshal(versionsBytes, &verCfg); err != nil {
		log.Fatalf("Failed to parse versions.yaml: %v", err)
	}

	version := strings.TrimSpace(verCfg.SDKVersion)
	if version == "" {
		log.Fatalf("sdk_version missing in versions.yaml")
	}
	log.Printf("Generating SDKs version: %s\n", version)

	// Update handwritten manifests
	manifests.UpdateManifests(outputDir, version)

	deps := make(map[string]string)
	for lang, langDeps := range verCfg.Dependencies {
		for key, val := range langDeps {
			switch v := val.(type) {
			case string:
				deps[lang+"."+key] = v
			case map[string]interface{}:
				for subKey, subVal := range v {
					if strVal, ok := subVal.(string); ok {
						deps[lang+"."+key+"."+subKey] = strVal
					}
				}
			}
		}
	}
	log.Printf("Loaded %d dependency versions from versions.yaml\n", len(deps))
	services := convertToServices(groups, manifest)

	// Parse response schema files into ModelIR for typed model generation
	responseModels := make(map[string]*jsonschema.ModelIR)
	// schemasDir should be .../generated/schemas/public
	schemasDir := filepath.Join(specsDir, "../schemas/public")
	for _, entry := range manifest {
		if entry.ResponseType != "" && entry.SchemaFile != "" {
			// Skip PaginatedResponse - this is specific to language implementation
			if entry.ResponseType == "PaginatedResponse" {
				continue
			}
			// Skip if already parsed
			if _, exists := responseModels[entry.ResponseType]; exists {
				continue
			}
			schemaPath := filepath.Join(schemasDir, entry.SchemaFile)
			if model, err := jsonschema.ParseSchemaFile(schemaPath); err == nil {
				// Use manifest ResponseType as the model name
				model.Name = entry.ResponseType
				responseModels[entry.ResponseType] = model
			} else {
				log.Printf("[WARN] Failed to parse schema %s for response type %s: %v", schemaPath, entry.ResponseType, err)
			}
		}
	}
	log.Printf("Parsed %d response model schemas.\n", len(responseModels))

	ctx := &common.GeneratorContext{
		OutputDir:      outputDir,
		ResponsesDir:   responsesDir,
		Version:        version,
		Services:       services,
		ResponseModels: responseModels,
		Deps:           deps,
	}

	generators := []common.SDKGenerator{
		&typescript.ClientGenerator{},
		&csharp.ClientGenerator{},
		&python.ClientGenerator{},
		&golang.ClientGenerator{},
		&java.ClientGenerator{},
		&kotlin.ClientGenerator{},
		&rust.ClientGenerator{},
		&csharp.WrapperGenerator{},
		&csharp.MainGenerator{},
		&golang.WrapperGenerator{},
		&python.WrapperGenerator{},
		&typescript.WrapperGenerator{},
		&java.WrapperGenerator{},
		&kotlin.WrapperGenerator{},
		&rust.WrapperGenerator{},
		&typescript.MainGenerator{},
		&python.MainGenerator{},
		&golang.MainGenerator{},
		&java.MainGenerator{},
		&kotlin.MainGenerator{},
		&rust.MainGenerator{},
	}

	for _, g := range generators {
		log.Printf("[NEW] Running %s Generator...\n", g.Name())
		if err := g.Generate(ctx); err != nil {
			log.Fatalf("Failed to run %s generator: %v", g.Name(), err)
		}
	}

	log.Println("SDK Generation Complete.")
}

func convertToServices(groups map[string][]bruno.Request, manifest map[string]*schema.ManifestEntry) []models.Service {
	var services []models.Service

	var keys []string
	for k := range groups {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	log.Printf("Groups found: %v\n", keys)

	for _, k := range keys {
		reqs := groups[k]
		service := models.Service{
			Name: k,
		}

		for _, req := range reqs {
			op := models.Operation{
				Name:         req.Name,
				Description:  req.Docs,
				Method:       req.Method,
				Path:         req.URL,
				BodyType:     req.BodyType,
				PathParams:   req.PathParams,
				BodySchema:   req.BodySchema,
				ReturnSchema: req.ReturnSchema,
			}

			// Populate fields from manifest keys = Group / Name
			manifestKey := fmt.Sprintf("%s/%s", k, req.Name)
			if entry, ok := manifest[manifestKey]; ok {
				op.ResponseType = entry.ResponseType
				op.IsArray = entry.IsArray
				op.IsPaginated = entry.IsPaginated
				op.RequestType = entry.RequestType
			}

			var queryParams []models.Parameter

			var paramKeys []string
			for pk := range req.Params {
				paramKeys = append(paramKeys, pk)
			}
			sort.Strings(paramKeys)

			for _, pk := range paramKeys {
				queryParams = append(queryParams, models.Parameter{
					Name: pk,
					Type: "string",
				})
			}
			op.QueryParams = queryParams
			service.Operations = append(service.Operations, op)
		}
		services = append(services, service)
	}

	return services
}
