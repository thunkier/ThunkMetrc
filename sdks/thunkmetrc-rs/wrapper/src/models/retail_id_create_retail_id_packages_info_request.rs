use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct RetailIdCreateRetailIdPackagesInfoRequest {
    #[serde(rename = "package_labels", skip_serializing_if = "Option::is_none")]
    pub package_labels: Option<Vec<String>>,
}
