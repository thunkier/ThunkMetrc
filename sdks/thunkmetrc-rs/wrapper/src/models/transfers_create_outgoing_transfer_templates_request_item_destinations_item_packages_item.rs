use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TransfersCreateOutgoingTransferTemplatesRequestItemDestinationsItemPackagesItem {
    #[serde(rename = "gross_unit_of_weight_name", skip_serializing_if = "Option::is_none")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "gross_weight", skip_serializing_if = "Option::is_none")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "package_label", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "wholesale_price", skip_serializing_if = "Option::is_none")]
    pub wholesale_price: Option<String>,
}
