package runner

import (
	"bufio"
	"encoding/json"
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/thunkmetrc/automation/internal/runner/common"
	"github.com/thunkmetrc/automation/internal/runner/util"
	"github.com/thunkmetrc/automation/pkg/bruno"
	"github.com/thunkmetrc/automation/pkg/metrc/client"
	"github.com/thunkmetrc/automation/pkg/metrc/models"
)

func ExecuteGroupWorkflow(cfg common.Config, group string, reqs []bruno.Request) bool {
	groupName := strings.TrimSpace(group)

	// Filter: only run endpoints we want to actively debug for now
	// TO DEBUG MORE GROUPS:
	// 1. Add `|| strings.Contains(groupName, "GroupName")` to the condition below.
	// 2. Ensure permissions.go allows this group for the facility type.
	// 3. If the group requires special dependency injection (e.g. specific IDs), add a case in `getDependencyInjector`.
	isTarget := strings.Contains(groupName, "Additives") ||
		strings.Contains(groupName, "Caregivers Status") ||
		strings.Contains(groupName, "Tags") ||
		strings.Contains(groupName, "Units Of Measure") ||
		strings.Contains(groupName, "Waste Methods") ||
		strings.Contains(groupName, "Locations") ||
		strings.Contains(groupName, "Sublocations") ||
		strings.Contains(groupName, "Employees") ||
		strings.Contains(groupName, "Strains") ||
		strings.Contains(groupName, "Items")

	if !isTarget {
		return false
	}

	// Strict Filter: Caregivers Status requires config
	if strings.Contains(groupName, "Caregivers Status") {
		if cfg.ExternalConfig.CaregiverLicenseNumber == "" {
			cfg.Logger.Log("  [SKIP] Group '%s' skipped (CaregiverLicenseNumber not configured).", groupName)
			return false
		}
	}

	injector := getDependencyInjector(cfg, group, cfg.ExternalConfig.CaregiverLicenseNumber)
	filter := func(n string) bool {
		// Strict Filter for Employees: Only run GetPermissions V2
		if groupName == "Employees" {
			if !strings.Contains(n, "GetPermissions V2") {
				return false
			}
			if cfg.ExternalConfig.SpecialEmployeeUserKey == "" {
				cfg.Logger.Log("  [SKIP] Skipping Employees endpoint because SpecialEmployeeUserKey is not configured.")
				return false
			}
		}
		return true
	}

	return runMultiVersionChain(cfg, group, reqs, filter, injector)
}

func getDependencyInjector(cfg common.Config, group, caregiverLicenseNumber string) func(*bruno.Request) bool {
	groupName := strings.TrimSpace(group)
	return func(req *bruno.Request) bool {
		switch groupName {
		case "Additives Templates":
			client.ApplyRandomization(req)
			return true
		case "Additives":
			client.ApplyRandomization(req)
			return true
		case "Caregivers Status":
			req.URL = strings.ReplaceAll(req.URL, "{caregiverLicenseNumber}", caregiverLicenseNumber)
			return true
		case "Locations":
			client.ApplyRandomization(req)
			return client.InjectDependency(cfg,
				func() ([]client.SimpleName, error) {
					return client.NewFetcher(cfg).GetLocationTypes()
				},
				func(item client.SimpleName) string { return item.Name },
				"LocationTypeName",
				req)
		case "Employees":
			if cfg.ExternalConfig.SpecialEmployeeUserKey != "" {
				sep := "?"
				if strings.Contains(req.URL, "?") {
					sep = "&"
				}
				req.URL = req.URL + sep + "employeeLicenseNumber=" + cfg.ExternalConfig.SpecialEmployeeUserKey
			}
			return true
		case "Items":
			client.ApplyRandomization(req)

			fetcher := client.NewFetcher(cfg)
			cats, err1 := fetcher.Items_GetCategoriesV1()
			uoms, err2 := fetcher.UnitsOfMeasure_GetActiveV1()

			if err1 == nil && err2 == nil && len(cats) > 0 {
				idx := 0
				selectedCat := cats[idx]

				var selectedUoM *models.UnitsOfMeasure
				for _, u := range uoms {
					if u.QuantityType == selectedCat.QuantityType {
						val := u
						selectedUoM = &val
						break
					}
				}

				if selectedUoM != nil {
					newBody, _ := util.TransformBody(req.Body, func(item map[string]interface{}) bool {
						item["ItemCategory"] = selectedCat.Name
						item["UnitOfMeasure"] = selectedUoM.Name

						for _, rule := range cfg.ExternalConfig.SanitizeRules {
							// Simple glob match
							if matched, _ := path.Match(rule.Pattern, req.URL); matched {
								for _, field := range rule.Fields {
									delete(item, field)
								}
							}
						}

						return true
					})
					req.Body = newBody
				} else {
					cfg.Logger.Log("  [SKIP] No matching UoM found for category '%s' (Type: %s)", selectedCat.Name, selectedCat.QuantityType)
					return false
				}
			} else {
				cfg.Logger.Log("  [SKIP] Failed to fetch Categories or UoMs for Items injection.")
				return false
			}

			client.InjectDependency(cfg,
				func() ([]models.Strain, error) {
					return client.NewFetcher(cfg).Strains_GetActiveV1()
				},
				func(s models.Strain) string { return s.Name },
				"Strain",
				req)

			return true
		default:
			client.ApplyRandomization(req)
			return true
		}
	}
}

func CategorizeRequests(reqs []bruno.Request, filter func(string) bool) (lists, creators, getters, updaters, delLists, delByIds []*bruno.Request) {
	for i := range reqs {
		r := &reqs[i]
		if !filter(r.Name) {
			continue
		}

		isById := strings.Contains(r.Name, "ById") ||
			strings.Contains(r.Name, "GetBy") ||
			strings.Contains(r.URL, "/{") ||
			strings.Contains(r.URL, "{")

		if strings.Contains(r.Name, "Get") {
			if isById {
				getters = append(getters, r)
			} else {
				lists = append(lists, r)
			}
		} else if strings.Contains(r.Name, "CreateUpdate") {
			updaters = append(updaters, r)
		} else if strings.Contains(r.Name, "Create") || strings.Contains(r.Name, "POST") {
			creators = append(creators, r)
		} else if strings.Contains(r.Name, "Update") || strings.Contains(r.Name, "PUT") {
			updaters = append(updaters, r)
		} else if strings.Contains(r.Name, "Delete") {
			if isById {
				delByIds = append(delByIds, r)
			} else {
				delLists = append(delLists, r)
			}
		}
	}

	if len(delByIds) == 0 && len(delLists) > 0 {
		delByIds = delLists
		delLists = nil
	}
	return
}

func runMultiVersionChain(cfg common.Config, group string, reqs []bruno.Request, filter func(string) bool, dependencyInjector func(*bruno.Request) bool) bool {
	permissionFilter := func(name string) bool {
		if !filter(name) {
			return false
		}

		// Find the request object to check info directly (inefficient loop but safe)
		for _, r := range reqs {
			if r.Name == name {
				required := extractRequiredPermissions(r.Docs)
				if len(required) > 0 {
					if cfg.ExternalConfig.SpecialEmployeeUserKey == "" {
						return true
					}
					// Check against cfg.EmployeePermissions
					for _, reqPerm := range required {
						// Apply Override if exists
						targetPerm := reqPerm
						if override, ok := cfg.PermissionOverrides[reqPerm]; ok {
							targetPerm = override
						}

						hasPerm := false
						for _, userPerm := range cfg.EmployeePermissions {
							// Exact Match
							if strings.EqualFold(userPerm, targetPerm) {
								hasPerm = true
								break
							}
						}
						if !hasPerm {
							cfg.Logger.Log("  [SKIP-PERM] '%s' requires '%s' (mapped: '%s'). User missing permission.", name, reqPerm, targetPerm)
							return false
						}
					}
				}
				return true
			}
		}
		return true
	}

	lists, creators, getters, updaters, _, delByIds := CategorizeRequests(reqs, permissionFilter)
	fmt.Printf("DEBUG: Group %s - Lists: %d, Creators: %d, Getters: %d\n", group, len(lists), len(creators), len(getters))
	dataFound := false

	for _, l := range lists {
		body := client.ExecuteRequest(cfg, l, "")
		if util.HasData(body) {
			dataFound = true
		}
	}

	if len(creators) > 0 {
		for _, createReq := range creators {
			createReq.Group = group

			if dependencyInjector != nil {
				if !dependencyInjector(createReq) {
					continue
				}
			}

			client.ApplyRandomization(createReq)

			cfg.Logger.Log("  [%s] Executing Creator: %s", group, createReq.Name)
			body := client.ExecuteRequest(cfg, createReq, "")
			cfg.Logger.Log("  [DEBUG] Creator Body found: %s", body)
			id := util.ExtractID(body)
			cfg.Logger.Log("  [DEBUG] Extracted ID: %s", id)

			if id != "" {
				runActionsOnID(cfg, group, id, body, getters, updaters, delByIds, dependencyInjector)
				dataFound = true
			}
		}
	} else if len(lists) > 0 {
		// Fallback: If no creators, at least try to run getters/updaters on an existing ID from a list?
		// For now, logging will show we visited Lists.
	}

	if len(creators) == 0 && len(getters) > 0 {
		for _, req := range getters {
			req.Group = group
			if dependencyInjector != nil {
				if !dependencyInjector(req) {
					continue
				}
			}

			cfg.Logger.Log("  [%s] Executing Getter: %s (%s)", group, req.Name, req.URL)
			client.ExecuteRequest(cfg, req, "")
			dataFound = true
		}
	}

	return dataFound
}

func runActionsOnID(cfg common.Config, group, id, createBody string, getters, updaters, deleters []*bruno.Request, dependencyInjector func(*bruno.Request) bool) {
	for _, req := range getters {
		req.Group = group
		client.ExecuteRequest(cfg, req, id)
	}

	// Execute ALL updaters (V1, V2...)
	for _, req := range updaters {
		req.Group = group

		// Inject Dependencies
		if dependencyInjector != nil {
			if !dependencyInjector(req) {
				continue
			}
		}

		client.ApplyRandomization(req)

		if strings.TrimSpace(group) == "Additives Templates" {
			newUrl := strings.ReplaceAll(req.URL, "/"+id, "")
			req.URL = newUrl

			idInt, _ := strconv.Atoi(id)
			if idInt == 0 {
				idInt = -1 // Fallback
			}

			// Try Array First (since Bru files often use lists)
			var bodyArr []map[string]interface{}
			if err := json.Unmarshal([]byte(req.Body), &bodyArr); err == nil && len(bodyArr) > 0 {
				// Use First Item, Inject ID
				bodyArr[0]["Id"] = idInt
				// Send as single item list (since we only have one ID)
				newBytes, _ := json.MarshalIndent([]map[string]interface{}{bodyArr[0]}, "", "  ")
				req.Body = string(newBytes)
				client.ExecuteRequest(cfg, req, "")
			} else {
				// Try Object
				var bodyObj map[string]interface{}
				if err := json.Unmarshal([]byte(req.Body), &bodyObj); err == nil {
					bodyObj["Id"] = idInt
					newBytes, _ := json.MarshalIndent(bodyObj, "", "  ")
					req.Body = string(newBytes)
					client.ExecuteRequest(cfg, req, "")
				} else {
					// Fallback
					client.ExecuteRequest(cfg, req, id)
				}
			}
		} else if strings.TrimSpace(group) == "Sublocations" || strings.TrimSpace(group) == "Locations" || strings.TrimSpace(group) == "Strains" || strings.TrimSpace(group) == "Items" {
			// Parse All IDs from Create Body
			var createResp struct {
				Ids []int `json:"Ids"`
			}
			err := json.Unmarshal([]byte(createBody), &createResp)
			cfg.Logger.Log("  [DEBUG] ID Injection Group=%s. CreateBodyLen=%d. ParseErr=%v. IDs=%v", group, len(createBody), err, createResp.Ids)

			if len(createResp.Ids) > 0 {
				// We have IDs, try to map them to the update body array
				var bodyArr []map[string]interface{}
				if err := json.Unmarshal([]byte(req.Body), &bodyArr); err == nil {
					// Update Ids in the request body
					for i := 0; i < len(bodyArr); i++ {
						if i < len(createResp.Ids) {
							oldId := bodyArr[i]["Id"]
							bodyArr[i]["Id"] = createResp.Ids[i]
							cfg.Logger.Log("  [DEBUG] Injected ID: %v -> %d (Index %d)", oldId, createResp.Ids[i], i)
						}
					}
					// Trim body array to match number of created IDs (if we have fewer created than in update body)
					if len(bodyArr) > len(createResp.Ids) {
						bodyArr = bodyArr[:len(createResp.Ids)]
					}

					newBytes, _ := json.MarshalIndent(bodyArr, "", "  ")
					req.Body = string(newBytes)
				} else {
					cfg.Logger.Log("  [DEBUG] Failed to unmarshal update body as array: %v", err)
				}
			}
			client.ExecuteRequest(cfg, req, id)
		} else {
			client.ExecuteRequest(cfg, req, id)
		}
	}

	// Execute Deletes
	for _, req := range deleters {
		req.Group = group
		client.ExecuteRequest(cfg, req, id)
	}
}

func extractRequiredPermissions(docs string) []string {
	var perms []string
	scanner := bufio.NewScanner(strings.NewReader(docs))
	inPerms := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, "Permissions Required:") {
			inPerms = true
			continue
		}
		if inPerms {
			if strings.HasPrefix(line, "- ") {
				perm := strings.TrimSpace(strings.TrimPrefix(line, "- "))
				if perm != "" {
					perms = append(perms, perm)
				}
			} else if line == "" {
				// Continue consuming empty lines
			} else {
				// Broken out of list
				break
			}
		}
	}
	return perms
}
