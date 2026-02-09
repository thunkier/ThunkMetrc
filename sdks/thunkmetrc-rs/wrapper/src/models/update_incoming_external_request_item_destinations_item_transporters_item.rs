use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateIncomingExternalRequestItemDestinationsItemTransportersItem {
    #[serde(rename = "DriverLayoverLeg", skip_serializing_if = "Option::is_none")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "IsLayover", skip_serializing_if = "Option::is_none")]
    pub is_layover: Option<bool>,
    #[serde(rename = "PhoneNumberForQuestions", skip_serializing_if = "Option::is_none")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterDetails", skip_serializing_if = "Option::is_none")]
    pub transporter_details: Option<Vec<serde_json::Value>>,
    #[serde(rename = "TransporterFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber", skip_serializing_if = "Option::is_none")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake", skip_serializing_if = "Option::is_none")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel", skip_serializing_if = "Option::is_none")]
    pub vehicle_model: Option<String>,
}
