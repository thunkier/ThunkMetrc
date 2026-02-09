use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesUpdateSaleDeliveriesHubVerifyIDRequestItem {
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "payment_type", skip_serializing_if = "Option::is_none")]
    pub payment_type: Option<String>,
}
