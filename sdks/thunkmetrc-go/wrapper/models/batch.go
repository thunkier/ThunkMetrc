package models

type Batch struct {
    Id int64 `json:"Id,omitempty"`
    ItemCategories []LabTestsBatchItemCategoriesItem `json:"ItemCategories,omitempty"`
    ItemCategoryCount int64 `json:"ItemCategoryCount,omitempty"`
    LabTestTypeCount int64 `json:"LabTestTypeCount,omitempty"`
    LabTestTypes []LabTestsBatchLabTestTypesItem `json:"LabTestTypes,omitempty"`
    Name string `json:"Name,omitempty"`
    RequiresAllFromLabTestBatch bool `json:"RequiresAllFromLabTestBatch,omitempty"`
}
