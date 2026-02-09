use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateIncomingExternalRequest {
    #[serde(rename = "Destinations", skip_serializing_if = "Option::is_none")]
    pub destinations: Option<Vec<serde_json::Value>>,
    #[serde(rename = "DriverLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions", skip_serializing_if = "Option::is_none")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "ShipperAddress1", skip_serializing_if = "Option::is_none")]
    pub shipper_address1: Option<String>,
    #[serde(rename = "ShipperAddress2", skip_serializing_if = "Option::is_none")]
    pub shipper_address2: Option<String>,
    #[serde(rename = "ShipperAddressCity", skip_serializing_if = "Option::is_none")]
    pub shipper_address_city: Option<String>,
    #[serde(rename = "ShipperAddressPostalCode", skip_serializing_if = "Option::is_none")]
    pub shipper_address_postal_code: Option<String>,
    #[serde(rename = "ShipperAddressState", skip_serializing_if = "Option::is_none")]
    pub shipper_address_state: Option<String>,
    #[serde(rename = "ShipperLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub shipper_license_number: Option<String>,
    #[serde(rename = "ShipperMainPhoneNumber", skip_serializing_if = "Option::is_none")]
    pub shipper_main_phone_number: Option<String>,
    #[serde(rename = "ShipperName", skip_serializing_if = "Option::is_none")]
    pub shipper_name: Option<String>,
    #[serde(rename = "TransferId", skip_serializing_if = "Option::is_none")]
    pub transfer_id: Option<i64>,
    #[serde(rename = "TransporterFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber", skip_serializing_if = "Option::is_none")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake", skip_serializing_if = "Option::is_none")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel", skip_serializing_if = "Option::is_none")]
    pub vehicle_model: Option<String>,
}
