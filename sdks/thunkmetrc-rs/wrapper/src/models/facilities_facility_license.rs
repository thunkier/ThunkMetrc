use serde::{Deserialize, Serialize};
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct FacilitiesFacilityLicense {
    #[serde(rename = "EmployeeLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub employee_license_number: Option<String>,
    #[serde(rename = "EndDate", skip_serializing_if = "Option::is_none")]
    pub end_date: Option<String>,
    #[serde(rename = "LicenseType", skip_serializing_if = "Option::is_none")]
    pub license_type: Option<String>,
    #[serde(rename = "Number", skip_serializing_if = "Option::is_none")]
    pub number: Option<String>,
    #[serde(rename = "StartDate", skip_serializing_if = "Option::is_none")]
    pub start_date: Option<String>,
}
