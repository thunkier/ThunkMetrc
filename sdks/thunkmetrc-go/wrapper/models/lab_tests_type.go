package models

type LabTestsType struct {
    AlwaysPasses bool `json:"AlwaysPasses,omitempty"`
    DependencyMode string `json:"DependencyMode,omitempty"`
    Id int64 `json:"Id,omitempty"`
    InformationalOnly bool `json:"InformationalOnly,omitempty"`
    LabTestResultExpirationDays *int64 `json:"LabTestResultExpirationDays,omitempty"`
    LabTestResultMaximum int64 `json:"LabTestResultMaximum,omitempty"`
    LabTestResultMinimum int64 `json:"LabTestResultMinimum,omitempty"`
    LabTestResultMode string `json:"LabTestResultMode,omitempty"`
    MaxAllowedFailureCount int64 `json:"MaxAllowedFailureCount,omitempty"`
    Name string `json:"Name,omitempty"`
    RequiresTestResult bool `json:"RequiresTestResult,omitempty"`
    ResearchAndDevelopment bool `json:"ResearchAndDevelopment,omitempty"`
}
