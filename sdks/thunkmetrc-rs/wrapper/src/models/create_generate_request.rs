use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateGenerateRequest {
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "Quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<f64>,
}
