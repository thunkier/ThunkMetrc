package models

type PackagesUpdatePackageLabtestsRequiredRequestItem struct {
    Label string `json:"Label,omitempty"`
    RequiredLabTestBatches []string `json:"RequiredLabTestBatches,omitempty"`
}
