use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantBatchesUpdatePlantBatchTagRequestItem {
    #[serde(rename = "group", skip_serializing_if = "Option::is_none")]
    pub group: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "new_tag", skip_serializing_if = "Option::is_none")]
    pub new_tag: Option<String>,
    #[serde(rename = "replace_date", skip_serializing_if = "Option::is_none")]
    pub replace_date: Option<String>,
    #[serde(rename = "tag_id", skip_serializing_if = "Option::is_none")]
    pub tag_id: Option<i64>,
}
