package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: bump <file> <part> (major|minor|patch)")
	}
	file, part := os.Args[1], os.Args[2]

	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read: %v", err)
	}

	v := strings.TrimSpace(string(content))
	parts := strings.Split(v, ".")
	if len(parts) != 3 {
		log.Fatalf("Invalid version: %s", v)
	}

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	switch part {
	case "major":
		major++
		minor = 0
		patch = 0
	case "minor":
		minor++
		patch = 0
	case "patch":
		patch++
	default:
		log.Fatal("Unknown part")
	}

	newVersion := fmt.Sprintf("%d.%d.%d", major, minor, patch)
	if err := os.WriteFile(file, []byte(newVersion), 0644); err != nil {
		log.Fatalf("Failed to write: %v", err)
	}
	fmt.Println(newVersion)

	updateManifests(newVersion)
}

func updateManifests(version string) {
	updateCSharp(version)
	updateTypeScript(version)
	updatePython(version)
	updateRust(version)
	updateJavaKotlin(version, "../sdks/java/thunkmetrc/pom.xml", "thunkmetrc", "thunkmetrc-wrapper")
	updateJavaKotlin(version, "../sdks/kotlin/thunkmetrc/pom.xml", "thunkmetrc", "thunkmetrc-kotlin-wrapper")
}

func updateCSharp(version string) {
	replaceInFile("../sdks/csharp/ThunkMetrc/ThunkMetrc.csproj",
		regexp.MustCompile(`<Version>.*?</Version>`),
		fmt.Sprintf("<Version>%s</Version>", version))
}

func updateTypeScript(version string) {
	replaceInFile("../sdks/typescript/thunkmetrc/package.json",
		regexp.MustCompile(`"version": ".*?"`),
		fmt.Sprintf(`"version": "%s"`, version))
}

func updatePython(version string) {
	path := "../sdks/python/thunkmetrc/pyproject.toml"
	replaceInFile(path, regexp.MustCompile(`version = ".*?"`), fmt.Sprintf(`version = "%s"`, version))
	replaceInFile(path, regexp.MustCompile(`thunkmetrc-wrapper==.*?(\"|')`), fmt.Sprintf(`thunkmetrc-wrapper==%s"`, version))
}

func updateRust(version string) {
	path := "../sdks/rust/thunkmetrc/Cargo.toml"
	content := readFile(path)

	// Package and dependency version updates
	content = regexp.MustCompile(`(?m)^version = ".*"`).ReplaceAllString(content, fmt.Sprintf(`version = "%s"`, version))
	content = regexp.MustCompile(`thunkmetrc-wrapper = \{ path = "\.\./wrapper", version = ".*?" \}`).
		ReplaceAllString(content, fmt.Sprintf(`thunkmetrc-wrapper = { path = "../wrapper", version = "%s" }`, version))

	writeFile(path, content)
}

func updateJavaKotlin(version, path, artifactId, depArtifactId string) {
	content := readFile(path)

	// Helper to replace matching XML version tag
	replaceVer := func(artId string, oldContent string) string {
		re := regexp.MustCompile(fmt.Sprintf(`(<artifactId>%s</artifactId>\s*<version>)(.*?)(</version>)`, regexp.QuoteMeta(artId)))
		return re.ReplaceAllString(oldContent, "${1}"+version+"${3}")
	}

	content = replaceVer(artifactId, content)    // Project version
	content = replaceVer(depArtifactId, content) // Dependency version
	writeFile(path, content)
}

func replaceInFile(path string, re *regexp.Regexp, replace string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Skip %s: %v", path, err)
		return
	}
	s := string(content)
	newS := re.ReplaceAllString(s, replace)
	if s != newS {
		if err := os.WriteFile(path, []byte(newS), 0644); err != nil {
			log.Printf("Err writing %s: %v", path, err)
		}
	}
}

func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Skip %s: %v", path, err)
		return ""
	}
	return string(content)
}

func writeFile(path string, content string) {
	if content == "" {
		return
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		log.Printf("Err writing %s: %v", path, err)
	}
}
