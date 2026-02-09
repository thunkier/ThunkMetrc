use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesUpdateSaleDeliveriesCompleteRequestItem {
    #[serde(rename = "accepted_packages", skip_serializing_if = "Option::is_none")]
    pub accepted_packages: Option<Vec<String>>,
    #[serde(rename = "actual_arrival_date_time", skip_serializing_if = "Option::is_none")]
    pub actual_arrival_date_time: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "payment_type", skip_serializing_if = "Option::is_none")]
    pub payment_type: Option<String>,
    #[serde(rename = "returned_packages", skip_serializing_if = "Option::is_none")]
    pub returned_packages: Option<Vec<SalesUpdateSaleDeliveriesCompleteRequestItemReturnedPackagesItem>>,
}
