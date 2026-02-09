use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesCreateSaleDeliveriesRetailerDepartRequestItem {
    #[serde(rename = "retailer_delivery_id", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_id: Option<i64>,
}
