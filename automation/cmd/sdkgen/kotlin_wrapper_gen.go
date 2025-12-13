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

func generateKotlinWrapper(groups map[string][]bruno.Request, version string) {
	log.Println("Generating Kotlin Wrapper...")
	outDir := filepath.Join(outputDir, "kotlin", "wrapper")
	os.MkdirAll(outDir, 0755)

	pkgDir := filepath.Join(outDir, "src", "main", "kotlin", "com", "thunkmetrc", "wrapper")
	os.MkdirAll(pkgDir, 0755)

	generateKotlinWrapperPom(outDir, version)
	generateKotlinRateLimiter(pkgDir)
	generateKotlinWrapperCode(pkgDir, groups)
}

func generateKotlinWrapperPom(dir string, version string) {
	content := fmt.Sprintf(`<project xmlns="http://maven.apache.org/POM/4.0.0" 
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>io.github.thunkier</groupId>
    <artifactId>thunkmetrc-kotlin-wrapper</artifactId>
    <version>%s</version>
    <packaging>jar</packaging>

    <name>ThunkMetrc Kotlin Wrapper</name>
    <description>Type-safe wrapper for ThunkMetrc Kotlin client with rate limiting.</description>
    <url>https://github.com/thunkmetrc/sdks</url>

    <licenses>
        <license>
            <name>MIT License</name>
            <url>https://opensource.org/licenses/MIT</url>
        </license>
    </licenses>

    <developers>
        <developer>
            <name>ThunkMetrc Team</name>
            <email>dev@thunkmetrc.com</email>
            <organization>ThunkMetrc</organization>
            <organizationUrl>https://thunkmetrc.com</organizationUrl>
        </developer>
    </developers>

    <scm>
        <connection>scm:git:git://github.com/thunkmetrc/sdks.git</connection>
        <developerConnection>scm:git:ssh://github.com:thunkmetrc/sdks.git</developerConnection>
        <url>https://github.com/thunkmetrc/sdks/tree/main</url>
    </scm>

    <properties>
        <kotlin.version>1.9.0</kotlin.version>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
        <maven.compiler.source>17</maven.compiler.source>
        <maven.compiler.target>17</maven.compiler.target>
    </properties>

    <dependencies>
        <dependency>
            <groupId>org.jetbrains.kotlin</groupId>
            <artifactId>kotlin-stdlib</artifactId>
            <version>${kotlin.version}</version>
        </dependency>
        <dependency>
            <groupId>org.jetbrains.kotlinx</groupId>
            <artifactId>kotlinx-coroutines-core</artifactId>
            <version>1.7.3</version>
        </dependency>
        <dependency>
            <groupId>io.github.thunkier</groupId>
            <artifactId>thunkmetrc-kotlin-client</artifactId>
            <version>%s</version>
        </dependency>
    </dependencies>

    <build>
        <sourceDirectory>${project.basedir}/src/main/kotlin</sourceDirectory>
        <plugins>
            <plugin>
                <groupId>org.jetbrains.kotlin</groupId>
                <artifactId>kotlin-maven-plugin</artifactId>
                <version>${kotlin.version}</version>
                <executions>
                    <execution>
                        <id>compile</id>
                        <phase>compile</phase>
                        <goals>
                            <goal>compile</goal>
                        </goals>
                    </execution>
                </executions>
            </plugin>
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-source-plugin</artifactId>
                <version>3.2.1</version>
                <executions>
                    <execution>
                        <id>attach-sources</id>
                        <goals>
                            <goal>jar-no-fork</goal>
                        </goals>
                    </execution>
                </executions>
            </plugin>
            <plugin>
                <groupId>org.jetbrains.dokka</groupId>
                <artifactId>dokka-maven-plugin</artifactId>
                <version>1.9.0</version>
                <executions>
                    <execution>
                        <phase>prepare-package</phase>
                        <goals>
                            <goal>javadocJar</goal>
                        </goals>
                    </execution>
                </executions>
            </plugin>
            <plugin>
                <groupId>org.sonatype.central</groupId>
                <artifactId>central-publishing-maven-plugin</artifactId>
                <version>0.9.0</version>
                <extensions>true</extensions>
                <configuration>
                    <publishingServerId>central</publishingServerId>
                    <tokenAuth>true</tokenAuth>
                    <autoPublish>true</autoPublish>
                </configuration>
            </plugin>
        </plugins>
    </build>

    <profiles>
        <profile>
            <id>ossrh</id>
            <build>
                <plugins>
                    <plugin>
                        <groupId>org.apache.maven.plugins</groupId>
                        <artifactId>maven-gpg-plugin</artifactId>
                        <version>3.1.0</version>
                        <executions>
                            <execution>
                                <id>sign-artifacts</id>
                                <phase>verify</phase>
                                <goals>
                                    <goal>sign</goal>
                                </goals>
                                <configuration>
                                    <gpgArguments>
                                        <arg>--pinentry-mode</arg>
                                        <arg>loopback</arg>
                                    </gpgArguments>
                                </configuration>
                            </execution>
                        </executions>
                    </plugin>
                </plugins>
            </build>
        </profile>
    </profiles>
</project>
`, version, version)
	os.WriteFile(filepath.Join(dir, "pom.xml"), []byte(content), 0644)
	os.Remove(filepath.Join(dir, "build.gradle.kts"))
}

func generateKotlinWrapperCode(dir string, groups map[string][]bruno.Request) {
	sb := strings.Builder{}
	modelsSb := &strings.Builder{}

	sb.WriteString("package com.thunkmetrc.wrapper\n\n")
	sb.WriteString("import com.thunkmetrc.client.MetrcClient\n\n")

	sb.WriteString("class MetrcWrapper(private val client: MetrcClient, private val config: RateLimiterConfig = RateLimiterConfig()) {\n\n")
	sb.WriteString("    private val rateLimiter = MetrcRateLimiter(config)\n\n")
	sb.WriteString("    constructor(baseUrl: String, vendorKey: String, userKey: String) : this(MetrcClient(baseUrl, vendorKey, userKey))\n\n")

	for group, reqs := range groups {
		safeGroup := ToCamelCase(cleanName(group))
		for _, req := range reqs {
			methodName := ToPascalCase(cleanName(req.Name))
			ktMethodName := safeGroup + methodName

			pathParams := GetPathParams(req.URL)

			var argsList []string
			for _, p := range pathParams {
				argsList = append(argsList, fmt.Sprintf("%s: String", ToCamelCase(p)))
			}

			// Query Params
			var sortedKeys []string
			for k := range req.Params {
				sortedKeys = append(sortedKeys, k)
			}
			sort.Strings(sortedKeys)

			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				ktParamName := ToCamelCase(paramName)
				argsList = append(argsList, fmt.Sprintf("%s: String? = null", ktParamName))
			}

			callBodyArg := "null"
			if req.Method == "POST" || req.Method == "PUT" {
				if req.BodySchema != nil {
					className := fmt.Sprintf("%s%sRequest", ToPascalCase(safeGroup), methodName)
					targetSchema := req.BodySchema
					if targetSchema.Type == bruno.TypeArray {
						targetSchema = targetSchema.ItemType
						className = fmt.Sprintf("%sItem", className)
						generateKotlinModel(className, targetSchema, modelsSb)

						argsList = append(argsList, fmt.Sprintf("body: List<%s>", className))
					} else {
						generateKotlinModel(className, targetSchema, modelsSb)
						argsList = append(argsList, fmt.Sprintf("body: %s", className))
					}
					callBodyArg = "body"
				} else {
					argsList = append(argsList, "body: Any? = null")
					callBodyArg = "body"
				}
			}

			sigArgs := strings.Join(argsList, ", ")

			// Call args
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
			sb.WriteString(fmt.Sprintf("     * %s %s\n     */\n", req.Method, req.Name))
			sb.WriteString(fmt.Sprintf("    suspend fun %s(%s): Any? {\n", ktMethodName, sigArgs))
			sb.WriteString(fmt.Sprintf("        return rateLimiter.execute(%s, %s) { client.%s(%s) }\n", facilityArg, isGet, ktMethodName, callArgsStr))
			sb.WriteString("    }\n\n")
		}
	}

	sb.WriteString(modelsSb.String())
	sb.WriteString("}\n")

	os.WriteFile(filepath.Join(dir, "MetrcWrapper.kt"), []byte(sb.String()), 0644)
}

func generateKotlinRateLimiter(dir string) {
	content := `package com.thunkmetrc.wrapper

import kotlinx.coroutines.sync.Semaphore
import kotlinx.coroutines.sync.Mutex
import kotlinx.coroutines.sync.withLock
import kotlinx.coroutines.delay
import java.util.concurrent.ConcurrentHashMap
import java.time.Instant

data class RateLimiterConfig(
    val enabled: Boolean = false,
    val maxGetPerSecondPerFacility: Double = 50.0,
    val maxGetPerSecondIntegrator: Double = 150.0,
    val maxConcurrentGetPerFacility: Int = 10,
    val maxConcurrentGetIntegrator: Int = 30
)

class MetrcRateLimiter(private val config: RateLimiterConfig = RateLimiterConfig()) {
    private val integratorRate = TokenBucket(config.maxGetPerSecondIntegrator, config.maxGetPerSecondIntegrator)
    private val integratorSem = Semaphore(config.maxConcurrentGetIntegrator)
    
    private val facilityRates = ConcurrentHashMap<String, TokenBucket>()
    private val facilitySems = ConcurrentHashMap<String, Semaphore>()
    private val semMutex = Mutex()

    suspend fun <T> execute(facility: String?, isGet: Boolean, op: suspend () -> T): T {
        if (!config.enabled || !isGet) {
            return op()
        }

        integratorSem.acquire()
        try {
            var facSem: Semaphore? = null
            if (facility != null) {
                // Double-checked locking pattern equivalent for suspension-safe map access if needed
                // But ConcurrentHashMap is thread-safe. ComputeIfAbsent is atomic.
                // However, creating Semaphore is cheap.
                facSem = facilitySems.computeIfAbsent(facility) { Semaphore(config.maxConcurrentGetPerFacility) }
                facSem.acquire()
            }

            try {
                integratorRate.acquire(1.0)
                if (facility != null) {
                    val facRate = facilityRates.computeIfAbsent(facility) { TokenBucket(config.maxGetPerSecondPerFacility, config.maxGetPerSecondPerFacility) }
                    facRate.acquire(1.0)
                }

                while (true) {
                    try {
                        return op()
                    } catch (e: Exception) {
                        if (e.message?.contains("429") == true) {
                            delay(1000)
                            continue
                        }
                        throw e
                    }
                }
            } finally {
                facSem?.release()
            }
        } finally {
            integratorSem.release()
        }
    }

    class TokenBucket(private val rate: Double, private val capacity: Double) {
        private var tokens = capacity
        private var lastRefill = System.nanoTime()
        private val mutex = Mutex()

        suspend fun acquire(amount: Double) {
            while (true) {
                val waitMs = mutex.withLock {
                    refill()
                    if (tokens >= amount) {
                        tokens -= amount
                        return
                    }
                    val missing = amount - tokens
                    (missing / rate * 1000).toLong()
                }
                if (waitMs > 0) delay(waitMs)
            }
        }

        private fun refill() {
            val now = System.nanoTime()
            val elapsed = (now - lastRefill) / 1_000_000_000.0
            tokens = Math.min(capacity, tokens + elapsed * rate)
            lastRefill = now
        }
    }
}
`
	os.WriteFile(filepath.Join(dir, "MetrcRateLimiter.kt"), []byte(content), 0644)
}

func generateKotlinModel(name string, schema *bruno.Schema, sb *strings.Builder) {
	if schema.Type == bruno.TypeObject {
		var keys []string
		for k := range schema.Properties {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		var fields []string

		for _, propName := range keys {
			propSchema := schema.Properties[propName]

			ktType, nestedSchema, nestedName := resolveKotlinType(propSchema, name, ToPascalCase(cleanName(propName)))
			if nestedSchema != nil {
				generateKotlinModel(nestedName, nestedSchema, sb)
			}

			fields = append(fields, fmt.Sprintf("val %s: %s? = null", propName, ktType))
		}

		sb.WriteString(fmt.Sprintf("\n    data class %s(\n", name))
		for i, field := range fields {
			comma := ","
			if i == len(fields)-1 {
				comma = ""
			}
			sb.WriteString(fmt.Sprintf("        %s%s\n", field, comma))
		}
		sb.WriteString("    )\n")
	}
}

func resolveKotlinType(s *bruno.Schema, parentName, propName string) (string, *bruno.Schema, string) {
	switch s.Type {
	case bruno.TypeString:
		return "String", nil, ""
	case bruno.TypeInt:
		return "Int", nil, ""
	case bruno.TypeFloat:
		return "Double", nil, ""
	case bruno.TypeBool:
		return "Boolean", nil, ""
	case bruno.TypeArray:
		innerType, innerSchema, innerName := "Any", (*bruno.Schema)(nil), ""
		if s.ItemType != nil {
			innerType, innerSchema, innerName = resolveKotlinType(s.ItemType, parentName, propName)
		}
		return fmt.Sprintf("List<%s>", innerType), innerSchema, innerName
	case bruno.TypeObject:
		subName := fmt.Sprintf("%s_%s", parentName, propName)
		return subName, s, subName
	default:
		return "Any", nil, ""
	}
}
