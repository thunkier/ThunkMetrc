package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type MetrcClient struct {
	BaseURL   string
	VendorKey string
	UserKey   string
	Client  *http.Client
}

func New(baseURL, vendorKey, userKey string) *MetrcClient {
	return &MetrcClient{
		BaseURL:   strings.TrimRight(baseURL, "/"),
		VendorKey: vendorKey,
		UserKey:   userKey,
		Client:    &http.Client{},
	}
}

func (c *MetrcClient) send(method, path string, queryParams map[string]string, body interface{}) (interface{}, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil { return nil, err }
		bodyReader = bytes.NewReader(jsonBytes)
	}

	u, err := url.Parse(c.BaseURL + path)
	if err != nil { return nil, err }
	q := u.Query()
	for k, v := range queryParams {
		if v != "" {
			q.Set(k, v)
		}
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil { return nil, err }

	req.SetBasicAuth(c.VendorKey, c.UserKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.Client.Do(req)
	if err != nil { return nil, err }
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API Error: %d", resp.StatusCode)
	}
	if resp.StatusCode == 204 {
		return nil, nil
	}

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// POST CreateAssociate V2
// Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
//   - WebApi Retail ID Read Write State (All or WriteOnly)
//   - Industry/View Packages
//   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
func (c *MetrcClient) RetailIdCreateAssociateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/retailid/v2/associate", queryParams, body)
}

// POST CreateGenerate V2
// Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
//   - WebApi Retail ID Read Write State (All or WriteOnly)
//   - Industry/View Packages
//   - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
func (c *MetrcClient) RetailIdCreateGenerateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/retailid/v2/generate", queryParams, body)
}

// POST CreateMerge V2
// Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
//   - WebApi Retail ID Read Write State (All or WriteOnly)
//   - Key Value Settings/Retail ID Merge Packages Enabled
func (c *MetrcClient) RetailIdCreateMergeV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/retailid/v2/merge", queryParams, body)
}

// POST CreatePackageInfo V2
// Retrieves Package information for given list of Package labels.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
//   - WebApi Retail ID Read Write State (All or WriteOnly)
//   - Industry/View Packages
//   - Admin/Employees/Packages Page/Product Labels(Manage)
func (c *MetrcClient) RetailIdCreatePackageInfoV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/retailid/v2/packages/info", queryParams, body)
}

// GET GetReceiveByLabel V2
// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Manage RetailId
//   - WebApi Retail ID Read Write State (All or ReadOnly)
//   - Industry/View Packages
//   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
func (c *MetrcClient) RetailIdGetReceiveByLabelV2(label string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/retailid/v2/receive/"+url.QueryEscape(label)+"", queryParams, nil)
}

// GET GetReceiveQrByShortCode V2
// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Manage RetailId
//   - WebApi Retail ID Read Write State (All or ReadOnly)
//   - Industry/View Packages
//   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
func (c *MetrcClient) RetailIdGetReceiveQrByShortCodeV2(shortcode string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/retailid/v2/receive/qr/"+url.QueryEscape(shortcode)+"", queryParams, nil)
}

// POST CreateIntegratorSetup V2
// This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) SandboxCreateIntegratorSetupV2(userkey string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if userkey != "" {
		queryParams["userKey"] = userkey
	}
	return c.send("POST", "/sandbox/v2/integrator/setup", queryParams, body)
}

// POST Create V1
// Permissions Required:
//   - Manage Strains
func (c *MetrcClient) StrainsCreateV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/strains/v1/create", queryParams, body)
}

// POST Create V2
// Creates new strain records for a specified Facility.
// 
//   Permissions Required:
//   - Manage Strains
func (c *MetrcClient) StrainsCreateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/strains/v2", queryParams, body)
}

// POST CreateUpdate V1
// Permissions Required:
//   - Manage Strains
func (c *MetrcClient) StrainsCreateUpdateV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/strains/v1/update", queryParams, body)
}

// DELETE Delete V1
// Permissions Required:
//   - Manage Strains
func (c *MetrcClient) StrainsDeleteV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/strains/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE Delete V2
// Archives an existing strain record for a Facility
// 
//   Permissions Required:
//   - Manage Strains
func (c *MetrcClient) StrainsDeleteV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/strains/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V1
// Permissions Required:
//   - Manage Strains
func (c *MetrcClient) StrainsGetV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/strains/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V2
// Retrieves a Strain record by its Id, with an optional license number.
// 
//   Permissions Required:
//   - Manage Strains
func (c *MetrcClient) StrainsGetV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/strains/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetActive V1
// Permissions Required:
//   - Manage Strains
func (c *MetrcClient) StrainsGetActiveV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/strains/v1/active", queryParams, nil)
}

// GET GetActive V2
// Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Strains
func (c *MetrcClient) StrainsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/strains/v2/active", queryParams, nil)
}

// GET GetInactive V2
// Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Strains
func (c *MetrcClient) StrainsGetInactiveV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/strains/v2/inactive", queryParams, nil)
}

// PUT Update V2
// Updates existing strain records for a specified Facility.
// 
//   Permissions Required:
//   - Manage Strains
func (c *MetrcClient) StrainsUpdateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/strains/v2", queryParams, body)
}

// GET GetPackageAvailable V2
// Returns a list of available package tags. NOTE: This is a premium endpoint.
// 
//   Permissions Required:
//   - Manage Tags
func (c *MetrcClient) TagsGetPackageAvailableV2(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/tags/v2/package/available", queryParams, nil)
}

// GET GetPlantAvailable V2
// Returns a list of available plant tags. NOTE: This is a premium endpoint.
// 
//   Permissions Required:
//   - Manage Tags
func (c *MetrcClient) TagsGetPlantAvailableV2(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/tags/v2/plant/available", queryParams, nil)
}

// GET GetStaged V2
// Returns a list of staged tags. NOTE: This is a premium endpoint.
// 
//   Permissions Required:
//   - Manage Tags
//   - RetailId.AllowPackageStaging Key Value enabled
func (c *MetrcClient) TagsGetStagedV2(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/tags/v2/staged", queryParams, nil)
}

// GET GetStatusesByPatientLicenseNumber V1
// Data returned by this endpoint is cached for up to one minute.
// 
//   Permissions Required:
//   - Lookup Patients
func (c *MetrcClient) PatientsStatusGetStatusesByPatientLicenseNumberV1(patientlicensenumber string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/patients/v1/statuses/"+url.QueryEscape(patientlicensenumber)+"", queryParams, nil)
}

// GET GetStatusesByPatientLicenseNumber V2
// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
// 
//   Permissions Required:
//   - Lookup Patients
func (c *MetrcClient) PatientsStatusGetStatusesByPatientLicenseNumberV2(patientlicensenumber string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/patients/v2/statuses/"+url.QueryEscape(patientlicensenumber)+"", queryParams, nil)
}

// POST CreateAdditives V1
// Permissions Required:
//   - Manage Plants Additives
func (c *MetrcClient) PlantBatchesCreateAdditivesV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v1/additives", queryParams, body)
}

// POST CreateAdditives V2
// Records Additive usage details for plant batches at a specific Facility.
// 
//   Permissions Required:
//   - Manage Plants Additives
func (c *MetrcClient) PlantBatchesCreateAdditivesV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v2/additives", queryParams, body)
}

// POST CreateAdditivesUsingtemplate V2
// Records Additive usage for plant batches at a Facility using predefined additive templates.
// 
//   Permissions Required:
//   - Manage Plants Additives
func (c *MetrcClient) PlantBatchesCreateAdditivesUsingtemplateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v2/additives/usingtemplate", queryParams, body)
}

// POST CreateAdjust V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (c *MetrcClient) PlantBatchesCreateAdjustV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v1/adjust", queryParams, body)
}

// POST CreateAdjust V2
// Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (c *MetrcClient) PlantBatchesCreateAdjustV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v2/adjust", queryParams, body)
}

// POST CreateChangegrowthphase V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantBatchesCreateChangegrowthphaseV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v1/changegrowthphase", queryParams, body)
}

// POST CreateGrowthphase V2
// Updates the growth phase of plants at a specified Facility based on tracking information.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantBatchesCreateGrowthphaseV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v2/growthphase", queryParams, body)
}

// POST CreatePackage V2
// Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PlantBatchesCreatePackageV2(isfrommotherplant string, licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if isfrommotherplant != "" {
		queryParams["isFromMotherPlant"] = isfrommotherplant
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v2/packages", queryParams, body)
}

// POST CreatePackageFrommotherplant V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PlantBatchesCreatePackageFrommotherplantV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v1/create/packages/frommotherplant", queryParams, body)
}

// POST CreatePackageFrommotherplant V2
// Creates packages from mother plants at the specified Facility.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PlantBatchesCreatePackageFrommotherplantV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v2/packages/frommotherplant", queryParams, body)
}

// POST CreatePlantings V2
// Creates new plantings for a Facility by generating plant batches based on provided planting details.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (c *MetrcClient) PlantBatchesCreatePlantingsV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v2/plantings", queryParams, body)
}

// POST CreateSplit V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (c *MetrcClient) PlantBatchesCreateSplitV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v1/split", queryParams, body)
}

// POST CreateSplit V2
// Splits an existing Plant Batch into multiple groups at the specified Facility.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (c *MetrcClient) PlantBatchesCreateSplitV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v2/split", queryParams, body)
}

// POST CreateWaste V1
// Permissions Required:
//   - Manage Plants Waste
func (c *MetrcClient) PlantBatchesCreateWasteV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v1/waste", queryParams, body)
}

// POST CreateWaste V2
// Records waste information for plant batches based on the submitted data for the specified Facility.
// 
//   Permissions Required:
//   - Manage Plants Waste
func (c *MetrcClient) PlantBatchesCreateWasteV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v2/waste", queryParams, body)
}

// POST Createpackages V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PlantBatchesCreatepackagesV1(isfrommotherplant string, licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if isfrommotherplant != "" {
		queryParams["isFromMotherPlant"] = isfrommotherplant
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v1/createpackages", queryParams, body)
}

// POST Createplantings V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
func (c *MetrcClient) PlantBatchesCreateplantingsV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plantbatches/v1/createplantings", queryParams, body)
}

// DELETE Delete V1
// Permissions Required:
//   - View Immature Plants
//   - Destroy Immature Plants
func (c *MetrcClient) PlantBatchesDeleteV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/plantbatches/v1", queryParams, nil)
}

// DELETE Delete V2
// Completes the destruction of plant batches based on the provided input data.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Destroy Immature Plants
func (c *MetrcClient) PlantBatchesDeleteV2(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/plantbatches/v2", queryParams, nil)
}

// GET Get V1
// Permissions Required:
//   - View Immature Plants
func (c *MetrcClient) PlantBatchesGetV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plantbatches/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V2
// Retrieves a Plant Batch by Id.
// 
//   Permissions Required:
//   - View Immature Plants
func (c *MetrcClient) PlantBatchesGetV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plantbatches/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetActive V1
// Permissions Required:
//   - View Immature Plants
func (c *MetrcClient) PlantBatchesGetActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plantbatches/v1/active", queryParams, nil)
}

// GET GetActive V2
// Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
// 
//   Permissions Required:
//   - View Immature Plants
func (c *MetrcClient) PlantBatchesGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plantbatches/v2/active", queryParams, nil)
}

// GET GetInactive V1
// Permissions Required:
//   - View Immature Plants
func (c *MetrcClient) PlantBatchesGetInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plantbatches/v1/inactive", queryParams, nil)
}

// GET GetInactive V2
// Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
// 
//   Permissions Required:
//   - View Immature Plants
func (c *MetrcClient) PlantBatchesGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plantbatches/v2/inactive", queryParams, nil)
}

// GET GetTypes V1
// Permissions Required:
//   - None
func (c *MetrcClient) PlantBatchesGetTypesV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/plantbatches/v1/types", queryParams, nil)
}

// GET GetTypes V2
// Retrieves a list of plant batch types.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) PlantBatchesGetTypesV2(pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plantbatches/v2/types", queryParams, nil)
}

// GET GetWaste V2
// Retrieves waste details associated with plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Plants Waste
func (c *MetrcClient) PlantBatchesGetWasteV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plantbatches/v2/waste", queryParams, nil)
}

// GET GetWasteReasons V1
// Permissions Required:
//   - None
func (c *MetrcClient) PlantBatchesGetWasteReasonsV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plantbatches/v1/waste/reasons", queryParams, nil)
}

// GET GetWasteReasons V2
// Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) PlantBatchesGetWasteReasonsV2(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plantbatches/v2/waste/reasons", queryParams, nil)
}

// PUT UpdateLocation V2
// Moves one or more plant batches to new locations with in a specified Facility.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants
func (c *MetrcClient) PlantBatchesUpdateLocationV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plantbatches/v2/location", queryParams, body)
}

// PUT UpdateMoveplantbatches V1
// Permissions Required:
//   - View Immature Plants
func (c *MetrcClient) PlantBatchesUpdateMoveplantbatchesV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plantbatches/v1/moveplantbatches", queryParams, body)
}

// PUT UpdateName V2
// Renames plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantBatchesUpdateNameV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plantbatches/v2/name", queryParams, body)
}

// PUT UpdateStrain V2
// Changes the strain of plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantBatchesUpdateStrainV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plantbatches/v2/strain", queryParams, body)
}

// PUT UpdateTag V2
// Replaces tags for plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantBatchesUpdateTagV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plantbatches/v2/tag", queryParams, body)
}

// POST CreateAdjust V1
// Permissions Required:
//   - ManageProcessingJobs
func (c *MetrcClient) ProcessingJobsCreateAdjustV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/processing/v1/adjust", queryParams, body)
}

// POST CreateAdjust V2
// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsCreateAdjustV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/processing/v2/adjust", queryParams, body)
}

// POST CreateJobtypes V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsCreateJobtypesV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/processing/v1/jobtypes", queryParams, body)
}

// POST CreateJobtypes V2
// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsCreateJobtypesV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/processing/v2/jobtypes", queryParams, body)
}

// POST CreateStart V1
// Permissions Required:
//   - ManageProcessingJobs
func (c *MetrcClient) ProcessingJobsCreateStartV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/processing/v1/start", queryParams, body)
}

// POST CreateStart V2
// Initiates new processing jobs at a Facility, including job details and associated packages.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsCreateStartV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/processing/v2/start", queryParams, body)
}

// POST Createpackages V1
// Permissions Required:
//   - ManageProcessingJobs
func (c *MetrcClient) ProcessingJobsCreatepackagesV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/processing/v1/createpackages", queryParams, body)
}

// POST Createpackages V2
// Creates packages from processing jobs at a Facility, including optional location and note assignments.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsCreatepackagesV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/processing/v2/createpackages", queryParams, body)
}

// DELETE Delete V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsDeleteV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/processing/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE Delete V2
// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsDeleteV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/processing/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteJobtypes V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsDeleteJobtypesV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/processing/v1/jobtypes/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteJobtypes V2
// Archives a Processing Job Type at a Facility, making it inactive for future use.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsDeleteJobtypesV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/processing/v2/jobtypes/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/processing/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V2
// Retrieves a ProcessingJob by Id.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/processing/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetActive V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/processing/v1/active", queryParams, nil)
}

// GET GetActive V2
// Retrieves active processing jobs at a specified Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/processing/v2/active", queryParams, nil)
}

// GET GetInactive V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/processing/v1/inactive", queryParams, nil)
}

// GET GetInactive V2
// Retrieves inactive processing jobs at a specified Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/processing/v2/inactive", queryParams, nil)
}

// GET GetJobtypesActive V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetJobtypesActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/processing/v1/jobtypes/active", queryParams, nil)
}

// GET GetJobtypesActive V2
// Retrieves a list of all active processing job types defined within a Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetJobtypesActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/processing/v2/jobtypes/active", queryParams, nil)
}

// GET GetJobtypesAttributes V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetJobtypesAttributesV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/processing/v1/jobtypes/attributes", queryParams, nil)
}

// GET GetJobtypesAttributes V2
// Retrieves all processing job attributes available for a Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetJobtypesAttributesV2(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/processing/v2/jobtypes/attributes", queryParams, nil)
}

// GET GetJobtypesCategories V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetJobtypesCategoriesV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/processing/v1/jobtypes/categories", queryParams, nil)
}

// GET GetJobtypesCategories V2
// Retrieves all processing job categories available for a specified Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetJobtypesCategoriesV2(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/processing/v2/jobtypes/categories", queryParams, nil)
}

// GET GetJobtypesInactive V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetJobtypesInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/processing/v1/jobtypes/inactive", queryParams, nil)
}

// GET GetJobtypesInactive V2
// Retrieves a list of all inactive processing job types defined within a Facility.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsGetJobtypesInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/processing/v2/jobtypes/inactive", queryParams, nil)
}

// PUT UpdateFinish V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsUpdateFinishV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/processing/v1/finish", queryParams, body)
}

// PUT UpdateFinish V2
// Completes processing jobs at a Facility by recording final notes and waste measurements.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsUpdateFinishV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/processing/v2/finish", queryParams, body)
}

// PUT UpdateJobtypes V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsUpdateJobtypesV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/processing/v1/jobtypes", queryParams, body)
}

// PUT UpdateJobtypes V2
// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsUpdateJobtypesV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/processing/v2/jobtypes", queryParams, body)
}

// PUT UpdateUnfinish V1
// Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsUpdateUnfinishV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/processing/v1/unfinish", queryParams, body)
}

// PUT UpdateUnfinish V2
// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
// 
//   Permissions Required:
//   - Manage Processing Job
func (c *MetrcClient) ProcessingJobsUpdateUnfinishV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/processing/v2/unfinish", queryParams, body)
}

// POST Create V2
// Creates new sublocation records for a Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) SublocationsCreateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sublocations/v2", queryParams, body)
}

// DELETE Delete V2
// Archives an existing Sublocation record for a Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) SublocationsDeleteV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/sublocations/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V2
// Retrieves a Sublocation by its Id, with an optional license number.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) SublocationsGetV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sublocations/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetActive V2
// Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) SublocationsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/sublocations/v2/active", queryParams, nil)
}

// GET GetInactive V2
// Retrieves a list of inactive sublocations for the specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) SublocationsGetInactiveV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/sublocations/v2/inactive", queryParams, nil)
}

// PUT Update V2
// Updates existing sublocation records for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) SublocationsUpdateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sublocations/v2", queryParams, body)
}

// POST CreateExternalIncoming V1
// Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersCreateExternalIncomingV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/transfers/v1/external/incoming", queryParams, body)
}

// POST CreateExternalIncoming V2
// Creates external incoming shipment plans for a Facility.
// 
//   Permissions Required:
//   - Manage Transfers
func (c *MetrcClient) TransfersCreateExternalIncomingV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/transfers/v2/external/incoming", queryParams, body)
}

// POST CreateTemplates V1
// Permissions Required:
//   - Transfer Templates
func (c *MetrcClient) TransfersCreateTemplatesV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/transfers/v1/templates", queryParams, body)
}

// POST CreateTemplatesOutgoing V2
// Creates new transfer templates for a Facility.
// 
//   Permissions Required:
//   - Manage Transfer Templates
func (c *MetrcClient) TransfersCreateTemplatesOutgoingV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/transfers/v2/templates/outgoing", queryParams, body)
}

// DELETE DeleteExternalIncoming V1
// Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersDeleteExternalIncomingV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/transfers/v1/external/incoming/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteExternalIncoming V2
// Voids an external incoming shipment plan for a Facility.
// 
//   Permissions Required:
//   - Manage Transfers
func (c *MetrcClient) TransfersDeleteExternalIncomingV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/transfers/v2/external/incoming/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteTemplates V1
// Permissions Required:
//   - Transfer Templates
func (c *MetrcClient) TransfersDeleteTemplatesV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/transfers/v1/templates/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteTemplatesOutgoing V2
// Archives a transfer template for a Facility.
// 
//   Permissions Required:
//   - Manage Transfer Templates
func (c *MetrcClient) TransfersDeleteTemplatesOutgoingV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/transfers/v2/templates/outgoing/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetDeliveriesPackagesStates V1
// Permissions Required:
//   - None
func (c *MetrcClient) TransfersGetDeliveriesPackagesStatesV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v1/deliveries/packages/states", queryParams, nil)
}

// GET GetDeliveriesPackagesStates V2
// Returns a list of available shipment Package states.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) TransfersGetDeliveriesPackagesStatesV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v2/deliveries/packages/states", queryParams, nil)
}

// GET GetDelivery V1
// Please note: that the {id} parameter above represents a Shipment Plan ID.
// 
//   Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersGetDeliveryV1(id string, no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v1/"+url.QueryEscape(id)+"/deliveries", queryParams, nil)
}

// GET GetDelivery V2
// Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (c *MetrcClient) TransfersGetDeliveryV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/"+url.QueryEscape(id)+"/deliveries", queryParams, nil)
}

// GET GetDeliveryPackage V1
// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersGetDeliveryPackageV1(id string, no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v1/deliveries/"+url.QueryEscape(id)+"/packages", queryParams, nil)
}

// GET GetDeliveryPackage V2
// Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (c *MetrcClient) TransfersGetDeliveryPackageV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/deliveries/"+url.QueryEscape(id)+"/packages", queryParams, nil)
}

// GET GetDeliveryPackageRequiredlabtestbatches V1
// Please note: The {id} parameter above represents a Transfer Delivery Package ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersGetDeliveryPackageRequiredlabtestbatchesV1(id string, no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v1/deliveries/package/"+url.QueryEscape(id)+"/requiredlabtestbatches", queryParams, nil)
}

// GET GetDeliveryPackageRequiredlabtestbatches V2
// Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (c *MetrcClient) TransfersGetDeliveryPackageRequiredlabtestbatchesV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/deliveries/package/"+url.QueryEscape(id)+"/requiredlabtestbatches", queryParams, nil)
}

// GET GetDeliveryPackageWholesale V1
// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersGetDeliveryPackageWholesaleV1(id string, no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v1/deliveries/"+url.QueryEscape(id)+"/packages/wholesale", queryParams, nil)
}

// GET GetDeliveryPackageWholesale V2
// Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (c *MetrcClient) TransfersGetDeliveryPackageWholesaleV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/deliveries/"+url.QueryEscape(id)+"/packages/wholesale", queryParams, nil)
}

// GET GetDeliveryTransporters V1
// Please note: that the {id} parameter above represents a Shipment Delivery ID.
// 
//   Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersGetDeliveryTransportersV1(id string, no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v1/deliveries/"+url.QueryEscape(id)+"/transporters", queryParams, nil)
}

// GET GetDeliveryTransporters V2
// Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (c *MetrcClient) TransfersGetDeliveryTransportersV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/deliveries/"+url.QueryEscape(id)+"/transporters", queryParams, nil)
}

// GET GetDeliveryTransportersDetails V1
// Please note: The {id} parameter above represents a Shipment Delivery ID.
// 
//   Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersGetDeliveryTransportersDetailsV1(id string, no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v1/deliveries/"+url.QueryEscape(id)+"/transporters/details", queryParams, nil)
}

// GET GetDeliveryTransportersDetails V2
// Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (c *MetrcClient) TransfersGetDeliveryTransportersDetailsV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/deliveries/"+url.QueryEscape(id)+"/transporters/details", queryParams, nil)
}

// GET GetHub V2
// Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (c *MetrcClient) TransfersGetHubV2(estimatedarrivalend string, estimatedarrivalstart string, lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if estimatedarrivalend != "" {
		queryParams["estimatedArrivalEnd"] = estimatedarrivalend
	}
	if estimatedarrivalstart != "" {
		queryParams["estimatedArrivalStart"] = estimatedarrivalstart
	}
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/hub", queryParams, nil)
}

// GET GetIncoming V1
// Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersGetIncomingV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/transfers/v1/incoming", queryParams, nil)
}

// GET GetIncoming V2
// Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (c *MetrcClient) TransfersGetIncomingV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/incoming", queryParams, nil)
}

// GET GetOutgoing V1
// Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersGetOutgoingV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/transfers/v1/outgoing", queryParams, nil)
}

// GET GetOutgoing V2
// Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (c *MetrcClient) TransfersGetOutgoingV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/outgoing", queryParams, nil)
}

// GET GetRejected V1
// Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersGetRejectedV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/transfers/v1/rejected", queryParams, nil)
}

// GET GetRejected V2
// Retrieves a list of shipments with rejected packages for a Facility.
// 
//   Permissions Required:
//   - Manage Transfers
//   - View Transfers
func (c *MetrcClient) TransfersGetRejectedV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/rejected", queryParams, nil)
}

// GET GetTemplates V1
// Permissions Required:
//   - Transfer Templates
func (c *MetrcClient) TransfersGetTemplatesV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/transfers/v1/templates", queryParams, nil)
}

// GET GetTemplatesDelivery V1
// Permissions Required:
//   - Transfer Templates
func (c *MetrcClient) TransfersGetTemplatesDeliveryV1(id string, no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v1/templates/"+url.QueryEscape(id)+"/deliveries", queryParams, nil)
}

// GET GetTemplatesDeliveryPackage V1
// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersGetTemplatesDeliveryPackageV1(id string, no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v1/templates/deliveries/"+url.QueryEscape(id)+"/packages", queryParams, nil)
}

// GET GetTemplatesDeliveryTransporters V1
// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfer Templates
func (c *MetrcClient) TransfersGetTemplatesDeliveryTransportersV1(id string, no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v1/templates/deliveries/"+url.QueryEscape(id)+"/transporters", queryParams, nil)
}

// GET GetTemplatesDeliveryTransportersDetails V1
// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
// 
//   Permissions Required:
//   - Transfer Templates
func (c *MetrcClient) TransfersGetTemplatesDeliveryTransportersDetailsV1(id string, no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v1/templates/deliveries/"+url.QueryEscape(id)+"/transporters/details", queryParams, nil)
}

// GET GetTemplatesOutgoing V2
// Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
// 
//   Permissions Required:
//   - Manage Transfer Templates
//   - View Transfer Templates
func (c *MetrcClient) TransfersGetTemplatesOutgoingV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/templates/outgoing", queryParams, nil)
}

// GET GetTemplatesOutgoingDelivery V2
// Retrieves a list of deliveries associated with a specific transfer template.
// 
//   Permissions Required:
//   - Manage Transfer Templates
//   - View Transfer Templates
func (c *MetrcClient) TransfersGetTemplatesOutgoingDeliveryV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/templates/outgoing/"+url.QueryEscape(id)+"/deliveries", queryParams, nil)
}

// GET GetTemplatesOutgoingDeliveryPackage V2
// Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfer Templates
//   - View Transfer Templates
func (c *MetrcClient) TransfersGetTemplatesOutgoingDeliveryPackageV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/templates/outgoing/deliveries/"+url.QueryEscape(id)+"/packages", queryParams, nil)
}

// GET GetTemplatesOutgoingDeliveryTransporters V2
// Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfer Templates
//   - View Transfer Templates
func (c *MetrcClient) TransfersGetTemplatesOutgoingDeliveryTransportersV2(id string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/templates/outgoing/deliveries/"+url.QueryEscape(id)+"/transporters", queryParams, nil)
}

// GET GetTemplatesOutgoingDeliveryTransportersDetails V2
// Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
// 
//   Permissions Required:
//   - Manage Transfer Templates
//   - View Transfer Templates
func (c *MetrcClient) TransfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id string, no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/transfers/v2/templates/outgoing/deliveries/"+url.QueryEscape(id)+"/transporters/details", queryParams, nil)
}

// GET GetTypes V1
// Permissions Required:
//   - None
func (c *MetrcClient) TransfersGetTypesV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/transfers/v1/types", queryParams, nil)
}

// GET GetTypes V2
// Retrieves a list of available transfer types for a Facility based on its license number.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) TransfersGetTypesV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transfers/v2/types", queryParams, nil)
}

// PUT UpdateExternalIncoming V1
// Permissions Required:
//   - Transfers
func (c *MetrcClient) TransfersUpdateExternalIncomingV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/transfers/v1/external/incoming", queryParams, body)
}

// PUT UpdateExternalIncoming V2
// Updates external incoming shipment plans for a Facility.
// 
//   Permissions Required:
//   - Manage Transfers
func (c *MetrcClient) TransfersUpdateExternalIncomingV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/transfers/v2/external/incoming", queryParams, body)
}

// PUT UpdateTemplates V1
// Permissions Required:
//   - Transfer Templates
func (c *MetrcClient) TransfersUpdateTemplatesV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/transfers/v1/templates", queryParams, body)
}

// PUT UpdateTemplatesOutgoing V2
// Updates existing transfer templates for a Facility.
// 
//   Permissions Required:
//   - Manage Transfer Templates
func (c *MetrcClient) TransfersUpdateTemplatesOutgoingV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/transfers/v2/templates/outgoing", queryParams, body)
}

// POST CreateAdditives V1
// Permissions Required:
//   - Manage Plants Additives
func (c *MetrcClient) PlantsCreateAdditivesV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v1/additives", queryParams, body)
}

// POST CreateAdditives V2
// Records additive usage details applied to specific plants at a Facility.
// 
//   Permissions Required:
//   - Manage Plants Additives
func (c *MetrcClient) PlantsCreateAdditivesV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v2/additives", queryParams, body)
}

// POST CreateAdditivesBylocation V1
// Permissions Required:
//   - Manage Plants
//   - Manage Plants Additives
func (c *MetrcClient) PlantsCreateAdditivesBylocationV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v1/additives/bylocation", queryParams, body)
}

// POST CreateAdditivesBylocation V2
// Records additive usage for plants based on their location within a specified Facility.
// 
//   Permissions Required:
//   - Manage Plants
//   - Manage Plants Additives
func (c *MetrcClient) PlantsCreateAdditivesBylocationV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v2/additives/bylocation", queryParams, body)
}

// POST CreateAdditivesBylocationUsingtemplate V2
// Records additive usage for plants by location using a predefined additive template at a specified Facility.
// 
//   Permissions Required:
//   - Manage Plants Additives
func (c *MetrcClient) PlantsCreateAdditivesBylocationUsingtemplateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v2/additives/bylocation/usingtemplate", queryParams, body)
}

// POST CreateAdditivesUsingtemplate V2
// Records additive usage for plants using predefined additive templates at a specified Facility.
// 
//   Permissions Required:
//   - Manage Plants Additives
func (c *MetrcClient) PlantsCreateAdditivesUsingtemplateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v2/additives/usingtemplate", queryParams, body)
}

// POST CreateChangegrowthphases V1
// Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantsCreateChangegrowthphasesV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v1/changegrowthphases", queryParams, body)
}

// POST CreateHarvestplants V1
// NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manicure/Harvest Veg/Flower Plants
func (c *MetrcClient) PlantsCreateHarvestplantsV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v1/harvestplants", queryParams, body)
}

// POST CreateManicure V2
// Creates harvest product records from plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manicure/Harvest Veg/Flower Plants
func (c *MetrcClient) PlantsCreateManicureV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v2/manicure", queryParams, body)
}

// POST CreateManicureplants V1
// Permissions Required:
//   - View Veg/Flower Plants
//   - Manicure/Harvest Veg/Flower Plants
func (c *MetrcClient) PlantsCreateManicureplantsV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v1/manicureplants", queryParams, body)
}

// POST CreateMoveplants V1
// Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantsCreateMoveplantsV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v1/moveplants", queryParams, body)
}

// POST CreatePlantbatchPackage V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PlantsCreatePlantbatchPackageV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v1/create/plantbatch/packages", queryParams, body)
}

// POST CreatePlantbatchPackage V2
// Creates packages from plant batches at a specified Facility.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PlantsCreatePlantbatchPackageV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v2/plantbatch/packages", queryParams, body)
}

// POST CreatePlantings V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantsCreatePlantingsV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v1/create/plantings", queryParams, body)
}

// POST CreatePlantings V2
// Creates new plant batches at a specified Facility from existing plant data.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants Inventory
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantsCreatePlantingsV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v2/plantings", queryParams, body)
}

// POST CreateWaste V1
// Permissions Required:
//   - Manage Plants Waste
func (c *MetrcClient) PlantsCreateWasteV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v1/waste", queryParams, body)
}

// POST CreateWaste V2
// Records waste events for plants at a Facility, including method, reason, and location details.
// 
//   Permissions Required:
//   - Manage Plants Waste
func (c *MetrcClient) PlantsCreateWasteV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/plants/v2/waste", queryParams, body)
}

// DELETE Delete V1
// Permissions Required:
//   - View Veg/Flower Plants
//   - Destroy Veg/Flower Plants
func (c *MetrcClient) PlantsDeleteV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/plants/v1", queryParams, nil)
}

// DELETE Delete V2
// Removes plants from a Facility’s inventory while recording the reason for their disposal.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Destroy Veg/Flower Plants
func (c *MetrcClient) PlantsDeleteV2(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/plants/v2", queryParams, nil)
}

// GET Get V1
// Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V2
// Retrieves a Plant by Id.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetAdditives V1
// Permissions Required:
//   - View/Manage Plants Additives
func (c *MetrcClient) PlantsGetAdditivesV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v1/additives", queryParams, nil)
}

// GET GetAdditives V2
// Retrieves additive records applied to plants at a specified Facility.
// 
//   Permissions Required:
//   - View/Manage Plants Additives
func (c *MetrcClient) PlantsGetAdditivesV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/additives", queryParams, nil)
}

// GET GetAdditivesTypes V1
// Permissions Required:
//   -
func (c *MetrcClient) PlantsGetAdditivesTypesV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/plants/v1/additives/types", queryParams, nil)
}

// GET GetAdditivesTypes V2
// Retrieves a list of all plant additive types defined within a Facility.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) PlantsGetAdditivesTypesV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/plants/v2/additives/types", queryParams, nil)
}

// GET GetByLabel V1
// Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetByLabelV1(label string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v1/"+url.QueryEscape(label)+"", queryParams, nil)
}

// GET GetByLabel V2
// Retrieves a Plant by label.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetByLabelV2(label string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v2/"+url.QueryEscape(label)+"", queryParams, nil)
}

// GET GetFlowering V1
// Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetFloweringV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v1/flowering", queryParams, nil)
}

// GET GetFlowering V2
// Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetFloweringV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/flowering", queryParams, nil)
}

// GET GetGrowthPhases V1
// Permissions Required:
//   - None
func (c *MetrcClient) PlantsGetGrowthPhasesV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v1/growthphases", queryParams, nil)
}

// GET GetGrowthPhases V2
// Retrieves the list of growth phases supported by a specified Facility.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) PlantsGetGrowthPhasesV2(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v2/growthphases", queryParams, nil)
}

// GET GetInactive V1
// Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v1/inactive", queryParams, nil)
}

// GET GetInactive V2
// Retrieves inactive plants at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/inactive", queryParams, nil)
}

// GET GetMother V2
// Retrieves mother-phase plants at a specified Facility.
// 
//   Permissions Required:
//   - View Mother Plants
func (c *MetrcClient) PlantsGetMotherV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/mother", queryParams, nil)
}

// GET GetMotherInactive V2
// Retrieves inactive mother-phase plants at a specified Facility.
// 
//   Permissions Required:
//   - View Mother Plants
func (c *MetrcClient) PlantsGetMotherInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/mother/inactive", queryParams, nil)
}

// GET GetMotherOnhold V2
// Retrieves mother-phase plants currently marked as on hold at a specified Facility.
// 
//   Permissions Required:
//   - View Mother Plants
func (c *MetrcClient) PlantsGetMotherOnholdV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/mother/onhold", queryParams, nil)
}

// GET GetOnhold V1
// Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetOnholdV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v1/onhold", queryParams, nil)
}

// GET GetOnhold V2
// Retrieves plants that are currently on hold at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetOnholdV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/onhold", queryParams, nil)
}

// GET GetVegetative V1
// Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetVegetativeV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v1/vegetative", queryParams, nil)
}

// GET GetVegetative V2
// Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
func (c *MetrcClient) PlantsGetVegetativeV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/vegetative", queryParams, nil)
}

// GET GetWaste V2
// Retrieves a list of recorded plant waste events for a specific Facility.
// 
//   Permissions Required:
//   - View Plants Waste
func (c *MetrcClient) PlantsGetWasteV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/waste", queryParams, nil)
}

// GET GetWasteMethodsAll V1
// Permissions Required:
//   - None
func (c *MetrcClient) PlantsGetWasteMethodsAllV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/plants/v1/waste/methods/all", queryParams, nil)
}

// GET GetWasteMethodsAll V2
// Retrieves a list of all available plant waste methods for use within a Facility.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) PlantsGetWasteMethodsAllV2(pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/waste/methods/all", queryParams, nil)
}

// GET GetWastePackage V2
// Retrieves a list of package records linked to the specified plantWasteId for a given facility.
// 
//   Permissions Required:
//   - View Plants Waste
func (c *MetrcClient) PlantsGetWastePackageV2(id string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/waste/"+url.QueryEscape(id)+"/package", queryParams, nil)
}

// GET GetWastePlant V2
// Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
// 
//   Permissions Required:
//   - View Plants Waste
func (c *MetrcClient) PlantsGetWastePlantV2(id string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/waste/"+url.QueryEscape(id)+"/plant", queryParams, nil)
}

// GET GetWasteReasons V1
// Permissions Required:
//   - None
func (c *MetrcClient) PlantsGetWasteReasonsV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/plants/v1/waste/reasons", queryParams, nil)
}

// GET GetWasteReasons V2
// Retriveves available reasons for recording mature plant waste at a specified Facility.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) PlantsGetWasteReasonsV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/plants/v2/waste/reasons", queryParams, nil)
}

// PUT UpdateAdjust V2
// Adjusts the recorded count of plants at a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantsUpdateAdjustV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plants/v2/adjust", queryParams, body)
}

// PUT UpdateGrowthphase V2
// Changes the growth phases of plants within a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantsUpdateGrowthphaseV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plants/v2/growthphase", queryParams, body)
}

// PUT UpdateHarvest V2
// Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manicure/Harvest Veg/Flower Plants
func (c *MetrcClient) PlantsUpdateHarvestV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plants/v2/harvest", queryParams, body)
}

// PUT UpdateLocation V2
// Moves plant batches to new locations within a specified Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantsUpdateLocationV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plants/v2/location", queryParams, body)
}

// PUT UpdateMerge V2
// Merges multiple plant groups into a single group within a Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manicure/Harvest Veg/Flower Plants
func (c *MetrcClient) PlantsUpdateMergeV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plants/v2/merge", queryParams, body)
}

// PUT UpdateSplit V2
// Splits an existing plant group into multiple groups within a Facility.
// 
//   Permissions Required:
//   - View Plant
func (c *MetrcClient) PlantsUpdateSplitV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plants/v2/split", queryParams, body)
}

// PUT UpdateStrain V2
// Updates the strain information for plants within a Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantsUpdateStrainV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plants/v2/strain", queryParams, body)
}

// PUT UpdateTag V2
// Replaces existing plant tags with new tags for plants within a Facility.
// 
//   Permissions Required:
//   - View Veg/Flower Plants
//   - Manage Veg/Flower Plants Inventory
func (c *MetrcClient) PlantsUpdateTagV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/plants/v2/tag", queryParams, body)
}

// POST Create V2
// Creates new additive templates for a specified Facility.
// 
//   Permissions Required:
//   - Manage Additives
func (c *MetrcClient) AdditivesTemplatesCreateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/additivestemplates/v2", queryParams, body)
}

// GET Get V2
// Retrieves an Additive Template by its Id.
// 
//   Permissions Required:
//   - Manage Additives
func (c *MetrcClient) AdditivesTemplatesGetV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/additivestemplates/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetActive V2
// Retrieves a list of active additive templates for a specified Facility.
// 
//   Permissions Required:
//   - Manage Additives
func (c *MetrcClient) AdditivesTemplatesGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/additivestemplates/v2/active", queryParams, nil)
}

// GET GetInactive V2
// Retrieves a list of inactive additive templates for a specified Facility.
// 
//   Permissions Required:
//   - Manage Additives
func (c *MetrcClient) AdditivesTemplatesGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/additivestemplates/v2/inactive", queryParams, nil)
}

// PUT Update V2
// Updates existing additive templates for a specified Facility.
// 
//   Permissions Required:
//   - Manage Additives
func (c *MetrcClient) AdditivesTemplatesUpdateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/additivestemplates/v2", queryParams, body)
}

// GET GetByCaregiverLicenseNumber V1
// Data returned by this endpoint is cached for up to one minute.
// 
//   Permissions Required:
//   - Lookup Caregivers
func (c *MetrcClient) CaregiversStatusGetByCaregiverLicenseNumberV1(caregiverlicensenumber string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/caregivers/v1/status/"+url.QueryEscape(caregiverlicensenumber)+"", queryParams, nil)
}

// GET GetByCaregiverLicenseNumber V2
// Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
// 
//   Permissions Required:
//   - Lookup Caregivers
func (c *MetrcClient) CaregiversStatusGetByCaregiverLicenseNumberV2(caregiverlicensenumber string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/caregivers/v2/status/"+url.QueryEscape(caregiverlicensenumber)+"", queryParams, nil)
}

// POST CreateFinish V1
// Permissions Required:
//   - View Harvests
//   - Finish/Discontinue Harvests
func (c *MetrcClient) HarvestsCreateFinishV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/harvests/v1/finish", queryParams, body)
}

// POST CreatePackage V1
// Permissions Required:
//   - View Harvests
//   - Manage Harvests
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) HarvestsCreatePackageV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/harvests/v1/create/packages", queryParams, body)
}

// POST CreatePackage V2
// Creates packages from harvested products for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Manage Harvests
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) HarvestsCreatePackageV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/harvests/v2/packages", queryParams, body)
}

// POST CreatePackageTesting V1
// Permissions Required:
//   - View Harvests
//   - Manage Harvests
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) HarvestsCreatePackageTestingV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/harvests/v1/create/packages/testing", queryParams, body)
}

// POST CreatePackageTesting V2
// Creates packages for testing from harvested products for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Manage Harvests
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) HarvestsCreatePackageTestingV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/harvests/v2/packages/testing", queryParams, body)
}

// POST CreateRemoveWaste V1
// Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (c *MetrcClient) HarvestsCreateRemoveWasteV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/harvests/v1/removewaste", queryParams, body)
}

// POST CreateUnfinish V1
// Permissions Required:
//   - View Harvests
//   - Finish/Discontinue Harvests
func (c *MetrcClient) HarvestsCreateUnfinishV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/harvests/v1/unfinish", queryParams, body)
}

// POST CreateWaste V2
// Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
// 
//   Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (c *MetrcClient) HarvestsCreateWasteV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/harvests/v2/waste", queryParams, body)
}

// DELETE DeleteWaste V2
// Discontinues a specific harvest waste record by Id for the specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Discontinue Harvest Waste
func (c *MetrcClient) HarvestsDeleteWasteV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/harvests/v2/waste/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V1
// Permissions Required:
//   - View Harvests
func (c *MetrcClient) HarvestsGetV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/harvests/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V2
// Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
// 
//   Permissions Required:
//   - View Harvests
func (c *MetrcClient) HarvestsGetV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/harvests/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetActive V1
// Permissions Required:
//   - View Harvests
func (c *MetrcClient) HarvestsGetActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/harvests/v1/active", queryParams, nil)
}

// GET GetActive V2
// Retrieves a list of active harvests for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
func (c *MetrcClient) HarvestsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/harvests/v2/active", queryParams, nil)
}

// GET GetInactive V1
// Permissions Required:
//   - View Harvests
func (c *MetrcClient) HarvestsGetInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/harvests/v1/inactive", queryParams, nil)
}

// GET GetInactive V2
// Retrieves a list of inactive harvests for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
func (c *MetrcClient) HarvestsGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/harvests/v2/inactive", queryParams, nil)
}

// GET GetOnhold V1
// Permissions Required:
//   - View Harvests
func (c *MetrcClient) HarvestsGetOnholdV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/harvests/v1/onhold", queryParams, nil)
}

// GET GetOnhold V2
// Retrieves a list of harvests on hold for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
func (c *MetrcClient) HarvestsGetOnholdV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/harvests/v2/onhold", queryParams, nil)
}

// GET GetWaste V2
// Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
// 
//   Permissions Required:
//   - View Harvests
func (c *MetrcClient) HarvestsGetWasteV2(harvestid string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if harvestid != "" {
		queryParams["harvestId"] = harvestid
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/harvests/v2/waste", queryParams, nil)
}

// GET GetWasteTypes V1
// Permissions Required:
//   - None
func (c *MetrcClient) HarvestsGetWasteTypesV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/harvests/v1/waste/types", queryParams, nil)
}

// GET GetWasteTypes V2
// Retrieves a list of Waste types for harvests.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) HarvestsGetWasteTypesV2(pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/harvests/v2/waste/types", queryParams, nil)
}

// PUT UpdateFinish V2
// Marks one or more harvests as finished for the specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Finish/Discontinue Harvests
func (c *MetrcClient) HarvestsUpdateFinishV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/harvests/v2/finish", queryParams, body)
}

// PUT UpdateLocation V2
// Updates the Location of Harvest for a specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (c *MetrcClient) HarvestsUpdateLocationV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/harvests/v2/location", queryParams, body)
}

// PUT UpdateMove V1
// Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (c *MetrcClient) HarvestsUpdateMoveV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/harvests/v1/move", queryParams, body)
}

// PUT UpdateRename V1
// Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (c *MetrcClient) HarvestsUpdateRenameV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/harvests/v1/rename", queryParams, body)
}

// PUT UpdateRename V2
// Renames one or more harvests for the specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Manage Harvests
func (c *MetrcClient) HarvestsUpdateRenameV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/harvests/v2/rename", queryParams, body)
}

// PUT UpdateRestoreHarvestedPlants V2
// Restores previously harvested plants to their original state for the specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Finish/Discontinue Harvests
func (c *MetrcClient) HarvestsUpdateRestoreHarvestedPlantsV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/harvests/v2/restore/harvestedplants", queryParams, body)
}

// PUT UpdateUnfinish V2
// Reopens one or more previously finished harvests for the specified Facility.
// 
//   Permissions Required:
//   - View Harvests
//   - Finish/Discontinue Harvests
func (c *MetrcClient) HarvestsUpdateUnfinishV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/harvests/v2/unfinish", queryParams, body)
}

// POST CreateRecord V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) LabTestsCreateRecordV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/labtests/v1/record", queryParams, body)
}

// POST CreateRecord V2
// Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) LabTestsCreateRecordV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/labtests/v2/record", queryParams, body)
}

// GET GetBatches V2
// Retrieves a list of Lab Test batches.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) LabTestsGetBatchesV2(pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/labtests/v2/batches", queryParams, nil)
}

// GET GetLabtestdocument V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) LabTestsGetLabtestdocumentV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/labtests/v1/labtestdocument/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetLabtestdocument V2
// Retrieves a specific Lab Test result document by its Id for a given Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) LabTestsGetLabtestdocumentV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/labtests/v2/labtestdocument/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetResults V1
// Permissions Required:
//   - View Packages
func (c *MetrcClient) LabTestsGetResultsV1(licensenumber string, packageid string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if packageid != "" {
		queryParams["packageId"] = packageid
	}
	return c.send("GET", "/labtests/v1/results", queryParams, nil)
}

// GET GetResults V2
// Retrieves Lab Test results for a specified Package.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) LabTestsGetResultsV2(licensenumber string, packageid string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if packageid != "" {
		queryParams["packageId"] = packageid
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/labtests/v2/results", queryParams, nil)
}

// GET GetStates V1
// Permissions Required:
//   - None
func (c *MetrcClient) LabTestsGetStatesV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/labtests/v1/states", queryParams, nil)
}

// GET GetStates V2
// Returns a list of all lab testing states.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) LabTestsGetStatesV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/labtests/v2/states", queryParams, nil)
}

// GET GetTypes V1
// Permissions Required:
//   - None
func (c *MetrcClient) LabTestsGetTypesV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/labtests/v1/types", queryParams, nil)
}

// GET GetTypes V2
// Returns a list of Lab Test types.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) LabTestsGetTypesV2(pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/labtests/v2/types", queryParams, nil)
}

// PUT UpdateLabtestdocument V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) LabTestsUpdateLabtestdocumentV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/labtests/v1/labtestdocument", queryParams, body)
}

// PUT UpdateLabtestdocument V2
// Updates one or more documents for previously submitted lab tests.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) LabTestsUpdateLabtestdocumentV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/labtests/v2/labtestdocument", queryParams, body)
}

// PUT UpdateResultRelease V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) LabTestsUpdateResultReleaseV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/labtests/v1/results/release", queryParams, body)
}

// PUT UpdateResultRelease V2
// Releases Lab Test results for one or more packages.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) LabTestsUpdateResultReleaseV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/labtests/v2/results/release", queryParams, body)
}

// POST CreateDelivery V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Sales Delivery
func (c *MetrcClient) SalesCreateDeliveryV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v1/deliveries", queryParams, body)
}

// POST CreateDelivery V2
// Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
func (c *MetrcClient) SalesCreateDeliveryV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v2/deliveries", queryParams, body)
}

// POST CreateDeliveryRetailer V1
// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Retailer Delivery
func (c *MetrcClient) SalesCreateDeliveryRetailerV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v1/deliveries/retailer", queryParams, body)
}

// POST CreateDeliveryRetailer V2
// Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
//   - Manage Retailer Delivery
func (c *MetrcClient) SalesCreateDeliveryRetailerV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v2/deliveries/retailer", queryParams, body)
}

// POST CreateDeliveryRetailerDepart V1
// Permissions Required:
//   - Retailer Delivery
func (c *MetrcClient) SalesCreateDeliveryRetailerDepartV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v1/deliveries/retailer/depart", queryParams, body)
}

// POST CreateDeliveryRetailerDepart V2
// Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - Manage Retailer Delivery
func (c *MetrcClient) SalesCreateDeliveryRetailerDepartV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v2/deliveries/retailer/depart", queryParams, body)
}

// POST CreateDeliveryRetailerEnd V1
// Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Retailer Delivery
func (c *MetrcClient) SalesCreateDeliveryRetailerEndV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v1/deliveries/retailer/end", queryParams, body)
}

// POST CreateDeliveryRetailerEnd V2
// Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - Manage Retailer Delivery
func (c *MetrcClient) SalesCreateDeliveryRetailerEndV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v2/deliveries/retailer/end", queryParams, body)
}

// POST CreateDeliveryRetailerRestock V1
// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Retailer Delivery
func (c *MetrcClient) SalesCreateDeliveryRetailerRestockV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v1/deliveries/retailer/restock", queryParams, body)
}

// POST CreateDeliveryRetailerRestock V2
// Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - Manage Retailer Delivery
func (c *MetrcClient) SalesCreateDeliveryRetailerRestockV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v2/deliveries/retailer/restock", queryParams, body)
}

// POST CreateDeliveryRetailerSale V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Retailer Delivery
func (c *MetrcClient) SalesCreateDeliveryRetailerSaleV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v1/deliveries/retailer/sale", queryParams, body)
}

// POST CreateDeliveryRetailerSale V2
// Records sales deliveries originating from a retailer delivery for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
func (c *MetrcClient) SalesCreateDeliveryRetailerSaleV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v2/deliveries/retailer/sale", queryParams, body)
}

// POST CreateReceipt V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Sales
func (c *MetrcClient) SalesCreateReceiptV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v1/receipts", queryParams, body)
}

// POST CreateReceipt V2
// Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales (Write)
//   - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
//   - Industry/Facility Type/Advanced Sales
//   - WebApi Sales Read Write State (All or WriteOnly)
//   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
func (c *MetrcClient) SalesCreateReceiptV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v2/receipts", queryParams, body)
}

// POST CreateTransactionByDate V1
// Permissions Required:
//   - Sales
func (c *MetrcClient) SalesCreateTransactionByDateV1(date string, licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/sales/v1/transactions/"+url.QueryEscape(date)+"", queryParams, body)
}

// DELETE DeleteDelivery V1
// Permissions Required:
//   - Sales Delivery
func (c *MetrcClient) SalesDeleteDeliveryV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/sales/v1/deliveries/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteDelivery V2
// Voids a sales delivery for a Facility using the provided License Number and delivery Id.
// 
//   Permissions Required:
//   - Manage Sales Delivery
func (c *MetrcClient) SalesDeleteDeliveryV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/sales/v2/deliveries/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteDeliveryRetailer V1
// Permissions Required:
//   - Retailer Delivery
func (c *MetrcClient) SalesDeleteDeliveryRetailerV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/sales/v1/deliveries/retailer/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteDeliveryRetailer V2
// Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - Manage Retailer Delivery
func (c *MetrcClient) SalesDeleteDeliveryRetailerV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/sales/v2/deliveries/retailer/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteReceipt V1
// Permissions Required:
//   - Sales
func (c *MetrcClient) SalesDeleteReceiptV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/sales/v1/receipts/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteReceipt V2
// Archives a sales receipt for a Facility using the provided License Number and receipt Id.
// 
//   Permissions Required:
//   - Manage Sales
func (c *MetrcClient) SalesDeleteReceiptV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/sales/v2/receipts/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetCounties V1
// Permissions Required:
//   - None
func (c *MetrcClient) SalesGetCountiesV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/sales/v1/counties", queryParams, nil)
}

// GET GetCounties V2
// Returns a list of counties available for sales deliveries.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) SalesGetCountiesV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/sales/v2/counties", queryParams, nil)
}

// GET GetCustomertypes V1
// Permissions Required:
//   - None
func (c *MetrcClient) SalesGetCustomertypesV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/sales/v1/customertypes", queryParams, nil)
}

// GET GetCustomertypes V2
// Returns a list of customer types.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) SalesGetCustomertypesV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/sales/v2/customertypes", queryParams, nil)
}

// GET GetDeliveriesActive V1
// Permissions Required:
//   - Sales Delivery
func (c *MetrcClient) SalesGetDeliveriesActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string, salesdateend string, salesdatestart string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if salesdateend != "" {
		queryParams["salesDateEnd"] = salesdateend
	}
	if salesdatestart != "" {
		queryParams["salesDateStart"] = salesdatestart
	}
	return c.send("GET", "/sales/v1/deliveries/active", queryParams, nil)
}

// GET GetDeliveriesActive V2
// Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
// 
//   Permissions Required:
//   - View Sales Delivery
//   - Manage Sales Delivery
func (c *MetrcClient) SalesGetDeliveriesActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string, salesdateend string, salesdatestart string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	if salesdateend != "" {
		queryParams["salesDateEnd"] = salesdateend
	}
	if salesdatestart != "" {
		queryParams["salesDateStart"] = salesdatestart
	}
	return c.send("GET", "/sales/v2/deliveries/active", queryParams, nil)
}

// GET GetDeliveriesInactive V1
// Permissions Required:
//   - Sales Delivery
func (c *MetrcClient) SalesGetDeliveriesInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string, salesdateend string, salesdatestart string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if salesdateend != "" {
		queryParams["salesDateEnd"] = salesdateend
	}
	if salesdatestart != "" {
		queryParams["salesDateStart"] = salesdatestart
	}
	return c.send("GET", "/sales/v1/deliveries/inactive", queryParams, nil)
}

// GET GetDeliveriesInactive V2
// Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
// 
//   Permissions Required:
//   - View Sales Delivery
//   - Manage Sales Delivery
func (c *MetrcClient) SalesGetDeliveriesInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string, salesdateend string, salesdatestart string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	if salesdateend != "" {
		queryParams["salesDateEnd"] = salesdateend
	}
	if salesdatestart != "" {
		queryParams["salesDateStart"] = salesdatestart
	}
	return c.send("GET", "/sales/v2/deliveries/inactive", queryParams, nil)
}

// GET GetDeliveriesRetailerActive V1
// Permissions Required:
//   - Retailer Delivery
func (c *MetrcClient) SalesGetDeliveriesRetailerActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v1/deliveries/retailer/active", queryParams, nil)
}

// GET GetDeliveriesRetailerActive V2
// Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
// 
//   Permissions Required:
//   - View Retailer Delivery
//   - Manage Retailer Delivery
func (c *MetrcClient) SalesGetDeliveriesRetailerActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/sales/v2/deliveries/retailer/active", queryParams, nil)
}

// GET GetDeliveriesRetailerInactive V1
// Permissions Required:
//   - Retailer Delivery
func (c *MetrcClient) SalesGetDeliveriesRetailerInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v1/deliveries/retailer/inactive", queryParams, nil)
}

// GET GetDeliveriesRetailerInactive V2
// Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
// 
//   Permissions Required:
//   - View Retailer Delivery
//   - Manage Retailer Delivery
func (c *MetrcClient) SalesGetDeliveriesRetailerInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/sales/v2/deliveries/retailer/inactive", queryParams, nil)
}

// GET GetDeliveriesReturnreasons V1
// Permissions Required:
//   -
func (c *MetrcClient) SalesGetDeliveriesReturnreasonsV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v1/deliveries/returnreasons", queryParams, nil)
}

// GET GetDeliveriesReturnreasons V2
// Returns a list of return reasons for sales deliveries based on the provided License Number.
// 
//   Permissions Required:
//   - Sales Delivery
func (c *MetrcClient) SalesGetDeliveriesReturnreasonsV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/sales/v2/deliveries/returnreasons", queryParams, nil)
}

// GET GetDelivery V1
// Permissions Required:
//   - Sales Delivery
func (c *MetrcClient) SalesGetDeliveryV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v1/deliveries/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetDelivery V2
// Retrieves a sales delivery record by its Id, with an optional License Number.
// 
//   Permissions Required:
//   - View Sales Delivery
//   - Manage Sales Delivery
func (c *MetrcClient) SalesGetDeliveryV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v2/deliveries/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetDeliveryRetailer V1
// Permissions Required:
//   - Retailer Delivery
func (c *MetrcClient) SalesGetDeliveryRetailerV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v1/deliveries/retailer/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetDeliveryRetailer V2
// Retrieves a retailer delivery record by its ID, with an optional License Number.
// 
//   Permissions Required:
//   - View Retailer Delivery
//   - Manage Retailer Delivery
func (c *MetrcClient) SalesGetDeliveryRetailerV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v2/deliveries/retailer/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetPatientRegistrationsLocations V1
// Permissions Required:
//   -
func (c *MetrcClient) SalesGetPatientRegistrationsLocationsV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/sales/v1/patientregistration/locations", queryParams, nil)
}

// GET GetPatientRegistrationsLocations V2
// Returns a list of valid Patient registration locations for sales.
// 
//   Permissions Required:
//   -
func (c *MetrcClient) SalesGetPatientRegistrationsLocationsV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/sales/v2/patientregistration/locations", queryParams, nil)
}

// GET GetPaymenttypes V1
// Permissions Required:
//   - Sales Delivery
func (c *MetrcClient) SalesGetPaymenttypesV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v1/paymenttypes", queryParams, nil)
}

// GET GetPaymenttypes V2
// Returns a list of available payment types for the specified License Number.
// 
//   Permissions Required:
//   - View Sales Delivery
//   - Manage Sales Delivery
func (c *MetrcClient) SalesGetPaymenttypesV2(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v2/paymenttypes", queryParams, nil)
}

// GET GetReceipt V1
// Permissions Required:
//   - Sales
func (c *MetrcClient) SalesGetReceiptV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v1/receipts/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetReceipt V2
// Retrieves a sales receipt by its Id, with an optional License Number.
// 
//   Permissions Required:
//   - View Sales
//   - Manage Sales
func (c *MetrcClient) SalesGetReceiptV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v2/receipts/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetReceiptsActive V1
// Permissions Required:
//   - Sales
func (c *MetrcClient) SalesGetReceiptsActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string, salesdateend string, salesdatestart string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if salesdateend != "" {
		queryParams["salesDateEnd"] = salesdateend
	}
	if salesdatestart != "" {
		queryParams["salesDateStart"] = salesdatestart
	}
	return c.send("GET", "/sales/v1/receipts/active", queryParams, nil)
}

// GET GetReceiptsActive V2
// Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
// 
//   Permissions Required:
//   - View Sales
//   - Manage Sales
func (c *MetrcClient) SalesGetReceiptsActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string, salesdateend string, salesdatestart string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	if salesdateend != "" {
		queryParams["salesDateEnd"] = salesdateend
	}
	if salesdatestart != "" {
		queryParams["salesDateStart"] = salesdatestart
	}
	return c.send("GET", "/sales/v2/receipts/active", queryParams, nil)
}

// GET GetReceiptsExternalByExternalNumber V2
// Retrieves a Sales Receipt by its external number, with an optional License Number.
// 
//   Permissions Required:
//   - View Sales
//   - Manage Sales
func (c *MetrcClient) SalesGetReceiptsExternalByExternalNumberV2(externalnumber string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v2/receipts/external/"+url.QueryEscape(externalnumber)+"", queryParams, nil)
}

// GET GetReceiptsInactive V1
// Permissions Required:
//   - Sales
func (c *MetrcClient) SalesGetReceiptsInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string, salesdateend string, salesdatestart string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if salesdateend != "" {
		queryParams["salesDateEnd"] = salesdateend
	}
	if salesdatestart != "" {
		queryParams["salesDateStart"] = salesdatestart
	}
	return c.send("GET", "/sales/v1/receipts/inactive", queryParams, nil)
}

// GET GetReceiptsInactive V2
// Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
// 
//   Permissions Required:
//   - View Sales
//   - Manage Sales
func (c *MetrcClient) SalesGetReceiptsInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string, salesdateend string, salesdatestart string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	if salesdateend != "" {
		queryParams["salesDateEnd"] = salesdateend
	}
	if salesdatestart != "" {
		queryParams["salesDateStart"] = salesdatestart
	}
	return c.send("GET", "/sales/v2/receipts/inactive", queryParams, nil)
}

// GET GetTransactions V1
// Permissions Required:
//   - Sales
func (c *MetrcClient) SalesGetTransactionsV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v1/transactions", queryParams, nil)
}

// GET GetTransactionsBySalesDateStartAndSalesDateEnd V1
// Permissions Required:
//   - Sales
func (c *MetrcClient) SalesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesdatestart string, salesdateend string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/sales/v1/transactions/"+url.QueryEscape(salesdatestart)+"/"+url.QueryEscape(salesdateend)+"", queryParams, nil)
}

// PUT UpdateDelivery V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Sales Delivery
func (c *MetrcClient) SalesUpdateDeliveryV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v1/deliveries", queryParams, body)
}

// PUT UpdateDelivery V2
// Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Manage Sales Delivery
func (c *MetrcClient) SalesUpdateDeliveryV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v2/deliveries", queryParams, body)
}

// PUT UpdateDeliveryComplete V1
// Permissions Required:
//   - Sales Delivery
func (c *MetrcClient) SalesUpdateDeliveryCompleteV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v1/deliveries/complete", queryParams, body)
}

// PUT UpdateDeliveryComplete V2
// Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
// 
//   Permissions Required:
//   - Manage Sales Delivery
func (c *MetrcClient) SalesUpdateDeliveryCompleteV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v2/deliveries/complete", queryParams, body)
}

// PUT UpdateDeliveryHub V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Sales Delivery
func (c *MetrcClient) SalesUpdateDeliveryHubV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v1/deliveries/hub", queryParams, body)
}

// PUT UpdateDeliveryHub V2
// Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Manage Sales Delivery, Manage Sales Delivery Hub
func (c *MetrcClient) SalesUpdateDeliveryHubV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v2/deliveries/hub", queryParams, body)
}

// PUT UpdateDeliveryHubAccept V1
// Permissions Required:
//   - Sales
func (c *MetrcClient) SalesUpdateDeliveryHubAcceptV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v1/deliveries/hub/accept", queryParams, body)
}

// PUT UpdateDeliveryHubAccept V2
// Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
// 
//   Permissions Required:
//   - Manage Sales Delivery Hub
func (c *MetrcClient) SalesUpdateDeliveryHubAcceptV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v2/deliveries/hub/accept", queryParams, body)
}

// PUT UpdateDeliveryHubDepart V1
// Permissions Required:
//   - Sales
func (c *MetrcClient) SalesUpdateDeliveryHubDepartV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v1/deliveries/hub/depart", queryParams, body)
}

// PUT UpdateDeliveryHubDepart V2
// Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
// 
//   Permissions Required:
//   - Manage Sales Delivery Hub
func (c *MetrcClient) SalesUpdateDeliveryHubDepartV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v2/deliveries/hub/depart", queryParams, body)
}

// PUT UpdateDeliveryHubVerifyID V1
// Permissions Required:
//   - Sales
func (c *MetrcClient) SalesUpdateDeliveryHubVerifyIdV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v1/deliveries/hub/verifyID", queryParams, body)
}

// PUT UpdateDeliveryHubVerifyID V2
// Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
// 
//   Permissions Required:
//   - Manage Sales Delivery Hub
func (c *MetrcClient) SalesUpdateDeliveryHubVerifyIdV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v2/deliveries/hub/verifyID", queryParams, body)
}

// PUT UpdateDeliveryRetailer V1
// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Retailer Delivery
func (c *MetrcClient) SalesUpdateDeliveryRetailerV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v1/deliveries/retailer", queryParams, body)
}

// PUT UpdateDeliveryRetailer V2
// Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
//   - Industry/Facility Type/Retailer Delivery
//   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
//   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
//   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
//   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
//   - Manage Retailer Delivery
func (c *MetrcClient) SalesUpdateDeliveryRetailerV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v2/deliveries/retailer", queryParams, body)
}

// PUT UpdateReceipt V1
// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Sales
func (c *MetrcClient) SalesUpdateReceiptV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v1/receipts", queryParams, body)
}

// PUT UpdateReceipt V2
// Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
// 
//   Permissions Required:
//   - Manage Sales
func (c *MetrcClient) SalesUpdateReceiptV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v2/receipts", queryParams, body)
}

// PUT UpdateReceiptFinalize V2
// Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
// 
//   Permissions Required:
//   - Manage Sales
func (c *MetrcClient) SalesUpdateReceiptFinalizeV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v2/receipts/finalize", queryParams, body)
}

// PUT UpdateReceiptUnfinalize V2
// Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
// 
//   Permissions Required:
//   - Manage Sales
func (c *MetrcClient) SalesUpdateReceiptUnfinalizeV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v2/receipts/unfinalize", queryParams, body)
}

// PUT UpdateTransactionByDate V1
// Permissions Required:
//   - Sales
func (c *MetrcClient) SalesUpdateTransactionByDateV1(date string, licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/sales/v1/transactions/"+url.QueryEscape(date)+"", queryParams, body)
}

// POST CreateDriver V2
// Creates new driver records for a Facility.
// 
//   Permissions Required:
//   - Manage Transporters
func (c *MetrcClient) TransportersCreateDriverV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/transporters/v2/drivers", queryParams, body)
}

// POST CreateVehicle V2
// Creates new vehicle records for a Facility.
// 
//   Permissions Required:
//   - Manage Transporters
func (c *MetrcClient) TransportersCreateVehicleV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/transporters/v2/vehicles", queryParams, body)
}

// DELETE DeleteDriver V2
// Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
// 
//   Permissions Required:
//   - Manage Transporters
func (c *MetrcClient) TransportersDeleteDriverV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/transporters/v2/drivers/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteVehicle V2
// Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
// 
//   Permissions Required:
//   - Manage Transporters
func (c *MetrcClient) TransportersDeleteVehicleV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/transporters/v2/vehicles/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetDriver V2
// Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
// 
//   Permissions Required:
//   - Transporters
func (c *MetrcClient) TransportersGetDriverV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/transporters/v2/drivers/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetDrivers V2
// Retrieves a list of drivers for a Facility.
// 
//   Permissions Required:
//   - Transporters
func (c *MetrcClient) TransportersGetDriversV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transporters/v2/drivers", queryParams, nil)
}

// GET GetVehicle V2
// Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
// 
//   Permissions Required:
//   - Transporters
func (c *MetrcClient) TransportersGetVehicleV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/transporters/v2/vehicles/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetVehicles V2
// Retrieves a list of vehicles for a Facility.
// 
//   Permissions Required:
//   - Transporters
func (c *MetrcClient) TransportersGetVehiclesV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/transporters/v2/vehicles", queryParams, nil)
}

// PUT UpdateDriver V2
// Updates existing driver records for a Facility.
// 
//   Permissions Required:
//   - Manage Transporters
func (c *MetrcClient) TransportersUpdateDriverV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/transporters/v2/drivers", queryParams, body)
}

// PUT UpdateVehicle V2
// Updates existing vehicle records for a facility.
// 
//   Permissions Required:
//   - Manage Transporters
func (c *MetrcClient) TransportersUpdateVehicleV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/transporters/v2/vehicles", queryParams, body)
}

// GET GetAll V1
// This endpoint provides a list of facilities for which the authenticated user has access.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) FacilitiesGetAllV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/facilities/v1", queryParams, nil)
}

// GET GetAll V2
// This endpoint provides a list of facilities for which the authenticated user has access.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) FacilitiesGetAllV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/facilities/v2", queryParams, nil)
}

// POST Create V1
// Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PackagesCreateV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v1/create", queryParams, body)
}

// POST Create V2
// Creates new packages for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PackagesCreateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v2", queryParams, body)
}

// POST CreateAdjust V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesCreateAdjustV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v1/adjust", queryParams, body)
}

// POST CreateAdjust V2
// Records a list of adjustments for packages at a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesCreateAdjustV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v2/adjust", queryParams, body)
}

// POST CreateChangeItem V1
// Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PackagesCreateChangeItemV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v1/change/item", queryParams, body)
}

// POST CreateChangeLocation V1
// Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PackagesCreateChangeLocationV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v1/change/locations", queryParams, body)
}

// POST CreateFinish V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesCreateFinishV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v1/finish", queryParams, body)
}

// POST CreatePlantings V1
// Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesCreatePlantingsV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v1/create/plantings", queryParams, body)
}

// POST CreatePlantings V2
// Creates new plantings from packages for a specified Facility.
// 
//   Permissions Required:
//   - View Immature Plants
//   - Manage Immature Plants
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesCreatePlantingsV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v2/plantings", queryParams, body)
}

// POST CreateRemediate V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesCreateRemediateV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v1/remediate", queryParams, body)
}

// POST CreateTesting V1
// Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PackagesCreateTestingV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v1/create/testing", queryParams, body)
}

// POST CreateTesting V2
// Creates new packages for testing for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PackagesCreateTestingV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v2/testing", queryParams, body)
}

// POST CreateUnfinish V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesCreateUnfinishV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/packages/v1/unfinish", queryParams, body)
}

// DELETE Delete V2
// Discontinues a Package at a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PackagesDeleteV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/packages/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V1
// Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/packages/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V2
// Retrieves a Package by its Id.
// 
//   Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/packages/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetActive V1
// Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetActiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/packages/v1/active", queryParams, nil)
}

// GET GetActive V2
// Retrieves a list of active packages for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/packages/v2/active", queryParams, nil)
}

// GET GetAdjustReasons V1
// Permissions Required:
//   - None
func (c *MetrcClient) PackagesGetAdjustReasonsV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/packages/v1/adjust/reasons", queryParams, nil)
}

// GET GetAdjustReasons V2
// Retrieves a list of adjustment reasons for packages at a specified Facility.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) PackagesGetAdjustReasonsV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/packages/v2/adjust/reasons", queryParams, nil)
}

// GET GetByLabel V1
// Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetByLabelV1(label string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/packages/v1/"+url.QueryEscape(label)+"", queryParams, nil)
}

// GET GetByLabel V2
// Retrieves a Package by its label.
// 
//   Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetByLabelV2(label string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/packages/v2/"+url.QueryEscape(label)+"", queryParams, nil)
}

// GET GetInactive V1
// Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetInactiveV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/packages/v1/inactive", queryParams, nil)
}

// GET GetInactive V2
// Retrieves a list of inactive packages for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetInactiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/packages/v2/inactive", queryParams, nil)
}

// GET GetIntransit V2
// Retrieves a list of packages in transit for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetIntransitV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/packages/v2/intransit", queryParams, nil)
}

// GET GetLabsamples V2
// Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetLabsamplesV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/packages/v2/labsamples", queryParams, nil)
}

// GET GetOnhold V1
// Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetOnholdV1(lastmodifiedend string, lastmodifiedstart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/packages/v1/onhold", queryParams, nil)
}

// GET GetOnhold V2
// Retrieves a list of packages on hold for a specified Facility.
// 
//   Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetOnholdV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/packages/v2/onhold", queryParams, nil)
}

// GET GetSourceHarvest V2
// Retrieves the source harvests for a Package by its Id.
// 
//   Permissions Required:
//   - View Package Source Harvests
func (c *MetrcClient) PackagesGetSourceHarvestV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/packages/v2/"+url.QueryEscape(id)+"/source/harvests", queryParams, nil)
}

// GET GetTransferred V2
// Retrieves a list of transferred packages for a specific Facility.
// 
//   Permissions Required:
//   - View Packages
func (c *MetrcClient) PackagesGetTransferredV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/packages/v2/transferred", queryParams, nil)
}

// GET GetTypes V1
// Permissions Required:
//   - None
func (c *MetrcClient) PackagesGetTypesV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/packages/v1/types", queryParams, nil)
}

// GET GetTypes V2
// Retrieves a list of available Package types.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) PackagesGetTypesV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/packages/v2/types", queryParams, nil)
}

// PUT UpdateAdjust V2
// Set the final quantity for a Package.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesUpdateAdjustV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/adjust", queryParams, body)
}

// PUT UpdateChangeNote V1
// Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
//   - Manage Package Notes
func (c *MetrcClient) PackagesUpdateChangeNoteV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v1/change/note", queryParams, body)
}

// PUT UpdateDecontaminate V2
// Updates the Product decontaminate information for a list of packages at a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesUpdateDecontaminateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/decontaminate", queryParams, body)
}

// PUT UpdateDonationFlag V2
// Flags one or more packages for donation at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesUpdateDonationFlagV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/donation/flag", queryParams, body)
}

// PUT UpdateDonationUnflag V2
// Removes the donation flag from one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesUpdateDonationUnflagV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/donation/unflag", queryParams, body)
}

// PUT UpdateExternalid V2
// Updates the external identifiers for one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Package Inventory
//   - External Id Enabled
func (c *MetrcClient) PackagesUpdateExternalidV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/externalid", queryParams, body)
}

// PUT UpdateFinish V2
// Updates a list of packages as finished for a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesUpdateFinishV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/finish", queryParams, body)
}

// PUT UpdateItem V2
// Updates the associated Item for one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PackagesUpdateItemV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/item", queryParams, body)
}

// PUT UpdateLabTestRequired V2
// Updates the list of required lab test batches for one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PackagesUpdateLabTestRequiredV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/labtests/required", queryParams, body)
}

// PUT UpdateLocation V2
// Updates the Location and Sublocation for one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PackagesUpdateLocationV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/location", queryParams, body)
}

// PUT UpdateNote V2
// Updates notes associated with one or more packages for the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
//   - Manage Package Notes
func (c *MetrcClient) PackagesUpdateNoteV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/note", queryParams, body)
}

// PUT UpdateRemediate V2
// Updates a list of Product remediations for packages at a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesUpdateRemediateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/remediate", queryParams, body)
}

// PUT UpdateTradesampleFlag V2
// Flags or unflags one or more packages at the specified Facility as trade samples.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesUpdateTradesampleFlagV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/tradesample/flag", queryParams, body)
}

// PUT UpdateTradesampleUnflag V2
// Removes the trade sample flag from one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesUpdateTradesampleUnflagV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/tradesample/unflag", queryParams, body)
}

// PUT UpdateUnfinish V2
// Updates a list of packages as unfinished for a specific Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Manage Packages Inventory
func (c *MetrcClient) PackagesUpdateUnfinishV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/unfinish", queryParams, body)
}

// PUT UpdateUsebydate V2
// Updates the use-by date for one or more packages at the specified Facility.
// 
//   Permissions Required:
//   - View Packages
//   - Create/Submit/Discontinue Packages
func (c *MetrcClient) PackagesUpdateUsebydateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/packages/v2/usebydate", queryParams, body)
}

// POST Create V1
// Permissions Required:
//   - ManagePatientsCheckIns
func (c *MetrcClient) PatientCheckInsCreateV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/patient-checkins/v1", queryParams, body)
}

// POST Create V2
// Records patient check-ins for a specified Facility.
// 
//   Permissions Required:
//   - ManagePatientsCheckIns
func (c *MetrcClient) PatientCheckInsCreateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/patient-checkins/v2", queryParams, body)
}

// DELETE Delete V1
// Permissions Required:
//   - ManagePatientsCheckIns
func (c *MetrcClient) PatientCheckInsDeleteV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/patient-checkins/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE Delete V2
// Archives a Patient Check-In, identified by its Id, for a specified Facility.
// 
//   Permissions Required:
//   - ManagePatientsCheckIns
func (c *MetrcClient) PatientCheckInsDeleteV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/patient-checkins/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetAll V1
// Permissions Required:
//   - ManagePatientsCheckIns
func (c *MetrcClient) PatientCheckInsGetAllV1(checkindateend string, checkindatestart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if checkindateend != "" {
		queryParams["checkinDateEnd"] = checkindateend
	}
	if checkindatestart != "" {
		queryParams["checkinDateStart"] = checkindatestart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/patient-checkins/v1", queryParams, nil)
}

// GET GetAll V2
// Retrieves a list of patient check-ins for a specified Facility.
// 
//   Permissions Required:
//   - ManagePatientsCheckIns
func (c *MetrcClient) PatientCheckInsGetAllV2(checkindateend string, checkindatestart string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if checkindateend != "" {
		queryParams["checkinDateEnd"] = checkindateend
	}
	if checkindatestart != "" {
		queryParams["checkinDateStart"] = checkindatestart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/patient-checkins/v2", queryParams, nil)
}

// GET GetLocations V1
// Permissions Required:
//   - None
func (c *MetrcClient) PatientCheckInsGetLocationsV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/patient-checkins/v1/locations", queryParams, nil)
}

// GET GetLocations V2
// Retrieves a list of Patient Check-In locations.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) PatientCheckInsGetLocationsV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/patient-checkins/v2/locations", queryParams, nil)
}

// PUT Update V1
// Permissions Required:
//   - ManagePatientsCheckIns
func (c *MetrcClient) PatientCheckInsUpdateV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/patient-checkins/v1", queryParams, body)
}

// PUT Update V2
// Updates patient check-ins for a specified Facility.
// 
//   Permissions Required:
//   - ManagePatientsCheckIns
func (c *MetrcClient) PatientCheckInsUpdateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/patient-checkins/v2", queryParams, body)
}

// GET GetActive V1
// Permissions Required:
//   - None
func (c *MetrcClient) UnitsOfMeasureGetActiveV1(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/unitsofmeasure/v1/active", queryParams, nil)
}

// GET GetActive V2
// Retrieves all active units of measure.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) UnitsOfMeasureGetActiveV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/unitsofmeasure/v2/active", queryParams, nil)
}

// GET GetInactive V2
// Retrieves all inactive units of measure.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) UnitsOfMeasureGetInactiveV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/unitsofmeasure/v2/inactive", queryParams, nil)
}

// GET GetAll V2
// Retrieves all available waste methods.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) WasteMethodsGetAllV2(no string) (interface{}, error) {
	queryParams := make(map[string]string)
	if no != "" {
		queryParams["No"] = no
	}
	return c.send("GET", "/wastemethods/v2", queryParams, nil)
}

// GET GetAll V1
// Permissions Required:
//   - Manage Employees
func (c *MetrcClient) EmployeesGetAllV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/employees/v1", queryParams, nil)
}

// GET GetAll V2
// Retrieves a list of employees for a specified Facility.
// 
//   Permissions Required:
//   - Manage Employees
//   - View Employees
func (c *MetrcClient) EmployeesGetAllV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/employees/v2", queryParams, nil)
}

// GET GetPermissions V2
// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
// 
//   Permissions Required:
//   - Manage Employees
func (c *MetrcClient) EmployeesGetPermissionsV2(employeelicensenumber string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if employeelicensenumber != "" {
		queryParams["employeeLicenseNumber"] = employeelicensenumber
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/employees/v2/permissions", queryParams, nil)
}

// POST Create V1
// NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsCreateV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/items/v1/create", queryParams, body)
}

// POST Create V2
// Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsCreateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/items/v2", queryParams, body)
}

// POST CreateBrand V2
// Creates one or more new item brands for the specified Facility identified by the License Number.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsCreateBrandV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/items/v2/brand", queryParams, body)
}

// POST CreateFile V2
// Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsCreateFileV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/items/v2/file", queryParams, body)
}

// POST CreatePhoto V1
// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsCreatePhotoV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/items/v1/photo", queryParams, body)
}

// POST CreatePhoto V2
// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsCreatePhotoV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/items/v2/photo", queryParams, body)
}

// POST CreateUpdate V1
// Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsCreateUpdateV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/items/v1/update", queryParams, body)
}

// DELETE Delete V1
// Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsDeleteV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/items/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE Delete V2
// Archives the specified Product by Id for the given Facility License Number.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsDeleteV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/items/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE DeleteBrand V2
// Archives the specified Item Brand by Id for the given Facility License Number.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsDeleteBrandV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/items/v2/brand/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V1
// Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsGetV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/items/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V2
// Retrieves detailed information about a specific Item by Id.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsGetV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/items/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetActive V1
// Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsGetActiveV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/items/v1/active", queryParams, nil)
}

// GET GetActive V2
// Returns a list of active items for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/items/v2/active", queryParams, nil)
}

// GET GetBrands V1
// Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsGetBrandsV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/items/v1/brands", queryParams, nil)
}

// GET GetBrands V2
// Retrieves a list of active item brands for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsGetBrandsV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/items/v2/brands", queryParams, nil)
}

// GET GetCategories V1
// Permissions Required:
//   - None
func (c *MetrcClient) ItemsGetCategoriesV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/items/v1/categories", queryParams, nil)
}

// GET GetCategories V2
// Retrieves a list of item categories.
// 
//   Permissions Required:
//   - None
func (c *MetrcClient) ItemsGetCategoriesV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/items/v2/categories", queryParams, nil)
}

// GET GetFile V2
// Retrieves a file by its Id for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsGetFileV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/items/v2/file/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetInactive V1
// Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsGetInactiveV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/items/v1/inactive", queryParams, nil)
}

// GET GetInactive V2
// Retrieves a list of inactive items for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsGetInactiveV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/items/v2/inactive", queryParams, nil)
}

// GET GetPhoto V1
// Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsGetPhotoV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/items/v1/photo/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetPhoto V2
// Retrieves an image by its Id for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsGetPhotoV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/items/v2/photo/"+url.QueryEscape(id)+"", queryParams, nil)
}

// PUT Update V2
// Updates one or more existing products for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsUpdateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/items/v2", queryParams, body)
}

// PUT UpdateBrand V2
// Updates one or more existing item brands for the specified Facility.
// 
//   Permissions Required:
//   - Manage Items
func (c *MetrcClient) ItemsUpdateBrandV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/items/v2/brand", queryParams, body)
}

// POST Create V1
// Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsCreateV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/locations/v1/create", queryParams, body)
}

// POST Create V2
// Creates new locations for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsCreateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/locations/v2", queryParams, body)
}

// POST CreateUpdate V1
// Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsCreateUpdateV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/locations/v1/update", queryParams, body)
}

// DELETE Delete V1
// Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsDeleteV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/locations/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE Delete V2
// Archives a specified Location, identified by its Id, for a Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsDeleteV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/locations/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V1
// Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsGetV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/locations/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V2
// Retrieves a Location by its Id.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsGetV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/locations/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetActive V1
// Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsGetActiveV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/locations/v1/active", queryParams, nil)
}

// GET GetActive V2
// Retrieves a list of active locations for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsGetActiveV2(lastmodifiedend string, lastmodifiedstart string, licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if lastmodifiedend != "" {
		queryParams["lastModifiedEnd"] = lastmodifiedend
	}
	if lastmodifiedstart != "" {
		queryParams["lastModifiedStart"] = lastmodifiedstart
	}
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/locations/v2/active", queryParams, nil)
}

// GET GetInactive V2
// Retrieves a list of inactive locations for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsGetInactiveV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/locations/v2/inactive", queryParams, nil)
}

// GET GetTypes V1
// Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsGetTypesV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/locations/v1/types", queryParams, nil)
}

// GET GetTypes V2
// Retrieves a list of active location types for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsGetTypesV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/locations/v2/types", queryParams, nil)
}

// PUT Update V2
// Updates existing locations for a specified Facility.
// 
//   Permissions Required:
//   - Manage Locations
func (c *MetrcClient) LocationsUpdateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/locations/v2", queryParams, body)
}

// POST Create V2
// Adds new patients to a specified Facility.
// 
//   Permissions Required:
//   - Manage Patients
func (c *MetrcClient) PatientsCreateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/patients/v2", queryParams, body)
}

// POST CreateAdd V1
// Permissions Required:
//   - Manage Patients
func (c *MetrcClient) PatientsCreateAddV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/patients/v1/add", queryParams, body)
}

// POST CreateUpdate V1
// Permissions Required:
//   - Manage Patients
func (c *MetrcClient) PatientsCreateUpdateV1(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("POST", "/patients/v1/update", queryParams, body)
}

// DELETE Delete V1
// Permissions Required:
//   - Manage Patients
func (c *MetrcClient) PatientsDeleteV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/patients/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// DELETE Delete V2
// Removes a Patient, identified by an Id, from a specified Facility.
// 
//   Permissions Required:
//   - Manage Patients
func (c *MetrcClient) PatientsDeleteV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("DELETE", "/patients/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V1
// Permissions Required:
//   - Manage Patients
func (c *MetrcClient) PatientsGetV1(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/patients/v1/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET Get V2
// Retrieves a Patient by Id.
// 
//   Permissions Required:
//   - Manage Patients
func (c *MetrcClient) PatientsGetV2(id string, licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/patients/v2/"+url.QueryEscape(id)+"", queryParams, nil)
}

// GET GetActive V1
// Permissions Required:
//   - Manage Patients
func (c *MetrcClient) PatientsGetActiveV1(licensenumber string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("GET", "/patients/v1/active", queryParams, nil)
}

// GET GetActive V2
// Retrieves a list of active patients for a specified Facility.
// 
//   Permissions Required:
//   - Manage Patients
func (c *MetrcClient) PatientsGetActiveV2(licensenumber string, pagenumber string, pagesize string) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	if pagenumber != "" {
		queryParams["pageNumber"] = pagenumber
	}
	if pagesize != "" {
		queryParams["pageSize"] = pagesize
	}
	return c.send("GET", "/patients/v2/active", queryParams, nil)
}

// PUT Update V2
// Updates Patient information for a specified Facility.
// 
//   Permissions Required:
//   - Manage Patients
func (c *MetrcClient) PatientsUpdateV2(licensenumber string, body interface{}) (interface{}, error) {
	queryParams := make(map[string]string)
	if licensenumber != "" {
		queryParams["licenseNumber"] = licensenumber
	}
	return c.send("PUT", "/patients/v2", queryParams, body)
}

