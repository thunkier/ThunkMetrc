package models

type Result struct {
    ExpirationDateTime *string `json:"ExpirationDateTime,omitempty"`
    LabFacilityLicenseNumber string `json:"LabFacilityLicenseNumber,omitempty"`
    LabFacilityName string `json:"LabFacilityName,omitempty"`
    LabTestDetailRevokedDate *string `json:"LabTestDetailRevokedDate,omitempty"`
    LabTestResultDocumentFileId *int64 `json:"LabTestResultDocumentFileId,omitempty"`
    LabTestResultId int64 `json:"LabTestResultId,omitempty"`
    OverallPassed bool `json:"OverallPassed,omitempty"`
    PackageId int64 `json:"PackageId,omitempty"`
    ProductCategoryName string `json:"ProductCategoryName,omitempty"`
    ProductName string `json:"ProductName,omitempty"`
    ResultReleaseDateTime string `json:"ResultReleaseDateTime,omitempty"`
    ResultReleased bool `json:"ResultReleased,omitempty"`
    RevokedDate *string `json:"RevokedDate,omitempty"`
    SourcePackageLabel string `json:"SourcePackageLabel,omitempty"`
    TestComment string `json:"TestComment,omitempty"`
    TestInformationalOnly bool `json:"TestInformationalOnly,omitempty"`
    TestPassed bool `json:"TestPassed,omitempty"`
    TestPerformedDate string `json:"TestPerformedDate,omitempty"`
    TestResultLevel float64 `json:"TestResultLevel,omitempty"`
    TestTypeName string `json:"TestTypeName,omitempty"`
}
