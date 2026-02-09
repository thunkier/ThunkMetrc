use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CaregiversStatus {
    #[serde(rename = "Active", skip_serializing_if = "Option::is_none")]
    pub active: Option<bool>,
    #[serde(rename = "CaregiverLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub caregiver_license_number: Option<String>,
    #[serde(rename = "Patients", skip_serializing_if = "Option::is_none")]
    pub patients: Option<Vec<String>>,
}
