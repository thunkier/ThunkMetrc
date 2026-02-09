use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Allotment {
    #[serde(rename = "Allotment", skip_serializing_if = "Option::is_none")]
    pub allotment: Option<i64>,
    #[serde(rename = "FacilityId", skip_serializing_if = "Option::is_none")]
    pub facility_id: Option<i64>,
    #[serde(rename = "FacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub facility_license_number: Option<String>,
}
