package internalgen

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/thunkier/thunkmetrc/tools/pkg/bruno"
	"github.com/thunkier/thunkmetrc/tools/pkg/schema"
)

var (
	brunoDir      string
	schemasDir    string
	modelsDir     string
	clientDir     string
	generatedFile = "generated_client.go"
)

type Config struct {
	BrunoDir   string
	SchemasDir string
	ModelsDir  string
	ClientDir  string
}

func Run(cfg Config) error {
	brunoDir = cfg.BrunoDir
	schemasDir = cfg.SchemasDir
	modelsDir = cfg.ModelsDir
	clientDir = cfg.ClientDir

	if err := run(); err != nil {
		return err
	}
	return nil
}

func run() error {
	if err := os.MkdirAll(modelsDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(clientDir, 0755); err != nil {
		return err
	}

	manifestPath := filepath.Join(schemasDir, "manifest.json")
	manifestData, err := os.ReadFile(manifestPath)
	if err != nil {
		return fmt.Errorf("reading manifest: %v", err)
	}
	var manifest map[string]*schema.ManifestEntry
	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		return fmt.Errorf("parsing manifest: %v", err)
	}
	fmt.Printf("Loaded manifest with %d entries\n", len(manifest))

	uniqueSchemas := make(map[string]bool)
	for _, entry := range manifest {
		if entry.SchemaFile != "" {
			uniqueSchemas[entry.SchemaFile] = true
		}
		if entry.RequestSchemaFile != "" {
			uniqueSchemas[entry.RequestSchemaFile] = true
		}
	}

	schemasByGroup := make(map[string][]string)
	for file := range uniqueSchemas {
		parts := strings.Split(file, "/")
		if len(parts) != 2 {
			continue
		}
		schemasByGroup[parts[0]] = append(schemasByGroup[parts[0]], file)
	}

	for group, files := range schemasByGroup {
		var groupStructs []schema.GoStruct
		for _, file := range files {
			schemaPath := filepath.Join(schemasDir, file)
			data, err := os.ReadFile(schemaPath)
			if err != nil {
				log.Printf("Warning: skipping schema %s: %v", file, err)
				continue
			}

			var js schema.JSONSchema
			if err := json.Unmarshal(data, &js); err != nil {
				return fmt.Errorf("parsing schema %s: %v", file, err)
			}
			s := js.ToSchema()
			baseName := strings.TrimSuffix(filepath.Base(file), ".schema.json")
			structs := schema.ToGoStructs(s, baseName)
			groupStructs = append(groupStructs, structs...)
		}

		if len(groupStructs) > 0 {
			if err := writeModelsFile(group, groupStructs); err != nil {
				return err
			}
		}
	}

	collection, err := bruno.ParseCollection(brunoDir)
	if err != nil {
		return fmt.Errorf("parsing bruno collection: %v", err)
	}

	var genReqs []GenRequest
	for _, req := range collection.Requests {
		if strings.HasPrefix(req.Group, ".") {
			continue
		}
		bruFileName := strings.TrimSuffix(filepath.Base(req.FilePath), ".bru")
		manifestKey := fmt.Sprintf("%s/%s", req.Group, bruFileName)

		entry := manifest[manifestKey]

		// Use extracted params from Bruno parser
		params := req.PathParams
		docParams := req.DocParams
		permissions := req.Permissions

		// Determine types from manifest entry
		returnType := "any"
		isList := false
		isPaginated := false
		bodyType := "any"
		hasBody := false

		if req.Body != "" && (req.Method == "POST" || req.Method == "PUT" || req.Method == "PATCH") {
			hasBody = true
		}

		if entry != nil && entry.ResponseType != "" {
			if entry.IsArray {
				returnType = "models." + entry.ResponseType
				isList = true
			} else if entry.IsPaginated {
				returnType = "models." + entry.ResponseType
				isPaginated = true
			} else {
				returnType = "models." + entry.ResponseType
			}
		} else {
			if !isList && !isPaginated {
				returnType = "map[string]any"
			}
		}

		if hasBody {
			if entry != nil && entry.RequestType != "" {
				if entry.IsRequestArray {
					bodyType = "[]models." + entry.RequestType
				} else {
					bodyType = "models." + entry.RequestType
				}
			}
		}

		var ipDocParams []ParamInfo
		for _, dp := range docParams {
			ipDocParams = append(ipDocParams, ParamInfo{
				Name:        dp.Name,
				Optional:    dp.Optional,
				Description: dp.Description,
			})
		}

		allParams := make([]string, 0, len(params)+len(ipDocParams))
		allParams = append(allParams, params...)
		for _, dp := range ipDocParams {
			allParams = append(allParams, dp.Name)
		}

		genReqs = append(genReqs, GenRequest{
			Name:        req.Name,
			Method:      req.Method,
			URL:         req.URL,
			ReturnType:  returnType,
			Group:       req.Group,
			IsList:      isList,
			IsPaginated: isPaginated,
			Params:      params,
			DocParams:   ipDocParams,
			AllParams:   allParams,
			Permissions: permissions,
			HasBody:     hasBody,
			BodyType:    bodyType,
		})
	}

	if err := generateFetcher(clientDir); err != nil {
		return fmt.Errorf("generating fetcher: %v", err)
	}

	if err := generateClient(clientDir); err != nil {
		return fmt.Errorf("generating client: %v", err)
	}

	reqsByGroup := make(map[string][]GenRequest)
	for _, req := range genReqs {
		reqsByGroup[req.Group] = append(reqsByGroup[req.Group], req)
	}

	for group, reqs := range reqsByGroup {
		// Sort requests in group
		sort.Slice(reqs, func(i, j int) bool {
			return reqs[i].Name < reqs[j].Name
		})

		if err := generateGroupFile(clientDir, group, reqs); err != nil {
			return fmt.Errorf("generating group %s: %v", group, err)
		}
	}

	fmt.Printf("Generated client files for %d groups\n", len(reqsByGroup))

	return nil
}

type GenRequest struct {
	Name        string
	Method      string
	URL         string
	ReturnType  string
	Group       string
	IsList      bool
	IsPaginated bool
	Params      []string
	DocParams   []ParamInfo
	AllParams   []string
	Permissions []string
	HasBody     bool
	BodyType    string
}

func writeModelsFile(group string, structs []schema.GoStruct) error {
	filename := strings.TrimPrefix(group, "_")
	filename = filepath.Join(modelsDir, filename+".go")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, "package models\n\n")

	written := make(map[string]bool)
	uniqueStructs := make([]schema.GoStruct, 0)
	for _, s := range structs {
		if !written[s.Name] {
			uniqueStructs = append(uniqueStructs, s)
			written[s.Name] = true
		}
	}

	fmt.Fprint(f, schema.RenderGoCode(uniqueStructs))
	return nil
}

func generateFetcher(dir string) error {
	const tpl = `package client

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
	Data          []T ` + "`json:\"Data\"`" + `
	Page          int ` + "`json:\"Page\"`" + `
	PageSize      int ` + "`json:\"PageSize\"`" + `
	Total         int ` + "`json:\"Total\"`" + `
	TotalPages    int ` + "`json:\"TotalPages\"`" + `
	TotalRecords  int ` + "`json:\"TotalRecords\"`" + `
	RecordsOnPage int ` + "`json:\"RecordsOnPage\"`" + `
	CurrentPage   int ` + "`json:\"CurrentPage\"`" + `
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
		Data []T ` + "`json:\"Data\"`" + `
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
`
	filename := filepath.Join(dir, "fetcher.go")
	return os.WriteFile(filename, []byte(tpl), 0644)
}

func generateGroupFile(dir, group string, reqs []GenRequest) error {
	const tpl = `package client

import (
{{ if .HasStrings }}	"strings"
{{ end -}}
{{ if .HasFmt }}	"fmt"
{{ end -}}
{{ if .HasModels }}	"github.com/thunkier/thunkmetrc/probe/pkg/metrc/models"
{{ end -}}
)

{{- range .Reqs }}

// {{ .Name }} ({{ .Group }})
// {{ .Method }} {{ .URL }}
{{- range .DocParams }}
//   {{.Name}} ({{if not .Optional}}required{{else}}optional{{end}}): {{.Description}}
{{- end }}
{{- $allParams := .AllParams -}}
{{- if .IsPaginated }}
func (f *Fetcher) {{ .Name }}({{ if .HasBody }}body {{ .BodyType }}{{ if $allParams }}, {{ end }}{{ end }}{{- range $i, $p := $allParams }}{{ if $i }}, {{ end }}{{ $p | toParam }} string{{ end }}) (PaginatedResponse[{{ .ReturnType }}], error) {
{{- else if .IsList }}
func (f *Fetcher) {{ .Name }}({{ if .HasBody }}body {{ .BodyType }}{{ if $allParams }}, {{ end }}{{ end }}{{- range $i, $p := $allParams }}{{ if $i }}, {{ end }}{{ $p | toParam }} string{{ end }}) ([]{{ .ReturnType }}, error) {
{{- else }}
func (f *Fetcher) {{ .Name }}({{ if .HasBody }}body {{ .BodyType }}{{ if $allParams }}, {{ end }}{{ end }}{{- range $i, $p := $allParams }}{{ if $i }}, {{ end }}{{ $p | toParam }} string{{ end }}) ({{ .ReturnType }}, error) {
{{- end }}
	url := "{{ .URL }}"
	{{- range .Params }}
	url = strings.ReplaceAll(url, "{"+"{{ . }}"+"}", {{ . | toParam }})
	{{- end }}
	{{- if .DocParams }}
	
	queryParams := make([]string, 0)
	{{- range .DocParams }}
	if {{ .Name | toParam }} != "" {
		queryParams = append(queryParams, fmt.Sprintf("{{ .Name }}=%s", {{ .Name | toParam }}))
	}
	{{- end }}

	if len(queryParams) > 0 {
		if strings.Contains(url, "?") {
			url += "&" + strings.Join(queryParams, "&")
		} else {
			url += "?" + strings.Join(queryParams, "&")
		}
	}
	{{- end }}
	
	{{- if .IsPaginated }}
	return fetchPaginated[{{ .ReturnType }}](f, "{{ .Group }}", "{{ .Name }}", "{{ .Method }}", url, {{ if .HasBody }}body{{ else }}nil{{ end }})
	{{- else if .IsList }}
	return fetchList[{{ .ReturnType }}](f, "{{ .Group }}", "{{ .Name }}", "{{ .Method }}", url, {{ if .HasBody }}body{{ else }}nil{{ end }})
	{{- else }}
	return fetchOne[{{ .ReturnType }}](f, "{{ .Group }}", "{{ .Name }}", "{{ .Method }}", url, {{ if .HasBody }}body{{ else }}nil{{ end }})
	{{- end }}
}
{{- end }}
`
	funcMap := template.FuncMap{
		"toExported": toExported,
		"toParam":    toParam,
	}
	t, err := template.New("client").Funcs(funcMap).Parse(tpl)
	if err != nil {
		return err
	}

	hasStrings := false
	hasFmt := false
	hasModels := false
	for _, r := range reqs {
		if len(r.AllParams) > 0 {
			hasStrings = true
		}
		if len(r.DocParams) > 0 {
			hasFmt = true
		}
		if strings.Contains(r.ReturnType, "models.") || strings.Contains(r.BodyType, "models.") {
			hasModels = true
		}
	}

	data := struct {
		Reqs       []GenRequest
		HasStrings bool
		HasFmt     bool
		HasModels  bool
	}{
		Reqs:       reqs,
		HasStrings: hasStrings,
		HasFmt:     hasFmt,
		HasModels:  hasModels,
	}

	filename := filepath.Join(dir, strings.ToLower(group)+".go")
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return t.Execute(f, data)
}

func generateClient(dir string) error {
	const tpl = `package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"os"
	"encoding/json"

	"github.com/thunkier/thunkmetrc/tools/pkg/bruno"
)

func ExecuteRequest(cfg Config, req *bruno.Request, userApiKey string) string {
	logger := cfg.Logger
	if logger == nil {
		logger = defaultLogger{}
	}

	url := req.URL
	url = strings.ReplaceAll(url, "{{baseUrl}}", cfg.BaseURL)
	url = strings.ReplaceAll(url, ":licenseNumber", cfg.FacilityLicenseNumber)

	if !strings.Contains(url, "licenseNumber=") {
		if strings.Contains(url, "?") {
			url += "&licenseNumber=" + cfg.FacilityLicenseNumber
		} else {
			url += "?licenseNumber=" + cfg.FacilityLicenseNumber
		}
	}

	for _, blocked := range cfg.BlockedEndpoints {
		if strings.Contains(req.Name, blocked) {
			if !cfg.SuppressConsoleLog {
				logger.Log("[BLOCKED] %s is in blocked endpoints list", req.Name)
			}
			return ""
		}
	}

	if cfg.UsedRequests != nil {
		key := fmt.Sprintf("%s/%s", req.Group, req.Name)
		cfg.UsedRequests.Store(key, true)
	}

	// Basic Auth: VendorKey:UserKey
	vendorKey := cfg.APIKey
	userKey := cfg.UserKey

	// Allow override if provided (though mostly unused now)
	if userApiKey != "" {
		userKey = userApiKey
	}

	var bodyReader io.Reader
	if req.Body != "" {
		bodyReader = bytes.NewBufferString(req.Body)
	}

	httpReq, err := http.NewRequest(req.Method, url, bodyReader)
	if err != nil {
		if cfg.ErrorReporter != nil {
			cfg.ErrorReporter.RecordError(cfg.FacilityLicenseNumber, cfg.CurrentWorkflow, req.Method, req.URL, 0)
		}
		return ""
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.SetBasicAuth(vendorKey, userKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		if cfg.ErrorReporter != nil {
			cfg.ErrorReporter.RecordError(cfg.FacilityLicenseNumber, cfg.CurrentWorkflow, req.Method, req.URL, 0)
		}
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if cfg.ErrorReporter != nil {
			cfg.ErrorReporter.RecordError(cfg.FacilityLicenseNumber, cfg.CurrentWorkflow, req.Method, req.URL, resp.StatusCode)
		}
		return ""
	}

	if resp.StatusCode >= 400 {
		if cfg.ErrorReporter != nil {
			cfg.ErrorReporter.RecordError(cfg.FacilityLicenseNumber, cfg.CurrentWorkflow, req.Method, req.URL, resp.StatusCode)
		}
		// Log error for debugging
		if !cfg.SuppressConsoleLog {
			logger.Log("[ERROR] %s returned %d: %s", req.Name, resp.StatusCode, string(body))
		}
		return ""
	}

	if cfg.ResponsesDir != "" && req.Method == "GET" && len(body) > 0 {
		saveResponse(cfg.ResponsesDir, req.Group, req.Name, body)
	}

	return string(body)
}

func saveResponse(dir, group, name string, data []byte) {
	groupDir := filepath.Join(dir, group)
	if err := os.MkdirAll(groupDir, 0755); err != nil {
		return
	}

	var formatted bytes.Buffer
	if json.Indent(&formatted, data, "", "  ") == nil {
		data = formatted.Bytes()
	}

	filePath := filepath.Join(groupDir, name+".json")
	os.WriteFile(filePath, data, 0644)
}
`
	filename := filepath.Join(dir, generatedFile)
	return os.WriteFile(filename, []byte(tpl), 0644)
}

