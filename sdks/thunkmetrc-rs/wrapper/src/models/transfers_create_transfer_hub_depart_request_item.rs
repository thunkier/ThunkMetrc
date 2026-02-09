use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TransfersCreateTransferHubDepartRequestItem {
    #[serde(rename = "shipment_delivery_id", skip_serializing_if = "Option::is_none")]
    pub shipment_delivery_id: Option<i64>,
    #[serde(rename = "transporter_direction", skip_serializing_if = "Option::is_none")]
    pub transporter_direction: Option<String>,
}
