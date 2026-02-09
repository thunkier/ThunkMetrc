use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateDeliveriesHubVerifyIDRequestItem {
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "PaymentType", skip_serializing_if = "Option::is_none")]
    pub payment_type: Option<String>,
}
