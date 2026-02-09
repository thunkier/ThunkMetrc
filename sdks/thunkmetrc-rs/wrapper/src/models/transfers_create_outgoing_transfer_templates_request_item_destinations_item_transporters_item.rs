use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TransfersCreateOutgoingTransferTemplatesRequestItemDestinationsItemTransportersItem {
    #[serde(rename = "driver_layover_leg", skip_serializing_if = "Option::is_none")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "driver_license_number", skip_serializing_if = "Option::is_none")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "driver_name", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "driver_occupational_license_number", skip_serializing_if = "Option::is_none")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "estimated_arrival_date_time", skip_serializing_if = "Option::is_none")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "estimated_departure_date_time", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "is_layover", skip_serializing_if = "Option::is_none")]
    pub is_layover: Option<bool>,
    #[serde(rename = "phone_number_for_questions", skip_serializing_if = "Option::is_none")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "transporter_details", skip_serializing_if = "Option::is_none")]
    pub transporter_details: Option<Vec<TransfersCreateOutgoingTransferTemplatesRequestItemDestinationsItemTransportersItemTransporterDetailsItem>>,
    #[serde(rename = "transporter_facility_license_number", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "vehicle_license_plate_number", skip_serializing_if = "Option::is_none")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "vehicle_make", skip_serializing_if = "Option::is_none")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "vehicle_model", skip_serializing_if = "Option::is_none")]
    pub vehicle_model: Option<String>,
}
