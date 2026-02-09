use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantBatchesCreatePlantBatchPlantingsRequestItem {
    #[serde(rename = "actual_date", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "count", skip_serializing_if = "Option::is_none")]
    pub count: Option<i64>,
    #[serde(rename = "location", skip_serializing_if = "Option::is_none")]
    pub location: Option<String>,
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "patient_license_number", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "source_plant_batches", skip_serializing_if = "Option::is_none")]
    pub source_plant_batches: Option<String>,
    #[serde(rename = "strain", skip_serializing_if = "Option::is_none")]
    pub strain: Option<String>,
    #[serde(rename = "sublocation", skip_serializing_if = "Option::is_none")]
    pub sublocation: Option<String>,
    #[serde(rename = "type_", skip_serializing_if = "Option::is_none")]
    pub type_: Option<String>,
}
