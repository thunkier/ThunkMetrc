use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DeliveryPackageWholesale {
    #[serde(rename = "PackageId", skip_serializing_if = "Option::is_none")]
    pub package_id: Option<i64>,
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "ReceiverWholesalePrice", skip_serializing_if = "Option::is_none")]
    pub receiver_wholesale_price: Option<String>,
    #[serde(rename = "ShipperWholesalePrice", skip_serializing_if = "Option::is_none")]
    pub shipper_wholesale_price: Option<String>,
}
