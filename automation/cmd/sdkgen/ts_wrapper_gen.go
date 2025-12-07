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

func generateTypeScriptWrapper(groups map[string][]bruno.Request, version string) {
	log.Println("Generating TypeScript Wrapper...")
	outDir := filepath.Join(outputDir, "typescript", "wrapper")
	os.MkdirAll(outDir, 0755)

	generateTsWrapperPackageJson(outDir, version)
	generateTsWrapperConfig(outDir)
	generateTsRateLimiter(outDir)
	generateTsWrapperCode(outDir, groups)
}

func generateTsWrapperPackageJson(dir string, version string) {
	content := fmt.Sprintf(`{
  "name": "@thunkmetrc/wrapper",
  "version": "%s",
  "description": "Type-safe wrapper for ThunkMetrc TypeScript client",
  "main": "dist/index.js",
  "types": "dist/index.d.ts",
  "files": [
    "dist"
  ],
  "scripts": {
    "build": "tsc",
    "prepublishOnly": "npm run build"
  },
  "dependencies": {
    "@thunkmetrc/client": "file:../client"
  },
  "devDependencies": {
    "typescript": "^5.0.0"
  }
}
`, version)
	os.WriteFile(filepath.Join(dir, "package.json"), []byte(content), 0644)
}

func generateTsWrapperConfig(dir string) {
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

func generateTsRateLimiter(dir string) {
	content := `export class RateLimiterConfig {
  enabled: boolean = false;
  maxGetPerSecondPerFacility: number = 50;
  maxGetPerSecondIntegrator: number = 150;
  maxConcurrentGetPerFacility: number = 10;
  maxConcurrentGetIntegrator: number = 30;
}

class Semaphore {
  private tasks: (() => void)[] = [];
  private count = 0;

  constructor(private max: number) {}

  async acquire(): Promise<void> {
    if (this.count < this.max) {
      this.count++;
      return;
    }
    return new Promise<void>(resolve => this.tasks.push(resolve));
  }

  release(): void {
    if (this.tasks.length > 0) {
      const next = this.tasks.shift()!;
      next();
    } else {
      this.count--;
    }
  }
}

class TokenBucket {
  private tokens: number;
  private lastRefill: number;

  constructor(private rate: number, private capacity: number) {
    this.tokens = capacity;
    this.lastRefill = Date.now();
  }

  async wait(): Promise<void> {
    this.refill();
    if (this.tokens >= 1) {
      this.tokens -= 1;
      return;
    }

    const missing = 1 - this.tokens;
    const waitMs = (missing / this.rate) * 1000;
    await new Promise(resolve => setTimeout(resolve, waitMs));
    return this.wait();
  }

  private refill() {
    const now = Date.now();
    const elapsedSec = (now - this.lastRefill) / 1000;
    this.tokens = Math.min(this.capacity, this.tokens + (elapsedSec * this.rate));
    this.lastRefill = now;
  }
}

export class MetrcRateLimiter {
  private config: RateLimiterConfig;
  
  private integratorRate: TokenBucket;
  private integratorSem: Semaphore;
  
  private facilityRates: Map<string, TokenBucket> = new Map();
  private facilitySems: Map<string, Semaphore> = new Map();

  constructor(config?: RateLimiterConfig) {
    this.config = config || new RateLimiterConfig();
    this.integratorRate = new TokenBucket(this.config.maxGetPerSecondIntegrator, this.config.maxGetPerSecondIntegrator);
    this.integratorSem = new Semaphore(this.config.maxConcurrentGetIntegrator);
  }

  public async execute<T>(facility: string | null, isGet: boolean, op: () => Promise<T>): Promise<T> {
    if (!this.config.enabled || !isGet) {
      return op();
    }

    await this.integratorSem.acquire();
    try {
      if (facility) {
        const sem = this.getFacilitySem(facility);
        await sem.acquire();
        try {
            return await this.executeRateLimited(facility, op);
        } finally {
            sem.release();
        }
      } else {
        return await this.executeRateLimited(facility, op);
      }
    } finally {
      this.integratorSem.release();
    }
  }

  private async executeRateLimited<T>(facility: string | null, op: () => Promise<T>): Promise<T> {
    await this.integratorRate.wait();
    if (facility) {
      await this.getFacilityRate(facility).wait();
    }

    // Retry loop
    while (true) {
      try {
        return await op();
      } catch (err: any) {
        // Check for 429
        if (err?.response?.status === 429 || (err?.message && err.message.includes('429'))) {
           // Backoff
           const retryAfterInfo = err?.response?.headers?.['retry-after'];
           let waitMs = 1000;
           if (retryAfterInfo) {
             const seconds = parseInt(retryAfterInfo, 10);
             if (!isNaN(seconds)) waitMs = seconds * 1000;
           }
           await new Promise(resolve => setTimeout(resolve, waitMs));
           continue;
        }
        throw err;
      }
    }
  }

  private getFacilityRate(facility: string): TokenBucket {
    let tb = this.facilityRates.get(facility);
    if (!tb) {
      tb = new TokenBucket(this.config.maxGetPerSecondPerFacility, this.config.maxGetPerSecondPerFacility);
      this.facilityRates.set(facility, tb);
    }
    return tb;
  }

  private getFacilitySem(facility: string): Semaphore {
    let sem = this.facilitySems.get(facility);
    if (!sem) {
      sem = new Semaphore(this.config.maxConcurrentGetPerFacility);
      this.facilitySems.set(facility, sem);
    }
    return sem;
  }
}
`
	os.WriteFile(filepath.Join(dir, "src", "RateLimiter.ts"), []byte(content), 0644)
}

func generateTsWrapperCode(dir string, groups map[string][]bruno.Request) {
	os.MkdirAll(filepath.Join(dir, "src"), 0755)

	sb := strings.Builder{}
	// Re-export Client so users can use MetrcClient from this package
	sb.WriteString("export * from '@thunkmetrc/client';\n")
	// Import for internal usage (renaming to avoid conflict with export *)
	sb.WriteString("import { MetrcClient as InternalMetrcClient } from '@thunkmetrc/client';\n")
	sb.WriteString("import { MetrcRateLimiter, RateLimiterConfig } from './RateLimiter';\n\n")

	// Model Generation Builder
	var modelsSb strings.Builder
	sb.WriteString(modelsSb.String())
	sb.WriteString("\n")

	sb.WriteString("export class MetrcWrapper {\n")
	sb.WriteString("  private rateLimiter: MetrcRateLimiter;\n\n")
	sb.WriteString("  constructor(public client: InternalMetrcClient, config?: RateLimiterConfig) {\n")
	sb.WriteString("    this.rateLimiter = new MetrcRateLimiter(config);\n")
	sb.WriteString("  }\n\n")

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

			callBodyArg := "undefined"

			if req.Method == "POST" || req.Method == "PUT" {
				if req.BodySchema != nil {
					interfaceName := fmt.Sprintf("%s%sRequest", ToPascalCase(safeGroup), methodName)
					targetSchema := req.BodySchema
					if targetSchema.Type == bruno.TypeArray {
						targetSchema = targetSchema.ItemType
						// ensure item model generated
						itemInterfaceName := fmt.Sprintf("%sItem", interfaceName)
						generateTSInterface(itemInterfaceName, targetSchema, &modelsSb)

						argsList = append(argsList, fmt.Sprintf("body: %s[]", itemInterfaceName))
					} else {
						generateTSInterface(interfaceName, targetSchema, &modelsSb)
						argsList = append(argsList, fmt.Sprintf("body: %s", interfaceName))
					}
					callBodyArg = "body"
				} else {
					argsList = append(argsList, "body?: any")
					callBodyArg = "body"
				}
			} else {
				argsList = append(argsList, "body?: any")
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
				argsList = append(argsList, fmt.Sprintf("%s?: string", paramName))
			}

			methodParams := strings.Join(argsList, ", ")

			// Build call args
			var callArgs []string
			for _, p := range pathParams {
				callArgs = append(callArgs, p)
			}
			callArgs = append(callArgs, callBodyArg)
			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				callArgs = append(callArgs, paramName)
			}

			callArgsStr := strings.Join(callArgs, ", ")

			// Determine Facility
			// Heuristic: Check for parameter named "licenseNumber"
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
			sb.WriteString(fmt.Sprintf("    return this.rateLimiter.execute(%s, %s, () => this.client.%s(%s));\n", facilityArg, isGetStr, tsMethodName, callArgsStr))
			sb.WriteString("  }\n\n")
		}
	}

	sb.WriteString("}\n\n")
	sb.WriteString("// --- Request Models ---\n")
	sb.WriteString(modelsSb.String())

	os.WriteFile(filepath.Join(dir, "src", "index.ts"), []byte(sb.String()), 0644)
}

func generateTSInterface(name string, schema *bruno.Schema, sb *strings.Builder) {
	if schema.Type == bruno.TypeObject {
		sb.WriteString(fmt.Sprintf("export interface %s {\n", name))

		nested := make(map[string]*bruno.Schema)

		var keys []string
		for k := range schema.Properties {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, propName := range keys {
			propSchema := schema.Properties[propName]

			tsType, nestedSchema, nestedName := resolveTSType(propSchema, name, ToPascalCase(cleanName(propName)))
			if nestedSchema != nil {
				nested[nestedName] = nestedSchema
			}

			sb.WriteString(fmt.Sprintf("  '%s'?: %s;\n", propName, tsType))
		}
		sb.WriteString("}\n\n")

		var nestedKeys []string
		for k := range nested {
			nestedKeys = append(nestedKeys, k)
		}
		sort.Strings(nestedKeys)

		for _, subName := range nestedKeys {
			generateTSInterface(subName, nested[subName], sb)
		}
	}
}

func resolveTSType(s *bruno.Schema, parentName, propName string) (string, *bruno.Schema, string) {
	switch s.Type {
	case bruno.TypeString:
		return "string", nil, ""
	case bruno.TypeInt, bruno.TypeFloat:
		return "number", nil, ""
	case bruno.TypeBool:
		return "boolean", nil, ""
	case bruno.TypeArray:
		innerType, innerSchema, innerName := "any", (*bruno.Schema)(nil), ""
		if s.ItemType != nil {
			innerType, innerSchema, innerName = resolveTSType(s.ItemType, parentName, propName)
		}
		return fmt.Sprintf("%s[]", innerType), innerSchema, innerName
	case bruno.TypeObject:
		subInterfaceName := fmt.Sprintf("%s_%s", parentName, propName)
		return subInterfaceName, s, subInterfaceName
	default:
		return "any", nil, ""
	}
}
