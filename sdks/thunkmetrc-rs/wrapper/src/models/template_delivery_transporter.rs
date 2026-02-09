use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TemplateDeliveryTransporter {
    #[serde(rename = "ShipmentDeliveryId", skip_serializing_if = "Option::is_none")]
    pub shipment_delivery_id: Option<i64>,
    #[serde(rename = "TransporterDirection", skip_serializing_if = "Option::is_none")]
    pub transporter_direction: Option<String>,
    #[serde(rename = "TransporterFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "TransporterFacilityName", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_name: Option<String>,
}
