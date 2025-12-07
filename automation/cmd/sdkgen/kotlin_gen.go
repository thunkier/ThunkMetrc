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

func generateKotlin(groups map[string][]bruno.Request, version string) {
	log.Println("Generating Kotlin SDK...")

	outDir := filepath.Join(outputDir, "kotlin", "client")
	os.MkdirAll(outDir, 0755)

	pkgDir := filepath.Join(outDir, "src", "main", "kotlin", "com", "thunkmetrc", "client")
	os.MkdirAll(pkgDir, 0755)

	generateKotlinPom(outDir, version)
	generateKotlinClient(pkgDir, groups)
}

func generateKotlinPom(dir string, version string) {
	content := fmt.Sprintf(`<project xmlns="http://maven.apache.org/POM/4.0.0" 
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.thunkmetrc</groupId>
    <artifactId>thunkmetrc-kotlin-client</artifactId>
    <version>%s</version>
    <packaging>jar</packaging>

    <properties>
        <kotlin.version>1.9.0</kotlin.version>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
        <maven.compiler.source>1.8</maven.compiler.source>
        <maven.compiler.target>1.8</maven.compiler.target>
    </properties>

    <dependencies>
        <dependency>
            <groupId>org.jetbrains.kotlin</groupId>
            <artifactId>kotlin-stdlib</artifactId>
            <version>${kotlin.version}</version>
        </dependency>
        <dependency>
            <groupId>com.squareup.okhttp3</groupId>
            <artifactId>okhttp</artifactId>
            <version>4.12.0</version>
        </dependency>
        <dependency>
            <groupId>com.fasterxml.jackson.module</groupId>
            <artifactId>jackson-module-kotlin</artifactId>
            <version>2.15.2</version>
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
        </plugins>
    </build>
</project>
`, version)
	os.WriteFile(filepath.Join(dir, "pom.xml"), []byte(content), 0644)
	// Remove gradle file if exists
	os.Remove(filepath.Join(dir, "build.gradle.kts"))
}

func generateKotlinClient(dir string, groups map[string][]bruno.Request) {
	sb := strings.Builder{}
	modelsSb := &strings.Builder{}
	sb.WriteString("package com.thunkmetrc.client\n\n")
	sb.WriteString("import okhttp3.*\n")
	sb.WriteString("import okhttp3.MediaType.Companion.toMediaType\n")
	sb.WriteString("import okhttp3.RequestBody.Companion.toRequestBody\n")
	sb.WriteString("import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper\n")
	sb.WriteString("import com.fasterxml.jackson.module.kotlin.jacksonObjectMapper\n")
	sb.WriteString("import java.net.URLEncoder\n")
	sb.WriteString("import java.nio.charset.StandardCharsets\n\n")

	sb.WriteString("class MetrcClient(private val baseUrl: String, vendorKey: String, userKey: String) {\n")
	sb.WriteString("    private val client = OkHttpClient()\n")
	sb.WriteString("    private val mapper = jacksonObjectMapper()\n")
	sb.WriteString("    private val auth = Credentials.basic(vendorKey, userKey)\n\n")

	sb.WriteString("    private inline fun <reified T> send(method: String, path: String, body: Any? = null): T? {\n")
	sb.WriteString("        val url = \"${baseUrl.trimEnd('/')}$path\"\n")
	sb.WriteString("        val builder = Request.Builder().url(url).header(\"Authorization\", auth)\n\n")
	sb.WriteString("        if (body != null && (method == \"POST\" || method == \"PUT\")) {\n")
	sb.WriteString("            val json = mapper.writeValueAsString(body)\n")
	sb.WriteString("            builder.method(method, json.toRequestBody(\"application/json\".toMediaType()))\n")
	sb.WriteString("        } else if (method == \"POST\" || method == \"PUT\") {\n")
	sb.WriteString("            builder.method(method, \"\".toRequestBody(null))\n")
	sb.WriteString("        } else {\n")
	sb.WriteString("            builder.method(method, null)\n")
	sb.WriteString("        }\n\n")
	sb.WriteString("        client.newCall(builder.build()).execute().use { response ->\n")
	sb.WriteString("            if (!response.isSuccessful) throw RuntimeException(\"Unexpected code $response\")\n")
	sb.WriteString("            if (response.code == 204) return null\n")
	sb.WriteString("            return mapper.readValue(response.body!!.string(), T::class.java)\n")
	sb.WriteString("        }\n")
	sb.WriteString("    }\n\n")

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

			if req.Method == "POST" || req.Method == "PUT" {
				// Decoupling: Use 'Any?' type for body
				argsList = append(argsList, "body: Any? = null")

			}

			sigArgs := strings.Join(argsList, ", ")

			ktUrl := req.URL
			for _, p := range pathParams {
				ktUrl = strings.ReplaceAll(ktUrl, "{"+p+"}", "${URLEncoder.encode("+ToCamelCase(p)+", StandardCharsets.UTF_8)}")
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
			sb.WriteString(fmt.Sprintf("    fun %s(%s): Any? {\n", ktMethodName, sigArgs))

			sb.WriteString(fmt.Sprintf("        val path = StringBuilder(\"%s\")\n", ktUrl))
			sb.WriteString("        val query = ArrayList<String>()\n")
			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				ktParamName := ToCamelCase(paramName)
				sb.WriteString(fmt.Sprintf("        if (%s != null) {\n", ktParamName))
				sb.WriteString(fmt.Sprintf("            query.add(\"%s=${URLEncoder.encode(%s, StandardCharsets.UTF_8)}\")\n", paramName, ktParamName))
				sb.WriteString("        }\n")
			}
			sb.WriteString("        if (query.isNotEmpty()) {\n")
			sb.WriteString("            path.append(\"?\").append(query.joinToString(\"&\"))\n")
			sb.WriteString("        }\n")

			sendBody := "null"
			if req.Method == "POST" || req.Method == "PUT" {
				sendBody = "body"
			}
			sb.WriteString(fmt.Sprintf("        return send<Any>(\"%s\", path.toString(), %s)\n", req.Method, sendBody))
			sb.WriteString("    }\n\n")
		}
	}

	sb.WriteString(modelsSb.String())
	sb.WriteString("}\n")
	os.WriteFile(filepath.Join(dir, "MetrcClient.kt"), []byte(sb.String()), 0644)
}
