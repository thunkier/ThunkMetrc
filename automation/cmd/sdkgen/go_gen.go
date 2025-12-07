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

func generateGo(groups map[string][]bruno.Request, version string) {
	log.Println("Generating Go SDK...")
	outDir := filepath.Join(outputDir, "go", "client")
	os.MkdirAll(outDir, 0755)

	generateGoMod(outDir)
	generateGoClient(outDir, groups)
}

func generateGoMod(dir string) {
	content := `module github.com/thunkmetrc/sdks/go/client

go 1.21
`
	os.WriteFile(filepath.Join(dir, "go.mod"), []byte(content), 0644)
}

func generateGoClient(dir string, groups map[string][]bruno.Request) {
	sb := strings.Builder{}
	sb.WriteString("package client\n\n")
	sb.WriteString("import (\n")
	sb.WriteString("\t\"bytes\"\n")
	sb.WriteString("\t\"encoding/json\"\n")
	sb.WriteString("\t\"fmt\"\n")
	sb.WriteString("\t\"io\"\n")
	sb.WriteString("\t\"net/http\"\n")
	sb.WriteString("\t\"net/url\"\n")
	sb.WriteString("\t\"strings\"\n")
	sb.WriteString(")\n\n")

	sb.WriteString("type MetrcClient struct {\n")
	sb.WriteString("\tBaseURL   string\n")
	sb.WriteString("\tVendorKey string\n")
	sb.WriteString("\tUserKey   string\n")
	sb.WriteString("\tClient  *http.Client\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func New(baseURL, vendorKey, userKey string) *MetrcClient {\n")
	sb.WriteString("\treturn &MetrcClient{\n")
	sb.WriteString("\t\tBaseURL:   strings.TrimRight(baseURL, \"/\"),\n")
	sb.WriteString("\t\tVendorKey: vendorKey,\n")
	sb.WriteString("\t\tUserKey:   userKey,\n")
	sb.WriteString("\t\tClient:    &http.Client{},\n")
	sb.WriteString("\t}\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func (c *MetrcClient) send(method, path string, queryParams map[string]string, body interface{}) (interface{}, error) {\n")
	sb.WriteString("\tvar bodyReader io.Reader\n")
	sb.WriteString("\tif body != nil {\n")
	sb.WriteString("\t\tjsonBytes, err := json.Marshal(body)\n")
	sb.WriteString("\t\tif err != nil { return nil, err }\n")
	sb.WriteString("\t\tbodyReader = bytes.NewReader(jsonBytes)\n")
	sb.WriteString("\t}\n\n")

	sb.WriteString("\tu, err := url.Parse(c.BaseURL + path)\n")
	sb.WriteString("\tif err != nil { return nil, err }\n")
	sb.WriteString("\tq := u.Query()\n")
	sb.WriteString("\tfor k, v := range queryParams {\n")
	sb.WriteString("\t\tif v != \"\" {\n")
	sb.WriteString("\t\t\tq.Set(k, v)\n")
	sb.WriteString("\t\t}\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\tu.RawQuery = q.Encode()\n\n")

	sb.WriteString("\treq, err := http.NewRequest(method, u.String(), bodyReader)\n")
	sb.WriteString("\tif err != nil { return nil, err }\n\n")

	sb.WriteString("\treq.SetBasicAuth(c.VendorKey, c.UserKey)\n")
	sb.WriteString("\tif body != nil {\n")
	sb.WriteString("\t\treq.Header.Set(\"Content-Type\", \"application/json\")\n")
	sb.WriteString("\t}\n\n")

	sb.WriteString("\tresp, err := c.Client.Do(req)\n")
	sb.WriteString("\tif err != nil { return nil, err }\n")
	sb.WriteString("\tdefer resp.Body.Close()\n\n")

	sb.WriteString("\tif resp.StatusCode >= 400 {\n")
	sb.WriteString("\t\treturn nil, fmt.Errorf(\"API Error: %d\", resp.StatusCode)\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\tif resp.StatusCode == 204 {\n")
	sb.WriteString("\t\treturn nil, nil\n")
	sb.WriteString("\t}\n\n")

	sb.WriteString("\tvar result interface{}\n")
	sb.WriteString("\tif err := json.NewDecoder(resp.Body).Decode(&result); err != nil {\n")
	sb.WriteString("\t\treturn nil, err\n")
	sb.WriteString("\t}\n")
	sb.WriteString("\treturn result, nil\n")
	sb.WriteString("}\n\n")

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

			sendBodyArg := "nil"
			if req.Method == "POST" || req.Method == "PUT" {
				argsList = append(argsList, "body interface{}")
				sendBodyArg = "body"
			}

			sigArgs := strings.Join(argsList, ", ")

			goUrl := req.URL
			for _, p := range pathParams {
				goUrl = strings.ReplaceAll(goUrl, "{"+p+"}", "\"+url.QueryEscape("+ToCamelCase(p)+")+\"")
			}

			sb.WriteString(fmt.Sprintf("// %s %s\n", req.Method, req.Name))
			docs := cleanDocs(req.Docs)
			if docs != "" {
				for _, line := range strings.Split(docs, "\n") {
					sb.WriteString(fmt.Sprintf("// %s\n", line))
				}
			}
			sb.WriteString(fmt.Sprintf("func (c *MetrcClient) %s(%s) (interface{}, error) {\n", goMethodName, sigArgs))

			sb.WriteString("\tqueryParams := make(map[string]string)\n")
			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				goParamName := ToCamelCase(paramName)
				sb.WriteString(fmt.Sprintf("\tif %s != \"\" {\n", goParamName))
				sb.WriteString(fmt.Sprintf("\t\tqueryParams[\"%s\"] = %s\n", paramName, goParamName))
				sb.WriteString("\t}\n")
			}

			sb.WriteString(fmt.Sprintf("\treturn c.send(\"%s\", \"%s\", queryParams, %s)\n", req.Method, goUrl, sendBodyArg))
			sb.WriteString("}\n\n")
		}
	}

	os.WriteFile(filepath.Join(dir, "client.go"), []byte(sb.String()), 0644)
}
