use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateOutgoingTemplatesRequestItem {
    #[serde(rename = "Destinations", skip_serializing_if = "Option::is_none")]
    pub destinations: Option<Vec<serde_json::Value>>,
    #[serde(rename = "DriverLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions", skip_serializing_if = "Option::is_none")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber", skip_serializing_if = "Option::is_none")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake", skip_serializing_if = "Option::is_none")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel", skip_serializing_if = "Option::is_none")]
    pub vehicle_model: Option<String>,
}
