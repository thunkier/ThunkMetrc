use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateGrowthPhaseRequestItem {
    #[serde(rename = "Count", skip_serializing_if = "Option::is_none")]
    pub count: Option<i64>,
    #[serde(rename = "GrowthDate", skip_serializing_if = "Option::is_none")]
    pub growth_date: Option<String>,
    #[serde(rename = "GrowthPhase", skip_serializing_if = "Option::is_none")]
    pub growth_phase: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "NewLocation", skip_serializing_if = "Option::is_none")]
    pub new_location: Option<String>,
    #[serde(rename = "NewSublocation", skip_serializing_if = "Option::is_none")]
    pub new_sublocation: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "StartingTag", skip_serializing_if = "Option::is_none")]
    pub starting_tag: Option<String>,
}
