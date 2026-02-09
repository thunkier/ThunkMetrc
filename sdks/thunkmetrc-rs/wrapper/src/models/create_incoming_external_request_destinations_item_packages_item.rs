use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateIncomingExternalRequestDestinationsItemPackagesItem {
    #[serde(rename = "ExpirationDate", skip_serializing_if = "Option::is_none")]
    pub expiration_date: Option<String>,
    #[serde(rename = "ExternalId", skip_serializing_if = "Option::is_none")]
    pub external_id: Option<String>,
    #[serde(rename = "GrossUnitOfWeightName", skip_serializing_if = "Option::is_none")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "GrossWeight", skip_serializing_if = "Option::is_none")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "ItemName", skip_serializing_if = "Option::is_none")]
    pub item_name: Option<String>,
    #[serde(rename = "PackagedDate", skip_serializing_if = "Option::is_none")]
    pub packaged_date: Option<String>,
    #[serde(rename = "Quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<i64>,
    #[serde(rename = "SellByDate", skip_serializing_if = "Option::is_none")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "UnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "UseByDate", skip_serializing_if = "Option::is_none")]
    pub use_by_date: Option<String>,
    #[serde(rename = "WholesalePrice", skip_serializing_if = "Option::is_none")]
    pub wholesale_price: Option<String>,
}
