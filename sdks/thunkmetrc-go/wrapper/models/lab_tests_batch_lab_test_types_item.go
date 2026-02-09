package models

type LabTestsBatchLabTestTypesItem struct {
    AlwaysPasses bool `json:"AlwaysPasses,omitempty"`
    DependencyMode int64 `json:"DependencyMode,omitempty"`
    Id int64 `json:"Id,omitempty"`
    InformationalOnly bool `json:"InformationalOnly,omitempty"`
    LabTestResultExpirationDays *int64 `json:"LabTestResultExpirationDays,omitempty"`
    LabTestResultMaximum *string `json:"LabTestResultMaximum,omitempty"`
    LabTestResultMinimum *string `json:"LabTestResultMinimum,omitempty"`
    LabTestResultMode int64 `json:"LabTestResultMode,omitempty"`
    MaxAllowedFailureCount int64 `json:"MaxAllowedFailureCount,omitempty"`
    Name string `json:"Name,omitempty"`
    RequiresTestResult bool `json:"RequiresTestResult,omitempty"`
    ResearchAndDevelopment bool `json:"ResearchAndDevelopment,omitempty"`
}
