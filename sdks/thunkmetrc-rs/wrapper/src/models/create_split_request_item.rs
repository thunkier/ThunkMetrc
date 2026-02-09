use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateSplitRequestItem {
    #[serde(rename = "ActualDate", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count", skip_serializing_if = "Option::is_none")]
    pub count: Option<i64>,
    #[serde(rename = "GroupName", skip_serializing_if = "Option::is_none")]
    pub group_name: Option<String>,
    #[serde(rename = "Location", skip_serializing_if = "Option::is_none")]
    pub location: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatch", skip_serializing_if = "Option::is_none")]
    pub plant_batch: Option<String>,
    #[serde(rename = "Strain", skip_serializing_if = "Option::is_none")]
    pub strain: Option<String>,
    #[serde(rename = "Sublocation", skip_serializing_if = "Option::is_none")]
    pub sublocation: Option<String>,
}
