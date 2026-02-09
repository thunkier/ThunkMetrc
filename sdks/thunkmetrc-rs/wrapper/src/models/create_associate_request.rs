use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateAssociateRequest {
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "QrUrls", skip_serializing_if = "Option::is_none")]
    pub qr_urls: Option<Vec<serde_json::Value>>,
}
