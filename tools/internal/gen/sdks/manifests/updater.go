package manifests

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func UpdateManifests(sdksDir, version string) {
	log.Printf("Updating handwritten manifests in %s to version %s...", sdksDir, version)
	updateCSharp(sdksDir, version)
	updateTypeScript(sdksDir, version)
	updatePython(sdksDir, version)
	updateRust(sdksDir, version)
	updateJava(sdksDir, version)
	updateKotlin(sdksDir, version)
	updateGo(sdksDir, version)
	updateGoWorkspace(sdksDir, version)
}

func updateCSharp(sdksDir, version string) {
	files := []string{
		filepath.Join(sdksDir, "thunkmetrc-cs", "ThunkMetrc", "ThunkMetrc.csproj"),
	}
	for _, f := range files {
		replaceInFile(f, regexp.MustCompile(`<Version>.*?</Version>`), fmt.Sprintf("<Version>%s</Version>", version))
	}
}

func updateTypeScript(sdksDir, version string) {
	files := []string{
		filepath.Join(sdksDir, "thunkmetrc-ts", "thunkmetrc", "package.json"),
	}
	for _, f := range files {
		replaceInFile(f, regexp.MustCompile(`"version": ".*?"`), fmt.Sprintf(`"version": "%s"`, version))
	}
}

func updatePython(sdksDir, version string) {
	files := []string{
		filepath.Join(sdksDir, "thunkmetrc-py", "thunkmetrc", "pyproject.toml"),
	}
	for _, f := range files {
		replaceInFile(f, regexp.MustCompile(`version = ".*?"`), fmt.Sprintf(`version = "%s"`, version))
		replaceInFile(f, regexp.MustCompile(`thunkmetrc-client==.*?"`), fmt.Sprintf(`thunkmetrc-client==%s"`, version))
		replaceInFile(f, regexp.MustCompile(`thunkmetrc-wrapper==.*?"`), fmt.Sprintf(`thunkmetrc-wrapper==%s"`, version))
	}
}

func updateRust(sdksDir, version string) {
	files := []string{
		filepath.Join(sdksDir, "thunkmetrc-rs", "thunkmetrc", "Cargo.toml"),
	}
	for _, f := range files {
		content := readFile(f)
		if content == "" {
			continue
		}

		content = regexp.MustCompile(`(?m)^version = ".*"`).ReplaceAllStringFunc(content, func(match string) string {
			return fmt.Sprintf(`version = "%s"`, version)
		})

		content = regexp.MustCompile(`thunkmetrc-client = \{ path = "\.\./client", version = ".*?" \}`).
			ReplaceAllString(content, fmt.Sprintf(`thunkmetrc-client = { path = "../client", version = "%s" }`, version))
		content = regexp.MustCompile(`thunkmetrc-wrapper = \{ path = "\.\./wrapper", version = ".*?" \}`).
			ReplaceAllString(content, fmt.Sprintf(`thunkmetrc-wrapper = { path = "../wrapper", version = "%s" }`, version))
		writeFile(f, content)
	}
}

func updateJava(sdksDir, version string) {
	files := []string{
		filepath.Join(sdksDir, "thunkmetrc-java", "thunkmetrc", "pom.xml"),
	}
	for _, f := range files {
		updatePomVersion(f, version)
	}
}

func updateKotlin(sdksDir, version string) {
	files := []string{
		filepath.Join(sdksDir, "thunkmetrc-kotlin", "thunkmetrc", "pom.xml"),
	}
	for _, f := range files {
		updatePomVersion(f, version)
	}
}

func updateGo(sdksDir, version string) {
	file := filepath.Join(sdksDir, "thunkmetrc-go", "thunkmetrc", "go.mod")
	content := readFile(file)
	if content == "" {
		return
	}

	goVersion := normalizeGoModuleVersion(version)

	content = regexp.MustCompile(`github\.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client\s+v[^\s]+`).
		ReplaceAllString(content, "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client "+goVersion)
	content = regexp.MustCompile(`github\.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper\s+v[^\s]+`).
		ReplaceAllString(content, "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper "+goVersion)

	content = regexp.MustCompile(`(?ms)^replace\s*\(\s*.*?\)\s*`).
		ReplaceAllString(content, "")
	content = regexp.MustCompile(`(?m)^replace\s+github\.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/(client|wrapper)\s+=>\s+\.\./(client|wrapper)\s*$`).
		ReplaceAllString(content, "")
	content = regexp.MustCompile(`\n{3,}`).
		ReplaceAllString(content, "\n\n")

	writeFile(file, strings.TrimSpace(content)+"\n")
}

func updateGoWorkspace(sdksDir, version string) {
	file := filepath.Join(sdksDir, "..", "go.work")
	content := readFile(file)
	if content == "" {
		return
	}

	goVersion := normalizeGoModuleVersion(version)
	replacements := []struct {
		modulePath string
		localPath  string
	}{
		{
			modulePath: "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/client",
			localPath:  "./sdks/thunkmetrc-go/client",
		},
		{
			modulePath: "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/wrapper",
			localPath:  "./sdks/thunkmetrc-go/wrapper",
		},
		{
			modulePath: "github.com/thunkier/thunkmetrc/sdks/thunkmetrc-go/thunkmetrc",
			localPath:  "./sdks/thunkmetrc-go/thunkmetrc",
		},
	}

	for _, r := range replacements {
		re := regexp.MustCompile(`(?m)^(replace ` + regexp.QuoteMeta(r.modulePath) + ` )v[^\s]+( => ` + regexp.QuoteMeta(r.localPath) + `)$`)
		content = re.ReplaceAllString(content, "${1}"+goVersion+"${2}")
	}

	writeFile(file, content)
}

func updatePomVersion(path, version string) {
	content := readFile(path)
	if content == "" {
		return
	}

	patterns := []string{
		"thunkmetrc-client",
		"thunkmetrc-wrapper",
		"thunkmetrc-kotlin-client",
		"thunkmetrc-kotlin-wrapper",
		"thunkmetrc-kotlin",
		"thunkmetrc",
	}

	for _, artId := range patterns {
		re := regexp.MustCompile(fmt.Sprintf(`(<artifactId>%s</artifactId>\s*<version>)(.*?)(</version>)`, regexp.QuoteMeta(artId)))
		content = re.ReplaceAllString(content, "${1}"+version+"${3}")
	}

	writeFile(path, content)
}

func normalizeGoModuleVersion(version string) string {
	version = strings.TrimSpace(version)
	if version == "" {
		return "v0.0.0"
	}
	if strings.HasPrefix(version, "v") {
		return version
	}
	return "v" + version
}

func replaceInFile(path string, re *regexp.Regexp, replace string) {
	content, err := os.ReadFile(path)
	if err != nil {
		return
	}
	s := string(content)
	newS := re.ReplaceAllString(s, replace)
	if s != newS {
		if err := os.WriteFile(path, []byte(newS), 0644); err != nil {
			log.Printf("Err writing %s: %v", path, err)
		} else {
			log.Printf("Updated: %s", path)
		}
	}
}

func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
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
	} else {
		log.Printf("Updated: %s", path)
	}
}
