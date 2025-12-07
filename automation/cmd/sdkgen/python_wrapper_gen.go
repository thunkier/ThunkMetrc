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

func generatePythonWrapper(groups map[string][]bruno.Request, version string) {
	log.Println("Generating Python Wrapper...")
	outDir := filepath.Join(outputDir, "python", "wrapper")
	os.MkdirAll(outDir, 0755)

	srcDir := filepath.Join(outDir, "src", "thunkmetrc", "wrapper")
	os.MkdirAll(srcDir, 0755)

	generatePyWrapperProject(outDir, version)
	generatePyWrapperInit(outDir)
	generatePyRateLimiter(srcDir)
	generatePyWrapperCode(srcDir, groups)
}

func generatePyRateLimiter(dir string) {
	content := `import asyncio
import time
from typing import Optional, Dict

class RateLimiterConfig:
    def __init__(self, 
                 enabled: bool = False,
                 max_get_per_second_per_facility: float = 50,
                 max_get_per_second_integrator: float = 150,
                 max_concurrent_get_per_facility: int = 10,
                 max_concurrent_get_integrator: int = 30):
        self.enabled = enabled
        self.max_get_per_second_per_facility = max_get_per_second_per_facility
        self.max_get_per_second_integrator = max_get_per_second_integrator
        self.max_concurrent_get_per_facility = max_concurrent_get_per_facility
        self.max_concurrent_get_integrator = max_concurrent_get_integrator

class TokenBucket:
    def __init__(self, rate: float, capacity: float):
        self.rate = rate
        self.capacity = capacity
        self.tokens = capacity
        self.last_refill = time.monotonic()
        self.lock = asyncio.Lock()

    async def wait(self):
        async with self.lock:
            now = time.monotonic()
            elapsed = now - self.last_refill
            self.tokens = min(self.capacity, self.tokens + elapsed * self.rate)
            self.last_refill = now

            if self.tokens >= 1:
                self.tokens -= 1
                return
            
            # Wait time
            missing = 1.0 - self.tokens
            wait_time = missing / self.rate
        
        await asyncio.sleep(wait_time)
        await self.wait() # Recurse to re-check

class MetrcRateLimiter:
    def __init__(self, config: Optional[RateLimiterConfig] = None):
        if config is None:
            config = RateLimiterConfig()
        self.config = config
        
        self.integrator_rate = TokenBucket(config.max_get_per_second_integrator, config.max_get_per_second_integrator)
        self.integrator_sem = asyncio.Semaphore(config.max_concurrent_get_integrator)
        
        self.facility_rates: Dict[str, TokenBucket] = {}
        self.facility_sems: Dict[str, asyncio.Semaphore] = {}
        self.lock = asyncio.Lock()

    async def execute(self, facility: Optional[str], is_get: bool, op):
        if not self.config.enabled or not is_get:
            return await op()

        async with self.integrator_sem:
            async with self.get_facility_sem(facility):
                await self.integrator_rate.wait()
                if facility:
                    await self.get_facility_rate(facility).wait()
                
                # Retry loop
                while True:
                    try:
                        return await op()
                    except Exception as e:
                        if "429" in str(e): # Checking generic error string, assume client raises exceptions
                             # TODO: Parse retry after
                             await asyncio.sleep(1.0)
                             continue
                        raise e

    def get_facility_rate(self, facility: str) -> TokenBucket:
        if facility not in self.facility_rates:
             self.facility_rates[facility] = TokenBucket(self.config.max_get_per_second_per_facility, self.config.max_get_per_second_per_facility)
        return self.facility_rates[facility]

    def get_facility_sem(self, facility: Optional[str]) -> asyncio.Semaphore:
        if not facility:
             # Return dummy semaphore or context manager
             return asyncio.Semaphore(9999) 
        if facility not in self.facility_sems:
             self.facility_sems[facility] = asyncio.Semaphore(self.config.max_concurrent_get_per_facility)
        return self.facility_sems[facility]
`
	os.WriteFile(filepath.Join(dir, "ratelimiter.py"), []byte(content), 0644)
}

func generatePyWrapperCode(dir string, groups map[string][]bruno.Request) {
	sb := strings.Builder{}
	sb.WriteString("from typing import Any, Optional, List, Callable, Awaitable\n")
	sb.WriteString("from typing_extensions import TypedDict\n")
	sb.WriteString("from thunkmetrc.client import MetrcClient\n")
	sb.WriteString("from .ratelimiter import MetrcRateLimiter, RateLimiterConfig\n\n")

	// Models builder
	var modelBuilder strings.Builder

	sb.WriteString("class MetrcWrapper:\n")
	sb.WriteString("    def __init__(self, client: MetrcClient, config: Optional[RateLimiterConfig] = None):\n")
	sb.WriteString("        self.client = client\n")
	sb.WriteString("        self.limiter = MetrcRateLimiter(config)\n\n")

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

			callBodyArg := "None"
			if req.Method == "POST" || req.Method == "PUT" {
				if req.BodySchema != nil {
					typeName := fmt.Sprintf("%s%sRequest", ToPascalCase(safeGroup), ToPascalCase(cleanName(req.Name)))
					targetSchema := req.BodySchema
					if targetSchema.Type == bruno.TypeArray {
						targetSchema = targetSchema.ItemType
						// Ensure item model generated
						itemTypeName := fmt.Sprintf("%sItem", typeName)
						generatePyModel(itemTypeName, targetSchema, &modelBuilder)

						typeName = fmt.Sprintf("List[%s]", itemTypeName)
					} else {
						generatePyModel(typeName, targetSchema, &modelBuilder)
					}
					// Type hint
					argsList = append(argsList, fmt.Sprintf("body: Optional['%s'] = None", typeName))
					callBodyArg = "body"
				} else {
					argsList = append(argsList, "body: Any = None")
					callBodyArg = "body"
				}
			} else {
				argsList = append(argsList, "body: Any = None")
				callBodyArg = "body"
			}

			// Query Params
			var sortedKeys []string
			for k := range req.Params {
				sortedKeys = append(sortedKeys, k)
			}
			sort.Strings(sortedKeys)

			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				snake := ToSnakeCase(paramName)
				argsList = append(argsList, fmt.Sprintf("%s: Optional[str] = None", snake))
			}

			sigArgs := strings.Join(argsList, ", ")
			if len(sigArgs) > 0 {
				sigArgs = "self, " + sigArgs
			} else {
				sigArgs = "self"
			}

			// Build call args
			var callArgs []string
			for _, p := range pathParams {
				callArgs = append(callArgs, ToSnakeCase(p))
			}
			callArgs = append(callArgs, "body="+callBodyArg)
			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				snake := ToSnakeCase(paramName)
				callArgs = append(callArgs, fmt.Sprintf("%s=%s", snake, snake))
			}
			callArgsStr := strings.Join(callArgs, ", ")

			// Determine Facility
			facilityArg := "None"
			for _, p := range pathParams {
				if strings.ToLower(cleanName(p)) == "licensenumber" {
					facilityArg = ToSnakeCase(p)
					break
				}
			}

			isGet := "False"
			if req.Method == "GET" {
				isGet = "True"
			}

			sb.WriteString(fmt.Sprintf("    async def %s(%s) -> Any:\n", pyMethodName, sigArgs))
			sb.WriteString(fmt.Sprintf("        \"\"\"\n"))
			docs := cleanDocs(req.Docs)
			if docs != "" {
				singleLineDocs := strings.ReplaceAll(docs, "\n", "\n        ")
				sb.WriteString(fmt.Sprintf("        %s\n\n", singleLineDocs))
			}
			sb.WriteString(fmt.Sprintf("        %s %s\n        \"\"\"\n", req.Method, req.Name))
			sb.WriteString(fmt.Sprintf("        async def op():\n"))
			sb.WriteString(fmt.Sprintf("            return await self.client.%s(%s)\n", pyMethodName, callArgsStr))
			sb.WriteString(fmt.Sprintf("        return await self.limiter.execute(%s, %s, op)\n\n", facilityArg, isGet))
		}
	}

	sb.WriteString("# --- Models ---\n")
	sb.WriteString(modelBuilder.String())

	os.WriteFile(filepath.Join(dir, "wrapper.py"), []byte(sb.String()), 0644)
}
func generatePyWrapperProject(dir string, version string) {
	content := fmt.Sprintf(`[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[project]
name = "thunkmetrc-wrapper"
version = "%s"
description = "Type-safe wrapper for ThunkMetrc Python client"
readme = "README.md"
requires-python = ">=3.10"
license = "MIT"
authors = [
  { name = "ThunkMetrc", email = "dev@thunkmetrc.com" },
]
dependencies = [
    "thunkmetrc-client==%s",
]

[tool.hatch.build.targets.wheel]
packages = ["src/thunkmetrc"]
`, version, version)
	os.WriteFile(filepath.Join(dir, "pyproject.toml"), []byte(content), 0644)
	os.WriteFile(filepath.Join(dir, "README.md"), []byte("# ThunkMetrc Wrapper\n\nType-safe wrapper for ThunkMetrc Python Client."), 0644)
}

func generatePyWrapperInit(dir string) {
	content := `from .wrapper import MetrcWrapper
from thunkmetrc.client import MetrcClient

__all__ = ["MetrcWrapper", "MetrcClient"]
`
	os.WriteFile(filepath.Join(dir, "src", "thunkmetrc", "wrapper", "__init__.py"), []byte(content), 0644)
}

var generatedPyModels = make(map[string]bool)

func generatePyModel(name string, schema *bruno.Schema, sb *strings.Builder) {
	if generatedPyModels[name] {
		return
	}
	generatedPyModels[name] = true

	if schema.Type == bruno.TypeObject {
		var keys []string
		for k := range schema.Properties {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var fieldsSb strings.Builder

		for _, propName := range keys {
			propSchema := schema.Properties[propName]
			pyType, nestedSchema, nestedName := resolvePyType(propSchema, name, ToPascalCase(cleanName(propName)))

			if nestedSchema != nil {
				generatePyModel(nestedName, nestedSchema, sb)
			}

			fieldsSb.WriteString(fmt.Sprintf("    %s: %s\n", propName, pyType))
		}

		sb.WriteString(fmt.Sprintf("\nclass %s(TypedDict, total=False):\n", name))
		if fieldsSb.Len() == 0 {
			sb.WriteString("    pass\n")
		} else {
			sb.WriteString(fieldsSb.String())
		}
	}
}

func resolvePyType(s *bruno.Schema, parentName, propName string) (string, *bruno.Schema, string) {
	switch s.Type {
	case bruno.TypeString:
		return "str", nil, ""
	case bruno.TypeInt:
		return "int", nil, ""
	case bruno.TypeFloat:
		return "float", nil, ""
	case bruno.TypeBool:
		return "bool", nil, ""
	case bruno.TypeArray:
		innerType, innerSchema, innerName := "Any", (*bruno.Schema)(nil), ""
		if s.ItemType != nil {
			innerType, innerSchema, innerName = resolvePyType(s.ItemType, parentName, propName)
		}
		return fmt.Sprintf("List[%s]", innerType), innerSchema, innerName
	case bruno.TypeObject:
		subName := fmt.Sprintf("%s_%s", parentName, propName)
		return subName, s, subName
	default:
		return "Any", nil, ""
	}
}
