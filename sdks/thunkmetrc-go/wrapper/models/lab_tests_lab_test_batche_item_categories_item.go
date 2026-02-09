package models

type LabTestsLabTestBatcheItemCategoriesItem struct {
    ProductCategory LabTestsLabTestBatcheItemCategoriesItemProductCategory `json:"ProductCategory,omitempty"`
    ProductCategoryId string `json:"ProductCategoryId,omitempty"`
    RequireLabTestBatch bool `json:"RequireLabTestBatch,omitempty"`
}
