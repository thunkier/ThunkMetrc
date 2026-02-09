package models

type LabTestBatche struct {
    Id int64 `json:"Id,omitempty"`
    ItemCategories []LabTestsLabTestBatcheItemCategoriesItem `json:"ItemCategories,omitempty"`
    ItemCategoryCount int64 `json:"ItemCategoryCount,omitempty"`
    LabTestTypeCount int64 `json:"LabTestTypeCount,omitempty"`
    LabTestTypes []LabTestsLabTestBatcheLabTestTypesItem `json:"LabTestTypes,omitempty"`
    Name string `json:"Name,omitempty"`
    RequiresAllFromLabTestBatch bool `json:"RequiresAllFromLabTestBatch,omitempty"`
}
