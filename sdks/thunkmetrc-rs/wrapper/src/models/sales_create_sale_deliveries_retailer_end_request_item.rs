use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesCreateSaleDeliveriesRetailerEndRequestItem {
    #[serde(rename = "actual_arrival_date_time", skip_serializing_if = "Option::is_none")]
    pub actual_arrival_date_time: Option<String>,
    #[serde(rename = "packages", skip_serializing_if = "Option::is_none")]
    pub packages: Option<Vec<SalesCreateSaleDeliveriesRetailerEndRequestItemPackagesItem>>,
    #[serde(rename = "retailer_delivery_id", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_id: Option<i64>,
}
