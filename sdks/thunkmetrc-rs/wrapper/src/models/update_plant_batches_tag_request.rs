use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdatePlantBatchesTagRequest {
    #[serde(rename = "Group", skip_serializing_if = "Option::is_none")]
    pub group: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "NewTag", skip_serializing_if = "Option::is_none")]
    pub new_tag: Option<String>,
    #[serde(rename = "ReplaceDate", skip_serializing_if = "Option::is_none")]
    pub replace_date: Option<String>,
    #[serde(rename = "TagId", skip_serializing_if = "Option::is_none")]
    pub tag_id: Option<i64>,
}
