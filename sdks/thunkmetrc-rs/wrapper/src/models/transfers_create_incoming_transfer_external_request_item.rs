use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TransfersCreateIncomingTransferExternalRequestItem {
    #[serde(rename = "destinations", skip_serializing_if = "Option::is_none")]
    pub destinations: Option<Vec<TransfersCreateIncomingTransferExternalRequestItemDestinationsItem>>,
    #[serde(rename = "driver_license_number", skip_serializing_if = "Option::is_none")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "driver_name", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "driver_occupational_license_number", skip_serializing_if = "Option::is_none")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "phone_number_for_questions", skip_serializing_if = "Option::is_none")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "shipper_address1", skip_serializing_if = "Option::is_none")]
    pub shipper_address1: Option<String>,
    #[serde(rename = "shipper_address2", skip_serializing_if = "Option::is_none")]
    pub shipper_address2: Option<String>,
    #[serde(rename = "shipper_address_city", skip_serializing_if = "Option::is_none")]
    pub shipper_address_city: Option<String>,
    #[serde(rename = "shipper_address_postal_code", skip_serializing_if = "Option::is_none")]
    pub shipper_address_postal_code: Option<String>,
    #[serde(rename = "shipper_address_state", skip_serializing_if = "Option::is_none")]
    pub shipper_address_state: Option<String>,
    #[serde(rename = "shipper_license_number", skip_serializing_if = "Option::is_none")]
    pub shipper_license_number: Option<String>,
    #[serde(rename = "shipper_main_phone_number", skip_serializing_if = "Option::is_none")]
    pub shipper_main_phone_number: Option<String>,
    #[serde(rename = "shipper_name", skip_serializing_if = "Option::is_none")]
    pub shipper_name: Option<String>,
    #[serde(rename = "transporter_facility_license_number", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "vehicle_license_plate_number", skip_serializing_if = "Option::is_none")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "vehicle_make", skip_serializing_if = "Option::is_none")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "vehicle_model", skip_serializing_if = "Option::is_none")]
    pub vehicle_model: Option<String>,
}
