package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/thunkmetrc/automation/pkg/bruno"
)

func generateRust(groups map[string][]bruno.Request, version string) {
	outDir := filepath.Join(outputDir, "rust", "client")
	os.MkdirAll(filepath.Join(outDir, "src"), 0755)

	generateCargoToml(outDir, version)
	generateRustClient(outDir, groups)
}

func generateCargoToml(dir string, version string) {
	content := fmt.Sprintf(`[package]
name = "thunkmetrc-client"
version = "%s"
edition = "2021"
description = "Auto-generated Rust client for ThunkMetrc"
license = "MIT"
repository = "https://github.com/thunkmetrc/sdks"
keywords = ["thunkmetrc", "cannabis", "compliance", "api", "sdk"]
categories = ["api-bindings", "web-programming::http-client"]

[dependencies]
reqwest = { version = "0.11", features = ["json", "blocking"] }
serde = { version = "1.0", features = ["derive"] }
serde_json = "1.0"
tokio = { version = "1.0", features = ["full"] }
urlencoding = "2"
`, version)
	os.WriteFile(filepath.Join(dir, "Cargo.toml"), []byte(content), 0644)
}

func generateRustClient(dir string, groups map[string][]bruno.Request) {
	sb := strings.Builder{}
	var modelBuilder strings.Builder
	sb.WriteString("use reqwest::Client;\n")
	sb.WriteString("use serde_json::Value;\n")
	sb.WriteString("use std::error::Error;\n\n")

	sb.WriteString("#[derive(Clone)]\n")
	sb.WriteString("pub struct MetrcClient {\n")
	sb.WriteString("    client: Client,\n")
	sb.WriteString("    base_url: String,\n")
	sb.WriteString("    vendor_key: String,\n")
	sb.WriteString("    user_key: String,\n")
	sb.WriteString("}\n\n")

	sb.WriteString("impl MetrcClient {\n")
	sb.WriteString("    pub fn new(base_url: &str, vendor_key: &str, user_key: &str) -> Self {\n")
	sb.WriteString("        MetrcClient {\n")
	sb.WriteString("            client: Client::new(),\n")
	sb.WriteString("            base_url: base_url.trim_end_matches('/').to_string(),\n")
	sb.WriteString("            vendor_key: vendor_key.to_string(),\n")
	sb.WriteString("            user_key: user_key.to_string(),\n")
	sb.WriteString("        }\n")
	sb.WriteString("    }\n\n")

	sb.WriteString("    async fn send(&self, method: reqwest::Method, path: &str, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {\n")
	sb.WriteString("        let url = format!(\"{}{}\", self.base_url, path);\n")
	sb.WriteString("        let mut req = self.client.request(method, &url);\n")
	sb.WriteString("        req = req.basic_auth(&self.vendor_key, Some(&self.user_key));\n\n")

	sb.WriteString("        if let Some(b) = body {\n")
	sb.WriteString("            req = req.json(b);\n")
	sb.WriteString("        }\n\n")

	sb.WriteString("        let resp = req.send().await?;\n")
	sb.WriteString("        let status = resp.status();\n")
	sb.WriteString("        if !status.is_success() {\n")
	sb.WriteString("            return Err(format!(\"API Error: {}\", status).into());\n")
	sb.WriteString("        }\n")
	sb.WriteString("        if status == reqwest::StatusCode::NO_CONTENT {\n")
	sb.WriteString("            return Ok(None);\n")
	sb.WriteString("        }\n")
	sb.WriteString("        let json: Value = resp.json().await?;\n")
	sb.WriteString("        Ok(Some(json))\n")
	sb.WriteString("    }\n\n")

	for group, reqs := range groups {
		safeGroup := ToSnakeCase(cleanName(group))
		for _, req := range reqs {
			methodName := ToSnakeCase(cleanName(req.Name))

			rustMethodName := fmt.Sprintf("%s_%s", safeGroup, methodName)

			methodEnum := "POST"
			if req.Method == "GET" {
				methodEnum = "GET"
			}
			if req.Method == "PUT" {
				methodEnum = "PUT"
			}
			if req.Method == "DELETE" {
				methodEnum = "DELETE"
			}

			pathParams := GetPathParams(req.URL)

			var argsList []string
			for _, p := range pathParams {
				argsList = append(argsList, fmt.Sprintf("%s: &str", ToSnakeCase(p)))
			}

			// Query Params
			var sortedKeys []string
			for k := range req.Params {
				sortedKeys = append(sortedKeys, k)
			}
			sort.Strings(sortedKeys)

			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				rustParamName := ToSnakeCase(paramName)
				argsList = append(argsList, fmt.Sprintf("%s: Option<String>", rustParamName))
			}

			if req.Method == "POST" || req.Method == "PUT" {
				// Decoupling: Use Option<&Value> for body
				argsList = append(argsList, "body: Option<&Value>")

			} else {
				argsList = append(argsList, "body: Option<&Value>")
			}

			sigArgs := strings.Join(argsList, ", ")
			if len(sigArgs) > 0 {
				sigArgs = "&self, " + sigArgs
			} else {
				sigArgs = "&self"
			}

			rustUrl := req.URL
			re := regexp.MustCompile(`\{[a-zA-Z0-9_]+\}`)
			rustUrl = re.ReplaceAllString(rustUrl, "{}")

			pathFmtStr := rustUrl
			var pathArgs []string
			for _, p := range pathParams {
				pathArgs = append(pathArgs, fmt.Sprintf("urlencoding::encode(%s).as_ref()", ToSnakeCase(p)))
			}

			sb.WriteString(fmt.Sprintf("    /// %s %s\n", req.Method, req.Name))
			docs := cleanDocs(req.Docs)
			if docs != "" {
				for _, line := range strings.Split(docs, "\n") {
					sb.WriteString(fmt.Sprintf("    /// %s\n", line))
				}
				sb.WriteString("    ///\n")
			}
			sb.WriteString(fmt.Sprintf("    pub async fn %s(%s) -> Result<Option<Value>, Box<dyn Error>> {\n", rustMethodName, sigArgs))

			pathArgsStr := ""
			if len(pathArgs) > 0 {
				pathArgsStr = ", " + strings.Join(pathArgs, ", ")
			}

			sb.WriteString(fmt.Sprintf("        let mut path = format!(\"%s\"%s);\n", pathFmtStr, pathArgsStr))

			if len(sortedKeys) > 0 {
				sb.WriteString("        let mut query_params = Vec::new();\n")
				for _, k := range sortedKeys {
					paramName := strings.TrimSuffix(k, "optional")
					rustParamName := ToSnakeCase(paramName)
					sb.WriteString(fmt.Sprintf("        if let Some(p) = %s {\n", rustParamName))
					sb.WriteString(fmt.Sprintf("            query_params.push(format!(\"%s={}\", urlencoding::encode(&p)));\n", paramName))
					sb.WriteString("        }\n")
				}
				sb.WriteString("        if !query_params.is_empty() {\n")
				sb.WriteString("            path.push_str(\"?\");\n")
				sb.WriteString("            path.push_str(&query_params.join(\"&\"));\n")
				sb.WriteString("        }\n")
			}

			sb.WriteString(fmt.Sprintf("        self.send(reqwest::Method::%s, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await\n", methodEnum))
			sb.WriteString("    }\n\n")
		}
	}

	sb.WriteString("}\n")

	sb.WriteString(modelBuilder.String())

	os.WriteFile(filepath.Join(dir, "src", "lib.rs"), []byte(sb.String()), 0644)
}
