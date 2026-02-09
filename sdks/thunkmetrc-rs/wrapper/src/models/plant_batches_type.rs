use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantBatchesType {
    #[serde(rename = "CanBeCloned", skip_serializing_if = "Option::is_none")]
    pub can_be_cloned: Option<bool>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
}
