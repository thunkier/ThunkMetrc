use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PatientCheckIn {
    #[serde(rename = "CheckInDate", skip_serializing_if = "Option::is_none")]
    pub check_in_date: Option<String>,
    #[serde(rename = "CheckInLocationId", skip_serializing_if = "Option::is_none")]
    pub check_in_location_id: Option<i64>,
    #[serde(rename = "CheckInLocationName", skip_serializing_if = "Option::is_none")]
    pub check_in_location_name: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "RegistrationExpiryDate", skip_serializing_if = "Option::is_none")]
    pub registration_expiry_date: Option<String>,
    #[serde(rename = "RegistrationStartDate", skip_serializing_if = "Option::is_none")]
    pub registration_start_date: Option<String>,
}
