use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateSplitRequestItem {
    #[serde(rename = "PlantCount", skip_serializing_if = "Option::is_none")]
    pub plant_count: Option<i64>,
    #[serde(rename = "SourcePlantLabel", skip_serializing_if = "Option::is_none")]
    pub source_plant_label: Option<String>,
    #[serde(rename = "SplitDate", skip_serializing_if = "Option::is_none")]
    pub split_date: Option<String>,
    #[serde(rename = "StrainLabel", skip_serializing_if = "Option::is_none")]
    pub strain_label: Option<String>,
    #[serde(rename = "TagLabel", skip_serializing_if = "Option::is_none")]
    pub tag_label: Option<String>,
}
