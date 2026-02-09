use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateGrowthPhaseRequestItem {
    #[serde(rename = "GrowthDate", skip_serializing_if = "Option::is_none")]
    pub growth_date: Option<String>,
    #[serde(rename = "GrowthPhase", skip_serializing_if = "Option::is_none")]
    pub growth_phase: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "NewLocation", skip_serializing_if = "Option::is_none")]
    pub new_location: Option<String>,
    #[serde(rename = "NewSublocation", skip_serializing_if = "Option::is_none")]
    pub new_sublocation: Option<String>,
    #[serde(rename = "NewTag", skip_serializing_if = "Option::is_none")]
    pub new_tag: Option<String>,
}
