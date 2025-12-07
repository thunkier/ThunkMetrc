package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/thunkmetrc/automation/internal/runner/common"
	"github.com/thunkmetrc/automation/internal/runner/util"
	"github.com/thunkmetrc/automation/pkg/bruno"
)

func PrepareRequestURL(cfg common.Config, req *bruno.Request, id string) string {
	rawUrl := strings.ReplaceAll(req.URL, "{{baseUrl}}", cfg.BaseURL)

	if strings.Contains(rawUrl, "{date}") {
		today := time.Now().Format("2006-01-02")
		rawUrl = strings.ReplaceAll(rawUrl, "{date}", today)
	}

	if strings.Contains(rawUrl, "/v1/") {
		if u, err := url.Parse(rawUrl); err == nil {
			q := u.Query()
			if q.Has("pageNumber") || q.Has("pageSize") {
				q.Del("pageNumber")
				q.Del("pageSize")
				u.RawQuery = q.Encode()
				rawUrl = u.String()
			}
		}
	}

	if id != "" {
		if strings.Contains(rawUrl, "{id}") {
			rawUrl = strings.ReplaceAll(rawUrl, "{id}", id)
		} else if strings.HasSuffix(rawUrl, "/") {
			rawUrl = rawUrl + id
		}
	}

	if strings.Contains(rawUrl, "{label}") {
		rawUrl = injectLabel(cfg, rawUrl)
	}

	if strings.Contains(rawUrl, "packageId=") {
		rawUrl = injectPackageID(cfg, req, rawUrl)
	}

	rawUrl = injectQueryParams(cfg, req, rawUrl)

	return rawUrl
}

func injectLabel(cfg common.Config, rawUrl string) string {
	f := NewFetcher(cfg)
	var labelToInject string

	if strings.Contains(rawUrl, "/packages/") {
		pkgs, _ := f.GetActivePackages()
		if len(pkgs) > 0 {
			labelToInject = pkgs[0].Label
		}
	} else if strings.Contains(rawUrl, "/plants/") {
		plants, _ := f.GetVegetativePlants()
		if len(plants) > 0 {
			labelToInject = plants[0].Label
		}
	} else if strings.Contains(rawUrl, "/retailid/") {
		pkgs, _ := f.GetActivePackages()
		if len(pkgs) > 0 {
			labelToInject = pkgs[0].Label
		}
	}

	if labelToInject != "" {
		return strings.ReplaceAll(rawUrl, "{label}", labelToInject)
	}

	cfg.Logger.Log("  [WARN] Could not find active entity to inject for {label} in %s", rawUrl)
	return rawUrl
}

func injectPackageID(cfg common.Config, req *bruno.Request, rawUrl string) string {
	if !strings.Contains(rawUrl, "/packages/") {
		f := NewFetcher(cfg)
		pkgs, _ := f.GetActivePackages()
		var validPackageId int

		for _, p := range pkgs {
			if p.LabTestingState != "NotSubmitted" && p.LabTestingState != "" {
				validPackageId = p.Id
				break
			}
		}

		if validPackageId == 0 {
			cfg.Logger.Log("  [SKIP] No packages found with available Lab Tests. Skipping request %s", req.Name)
			return rawUrl
		}

		re := regexp.MustCompile(`packageId=[^&]+`)
		return re.ReplaceAllString(rawUrl, fmt.Sprintf("packageId=%d", validPackageId))
	}
	return rawUrl
}

func injectQueryParams(cfg common.Config, req *bruno.Request, rawUrl string) string {
	addParam := func(u string, key, val string) string {
		separator := "&"
		if !strings.Contains(u, "?") {
			separator = "?"
		}
		return fmt.Sprintf("%s%s%s=%s", u, separator, key, val)
	}

	if !strings.Contains(rawUrl, "licenseNumber=") && cfg.FacilityLicenseNumber != "" {
		rawUrl = addParam(rawUrl, "licenseNumber", cfg.FacilityLicenseNumber)
	}

	for k, v := range req.Params {
		val := v
		if k == "licenseNumber" {
			continue
		}

		if strings.Contains(rawUrl, "/v1/") {
			if k == "pageNumber" || k == "pageSize" {
				continue
			}
		}

		if val != "" {
			rawUrl = addParam(rawUrl, k, val)
		}
	}
	return rawUrl
}

func InjectDependency[T any](
	cfg common.Config,
	fetchFunc func() ([]T, error),
	extractor func(T) string,
	targetField string,
	req *bruno.Request,
) bool {
	items, err := fetchFunc()

	var value string
	if err == nil && len(items) > 0 {
		value = extractor(items[0])
	} else {
		cfg.Logger.Log("  [WARN] Failed to fetch dependency for field '%s': %v", targetField, err)
	}

	if value == "" {
		cfg.Logger.Log("  [SKIP] Missing dependency for '%s'. Skipping chain.", targetField)
		return false
	}

	req.Body, _ = util.TransformBody(req.Body, func(item map[string]interface{}) bool {
		item[targetField] = value
		return true
	})
	return true
}

func InjectTagOrDelete(cfg common.Config, f *Fetcher, jsonBody *string, invalidTagCallback func()) bool {
	newBody, _ := util.TransformBody(*jsonBody, func(item map[string]interface{}) bool {
		if _, ok := item["Tag"].(string); ok {
			if cfg.Facility.FacilityType.CanUseComplianceLabel {
				tags, _ := f.GetPackageTags()
				if len(tags) > 0 {
					item["Tag"] = tags[0].Tag
				} else {
					if invalidTagCallback != nil {
						invalidTagCallback()
					}
					return false
				}
			} else {
				delete(item, "Tag")
			}
		}
		return true
	})

	if newBody == "[]" {
		return false
	}
	*jsonBody = newBody
	return true
}

func GetValidLocation(f *Fetcher) string {
	locs, _ := f.GetActiveLocations()
	if len(locs) > 0 {
		return locs[0].Name
	}
	return ""
}

func GetValidItem(f *Fetcher) string {
	items, _ := f.GetActiveItems()
	if len(items) > 0 {
		return items[0].Name
	}
	return ""
}

func ApplyRandomization(create *bruno.Request) {
	var bodyArr []map[string]interface{}
	if err := json.Unmarshal([]byte(create.Body), &bodyArr); err != nil {
		var bodyObj map[string]interface{}
		if err2 := json.Unmarshal([]byte(create.Body), &bodyObj); err2 == nil {
			bodyArr = []map[string]interface{}{bodyObj}
		} else {
			return
		}
	}

	suffix := util.GenerateRandomString(4)

	for i := range bodyArr {
		if name, ok := bodyArr[i]["Name"].(string); ok {
			bodyArr[i]["Name"] = name + " " + suffix
		}
		if plate, ok := bodyArr[i]["LicensePlateNumber"].(string); ok {
			bodyArr[i]["LicensePlateNumber"] = plate + suffix
		}
		if dl, ok := bodyArr[i]["DriversLicenseNumber"].(string); ok {
			bodyArr[i]["DriversLicenseNumber"] = dl + suffix
		}
		if empId, ok := bodyArr[i]["EmployeeId"].(string); ok {
			bodyArr[i]["EmployeeId"] = empId + suffix
		}
		if ln, ok := bodyArr[i]["LicenseNumber"].(string); ok {
			bodyArr[i]["LicenseNumber"] = ln + suffix
		}
	}

	var bodyObj map[string]interface{}
	if json.Unmarshal([]byte(create.Body), &bodyObj) == nil {
		if len(bodyArr) > 0 {
			newBytes, _ := json.MarshalIndent(bodyArr[0], "", "  ")
			create.Body = string(newBytes)
		}
	} else {
		newBytes, _ := json.MarshalIndent(bodyArr, "", "  ")
		create.Body = string(newBytes)
	}
}
