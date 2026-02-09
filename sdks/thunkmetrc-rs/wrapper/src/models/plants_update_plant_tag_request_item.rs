use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsUpdatePlantTagRequestItem {
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "new_tag", skip_serializing_if = "Option::is_none")]
    pub new_tag: Option<String>,
    #[serde(rename = "replace_date", skip_serializing_if = "Option::is_none")]
    pub replace_date: Option<String>,
    #[serde(rename = "tag_id", skip_serializing_if = "Option::is_none")]
    pub tag_id: Option<i64>,
}
