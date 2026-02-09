package specgen

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func createBrunoConfig(dir string) {
	config := `{
  "version": "1",
  "name": "Metrc API",
  "type": "collection",
  "ignore": [".git"]
}`
	os.WriteFile(filepath.Join(dir, "bruno.json"), []byte(config), 0644)
}

func GenerateBrunoFiles(ops []ParsedOperation, outDir string, permOverrides map[string]string) {
	// First pass: detect duplicate filenames across groups
	// Map of cleanFileName(op.Name) -> usage count
	nameCounts := make(map[string]int)
	for _, op := range ops {
		if isV1Endpoint(op.URL) {
			continue
		}
		cleanName := cleanFileName(op.Name)
		nameCounts[cleanName]++
	}

	for _, op := range ops {
		if isV1Endpoint(op.URL) {
			continue
		}

		groupDir := filepath.Join(outDir, cleanFileName(op.Group))
		os.MkdirAll(groupDir, 0755)

		// Generate the filename
		// Logic: If duplicate name detected, use [Prefix][Group][Suffix]
		// Otherwise, use original name
		var finalName string
		baseName := cleanFileName(op.Name)

		if count, ok := nameCounts[baseName]; ok && count > 1 {
			finalName = getBrunoFileName(op.Name, op.Group)
		} else {
			finalName = op.Name
		}

		fileName := fmt.Sprintf("%s.bru", cleanFileName(finalName))
		path := filepath.Join(groupDir, fileName)

		f, err := os.Create(path)
		if err != nil {
			continue
		}

		// Update op.Name to match the filename/metadata so internalgen picks up the unique name
		op.Name = finalName

		writeMeta(f, op)
		writeRequest(f, op)
		writeParams(f, op)
		writeAuth(f)
		writeBody(f, op)
		writeDocs(f, op, permOverrides)

		f.Close()
	}
}

func writeMeta(w io.Writer, op ParsedOperation) {
	fmt.Fprintf(w, "meta {\n  name: %s\n  type: http\n  seq: 1\n}\n\n", op.Name)
}

func writeRequest(w io.Writer, op ParsedOperation) {
	fmt.Fprintf(w, "%s {\n  url: {{baseUrl}}%s\n  body: json\n  auth: basic\n}\n\n", strings.ToLower(op.Method), op.URL)
}

func writeParams(w io.Writer, op ParsedOperation) {
	if len(op.Params) > 0 {
		fmt.Fprintln(w, "params:query {")
		for _, p := range op.Params {
			fmt.Fprintf(w, "  %s: \n", p.Name)
		}
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "")
	}
}

func writeAuth(w io.Writer) {
	fmt.Fprintln(w, "auth:basic {\n  username: {{username}}\n  password: {{password}}\n}")
	fmt.Fprintln(w, "")
}

func writeBody(w io.Writer, op ParsedOperation) {
	if op.Body != "" {
		fmt.Fprintln(w, "body:json {")
		fmt.Fprintln(w, op.Body)
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "")
	}
}

func writeDocs(w io.Writer, op ParsedOperation, overrides map[string]string) {
	fmt.Fprintln(w, "docs {")
	fmt.Fprintln(w, "  "+op.Description)
	if len(op.Permissions) > 0 {
		fmt.Fprintln(w, "\n  Permissions Required:")
		for _, p := range op.Permissions {
			permName := p
			if val, ok := overrides[p]; ok {
				permName = val
			}
			fmt.Fprintf(w, "  - %s\n", permName)
		}
	}
	if len(op.Params) > 0 {
		fmt.Fprintln(w, "\n  Parameters:")
		for _, p := range op.Params {
			optLabel := "Required"
			if !p.Required {
				optLabel = "Optional"
			}

			fmt.Fprintf(w, "  - %s (%s): %s\n", p.Name, optLabel, p.Description)
		}
	}

	if op.Response != "" {
		fmt.Fprintln(w, "\n  Example Response:")
		fmt.Fprintln(w, "  "+strings.ReplaceAll(op.Response, "\n", "\n  "))
	}
	fmt.Fprintln(w, "}")
}

func getBrunoFileName(opName, groupName string) string {
	prefixes := []string{"Create", "Update", "Delete", "Get", "Start", "Finish", "Unfinish", "Adjust"}

	// Clean up group name: remove spaces, ensure TitleCase if possible (simple approach)
	cleanGroup := strings.ReplaceAll(groupName, " ", "")
	cleanGroup = strings.ReplaceAll(cleanGroup, "-", "")
	// Ensure first letter is capitalized
	if len(cleanGroup) > 0 {
		cleanGroup = casesTitle(cleanGroup)
	}

	for _, prefix := range prefixes {
		if strings.HasPrefix(opName, prefix) {
			suffix := opName[len(prefix):]
			// Handle case where suffix might be empty or start with lowercase (though usually PascalCase)
			return prefix + cleanGroup + suffix
		}
	}

	// Fallback: Group + Name
	return cleanGroup + opName
}
