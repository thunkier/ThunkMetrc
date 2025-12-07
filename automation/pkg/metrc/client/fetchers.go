package client

import (
	"strings"

	"github.com/thunkmetrc/automation/internal/runner/common"
	"github.com/thunkmetrc/automation/internal/runner/util"
	"github.com/thunkmetrc/automation/pkg/bruno"
	"github.com/thunkmetrc/automation/pkg/metrc/models"
)

type Fetcher struct {
	Cfg common.Config
}

func NewFetcher(cfg common.Config) *Fetcher {
	return &Fetcher{Cfg: cfg}
}

func fetchList[T any](f *Fetcher, name, method, url string) ([]T, error) {
	req := &bruno.Request{Name: name, Method: method, URL: url}

	if parts := strings.Split(name, "_"); len(parts) > 0 {
		req.Group = parts[0]
	}

	body := ExecuteRequest(f.Cfg, req, "")
	var list []T
	err := util.ParseList(body, &list)
	return list, err
}

type SimpleName struct {
	Id   int    `json:"Id,omitempty"`
	Name string `json:"Name"`
}

type SimpleIdName struct {
	Id   int    `json:"Id"`
	Name string `json:"Name"`
}

type SimpleLabel struct {
	Label string `json:"Label"`
}

type SimpleTag struct {
	Tag   string `json:"Tag"`
	Label string `json:"Label"`
}

func (f *Fetcher) GetLocationTypes() ([]SimpleName, error) {
	return fetchList[SimpleName](f, "GetLocationTypes", "GET", "{{baseUrl}}/locations/v1/types")
}

func (f *Fetcher) GetActiveLocations() ([]SimpleName, error) {
	locs, err := f.Locations_GetActiveV2()
	if err != nil {
		return nil, err
	}
	var simple []SimpleName
	for _, l := range locs {
		simple = append(simple, SimpleName{
			Id:   l.Id,
			Name: l.Name,
		})
	}
	return simple, nil
}

func (f *Fetcher) GetStrains() ([]SimpleName, error) {
	list, err := f.Strains_GetActiveV1()
	if err != nil {
		return nil, err
	}
	var simple []SimpleName
	for _, x := range list {
		simple = append(simple, SimpleName{
			Id:   x.Id,
			Name: x.Name,
		})
	}
	return simple, nil
}

func (f *Fetcher) GetActiveItems() ([]models.Item, error) {
	return f.Items_GetActiveV2()
}

type ItemCategory struct {
	Name                string `json:"Name"`
	ProductCategoryType string `json:"ProductCategoryType"`
}

func (f *Fetcher) GetItemCategories() ([]ItemCategory, error) {
	list, err := f.Items_GetCategoriesV1()
	if err != nil {
		return nil, err
	}
	var res []ItemCategory
	for _, x := range list {
		res = append(res, ItemCategory{
			Name:                x.Name,
			ProductCategoryType: x.ProductCategoryType,
		})
	}
	return res, nil
}

func (f *Fetcher) GetPackageTags() ([]SimpleTag, error) {
	list, err := f.Tags_GetPackageAvailableV2()
	if err != nil {
		return nil, err
	}
	var res []SimpleTag
	for _, x := range list {
		res = append(res, SimpleTag{
			Tag:   x.Label,
			Label: x.Label,
		})
	}
	return res, nil
}

func (f *Fetcher) GetPlantTags() ([]SimpleTag, error) {
	list, err := f.Tags_GetPlantAvailableV2()
	if err != nil {
		return nil, err
	}
	var res []SimpleTag
	for _, x := range list {
		res = append(res, SimpleTag{
			Tag:   x.Label,
			Label: x.Label,
		})
	}
	return res, nil
}

func (f *Fetcher) GetWasteMethods() ([]SimpleName, error) {
	list, err := f.WasteMethods_GetAllV2()
	if err != nil {
		return nil, err
	}
	var res []SimpleName
	for _, x := range list {
		res = append(res, SimpleName{
			Id:   0,
			Name: x.Name,
		})
	}
	return res, nil
}

func (f *Fetcher) GetVegetativePlants() ([]SimpleLabel, error) {
	list, err := f.Plants_GetVegetativeV1()
	if err != nil {
		return nil, err
	}
	var res []SimpleLabel
	for _, x := range list {
		res = append(res, SimpleLabel{
			Label: x.Label,
		})
	}
	return res, nil
}

func (f *Fetcher) GetActiveHarvests() ([]SimpleIdName, error) {
	list, err := f.Harvests_GetActiveV1()
	if err != nil {
		return nil, err
	}
	var res []SimpleIdName
	for _, x := range list {
		res = append(res, SimpleIdName{
			Id:   x.Id,
			Name: x.Name,
		})
	}
	return res, nil
}

func (f *Fetcher) GetActivePlantBatches() ([]SimpleName, error) {
	list, err := f.PlantBatches_GetActiveV1()
	if err != nil {
		return nil, err
	}
	var res []SimpleName
	for _, x := range list {
		res = append(res, SimpleName{
			Id:   x.Id,
			Name: x.Name,
		})
	}
	return res, nil
}

func (f *Fetcher) GetActivePackages() ([]models.Package, error) {
	return f.Packages_GetActiveV2()
}

func (f *Fetcher) GetProcessingJobTypes() ([]models.ProcessingJobData, error) {
	list, err := f.ProcessingJobs_GetJobtypesActiveV2()
	if err != nil {
		return nil, err
	}
	var res []models.ProcessingJobData
	for _, x := range list {
		res = append(res, models.ProcessingJobData{
			Id:   x.Id,
			Name: x.Name,
		})
	}
	return res, nil
}

func (f *Fetcher) GetJobTypeCategories() ([]SimpleName, error) {
	list, err := f.ProcessingJobs_GetJobtypesCategoriesV2()
	if err != nil {
		return nil, err
	}
	var res []SimpleName
	for _, x := range list {
		res = append(res, SimpleName{
			Name: x.Name,
		})
	}
	return res, nil
}

func (f *Fetcher) GetJobTypeAttributes() ([]SimpleName, error) {
	list, err := f.ProcessingJobs_GetJobtypesAttributesV2()
	if err != nil {
		return nil, err
	}
	var res []SimpleName
	for _, x := range list {
		res = append(res, SimpleName{
			Name: x.Name,
		})
	}
	return res, nil
}

func (f *Fetcher) GetLabTestTypes() ([]SimpleName, error) {
	list, err := f.LabTests_GetTypesV1()
	if err != nil {
		return nil, err
	}
	var res []SimpleName
	for _, x := range list {
		res = append(res, SimpleName{
			Name: x.Name,
		})
	}
	return res, nil
}

func (f *Fetcher) GetTransferTypes() ([]models.TransferData, error) {
	list, err := f.Transfers_GetTypesV1()
	if err != nil {
		return nil, err
	}
	var res []models.TransferData
	for _, x := range list {
		res = append(res, models.TransferData{
			Id:   x.Id,
			Name: x.Name,
		})
	}
	return res, nil
}

func (f *Fetcher) GetFacilities() ([]models.Facility, error) {
	f.Facilities_GetAllV2()
	return f.Facilities_GetAllV1()
}
