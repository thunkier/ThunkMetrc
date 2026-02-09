use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdatePlantsTagRequest {
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "NewTag", skip_serializing_if = "Option::is_none")]
    pub new_tag: Option<String>,
    #[serde(rename = "ReplaceDate", skip_serializing_if = "Option::is_none")]
    pub replace_date: Option<String>,
    #[serde(rename = "TagId", skip_serializing_if = "Option::is_none")]
    pub tag_id: Option<i64>,
}
