use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateDeliveriesRetailerRestockRequestItem {
    #[serde(rename = "DateTime", skip_serializing_if = "Option::is_none")]
    pub date_time: Option<String>,
    #[serde(rename = "Destinations", skip_serializing_if = "Option::is_none")]
    pub destinations: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Packages", skip_serializing_if = "Option::is_none")]
    pub packages: Option<Vec<serde_json::Value>>,
    #[serde(rename = "RetailerDeliveryId", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_id: Option<i64>,
}
