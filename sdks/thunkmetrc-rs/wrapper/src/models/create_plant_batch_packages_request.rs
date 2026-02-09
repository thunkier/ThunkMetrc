use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreatePlantBatchPackagesRequest {
    #[serde(rename = "ActualDate", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count", skip_serializing_if = "Option::is_none")]
    pub count: Option<i64>,
    #[serde(rename = "IsDonation", skip_serializing_if = "Option::is_none")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsTradeSample", skip_serializing_if = "Option::is_none")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item", skip_serializing_if = "Option::is_none")]
    pub item: Option<String>,
    #[serde(rename = "Location", skip_serializing_if = "Option::is_none")]
    pub location: Option<String>,
    #[serde(rename = "Note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "PackageTag", skip_serializing_if = "Option::is_none")]
    pub package_tag: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatchType", skip_serializing_if = "Option::is_none")]
    pub plant_batch_type: Option<String>,
    #[serde(rename = "PlantLabel", skip_serializing_if = "Option::is_none")]
    pub plant_label: Option<String>,
    #[serde(rename = "Sublocation", skip_serializing_if = "Option::is_none")]
    pub sublocation: Option<String>,
}
