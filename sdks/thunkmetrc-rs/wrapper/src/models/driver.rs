use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Driver {
    #[serde(rename = "DriversLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EmployeeId", skip_serializing_if = "Option::is_none")]
    pub employee_id: Option<String>,
    #[serde(rename = "FacilityId", skip_serializing_if = "Option::is_none")]
    pub facility_id: Option<i64>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "IsArchived", skip_serializing_if = "Option::is_none")]
    pub is_archived: Option<bool>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
}
