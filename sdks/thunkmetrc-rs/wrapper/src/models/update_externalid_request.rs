use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateExternalidRequest {
    #[serde(rename = "ExternalId", skip_serializing_if = "Option::is_none")]
    pub external_id: Option<String>,
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
}
