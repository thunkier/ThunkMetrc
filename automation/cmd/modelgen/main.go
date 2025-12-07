package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

var (
	brunoDir  = "../specs/metrc-bruno"
	outputDir = "pkg/metrc/models"
)

func main() {
	flag.StringVar(&brunoDir, "input", brunoDir, "Path to Bruno specs directory")
	flag.StringVar(&outputDir, "output", outputDir, "Path to output Go models directory")
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	groups, err := os.ReadDir(brunoDir)
	if err != nil {
		return err
	}

	typeMap := make(map[string]string)

	for _, group := range groups {
		if !group.IsDir() || strings.HasPrefix(group.Name(), ".") {
			continue
		}

		groupName := group.Name()
		fmt.Printf("Processing Group: %s\n", groupName)

		files, err := os.ReadDir(filepath.Join(brunoDir, groupName))
		if err != nil {
			continue
		}

		var structs []GoStruct

		rootName := singularize(cleanName(groupName))

		groupTypeMap := make(map[string]string)

		for _, f := range files {
			if !strings.HasSuffix(f.Name(), ".bru") {
				continue
			}

			path := filepath.Join(brunoDir, groupName, f.Name())
			content, err := os.ReadFile(path)
			if err != nil {
				continue
			}

			reqName := extractMetaName(string(content))

			resp := extractExampleResponse(string(content))
			if resp == "" {
				continue
			}

			if reqName != "" {
				key := fmt.Sprintf("%s:%s", groupName, reqName)
				groupTypeMap[key] = rootName
			}

			var data interface{}
			if err := json.Unmarshal([]byte(resp), &data); err != nil {
				fmt.Printf("  [WARN] Failed to parse JSON in %s: %v\n", f.Name(), err)
				continue
			}

			generated := generateStructs(rootName, data)
			structs = append(structs, generated...)
		}

		if len(structs) > 0 {
			writeModels(groupName, structs)

			for k, v := range groupTypeMap {
				typeMap[k] = v
			}
		} else {
			fmt.Printf("  [WARN] No structs generated for %s (skipping type map)\n", groupName)
		}
	}

	if err := writeTypeMap(typeMap); err != nil {
		return err
	}

	return nil
}

func extractMetaName(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "name:") {
			return strings.TrimSpace(strings.TrimPrefix(line, "name:"))
		}
	}
	return ""
}

func writeTypeMap(typeMap map[string]string) error {
	path := filepath.Join(outputDir, "type_map.json")
	data, err := json.MarshalIndent(typeMap, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		return err
	}
	fmt.Printf("Generated Type Map: %s (%d entries)\n", path, len(typeMap))
	return nil
}

func extractExampleResponse(content string) string {
	docsStart := strings.Index(content, "docs {")
	if docsStart == -1 {
		return ""
	}
	remaining := content[docsStart:]

	exStart := strings.Index(remaining, "Example Response:")
	if exStart == -1 {
		return ""
	}

	jsonStart := exStart + len("Example Response:")
	jsonText := remaining[jsonStart:]

	firstBrace := strings.IndexAny(jsonText, "{[")
	if firstBrace == -1 {
		return ""
	}
	jsonText = jsonText[firstBrace:]

	openChar := jsonText[0]
	closeChar := '}'
	if openChar == '[' {
		closeChar = ']'
	}

	balance := 0
	endIndex := -1

	for i, r := range jsonText {
		if r == rune(openChar) {
			balance++
		} else if r == closeChar {
			balance--
			if balance == 0 {
				endIndex = i + 1
				break
			}
		}
	}

	if endIndex != -1 {
		return jsonText[:endIndex]
	}

	return ""
}

type GoStruct struct {
	Name   string
	Fields []GoField
}

type GoField struct {
	Name string
	Type string
	JSON string
}

func generateStructs(rootName string, data interface{}) []GoStruct {
	var structs []GoStruct

	if arr, ok := data.([]interface{}); ok {
		if len(arr) > 0 {
			data = arr[0]
		} else {
			return nil
		}
	}

	obj, ok := data.(map[string]interface{})
	if !ok {
		return nil
	}

	fields := []GoField{}

	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := obj[k]
		fieldName := toExported(k)
		typeName := "string"

		switch val := v.(type) {
		case bool:
			typeName = "bool"
		case float64:
			if strings.HasSuffix(fieldName, "Percentage") ||
				strings.HasSuffix(fieldName, "Percent") ||
				strings.HasSuffix(fieldName, "Level") ||
				strings.HasSuffix(fieldName, "Content") ||
				strings.HasSuffix(fieldName, "Volume") ||
				strings.HasSuffix(fieldName, "Weight") ||
				strings.HasSuffix(fieldName, "Dose") {
				typeName = "float64"
			} else {
				if val == float64(int(val)) {
					typeName = "int"
				} else {
					typeName = "float64"
				}
			}
		case string:
			typeName = "string"
		case []interface{}:
			if len(val) > 0 {
				switch val[0].(type) {
				case string:
					typeName = "[]string"
				case float64:
					typeName = "[]float64"
				case map[string]interface{}:
					subName := rootName + singularize(toExported(k))
					typeName = "[]" + subName
					structs = append(structs, generateStructs(subName, val[0])...)
				}
			} else {
				typeName = "[]interface{}"
			}
		case map[string]interface{}:
			subName := rootName + toExported(k)
			if toExported(k) == "FacilityType" {
				subName = "FacilityType"
			}
			typeName = subName
			structs = append(structs, generateStructs(subName, val)...)
		case nil:
			if strings.HasSuffix(fieldName, "Hours") ||
				strings.HasSuffix(fieldName, "Days") ||
				strings.HasSuffix(fieldName, "Allowed") ||
				strings.HasSuffix(fieldName, "Count") ||
				strings.HasSuffix(fieldName, "Id") {
				typeName = "*int"
			} else if strings.HasSuffix(fieldName, "Level") ||
				strings.HasSuffix(fieldName, "Percentage") ||
				strings.HasSuffix(fieldName, "Percent") ||
				strings.HasSuffix(fieldName, "Content") ||
				strings.HasSuffix(fieldName, "Volume") ||
				strings.HasSuffix(fieldName, "Weight") ||
				strings.HasSuffix(fieldName, "Dose") {
				typeName = "*float64"
			} else if strings.HasPrefix(fieldName, "Allow") ||
				strings.HasPrefix(fieldName, "Is") ||
				strings.HasPrefix(fieldName, "Can") ||
				strings.HasPrefix(fieldName, "Enable") ||
				strings.HasPrefix(fieldName, "Require") {
				typeName = "*bool"
			} else {
				typeName = "*string"
			}
		}

		fields = append(fields, GoField{
			Name: fieldName,
			Type: typeName,
			JSON: k,
		})
	}

	structs = append(structs, GoStruct{
		Name:   rootName,
		Fields: fields,
	})

	return structs
}

func writeModels(group string, structs []GoStruct) {
	unique := make(map[string]GoStruct)
	for _, s := range structs {
		if existing, ok := unique[s.Name]; ok {
			for _, f := range s.Fields {
				found := false
				for _, ef := range existing.Fields {
					if ef.Name == f.Name {
						found = true
						break
					}
				}
				if !found {
					existing.Fields = append(existing.Fields, f)
				}
			}
			unique[s.Name] = existing
		} else {
			unique[s.Name] = s
		}
	}

	filename := filepath.Join(outputDir, group+".go")
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	fmt.Fprintf(f, "package models\n\n")

	names := make([]string, 0, len(unique))
	for n := range unique {
		names = append(names, n)
	}
	sort.Strings(names)

	for _, name := range names {
		s := unique[name]
		fmt.Fprintf(f, "type %s struct {\n", s.Name)
		for _, fld := range s.Fields {
			fmt.Fprintf(f, "\t%s %s `json:\"%s\"`\n", fld.Name, fld.Type, fld.JSON)
		}
		fmt.Fprintf(f, "}\n\n")
	}

	fmt.Printf("Generated %s (%d structs)\n", filename, len(unique))
}

func toExported(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func cleanName(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "/", "")
	return toExported(s)
}

func singularize(s string) string {
	if strings.HasSuffix(s, "ies") {
		return strings.TrimSuffix(s, "ies") + "y"
	}
	if strings.HasSuffix(s, "s") && !strings.HasSuffix(s, "ss") {
		return strings.TrimSuffix(s, "s")
	}
	return s
}
