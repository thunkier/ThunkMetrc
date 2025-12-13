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

func generateTypeScript(groups map[string][]bruno.Request, version string) {
	log.Println("Generating TypeScript SDK...")

	outDir := filepath.Join(outputDir, "typescript", "client")
	os.MkdirAll(outDir, 0755)

	generateTsPackageJson(outDir, version)
	generateTsConfig(outDir)
	generateTsClient(outDir, groups)
}

func generateTsPackageJson(dir string, version string) {
	content := fmt.Sprintf(`{
  "name": "@thunkier/thunkmetrc-client",
  "version": "%s",
  "description": "Auto-generated TypeScript client for ThunkMetrc",
  "main": "dist/index.js",
  "types": "dist/index.d.ts",
  "files": [
    "dist"
  ],
  "scripts": {
    "build": "tsc",
    "prepublishOnly": "npm run build"
  },
  "publishConfig": {
    "access": "public"
  },
  "repository": {
    "type": "git",
    "url": "https://github.com/thunkmetrc/sdks.git"
  },
  "keywords": [
    "thunkmetrc",
    "cannabis",
    "compliance",
    "api",
    "sdk"
  ],
  "author": "ThunkMetrc",
  "license": "MIT",
  "bugs": {
    "url": "https://github.com/thunkmetrc/sdks/issues"
  },
  "homepage": "https://github.com/thunkmetrc/sdks#readme",
  "dependencies": {
    "axios": "^1.6.0"
  },
  "devDependencies": {
    "typescript": "^5.0.0"
  }
}
`, version)
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(content), 0644)
}

func generateTsConfig(dir string) {
	content := `{
  "compilerOptions": {
    "target": "ES2020",
    "module": "commonjs",
    "declaration": true,
    "outDir": "./dist",
    "strict": true,
    "esModuleInterop": true
  },
  "include": ["src/**/*"]
}
`
	os.WriteFile(filepath.Join(dir, "tsconfig.json"), []byte(content), 0644)
}

func generateTsClient(dir string, groups map[string][]bruno.Request) {
	if err := os.MkdirAll(filepath.Join(dir, "src"), 0755); err != nil {
		log.Printf("Warning: failed to create src dir: %v", err)
	}

	sb := strings.Builder{}
	sb.WriteString("import axios, { AxiosInstance, AxiosResponse } from 'axios';\n\n")
	sb.WriteString("export class MetrcClient {\n")
	sb.WriteString("  private client: AxiosInstance;\n\n")
	sb.WriteString("  constructor(baseUrl: string, vendorKey: string, userKey: string) {\n")
	sb.WriteString("    this.client = axios.create({\n")
	sb.WriteString("      baseURL: baseUrl,\n")
	sb.WriteString("      auth: {\n")
	sb.WriteString("        username: vendorKey,\n")
	sb.WriteString("        password: userKey\n")
	sb.WriteString("      }\n")
	sb.WriteString("    });\n")
	sb.WriteString("  }\n\n")

	var interfaces strings.Builder

	for group, reqs := range groups {
		safeGroup := ToCamelCase(cleanName(group))
		for _, req := range reqs {
			methodName := ToPascalCase(cleanName(req.Name))
			tsMethodName := safeGroup + methodName

			pathParams := GetPathParams(req.URL)

			var argsList []string
			for _, p := range pathParams {
				argsList = append(argsList, fmt.Sprintf("%s: string", p))
			}

			if req.Method == "POST" || req.Method == "PUT" {
				// Decoupling: Use 'any' type for body
				argsList = append(argsList, "body?: any")

			} else {
				argsList = append(argsList, "body?: any")
			}

			// Query Params
			var sortedKeys []string
			for k := range req.Params {
				sortedKeys = append(sortedKeys, k)
			}
			sort.Strings(sortedKeys)

			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				argsList = append(argsList, fmt.Sprintf("%s?: string", paramName))
			}

			methodParams := strings.Join(argsList, ", ")

			tsUrl := req.URL
			for _, p := range pathParams {
				tsUrl = strings.ReplaceAll(tsUrl, "{"+p+"}", "${encodeURIComponent("+p+")}")
			}

			sb.WriteString(fmt.Sprintf("  /**\n"))
			docs := cleanDocs(req.Docs)
			if docs != "" {
				for _, line := range strings.Split(docs, "\n") {
					sb.WriteString(fmt.Sprintf("   * %s\n", line))
				}
				sb.WriteString("   *\n")
			}
			sb.WriteString(fmt.Sprintf("   * %s %s\n   */\n", req.Method, req.Name))
			sb.WriteString(fmt.Sprintf("  public async %s(%s): Promise<any> {\n", tsMethodName, methodParams))

			sb.WriteString("    const query = new URLSearchParams();\n")
			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				sb.WriteString(fmt.Sprintf("    if (%s) query.append('%s', %s);\n", paramName, paramName, paramName))
			}
			sb.WriteString("    const queryStr = query.toString();\n")
			sb.WriteString(fmt.Sprintf("    const fullUrl = queryStr ? `%s?${queryStr}` : `%s`;\n", tsUrl, tsUrl))

			sb.WriteString(fmt.Sprintf("    const { data } = await this.client.%s(fullUrl, body);\n", strings.ToLower(req.Method)))
			sb.WriteString("    return data;\n")
			sb.WriteString("  }\n\n")
		}
	}

	sb.WriteString("}\n\n")
	sb.WriteString("// --- Request Models ---\n")
	sb.WriteString(interfaces.String())

	log.Printf("Writing TS Client index.ts to %s with %d groups", filepath.Join(dir, "src", "index.ts"), len(groups))
	if err := os.WriteFile(filepath.Join(dir, "src", "index.ts"), []byte(sb.String()), 0644); err != nil {
		log.Fatalf("Failed to write TS index.ts: %v", err)
	}
}
