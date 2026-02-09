use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsCreatePlantWasteRequestItem {
    #[serde(rename = "location_name", skip_serializing_if = "Option::is_none")]
    pub location_name: Option<String>,
    #[serde(rename = "mixed_material", skip_serializing_if = "Option::is_none")]
    pub mixed_material: Option<String>,
    #[serde(rename = "note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "plant_labels", skip_serializing_if = "Option::is_none")]
    pub plant_labels: Option<Vec<serde_json::Value>>,
    #[serde(rename = "reason_name", skip_serializing_if = "Option::is_none")]
    pub reason_name: Option<String>,
    #[serde(rename = "sublocation_name", skip_serializing_if = "Option::is_none")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "waste_date", skip_serializing_if = "Option::is_none")]
    pub waste_date: Option<String>,
    #[serde(rename = "waste_method_name", skip_serializing_if = "Option::is_none")]
    pub waste_method_name: Option<String>,
    #[serde(rename = "waste_weight", skip_serializing_if = "Option::is_none")]
    pub waste_weight: Option<f64>,
}
