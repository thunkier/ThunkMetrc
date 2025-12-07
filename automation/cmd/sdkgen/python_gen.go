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

func generatePython(groups map[string][]bruno.Request, version string) {
	log.Println("Generating Python SDK...")

	outDir := filepath.Join(outputDir, "python", "client")
	os.MkdirAll(outDir, 0755)
	os.MkdirAll(filepath.Join(outDir, "src", "thunkmetrc", "client"), 0755)

	generatePyProject(outDir, version)
	generatePyClient(outDir, groups)
}

func generatePyProject(dir string, version string) {
	content := fmt.Sprintf(`[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[project]
name = "thunkmetrc-client"
version = "%s"
description = "Auto-generated Python client for ThunkMetrc"
readme = "README.md"
requires-python = ">=3.10"
license = "MIT"
authors = [
  { name = "ThunkMetrc", email = "dev@thunkmetrc.com" },
]
classifiers = [
  "Development Status :: 4 - Beta",
  "Programming Language :: Python",
  "Programming Language :: Python :: 3.10",
  "Programming Language :: Python :: 3.11",
  "Programming Language :: Python :: Implementation :: CPython",
  "Programming Language :: Python :: Implementation :: PyPy",
]
dependencies = [
    "httpx>=0.24.0",
]

[project.urls]
Documentation = "https://github.com/thunkmetrc/sdks#readme"
Issues = "https://github.com/thunkmetrc/sdks/issues"
Source = "https://github.com/thunkmetrc/sdks"

[tool.hatch.build.targets.wheel]
packages = ["src/thunkmetrc"]
`, version)
	os.WriteFile(filepath.Join(dir, "pyproject.toml"), []byte(content), 0644)

	readmeContent := `# ThunkMetrc Client

Auto-generated Python client for ThunkMetrc API.

## Installation

` + "```bash" + `
pip install thunkmetrc-client
` + "```" + `

## Usage

swag
`
	os.WriteFile(filepath.Join(dir, "README.md"), []byte(readmeContent), 0644)
}

func generatePyClient(dir string, groups map[string][]bruno.Request) {
	srcDir := filepath.Join(dir, "src", "thunkmetrc", "client")

	os.WriteFile(filepath.Join(filepath.Join(dir, "src", "thunkmetrc"), "__init__.py"), []byte(""), 0644)
	os.WriteFile(filepath.Join(srcDir, "__init__.py"), []byte("from .client import MetrcClient\n"), 0644)

	sb := strings.Builder{}
	sb.WriteString(`import httpx
import base64
import urllib.parse
from typing import Any, Optional, Dict, List
from typing_extensions import TypedDict 
from typing import TypedDict

# --- Models ---


class MetrcClient:
    def __init__(self, base_url: str, vendor_key: str, user_key: str, client: Optional[httpx.Client] = None):
        self.base_url = base_url.rstrip('/')
        self.vendor_key = vendor_key
        self.user_key = user_key
        self.client = client or httpx.Client()
        self.client.auth = (vendor_key, user_key)

    def _send(self, method: str, path: str, body: Any = None, query_params: Dict[str, Any] = None) -> Any:
        url = f"{self.base_url}{path}"
        response = self.client.request(method, url, json=body, params=query_params)
        response.raise_for_status()
        if response.status_code == 204:
            return None
        return response.json()

`)

	var modelBuilder strings.Builder

	for group, reqs := range groups {
		safeGroup := ToSnakeCase(cleanName(group))
		for _, req := range reqs {
			methodName := ToSnakeCase(cleanName(req.Name))
			pyMethodName := fmt.Sprintf("%s_%s", safeGroup, methodName)

			pathParams := GetPathParams(req.URL)

			var argsList []string
			for _, p := range pathParams {
				argsList = append(argsList, fmt.Sprintf("%s: str", ToSnakeCase(p)))
			}

			if req.Method == "POST" || req.Method == "PUT" {
				// Decoupling: Use 'Any' type for body
				argsList = append(argsList, "body: Any = None")

			} else {
				argsList = append(argsList, "body: Any = None")
			}

			sigArgs := strings.Join(argsList, ", ")
			if len(sigArgs) > 0 {
				sigArgs = "self, " + sigArgs
			} else {
				sigArgs = "self"
			}

			pyUrl := req.URL
			for _, p := range pathParams {
				snake := ToSnakeCase(p)
				pyUrl = strings.ReplaceAll(pyUrl, "{"+p+"}", "{urllib.parse.quote("+snake+")}")
			}

			// Query Params
			var queryArgs []string
			var sortedKeys []string
			for k := range req.Params {
				sortedKeys = append(sortedKeys, k)
			}
			sort.Strings(sortedKeys)

			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				snake := ToSnakeCase(paramName)
				queryArgs = append(queryArgs, fmt.Sprintf("%s: Optional[str] = None", snake))
			}

			if len(queryArgs) > 0 {
				if sigArgs != "" {
					sigArgs += ", "
				}
				sigArgs += strings.Join(queryArgs, ", ")
			}

			sb.WriteString(fmt.Sprintf("    def %s(%s) -> Any:\n", pyMethodName, sigArgs))
			sb.WriteString(fmt.Sprintf("        \"\"\"\n"))
			docs := cleanDocs(req.Docs)
			if docs != "" {
				singleLineDocs := strings.ReplaceAll(docs, "\n", "\n        ")
				sb.WriteString(fmt.Sprintf("        %s\n\n", singleLineDocs))
			}
			sb.WriteString(fmt.Sprintf("        %s %s\n        \"\"\"\n", req.Method, req.Name))

			sb.WriteString("        query_params = {}\n")
			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				snake := ToSnakeCase(paramName)
				sb.WriteString(fmt.Sprintf("        if %s is not None:\n", snake))
				sb.WriteString(fmt.Sprintf("            query_params[\"%s\"] = %s\n", paramName, snake))
			}

			sb.WriteString(fmt.Sprintf("        path = f\"%s\"\n", pyUrl))
			sb.WriteString(fmt.Sprintf("        return self._send(\"%s\", path, body, query_params)\n\n", req.Method))
		}
	}

	finalContent := strings.Replace(sb.String(), "# --- Models ---", "# --- Models ---\n"+modelBuilder.String(), 1)
	os.WriteFile(filepath.Join(srcDir, "client.py"), []byte(finalContent), 0644)
}
