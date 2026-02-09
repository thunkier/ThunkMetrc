use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct WriteResponse {
    #[serde(rename = "Ids", skip_serializing_if = "Option::is_none")]
    pub ids: Option<Vec<i64>>,
    #[serde(rename = "Warnings", skip_serializing_if = "Option::is_none")]
    pub warnings: Option<String>,
}
