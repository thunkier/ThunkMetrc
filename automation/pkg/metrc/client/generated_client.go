package client

import (
	"fmt"
	"github.com/thunkmetrc/automation/pkg/metrc/models"
)

// Auto-Generated Client Methods

// AdditivesTemplates_CreateV2 (Additives Templates)
// POST {{baseUrl}}/additivestemplates/v2
func (f *Fetcher) AdditivesTemplates_CreateV2() (models.AdditivesTemplate, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.AdditivesTemplate](f, "AdditivesTemplates_CreateV2", "POST", "{{baseUrl}}/additivestemplates/v2")
	if err != nil {
		var empty models.AdditivesTemplate
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.AdditivesTemplate
	return empty, fmt.Errorf("not found")
}

// AdditivesTemplates_GetActiveV2 (Additives Templates)
// GET {{baseUrl}}/additivestemplates/v2/active
func (f *Fetcher) AdditivesTemplates_GetActiveV2() ([]models.AdditivesTemplate, error) {
	return fetchList[models.AdditivesTemplate](f, "AdditivesTemplates_GetActiveV2", "GET", "{{baseUrl}}/additivestemplates/v2/active")
}

// AdditivesTemplates_GetInactiveV2 (Additives Templates)
// GET {{baseUrl}}/additivestemplates/v2/inactive
func (f *Fetcher) AdditivesTemplates_GetInactiveV2() ([]models.AdditivesTemplate, error) {
	return fetchList[models.AdditivesTemplate](f, "AdditivesTemplates_GetInactiveV2", "GET", "{{baseUrl}}/additivestemplates/v2/inactive")
}

// AdditivesTemplates_GetV2 (Additives Templates)
// GET {{baseUrl}}/additivestemplates/v2/{id}
func (f *Fetcher) AdditivesTemplates_GetV2() (models.AdditivesTemplate, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.AdditivesTemplate](f, "AdditivesTemplates_GetV2", "GET", "{{baseUrl}}/additivestemplates/v2/{id}")
	if err != nil {
		var empty models.AdditivesTemplate
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.AdditivesTemplate
	return empty, fmt.Errorf("not found")
}

// AdditivesTemplates_UpdateV2 (Additives Templates)
// PUT {{baseUrl}}/additivestemplates/v2
func (f *Fetcher) AdditivesTemplates_UpdateV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "AdditivesTemplates_UpdateV2", "PUT", "{{baseUrl}}/additivestemplates/v2")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// CaregiversStatus_GetByCaregiverLicenseNumberV1 (Caregivers Status)
// GET {{baseUrl}}/caregivers/v1/status/{caregiverLicenseNumber}
func (f *Fetcher) CaregiversStatus_GetByCaregiverLicenseNumberV1() (models.CaregiversStatu, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.CaregiversStatu](f, "CaregiversStatus_GetByCaregiverLicenseNumberV1", "GET", "{{baseUrl}}/caregivers/v1/status/{caregiverLicenseNumber}")
	if err != nil {
		var empty models.CaregiversStatu
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.CaregiversStatu
	return empty, fmt.Errorf("not found")
}

// CaregiversStatus_GetByCaregiverLicenseNumberV2 (Caregivers Status)
// GET {{baseUrl}}/caregivers/v2/status/{caregiverLicenseNumber}
func (f *Fetcher) CaregiversStatus_GetByCaregiverLicenseNumberV2() (models.CaregiversStatu, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.CaregiversStatu](f, "CaregiversStatus_GetByCaregiverLicenseNumberV2", "GET", "{{baseUrl}}/caregivers/v2/status/{caregiverLicenseNumber}")
	if err != nil {
		var empty models.CaregiversStatu
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.CaregiversStatu
	return empty, fmt.Errorf("not found")
}

// Employees_GetAllV1 (Employees)
// GET {{baseUrl}}/employees/v1
func (f *Fetcher) Employees_GetAllV1() ([]models.Employee, error) {
	return fetchList[models.Employee](f, "Employees_GetAllV1", "GET", "{{baseUrl}}/employees/v1")
}

// Employees_GetAllV2 (Employees)
// GET {{baseUrl}}/employees/v2
func (f *Fetcher) Employees_GetAllV2() ([]models.Employee, error) {
	return fetchList[models.Employee](f, "Employees_GetAllV2", "GET", "{{baseUrl}}/employees/v2")
}

// Employees_GetPermissionsV2 (Employees)
// GET {{baseUrl}}/employees/v2/permissions
func (f *Fetcher) Employees_GetPermissionsV2() ([]models.Employee, error) {
	return fetchList[models.Employee](f, "Employees_GetPermissionsV2", "GET", "{{baseUrl}}/employees/v2/permissions")
}

// Facilities_GetAllV1 (Facilities)
// GET {{baseUrl}}/facilities/v1
func (f *Fetcher) Facilities_GetAllV1() ([]models.Facility, error) {
	return fetchList[models.Facility](f, "Facilities_GetAllV1", "GET", "{{baseUrl}}/facilities/v1")
}

// Facilities_GetAllV2 (Facilities)
// GET {{baseUrl}}/facilities/v2
func (f *Fetcher) Facilities_GetAllV2() ([]models.Facility, error) {
	return fetchList[models.Facility](f, "Facilities_GetAllV2", "GET", "{{baseUrl}}/facilities/v2")
}

// Harvests_CreateFinishV1 (Harvests)
// POST {{baseUrl}}/harvests/v1/finish
func (f *Fetcher) Harvests_CreateFinishV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Harvests_CreateFinishV1", "POST", "{{baseUrl}}/harvests/v1/finish")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Harvests_CreatePackageTestingV1 (Harvests)
// POST {{baseUrl}}/harvests/v1/create/packages/testing
func (f *Fetcher) Harvests_CreatePackageTestingV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Harvests_CreatePackageTestingV1", "POST", "{{baseUrl}}/harvests/v1/create/packages/testing")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Harvests_CreatePackageTestingV2 (Harvests)
// POST {{baseUrl}}/harvests/v2/packages/testing
func (f *Fetcher) Harvests_CreatePackageTestingV2() (models.Harvest, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Harvest](f, "Harvests_CreatePackageTestingV2", "POST", "{{baseUrl}}/harvests/v2/packages/testing")
	if err != nil {
		var empty models.Harvest
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Harvest
	return empty, fmt.Errorf("not found")
}

// Harvests_CreatePackageV1 (Harvests)
// POST {{baseUrl}}/harvests/v1/create/packages
func (f *Fetcher) Harvests_CreatePackageV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Harvests_CreatePackageV1", "POST", "{{baseUrl}}/harvests/v1/create/packages")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Harvests_CreatePackageV2 (Harvests)
// POST {{baseUrl}}/harvests/v2/packages
func (f *Fetcher) Harvests_CreatePackageV2() (models.Harvest, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Harvest](f, "Harvests_CreatePackageV2", "POST", "{{baseUrl}}/harvests/v2/packages")
	if err != nil {
		var empty models.Harvest
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Harvest
	return empty, fmt.Errorf("not found")
}

// Harvests_CreateRemoveWasteV1 (Harvests)
// POST {{baseUrl}}/harvests/v1/removewaste
func (f *Fetcher) Harvests_CreateRemoveWasteV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Harvests_CreateRemoveWasteV1", "POST", "{{baseUrl}}/harvests/v1/removewaste")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Harvests_CreateUnfinishV1 (Harvests)
// POST {{baseUrl}}/harvests/v1/unfinish
func (f *Fetcher) Harvests_CreateUnfinishV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Harvests_CreateUnfinishV1", "POST", "{{baseUrl}}/harvests/v1/unfinish")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Harvests_CreateWasteV2 (Harvests)
// POST {{baseUrl}}/harvests/v2/waste
func (f *Fetcher) Harvests_CreateWasteV2() (models.Harvest, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Harvest](f, "Harvests_CreateWasteV2", "POST", "{{baseUrl}}/harvests/v2/waste")
	if err != nil {
		var empty models.Harvest
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Harvest
	return empty, fmt.Errorf("not found")
}

// Harvests_DeleteWasteV2 (Harvests)
// DELETE {{baseUrl}}/harvests/v2/waste/{id}
func (f *Fetcher) Harvests_DeleteWasteV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Harvests_DeleteWasteV2", "DELETE", "{{baseUrl}}/harvests/v2/waste/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Harvests_GetActiveV1 (Harvests)
// GET {{baseUrl}}/harvests/v1/active
func (f *Fetcher) Harvests_GetActiveV1() ([]models.Harvest, error) {
	return fetchList[models.Harvest](f, "Harvests_GetActiveV1", "GET", "{{baseUrl}}/harvests/v1/active")
}

// Harvests_GetActiveV2 (Harvests)
// GET {{baseUrl}}/harvests/v2/active
func (f *Fetcher) Harvests_GetActiveV2() ([]models.Harvest, error) {
	return fetchList[models.Harvest](f, "Harvests_GetActiveV2", "GET", "{{baseUrl}}/harvests/v2/active")
}

// Harvests_GetInactiveV1 (Harvests)
// GET {{baseUrl}}/harvests/v1/inactive
func (f *Fetcher) Harvests_GetInactiveV1() ([]models.Harvest, error) {
	return fetchList[models.Harvest](f, "Harvests_GetInactiveV1", "GET", "{{baseUrl}}/harvests/v1/inactive")
}

// Harvests_GetInactiveV2 (Harvests)
// GET {{baseUrl}}/harvests/v2/inactive
func (f *Fetcher) Harvests_GetInactiveV2() ([]models.Harvest, error) {
	return fetchList[models.Harvest](f, "Harvests_GetInactiveV2", "GET", "{{baseUrl}}/harvests/v2/inactive")
}

// Harvests_GetOnholdV1 (Harvests)
// GET {{baseUrl}}/harvests/v1/onhold
func (f *Fetcher) Harvests_GetOnholdV1() ([]models.Harvest, error) {
	return fetchList[models.Harvest](f, "Harvests_GetOnholdV1", "GET", "{{baseUrl}}/harvests/v1/onhold")
}

// Harvests_GetOnholdV2 (Harvests)
// GET {{baseUrl}}/harvests/v2/onhold
func (f *Fetcher) Harvests_GetOnholdV2() ([]models.Harvest, error) {
	return fetchList[models.Harvest](f, "Harvests_GetOnholdV2", "GET", "{{baseUrl}}/harvests/v2/onhold")
}

// Harvests_GetV1 (Harvests)
// GET {{baseUrl}}/harvests/v1/{id}
func (f *Fetcher) Harvests_GetV1() (models.Harvest, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Harvest](f, "Harvests_GetV1", "GET", "{{baseUrl}}/harvests/v1/{id}")
	if err != nil {
		var empty models.Harvest
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Harvest
	return empty, fmt.Errorf("not found")
}

// Harvests_GetV2 (Harvests)
// GET {{baseUrl}}/harvests/v2/{id}
func (f *Fetcher) Harvests_GetV2() (models.Harvest, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Harvest](f, "Harvests_GetV2", "GET", "{{baseUrl}}/harvests/v2/{id}")
	if err != nil {
		var empty models.Harvest
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Harvest
	return empty, fmt.Errorf("not found")
}

// Harvests_GetWasteTypesV1 (Harvests)
// GET {{baseUrl}}/harvests/v1/waste/types
func (f *Fetcher) Harvests_GetWasteTypesV1() ([]models.Harvest, error) {
	return fetchList[models.Harvest](f, "Harvests_GetWasteTypesV1", "GET", "{{baseUrl}}/harvests/v1/waste/types")
}

// Harvests_GetWasteTypesV2 (Harvests)
// GET {{baseUrl}}/harvests/v2/waste/types
func (f *Fetcher) Harvests_GetWasteTypesV2() ([]models.Harvest, error) {
	return fetchList[models.Harvest](f, "Harvests_GetWasteTypesV2", "GET", "{{baseUrl}}/harvests/v2/waste/types")
}

// Harvests_GetWasteV2 (Harvests)
// GET {{baseUrl}}/harvests/v2/waste
func (f *Fetcher) Harvests_GetWasteV2() ([]models.Harvest, error) {
	return fetchList[models.Harvest](f, "Harvests_GetWasteV2", "GET", "{{baseUrl}}/harvests/v2/waste")
}

// Harvests_UpdateFinishV2 (Harvests)
// PUT {{baseUrl}}/harvests/v2/finish
func (f *Fetcher) Harvests_UpdateFinishV2() (models.Harvest, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Harvest](f, "Harvests_UpdateFinishV2", "PUT", "{{baseUrl}}/harvests/v2/finish")
	if err != nil {
		var empty models.Harvest
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Harvest
	return empty, fmt.Errorf("not found")
}

// Harvests_UpdateLocationV2 (Harvests)
// PUT {{baseUrl}}/harvests/v2/location
func (f *Fetcher) Harvests_UpdateLocationV2() (models.Harvest, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Harvest](f, "Harvests_UpdateLocationV2", "PUT", "{{baseUrl}}/harvests/v2/location")
	if err != nil {
		var empty models.Harvest
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Harvest
	return empty, fmt.Errorf("not found")
}

// Harvests_UpdateMoveV1 (Harvests)
// PUT {{baseUrl}}/harvests/v1/move
func (f *Fetcher) Harvests_UpdateMoveV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Harvests_UpdateMoveV1", "PUT", "{{baseUrl}}/harvests/v1/move")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Harvests_UpdateRenameV1 (Harvests)
// PUT {{baseUrl}}/harvests/v1/rename
func (f *Fetcher) Harvests_UpdateRenameV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Harvests_UpdateRenameV1", "PUT", "{{baseUrl}}/harvests/v1/rename")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Harvests_UpdateRenameV2 (Harvests)
// PUT {{baseUrl}}/harvests/v2/rename
func (f *Fetcher) Harvests_UpdateRenameV2() (models.Harvest, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Harvest](f, "Harvests_UpdateRenameV2", "PUT", "{{baseUrl}}/harvests/v2/rename")
	if err != nil {
		var empty models.Harvest
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Harvest
	return empty, fmt.Errorf("not found")
}

// Harvests_UpdateRestoreHarvestedPlantsV2 (Harvests)
// PUT {{baseUrl}}/harvests/v2/restore/harvestedplants
func (f *Fetcher) Harvests_UpdateRestoreHarvestedPlantsV2() (models.Harvest, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Harvest](f, "Harvests_UpdateRestoreHarvestedPlantsV2", "PUT", "{{baseUrl}}/harvests/v2/restore/harvestedplants")
	if err != nil {
		var empty models.Harvest
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Harvest
	return empty, fmt.Errorf("not found")
}

// Harvests_UpdateUnfinishV2 (Harvests)
// PUT {{baseUrl}}/harvests/v2/unfinish
func (f *Fetcher) Harvests_UpdateUnfinishV2() (models.Harvest, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Harvest](f, "Harvests_UpdateUnfinishV2", "PUT", "{{baseUrl}}/harvests/v2/unfinish")
	if err != nil {
		var empty models.Harvest
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Harvest
	return empty, fmt.Errorf("not found")
}

// Items_CreateBrandV2 (Items)
// POST {{baseUrl}}/items/v2/brand
func (f *Fetcher) Items_CreateBrandV2() (models.Item, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Item](f, "Items_CreateBrandV2", "POST", "{{baseUrl}}/items/v2/brand")
	if err != nil {
		var empty models.Item
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Item
	return empty, fmt.Errorf("not found")
}

// Items_CreateFileV2 (Items)
// POST {{baseUrl}}/items/v2/file
func (f *Fetcher) Items_CreateFileV2() (models.Item, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Item](f, "Items_CreateFileV2", "POST", "{{baseUrl}}/items/v2/file")
	if err != nil {
		var empty models.Item
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Item
	return empty, fmt.Errorf("not found")
}

// Items_CreatePhotoV1 (Items)
// POST {{baseUrl}}/items/v1/photo
func (f *Fetcher) Items_CreatePhotoV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Items_CreatePhotoV1", "POST", "{{baseUrl}}/items/v1/photo")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Items_CreatePhotoV2 (Items)
// POST {{baseUrl}}/items/v2/photo
func (f *Fetcher) Items_CreatePhotoV2() (models.Item, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Item](f, "Items_CreatePhotoV2", "POST", "{{baseUrl}}/items/v2/photo")
	if err != nil {
		var empty models.Item
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Item
	return empty, fmt.Errorf("not found")
}

// Items_CreateUpdateV1 (Items)
// POST {{baseUrl}}/items/v1/update
func (f *Fetcher) Items_CreateUpdateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Items_CreateUpdateV1", "POST", "{{baseUrl}}/items/v1/update")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Items_CreateV1 (Items)
// POST {{baseUrl}}/items/v1/create
func (f *Fetcher) Items_CreateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Items_CreateV1", "POST", "{{baseUrl}}/items/v1/create")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Items_CreateV2 (Items)
// POST {{baseUrl}}/items/v2
func (f *Fetcher) Items_CreateV2() (models.Item, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Item](f, "Items_CreateV2", "POST", "{{baseUrl}}/items/v2")
	if err != nil {
		var empty models.Item
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Item
	return empty, fmt.Errorf("not found")
}

// Items_DeleteBrandV2 (Items)
// DELETE {{baseUrl}}/items/v2/brand/{id}
func (f *Fetcher) Items_DeleteBrandV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Items_DeleteBrandV2", "DELETE", "{{baseUrl}}/items/v2/brand/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Items_DeleteV1 (Items)
// DELETE {{baseUrl}}/items/v1/{id}
func (f *Fetcher) Items_DeleteV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Items_DeleteV1", "DELETE", "{{baseUrl}}/items/v1/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Items_DeleteV2 (Items)
// DELETE {{baseUrl}}/items/v2/{id}
func (f *Fetcher) Items_DeleteV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Items_DeleteV2", "DELETE", "{{baseUrl}}/items/v2/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Items_GetActiveV1 (Items)
// GET {{baseUrl}}/items/v1/active
func (f *Fetcher) Items_GetActiveV1() ([]models.Item, error) {
	return fetchList[models.Item](f, "Items_GetActiveV1", "GET", "{{baseUrl}}/items/v1/active")
}

// Items_GetActiveV2 (Items)
// GET {{baseUrl}}/items/v2/active
func (f *Fetcher) Items_GetActiveV2() ([]models.Item, error) {
	return fetchList[models.Item](f, "Items_GetActiveV2", "GET", "{{baseUrl}}/items/v2/active")
}

// Items_GetBrandsV1 (Items)
// GET {{baseUrl}}/items/v1/brands
func (f *Fetcher) Items_GetBrandsV1() ([]models.Item, error) {
	return fetchList[models.Item](f, "Items_GetBrandsV1", "GET", "{{baseUrl}}/items/v1/brands")
}

// Items_GetBrandsV2 (Items)
// GET {{baseUrl}}/items/v2/brands
func (f *Fetcher) Items_GetBrandsV2() ([]models.Item, error) {
	return fetchList[models.Item](f, "Items_GetBrandsV2", "GET", "{{baseUrl}}/items/v2/brands")
}

// Items_GetCategoriesV1 (Items)
// GET {{baseUrl}}/items/v1/categories
func (f *Fetcher) Items_GetCategoriesV1() ([]models.Item, error) {
	return fetchList[models.Item](f, "Items_GetCategoriesV1", "GET", "{{baseUrl}}/items/v1/categories")
}

// Items_GetCategoriesV2 (Items)
// GET {{baseUrl}}/items/v2/categories
func (f *Fetcher) Items_GetCategoriesV2() ([]models.Item, error) {
	return fetchList[models.Item](f, "Items_GetCategoriesV2", "GET", "{{baseUrl}}/items/v2/categories")
}

// Items_GetFileV2 (Items)
// GET {{baseUrl}}/items/v2/file/{id}
func (f *Fetcher) Items_GetFileV2() (models.Item, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Item](f, "Items_GetFileV2", "GET", "{{baseUrl}}/items/v2/file/{id}")
	if err != nil {
		var empty models.Item
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Item
	return empty, fmt.Errorf("not found")
}

// Items_GetInactiveV1 (Items)
// GET {{baseUrl}}/items/v1/inactive
func (f *Fetcher) Items_GetInactiveV1() ([]models.Item, error) {
	return fetchList[models.Item](f, "Items_GetInactiveV1", "GET", "{{baseUrl}}/items/v1/inactive")
}

// Items_GetInactiveV2 (Items)
// GET {{baseUrl}}/items/v2/inactive
func (f *Fetcher) Items_GetInactiveV2() ([]models.Item, error) {
	return fetchList[models.Item](f, "Items_GetInactiveV2", "GET", "{{baseUrl}}/items/v2/inactive")
}

// Items_GetPhotoV1 (Items)
// GET {{baseUrl}}/items/v1/photo/{id}
func (f *Fetcher) Items_GetPhotoV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Items_GetPhotoV1", "GET", "{{baseUrl}}/items/v1/photo/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Items_GetPhotoV2 (Items)
// GET {{baseUrl}}/items/v2/photo/{id}
func (f *Fetcher) Items_GetPhotoV2() (models.Item, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Item](f, "Items_GetPhotoV2", "GET", "{{baseUrl}}/items/v2/photo/{id}")
	if err != nil {
		var empty models.Item
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Item
	return empty, fmt.Errorf("not found")
}

// Items_GetV1 (Items)
// GET {{baseUrl}}/items/v1/{id}
func (f *Fetcher) Items_GetV1() (models.Item, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Item](f, "Items_GetV1", "GET", "{{baseUrl}}/items/v1/{id}")
	if err != nil {
		var empty models.Item
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Item
	return empty, fmt.Errorf("not found")
}

// Items_GetV2 (Items)
// GET {{baseUrl}}/items/v2/{id}
func (f *Fetcher) Items_GetV2() (models.Item, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Item](f, "Items_GetV2", "GET", "{{baseUrl}}/items/v2/{id}")
	if err != nil {
		var empty models.Item
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Item
	return empty, fmt.Errorf("not found")
}

// Items_UpdateBrandV2 (Items)
// PUT {{baseUrl}}/items/v2/brand
func (f *Fetcher) Items_UpdateBrandV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Items_UpdateBrandV2", "PUT", "{{baseUrl}}/items/v2/brand")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Items_UpdateV2 (Items)
// PUT {{baseUrl}}/items/v2
func (f *Fetcher) Items_UpdateV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Items_UpdateV2", "PUT", "{{baseUrl}}/items/v2")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// LabTests_CreateRecordV1 (Lab Tests)
// POST {{baseUrl}}/labtests/v1/record
func (f *Fetcher) LabTests_CreateRecordV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "LabTests_CreateRecordV1", "POST", "{{baseUrl}}/labtests/v1/record")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// LabTests_CreateRecordV2 (Lab Tests)
// POST {{baseUrl}}/labtests/v2/record
func (f *Fetcher) LabTests_CreateRecordV2() (models.LabTest, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.LabTest](f, "LabTests_CreateRecordV2", "POST", "{{baseUrl}}/labtests/v2/record")
	if err != nil {
		var empty models.LabTest
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.LabTest
	return empty, fmt.Errorf("not found")
}

// LabTests_GetBatchesV2 (Lab Tests)
// GET {{baseUrl}}/labtests/v2/batches
func (f *Fetcher) LabTests_GetBatchesV2() ([]models.LabTest, error) {
	return fetchList[models.LabTest](f, "LabTests_GetBatchesV2", "GET", "{{baseUrl}}/labtests/v2/batches")
}

// LabTests_GetLabtestdocumentV1 (Lab Tests)
// GET {{baseUrl}}/labtests/v1/labtestdocument/{id}
func (f *Fetcher) LabTests_GetLabtestdocumentV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "LabTests_GetLabtestdocumentV1", "GET", "{{baseUrl}}/labtests/v1/labtestdocument/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// LabTests_GetLabtestdocumentV2 (Lab Tests)
// GET {{baseUrl}}/labtests/v2/labtestdocument/{id}
func (f *Fetcher) LabTests_GetLabtestdocumentV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "LabTests_GetLabtestdocumentV2", "GET", "{{baseUrl}}/labtests/v2/labtestdocument/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// LabTests_GetResultsV1 (Lab Tests)
// GET {{baseUrl}}/labtests/v1/results
func (f *Fetcher) LabTests_GetResultsV1() ([]models.LabTest, error) {
	return fetchList[models.LabTest](f, "LabTests_GetResultsV1", "GET", "{{baseUrl}}/labtests/v1/results")
}

// LabTests_GetResultsV2 (Lab Tests)
// GET {{baseUrl}}/labtests/v2/results
func (f *Fetcher) LabTests_GetResultsV2() ([]models.LabTest, error) {
	return fetchList[models.LabTest](f, "LabTests_GetResultsV2", "GET", "{{baseUrl}}/labtests/v2/results")
}

// LabTests_GetStatesV1 (Lab Tests)
// GET {{baseUrl}}/labtests/v1/states
func (f *Fetcher) LabTests_GetStatesV1() ([]models.LabTest, error) {
	return fetchList[models.LabTest](f, "LabTests_GetStatesV1", "GET", "{{baseUrl}}/labtests/v1/states")
}

// LabTests_GetStatesV2 (Lab Tests)
// GET {{baseUrl}}/labtests/v2/states
func (f *Fetcher) LabTests_GetStatesV2() ([]models.LabTest, error) {
	return fetchList[models.LabTest](f, "LabTests_GetStatesV2", "GET", "{{baseUrl}}/labtests/v2/states")
}

// LabTests_GetTypesV1 (Lab Tests)
// GET {{baseUrl}}/labtests/v1/types
func (f *Fetcher) LabTests_GetTypesV1() ([]models.LabTest, error) {
	return fetchList[models.LabTest](f, "LabTests_GetTypesV1", "GET", "{{baseUrl}}/labtests/v1/types")
}

// LabTests_GetTypesV2 (Lab Tests)
// GET {{baseUrl}}/labtests/v2/types
func (f *Fetcher) LabTests_GetTypesV2() ([]models.LabTest, error) {
	return fetchList[models.LabTest](f, "LabTests_GetTypesV2", "GET", "{{baseUrl}}/labtests/v2/types")
}

// LabTests_UpdateLabtestdocumentV1 (Lab Tests)
// PUT {{baseUrl}}/labtests/v1/labtestdocument
func (f *Fetcher) LabTests_UpdateLabtestdocumentV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "LabTests_UpdateLabtestdocumentV1", "PUT", "{{baseUrl}}/labtests/v1/labtestdocument")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// LabTests_UpdateLabtestdocumentV2 (Lab Tests)
// PUT {{baseUrl}}/labtests/v2/labtestdocument
func (f *Fetcher) LabTests_UpdateLabtestdocumentV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "LabTests_UpdateLabtestdocumentV2", "PUT", "{{baseUrl}}/labtests/v2/labtestdocument")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// LabTests_UpdateResultReleaseV1 (Lab Tests)
// PUT {{baseUrl}}/labtests/v1/results/release
func (f *Fetcher) LabTests_UpdateResultReleaseV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "LabTests_UpdateResultReleaseV1", "PUT", "{{baseUrl}}/labtests/v1/results/release")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// LabTests_UpdateResultReleaseV2 (Lab Tests)
// PUT {{baseUrl}}/labtests/v2/results/release
func (f *Fetcher) LabTests_UpdateResultReleaseV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "LabTests_UpdateResultReleaseV2", "PUT", "{{baseUrl}}/labtests/v2/results/release")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Locations_CreateUpdateV1 (Locations)
// POST {{baseUrl}}/locations/v1/update
func (f *Fetcher) Locations_CreateUpdateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Locations_CreateUpdateV1", "POST", "{{baseUrl}}/locations/v1/update")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Locations_CreateV1 (Locations)
// POST {{baseUrl}}/locations/v1/create
func (f *Fetcher) Locations_CreateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Locations_CreateV1", "POST", "{{baseUrl}}/locations/v1/create")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Locations_CreateV2 (Locations)
// POST {{baseUrl}}/locations/v2
func (f *Fetcher) Locations_CreateV2() (models.Location, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Location](f, "Locations_CreateV2", "POST", "{{baseUrl}}/locations/v2")
	if err != nil {
		var empty models.Location
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Location
	return empty, fmt.Errorf("not found")
}

// Locations_DeleteV1 (Locations)
// DELETE {{baseUrl}}/locations/v1/{id}
func (f *Fetcher) Locations_DeleteV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Locations_DeleteV1", "DELETE", "{{baseUrl}}/locations/v1/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Locations_DeleteV2 (Locations)
// DELETE {{baseUrl}}/locations/v2/{id}
func (f *Fetcher) Locations_DeleteV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Locations_DeleteV2", "DELETE", "{{baseUrl}}/locations/v2/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Locations_GetActiveV1 (Locations)
// GET {{baseUrl}}/locations/v1/active
func (f *Fetcher) Locations_GetActiveV1() ([]models.Location, error) {
	return fetchList[models.Location](f, "Locations_GetActiveV1", "GET", "{{baseUrl}}/locations/v1/active")
}

// Locations_GetActiveV2 (Locations)
// GET {{baseUrl}}/locations/v2/active
func (f *Fetcher) Locations_GetActiveV2() ([]models.Location, error) {
	return fetchList[models.Location](f, "Locations_GetActiveV2", "GET", "{{baseUrl}}/locations/v2/active")
}

// Locations_GetInactiveV2 (Locations)
// GET {{baseUrl}}/locations/v2/inactive
func (f *Fetcher) Locations_GetInactiveV2() ([]models.Location, error) {
	return fetchList[models.Location](f, "Locations_GetInactiveV2", "GET", "{{baseUrl}}/locations/v2/inactive")
}

// Locations_GetTypesV1 (Locations)
// GET {{baseUrl}}/locations/v1/types
func (f *Fetcher) Locations_GetTypesV1() ([]models.Location, error) {
	return fetchList[models.Location](f, "Locations_GetTypesV1", "GET", "{{baseUrl}}/locations/v1/types")
}

// Locations_GetTypesV2 (Locations)
// GET {{baseUrl}}/locations/v2/types
func (f *Fetcher) Locations_GetTypesV2() ([]models.Location, error) {
	return fetchList[models.Location](f, "Locations_GetTypesV2", "GET", "{{baseUrl}}/locations/v2/types")
}

// Locations_GetV1 (Locations)
// GET {{baseUrl}}/locations/v1/{id}
func (f *Fetcher) Locations_GetV1() (models.Location, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Location](f, "Locations_GetV1", "GET", "{{baseUrl}}/locations/v1/{id}")
	if err != nil {
		var empty models.Location
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Location
	return empty, fmt.Errorf("not found")
}

// Locations_GetV2 (Locations)
// GET {{baseUrl}}/locations/v2/{id}
func (f *Fetcher) Locations_GetV2() (models.Location, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Location](f, "Locations_GetV2", "GET", "{{baseUrl}}/locations/v2/{id}")
	if err != nil {
		var empty models.Location
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Location
	return empty, fmt.Errorf("not found")
}

// Locations_UpdateV2 (Locations)
// PUT {{baseUrl}}/locations/v2
func (f *Fetcher) Locations_UpdateV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Locations_UpdateV2", "PUT", "{{baseUrl}}/locations/v2")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_CreateAdjustV1 (Packages)
// POST {{baseUrl}}/packages/v1/adjust
func (f *Fetcher) Packages_CreateAdjustV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_CreateAdjustV1", "POST", "{{baseUrl}}/packages/v1/adjust")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_CreateAdjustV2 (Packages)
// POST {{baseUrl}}/packages/v2/adjust
func (f *Fetcher) Packages_CreateAdjustV2() (models.Package, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Package](f, "Packages_CreateAdjustV2", "POST", "{{baseUrl}}/packages/v2/adjust")
	if err != nil {
		var empty models.Package
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Package
	return empty, fmt.Errorf("not found")
}

// Packages_CreateChangeItemV1 (Packages)
// POST {{baseUrl}}/packages/v1/change/item
func (f *Fetcher) Packages_CreateChangeItemV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_CreateChangeItemV1", "POST", "{{baseUrl}}/packages/v1/change/item")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_CreateChangeLocationV1 (Packages)
// POST {{baseUrl}}/packages/v1/change/locations
func (f *Fetcher) Packages_CreateChangeLocationV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_CreateChangeLocationV1", "POST", "{{baseUrl}}/packages/v1/change/locations")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_CreateFinishV1 (Packages)
// POST {{baseUrl}}/packages/v1/finish
func (f *Fetcher) Packages_CreateFinishV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_CreateFinishV1", "POST", "{{baseUrl}}/packages/v1/finish")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_CreatePlantingsV1 (Packages)
// POST {{baseUrl}}/packages/v1/create/plantings
func (f *Fetcher) Packages_CreatePlantingsV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_CreatePlantingsV1", "POST", "{{baseUrl}}/packages/v1/create/plantings")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_CreatePlantingsV2 (Packages)
// POST {{baseUrl}}/packages/v2/plantings
func (f *Fetcher) Packages_CreatePlantingsV2() (models.Package, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Package](f, "Packages_CreatePlantingsV2", "POST", "{{baseUrl}}/packages/v2/plantings")
	if err != nil {
		var empty models.Package
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Package
	return empty, fmt.Errorf("not found")
}

// Packages_CreateRemediateV1 (Packages)
// POST {{baseUrl}}/packages/v1/remediate
func (f *Fetcher) Packages_CreateRemediateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_CreateRemediateV1", "POST", "{{baseUrl}}/packages/v1/remediate")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_CreateTestingV1 (Packages)
// POST {{baseUrl}}/packages/v1/create/testing
func (f *Fetcher) Packages_CreateTestingV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_CreateTestingV1", "POST", "{{baseUrl}}/packages/v1/create/testing")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_CreateTestingV2 (Packages)
// POST {{baseUrl}}/packages/v2/testing
func (f *Fetcher) Packages_CreateTestingV2() (models.Package, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Package](f, "Packages_CreateTestingV2", "POST", "{{baseUrl}}/packages/v2/testing")
	if err != nil {
		var empty models.Package
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Package
	return empty, fmt.Errorf("not found")
}

// Packages_CreateUnfinishV1 (Packages)
// POST {{baseUrl}}/packages/v1/unfinish
func (f *Fetcher) Packages_CreateUnfinishV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_CreateUnfinishV1", "POST", "{{baseUrl}}/packages/v1/unfinish")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_CreateV1 (Packages)
// POST {{baseUrl}}/packages/v1/create
func (f *Fetcher) Packages_CreateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_CreateV1", "POST", "{{baseUrl}}/packages/v1/create")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_CreateV2 (Packages)
// POST {{baseUrl}}/packages/v2
func (f *Fetcher) Packages_CreateV2() (models.Package, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Package](f, "Packages_CreateV2", "POST", "{{baseUrl}}/packages/v2")
	if err != nil {
		var empty models.Package
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Package
	return empty, fmt.Errorf("not found")
}

// Packages_DeleteV2 (Packages)
// DELETE {{baseUrl}}/packages/v2/{id}
func (f *Fetcher) Packages_DeleteV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_DeleteV2", "DELETE", "{{baseUrl}}/packages/v2/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_GetActiveV1 (Packages)
// GET {{baseUrl}}/packages/v1/active
func (f *Fetcher) Packages_GetActiveV1() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetActiveV1", "GET", "{{baseUrl}}/packages/v1/active")
}

// Packages_GetActiveV2 (Packages)
// GET {{baseUrl}}/packages/v2/active
func (f *Fetcher) Packages_GetActiveV2() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetActiveV2", "GET", "{{baseUrl}}/packages/v2/active")
}

// Packages_GetAdjustReasonsV1 (Packages)
// GET {{baseUrl}}/packages/v1/adjust/reasons
func (f *Fetcher) Packages_GetAdjustReasonsV1() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetAdjustReasonsV1", "GET", "{{baseUrl}}/packages/v1/adjust/reasons")
}

// Packages_GetAdjustReasonsV2 (Packages)
// GET {{baseUrl}}/packages/v2/adjust/reasons
func (f *Fetcher) Packages_GetAdjustReasonsV2() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetAdjustReasonsV2", "GET", "{{baseUrl}}/packages/v2/adjust/reasons")
}

// Packages_GetByLabelV1 (Packages)
// GET {{baseUrl}}/packages/v1/{label}
func (f *Fetcher) Packages_GetByLabelV1() (models.Package, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Package](f, "Packages_GetByLabelV1", "GET", "{{baseUrl}}/packages/v1/{label}")
	if err != nil {
		var empty models.Package
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Package
	return empty, fmt.Errorf("not found")
}

// Packages_GetByLabelV2 (Packages)
// GET {{baseUrl}}/packages/v2/{label}
func (f *Fetcher) Packages_GetByLabelV2() (models.Package, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Package](f, "Packages_GetByLabelV2", "GET", "{{baseUrl}}/packages/v2/{label}")
	if err != nil {
		var empty models.Package
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Package
	return empty, fmt.Errorf("not found")
}

// Packages_GetInactiveV1 (Packages)
// GET {{baseUrl}}/packages/v1/inactive
func (f *Fetcher) Packages_GetInactiveV1() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetInactiveV1", "GET", "{{baseUrl}}/packages/v1/inactive")
}

// Packages_GetInactiveV2 (Packages)
// GET {{baseUrl}}/packages/v2/inactive
func (f *Fetcher) Packages_GetInactiveV2() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetInactiveV2", "GET", "{{baseUrl}}/packages/v2/inactive")
}

// Packages_GetIntransitV2 (Packages)
// GET {{baseUrl}}/packages/v2/intransit
func (f *Fetcher) Packages_GetIntransitV2() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetIntransitV2", "GET", "{{baseUrl}}/packages/v2/intransit")
}

// Packages_GetLabsamplesV2 (Packages)
// GET {{baseUrl}}/packages/v2/labsamples
func (f *Fetcher) Packages_GetLabsamplesV2() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetLabsamplesV2", "GET", "{{baseUrl}}/packages/v2/labsamples")
}

// Packages_GetOnholdV1 (Packages)
// GET {{baseUrl}}/packages/v1/onhold
func (f *Fetcher) Packages_GetOnholdV1() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetOnholdV1", "GET", "{{baseUrl}}/packages/v1/onhold")
}

// Packages_GetOnholdV2 (Packages)
// GET {{baseUrl}}/packages/v2/onhold
func (f *Fetcher) Packages_GetOnholdV2() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetOnholdV2", "GET", "{{baseUrl}}/packages/v2/onhold")
}

// Packages_GetSourceHarvestV2 (Packages)
// GET {{baseUrl}}/packages/v2/{id}/source/harvests
func (f *Fetcher) Packages_GetSourceHarvestV2() (models.Package, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Package](f, "Packages_GetSourceHarvestV2", "GET", "{{baseUrl}}/packages/v2/{id}/source/harvests")
	if err != nil {
		var empty models.Package
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Package
	return empty, fmt.Errorf("not found")
}

// Packages_GetTransferredV2 (Packages)
// GET {{baseUrl}}/packages/v2/transferred
func (f *Fetcher) Packages_GetTransferredV2() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetTransferredV2", "GET", "{{baseUrl}}/packages/v2/transferred")
}

// Packages_GetTypesV1 (Packages)
// GET {{baseUrl}}/packages/v1/types
func (f *Fetcher) Packages_GetTypesV1() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetTypesV1", "GET", "{{baseUrl}}/packages/v1/types")
}

// Packages_GetTypesV2 (Packages)
// GET {{baseUrl}}/packages/v2/types
func (f *Fetcher) Packages_GetTypesV2() ([]models.Package, error) {
	return fetchList[models.Package](f, "Packages_GetTypesV2", "GET", "{{baseUrl}}/packages/v2/types")
}

// Packages_GetV1 (Packages)
// GET {{baseUrl}}/packages/v1/{id}
func (f *Fetcher) Packages_GetV1() (models.Package, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Package](f, "Packages_GetV1", "GET", "{{baseUrl}}/packages/v1/{id}")
	if err != nil {
		var empty models.Package
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Package
	return empty, fmt.Errorf("not found")
}

// Packages_GetV2 (Packages)
// GET {{baseUrl}}/packages/v2/{id}
func (f *Fetcher) Packages_GetV2() (models.Package, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Package](f, "Packages_GetV2", "GET", "{{baseUrl}}/packages/v2/{id}")
	if err != nil {
		var empty models.Package
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Package
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateAdjustV2 (Packages)
// PUT {{baseUrl}}/packages/v2/adjust
func (f *Fetcher) Packages_UpdateAdjustV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateAdjustV2", "PUT", "{{baseUrl}}/packages/v2/adjust")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateChangeNoteV1 (Packages)
// PUT {{baseUrl}}/packages/v1/change/note
func (f *Fetcher) Packages_UpdateChangeNoteV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateChangeNoteV1", "PUT", "{{baseUrl}}/packages/v1/change/note")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateDecontaminateV2 (Packages)
// PUT {{baseUrl}}/packages/v2/decontaminate
func (f *Fetcher) Packages_UpdateDecontaminateV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateDecontaminateV2", "PUT", "{{baseUrl}}/packages/v2/decontaminate")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateDonationFlagV2 (Packages)
// PUT {{baseUrl}}/packages/v2/donation/flag
func (f *Fetcher) Packages_UpdateDonationFlagV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateDonationFlagV2", "PUT", "{{baseUrl}}/packages/v2/donation/flag")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateDonationUnflagV2 (Packages)
// PUT {{baseUrl}}/packages/v2/donation/unflag
func (f *Fetcher) Packages_UpdateDonationUnflagV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateDonationUnflagV2", "PUT", "{{baseUrl}}/packages/v2/donation/unflag")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateExternalidV2 (Packages)
// PUT {{baseUrl}}/packages/v2/externalid
func (f *Fetcher) Packages_UpdateExternalidV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateExternalidV2", "PUT", "{{baseUrl}}/packages/v2/externalid")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateFinishV2 (Packages)
// PUT {{baseUrl}}/packages/v2/finish
func (f *Fetcher) Packages_UpdateFinishV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateFinishV2", "PUT", "{{baseUrl}}/packages/v2/finish")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateItemV2 (Packages)
// PUT {{baseUrl}}/packages/v2/item
func (f *Fetcher) Packages_UpdateItemV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateItemV2", "PUT", "{{baseUrl}}/packages/v2/item")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateLabTestRequiredV2 (Packages)
// PUT {{baseUrl}}/packages/v2/labtests/required
func (f *Fetcher) Packages_UpdateLabTestRequiredV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateLabTestRequiredV2", "PUT", "{{baseUrl}}/packages/v2/labtests/required")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateLocationV2 (Packages)
// PUT {{baseUrl}}/packages/v2/location
func (f *Fetcher) Packages_UpdateLocationV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateLocationV2", "PUT", "{{baseUrl}}/packages/v2/location")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateNoteV2 (Packages)
// PUT {{baseUrl}}/packages/v2/note
func (f *Fetcher) Packages_UpdateNoteV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateNoteV2", "PUT", "{{baseUrl}}/packages/v2/note")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateRemediateV2 (Packages)
// PUT {{baseUrl}}/packages/v2/remediate
func (f *Fetcher) Packages_UpdateRemediateV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateRemediateV2", "PUT", "{{baseUrl}}/packages/v2/remediate")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateTradesampleFlagV2 (Packages)
// PUT {{baseUrl}}/packages/v2/tradesample/flag
func (f *Fetcher) Packages_UpdateTradesampleFlagV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateTradesampleFlagV2", "PUT", "{{baseUrl}}/packages/v2/tradesample/flag")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateTradesampleUnflagV2 (Packages)
// PUT {{baseUrl}}/packages/v2/tradesample/unflag
func (f *Fetcher) Packages_UpdateTradesampleUnflagV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateTradesampleUnflagV2", "PUT", "{{baseUrl}}/packages/v2/tradesample/unflag")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateUnfinishV2 (Packages)
// PUT {{baseUrl}}/packages/v2/unfinish
func (f *Fetcher) Packages_UpdateUnfinishV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateUnfinishV2", "PUT", "{{baseUrl}}/packages/v2/unfinish")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Packages_UpdateUsebydateV2 (Packages)
// PUT {{baseUrl}}/packages/v2/usebydate
func (f *Fetcher) Packages_UpdateUsebydateV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Packages_UpdateUsebydateV2", "PUT", "{{baseUrl}}/packages/v2/usebydate")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PatientCheckIns_CreateV1 (Patient Check Ins)
// POST {{baseUrl}}/patient-checkins/v1
func (f *Fetcher) PatientCheckIns_CreateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PatientCheckIns_CreateV1", "POST", "{{baseUrl}}/patient-checkins/v1")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PatientCheckIns_CreateV2 (Patient Check Ins)
// POST {{baseUrl}}/patient-checkins/v2
func (f *Fetcher) PatientCheckIns_CreateV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PatientCheckIns_CreateV2", "POST", "{{baseUrl}}/patient-checkins/v2")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PatientCheckIns_DeleteV1 (Patient Check Ins)
// DELETE {{baseUrl}}/patient-checkins/v1/{id}
func (f *Fetcher) PatientCheckIns_DeleteV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PatientCheckIns_DeleteV1", "DELETE", "{{baseUrl}}/patient-checkins/v1/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PatientCheckIns_DeleteV2 (Patient Check Ins)
// DELETE {{baseUrl}}/patient-checkins/v2/{id}
func (f *Fetcher) PatientCheckIns_DeleteV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PatientCheckIns_DeleteV2", "DELETE", "{{baseUrl}}/patient-checkins/v2/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PatientCheckIns_GetAllV1 (Patient Check Ins)
// GET {{baseUrl}}/patient-checkins/v1
func (f *Fetcher) PatientCheckIns_GetAllV1() ([]models.PatientCheckIn, error) {
	return fetchList[models.PatientCheckIn](f, "PatientCheckIns_GetAllV1", "GET", "{{baseUrl}}/patient-checkins/v1")
}

// PatientCheckIns_GetAllV2 (Patient Check Ins)
// GET {{baseUrl}}/patient-checkins/v2
func (f *Fetcher) PatientCheckIns_GetAllV2() ([]models.PatientCheckIn, error) {
	return fetchList[models.PatientCheckIn](f, "PatientCheckIns_GetAllV2", "GET", "{{baseUrl}}/patient-checkins/v2")
}

// PatientCheckIns_GetLocationsV1 (Patient Check Ins)
// GET {{baseUrl}}/patient-checkins/v1/locations
func (f *Fetcher) PatientCheckIns_GetLocationsV1() ([]models.PatientCheckIn, error) {
	return fetchList[models.PatientCheckIn](f, "PatientCheckIns_GetLocationsV1", "GET", "{{baseUrl}}/patient-checkins/v1/locations")
}

// PatientCheckIns_GetLocationsV2 (Patient Check Ins)
// GET {{baseUrl}}/patient-checkins/v2/locations
func (f *Fetcher) PatientCheckIns_GetLocationsV2() ([]models.PatientCheckIn, error) {
	return fetchList[models.PatientCheckIn](f, "PatientCheckIns_GetLocationsV2", "GET", "{{baseUrl}}/patient-checkins/v2/locations")
}

// PatientCheckIns_UpdateV1 (Patient Check Ins)
// PUT {{baseUrl}}/patient-checkins/v1
func (f *Fetcher) PatientCheckIns_UpdateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PatientCheckIns_UpdateV1", "PUT", "{{baseUrl}}/patient-checkins/v1")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PatientCheckIns_UpdateV2 (Patient Check Ins)
// PUT {{baseUrl}}/patient-checkins/v2
func (f *Fetcher) PatientCheckIns_UpdateV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PatientCheckIns_UpdateV2", "PUT", "{{baseUrl}}/patient-checkins/v2")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PatientsStatus_GetStatusesByPatientLicenseNumberV1 (Patients Status)
// GET {{baseUrl}}/patients/v1/statuses/{patientLicenseNumber}
func (f *Fetcher) PatientsStatus_GetStatusesByPatientLicenseNumberV1() (models.PatientsStatu, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PatientsStatu](f, "PatientsStatus_GetStatusesByPatientLicenseNumberV1", "GET", "{{baseUrl}}/patients/v1/statuses/{patientLicenseNumber}")
	if err != nil {
		var empty models.PatientsStatu
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PatientsStatu
	return empty, fmt.Errorf("not found")
}

// PatientsStatus_GetStatusesByPatientLicenseNumberV2 (Patients Status)
// GET {{baseUrl}}/patients/v2/statuses/{patientLicenseNumber}
func (f *Fetcher) PatientsStatus_GetStatusesByPatientLicenseNumberV2() (models.PatientsStatu, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PatientsStatu](f, "PatientsStatus_GetStatusesByPatientLicenseNumberV2", "GET", "{{baseUrl}}/patients/v2/statuses/{patientLicenseNumber}")
	if err != nil {
		var empty models.PatientsStatu
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PatientsStatu
	return empty, fmt.Errorf("not found")
}

// Patients_CreateAddV1 (Patients)
// POST {{baseUrl}}/patients/v1/add
func (f *Fetcher) Patients_CreateAddV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Patients_CreateAddV1", "POST", "{{baseUrl}}/patients/v1/add")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Patients_CreateUpdateV1 (Patients)
// POST {{baseUrl}}/patients/v1/update
func (f *Fetcher) Patients_CreateUpdateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Patients_CreateUpdateV1", "POST", "{{baseUrl}}/patients/v1/update")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Patients_CreateV2 (Patients)
// POST {{baseUrl}}/patients/v2
func (f *Fetcher) Patients_CreateV2() (models.Patient, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Patient](f, "Patients_CreateV2", "POST", "{{baseUrl}}/patients/v2")
	if err != nil {
		var empty models.Patient
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Patient
	return empty, fmt.Errorf("not found")
}

// Patients_DeleteV1 (Patients)
// DELETE {{baseUrl}}/patients/v1/{id}
func (f *Fetcher) Patients_DeleteV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Patients_DeleteV1", "DELETE", "{{baseUrl}}/patients/v1/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Patients_DeleteV2 (Patients)
// DELETE {{baseUrl}}/patients/v2/{id}
func (f *Fetcher) Patients_DeleteV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Patients_DeleteV2", "DELETE", "{{baseUrl}}/patients/v2/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Patients_GetActiveV1 (Patients)
// GET {{baseUrl}}/patients/v1/active
func (f *Fetcher) Patients_GetActiveV1() ([]models.Patient, error) {
	return fetchList[models.Patient](f, "Patients_GetActiveV1", "GET", "{{baseUrl}}/patients/v1/active")
}

// Patients_GetActiveV2 (Patients)
// GET {{baseUrl}}/patients/v2/active
func (f *Fetcher) Patients_GetActiveV2() ([]models.Patient, error) {
	return fetchList[models.Patient](f, "Patients_GetActiveV2", "GET", "{{baseUrl}}/patients/v2/active")
}

// Patients_GetV1 (Patients)
// GET {{baseUrl}}/patients/v1/{id}
func (f *Fetcher) Patients_GetV1() (models.Patient, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Patient](f, "Patients_GetV1", "GET", "{{baseUrl}}/patients/v1/{id}")
	if err != nil {
		var empty models.Patient
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Patient
	return empty, fmt.Errorf("not found")
}

// Patients_GetV2 (Patients)
// GET {{baseUrl}}/patients/v2/{id}
func (f *Fetcher) Patients_GetV2() (models.Patient, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Patient](f, "Patients_GetV2", "GET", "{{baseUrl}}/patients/v2/{id}")
	if err != nil {
		var empty models.Patient
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Patient
	return empty, fmt.Errorf("not found")
}

// Patients_UpdateV2 (Patients)
// PUT {{baseUrl}}/patients/v2
func (f *Fetcher) Patients_UpdateV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Patients_UpdateV2", "PUT", "{{baseUrl}}/patients/v2")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateAdditivesUsingtemplateV2 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v2/additives/usingtemplate
func (f *Fetcher) PlantBatches_CreateAdditivesUsingtemplateV2() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_CreateAdditivesUsingtemplateV2", "POST", "{{baseUrl}}/plantbatches/v2/additives/usingtemplate")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateAdditivesV1 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v1/additives
func (f *Fetcher) PlantBatches_CreateAdditivesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_CreateAdditivesV1", "POST", "{{baseUrl}}/plantbatches/v1/additives")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateAdditivesV2 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v2/additives
func (f *Fetcher) PlantBatches_CreateAdditivesV2() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_CreateAdditivesV2", "POST", "{{baseUrl}}/plantbatches/v2/additives")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateAdjustV1 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v1/adjust
func (f *Fetcher) PlantBatches_CreateAdjustV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_CreateAdjustV1", "POST", "{{baseUrl}}/plantbatches/v1/adjust")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateAdjustV2 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v2/adjust
func (f *Fetcher) PlantBatches_CreateAdjustV2() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_CreateAdjustV2", "POST", "{{baseUrl}}/plantbatches/v2/adjust")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateChangegrowthphaseV1 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v1/changegrowthphase
func (f *Fetcher) PlantBatches_CreateChangegrowthphaseV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_CreateChangegrowthphaseV1", "POST", "{{baseUrl}}/plantbatches/v1/changegrowthphase")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateGrowthphaseV2 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v2/growthphase
func (f *Fetcher) PlantBatches_CreateGrowthphaseV2() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_CreateGrowthphaseV2", "POST", "{{baseUrl}}/plantbatches/v2/growthphase")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreatePackageFrommotherplantV1 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v1/create/packages/frommotherplant
func (f *Fetcher) PlantBatches_CreatePackageFrommotherplantV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_CreatePackageFrommotherplantV1", "POST", "{{baseUrl}}/plantbatches/v1/create/packages/frommotherplant")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreatePackageFrommotherplantV2 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v2/packages/frommotherplant
func (f *Fetcher) PlantBatches_CreatePackageFrommotherplantV2() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_CreatePackageFrommotherplantV2", "POST", "{{baseUrl}}/plantbatches/v2/packages/frommotherplant")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreatePackageV2 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v2/packages
func (f *Fetcher) PlantBatches_CreatePackageV2() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_CreatePackageV2", "POST", "{{baseUrl}}/plantbatches/v2/packages")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreatePlantingsV2 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v2/plantings
func (f *Fetcher) PlantBatches_CreatePlantingsV2() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_CreatePlantingsV2", "POST", "{{baseUrl}}/plantbatches/v2/plantings")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateSplitV1 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v1/split
func (f *Fetcher) PlantBatches_CreateSplitV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_CreateSplitV1", "POST", "{{baseUrl}}/plantbatches/v1/split")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateSplitV2 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v2/split
func (f *Fetcher) PlantBatches_CreateSplitV2() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_CreateSplitV2", "POST", "{{baseUrl}}/plantbatches/v2/split")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateWasteV1 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v1/waste
func (f *Fetcher) PlantBatches_CreateWasteV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_CreateWasteV1", "POST", "{{baseUrl}}/plantbatches/v1/waste")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateWasteV2 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v2/waste
func (f *Fetcher) PlantBatches_CreateWasteV2() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_CreateWasteV2", "POST", "{{baseUrl}}/plantbatches/v2/waste")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreatepackagesV1 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v1/createpackages
func (f *Fetcher) PlantBatches_CreatepackagesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_CreatepackagesV1", "POST", "{{baseUrl}}/plantbatches/v1/createpackages")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_CreateplantingsV1 (Plant Batches)
// POST {{baseUrl}}/plantbatches/v1/createplantings
func (f *Fetcher) PlantBatches_CreateplantingsV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_CreateplantingsV1", "POST", "{{baseUrl}}/plantbatches/v1/createplantings")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_DeleteV1 (Plant Batches)
// DELETE {{baseUrl}}/plantbatches/v1
func (f *Fetcher) PlantBatches_DeleteV1() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_DeleteV1", "DELETE", "{{baseUrl}}/plantbatches/v1")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_DeleteV2 (Plant Batches)
// DELETE {{baseUrl}}/plantbatches/v2
func (f *Fetcher) PlantBatches_DeleteV2() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_DeleteV2", "DELETE", "{{baseUrl}}/plantbatches/v2")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_GetActiveV1 (Plant Batches)
// GET {{baseUrl}}/plantbatches/v1/active
func (f *Fetcher) PlantBatches_GetActiveV1() ([]models.PlantBatche, error) {
	return fetchList[models.PlantBatche](f, "PlantBatches_GetActiveV1", "GET", "{{baseUrl}}/plantbatches/v1/active")
}

// PlantBatches_GetActiveV2 (Plant Batches)
// GET {{baseUrl}}/plantbatches/v2/active
func (f *Fetcher) PlantBatches_GetActiveV2() ([]models.PlantBatche, error) {
	return fetchList[models.PlantBatche](f, "PlantBatches_GetActiveV2", "GET", "{{baseUrl}}/plantbatches/v2/active")
}

// PlantBatches_GetInactiveV1 (Plant Batches)
// GET {{baseUrl}}/plantbatches/v1/inactive
func (f *Fetcher) PlantBatches_GetInactiveV1() ([]models.PlantBatche, error) {
	return fetchList[models.PlantBatche](f, "PlantBatches_GetInactiveV1", "GET", "{{baseUrl}}/plantbatches/v1/inactive")
}

// PlantBatches_GetInactiveV2 (Plant Batches)
// GET {{baseUrl}}/plantbatches/v2/inactive
func (f *Fetcher) PlantBatches_GetInactiveV2() ([]models.PlantBatche, error) {
	return fetchList[models.PlantBatche](f, "PlantBatches_GetInactiveV2", "GET", "{{baseUrl}}/plantbatches/v2/inactive")
}

// PlantBatches_GetTypesV1 (Plant Batches)
// GET {{baseUrl}}/plantbatches/v1/types
func (f *Fetcher) PlantBatches_GetTypesV1() ([]models.PlantBatche, error) {
	return fetchList[models.PlantBatche](f, "PlantBatches_GetTypesV1", "GET", "{{baseUrl}}/plantbatches/v1/types")
}

// PlantBatches_GetTypesV2 (Plant Batches)
// GET {{baseUrl}}/plantbatches/v2/types
func (f *Fetcher) PlantBatches_GetTypesV2() ([]models.PlantBatche, error) {
	return fetchList[models.PlantBatche](f, "PlantBatches_GetTypesV2", "GET", "{{baseUrl}}/plantbatches/v2/types")
}

// PlantBatches_GetV1 (Plant Batches)
// GET {{baseUrl}}/plantbatches/v1/{id}
func (f *Fetcher) PlantBatches_GetV1() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_GetV1", "GET", "{{baseUrl}}/plantbatches/v1/{id}")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_GetV2 (Plant Batches)
// GET {{baseUrl}}/plantbatches/v2/{id}
func (f *Fetcher) PlantBatches_GetV2() (models.PlantBatche, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.PlantBatche](f, "PlantBatches_GetV2", "GET", "{{baseUrl}}/plantbatches/v2/{id}")
	if err != nil {
		var empty models.PlantBatche
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.PlantBatche
	return empty, fmt.Errorf("not found")
}

// PlantBatches_GetWasteReasonsV1 (Plant Batches)
// GET {{baseUrl}}/plantbatches/v1/waste/reasons
func (f *Fetcher) PlantBatches_GetWasteReasonsV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_GetWasteReasonsV1", "GET", "{{baseUrl}}/plantbatches/v1/waste/reasons")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_GetWasteReasonsV2 (Plant Batches)
// GET {{baseUrl}}/plantbatches/v2/waste/reasons
func (f *Fetcher) PlantBatches_GetWasteReasonsV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_GetWasteReasonsV2", "GET", "{{baseUrl}}/plantbatches/v2/waste/reasons")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_GetWasteV2 (Plant Batches)
// GET {{baseUrl}}/plantbatches/v2/waste
func (f *Fetcher) PlantBatches_GetWasteV2() ([]models.PlantBatche, error) {
	return fetchList[models.PlantBatche](f, "PlantBatches_GetWasteV2", "GET", "{{baseUrl}}/plantbatches/v2/waste")
}

// PlantBatches_UpdateLocationV2 (Plant Batches)
// PUT {{baseUrl}}/plantbatches/v2/location
func (f *Fetcher) PlantBatches_UpdateLocationV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_UpdateLocationV2", "PUT", "{{baseUrl}}/plantbatches/v2/location")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_UpdateMoveplantbatchesV1 (Plant Batches)
// PUT {{baseUrl}}/plantbatches/v1/moveplantbatches
func (f *Fetcher) PlantBatches_UpdateMoveplantbatchesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_UpdateMoveplantbatchesV1", "PUT", "{{baseUrl}}/plantbatches/v1/moveplantbatches")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_UpdateNameV2 (Plant Batches)
// PUT {{baseUrl}}/plantbatches/v2/name
func (f *Fetcher) PlantBatches_UpdateNameV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_UpdateNameV2", "PUT", "{{baseUrl}}/plantbatches/v2/name")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_UpdateStrainV2 (Plant Batches)
// PUT {{baseUrl}}/plantbatches/v2/strain
func (f *Fetcher) PlantBatches_UpdateStrainV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_UpdateStrainV2", "PUT", "{{baseUrl}}/plantbatches/v2/strain")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// PlantBatches_UpdateTagV2 (Plant Batches)
// PUT {{baseUrl}}/plantbatches/v2/tag
func (f *Fetcher) PlantBatches_UpdateTagV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "PlantBatches_UpdateTagV2", "PUT", "{{baseUrl}}/plantbatches/v2/tag")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_CreateAdditivesBylocationUsingtemplateV2 (Plants)
// POST {{baseUrl}}/plants/v2/additives/bylocation/usingtemplate
func (f *Fetcher) Plants_CreateAdditivesBylocationUsingtemplateV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_CreateAdditivesBylocationUsingtemplateV2", "POST", "{{baseUrl}}/plants/v2/additives/bylocation/usingtemplate")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_CreateAdditivesBylocationV1 (Plants)
// POST {{baseUrl}}/plants/v1/additives/bylocation
func (f *Fetcher) Plants_CreateAdditivesBylocationV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_CreateAdditivesBylocationV1", "POST", "{{baseUrl}}/plants/v1/additives/bylocation")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_CreateAdditivesBylocationV2 (Plants)
// POST {{baseUrl}}/plants/v2/additives/bylocation
func (f *Fetcher) Plants_CreateAdditivesBylocationV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_CreateAdditivesBylocationV2", "POST", "{{baseUrl}}/plants/v2/additives/bylocation")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_CreateAdditivesUsingtemplateV2 (Plants)
// POST {{baseUrl}}/plants/v2/additives/usingtemplate
func (f *Fetcher) Plants_CreateAdditivesUsingtemplateV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_CreateAdditivesUsingtemplateV2", "POST", "{{baseUrl}}/plants/v2/additives/usingtemplate")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_CreateAdditivesV1 (Plants)
// POST {{baseUrl}}/plants/v1/additives
func (f *Fetcher) Plants_CreateAdditivesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_CreateAdditivesV1", "POST", "{{baseUrl}}/plants/v1/additives")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_CreateAdditivesV2 (Plants)
// POST {{baseUrl}}/plants/v2/additives
func (f *Fetcher) Plants_CreateAdditivesV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_CreateAdditivesV2", "POST", "{{baseUrl}}/plants/v2/additives")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_CreateChangegrowthphasesV1 (Plants)
// POST {{baseUrl}}/plants/v1/changegrowthphases
func (f *Fetcher) Plants_CreateChangegrowthphasesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_CreateChangegrowthphasesV1", "POST", "{{baseUrl}}/plants/v1/changegrowthphases")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_CreateHarvestplantsV1 (Plants)
// POST {{baseUrl}}/plants/v1/harvestplants
func (f *Fetcher) Plants_CreateHarvestplantsV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_CreateHarvestplantsV1", "POST", "{{baseUrl}}/plants/v1/harvestplants")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_CreateManicureV2 (Plants)
// POST {{baseUrl}}/plants/v2/manicure
func (f *Fetcher) Plants_CreateManicureV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_CreateManicureV2", "POST", "{{baseUrl}}/plants/v2/manicure")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_CreateManicureplantsV1 (Plants)
// POST {{baseUrl}}/plants/v1/manicureplants
func (f *Fetcher) Plants_CreateManicureplantsV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_CreateManicureplantsV1", "POST", "{{baseUrl}}/plants/v1/manicureplants")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_CreateMoveplantsV1 (Plants)
// POST {{baseUrl}}/plants/v1/moveplants
func (f *Fetcher) Plants_CreateMoveplantsV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_CreateMoveplantsV1", "POST", "{{baseUrl}}/plants/v1/moveplants")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_CreatePlantbatchPackageV1 (Plants)
// POST {{baseUrl}}/plants/v1/create/plantbatch/packages
func (f *Fetcher) Plants_CreatePlantbatchPackageV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_CreatePlantbatchPackageV1", "POST", "{{baseUrl}}/plants/v1/create/plantbatch/packages")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_CreatePlantbatchPackageV2 (Plants)
// POST {{baseUrl}}/plants/v2/plantbatch/packages
func (f *Fetcher) Plants_CreatePlantbatchPackageV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_CreatePlantbatchPackageV2", "POST", "{{baseUrl}}/plants/v2/plantbatch/packages")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_CreatePlantingsV1 (Plants)
// POST {{baseUrl}}/plants/v1/create/plantings
func (f *Fetcher) Plants_CreatePlantingsV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_CreatePlantingsV1", "POST", "{{baseUrl}}/plants/v1/create/plantings")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_CreatePlantingsV2 (Plants)
// POST {{baseUrl}}/plants/v2/plantings
func (f *Fetcher) Plants_CreatePlantingsV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_CreatePlantingsV2", "POST", "{{baseUrl}}/plants/v2/plantings")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_CreateWasteV1 (Plants)
// POST {{baseUrl}}/plants/v1/waste
func (f *Fetcher) Plants_CreateWasteV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_CreateWasteV1", "POST", "{{baseUrl}}/plants/v1/waste")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_CreateWasteV2 (Plants)
// POST {{baseUrl}}/plants/v2/waste
func (f *Fetcher) Plants_CreateWasteV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_CreateWasteV2", "POST", "{{baseUrl}}/plants/v2/waste")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_DeleteV1 (Plants)
// DELETE {{baseUrl}}/plants/v1
func (f *Fetcher) Plants_DeleteV1() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_DeleteV1", "DELETE", "{{baseUrl}}/plants/v1")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_DeleteV2 (Plants)
// DELETE {{baseUrl}}/plants/v2
func (f *Fetcher) Plants_DeleteV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_DeleteV2", "DELETE", "{{baseUrl}}/plants/v2")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_GetAdditivesTypesV1 (Plants)
// GET {{baseUrl}}/plants/v1/additives/types
func (f *Fetcher) Plants_GetAdditivesTypesV1() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetAdditivesTypesV1", "GET", "{{baseUrl}}/plants/v1/additives/types")
}

// Plants_GetAdditivesTypesV2 (Plants)
// GET {{baseUrl}}/plants/v2/additives/types
func (f *Fetcher) Plants_GetAdditivesTypesV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetAdditivesTypesV2", "GET", "{{baseUrl}}/plants/v2/additives/types")
}

// Plants_GetAdditivesV1 (Plants)
// GET {{baseUrl}}/plants/v1/additives
func (f *Fetcher) Plants_GetAdditivesV1() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetAdditivesV1", "GET", "{{baseUrl}}/plants/v1/additives")
}

// Plants_GetAdditivesV2 (Plants)
// GET {{baseUrl}}/plants/v2/additives
func (f *Fetcher) Plants_GetAdditivesV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetAdditivesV2", "GET", "{{baseUrl}}/plants/v2/additives")
}

// Plants_GetByLabelV1 (Plants)
// GET {{baseUrl}}/plants/v1/{label}
func (f *Fetcher) Plants_GetByLabelV1() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_GetByLabelV1", "GET", "{{baseUrl}}/plants/v1/{label}")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_GetByLabelV2 (Plants)
// GET {{baseUrl}}/plants/v2/{label}
func (f *Fetcher) Plants_GetByLabelV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_GetByLabelV2", "GET", "{{baseUrl}}/plants/v2/{label}")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_GetFloweringV1 (Plants)
// GET {{baseUrl}}/plants/v1/flowering
func (f *Fetcher) Plants_GetFloweringV1() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetFloweringV1", "GET", "{{baseUrl}}/plants/v1/flowering")
}

// Plants_GetFloweringV2 (Plants)
// GET {{baseUrl}}/plants/v2/flowering
func (f *Fetcher) Plants_GetFloweringV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetFloweringV2", "GET", "{{baseUrl}}/plants/v2/flowering")
}

// Plants_GetGrowthPhasesV1 (Plants)
// GET {{baseUrl}}/plants/v1/growthphases
func (f *Fetcher) Plants_GetGrowthPhasesV1() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetGrowthPhasesV1", "GET", "{{baseUrl}}/plants/v1/growthphases")
}

// Plants_GetGrowthPhasesV2 (Plants)
// GET {{baseUrl}}/plants/v2/growthphases
func (f *Fetcher) Plants_GetGrowthPhasesV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetGrowthPhasesV2", "GET", "{{baseUrl}}/plants/v2/growthphases")
}

// Plants_GetInactiveV1 (Plants)
// GET {{baseUrl}}/plants/v1/inactive
func (f *Fetcher) Plants_GetInactiveV1() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetInactiveV1", "GET", "{{baseUrl}}/plants/v1/inactive")
}

// Plants_GetInactiveV2 (Plants)
// GET {{baseUrl}}/plants/v2/inactive
func (f *Fetcher) Plants_GetInactiveV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetInactiveV2", "GET", "{{baseUrl}}/plants/v2/inactive")
}

// Plants_GetMotherInactiveV2 (Plants)
// GET {{baseUrl}}/plants/v2/mother/inactive
func (f *Fetcher) Plants_GetMotherInactiveV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetMotherInactiveV2", "GET", "{{baseUrl}}/plants/v2/mother/inactive")
}

// Plants_GetMotherOnholdV2 (Plants)
// GET {{baseUrl}}/plants/v2/mother/onhold
func (f *Fetcher) Plants_GetMotherOnholdV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetMotherOnholdV2", "GET", "{{baseUrl}}/plants/v2/mother/onhold")
}

// Plants_GetMotherV2 (Plants)
// GET {{baseUrl}}/plants/v2/mother
func (f *Fetcher) Plants_GetMotherV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetMotherV2", "GET", "{{baseUrl}}/plants/v2/mother")
}

// Plants_GetOnholdV1 (Plants)
// GET {{baseUrl}}/plants/v1/onhold
func (f *Fetcher) Plants_GetOnholdV1() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetOnholdV1", "GET", "{{baseUrl}}/plants/v1/onhold")
}

// Plants_GetOnholdV2 (Plants)
// GET {{baseUrl}}/plants/v2/onhold
func (f *Fetcher) Plants_GetOnholdV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetOnholdV2", "GET", "{{baseUrl}}/plants/v2/onhold")
}

// Plants_GetV1 (Plants)
// GET {{baseUrl}}/plants/v1/{id}
func (f *Fetcher) Plants_GetV1() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_GetV1", "GET", "{{baseUrl}}/plants/v1/{id}")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_GetV2 (Plants)
// GET {{baseUrl}}/plants/v2/{id}
func (f *Fetcher) Plants_GetV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_GetV2", "GET", "{{baseUrl}}/plants/v2/{id}")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_GetVegetativeV1 (Plants)
// GET {{baseUrl}}/plants/v1/vegetative
func (f *Fetcher) Plants_GetVegetativeV1() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetVegetativeV1", "GET", "{{baseUrl}}/plants/v1/vegetative")
}

// Plants_GetVegetativeV2 (Plants)
// GET {{baseUrl}}/plants/v2/vegetative
func (f *Fetcher) Plants_GetVegetativeV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetVegetativeV2", "GET", "{{baseUrl}}/plants/v2/vegetative")
}

// Plants_GetWasteMethodsAllV1 (Plants)
// GET {{baseUrl}}/plants/v1/waste/methods/all
func (f *Fetcher) Plants_GetWasteMethodsAllV1() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetWasteMethodsAllV1", "GET", "{{baseUrl}}/plants/v1/waste/methods/all")
}

// Plants_GetWasteMethodsAllV2 (Plants)
// GET {{baseUrl}}/plants/v2/waste/methods/all
func (f *Fetcher) Plants_GetWasteMethodsAllV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetWasteMethodsAllV2", "GET", "{{baseUrl}}/plants/v2/waste/methods/all")
}

// Plants_GetWastePackageV2 (Plants)
// GET {{baseUrl}}/plants/v2/waste/{id}/package
func (f *Fetcher) Plants_GetWastePackageV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_GetWastePackageV2", "GET", "{{baseUrl}}/plants/v2/waste/{id}/package")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_GetWastePlantV2 (Plants)
// GET {{baseUrl}}/plants/v2/waste/{id}/plant
func (f *Fetcher) Plants_GetWastePlantV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_GetWastePlantV2", "GET", "{{baseUrl}}/plants/v2/waste/{id}/plant")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_GetWasteReasonsV1 (Plants)
// GET {{baseUrl}}/plants/v1/waste/reasons
func (f *Fetcher) Plants_GetWasteReasonsV1() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetWasteReasonsV1", "GET", "{{baseUrl}}/plants/v1/waste/reasons")
}

// Plants_GetWasteReasonsV2 (Plants)
// GET {{baseUrl}}/plants/v2/waste/reasons
func (f *Fetcher) Plants_GetWasteReasonsV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetWasteReasonsV2", "GET", "{{baseUrl}}/plants/v2/waste/reasons")
}

// Plants_GetWasteV2 (Plants)
// GET {{baseUrl}}/plants/v2/waste
func (f *Fetcher) Plants_GetWasteV2() ([]models.Plant, error) {
	return fetchList[models.Plant](f, "Plants_GetWasteV2", "GET", "{{baseUrl}}/plants/v2/waste")
}

// Plants_UpdateAdjustV2 (Plants)
// PUT {{baseUrl}}/plants/v2/adjust
func (f *Fetcher) Plants_UpdateAdjustV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_UpdateAdjustV2", "PUT", "{{baseUrl}}/plants/v2/adjust")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_UpdateGrowthphaseV2 (Plants)
// PUT {{baseUrl}}/plants/v2/growthphase
func (f *Fetcher) Plants_UpdateGrowthphaseV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_UpdateGrowthphaseV2", "PUT", "{{baseUrl}}/plants/v2/growthphase")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_UpdateHarvestV2 (Plants)
// PUT {{baseUrl}}/plants/v2/harvest
func (f *Fetcher) Plants_UpdateHarvestV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_UpdateHarvestV2", "PUT", "{{baseUrl}}/plants/v2/harvest")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_UpdateLocationV2 (Plants)
// PUT {{baseUrl}}/plants/v2/location
func (f *Fetcher) Plants_UpdateLocationV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_UpdateLocationV2", "PUT", "{{baseUrl}}/plants/v2/location")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_UpdateMergeV2 (Plants)
// PUT {{baseUrl}}/plants/v2/merge
func (f *Fetcher) Plants_UpdateMergeV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_UpdateMergeV2", "PUT", "{{baseUrl}}/plants/v2/merge")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_UpdateSplitV2 (Plants)
// PUT {{baseUrl}}/plants/v2/split
func (f *Fetcher) Plants_UpdateSplitV2() (models.Plant, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Plant](f, "Plants_UpdateSplitV2", "PUT", "{{baseUrl}}/plants/v2/split")
	if err != nil {
		var empty models.Plant
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Plant
	return empty, fmt.Errorf("not found")
}

// Plants_UpdateStrainV2 (Plants)
// PUT {{baseUrl}}/plants/v2/strain
func (f *Fetcher) Plants_UpdateStrainV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_UpdateStrainV2", "PUT", "{{baseUrl}}/plants/v2/strain")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Plants_UpdateTagV2 (Plants)
// PUT {{baseUrl}}/plants/v2/tag
func (f *Fetcher) Plants_UpdateTagV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Plants_UpdateTagV2", "PUT", "{{baseUrl}}/plants/v2/tag")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_CreateAdjustV1 (Processing Jobs)
// POST {{baseUrl}}/processing/v1/adjust
func (f *Fetcher) ProcessingJobs_CreateAdjustV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_CreateAdjustV1", "POST", "{{baseUrl}}/processing/v1/adjust")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_CreateAdjustV2 (Processing Jobs)
// POST {{baseUrl}}/processing/v2/adjust
func (f *Fetcher) ProcessingJobs_CreateAdjustV2() (models.ProcessingJob, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.ProcessingJob](f, "ProcessingJobs_CreateAdjustV2", "POST", "{{baseUrl}}/processing/v2/adjust")
	if err != nil {
		var empty models.ProcessingJob
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.ProcessingJob
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_CreateJobtypesV1 (Processing Jobs)
// POST {{baseUrl}}/processing/v1/jobtypes
func (f *Fetcher) ProcessingJobs_CreateJobtypesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_CreateJobtypesV1", "POST", "{{baseUrl}}/processing/v1/jobtypes")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_CreateJobtypesV2 (Processing Jobs)
// POST {{baseUrl}}/processing/v2/jobtypes
func (f *Fetcher) ProcessingJobs_CreateJobtypesV2() (models.ProcessingJob, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.ProcessingJob](f, "ProcessingJobs_CreateJobtypesV2", "POST", "{{baseUrl}}/processing/v2/jobtypes")
	if err != nil {
		var empty models.ProcessingJob
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.ProcessingJob
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_CreateStartV1 (Processing Jobs)
// POST {{baseUrl}}/processing/v1/start
func (f *Fetcher) ProcessingJobs_CreateStartV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_CreateStartV1", "POST", "{{baseUrl}}/processing/v1/start")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_CreateStartV2 (Processing Jobs)
// POST {{baseUrl}}/processing/v2/start
func (f *Fetcher) ProcessingJobs_CreateStartV2() (models.ProcessingJob, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.ProcessingJob](f, "ProcessingJobs_CreateStartV2", "POST", "{{baseUrl}}/processing/v2/start")
	if err != nil {
		var empty models.ProcessingJob
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.ProcessingJob
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_CreatepackagesV1 (Processing Jobs)
// POST {{baseUrl}}/processing/v1/createpackages
func (f *Fetcher) ProcessingJobs_CreatepackagesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_CreatepackagesV1", "POST", "{{baseUrl}}/processing/v1/createpackages")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_CreatepackagesV2 (Processing Jobs)
// POST {{baseUrl}}/processing/v2/createpackages
func (f *Fetcher) ProcessingJobs_CreatepackagesV2() (models.ProcessingJob, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.ProcessingJob](f, "ProcessingJobs_CreatepackagesV2", "POST", "{{baseUrl}}/processing/v2/createpackages")
	if err != nil {
		var empty models.ProcessingJob
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.ProcessingJob
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_DeleteJobtypesV1 (Processing Jobs)
// DELETE {{baseUrl}}/processing/v1/jobtypes/{id}
func (f *Fetcher) ProcessingJobs_DeleteJobtypesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_DeleteJobtypesV1", "DELETE", "{{baseUrl}}/processing/v1/jobtypes/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_DeleteJobtypesV2 (Processing Jobs)
// DELETE {{baseUrl}}/processing/v2/jobtypes/{id}
func (f *Fetcher) ProcessingJobs_DeleteJobtypesV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_DeleteJobtypesV2", "DELETE", "{{baseUrl}}/processing/v2/jobtypes/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_DeleteV1 (Processing Jobs)
// DELETE {{baseUrl}}/processing/v1/{id}
func (f *Fetcher) ProcessingJobs_DeleteV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_DeleteV1", "DELETE", "{{baseUrl}}/processing/v1/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_DeleteV2 (Processing Jobs)
// DELETE {{baseUrl}}/processing/v2/{id}
func (f *Fetcher) ProcessingJobs_DeleteV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_DeleteV2", "DELETE", "{{baseUrl}}/processing/v2/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_GetActiveV1 (Processing Jobs)
// GET {{baseUrl}}/processing/v1/active
func (f *Fetcher) ProcessingJobs_GetActiveV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_GetActiveV1", "GET", "{{baseUrl}}/processing/v1/active")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_GetActiveV2 (Processing Jobs)
// GET {{baseUrl}}/processing/v2/active
func (f *Fetcher) ProcessingJobs_GetActiveV2() ([]models.ProcessingJob, error) {
	return fetchList[models.ProcessingJob](f, "ProcessingJobs_GetActiveV2", "GET", "{{baseUrl}}/processing/v2/active")
}

// ProcessingJobs_GetInactiveV1 (Processing Jobs)
// GET {{baseUrl}}/processing/v1/inactive
func (f *Fetcher) ProcessingJobs_GetInactiveV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_GetInactiveV1", "GET", "{{baseUrl}}/processing/v1/inactive")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_GetInactiveV2 (Processing Jobs)
// GET {{baseUrl}}/processing/v2/inactive
func (f *Fetcher) ProcessingJobs_GetInactiveV2() ([]models.ProcessingJob, error) {
	return fetchList[models.ProcessingJob](f, "ProcessingJobs_GetInactiveV2", "GET", "{{baseUrl}}/processing/v2/inactive")
}

// ProcessingJobs_GetJobtypesActiveV1 (Processing Jobs)
// GET {{baseUrl}}/processing/v1/jobtypes/active
func (f *Fetcher) ProcessingJobs_GetJobtypesActiveV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_GetJobtypesActiveV1", "GET", "{{baseUrl}}/processing/v1/jobtypes/active")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_GetJobtypesActiveV2 (Processing Jobs)
// GET {{baseUrl}}/processing/v2/jobtypes/active
func (f *Fetcher) ProcessingJobs_GetJobtypesActiveV2() ([]models.ProcessingJob, error) {
	return fetchList[models.ProcessingJob](f, "ProcessingJobs_GetJobtypesActiveV2", "GET", "{{baseUrl}}/processing/v2/jobtypes/active")
}

// ProcessingJobs_GetJobtypesAttributesV1 (Processing Jobs)
// GET {{baseUrl}}/processing/v1/jobtypes/attributes
func (f *Fetcher) ProcessingJobs_GetJobtypesAttributesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_GetJobtypesAttributesV1", "GET", "{{baseUrl}}/processing/v1/jobtypes/attributes")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_GetJobtypesAttributesV2 (Processing Jobs)
// GET {{baseUrl}}/processing/v2/jobtypes/attributes
func (f *Fetcher) ProcessingJobs_GetJobtypesAttributesV2() ([]models.ProcessingJob, error) {
	return fetchList[models.ProcessingJob](f, "ProcessingJobs_GetJobtypesAttributesV2", "GET", "{{baseUrl}}/processing/v2/jobtypes/attributes")
}

// ProcessingJobs_GetJobtypesCategoriesV1 (Processing Jobs)
// GET {{baseUrl}}/processing/v1/jobtypes/categories
func (f *Fetcher) ProcessingJobs_GetJobtypesCategoriesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_GetJobtypesCategoriesV1", "GET", "{{baseUrl}}/processing/v1/jobtypes/categories")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_GetJobtypesCategoriesV2 (Processing Jobs)
// GET {{baseUrl}}/processing/v2/jobtypes/categories
func (f *Fetcher) ProcessingJobs_GetJobtypesCategoriesV2() ([]models.ProcessingJob, error) {
	return fetchList[models.ProcessingJob](f, "ProcessingJobs_GetJobtypesCategoriesV2", "GET", "{{baseUrl}}/processing/v2/jobtypes/categories")
}

// ProcessingJobs_GetJobtypesInactiveV1 (Processing Jobs)
// GET {{baseUrl}}/processing/v1/jobtypes/inactive
func (f *Fetcher) ProcessingJobs_GetJobtypesInactiveV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_GetJobtypesInactiveV1", "GET", "{{baseUrl}}/processing/v1/jobtypes/inactive")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_GetJobtypesInactiveV2 (Processing Jobs)
// GET {{baseUrl}}/processing/v2/jobtypes/inactive
func (f *Fetcher) ProcessingJobs_GetJobtypesInactiveV2() ([]models.ProcessingJob, error) {
	return fetchList[models.ProcessingJob](f, "ProcessingJobs_GetJobtypesInactiveV2", "GET", "{{baseUrl}}/processing/v2/jobtypes/inactive")
}

// ProcessingJobs_GetV1 (Processing Jobs)
// GET {{baseUrl}}/processing/v1/{id}
func (f *Fetcher) ProcessingJobs_GetV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_GetV1", "GET", "{{baseUrl}}/processing/v1/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_GetV2 (Processing Jobs)
// GET {{baseUrl}}/processing/v2/{id}
func (f *Fetcher) ProcessingJobs_GetV2() (models.ProcessingJob, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.ProcessingJob](f, "ProcessingJobs_GetV2", "GET", "{{baseUrl}}/processing/v2/{id}")
	if err != nil {
		var empty models.ProcessingJob
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.ProcessingJob
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_UpdateFinishV1 (Processing Jobs)
// PUT {{baseUrl}}/processing/v1/finish
func (f *Fetcher) ProcessingJobs_UpdateFinishV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_UpdateFinishV1", "PUT", "{{baseUrl}}/processing/v1/finish")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_UpdateFinishV2 (Processing Jobs)
// PUT {{baseUrl}}/processing/v2/finish
func (f *Fetcher) ProcessingJobs_UpdateFinishV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_UpdateFinishV2", "PUT", "{{baseUrl}}/processing/v2/finish")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_UpdateJobtypesV1 (Processing Jobs)
// PUT {{baseUrl}}/processing/v1/jobtypes
func (f *Fetcher) ProcessingJobs_UpdateJobtypesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_UpdateJobtypesV1", "PUT", "{{baseUrl}}/processing/v1/jobtypes")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_UpdateJobtypesV2 (Processing Jobs)
// PUT {{baseUrl}}/processing/v2/jobtypes
func (f *Fetcher) ProcessingJobs_UpdateJobtypesV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_UpdateJobtypesV2", "PUT", "{{baseUrl}}/processing/v2/jobtypes")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_UpdateUnfinishV1 (Processing Jobs)
// PUT {{baseUrl}}/processing/v1/unfinish
func (f *Fetcher) ProcessingJobs_UpdateUnfinishV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_UpdateUnfinishV1", "PUT", "{{baseUrl}}/processing/v1/unfinish")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// ProcessingJobs_UpdateUnfinishV2 (Processing Jobs)
// PUT {{baseUrl}}/processing/v2/unfinish
func (f *Fetcher) ProcessingJobs_UpdateUnfinishV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "ProcessingJobs_UpdateUnfinishV2", "PUT", "{{baseUrl}}/processing/v2/unfinish")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// RetailId_CreateAssociateV2 (Retail Id)
// POST {{baseUrl}}/retailid/v2/associate
func (f *Fetcher) RetailId_CreateAssociateV2() (models.RetailId, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.RetailId](f, "RetailId_CreateAssociateV2", "POST", "{{baseUrl}}/retailid/v2/associate")
	if err != nil {
		var empty models.RetailId
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.RetailId
	return empty, fmt.Errorf("not found")
}

// RetailId_CreateGenerateV2 (Retail Id)
// POST {{baseUrl}}/retailid/v2/generate
func (f *Fetcher) RetailId_CreateGenerateV2() (models.RetailId, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.RetailId](f, "RetailId_CreateGenerateV2", "POST", "{{baseUrl}}/retailid/v2/generate")
	if err != nil {
		var empty models.RetailId
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.RetailId
	return empty, fmt.Errorf("not found")
}

// RetailId_CreateMergeV2 (Retail Id)
// POST {{baseUrl}}/retailid/v2/merge
func (f *Fetcher) RetailId_CreateMergeV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "RetailId_CreateMergeV2", "POST", "{{baseUrl}}/retailid/v2/merge")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// RetailId_CreatePackageInfoV2 (Retail Id)
// POST {{baseUrl}}/retailid/v2/packages/info
func (f *Fetcher) RetailId_CreatePackageInfoV2() (models.RetailId, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.RetailId](f, "RetailId_CreatePackageInfoV2", "POST", "{{baseUrl}}/retailid/v2/packages/info")
	if err != nil {
		var empty models.RetailId
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.RetailId
	return empty, fmt.Errorf("not found")
}

// RetailId_GetReceiveByLabelV2 (Retail Id)
// GET {{baseUrl}}/retailid/v2/receive/{label}
func (f *Fetcher) RetailId_GetReceiveByLabelV2() (models.RetailId, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.RetailId](f, "RetailId_GetReceiveByLabelV2", "GET", "{{baseUrl}}/retailid/v2/receive/{label}")
	if err != nil {
		var empty models.RetailId
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.RetailId
	return empty, fmt.Errorf("not found")
}

// RetailId_GetReceiveQrByShortCodeV2 (Retail Id)
// GET {{baseUrl}}/retailid/v2/receive/qr/{shortCode}
func (f *Fetcher) RetailId_GetReceiveQrByShortCodeV2() (models.RetailId, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.RetailId](f, "RetailId_GetReceiveQrByShortCodeV2", "GET", "{{baseUrl}}/retailid/v2/receive/qr/{shortCode}")
	if err != nil {
		var empty models.RetailId
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.RetailId
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryRetailerDepartV1 (Sales)
// POST {{baseUrl}}/sales/v1/deliveries/retailer/depart
func (f *Fetcher) Sales_CreateDeliveryRetailerDepartV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_CreateDeliveryRetailerDepartV1", "POST", "{{baseUrl}}/sales/v1/deliveries/retailer/depart")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryRetailerDepartV2 (Sales)
// POST {{baseUrl}}/sales/v2/deliveries/retailer/depart
func (f *Fetcher) Sales_CreateDeliveryRetailerDepartV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_CreateDeliveryRetailerDepartV2", "POST", "{{baseUrl}}/sales/v2/deliveries/retailer/depart")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryRetailerEndV1 (Sales)
// POST {{baseUrl}}/sales/v1/deliveries/retailer/end
func (f *Fetcher) Sales_CreateDeliveryRetailerEndV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_CreateDeliveryRetailerEndV1", "POST", "{{baseUrl}}/sales/v1/deliveries/retailer/end")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryRetailerEndV2 (Sales)
// POST {{baseUrl}}/sales/v2/deliveries/retailer/end
func (f *Fetcher) Sales_CreateDeliveryRetailerEndV2() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_CreateDeliveryRetailerEndV2", "POST", "{{baseUrl}}/sales/v2/deliveries/retailer/end")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryRetailerRestockV1 (Sales)
// POST {{baseUrl}}/sales/v1/deliveries/retailer/restock
func (f *Fetcher) Sales_CreateDeliveryRetailerRestockV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_CreateDeliveryRetailerRestockV1", "POST", "{{baseUrl}}/sales/v1/deliveries/retailer/restock")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryRetailerRestockV2 (Sales)
// POST {{baseUrl}}/sales/v2/deliveries/retailer/restock
func (f *Fetcher) Sales_CreateDeliveryRetailerRestockV2() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_CreateDeliveryRetailerRestockV2", "POST", "{{baseUrl}}/sales/v2/deliveries/retailer/restock")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryRetailerSaleV1 (Sales)
// POST {{baseUrl}}/sales/v1/deliveries/retailer/sale
func (f *Fetcher) Sales_CreateDeliveryRetailerSaleV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_CreateDeliveryRetailerSaleV1", "POST", "{{baseUrl}}/sales/v1/deliveries/retailer/sale")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryRetailerSaleV2 (Sales)
// POST {{baseUrl}}/sales/v2/deliveries/retailer/sale
func (f *Fetcher) Sales_CreateDeliveryRetailerSaleV2() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_CreateDeliveryRetailerSaleV2", "POST", "{{baseUrl}}/sales/v2/deliveries/retailer/sale")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryRetailerV1 (Sales)
// POST {{baseUrl}}/sales/v1/deliveries/retailer
func (f *Fetcher) Sales_CreateDeliveryRetailerV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_CreateDeliveryRetailerV1", "POST", "{{baseUrl}}/sales/v1/deliveries/retailer")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryRetailerV2 (Sales)
// POST {{baseUrl}}/sales/v2/deliveries/retailer
func (f *Fetcher) Sales_CreateDeliveryRetailerV2() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_CreateDeliveryRetailerV2", "POST", "{{baseUrl}}/sales/v2/deliveries/retailer")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryV1 (Sales)
// POST {{baseUrl}}/sales/v1/deliveries
func (f *Fetcher) Sales_CreateDeliveryV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_CreateDeliveryV1", "POST", "{{baseUrl}}/sales/v1/deliveries")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_CreateDeliveryV2 (Sales)
// POST {{baseUrl}}/sales/v2/deliveries
func (f *Fetcher) Sales_CreateDeliveryV2() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_CreateDeliveryV2", "POST", "{{baseUrl}}/sales/v2/deliveries")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_CreateReceiptV1 (Sales)
// POST {{baseUrl}}/sales/v1/receipts
func (f *Fetcher) Sales_CreateReceiptV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_CreateReceiptV1", "POST", "{{baseUrl}}/sales/v1/receipts")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_CreateReceiptV2 (Sales)
// POST {{baseUrl}}/sales/v2/receipts
func (f *Fetcher) Sales_CreateReceiptV2() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_CreateReceiptV2", "POST", "{{baseUrl}}/sales/v2/receipts")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_CreateTransactionByDateV1 (Sales)
// POST {{baseUrl}}/sales/v1/transactions/{date}
func (f *Fetcher) Sales_CreateTransactionByDateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_CreateTransactionByDateV1", "POST", "{{baseUrl}}/sales/v1/transactions/{date}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_DeleteDeliveryRetailerV1 (Sales)
// DELETE {{baseUrl}}/sales/v1/deliveries/retailer/{id}
func (f *Fetcher) Sales_DeleteDeliveryRetailerV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_DeleteDeliveryRetailerV1", "DELETE", "{{baseUrl}}/sales/v1/deliveries/retailer/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_DeleteDeliveryRetailerV2 (Sales)
// DELETE {{baseUrl}}/sales/v2/deliveries/retailer/{id}
func (f *Fetcher) Sales_DeleteDeliveryRetailerV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_DeleteDeliveryRetailerV2", "DELETE", "{{baseUrl}}/sales/v2/deliveries/retailer/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_DeleteDeliveryV1 (Sales)
// DELETE {{baseUrl}}/sales/v1/deliveries/{id}
func (f *Fetcher) Sales_DeleteDeliveryV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_DeleteDeliveryV1", "DELETE", "{{baseUrl}}/sales/v1/deliveries/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_DeleteDeliveryV2 (Sales)
// DELETE {{baseUrl}}/sales/v2/deliveries/{id}
func (f *Fetcher) Sales_DeleteDeliveryV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_DeleteDeliveryV2", "DELETE", "{{baseUrl}}/sales/v2/deliveries/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_DeleteReceiptV1 (Sales)
// DELETE {{baseUrl}}/sales/v1/receipts/{id}
func (f *Fetcher) Sales_DeleteReceiptV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_DeleteReceiptV1", "DELETE", "{{baseUrl}}/sales/v1/receipts/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_DeleteReceiptV2 (Sales)
// DELETE {{baseUrl}}/sales/v2/receipts/{id}
func (f *Fetcher) Sales_DeleteReceiptV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_DeleteReceiptV2", "DELETE", "{{baseUrl}}/sales/v2/receipts/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_GetCountiesV1 (Sales)
// GET {{baseUrl}}/sales/v1/counties
func (f *Fetcher) Sales_GetCountiesV1() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetCountiesV1", "GET", "{{baseUrl}}/sales/v1/counties")
}

// Sales_GetCountiesV2 (Sales)
// GET {{baseUrl}}/sales/v2/counties
func (f *Fetcher) Sales_GetCountiesV2() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetCountiesV2", "GET", "{{baseUrl}}/sales/v2/counties")
}

// Sales_GetCustomertypesV1 (Sales)
// GET {{baseUrl}}/sales/v1/customertypes
func (f *Fetcher) Sales_GetCustomertypesV1() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetCustomertypesV1", "GET", "{{baseUrl}}/sales/v1/customertypes")
}

// Sales_GetCustomertypesV2 (Sales)
// GET {{baseUrl}}/sales/v2/customertypes
func (f *Fetcher) Sales_GetCustomertypesV2() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetCustomertypesV2", "GET", "{{baseUrl}}/sales/v2/customertypes")
}

// Sales_GetDeliveriesActiveV1 (Sales)
// GET {{baseUrl}}/sales/v1/deliveries/active
func (f *Fetcher) Sales_GetDeliveriesActiveV1() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetDeliveriesActiveV1", "GET", "{{baseUrl}}/sales/v1/deliveries/active")
}

// Sales_GetDeliveriesActiveV2 (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/active
func (f *Fetcher) Sales_GetDeliveriesActiveV2() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetDeliveriesActiveV2", "GET", "{{baseUrl}}/sales/v2/deliveries/active")
}

// Sales_GetDeliveriesInactiveV1 (Sales)
// GET {{baseUrl}}/sales/v1/deliveries/inactive
func (f *Fetcher) Sales_GetDeliveriesInactiveV1() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetDeliveriesInactiveV1", "GET", "{{baseUrl}}/sales/v1/deliveries/inactive")
}

// Sales_GetDeliveriesInactiveV2 (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/inactive
func (f *Fetcher) Sales_GetDeliveriesInactiveV2() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetDeliveriesInactiveV2", "GET", "{{baseUrl}}/sales/v2/deliveries/inactive")
}

// Sales_GetDeliveriesRetailerActiveV1 (Sales)
// GET {{baseUrl}}/sales/v1/deliveries/retailer/active
func (f *Fetcher) Sales_GetDeliveriesRetailerActiveV1() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetDeliveriesRetailerActiveV1", "GET", "{{baseUrl}}/sales/v1/deliveries/retailer/active")
}

// Sales_GetDeliveriesRetailerActiveV2 (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/retailer/active
func (f *Fetcher) Sales_GetDeliveriesRetailerActiveV2() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetDeliveriesRetailerActiveV2", "GET", "{{baseUrl}}/sales/v2/deliveries/retailer/active")
}

// Sales_GetDeliveriesRetailerInactiveV1 (Sales)
// GET {{baseUrl}}/sales/v1/deliveries/retailer/inactive
func (f *Fetcher) Sales_GetDeliveriesRetailerInactiveV1() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetDeliveriesRetailerInactiveV1", "GET", "{{baseUrl}}/sales/v1/deliveries/retailer/inactive")
}

// Sales_GetDeliveriesRetailerInactiveV2 (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/retailer/inactive
func (f *Fetcher) Sales_GetDeliveriesRetailerInactiveV2() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetDeliveriesRetailerInactiveV2", "GET", "{{baseUrl}}/sales/v2/deliveries/retailer/inactive")
}

// Sales_GetDeliveriesReturnreasonsV1 (Sales)
// GET {{baseUrl}}/sales/v1/deliveries/returnreasons
func (f *Fetcher) Sales_GetDeliveriesReturnreasonsV1() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetDeliveriesReturnreasonsV1", "GET", "{{baseUrl}}/sales/v1/deliveries/returnreasons")
}

// Sales_GetDeliveriesReturnreasonsV2 (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/returnreasons
func (f *Fetcher) Sales_GetDeliveriesReturnreasonsV2() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetDeliveriesReturnreasonsV2", "GET", "{{baseUrl}}/sales/v2/deliveries/returnreasons")
}

// Sales_GetDeliveryRetailerV1 (Sales)
// GET {{baseUrl}}/sales/v1/deliveries/retailer/{id}
func (f *Fetcher) Sales_GetDeliveryRetailerV1() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_GetDeliveryRetailerV1", "GET", "{{baseUrl}}/sales/v1/deliveries/retailer/{id}")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_GetDeliveryRetailerV2 (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/retailer/{id}
func (f *Fetcher) Sales_GetDeliveryRetailerV2() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_GetDeliveryRetailerV2", "GET", "{{baseUrl}}/sales/v2/deliveries/retailer/{id}")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_GetDeliveryV1 (Sales)
// GET {{baseUrl}}/sales/v1/deliveries/{id}
func (f *Fetcher) Sales_GetDeliveryV1() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_GetDeliveryV1", "GET", "{{baseUrl}}/sales/v1/deliveries/{id}")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_GetDeliveryV2 (Sales)
// GET {{baseUrl}}/sales/v2/deliveries/{id}
func (f *Fetcher) Sales_GetDeliveryV2() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_GetDeliveryV2", "GET", "{{baseUrl}}/sales/v2/deliveries/{id}")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_GetPatientRegistrationsLocationsV1 (Sales)
// GET {{baseUrl}}/sales/v1/patientregistration/locations
func (f *Fetcher) Sales_GetPatientRegistrationsLocationsV1() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetPatientRegistrationsLocationsV1", "GET", "{{baseUrl}}/sales/v1/patientregistration/locations")
}

// Sales_GetPatientRegistrationsLocationsV2 (Sales)
// GET {{baseUrl}}/sales/v2/patientregistration/locations
func (f *Fetcher) Sales_GetPatientRegistrationsLocationsV2() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetPatientRegistrationsLocationsV2", "GET", "{{baseUrl}}/sales/v2/patientregistration/locations")
}

// Sales_GetPaymenttypesV1 (Sales)
// GET {{baseUrl}}/sales/v1/paymenttypes
func (f *Fetcher) Sales_GetPaymenttypesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_GetPaymenttypesV1", "GET", "{{baseUrl}}/sales/v1/paymenttypes")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_GetPaymenttypesV2 (Sales)
// GET {{baseUrl}}/sales/v2/paymenttypes
func (f *Fetcher) Sales_GetPaymenttypesV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_GetPaymenttypesV2", "GET", "{{baseUrl}}/sales/v2/paymenttypes")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_GetReceiptV1 (Sales)
// GET {{baseUrl}}/sales/v1/receipts/{id}
func (f *Fetcher) Sales_GetReceiptV1() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_GetReceiptV1", "GET", "{{baseUrl}}/sales/v1/receipts/{id}")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_GetReceiptV2 (Sales)
// GET {{baseUrl}}/sales/v2/receipts/{id}
func (f *Fetcher) Sales_GetReceiptV2() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_GetReceiptV2", "GET", "{{baseUrl}}/sales/v2/receipts/{id}")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_GetReceiptsActiveV1 (Sales)
// GET {{baseUrl}}/sales/v1/receipts/active
func (f *Fetcher) Sales_GetReceiptsActiveV1() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetReceiptsActiveV1", "GET", "{{baseUrl}}/sales/v1/receipts/active")
}

// Sales_GetReceiptsActiveV2 (Sales)
// GET {{baseUrl}}/sales/v2/receipts/active
func (f *Fetcher) Sales_GetReceiptsActiveV2() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetReceiptsActiveV2", "GET", "{{baseUrl}}/sales/v2/receipts/active")
}

// Sales_GetReceiptsExternalByExternalNumberV2 (Sales)
// GET {{baseUrl}}/sales/v2/receipts/external/{externalNumber}
func (f *Fetcher) Sales_GetReceiptsExternalByExternalNumberV2() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_GetReceiptsExternalByExternalNumberV2", "GET", "{{baseUrl}}/sales/v2/receipts/external/{externalNumber}")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_GetReceiptsInactiveV1 (Sales)
// GET {{baseUrl}}/sales/v1/receipts/inactive
func (f *Fetcher) Sales_GetReceiptsInactiveV1() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetReceiptsInactiveV1", "GET", "{{baseUrl}}/sales/v1/receipts/inactive")
}

// Sales_GetReceiptsInactiveV2 (Sales)
// GET {{baseUrl}}/sales/v2/receipts/inactive
func (f *Fetcher) Sales_GetReceiptsInactiveV2() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetReceiptsInactiveV2", "GET", "{{baseUrl}}/sales/v2/receipts/inactive")
}

// Sales_GetTransactionsBySalesDateStartAndSalesDateEndV1 (Sales)
// GET {{baseUrl}}/sales/v1/transactions/{salesDateStart}/{salesDateEnd}
func (f *Fetcher) Sales_GetTransactionsBySalesDateStartAndSalesDateEndV1() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_GetTransactionsBySalesDateStartAndSalesDateEndV1", "GET", "{{baseUrl}}/sales/v1/transactions/{salesDateStart}/{salesDateEnd}")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_GetTransactionsV1 (Sales)
// GET {{baseUrl}}/sales/v1/transactions
func (f *Fetcher) Sales_GetTransactionsV1() ([]models.Sale, error) {
	return fetchList[models.Sale](f, "Sales_GetTransactionsV1", "GET", "{{baseUrl}}/sales/v1/transactions")
}

// Sales_UpdateDeliveryCompleteV1 (Sales)
// PUT {{baseUrl}}/sales/v1/deliveries/complete
func (f *Fetcher) Sales_UpdateDeliveryCompleteV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryCompleteV1", "PUT", "{{baseUrl}}/sales/v1/deliveries/complete")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryCompleteV2 (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/complete
func (f *Fetcher) Sales_UpdateDeliveryCompleteV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryCompleteV2", "PUT", "{{baseUrl}}/sales/v2/deliveries/complete")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryHubAcceptV1 (Sales)
// PUT {{baseUrl}}/sales/v1/deliveries/hub/accept
func (f *Fetcher) Sales_UpdateDeliveryHubAcceptV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryHubAcceptV1", "PUT", "{{baseUrl}}/sales/v1/deliveries/hub/accept")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryHubAcceptV2 (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/hub/accept
func (f *Fetcher) Sales_UpdateDeliveryHubAcceptV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryHubAcceptV2", "PUT", "{{baseUrl}}/sales/v2/deliveries/hub/accept")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryHubDepartV1 (Sales)
// PUT {{baseUrl}}/sales/v1/deliveries/hub/depart
func (f *Fetcher) Sales_UpdateDeliveryHubDepartV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryHubDepartV1", "PUT", "{{baseUrl}}/sales/v1/deliveries/hub/depart")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryHubDepartV2 (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/hub/depart
func (f *Fetcher) Sales_UpdateDeliveryHubDepartV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryHubDepartV2", "PUT", "{{baseUrl}}/sales/v2/deliveries/hub/depart")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryHubV1 (Sales)
// PUT {{baseUrl}}/sales/v1/deliveries/hub
func (f *Fetcher) Sales_UpdateDeliveryHubV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryHubV1", "PUT", "{{baseUrl}}/sales/v1/deliveries/hub")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryHubV2 (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/hub
func (f *Fetcher) Sales_UpdateDeliveryHubV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryHubV2", "PUT", "{{baseUrl}}/sales/v2/deliveries/hub")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryHubVerifyIDV1 (Sales)
// PUT {{baseUrl}}/sales/v1/deliveries/hub/verifyID
func (f *Fetcher) Sales_UpdateDeliveryHubVerifyIDV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryHubVerifyIDV1", "PUT", "{{baseUrl}}/sales/v1/deliveries/hub/verifyID")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryHubVerifyIDV2 (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/hub/verifyID
func (f *Fetcher) Sales_UpdateDeliveryHubVerifyIDV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryHubVerifyIDV2", "PUT", "{{baseUrl}}/sales/v2/deliveries/hub/verifyID")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryRetailerV1 (Sales)
// PUT {{baseUrl}}/sales/v1/deliveries/retailer
func (f *Fetcher) Sales_UpdateDeliveryRetailerV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryRetailerV1", "PUT", "{{baseUrl}}/sales/v1/deliveries/retailer")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryRetailerV2 (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries/retailer
func (f *Fetcher) Sales_UpdateDeliveryRetailerV2() (models.Sale, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sale](f, "Sales_UpdateDeliveryRetailerV2", "PUT", "{{baseUrl}}/sales/v2/deliveries/retailer")
	if err != nil {
		var empty models.Sale
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sale
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryV1 (Sales)
// PUT {{baseUrl}}/sales/v1/deliveries
func (f *Fetcher) Sales_UpdateDeliveryV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryV1", "PUT", "{{baseUrl}}/sales/v1/deliveries")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateDeliveryV2 (Sales)
// PUT {{baseUrl}}/sales/v2/deliveries
func (f *Fetcher) Sales_UpdateDeliveryV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateDeliveryV2", "PUT", "{{baseUrl}}/sales/v2/deliveries")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateReceiptFinalizeV2 (Sales)
// PUT {{baseUrl}}/sales/v2/receipts/finalize
func (f *Fetcher) Sales_UpdateReceiptFinalizeV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateReceiptFinalizeV2", "PUT", "{{baseUrl}}/sales/v2/receipts/finalize")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateReceiptUnfinalizeV2 (Sales)
// PUT {{baseUrl}}/sales/v2/receipts/unfinalize
func (f *Fetcher) Sales_UpdateReceiptUnfinalizeV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateReceiptUnfinalizeV2", "PUT", "{{baseUrl}}/sales/v2/receipts/unfinalize")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateReceiptV1 (Sales)
// PUT {{baseUrl}}/sales/v1/receipts
func (f *Fetcher) Sales_UpdateReceiptV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateReceiptV1", "PUT", "{{baseUrl}}/sales/v1/receipts")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateReceiptV2 (Sales)
// PUT {{baseUrl}}/sales/v2/receipts
func (f *Fetcher) Sales_UpdateReceiptV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateReceiptV2", "PUT", "{{baseUrl}}/sales/v2/receipts")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sales_UpdateTransactionByDateV1 (Sales)
// PUT {{baseUrl}}/sales/v1/transactions/{date}
func (f *Fetcher) Sales_UpdateTransactionByDateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sales_UpdateTransactionByDateV1", "PUT", "{{baseUrl}}/sales/v1/transactions/{date}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sandbox_CreateIntegratorSetupV2 (Sandbox)
// POST {{baseUrl}}/sandbox/v2/integrator/setup
func (f *Fetcher) Sandbox_CreateIntegratorSetupV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sandbox_CreateIntegratorSetupV2", "POST", "{{baseUrl}}/sandbox/v2/integrator/setup")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Strains_CreateUpdateV1 (Strains)
// POST {{baseUrl}}/strains/v1/update
func (f *Fetcher) Strains_CreateUpdateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Strains_CreateUpdateV1", "POST", "{{baseUrl}}/strains/v1/update")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Strains_CreateV1 (Strains)
// POST {{baseUrl}}/strains/v1/create
func (f *Fetcher) Strains_CreateV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Strains_CreateV1", "POST", "{{baseUrl}}/strains/v1/create")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Strains_CreateV2 (Strains)
// POST {{baseUrl}}/strains/v2
func (f *Fetcher) Strains_CreateV2() (models.Strain, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Strain](f, "Strains_CreateV2", "POST", "{{baseUrl}}/strains/v2")
	if err != nil {
		var empty models.Strain
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Strain
	return empty, fmt.Errorf("not found")
}

// Strains_DeleteV1 (Strains)
// DELETE {{baseUrl}}/strains/v1/{id}
func (f *Fetcher) Strains_DeleteV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Strains_DeleteV1", "DELETE", "{{baseUrl}}/strains/v1/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Strains_DeleteV2 (Strains)
// DELETE {{baseUrl}}/strains/v2/{id}
func (f *Fetcher) Strains_DeleteV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Strains_DeleteV2", "DELETE", "{{baseUrl}}/strains/v2/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Strains_GetActiveV1 (Strains)
// GET {{baseUrl}}/strains/v1/active
func (f *Fetcher) Strains_GetActiveV1() ([]models.Strain, error) {
	return fetchList[models.Strain](f, "Strains_GetActiveV1", "GET", "{{baseUrl}}/strains/v1/active")
}

// Strains_GetActiveV2 (Strains)
// GET {{baseUrl}}/strains/v2/active
func (f *Fetcher) Strains_GetActiveV2() ([]models.Strain, error) {
	return fetchList[models.Strain](f, "Strains_GetActiveV2", "GET", "{{baseUrl}}/strains/v2/active")
}

// Strains_GetInactiveV2 (Strains)
// GET {{baseUrl}}/strains/v2/inactive
func (f *Fetcher) Strains_GetInactiveV2() ([]models.Strain, error) {
	return fetchList[models.Strain](f, "Strains_GetInactiveV2", "GET", "{{baseUrl}}/strains/v2/inactive")
}

// Strains_GetV1 (Strains)
// GET {{baseUrl}}/strains/v1/{id}
func (f *Fetcher) Strains_GetV1() (models.Strain, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Strain](f, "Strains_GetV1", "GET", "{{baseUrl}}/strains/v1/{id}")
	if err != nil {
		var empty models.Strain
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Strain
	return empty, fmt.Errorf("not found")
}

// Strains_GetV2 (Strains)
// GET {{baseUrl}}/strains/v2/{id}
func (f *Fetcher) Strains_GetV2() (models.Strain, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Strain](f, "Strains_GetV2", "GET", "{{baseUrl}}/strains/v2/{id}")
	if err != nil {
		var empty models.Strain
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Strain
	return empty, fmt.Errorf("not found")
}

// Strains_UpdateV2 (Strains)
// PUT {{baseUrl}}/strains/v2
func (f *Fetcher) Strains_UpdateV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Strains_UpdateV2", "PUT", "{{baseUrl}}/strains/v2")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sublocations_CreateV2 (Sublocations)
// POST {{baseUrl}}/sublocations/v2
func (f *Fetcher) Sublocations_CreateV2() (models.Sublocation, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sublocation](f, "Sublocations_CreateV2", "POST", "{{baseUrl}}/sublocations/v2")
	if err != nil {
		var empty models.Sublocation
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sublocation
	return empty, fmt.Errorf("not found")
}

// Sublocations_DeleteV2 (Sublocations)
// DELETE {{baseUrl}}/sublocations/v2/{id}
func (f *Fetcher) Sublocations_DeleteV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sublocations_DeleteV2", "DELETE", "{{baseUrl}}/sublocations/v2/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Sublocations_GetActiveV2 (Sublocations)
// GET {{baseUrl}}/sublocations/v2/active
func (f *Fetcher) Sublocations_GetActiveV2() ([]models.Sublocation, error) {
	return fetchList[models.Sublocation](f, "Sublocations_GetActiveV2", "GET", "{{baseUrl}}/sublocations/v2/active")
}

// Sublocations_GetInactiveV2 (Sublocations)
// GET {{baseUrl}}/sublocations/v2/inactive
func (f *Fetcher) Sublocations_GetInactiveV2() ([]models.Sublocation, error) {
	return fetchList[models.Sublocation](f, "Sublocations_GetInactiveV2", "GET", "{{baseUrl}}/sublocations/v2/inactive")
}

// Sublocations_GetV2 (Sublocations)
// GET {{baseUrl}}/sublocations/v2/{id}
func (f *Fetcher) Sublocations_GetV2() (models.Sublocation, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Sublocation](f, "Sublocations_GetV2", "GET", "{{baseUrl}}/sublocations/v2/{id}")
	if err != nil {
		var empty models.Sublocation
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Sublocation
	return empty, fmt.Errorf("not found")
}

// Sublocations_UpdateV2 (Sublocations)
// PUT {{baseUrl}}/sublocations/v2
func (f *Fetcher) Sublocations_UpdateV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Sublocations_UpdateV2", "PUT", "{{baseUrl}}/sublocations/v2")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Tags_GetPackageAvailableV2 (Tags)
// GET {{baseUrl}}/tags/v2/package/available
func (f *Fetcher) Tags_GetPackageAvailableV2() ([]models.Tag, error) {
	return fetchList[models.Tag](f, "Tags_GetPackageAvailableV2", "GET", "{{baseUrl}}/tags/v2/package/available")
}

// Tags_GetPlantAvailableV2 (Tags)
// GET {{baseUrl}}/tags/v2/plant/available
func (f *Fetcher) Tags_GetPlantAvailableV2() ([]models.Tag, error) {
	return fetchList[models.Tag](f, "Tags_GetPlantAvailableV2", "GET", "{{baseUrl}}/tags/v2/plant/available")
}

// Tags_GetStagedV2 (Tags)
// GET {{baseUrl}}/tags/v2/staged
func (f *Fetcher) Tags_GetStagedV2() ([]models.Tag, error) {
	return fetchList[models.Tag](f, "Tags_GetStagedV2", "GET", "{{baseUrl}}/tags/v2/staged")
}

// Transfers_CreateExternalIncomingV1 (Transfers)
// POST {{baseUrl}}/transfers/v1/external/incoming
func (f *Fetcher) Transfers_CreateExternalIncomingV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transfers_CreateExternalIncomingV1", "POST", "{{baseUrl}}/transfers/v1/external/incoming")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transfers_CreateExternalIncomingV2 (Transfers)
// POST {{baseUrl}}/transfers/v2/external/incoming
func (f *Fetcher) Transfers_CreateExternalIncomingV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_CreateExternalIncomingV2", "POST", "{{baseUrl}}/transfers/v2/external/incoming")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_CreateTemplatesOutgoingV2 (Transfers)
// POST {{baseUrl}}/transfers/v2/templates/outgoing
func (f *Fetcher) Transfers_CreateTemplatesOutgoingV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_CreateTemplatesOutgoingV2", "POST", "{{baseUrl}}/transfers/v2/templates/outgoing")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_CreateTemplatesV1 (Transfers)
// POST {{baseUrl}}/transfers/v1/templates
func (f *Fetcher) Transfers_CreateTemplatesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transfers_CreateTemplatesV1", "POST", "{{baseUrl}}/transfers/v1/templates")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transfers_DeleteExternalIncomingV1 (Transfers)
// DELETE {{baseUrl}}/transfers/v1/external/incoming/{id}
func (f *Fetcher) Transfers_DeleteExternalIncomingV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transfers_DeleteExternalIncomingV1", "DELETE", "{{baseUrl}}/transfers/v1/external/incoming/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transfers_DeleteExternalIncomingV2 (Transfers)
// DELETE {{baseUrl}}/transfers/v2/external/incoming/{id}
func (f *Fetcher) Transfers_DeleteExternalIncomingV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transfers_DeleteExternalIncomingV2", "DELETE", "{{baseUrl}}/transfers/v2/external/incoming/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transfers_DeleteTemplatesOutgoingV2 (Transfers)
// DELETE {{baseUrl}}/transfers/v2/templates/outgoing/{id}
func (f *Fetcher) Transfers_DeleteTemplatesOutgoingV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transfers_DeleteTemplatesOutgoingV2", "DELETE", "{{baseUrl}}/transfers/v2/templates/outgoing/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transfers_DeleteTemplatesV1 (Transfers)
// DELETE {{baseUrl}}/transfers/v1/templates/{id}
func (f *Fetcher) Transfers_DeleteTemplatesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transfers_DeleteTemplatesV1", "DELETE", "{{baseUrl}}/transfers/v1/templates/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveriesPackagesStatesV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/deliveries/packages/states
func (f *Fetcher) Transfers_GetDeliveriesPackagesStatesV1() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetDeliveriesPackagesStatesV1", "GET", "{{baseUrl}}/transfers/v1/deliveries/packages/states")
}

// Transfers_GetDeliveriesPackagesStatesV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/packages/states
func (f *Fetcher) Transfers_GetDeliveriesPackagesStatesV2() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetDeliveriesPackagesStatesV2", "GET", "{{baseUrl}}/transfers/v2/deliveries/packages/states")
}

// Transfers_GetDeliveryPackageRequiredlabtestbatchesV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/deliveries/package/{id}/requiredlabtestbatches
func (f *Fetcher) Transfers_GetDeliveryPackageRequiredlabtestbatchesV1() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryPackageRequiredlabtestbatchesV1", "GET", "{{baseUrl}}/transfers/v1/deliveries/package/{id}/requiredlabtestbatches")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveryPackageRequiredlabtestbatchesV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/package/{id}/requiredlabtestbatches
func (f *Fetcher) Transfers_GetDeliveryPackageRequiredlabtestbatchesV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryPackageRequiredlabtestbatchesV2", "GET", "{{baseUrl}}/transfers/v2/deliveries/package/{id}/requiredlabtestbatches")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveryPackageV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/deliveries/{id}/packages
func (f *Fetcher) Transfers_GetDeliveryPackageV1() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryPackageV1", "GET", "{{baseUrl}}/transfers/v1/deliveries/{id}/packages")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveryPackageV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/{id}/packages
func (f *Fetcher) Transfers_GetDeliveryPackageV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryPackageV2", "GET", "{{baseUrl}}/transfers/v2/deliveries/{id}/packages")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveryPackageWholesaleV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/deliveries/{id}/packages/wholesale
func (f *Fetcher) Transfers_GetDeliveryPackageWholesaleV1() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryPackageWholesaleV1", "GET", "{{baseUrl}}/transfers/v1/deliveries/{id}/packages/wholesale")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveryPackageWholesaleV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/{id}/packages/wholesale
func (f *Fetcher) Transfers_GetDeliveryPackageWholesaleV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryPackageWholesaleV2", "GET", "{{baseUrl}}/transfers/v2/deliveries/{id}/packages/wholesale")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveryTransportersDetailsV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/deliveries/{id}/transporters/details
func (f *Fetcher) Transfers_GetDeliveryTransportersDetailsV1() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryTransportersDetailsV1", "GET", "{{baseUrl}}/transfers/v1/deliveries/{id}/transporters/details")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveryTransportersDetailsV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/{id}/transporters/details
func (f *Fetcher) Transfers_GetDeliveryTransportersDetailsV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryTransportersDetailsV2", "GET", "{{baseUrl}}/transfers/v2/deliveries/{id}/transporters/details")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveryTransportersV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/deliveries/{id}/transporters
func (f *Fetcher) Transfers_GetDeliveryTransportersV1() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryTransportersV1", "GET", "{{baseUrl}}/transfers/v1/deliveries/{id}/transporters")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveryTransportersV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/deliveries/{id}/transporters
func (f *Fetcher) Transfers_GetDeliveryTransportersV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryTransportersV2", "GET", "{{baseUrl}}/transfers/v2/deliveries/{id}/transporters")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveryV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/{id}/deliveries
func (f *Fetcher) Transfers_GetDeliveryV1() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryV1", "GET", "{{baseUrl}}/transfers/v1/{id}/deliveries")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetDeliveryV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/{id}/deliveries
func (f *Fetcher) Transfers_GetDeliveryV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetDeliveryV2", "GET", "{{baseUrl}}/transfers/v2/{id}/deliveries")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetHubV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/hub
func (f *Fetcher) Transfers_GetHubV2() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetHubV2", "GET", "{{baseUrl}}/transfers/v2/hub")
}

// Transfers_GetIncomingV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/incoming
func (f *Fetcher) Transfers_GetIncomingV1() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetIncomingV1", "GET", "{{baseUrl}}/transfers/v1/incoming")
}

// Transfers_GetIncomingV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/incoming
func (f *Fetcher) Transfers_GetIncomingV2() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetIncomingV2", "GET", "{{baseUrl}}/transfers/v2/incoming")
}

// Transfers_GetOutgoingV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/outgoing
func (f *Fetcher) Transfers_GetOutgoingV1() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetOutgoingV1", "GET", "{{baseUrl}}/transfers/v1/outgoing")
}

// Transfers_GetOutgoingV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/outgoing
func (f *Fetcher) Transfers_GetOutgoingV2() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetOutgoingV2", "GET", "{{baseUrl}}/transfers/v2/outgoing")
}

// Transfers_GetRejectedV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/rejected
func (f *Fetcher) Transfers_GetRejectedV1() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetRejectedV1", "GET", "{{baseUrl}}/transfers/v1/rejected")
}

// Transfers_GetRejectedV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/rejected
func (f *Fetcher) Transfers_GetRejectedV2() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetRejectedV2", "GET", "{{baseUrl}}/transfers/v2/rejected")
}

// Transfers_GetTemplatesDeliveryPackageV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/templates/deliveries/{id}/packages
func (f *Fetcher) Transfers_GetTemplatesDeliveryPackageV1() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetTemplatesDeliveryPackageV1", "GET", "{{baseUrl}}/transfers/v1/templates/deliveries/{id}/packages")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetTemplatesDeliveryTransportersDetailsV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/templates/deliveries/{id}/transporters/details
func (f *Fetcher) Transfers_GetTemplatesDeliveryTransportersDetailsV1() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetTemplatesDeliveryTransportersDetailsV1", "GET", "{{baseUrl}}/transfers/v1/templates/deliveries/{id}/transporters/details")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetTemplatesDeliveryTransportersV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/templates/deliveries/{id}/transporters
func (f *Fetcher) Transfers_GetTemplatesDeliveryTransportersV1() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetTemplatesDeliveryTransportersV1", "GET", "{{baseUrl}}/transfers/v1/templates/deliveries/{id}/transporters")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetTemplatesDeliveryV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/templates/{id}/deliveries
func (f *Fetcher) Transfers_GetTemplatesDeliveryV1() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetTemplatesDeliveryV1", "GET", "{{baseUrl}}/transfers/v1/templates/{id}/deliveries")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetTemplatesOutgoingDeliveryPackageV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/packages
func (f *Fetcher) Transfers_GetTemplatesOutgoingDeliveryPackageV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetTemplatesOutgoingDeliveryPackageV2", "GET", "{{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/packages")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetTemplatesOutgoingDeliveryTransportersDetailsV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/transporters/details
func (f *Fetcher) Transfers_GetTemplatesOutgoingDeliveryTransportersDetailsV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetTemplatesOutgoingDeliveryTransportersDetailsV2", "GET", "{{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/transporters/details")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetTemplatesOutgoingDeliveryTransportersV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/transporters
func (f *Fetcher) Transfers_GetTemplatesOutgoingDeliveryTransportersV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetTemplatesOutgoingDeliveryTransportersV2", "GET", "{{baseUrl}}/transfers/v2/templates/outgoing/deliveries/{id}/transporters")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetTemplatesOutgoingDeliveryV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/templates/outgoing/{id}/deliveries
func (f *Fetcher) Transfers_GetTemplatesOutgoingDeliveryV2() (models.Transfer, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transfer](f, "Transfers_GetTemplatesOutgoingDeliveryV2", "GET", "{{baseUrl}}/transfers/v2/templates/outgoing/{id}/deliveries")
	if err != nil {
		var empty models.Transfer
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transfer
	return empty, fmt.Errorf("not found")
}

// Transfers_GetTemplatesOutgoingV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/templates/outgoing
func (f *Fetcher) Transfers_GetTemplatesOutgoingV2() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetTemplatesOutgoingV2", "GET", "{{baseUrl}}/transfers/v2/templates/outgoing")
}

// Transfers_GetTemplatesV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/templates
func (f *Fetcher) Transfers_GetTemplatesV1() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetTemplatesV1", "GET", "{{baseUrl}}/transfers/v1/templates")
}

// Transfers_GetTypesV1 (Transfers)
// GET {{baseUrl}}/transfers/v1/types
func (f *Fetcher) Transfers_GetTypesV1() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetTypesV1", "GET", "{{baseUrl}}/transfers/v1/types")
}

// Transfers_GetTypesV2 (Transfers)
// GET {{baseUrl}}/transfers/v2/types
func (f *Fetcher) Transfers_GetTypesV2() ([]models.Transfer, error) {
	return fetchList[models.Transfer](f, "Transfers_GetTypesV2", "GET", "{{baseUrl}}/transfers/v2/types")
}

// Transfers_UpdateExternalIncomingV1 (Transfers)
// PUT {{baseUrl}}/transfers/v1/external/incoming
func (f *Fetcher) Transfers_UpdateExternalIncomingV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transfers_UpdateExternalIncomingV1", "PUT", "{{baseUrl}}/transfers/v1/external/incoming")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transfers_UpdateExternalIncomingV2 (Transfers)
// PUT {{baseUrl}}/transfers/v2/external/incoming
func (f *Fetcher) Transfers_UpdateExternalIncomingV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transfers_UpdateExternalIncomingV2", "PUT", "{{baseUrl}}/transfers/v2/external/incoming")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transfers_UpdateTemplatesOutgoingV2 (Transfers)
// PUT {{baseUrl}}/transfers/v2/templates/outgoing
func (f *Fetcher) Transfers_UpdateTemplatesOutgoingV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transfers_UpdateTemplatesOutgoingV2", "PUT", "{{baseUrl}}/transfers/v2/templates/outgoing")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transfers_UpdateTemplatesV1 (Transfers)
// PUT {{baseUrl}}/transfers/v1/templates
func (f *Fetcher) Transfers_UpdateTemplatesV1() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transfers_UpdateTemplatesV1", "PUT", "{{baseUrl}}/transfers/v1/templates")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transporters_CreateDriverV2 (Transporters)
// POST {{baseUrl}}/transporters/v2/drivers
func (f *Fetcher) Transporters_CreateDriverV2() (models.Transporter, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transporter](f, "Transporters_CreateDriverV2", "POST", "{{baseUrl}}/transporters/v2/drivers")
	if err != nil {
		var empty models.Transporter
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transporter
	return empty, fmt.Errorf("not found")
}

// Transporters_CreateVehicleV2 (Transporters)
// POST {{baseUrl}}/transporters/v2/vehicles
func (f *Fetcher) Transporters_CreateVehicleV2() (models.Transporter, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transporter](f, "Transporters_CreateVehicleV2", "POST", "{{baseUrl}}/transporters/v2/vehicles")
	if err != nil {
		var empty models.Transporter
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transporter
	return empty, fmt.Errorf("not found")
}

// Transporters_DeleteDriverV2 (Transporters)
// DELETE {{baseUrl}}/transporters/v2/drivers/{id}
func (f *Fetcher) Transporters_DeleteDriverV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transporters_DeleteDriverV2", "DELETE", "{{baseUrl}}/transporters/v2/drivers/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transporters_DeleteVehicleV2 (Transporters)
// DELETE {{baseUrl}}/transporters/v2/vehicles/{id}
func (f *Fetcher) Transporters_DeleteVehicleV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transporters_DeleteVehicleV2", "DELETE", "{{baseUrl}}/transporters/v2/vehicles/{id}")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transporters_GetDriverV2 (Transporters)
// GET {{baseUrl}}/transporters/v2/drivers/{id}
func (f *Fetcher) Transporters_GetDriverV2() (models.Transporter, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transporter](f, "Transporters_GetDriverV2", "GET", "{{baseUrl}}/transporters/v2/drivers/{id}")
	if err != nil {
		var empty models.Transporter
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transporter
	return empty, fmt.Errorf("not found")
}

// Transporters_GetDriversV2 (Transporters)
// GET {{baseUrl}}/transporters/v2/drivers
func (f *Fetcher) Transporters_GetDriversV2() ([]models.Transporter, error) {
	return fetchList[models.Transporter](f, "Transporters_GetDriversV2", "GET", "{{baseUrl}}/transporters/v2/drivers")
}

// Transporters_GetVehicleV2 (Transporters)
// GET {{baseUrl}}/transporters/v2/vehicles/{id}
func (f *Fetcher) Transporters_GetVehicleV2() (models.Transporter, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[models.Transporter](f, "Transporters_GetVehicleV2", "GET", "{{baseUrl}}/transporters/v2/vehicles/{id}")
	if err != nil {
		var empty models.Transporter
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty models.Transporter
	return empty, fmt.Errorf("not found")
}

// Transporters_GetVehiclesV2 (Transporters)
// GET {{baseUrl}}/transporters/v2/vehicles
func (f *Fetcher) Transporters_GetVehiclesV2() ([]models.Transporter, error) {
	return fetchList[models.Transporter](f, "Transporters_GetVehiclesV2", "GET", "{{baseUrl}}/transporters/v2/vehicles")
}

// Transporters_UpdateDriverV2 (Transporters)
// PUT {{baseUrl}}/transporters/v2/drivers
func (f *Fetcher) Transporters_UpdateDriverV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transporters_UpdateDriverV2", "PUT", "{{baseUrl}}/transporters/v2/drivers")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// Transporters_UpdateVehicleV2 (Transporters)
// PUT {{baseUrl}}/transporters/v2/vehicles
func (f *Fetcher) Transporters_UpdateVehicleV2() (map[string]any, error) {
	// Single Item Fetch logic (reuse list for now or generic)
	// Fallback to List for robustness since API often returns arrays
	list, err := fetchList[map[string]any](f, "Transporters_UpdateVehicleV2", "PUT", "{{baseUrl}}/transporters/v2/vehicles")
	if err != nil {
		var empty map[string]any
		return empty, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	var empty map[string]any
	return empty, fmt.Errorf("not found")
}

// UnitsOfMeasure_GetActiveV1 (Units Of Measure)
// GET {{baseUrl}}/unitsofmeasure/v1/active
func (f *Fetcher) UnitsOfMeasure_GetActiveV1() ([]models.UnitsOfMeasure, error) {
	return fetchList[models.UnitsOfMeasure](f, "UnitsOfMeasure_GetActiveV1", "GET", "{{baseUrl}}/unitsofmeasure/v1/active")
}

// UnitsOfMeasure_GetActiveV2 (Units Of Measure)
// GET {{baseUrl}}/unitsofmeasure/v2/active
func (f *Fetcher) UnitsOfMeasure_GetActiveV2() ([]models.UnitsOfMeasure, error) {
	return fetchList[models.UnitsOfMeasure](f, "UnitsOfMeasure_GetActiveV2", "GET", "{{baseUrl}}/unitsofmeasure/v2/active")
}

// UnitsOfMeasure_GetInactiveV2 (Units Of Measure)
// GET {{baseUrl}}/unitsofmeasure/v2/inactive
func (f *Fetcher) UnitsOfMeasure_GetInactiveV2() ([]models.UnitsOfMeasure, error) {
	return fetchList[models.UnitsOfMeasure](f, "UnitsOfMeasure_GetInactiveV2", "GET", "{{baseUrl}}/unitsofmeasure/v2/inactive")
}

// WasteMethods_GetAllV2 (Waste Methods)
// GET {{baseUrl}}/wastemethods/v2
func (f *Fetcher) WasteMethods_GetAllV2() ([]models.WasteMethod, error) {
	return fetchList[models.WasteMethod](f, "WasteMethods_GetAllV2", "GET", "{{baseUrl}}/wastemethods/v2")
}
