use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsUpdatePlantMergeRequestItem {
    #[serde(rename = "merge_date", skip_serializing_if = "Option::is_none")]
    pub merge_date: Option<String>,
    #[serde(rename = "source_plant_group_label", skip_serializing_if = "Option::is_none")]
    pub source_plant_group_label: Option<String>,
    #[serde(rename = "target_plant_group_label", skip_serializing_if = "Option::is_none")]
    pub target_plant_group_label: Option<String>,
}
