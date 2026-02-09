package models

type PackagesUpdateRemediateRequestItem struct {
    PackageLabel string `json:"PackageLabel,omitempty"`
    RemediationDate string `json:"RemediationDate,omitempty"`
    RemediationMethodName string `json:"RemediationMethodName,omitempty"`
    RemediationSteps string `json:"RemediationSteps,omitempty"`
}
