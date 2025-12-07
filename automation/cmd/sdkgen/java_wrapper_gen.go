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

func generateJavaWrapper(groups map[string][]bruno.Request, version string) {
	log.Println("Generating Java Wrapper...")
	dir := filepath.Join(outputDir, "java", "wrapper")
	os.MkdirAll(dir, 0755)

	pkgDir := filepath.Join(dir, "src", "main", "java", "com", "thunkmetrc", "wrapper")
	os.MkdirAll(pkgDir, 0755)

	generateJavaWrapperPom(dir, version)
	generateJavaRateLimiter(pkgDir)
	generateJavaWrapperCode(pkgDir, groups)
}

func generateJavaWrapperPom(dir string, version string) {
	content := fmt.Sprintf(`<project xmlns="http://maven.apache.org/POM/4.0.0" 
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>io.github.thunkier</groupId>
    <artifactId>thunkmetrc-wrapper</artifactId>
    <version>%s</version>
    <packaging>jar</packaging>

    <properties>
        <maven.compiler.source>17</maven.compiler.source>
        <maven.compiler.target>17</maven.compiler.target>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
    </properties>

    <dependencies>
        <dependency>
            <groupId>io.github.thunkier</groupId>
            <artifactId>thunkmetrc-client</artifactId>
            <version>%s</version>
        </dependency>
    </dependencies>
</project>
`, version, version)
	os.WriteFile(filepath.Join(dir, "pom.xml"), []byte(content), 0644)
}

func generateJavaWrapperCode(dir string, groups map[string][]bruno.Request) {
	sb := strings.Builder{}
	sb.WriteString("package com.thunkmetrc.wrapper;\n\n")
	sb.WriteString("import com.thunkmetrc.client.MetrcClient;\n")
	sb.WriteString("import java.io.IOException;\n")
	sb.WriteString("import java.util.List;\n\n")

	sb.WriteString("/**\n * Metrc Wrapper with rate limiting.\n */\n")
	sb.WriteString("public class MetrcWrapper {\n")
	sb.WriteString("    private final MetrcClient client;\n")
	sb.WriteString("    private final RateLimiter rateLimiter;\n\n")

	sb.WriteString("    /**\n     * Constructor.\n     * @param client MetrcClient instance\n     */\n")
	sb.WriteString("    public MetrcWrapper(MetrcClient client) {\n")
	sb.WriteString("        this(client, new RateLimiter.Config());\n")
	sb.WriteString("    }\n\n")

	sb.WriteString("    /**\n     * Constructor.\n     * @param client MetrcClient instance\n     * @param config RateLimiter configuration\n     */\n")
	sb.WriteString("    public MetrcWrapper(MetrcClient client, RateLimiter.Config config) {\n")
	sb.WriteString("        this.client = client;\n")
	sb.WriteString("        this.rateLimiter = new RateLimiter(config);\n")
	sb.WriteString("    }\n\n")

	for group, reqs := range groups {
		safeGroup := ToCamelCase(cleanName(group))
		for _, req := range reqs {
			methodName := ToPascalCase(cleanName(req.Name))
			javaMethodName := safeGroup + methodName

			pathParams := GetPathParams(req.URL)

			var argsList []string
			for _, p := range pathParams {
				argsList = append(argsList, fmt.Sprintf("String %s", ToCamelCase(p)))
			}

			// Query Params
			var sortedKeys []string
			for k := range req.Params {
				sortedKeys = append(sortedKeys, k)
			}
			sort.Strings(sortedKeys)

			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				javaParamName := ToCamelCase(paramName)
				argsList = append(argsList, fmt.Sprintf("String %s", javaParamName))
			}

			callBodyArg := "null"
			if req.Method == "POST" || req.Method == "PUT" {
				if req.BodySchema != nil {
					className := fmt.Sprintf("%s%sRequest", ToPascalCase(safeGroup), methodName)
					targetSchema := req.BodySchema
					if targetSchema.Type == bruno.TypeArray {
						targetSchema = targetSchema.ItemType
						className = fmt.Sprintf("%sItem", className)
						className = fmt.Sprintf("%sItem", fmt.Sprintf("%s%sRequest", ToPascalCase(safeGroup), methodName))
						generateJavaModel(className, targetSchema, &sb)

						argsList = append(argsList, fmt.Sprintf("List<%s> body", className))
					} else {
						generateJavaModel(className, targetSchema, &sb)
						argsList = append(argsList, fmt.Sprintf("%s body", className))
					}
					callBodyArg = "body"
				} else {
					argsList = append(argsList, "Object body")
					callBodyArg = "body"
				}
			}

			sigArgs := strings.Join(argsList, ", ")

			// Build call args
			var callArgs []string
			for _, p := range pathParams {
				callArgs = append(callArgs, ToCamelCase(p))
			}
			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				callArgs = append(callArgs, ToCamelCase(paramName))
			}
			if req.Method == "POST" || req.Method == "PUT" {
				callArgs = append(callArgs, callBodyArg)
			}
			callArgsStr := strings.Join(callArgs, ", ")

			// Determine Facility
			facilityArg := "null"
			for _, p := range pathParams {
				if strings.ToLower(cleanName(p)) == "licensenumber" {
					facilityArg = ToCamelCase(p)
					break
				}
			}
			isGet := "false"
			if req.Method == "GET" {
				isGet = "true"
			}

			sb.WriteString(fmt.Sprintf("    /**\n"))
			docs := cleanDocs(req.Docs)
			if docs != "" {
				for _, line := range strings.Split(docs, "\n") {
					sb.WriteString(fmt.Sprintf("     * %s\n", line))
				}
				sb.WriteString("     *\n")
			}
			sb.WriteString(fmt.Sprintf("     * %s %s\n     * @throws Exception If request fails\n     * @return Response object\n     */\n", req.Method, req.Name))
			sb.WriteString(fmt.Sprintf("    public Object %s(%s) throws Exception {\n", javaMethodName, sigArgs))
			sb.WriteString(fmt.Sprintf("        return rateLimiter.execute(%s, %s, () -> client.%s(%s));\n", facilityArg, isGet, javaMethodName, callArgsStr))
			sb.WriteString("    }\n\n")
		}
	}

	sb.WriteString("}\n")
	os.WriteFile(filepath.Join(dir, "MetrcWrapper.java"), []byte(sb.String()), 0644)
}

func generateJavaRateLimiter(dir string) {
	content := `package com.thunkmetrc.wrapper;

import java.util.concurrent.Semaphore;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.TimeUnit;
import java.util.Map;
import java.util.concurrent.Callable;

public class RateLimiter {
    public static class Config {
        public boolean enabled = false;
        public double maxGetPerSecondPerFacility = 50.0;
        public double maxGetPerSecondIntegrator = 150.0;
        public int maxConcurrentGetPerFacility = 10;
        public int maxConcurrentGetIntegrator = 30;
    }

    private final Config config;
    private final Semaphore integratorSem;
    private final TokenBucket integratorRate;
    private final Map<String, Semaphore> facilitySems = new ConcurrentHashMap<>();
    private final Map<String, TokenBucket> facilityRates = new ConcurrentHashMap<>();

    public RateLimiter(Config config) {
        if (config == null) config = new Config();
        this.config = config;
        this.integratorSem = new Semaphore(config.maxConcurrentGetIntegrator);
        this.integratorRate = new TokenBucket(config.maxGetPerSecondIntegrator, config.maxGetPerSecondIntegrator);
    }

    public <T> T execute(String facility, boolean isGet, Callable<T> op) throws Exception {
        if (!config.enabled || !isGet) {
            return op.call();
        }

        integratorSem.acquire();
        try {
            Semaphore facSem = null;
            if (facility != null) {
                facSem = facilitySems.computeIfAbsent(facility, k -> new Semaphore(config.maxConcurrentGetPerFacility));
                facSem.acquire();
            }

            try {
                integratorRate.acquire(1);
                if (facility != null) {
                    TokenBucket facRate = facilityRates.computeIfAbsent(facility, k -> new TokenBucket(config.maxGetPerSecondPerFacility, config.maxGetPerSecondPerFacility));
                    facRate.acquire(1);
                }

                while (true) {
                    try {
                        return op.call();
                    } catch (Exception e) {
                        if (e.getMessage() != null && e.getMessage().contains("429")) {
                            Thread.sleep(1000);
                            continue;
                        }
                        throw e;
                    }
                }
            } finally {
                if (facSem != null) facSem.release();
            }
        } finally {
            integratorSem.release();
        }
    }

    static class TokenBucket {
        private final double rate;
        private final double capacity;
        private double tokens;
        private long lastRefill;

        public TokenBucket(double rate, double capacity) {
            this.rate = rate;
            this.capacity = capacity;
            this.tokens = capacity;
            this.lastRefill = System.nanoTime();
        }

        public synchronized void acquire(double amount) throws InterruptedException {
            while (true) {
                refill();
                if (tokens >= amount) {
                    tokens -= amount;
                    return;
                }
                double missing = amount - tokens;
                long waitNs = (long) (missing / rate * 1_000_000_000.0);
                Thread.sleep(waitNs / 1_000_000, (int) (waitNs % 1_000_000));
            }
        }

        private void refill() {
            long now = System.nanoTime();
            double elapsed = (now - lastRefill) / 1_000_000_000.0;
            tokens = Math.min(capacity, tokens + elapsed * rate);
            lastRefill = now;
        }
    }
}
`
	os.WriteFile(filepath.Join(dir, "RateLimiter.java"), []byte(content), 0644)
}

func generateJavaModel(name string, schema *bruno.Schema, sb *strings.Builder) {
	if schema.Type == bruno.TypeObject {
		sb.WriteString(fmt.Sprintf("    public static class %s {\n", name))

		nested := make(map[string]*bruno.Schema)

		var keys []string
		for k := range schema.Properties {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, propName := range keys {
			propSchema := schema.Properties[propName]
			javaField := propName
			javaType, nestedSchema, nestedName := resolveJavaType(propSchema, name, ToPascalCase(cleanName(propName)))
			if nestedSchema != nil {
				nested[nestedName] = nestedSchema
			}

			sb.WriteString(fmt.Sprintf("        public %s %s;\n", javaType, javaField))
		}
		sb.WriteString("    }\n\n")

		var nestedKeys []string
		for k := range nested {
			nestedKeys = append(nestedKeys, k)
		}
		sort.Strings(nestedKeys)

		for _, subName := range nestedKeys {
			generateJavaModel(subName, nested[subName], sb)
		}
	}
}

func resolveJavaType(s *bruno.Schema, parentName, propName string) (string, *bruno.Schema, string) {
	switch s.Type {
	case bruno.TypeString:
		return "String", nil, ""
	case bruno.TypeInt:
		return "Integer", nil, ""
	case bruno.TypeFloat:
		return "Double", nil, ""
	case bruno.TypeBool:
		return "Boolean", nil, ""
	case bruno.TypeArray:
		innerType, innerSchema, innerName := "Object", (*bruno.Schema)(nil), ""
		if s.ItemType != nil {
			innerType, innerSchema, innerName = resolveJavaType(s.ItemType, parentName, propName)
		}
		return fmt.Sprintf("List<%s>", innerType), innerSchema, innerName
	case bruno.TypeObject:
		subStructName := fmt.Sprintf("%s_%s", parentName, propName)
		return subStructName, s, subStructName
	default:
		return "Object", nil, ""
	}
}
