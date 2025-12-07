package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"unicode"
)

var (
	brunoDir    = "../specs/metrc-bruno"
	typeMapFile = "pkg/metrc/models/type_map.json"
	outputFile  = "pkg/metrc/client/generated_client.go"
)

type GenRequest struct {
	Name       string
	Method     string
	URL        string
	ReturnType string
	Group      string
	IsList     bool
}

func main() {
	flag.StringVar(&brunoDir, "input", brunoDir, "Path to Bruno specs directory")
	flag.StringVar(&typeMapFile, "typemap", typeMapFile, "Path to Type Map JSON")
	flag.StringVar(&outputFile, "output", outputFile, "Path to output Go client file")
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	typeMap, err := loadTypeMap(typeMapFile)
	if err != nil {
		return fmt.Errorf("loading type map: %v", err)
	}

	requests, err := parseRequests(brunoDir, typeMap)
	if err != nil {
		return fmt.Errorf("parsing requests: %v", err)
	}

	if err := generateCode(outputFile, requests); err != nil {
		return fmt.Errorf("generating code: %v", err)
	}

	fmt.Printf("Generated %s (%d methods)\n", outputFile, len(requests))
	return nil
}

func loadTypeMap(path string) (map[string]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var m map[string]string
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func parseRequests(rootDir string, typeMap map[string]string) ([]GenRequest, error) {
	var reqs []GenRequest

	groups, err := os.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		if !group.IsDir() || strings.HasPrefix(group.Name(), ".") {
			continue
		}

		groupName := group.Name()
		files, err := os.ReadDir(filepath.Join(rootDir, groupName))
		if err != nil {
			continue
		}

		for _, f := range files {
			if !strings.HasSuffix(f.Name(), ".bru") {
				continue
			}
			fmt.Printf("DEBUG: Processing request file: %s/%s\n", groupName, f.Name())

			path := filepath.Join(rootDir, groupName, f.Name())
			content, err := os.ReadFile(path)
			if err != nil {
				continue
			}

			metaName := extractMeta(string(content), "name")
			method := extractMethod(string(content))
			url := extractURL(string(content))

			if metaName == "" || method == "" {
				continue
			}

			returnType := "map[string]any"
			isList := false

			key := fmt.Sprintf("%s:%s", groupName, metaName)
			modelName, ok := typeMap[key]

			if !ok {
				modelName, ok = typeMap[metaName]
			}

			if ok && modelName != "" {
				returnType = "models." + modelName
				if strings.ToUpper(method) == "GET" {
					isList = true
					checkUrl := strings.ReplaceAll(url, "{{", "")
					checkUrl = strings.ReplaceAll(checkUrl, "}}", "")

					if strings.Contains(metaName, "ById") || strings.Contains(checkUrl, "{") {
						isList = false
					}
				}
			}

			cleanName := removeMethodSuffix(metaName)

			baseName := sanitizeName(cleanName)

			if len(baseName) > 0 {
				r := []rune(baseName)
				r[0] = unicode.ToUpper(r[0])
				baseName = string(r)
			}

			groupPrefix := sanitizeName(groupName)
			if len(groupPrefix) > 0 {
				r := []rune(groupPrefix)
				r[0] = unicode.ToUpper(r[0])
				groupPrefix = string(r)
			}

			funcName := groupPrefix + "_" + baseName

			reqs = append(reqs, GenRequest{
				Name:       funcName,
				Method:     method,
				URL:        url,
				ReturnType: returnType,
				Group:      groupName,
				IsList:     isList,
			})
		}
	}

	sort.Slice(reqs, func(i, j int) bool {
		return reqs[i].Name < reqs[j].Name
	})

	seen := make(map[string]bool)
	var uniqueReqs []GenRequest
	for _, r := range reqs {
		if seen[r.Name] {
			r.Name = r.Name + sanitizeName(r.Group)

			if seen[r.Name] {
				r.Name = r.Name + r.Method
			}
		}

		seen[r.Name] = true
		uniqueReqs = append(uniqueReqs, r)
	}

	return uniqueReqs, nil
}

func removeMethodSuffix(s string) string {
	s = strings.ReplaceAll(s, "(POST)", "")
	s = strings.ReplaceAll(s, "(GET)", "")
	s = strings.ReplaceAll(s, "(PUT)", "")
	s = strings.ReplaceAll(s, "(DELETE)", "")
	s = strings.ReplaceAll(s, "(post)", "")
	s = strings.ReplaceAll(s, "(get)", "")
	s = strings.ReplaceAll(s, "(put)", "")
	s = strings.ReplaceAll(s, "(delete)", "")
	return strings.TrimSpace(s)
}

func sanitizeName(s string) string {
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "/", "")
	return s
}

func generateCode(path string, reqs []GenRequest) error {
	const tpl = `package client

import (
	"fmt"
	"github.com/thunkmetrc/automation/pkg/metrc/models"
)

// Auto-Generated Client Methods

{{- range . }}

// {{ .Name }} ({{ .Group }})
// {{ .Method }} {{ .URL }}
func (f *Fetcher) {{ .Name }}() ({{ if .IsList }}[]{{ end }}{{ .ReturnType }}, error) {
	{{- if .IsList }}
	return fetchList[{{ .ReturnType }}](f, "{{ .Name }}", "{{ .Method }}", "{{ .URL }}")
	{{- else }}
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[{{ .ReturnType }}](f, "{{ .Name }}", "{{ .Method }}", "{{ .URL }}")
	if err != nil {
		var empty {{ .ReturnType }}
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty {{ .ReturnType }}
	return empty, fmt.Errorf("not found")
	{{- end }}
}
{{- end }}
`

	t, err := template.New("client").Parse(tpl)
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, reqs)
}

func extractMeta(content, key string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, key+":") {
			return strings.TrimSpace(strings.TrimPrefix(line, key+":"))
		}
	}
	return ""
}

func extractMethod(content string) string {
	methods := []string{"get", "post", "put", "delete"}
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(strings.ToLower(line))
		for _, m := range methods {
			if strings.HasPrefix(line, m+" {") {
				return strings.ToUpper(m)
			}
		}
	}
	return ""
}

func extractURL(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "url:") {
			u := strings.TrimSpace(strings.TrimPrefix(line, "url:"))
			return u
		}
	}
	return ""
}
