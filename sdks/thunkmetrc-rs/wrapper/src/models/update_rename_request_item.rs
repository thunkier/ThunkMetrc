use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateRenameRequestItem {
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "NewName", skip_serializing_if = "Option::is_none")]
    pub new_name: Option<String>,
    #[serde(rename = "OldName", skip_serializing_if = "Option::is_none")]
    pub old_name: Option<String>,
}
