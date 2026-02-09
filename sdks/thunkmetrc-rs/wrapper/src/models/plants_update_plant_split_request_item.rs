use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsUpdatePlantSplitRequestItem {
    #[serde(rename = "plant_count", skip_serializing_if = "Option::is_none")]
    pub plant_count: Option<i64>,
    #[serde(rename = "source_plant_label", skip_serializing_if = "Option::is_none")]
    pub source_plant_label: Option<String>,
    #[serde(rename = "split_date", skip_serializing_if = "Option::is_none")]
    pub split_date: Option<String>,
    #[serde(rename = "strain_label", skip_serializing_if = "Option::is_none")]
    pub strain_label: Option<String>,
    #[serde(rename = "tag_label", skip_serializing_if = "Option::is_none")]
    pub tag_label: Option<String>,
}
