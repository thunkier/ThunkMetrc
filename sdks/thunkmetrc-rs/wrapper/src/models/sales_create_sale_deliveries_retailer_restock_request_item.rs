use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesCreateSaleDeliveriesRetailerRestockRequestItem {
    #[serde(rename = "date_time", skip_serializing_if = "Option::is_none")]
    pub date_time: Option<String>,
    #[serde(rename = "destinations", skip_serializing_if = "Option::is_none")]
    pub destinations: Option<String>,
    #[serde(rename = "estimated_departure_date_time", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "packages", skip_serializing_if = "Option::is_none")]
    pub packages: Option<Vec<SalesCreateSaleDeliveriesRetailerRestockRequestItemPackagesItem>>,
    #[serde(rename = "retailer_delivery_id", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_id: Option<i64>,
}
