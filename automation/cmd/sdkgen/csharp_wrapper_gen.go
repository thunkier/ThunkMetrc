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

func generateCSharpWrapper(groups map[string][]bruno.Request, version string) {
	log.Println("Generating C# Wrapper...")
	outDir := filepath.Join(outputDir, "csharp", "ThunkMetrc.Wrapper")
	os.MkdirAll(outDir, 0755)

	generateCsWrapperProj(outDir, version)
	generateCsRateLimiter(outDir)
	generateCsWrapperCode(outDir, groups)
}

func generateCsWrapperProj(dir string, version string) {
	content := fmt.Sprintf(`<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <TargetFramework>net8.0</TargetFramework>
    <ImplicitUsings>enable</ImplicitUsings>
    <Nullable>enable</Nullable>
    <Version>%s</Version>
    <GeneratePackageOnBuild>true</GeneratePackageOnBuild>
    <PackageId>ThunkMetrc.Wrapper</PackageId>
    <Authors>ThunkMetrc</Authors>
    <Description>Type-safe wrapper for ThunkMetrc C# client.</Description>
    <PackageTags>ThunkMetrc;Cannabis;Compliance;API;SDK;Wrapper</PackageTags>
  </PropertyGroup>

  <ItemGroup>
    <ProjectReference Include="..\ThunkMetrc.Client\ThunkMetrc.Client.csproj" />
    <PackageReference Include="System.Threading.RateLimiting" Version="8.0.0" />
  </ItemGroup>

</Project>
`, version)
	os.WriteFile(filepath.Join(dir, "ThunkMetrc.Wrapper.csproj"), []byte(content), 0644)
}

func generateCsRateLimiter(dir string) {
	content := `using System;
using System.Collections.Concurrent;
using System.Net.Http;
using System.Threading;
using System.Threading.RateLimiting;
using System.Threading.Tasks;

namespace ThunkMetrc.Wrapper
{
    public class RateLimiterConfig
    {
        public bool Enabled { get; set; } = false;
        public int MaxGetPerSecondPerFacility { get; set; } = 50;
        public int MaxGetPerSecondIntegrator { get; set; } = 150;
        public int MaxConcurrentGetPerFacility { get; set; } = 10;
        public int MaxConcurrentGetIntegrator { get; set; } = 30;
    }

    public class MetrcRateLimiter
    {
        private readonly RateLimiterConfig _config;
        
        // Integrator Limits
        private readonly TokenBucketRateLimiter _integratorRateLimiter;
        private readonly ConcurrencyLimiter _integratorConcurrencyLimiter;

        // Facility Limits
        private readonly ConcurrentDictionary<string, TokenBucketRateLimiter> _facilityRateLimiters = new();
        private readonly ConcurrentDictionary<string, ConcurrencyLimiter> _facilityConcurrencyLimiters = new();

        public MetrcRateLimiter(RateLimiterConfig? config = null)
        {
            _config = config ?? new RateLimiterConfig();

            // Setup global integrator limits
            _integratorRateLimiter = new TokenBucketRateLimiter(new TokenBucketRateLimiterOptions
            {
                TokenLimit = _config.MaxGetPerSecondIntegrator,
                QueueProcessingOrder = QueueProcessingOrder.OldestFirst,
                QueueLimit = int.MaxValue,
                ReplenishmentPeriod = TimeSpan.FromSeconds(1),
                TokensPerPeriod = _config.MaxGetPerSecondIntegrator,
                AutoReplenishment = true
            });

            _integratorConcurrencyLimiter = new ConcurrencyLimiter(new ConcurrencyLimiterOptions
            {
                PermitLimit = _config.MaxConcurrentGetIntegrator,
                QueueProcessingOrder = QueueProcessingOrder.OldestFirst,
                QueueLimit = int.MaxValue
            });
        }

        public async Task<T> ExecuteAsync<T>(string? facility, bool isGet, Func<Task<T>> op)
        {
            if (!_config.Enabled || !isGet)
            {
                return await op();
            }

            // 1. Acquire Integrator Concurrency
            using var integratorConcurrencyLease = await _integratorConcurrencyLimiter.AcquireAsync(1);
            if (!integratorConcurrencyLease.IsAcquired) throw new InvalidOperationException("Failed to acquire integrator concurrency permit");

            // 2. Acquire Facility Concurrency
            IDisposable? facilityConcurrencyLease = null;
            if (!string.IsNullOrEmpty(facility))
            {
                var fl = GetFacilityConcurrency(facility);
                facilityConcurrencyLease = await fl.AcquireAsync(1);
                if (facilityConcurrencyLease is RateLimitLease lease && !lease.IsAcquired)
                     throw new InvalidOperationException($"Failed to acquire facility concurrency permit for {facility}");
            }

            try
            {
                // 3. Acquire Integrator Rate
                using var integratorRateLease = await _integratorRateLimiter.AcquireAsync(1);
                if (!integratorRateLease.IsAcquired) 
                {
                     // Should wait in queue, but if rejected:
                     await Task.Delay(100); // Simple fallback check
                }

                // 4. Acquire Facility Rate
                if (!string.IsNullOrEmpty(facility))
                {
                    var fl = GetFacilityRate(facility);
                    using var facilityRateLease = await fl.AcquireAsync(1);
                    if (!facilityRateLease.IsAcquired)
                    {
                         // queue logic handled by limiter usually
                    }
                }

                // 5. Retry Loop for 429
                while (true)
                {
                    try
                    {
                        return await op();
                    }
                    catch (Exception ex)
                    {
                        // Check for 429
                        // The generated client currently throws generic exceptions or returns object
                        // We need to check exception message if Client isn't strongly typed error
                        // Assuming Client throws HttpRequestException or similar with status code info
                        if (ex.Message.Contains("429") || (ex is HttpRequestException hre && hre.StatusCode == System.Net.HttpStatusCode.TooManyRequests))
                        {
                            // Backoff
                            await Task.Delay(1000);
                            continue;
                        }
                        throw;
                    }
                }
            }
            finally
            {
                facilityConcurrencyLease?.Dispose();
            }
        }
        
        private TokenBucketRateLimiter GetFacilityRate(string facility)
        {
            return _facilityRateLimiters.GetOrAdd(facility, _ => new TokenBucketRateLimiter(new TokenBucketRateLimiterOptions
            {
                TokenLimit = _config.MaxGetPerSecondPerFacility,
                QueueProcessingOrder = QueueProcessingOrder.OldestFirst,
                QueueLimit = int.MaxValue,
                ReplenishmentPeriod = TimeSpan.FromSeconds(1),
                TokensPerPeriod = _config.MaxGetPerSecondPerFacility,
                AutoReplenishment = true
            }));
        }

        private ConcurrencyLimiter GetFacilityConcurrency(string facility)
        {
             return _facilityConcurrencyLimiters.GetOrAdd(facility, _ => new ConcurrencyLimiter(new ConcurrencyLimiterOptions
            {
                PermitLimit = _config.MaxConcurrentGetPerFacility,
                QueueProcessingOrder = QueueProcessingOrder.OldestFirst,
                QueueLimit = int.MaxValue
            }));
        }
    }
}
`
	os.WriteFile(filepath.Join(dir, "MetrcRateLimiter.cs"), []byte(content), 0644)
}

func generateCsWrapperCode(dir string, groups map[string][]bruno.Request) {
	sb := strings.Builder{}
	sb.WriteString("using System;\nusing System.Collections.Generic;\nusing System.Threading.Tasks;\nusing ThunkMetrc.Client;\n\n")
	sb.WriteString("namespace ThunkMetrc.Wrapper\n{\n")

	sb.WriteString("    public class MetrcWrapper\n    {\n")
	sb.WriteString("        private readonly MetrcClient _client;\n")
	sb.WriteString("        private readonly MetrcRateLimiter _rateLimiter;\n\n")

	sb.WriteString("        public MetrcWrapper(MetrcClient client, RateLimiterConfig? config = null)\n        {\n")
	sb.WriteString("            _client = client;\n")
	sb.WriteString("            _rateLimiter = new MetrcRateLimiter(config);\n")
	sb.WriteString("        }\n\n")

	modelsSb := strings.Builder{}
	modelsSb.WriteString("using System;\nusing System.Collections.Generic;\n\n")
	modelsSb.WriteString("namespace ThunkMetrc.Wrapper\n{\n")

	for group, reqs := range groups {
		if len(reqs) == 0 {
			continue
		}
		safeGroup := ToPascalCase(cleanName(group))

		for _, req := range reqs {
			methodName := ToPascalCase(cleanName(req.Name))
			finalName := safeGroup + methodName

			isVoid := req.Method != "GET"
			returnType := "object"

			pathParams := GetPathParams(req.URL)
			var argsList []string
			for _, p := range pathParams {
				argsList = append(argsList, fmt.Sprintf("string %s", p))
			}

			callBodyArg := "null"

			if req.Method == "POST" || req.Method == "PUT" {
				if req.BodySchema != nil {
					baseTypeName := fmt.Sprintf("%s%sRequest", safeGroup, methodName)
					typeName := baseTypeName

					targetSchema := req.BodySchema
					if req.BodySchema.Type == bruno.TypeArray {
						targetSchema = req.BodySchema.ItemType
						itemTypeName := fmt.Sprintf("%sItem", baseTypeName)
						generateCSharpModel(itemTypeName, targetSchema, &modelsSb)
						typeName = fmt.Sprintf("List<%s>", itemTypeName)
					} else {
						generateCSharpModel(baseTypeName, targetSchema, &modelsSb)
					}

					argsList = append(argsList, fmt.Sprintf("%s body", typeName))
					callBodyArg = "body"
				} else {
					argsList = append(argsList, "object? body = null")
					callBodyArg = "body"
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
				argsList = append(argsList, fmt.Sprintf("string? %s = null", paramName))
			}

			methodParams := strings.Join(argsList, ", ")

			var passArgs []string
			for _, p := range pathParams {
				passArgs = append(passArgs, p)
			}
			if req.Method == "POST" || req.Method == "PUT" {
				passArgs = append(passArgs, callBodyArg)
			}
			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				passArgs = append(passArgs, paramName)
			}
			passArgsStr := strings.Join(passArgs, ", ")

			// Determine Facility
			facilityArg := "null"
			for _, p := range pathParams {
				if strings.ToLower(cleanName(p)) == "licensenumber" {
					facilityArg = p
					break
				}
			}

			isGetStr := "false"
			if req.Method == "GET" {
				isGetStr = "true"
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
				sb.WriteString(fmt.Sprintf("            await _rateLimiter.ExecuteAsync<object?>(%s, %s, async () => { await _client.%s(%s); return null; });\n", facilityArg, isGetStr, finalName, passArgsStr))
				sb.WriteString("        }\n\n")
			} else {
				sb.WriteString(fmt.Sprintf("        public async Task<%s> %s(%s)\n", returnType, finalName, methodParams))
				sb.WriteString("        {\n")
				sb.WriteString(fmt.Sprintf("            return await _rateLimiter.ExecuteAsync<%s>(%s, %s, () => _client.%s(%s));\n", returnType, facilityArg, isGetStr, finalName, passArgsStr))
				sb.WriteString("        }\n\n")
			}
		}
	}

	sb.WriteString("    }\n}\n")
	modelsSb.WriteString("}\n")

	os.WriteFile(filepath.Join(dir, "MetrcWrapper.cs"), []byte(sb.String()), 0644)
	os.WriteFile(filepath.Join(dir, "Models.cs"), []byte(modelsSb.String()), 0644)
}

func generateCSharpModel(name string, schema *bruno.Schema, sb *strings.Builder) {
	if schema.Type == bruno.TypeObject {
		sb.WriteString(fmt.Sprintf("\n    public class %s\n    {\n", name))

		nested := make(map[string]*bruno.Schema)

		var keys []string
		for k := range schema.Properties {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, propName := range keys {
			propSchema := schema.Properties[propName]
			pascalProp := ToPascalCase(cleanName(propName))
			if len(pascalProp) > 0 && pascalProp[0] >= '0' && pascalProp[0] <= '9' {
				pascalProp = "_" + pascalProp
			}

			csType := getCSharpType(propSchema, name, pascalProp)

			if propSchema.Type == bruno.TypeObject {
				subClassName := fmt.Sprintf("%s%s", name, pascalProp)
				csType = subClassName + "?"
				nested[subClassName] = propSchema
			} else if propSchema.Type == bruno.TypeArray && propSchema.ItemType != nil && propSchema.ItemType.Type == bruno.TypeObject {
				subClassName := fmt.Sprintf("%s%s", name, pascalProp) // e.g. RequestIngredient
				csType = fmt.Sprintf("List<%s>?", subClassName)
				nested[subClassName] = propSchema.ItemType
			}

			sb.WriteString(fmt.Sprintf("        public %s %s { get; set; }\n", csType, pascalProp))
		}
		sb.WriteString("    }\n")

		var nestedKeys []string
		for k := range nested {
			nestedKeys = append(nestedKeys, k)
		}
		sort.Strings(nestedKeys)

		for _, subName := range nestedKeys {
			generateCSharpModel(subName, nested[subName], sb)
		}
	}
}

func getCSharpType(s *bruno.Schema, parentName, propName string) string {
	switch s.Type {
	case bruno.TypeString:
		return "string?"
	case bruno.TypeInt:
		if s.IsNullable {
			return "int?"
		}
		return "int"
	case bruno.TypeFloat:
		if s.IsNullable {
			return "double?"
		}
		return "double"
	case bruno.TypeBool:
		if s.IsNullable {
			return "bool?"
		}
		return "bool"
	case bruno.TypeArray:
		inner := "object"
		if s.ItemType != nil {
			inner = getCSharpType(s.ItemType, parentName, propName)
		}
		return fmt.Sprintf("List<%s>?", inner)
	case bruno.TypeObject:
		return "object?"
	default:
		return "object?"
	}
}
