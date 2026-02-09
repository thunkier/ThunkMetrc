use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateDeliveriesRetailerDepartRequest {
    #[serde(rename = "RetailerDeliveryId", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_id: Option<i64>,
}
