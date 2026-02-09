use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PackagesUpdatePackageExternalidRequestItem {
    #[serde(rename = "external_id", skip_serializing_if = "Option::is_none")]
    pub external_id: Option<String>,
    #[serde(rename = "package_label", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
}
