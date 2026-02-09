use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TransfersCreateIncomingTransferExternalRequestItemDestinationsItemPackagesItem {
    #[serde(rename = "expiration_date", skip_serializing_if = "Option::is_none")]
    pub expiration_date: Option<String>,
    #[serde(rename = "external_id", skip_serializing_if = "Option::is_none")]
    pub external_id: Option<String>,
    #[serde(rename = "gross_unit_of_weight_name", skip_serializing_if = "Option::is_none")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "gross_weight", skip_serializing_if = "Option::is_none")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "item_name", skip_serializing_if = "Option::is_none")]
    pub item_name: Option<String>,
    #[serde(rename = "packaged_date", skip_serializing_if = "Option::is_none")]
    pub packaged_date: Option<String>,
    #[serde(rename = "quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<i64>,
    #[serde(rename = "sell_by_date", skip_serializing_if = "Option::is_none")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "use_by_date", skip_serializing_if = "Option::is_none")]
    pub use_by_date: Option<String>,
    #[serde(rename = "wholesale_price", skip_serializing_if = "Option::is_none")]
    pub wholesale_price: Option<String>,
}
