package schemas

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/thunkier/thunkmetrc/tools/pkg/schema"
)

var (
	responsesDir string
	manualDir    string
	brunoDir     string
	outputDir    string
	internalMode bool
)

type Config struct {
	ResponsesDir string
	ManualDir    string
	BrunoDir     string
	OutputDir    string
	InternalMode bool
}

type CandidateModel struct {
	Group       string
	OriginalOp  string
	BaseName    string
	Schema      *schema.Schema
	IsRequest   bool
	IsPaginated bool
}

func Run(cfg Config) error {

	responsesDir = cfg.ResponsesDir
	manualDir = cfg.ManualDir
	brunoDir = cfg.BrunoDir
	outputDir = cfg.OutputDir
	internalMode = cfg.InternalMode

	if err := run(); err != nil {
		return err
	}
	return nil
}

func scanBruno(dir string, schemas map[string]*schemaEntry) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(info.Name(), ".bru") {
			return nil
		}

		relPath, _ := filepath.Rel(dir, path)
		parts := strings.Split(filepath.ToSlash(relPath), "/")
		if len(parts) < 2 {
			return nil
		}

		group := parts[len(parts)-2]
		opName := strings.TrimSuffix(parts[len(parts)-1], ".bru")
		key := fmt.Sprintf("%s/%s", group, opName)

		if group == "Packages" {
			log.Printf("DEBUG: Visiting %s. Key: %s", path, key)
		}

		content, err := os.ReadFile(path)
		if err != nil {
			log.Printf("[WARN] Could not read %s: %v", path, err)
			return nil
		}
		contentStr := string(content)

		reqJsonStr := extractRequestBody(contentStr)
		var reqSchema *schema.Schema
		var isRequestArray bool
		if reqJsonStr != "" {
			var reqData interface{}
			if err := json.Unmarshal([]byte(reqJsonStr), &reqData); err != nil {
				log.Printf("[WARN] Could not parse Request JSON in %s: %v", path, err)
			} else {
				inferred := schema.InferFromJSON(reqData)
				if inferred.Type == schema.TypeArray {
					isRequestArray = true
					reqSchema = inferred.UnwrapArray()
				} else {
					reqSchema = inferred
				}
			}
		}

		entry := schemas[key]

		if entry != nil && entry.schema != nil && !entry.schema.IsEmpty() {
			if reqSchema != nil {
				if entry.reqSchema == nil {
					entry.reqSchema = reqSchema
					entry.isRequestArray = isRequestArray
				}
			}

			return nil
		}

		jsonStr := extractExampleResponse(contentStr)

		if jsonStr == "" {
			if reqSchema == nil {
				return nil
			}
		}

		var itemSchema *schema.Schema
		var isPaginated bool
		var isWrite bool

		if jsonStr != "" {
			var jsonData interface{}
			if err := json.Unmarshal([]byte(jsonStr), &jsonData); err != nil {
				log.Printf("[WARN] Could not parse Response JSON in %s: %v", path, err)
			} else {
				inferred := schema.InferFromJSON(jsonData)

				if inferred.IsPaginated() {
					isPaginated = true
					itemSchema = inferred.UnwrapPaginated()
				} else if inferred.Type == schema.TypeArray {
					itemSchema = inferred.UnwrapArray()
				} else if inferred.IsWriteResponse() {
					isWrite = true
					itemSchema = inferred
				} else {
					itemSchema = inferred
				}
			}
		}

		newEntry := &schemaEntry{
			schema:         itemSchema,
			reqSchema:      reqSchema,
			isPaginated:    isPaginated,
			isRequestArray: isRequestArray,
			isWrite:        isWrite,
		}

		schemas[key] = newEntry
		return nil
	})
}

func extractRequestBody(content string) string {
	marker := "body:json {"
	idx := strings.Index(content, marker)
	if idx == -1 {
		return ""
	}

	jsonStart := -1
	for i := idx + len(marker); i < len(content); i++ {
		c := content[i]
		if !strings.ContainsRune(" \t\r\n", rune(c)) {
			if c == '{' || c == '[' {
				jsonStart = i
				break
			}
			if c == '}' {
				return ""
			}
		}
	}

	if jsonStart == -1 {
		return ""
	}

	return extractBalancedBlock(content[jsonStart:])
}

func extractExampleResponse(content string) string {
	docsStart := strings.Index(content, "docs {")
	if docsStart == -1 {
		return ""
	}

	exampleMarker := "Example Response:"
	idx := strings.Index(content[docsStart:], exampleMarker)
	if idx == -1 {
		return ""
	}

	startIdx := docsStart + idx + len(exampleMarker)

	jsonStart := -1
	for i := startIdx; i < len(content); i++ {
		c := content[i]
		if c == '{' || c == '[' {
			jsonStart = i
			break
		}
		if c == '}' && i > docsStart {
		}
	}

	if jsonStart == -1 {
		return ""
	}

	return extractBalancedBlock(content[jsonStart:])
}

func extractBalancedBlock(s string) string {
	if len(s) == 0 {
		return ""
	}
	startChar := s[0]
	endChar := byte(0)
	switch startChar {
	case '{':
		endChar = '}'
	case '[':
		endChar = ']'
	default:
		return ""
	}

	balance := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case startChar:
			balance++
		case endChar:
			balance--
		}

		if balance == 0 {
			return s[:i+1]
		}
	}
	return ""
}

type schemaEntry struct {
	schema         *schema.Schema
	reqSchema      *schema.Schema
	isPaginated    bool
	isWrite        bool
	isRequestArray bool
	innerSchema    *schema.Schema
}

func run() error {
	log.Printf("Generating JSON Schemas from:")
	log.Printf("  Manual: %s", manualDir)
	log.Printf("  Captured: %s", responsesDir)

	log.Printf("Output directory: %s", outputDir)

	manifest := make(map[string]*schema.ManifestEntry)
	schemas := make(map[string]*schemaEntry)
	hasWriteResponse := false

	scanDir := func(dir string, isManual bool) error {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if isManual {
				log.Printf("Manual directory %s does not exist, skipping", dir)
				return nil
			}
			return err
		}

		return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}

			if !strings.HasSuffix(info.Name(), ".json") {
				return nil
			}

			relPath, _ := filepath.Rel(dir, path)
			parts := strings.Split(filepath.ToSlash(relPath), "/")

			var group, opName string

			if len(parts) == 2 {
				group = parts[0]
				opName = strings.TrimSuffix(parts[1], ".json")
			} else if len(parts) == 3 {
				group = parts[1]
				opName = strings.TrimSuffix(parts[2], ".json")
			} else {
				return nil
			}

			if group == "Tags" {
				if strings.HasPrefix(opName, "GetAvailableTag") {
					opName = strings.Replace(opName, "GetAvailableTag", "GetAvailable", 1)
				}
			}

			key := fmt.Sprintf("%s/%s", group, opName)

			if _, exists := schemas[key]; exists {
				return nil
			}

			data, err := os.ReadFile(path)
			if err != nil {
				log.Printf("[WARN] Could not read %s: %v", path, err)
				return nil
			}

			var jsonData interface{}
			if err := json.Unmarshal(data, &jsonData); err != nil {
				log.Printf("[WARN] Could not parse %s: %v", path, err)
				return nil
			}

			inferred := schema.InferFromJSON(jsonData)

			itemSchema := inferred.UnwrapArray()

			entry := &schemaEntry{schema: itemSchema}

			if itemSchema.IsWriteResponse() {
				entry.isWrite = true
				hasWriteResponse = true
			} else if itemSchema.IsPaginated() {
				entry.isPaginated = true
				entry.innerSchema = itemSchema.UnwrapPaginated()
			}

			schemas[key] = entry
			return nil
		})
	}

	operationPaginated := make(map[string]bool)

	if err := scanDir(manualDir, true); err != nil {
		return fmt.Errorf("scan manual: %w", err)
	}

	if !internalMode {
		if err := scanDir(responsesDir, false); err != nil {
			return fmt.Errorf("scan responses: %w", err)
		}
	}

	if err := scanBruno(brunoDir, schemas); err != nil {
		return fmt.Errorf("scan bruno: %w", err)
	}

	for _, entry := range schemas {
		if entry.isWrite {
			hasWriteResponse = true
			break
		}
	}

	_ = os.RemoveAll(outputDir)

	if hasWriteResponse {
		writeSchema := &schema.Schema{
			Type: schema.TypeObject,
			Properties: map[string]*schema.Schema{
				"Ids":      {Type: schema.TypeArray, Items: &schema.Schema{Type: schema.TypeInteger}},
				"Warnings": {Type: schema.TypeString, Nullable: true},
			},
		}
		if err := writeSchemaFile(outputDir, "_common", "WriteResponse", writeSchema); err != nil {
			return err
		}
		log.Printf("Generated common WriteResponse schema")
	}

	paginatedSchema := &schema.Schema{
		Type: schema.TypeObject,
		Properties: map[string]*schema.Schema{
			"Data":          {Type: schema.TypeArray, Description: "The array of items"},
			"CurrentPage":   {Type: schema.TypeInteger},
			"Page":          {Type: schema.TypeInteger},
			"PageSize":      {Type: schema.TypeInteger},
			"RecordsOnPage": {Type: schema.TypeInteger},
			"Total":         {Type: schema.TypeInteger},
			"TotalPages":    {Type: schema.TypeInteger},
			"TotalRecords":  {Type: schema.TypeInteger},
		},
	}
	if err := writeSchemaFile(outputDir, "_common", "PaginatedResponse", paginatedSchema); err != nil {
		return err
	}
	log.Printf("Generated common PaginatedResponse schema")

	registry := make(map[string][]*CandidateModel)

	addCandidate := func(opKey string, entry *schemaEntry, isRequest bool) {
		parts := strings.Split(opKey, "/")
		group := parts[0]
		opName := parts[1]

		var baseName string
		var s *schema.Schema
		var isPaginated bool

		if isRequest {
			if entry.reqSchema == nil {
				return
			}
			baseName = deriveModelName(group, opName+"Request")
			s = entry.reqSchema
		} else {
			if entry.isWrite || entry.schema == nil {
				return
			}
			baseName = deriveModelName(group, opName)
			s = entry.schema
			if entry.isPaginated && entry.innerSchema != nil {
				s = entry.innerSchema
				isPaginated = true
			}
		}

		if s == nil || s.IsEmpty() {
			return
		}

		if len(baseName) > 0 {
			r := []rune(baseName)
			r[0] = unicode.ToUpper(r[0])
			baseName = string(r)
		}

		c := &CandidateModel{
			Group:       group,
			OriginalOp:  opKey,
			BaseName:    baseName,
			Schema:      s,
			IsRequest:   isRequest,
			IsPaginated: isPaginated,
		}
		registry[baseName] = append(registry[baseName], c)
	}

	for key, entry := range schemas {
		if !entry.isWrite {
			addCandidate(key, entry, false)
		}
		addCandidate(key, entry, true)

		if _, exists := manifest[key]; !exists {
			if entry.schema != nil || entry.reqSchema != nil {
				manifest[key] = &schema.ManifestEntry{
					IsWrite:     entry.isWrite,
					IsPaginated: entry.isPaginated,
				}
			}
		}
		if entry.isPaginated {
			operationPaginated[key] = true
		}
	}

	count := 0
	skipped := 0

	for name, candidates := range registry {
		bySig := make(map[string][]*CandidateModel)
		for _, c := range candidates {
			sig := c.Schema.Signature()
			bySig[sig] = append(bySig[sig], c)
		}

		if len(bySig) == 1 && len(candidates) > 1 {
			template := candidates[0]

			if err := writeSchemaFile(outputDir, "_common", name, template.Schema); err != nil {
				log.Printf("Error writing shared schema %s: %v", name, err)
				continue
			}

			schemaFile := fmt.Sprintf("_common/%s.schema.json", name)
			log.Printf("Consolidated %d usages of %s into shared schema", len(candidates), name)

			for _, c := range candidates {
				updateManifest(manifest, c, name, schemaFile, schemas)
			}
			count++
		} else if len(bySig) == 1 && len(candidates) == 1 {
			forceCommon := map[string]bool{
				"WasteMethod": true,
			}

			if forceCommon[name] {
				template := candidates[0]
				if err := writeSchemaFile(outputDir, "_common", name, template.Schema); err != nil {
					log.Printf("Error writing forced common schema %s: %v", name, err)
					continue
				}
				schemaFile := fmt.Sprintf("_common/%s.schema.json", name)
				updateManifest(manifest, template, name, schemaFile, schemas)
				count++
			} else {
				c := candidates[0]
				if err := writeSchemaFile(outputDir, c.Group, name, c.Schema); err != nil {
					continue
				}
				schemaFile := fmt.Sprintf("%s/%s.schema.json", c.Group, name)
				updateManifest(manifest, c, name, schemaFile, schemas)
				count++
			}
		} else {

			for sig, sigCandidates := range bySig {
				_ = sig
				for _, c := range sigCandidates {
					disambiguatedName := c.Group + name

					if err := writeSchemaFile(outputDir, c.Group, disambiguatedName, c.Schema); err != nil {
						continue
					}
					schemaFile := fmt.Sprintf("%s/%s.schema.json", c.Group, disambiguatedName)
					updateManifest(manifest, c, disambiguatedName, schemaFile, schemas)
				}
				count++
			}
		}
	}

	writeExclusions := map[string]bool{
		"Packages/FinishPackages":                true,
		"Packages/UnfinishPackages":              true,
		"ProcessingJob/FinishProcessingJob":      true,
		"ProcessingJob/UnfinishProcessingJob":    true,
		"ProcessingJob/CreateProcessingJobStart": true,
		"ProcessingJob/CreateProcessingJobTypes": true,
		"RetailId/CreateRetailIdGenerate":        true,
		"RetailId/CreateRetailIdPackagesInfo":    true,
	}
	writeVerbPrefixes := []string{"Create", "Update", "Delete", "Finish", "Unfinish", "Start", "Record"}
	for key, entry := range schemas {
		if writeExclusions[key] {
			continue
		}
		opName := strings.Split(key, "/")[1]
		isWriteOp := entry.isWrite
		if !isWriteOp {
			for _, prefix := range writeVerbPrefixes {
				if strings.HasPrefix(opName, prefix) {
					isWriteOp = true
					break
				}
			}
		}
		if isWriteOp {
			if manifest[key] == nil {
				if !internalMode {
					continue
				}
				manifest[key] = &schema.ManifestEntry{}
			}
			manifest[key].ResponseType = "WriteResponse"
			manifest[key].SchemaFile = "_common/WriteResponse.schema.json"
			manifest[key].IsWrite = true
		}
	}

	paginatedPrefixes := []string{
		"GetActive", "GetInactive", "GetOnHold", "GetTransferred",
		"GetIncoming", "GetOutgoing", "GetRejected",
		"GetAdmitted", "GetDiscontinued", "GetRestocked",
	}

	for key, entry := range manifest {
		if entry.IsPaginated {
			continue
		}

		if !entry.IsArray {
			continue
		}

		opName := strings.Split(key, "/")[1]

		isPaginatedName := false
		for _, prefix := range paginatedPrefixes {
			if strings.HasPrefix(opName, prefix) {
				isPaginatedName = true
				break
			}
		}

		if isPaginatedName {
			entry.IsPaginated = true
			log.Printf("Auto-inferred Pagination for %s", key)
		}
	}

	for opKey, entry := range manifest {

		parts := strings.Split(opKey, "/")
		if len(parts) < 2 {
			continue
		}
		group := parts[0]
		opName := parts[1]

		if entry.SchemaFile == "" {
			continue
		}

		var targetBase string
		if strings.HasPrefix(opName, "GetActive") {
			targetBase = strings.TrimPrefix(opName, "GetActive")
		} else if strings.HasPrefix(opName, "GetInactive") {
			targetBase = strings.TrimPrefix(opName, "GetInactive")
		} else {
			continue
		}

		singular := singularize(targetBase)
		byIdOpName := fmt.Sprintf("Get%sById", singular)
		byIdKey := fmt.Sprintf("%s/%s", group, byIdOpName)

		if manifest[byIdKey] == nil {
			manifest[byIdKey] = &schema.ManifestEntry{}
		}
		manifest[byIdKey].ResponseType = entry.ResponseType
		manifest[byIdKey].SchemaFile = entry.SchemaFile
		manifest[byIdKey].IsPaginated = false
		manifest[byIdKey].IsArray = false
		manifest[byIdKey].IsWrite = false

		log.Printf("Auto-mapped %s -> %s (Schema: %s)", opKey, byIdKey, entry.SchemaFile)
	}

	manifestBytes, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal manifest: %w", err)
	}
	manifestPath := filepath.Join(outputDir, "manifest.json")
	if err := os.WriteFile(manifestPath, manifestBytes, 0644); err != nil {
		return fmt.Errorf("failed to write manifest: %w", err)
	}
	log.Printf("Generated manifest with %d entries", len(manifest))

	log.Printf("Generated %d model schemas, skipped %d write responses", count, skipped)
	return nil
}

func updateManifest(manifest map[string]*schema.ManifestEntry, c *CandidateModel, finalName, schemaFile string, schemas map[string]*schemaEntry) {
	entry := manifest[c.OriginalOp]
	if entry == nil {
		entry = &schema.ManifestEntry{}
		manifest[c.OriginalOp] = entry
	}

	if c.IsRequest {
		entry.RequestType = finalName
		entry.RequestSchemaFile = schemaFile
		if originalEntry, ok := schemas[c.OriginalOp]; ok {
			entry.IsRequestArray = originalEntry.isRequestArray
		}
	} else {
		entry.ResponseType = finalName
		entry.SchemaFile = schemaFile

		opNamePart := strings.Split(c.OriginalOp, "/")[1]
		entry.IsArray = !strings.Contains(opNamePart, "ById")
	}
}

func writeSchemaFile(outputDir, group, name string, s *schema.Schema) error {
	if group != "_common" && s.IsEmpty() {
		return nil
	}

	title := group + name
	if group == "_common" || singularize(group) == name {
		title = name
	}
	schemaBytes, err := schema.ToJSONSchemaBytes(s, title)
	if err != nil {
		return err
	}

	outPath := filepath.Join(outputDir, group, name+".schema.json")
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return err
	}

	return os.WriteFile(outPath, schemaBytes, 0644)
}

func singularize(word string) string {
	overrides := map[string]string{
		"Statuses":         "Status",
		"PatientsStatuses": "PatientsStatus",
		"CaregiversStatus": "CaregiversStatus",
		"PatientsStatus":   "PatientsStatus",
		"Facilities":       "Facility",
		"PlantBatches":     "PlantBatch",
		"WasteMethods":     "WasteMethod",
		"UnitsOfMeasure":   "UnitOfMeasure",
	}
	if override, ok := overrides[word]; ok {
		return override
	}

	if strings.HasSuffix(word, "ies") {
		return strings.TrimSuffix(word, "ies") + "y"
	}
	if strings.HasSuffix(word, "uses") {
		return strings.TrimSuffix(word, "es")
	}
	if strings.HasSuffix(word, "ches") || strings.HasSuffix(word, "shes") || strings.HasSuffix(word, "xes") {
		return strings.TrimSuffix(word, "es")
	}
	if strings.HasSuffix(word, "sses") {
		return strings.TrimSuffix(word, "es")
	}
	if strings.HasSuffix(word, "s") && !strings.HasSuffix(word, "ss") && !strings.HasSuffix(word, "us") {
		return strings.TrimSuffix(word, "s")
	}
	return word
}

func deriveModelName(group, opName string) string {
	if group == "Tags" {
		if strings.HasPrefix(opName, "GetAvailable") {
			return "Tag"
		}
	}

	groupSingular := singularize(group)

	name := opName

	if before, ok := strings.CutSuffix(name, group); ok {
		name = before
	}

	for _, prefix := range []string{"Get", "Create", "Update", "Delete", "Record"} {
		if strings.HasPrefix(name, prefix) {
			name = strings.TrimPrefix(name, prefix)
			break
		}
	}

	for _, prefix := range []string{
		"Incoming", "Outgoing", "Rejected",
		"Flowering", "Vegetative",
	} {
		if strings.HasPrefix(name, prefix) {
			name = strings.TrimPrefix(name, prefix)
			break
		}
	}

	suffixes := []string{
		"ById", "ByLabel", "ByFacilityId", "ByDay", "ByPackageId",
		"Active", "Inactive", "OnHold", "Intransit", "Transferred",
		"LabSamples", "SourceHarvestById",
		"AdjustmentReasons", "ByCaregiverLicenseNumber", "ByPatientLicenseNumber",
	}
	for _, suffix := range suffixes {
		if strings.HasSuffix(name, suffix) {
			name = strings.TrimSuffix(name, suffix)
			break
		}
	}

	if strings.HasSuffix(name, "Reasons") {
		name = strings.TrimSuffix(name, "s")
	}

	if strings.HasPrefix(name, groupSingular) {
		rest := strings.TrimPrefix(name, groupSingular)
		if len(rest) <= 1 {
			return groupSingular
		}

		if strings.HasPrefix(rest, "By") {
			name = rest
		} else if strings.HasPrefix(rest, "s") {
			if len(rest) > 1 && unicode.IsUpper(rune(rest[1])) {
				name = strings.TrimPrefix(rest, "s")
			} else {
				name = strings.TrimPrefix(rest, "s")
			}
		}
	}

	name = singularize(name)

	if len(name) > 0 {
		r := []rune(name)
		r[0] = unicode.ToUpper(r[0])
		name = string(r)
	}

	if name == "" || name == groupSingular {
		return groupSingular
	}

	log.Printf("deriveModelName: %s/%s -> %s", group, opName, name)
	return name
}
