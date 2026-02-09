use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct HarvestsUpdateHarvestRenameRequestItem {
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "new_name", skip_serializing_if = "Option::is_none")]
    pub new_name: Option<String>,
    #[serde(rename = "old_name", skip_serializing_if = "Option::is_none")]
    pub old_name: Option<String>,
}
