use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsUpdatePlantGrowthPhaseRequestItem {
    #[serde(rename = "growth_date", skip_serializing_if = "Option::is_none")]
    pub growth_date: Option<String>,
    #[serde(rename = "growth_phase", skip_serializing_if = "Option::is_none")]
    pub growth_phase: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "new_location", skip_serializing_if = "Option::is_none")]
    pub new_location: Option<String>,
    #[serde(rename = "new_sublocation", skip_serializing_if = "Option::is_none")]
    pub new_sublocation: Option<String>,
    #[serde(rename = "new_tag", skip_serializing_if = "Option::is_none")]
    pub new_tag: Option<String>,
}
