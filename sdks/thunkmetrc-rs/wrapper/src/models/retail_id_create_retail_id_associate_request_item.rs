use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct RetailIdCreateRetailIdAssociateRequestItem {
    #[serde(rename = "package_label", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "qr_urls", skip_serializing_if = "Option::is_none")]
    pub qr_urls: Option<Vec<String>>,
}
