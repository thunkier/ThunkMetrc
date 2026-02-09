use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateHubCheckinRequest {
    #[serde(rename = "ShipmentDeliveryId", skip_serializing_if = "Option::is_none")]
    pub shipment_delivery_id: Option<i64>,
    #[serde(rename = "TransporterDirection", skip_serializing_if = "Option::is_none")]
    pub transporter_direction: Option<String>,
}
