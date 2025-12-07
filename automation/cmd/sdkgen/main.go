package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/thunkmetrc/automation/pkg/bruno"
)

var (
	specsDir  = "../specs/metrc-bruno"
	outputDir = "../sdks"
)

func main() {
	flag.StringVar(&specsDir, "specs", specsDir, "Path to Bruno specs directory")
	flag.StringVar(&outputDir, "out", outputDir, "Path to output SDKs directory")
	flag.Parse()

	log.Println("Starting SDK Generation...")
	log.Printf("Specs: %s\n", specsDir)
	log.Printf("Output: %s\n", outputDir)

	collection, err := bruno.ParseCollection(specsDir)
	if err != nil {
		log.Fatalf("Failed to parse specs: %v", err)
	}
	log.Printf("Found %d requests in collection.\n", len(collection.Requests))

	groups := make(map[string][]bruno.Request)
	for _, req := range collection.Requests {
		req.URL = strings.ReplaceAll(req.URL, "{{baseUrl}}", "")
		req.URL = strings.ReplaceAll(req.URL, "{{url}}", "")

		if req.BodyType == "json" && req.Body != "" {
			schema, err := InferSchema(req.Body)
			if err != nil {
				log.Printf("[WARN] Failed to infer schema for %s: %v", req.Name, err)
			} else {
				req.BodySchema = schema
			}
		}

		if req.Group != "" {
			groups[req.Group] = append(groups[req.Group], req)
		}
	}
	log.Printf("Grouped into %d categories.\n", len(groups))

	for groupName, reqs := range groups {
		sort.Slice(reqs, func(i, j int) bool {
			return reqs[i].Name < reqs[j].Name
		})

		seen := make(map[string]int)
		for i, req := range reqs {
			baseName := cleanName(req.Name)
			checkName := strings.ToLower(baseName)

			if count, exists := seen[checkName]; exists {
				seen[checkName] = count + 1
				reqs[i].Name = fmt.Sprintf("%s_%d", baseName, count+1)
			} else {
				seen[checkName] = 1
			}
		}
		groups[groupName] = reqs
	}

	// Read version from VERSION file
	versionBytes, err := os.ReadFile("../VERSION")
	if err != nil {
		log.Fatalf("Failed to read VERSION file: %v", err)
	}
	version := strings.TrimSpace(string(versionBytes))
	log.Printf("Generating SDKs version: %s\n", version)

	generateCSharp(groups, version)
	generatePython(groups, version)
	generateTypeScript(groups, version)
	generateJava(groups, version)
	generateKotlin(groups, version)
	generateGo(groups, version)
	generateRust(groups, version)

	// Wrappers
	generateGoWrapper(groups, version)
	generateCSharpWrapper(groups, version)
	generateTypeScriptWrapper(groups, version)
	generatePythonWrapper(groups, version)
	generateJavaWrapper(groups, version)
	generateKotlinWrapper(groups, version)
	generateRustWrapper(groups, version)

	log.Println("SDK Generation Complete.")
}
