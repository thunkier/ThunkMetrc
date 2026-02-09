use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreatePackagesInfoRequest {
    #[serde(rename = "packageLabels", skip_serializing_if = "Option::is_none")]
    pub package_labels: Option<Vec<serde_json::Value>>,
}
