use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateDriversRequest {
    #[serde(rename = "DriversLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EmployeeId", skip_serializing_if = "Option::is_none")]
    pub employee_id: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
}
