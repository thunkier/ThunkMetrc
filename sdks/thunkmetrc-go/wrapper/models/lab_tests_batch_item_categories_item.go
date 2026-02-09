package models

type LabTestsBatchItemCategoriesItem struct {
    ProductCategory LabTestsBatchItemCategoriesItemProductCategory `json:"ProductCategory,omitempty"`
    ProductCategoryId string `json:"ProductCategoryId,omitempty"`
    RequireLabTestBatch bool `json:"RequireLabTestBatch,omitempty"`
}
