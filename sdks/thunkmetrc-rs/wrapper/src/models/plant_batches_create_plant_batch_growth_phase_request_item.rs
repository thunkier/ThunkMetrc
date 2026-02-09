use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantBatchesCreatePlantBatchGrowthPhaseRequestItem {
    #[serde(rename = "count", skip_serializing_if = "Option::is_none")]
    pub count: Option<i64>,
    #[serde(rename = "growth_date", skip_serializing_if = "Option::is_none")]
    pub growth_date: Option<String>,
    #[serde(rename = "growth_phase", skip_serializing_if = "Option::is_none")]
    pub growth_phase: Option<String>,
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "new_location", skip_serializing_if = "Option::is_none")]
    pub new_location: Option<String>,
    #[serde(rename = "new_sublocation", skip_serializing_if = "Option::is_none")]
    pub new_sublocation: Option<String>,
    #[serde(rename = "patient_license_number", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "starting_tag", skip_serializing_if = "Option::is_none")]
    pub starting_tag: Option<String>,
}
