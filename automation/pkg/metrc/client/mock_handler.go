package client

import (
	"strings"

	"github.com/thunkmetrc/automation/internal/runner/common"
	"github.com/thunkmetrc/automation/pkg/bruno"
)

func GetMockResponse(req *bruno.Request) string {
	if strings.Contains(req.URL, "items/v2/brand") && req.Method == "GET" {
		return common.Mocks.Brands
	}
	if strings.Contains(req.URL, "packages/v2/active") && req.Method == "GET" {
		return common.Mocks.Empty
	}
	if strings.Contains(req.URL, "items/v2/active") {
		return common.Mocks.ActiveItems
	}
	if strings.Contains(req.URL, "harvests/v1/active") {
		return common.Mocks.ActiveHarvests
	}
	if strings.Contains(req.URL, "plants/v1/vegetative") {
		return common.Mocks.Empty
	}
	if strings.Contains(req.URL, "tags/v2/package/available") {
		return common.Mocks.PackageTags
	}
	if strings.Contains(req.URL, "tags/v2/plant/available") {
		return common.Mocks.PlantTags
	}
	if strings.Contains(req.URL, "locations/v1/types") {
		return common.Mocks.LocationTypes
	}
	if strings.Contains(req.URL, "wastemethods") || strings.Contains(req.URL, "waste/types") {
		return common.Mocks.WasteMethods
	}
	if strings.Contains(req.URL, "labtests/v1/types") {
		return common.Mocks.LabTestTypes
	}
	if strings.Contains(req.URL, "processing/v2/jobtypes/categories") {
		return common.Mocks.JobTypeCategories
	}
	if strings.Contains(req.URL, "processing/v2/jobtypes/attributes") {
		return common.Mocks.JobTypeAttributes
	}
	if strings.Contains(req.URL, "processing/v2/jobtypes/active") {
		return common.Mocks.ProcessingJobTypes
	}
	if strings.Contains(req.URL, "items/v1/categories") {
		return common.Mocks.ItemCategories
	}
	if strings.Contains(req.URL, "strains/v1/active") {
		return common.Mocks.Strains
	}
	if strings.Contains(req.URL, "locations/v2/active") {
		return common.Mocks.ActiveLocations
	}
	if strings.Contains(req.URL, "facilities/v") {
		return common.Mocks.Facilities
	}

	if strings.Contains(req.URL, "?pageSize=") || strings.Contains(req.URL, "active") {
		return common.Mocks.GenericList
	}

	return `{"Id": 999999, "Name": "Mock Created Item"}`
}
