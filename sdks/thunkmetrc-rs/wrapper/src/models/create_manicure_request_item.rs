use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateManicureRequestItem {
    #[serde(rename = "ActualDate", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "DryingLocation", skip_serializing_if = "Option::is_none")]
    pub drying_location: Option<String>,
    #[serde(rename = "DryingSublocation", skip_serializing_if = "Option::is_none")]
    pub drying_sublocation: Option<String>,
    #[serde(rename = "HarvestName", skip_serializing_if = "Option::is_none")]
    pub harvest_name: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "Plant", skip_serializing_if = "Option::is_none")]
    pub plant: Option<String>,
    #[serde(rename = "PlantCount", skip_serializing_if = "Option::is_none")]
    pub plant_count: Option<i64>,
    #[serde(rename = "UnitOfWeight", skip_serializing_if = "Option::is_none")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "Weight", skip_serializing_if = "Option::is_none")]
    pub weight: Option<f64>,
}
