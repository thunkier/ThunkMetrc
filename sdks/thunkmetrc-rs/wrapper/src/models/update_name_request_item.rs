use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateNameRequestItem {
    #[serde(rename = "Group", skip_serializing_if = "Option::is_none")]
    pub group: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "NewGroup", skip_serializing_if = "Option::is_none")]
    pub new_group: Option<String>,
}
