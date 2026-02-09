use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsCreatePlantManicureRequestItem {
    #[serde(rename = "actual_date", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "drying_location", skip_serializing_if = "Option::is_none")]
    pub drying_location: Option<String>,
    #[serde(rename = "drying_sublocation", skip_serializing_if = "Option::is_none")]
    pub drying_sublocation: Option<String>,
    #[serde(rename = "harvest_name", skip_serializing_if = "Option::is_none")]
    pub harvest_name: Option<String>,
    #[serde(rename = "patient_license_number", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "plant", skip_serializing_if = "Option::is_none")]
    pub plant: Option<String>,
    #[serde(rename = "plant_count", skip_serializing_if = "Option::is_none")]
    pub plant_count: Option<i64>,
    #[serde(rename = "unit_of_weight", skip_serializing_if = "Option::is_none")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "weight", skip_serializing_if = "Option::is_none")]
    pub weight: Option<f64>,
}
