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

const (
	CsNamespace = "ThunkMetrc.Client"
)

func generateCSharp(groups map[string][]bruno.Request, version string) {
	log.Println("Generating C# SDK...")

	outDir := filepath.Join(outputDir, "csharp", "ThunkMetrc.Client")
	os.MkdirAll(outDir, 0755)

	csprojContent := fmt.Sprintf(`<Project Sdk="Microsoft.NET.Sdk">
	
  <PropertyGroup>
    <TargetFramework>net8.0</TargetFramework>
    <ImplicitUsings>enable</ImplicitUsings>
    <Nullable>enable</Nullable>
    <Version>%s</Version>
    <GeneratePackageOnBuild>true</GeneratePackageOnBuild>
    <PackageId>ThunkMetrc.Client</PackageId>
    <Authors>ThunkMetrc</Authors>
    <Description>Auto-generated C# client for ThunkMetrc API.</Description>
    <RepositoryUrl>https://github.com/thunkmetrc/sdks</RepositoryUrl>
    <PackageLicenseExpression>MIT</PackageLicenseExpression>
    <PackageTags>ThunkMetrc;Cannabis;Compliance;API;SDK</PackageTags>
  </PropertyGroup>


</Project>
`, version)
	os.WriteFile(filepath.Join(outDir, "ThunkMetrc.Client.csproj"), []byte(csprojContent), 0644)

	generateCsClient(outDir, groups)
}

func generateCsClient(dir string, groups map[string][]bruno.Request) {
	mainClient := fmt.Sprintf(`using System;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Threading.Tasks;

namespace %s
{
    public partial class MetrcClient
    {
        private readonly HttpClient _http;
        private readonly string _baseUrl;
        private readonly JsonSerializerOptions _jsonOptions;

        public MetrcClient(string baseUrl, string vendorKey, string userKey, HttpClient? http = null)
        {
            _baseUrl = baseUrl.TrimEnd('/');
            _http = http ?? new HttpClient();
            var basicAuth = Convert.ToBase64String(Encoding.UTF8.GetBytes($"{vendorKey}:{userKey}"));
            _http.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Basic", basicAuth);
            _jsonOptions = new JsonSerializerOptions
            {
                PropertyNamingPolicy = JsonNamingPolicy.CamelCase,
                DefaultIgnoreCondition = JsonIgnoreCondition.WhenWritingNull,
                WriteIndented = true
            };
        }

        private async Task<T> SendAsync<T>(HttpMethod method, string path, object? body = null, Dictionary<string, string>? queryParams = null)
        {
            var url = $"{_baseUrl}{path}";
            if (queryParams != null && queryParams.Count > 0)
            {
                var q = string.Join("&", queryParams.Where(kv => !string.IsNullOrEmpty(kv.Value)).Select(kv => $"{kv.Key}={Uri.EscapeDataString(kv.Value)}"));
                if (!string.IsNullOrEmpty(q)) url += $"?{q}";
            }

            var req = new HttpRequestMessage(method, url);

            if (body != null)
            {
                var json = JsonSerializer.Serialize(body, _jsonOptions);
                req.Content = new StringContent(json, Encoding.UTF8, "application/json");
            }

            var resp = await _http.SendAsync(req);
            resp.EnsureSuccessStatusCode();

            var content = await resp.Content.ReadAsStringAsync();
            if (string.IsNullOrWhiteSpace(content)) return default!;

            return JsonSerializer.Deserialize<T>(content, _jsonOptions)!;
        }

        private async Task SendAsync(HttpMethod method, string path, object? body = null, Dictionary<string, string>? queryParams = null)
        {
             var url = $"{_baseUrl}{path}";
            if (queryParams != null && queryParams.Count > 0)
            {
                var q = string.Join("&", queryParams.Where(kv => !string.IsNullOrEmpty(kv.Value)).Select(kv => $"{kv.Key}={Uri.EscapeDataString(kv.Value)}"));
                if (!string.IsNullOrEmpty(q)) url += $"?{q}";
            }

            var req = new HttpRequestMessage(method, url);

            if (body != null)
            {
                var json = JsonSerializer.Serialize(body, _jsonOptions);
                req.Content = new StringContent(json, Encoding.UTF8, "application/json");
            }

            var resp = await _http.SendAsync(req);
            resp.EnsureSuccessStatusCode();
        }
    }
}
`, CsNamespace)
	os.WriteFile(filepath.Join(dir, "MetrcClient.cs"), []byte(mainClient), 0644)

	for group, reqs := range groups {
		if len(reqs) == 0 {
			continue
		}

		safeGroup := ToPascalCase(cleanName(group))

		sb := strings.Builder{}
		sb.WriteString(fmt.Sprintf("using System;\nusing System.Collections.Generic;\nusing System.Net.Http;\nusing System.Threading.Tasks;\n\nnamespace %s\n{\n    public partial class MetrcClient\n    {\n", CsNamespace))

		for _, req := range reqs {
			methodName := ToPascalCase(cleanName(req.Name))
			methodName = strings.TrimSpace(methodName)

			returnType := "object"
			isVoid := req.Method != "GET"

			sendArgs := "null"
			pathParams := GetPathParams(req.URL)

			var argsList []string
			for _, p := range pathParams {
				argsList = append(argsList, fmt.Sprintf("string %s", p))
			}

			if req.Method == "POST" || req.Method == "PUT" {
				argsList = append(argsList, "object? body = null")
				sendArgs = "body"

				// Still generate the model classes if needed
				if req.BodySchema != nil {
					// Models are generated in a separate loop below
				}
			}

			// Query Params
			var sortedKeys []string
			for k := range req.Params {
				sortedKeys = append(sortedKeys, k)
			}
			sort.Strings(sortedKeys)

			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				// C# params camelCase
				argsList = append(argsList, fmt.Sprintf("string? %s = null", paramName))
			}

			methodParams := strings.Join(argsList, ", ")
			finalName := safeGroup + methodName

			csUrl := req.URL
			for _, p := range pathParams {
				csUrl = strings.ReplaceAll(csUrl, "{"+p+"}", "{Uri.EscapeDataString("+p+")}")
			}

			sb.WriteString(fmt.Sprintf("        /// <summary>\n"))
			docs := cleanDocs(req.Docs)
			if docs != "" {
				for _, line := range strings.Split(docs, "\n") {
					sb.WriteString(fmt.Sprintf("        /// %s\n", line))
				}
			}
			sb.WriteString(fmt.Sprintf("        /// </summary>\n"))
			if isVoid {
				sb.WriteString(fmt.Sprintf("        public async Task %s(%s)\n", finalName, methodParams))
				sb.WriteString("        {\n")
				sb.WriteString("            var queryParams = new Dictionary<string, string>();\n")
				for _, k := range sortedKeys {
					paramName := strings.TrimSuffix(k, "optional")
					sb.WriteString(fmt.Sprintf("            if (!string.IsNullOrEmpty(%s)) queryParams[\"%s\"] = %s;\n", paramName, strings.TrimSuffix(k, "optional"), paramName))
				}
				sb.WriteString(fmt.Sprintf("            await SendAsync(new HttpMethod(\"%s\"), $\"%s\", %s, queryParams);\n", req.Method, csUrl, sendArgs))
				sb.WriteString("        }\n\n")
			} else {
				sb.WriteString(fmt.Sprintf("        public async Task<%s> %s(%s)\n", returnType, finalName, methodParams))
				sb.WriteString("        {\n")
				sb.WriteString("            var queryParams = new Dictionary<string, string>();\n")
				for _, k := range sortedKeys {
					paramName := strings.TrimSuffix(k, "optional")
					sb.WriteString(fmt.Sprintf("            if (!string.IsNullOrEmpty(%s)) queryParams[\"%s\"] = %s;\n", paramName, strings.TrimSuffix(k, "optional"), paramName))
				}
				sb.WriteString(fmt.Sprintf("            return await SendAsync<%s>(new HttpMethod(\"%s\"), $\"%s\", %s, queryParams);\n", returnType, req.Method, csUrl, sendArgs))
				sb.WriteString("        }\n\n")
			}
		}

		sb.WriteString("    }\n")

		sb.WriteString("}\n")

		os.WriteFile(filepath.Join(dir, fmt.Sprintf("MetrcClient.%s.cs", safeGroup)), []byte(sb.String()), 0644)
	}
}
