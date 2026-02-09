package models

type PackagesUpdateLabtestsRequiredRequestItem struct {
    Label string `json:"Label,omitempty"`
    RequiredLabTestBatches []string `json:"RequiredLabTestBatches,omitempty"`
}
