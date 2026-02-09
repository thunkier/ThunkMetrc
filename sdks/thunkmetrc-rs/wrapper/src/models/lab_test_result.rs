use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct LabTestResult {
    #[serde(rename = "ExpirationDateTime", skip_serializing_if = "Option::is_none")]
    pub expiration_date_time: Option<String>,
    #[serde(rename = "LabFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub lab_facility_license_number: Option<String>,
    #[serde(rename = "LabFacilityName", skip_serializing_if = "Option::is_none")]
    pub lab_facility_name: Option<String>,
    #[serde(rename = "LabTestDetailRevokedDate", skip_serializing_if = "Option::is_none")]
    pub lab_test_detail_revoked_date: Option<String>,
    #[serde(rename = "LabTestResultDocumentFileId", skip_serializing_if = "Option::is_none")]
    pub lab_test_result_document_file_id: Option<i64>,
    #[serde(rename = "LabTestResultId", skip_serializing_if = "Option::is_none")]
    pub lab_test_result_id: Option<i64>,
    #[serde(rename = "OverallPassed", skip_serializing_if = "Option::is_none")]
    pub overall_passed: Option<bool>,
    #[serde(rename = "PackageId", skip_serializing_if = "Option::is_none")]
    pub package_id: Option<i64>,
    #[serde(rename = "ProductCategoryName", skip_serializing_if = "Option::is_none")]
    pub product_category_name: Option<String>,
    #[serde(rename = "ProductName", skip_serializing_if = "Option::is_none")]
    pub product_name: Option<String>,
    #[serde(rename = "ResultReleaseDateTime", skip_serializing_if = "Option::is_none")]
    pub result_release_date_time: Option<String>,
    #[serde(rename = "ResultReleased", skip_serializing_if = "Option::is_none")]
    pub result_released: Option<bool>,
    #[serde(rename = "RevokedDate", skip_serializing_if = "Option::is_none")]
    pub revoked_date: Option<String>,
    #[serde(rename = "SourcePackageLabel", skip_serializing_if = "Option::is_none")]
    pub source_package_label: Option<String>,
    #[serde(rename = "TestComment", skip_serializing_if = "Option::is_none")]
    pub test_comment: Option<String>,
    #[serde(rename = "TestInformationalOnly", skip_serializing_if = "Option::is_none")]
    pub test_informational_only: Option<bool>,
    #[serde(rename = "TestPassed", skip_serializing_if = "Option::is_none")]
    pub test_passed: Option<bool>,
    #[serde(rename = "TestPerformedDate", skip_serializing_if = "Option::is_none")]
    pub test_performed_date: Option<String>,
    #[serde(rename = "TestResultLevel", skip_serializing_if = "Option::is_none")]
    pub test_result_level: Option<f64>,
    #[serde(rename = "TestTypeName", skip_serializing_if = "Option::is_none")]
    pub test_type_name: Option<String>,
}
