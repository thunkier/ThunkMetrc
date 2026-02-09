use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PatientCheckInsCreatePatientCheckInsRequestItem {
    #[serde(rename = "check_in_date", skip_serializing_if = "Option::is_none")]
    pub check_in_date: Option<String>,
    #[serde(rename = "check_in_location_id", skip_serializing_if = "Option::is_none")]
    pub check_in_location_id: Option<i64>,
    #[serde(rename = "patient_license_number", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "registration_expiry_date", skip_serializing_if = "Option::is_none")]
    pub registration_expiry_date: Option<String>,
    #[serde(rename = "registration_start_date", skip_serializing_if = "Option::is_none")]
    pub registration_start_date: Option<String>,
}
