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

func generateJava(groups map[string][]bruno.Request, version string) {
	log.Println("Generating Java SDK...")

	dir := filepath.Join(outputDir, "java", "client")
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Failed to create Java dir: %v", err)
	}
	pkgDir := filepath.Join(dir, "src", "main", "java", "com", "thunkmetrc", "client")
	os.MkdirAll(pkgDir, 0755)

	generatePomXml(dir, version)
	generateJavaClient(pkgDir, groups)
}

func generatePomXml(dir string, version string) {
	content := fmt.Sprintf(`<project xmlns="http://maven.apache.org/POM/4.0.0" 
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>io.github.thunkier</groupId>
    <artifactId>thunkmetrc-client</artifactId>
    <version>%s</version>
    <packaging>jar</packaging>

    <name>ThunkMetrc Client</name>
    <description>Auto-generated Java client for ThunkMetrc API.</description>
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
        <maven.compiler.source>17</maven.compiler.source>
        <maven.compiler.target>17</maven.compiler.target>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
    </properties>

    <dependencies>
        <dependency>
            <groupId>com.squareup.okhttp3</groupId>
            <artifactId>okhttp</artifactId>
            <version>4.12.0</version>
        </dependency>
        <dependency>
            <groupId>com.fasterxml.jackson.core</groupId>
            <artifactId>jackson-databind</artifactId>
            <version>2.15.2</version>
        </dependency>
        <dependency>
            <groupId>org.junit.jupiter</groupId>
            <artifactId>junit-jupiter</artifactId>
            <version>5.10.0</version>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>io.github.cdimascio</groupId>
            <artifactId>java-dotenv</artifactId>
            <version>5.2.2</version>
            <scope>test</scope>
        </dependency>
    </dependencies>

    <build>
        <plugins>
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
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-javadoc-plugin</artifactId>
                <version>3.5.0</version>
                <executions>
                    <execution>
                        <id>attach-javadocs</id>
                        <goals>
                            <goal>jar</goal>
                        </goals>
                    </execution>
                </executions>
            </plugin>
            <plugin>
                <groupId>org.sonatype.central</groupId>
                <artifactId>central-publishing-maven-plugin</artifactId>
                <version>0.6.0</version>
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
`, version)
	os.WriteFile(filepath.Join(dir, "pom.xml"), []byte(content), 0644)
}

func generateJavaClient(dir string, groups map[string][]bruno.Request) {
	sb := strings.Builder{}
	sb.WriteString("package com.thunkmetrc.client;\n\n")
	sb.WriteString("import okhttp3.*;\n")
	sb.WriteString("import com.fasterxml.jackson.databind.ObjectMapper;\n")
	sb.WriteString("import java.io.IOException;\n")
	sb.WriteString("import java.util.Map;\n")
	sb.WriteString("import java.util.List;\n")
	sb.WriteString("import java.net.URLEncoder;\n")
	sb.WriteString("import java.nio.charset.StandardCharsets;\n\n")

	sb.WriteString("/**\n * Metrc Client for interacting with the Metrc API.\n */\n")
	sb.WriteString("public class MetrcClient {\n")
	sb.WriteString("    private final OkHttpClient client;\n")
	sb.WriteString("    private final ObjectMapper mapper;\n")
	sb.WriteString("    private final String baseUrl;\n")
	sb.WriteString("    private final String apiKey;\n\n")

	sb.WriteString("    /**\n     * Constructor.\n     * @param baseUrl Base URL for the API\n     * @param vendorKey Vendor API Key\n     * @param userKey User API Key\n     */\n")
	sb.WriteString("    public MetrcClient(String baseUrl, String vendorKey, String userKey) {\n")
	sb.WriteString("        this.baseUrl = baseUrl.replaceAll(\"/$\", \"\");\n")
	sb.WriteString("        this.apiKey = Credentials.basic(vendorKey, userKey);\n")
	sb.WriteString("        this.client = new OkHttpClient();\n")
	sb.WriteString("        this.mapper = new ObjectMapper();\n")
	sb.WriteString("    }\n\n")

	sb.WriteString("    private Object send(String method, String path, Object body) throws IOException {\n")
	sb.WriteString("        String url = this.baseUrl + path;\n")
	sb.WriteString("        Request.Builder builder = new Request.Builder()\n")
	sb.WriteString("            .url(url)\n")
	sb.WriteString("            .header(\"Authorization\", this.apiKey);\n\n")

	sb.WriteString("        if (body != null && (method.equals(\"POST\") || method.equals(\"PUT\"))) {\n")
	sb.WriteString("            String json = this.mapper.writeValueAsString(body);\n")
	sb.WriteString("            builder.method(method, RequestBody.create(json, MediaType.parse(\"application/json\")));\n")
	sb.WriteString("        } else if (method.equals(\"POST\") || method.equals(\"PUT\")) {\n")
	sb.WriteString("            builder.method(method, RequestBody.create(\"\", null));\n")
	sb.WriteString("        } else {\n")
	sb.WriteString("            builder.method(method, null);\n")
	sb.WriteString("        }\n\n")

	sb.WriteString("        try (Response response = client.newCall(builder.build()).execute()) {\n")
	sb.WriteString("            if (!response.isSuccessful()) throw new IOException(\"Unexpected code \" + response);\n")
	sb.WriteString("            if (response.code() == 204) return null;\n")
	sb.WriteString("            return this.mapper.readValue(response.body().string(), Object.class);\n")
	sb.WriteString("        }\n")
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

			if req.Method == "POST" || req.Method == "PUT" {
				// Decoupling: Use 'Object' type for body
				argsList = append(argsList, "Object body")

			}

			sigArgs := strings.Join(argsList, ", ")

			javaUrl := req.URL
			for _, p := range pathParams {
				javaUrl = strings.ReplaceAll(javaUrl, "{"+p+"}", "\"+URLEncoder.encode("+ToCamelCase(p)+", StandardCharsets.UTF_8)+\"")
			}

			sb.WriteString(fmt.Sprintf("    /**\n"))
			docs := cleanDocs(req.Docs)
			if docs != "" {
				for _, line := range strings.Split(docs, "\n") {
					sb.WriteString(fmt.Sprintf("     * %s\n", line))
				}
				sb.WriteString("     *\n")
			}
			sb.WriteString(fmt.Sprintf("     * %s %s\n     * @throws IOException If request fails\n     * @return Response object\n     */\n", req.Method, req.Name))
			sb.WriteString(fmt.Sprintf("    public Object %s(%s) throws IOException {\n", javaMethodName, sigArgs))

			sb.WriteString(fmt.Sprintf("        StringBuilder path = new StringBuilder(\"%s\");\n", javaUrl))
			sb.WriteString("        StringBuilder query = new StringBuilder();\n")
			for _, k := range sortedKeys {
				paramName := strings.TrimSuffix(k, "optional")
				javaParamName := ToCamelCase(paramName)
				sb.WriteString(fmt.Sprintf("        if (%s != null) {\n", javaParamName))
				sb.WriteString(fmt.Sprintf("            if (query.length() > 0) query.append(\"&\");\n"))
				sb.WriteString(fmt.Sprintf("            query.append(\"%s=\").append(URLEncoder.encode(%s, StandardCharsets.UTF_8));\n", paramName, javaParamName))
				sb.WriteString("        }\n")
			}
			sb.WriteString("        if (query.length() > 0) path.append(\"?\").append(query);\n")

			sendBody := "null"
			if req.Method == "POST" || req.Method == "PUT" {
				sendBody = "body"
			}
			sb.WriteString(fmt.Sprintf("        return send(\"%s\", path.toString(), %s);\n", req.Method, sendBody))
			sb.WriteString("    }\n\n")
		}
	}

	sb.WriteString("}\n")
	os.WriteFile(filepath.Join(dir, "MetrcClient.java"), []byte(sb.String()), 0644)
}
