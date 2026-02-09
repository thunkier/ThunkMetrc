package models

type PackagesUpdatePackageDecontaminateRequestItem struct {
    DecontaminationDate string `json:"DecontaminationDate,omitempty"`
    DecontaminationMethodName string `json:"DecontaminationMethodName,omitempty"`
    DecontaminationSteps string `json:"DecontaminationSteps,omitempty"`
    PackageLabel string `json:"PackageLabel,omitempty"`
}
