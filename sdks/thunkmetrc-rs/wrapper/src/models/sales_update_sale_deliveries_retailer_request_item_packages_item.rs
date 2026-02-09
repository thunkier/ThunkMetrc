use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesUpdateSaleDeliveriesRetailerRequestItemPackagesItem {
    #[serde(rename = "date_time", skip_serializing_if = "Option::is_none")]
    pub date_time: Option<String>,
    #[serde(rename = "package_label", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<i64>,
    #[serde(rename = "total_price", skip_serializing_if = "Option::is_none")]
    pub total_price: Option<f64>,
    #[serde(rename = "unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure: Option<String>,
}
