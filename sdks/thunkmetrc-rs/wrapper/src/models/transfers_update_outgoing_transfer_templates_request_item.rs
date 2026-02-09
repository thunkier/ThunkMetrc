use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TransfersUpdateOutgoingTransferTemplatesRequestItem {
    #[serde(rename = "destinations", skip_serializing_if = "Option::is_none")]
    pub destinations: Option<Vec<TransfersUpdateOutgoingTransferTemplatesRequestItemDestinationsItem>>,
    #[serde(rename = "driver_license_number", skip_serializing_if = "Option::is_none")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "driver_name", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "driver_occupational_license_number", skip_serializing_if = "Option::is_none")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "phone_number_for_questions", skip_serializing_if = "Option::is_none")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "transfer_template_id", skip_serializing_if = "Option::is_none")]
    pub transfer_template_id: Option<i64>,
    #[serde(rename = "transporter_facility_license_number", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "vehicle_license_plate_number", skip_serializing_if = "Option::is_none")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "vehicle_make", skip_serializing_if = "Option::is_none")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "vehicle_model", skip_serializing_if = "Option::is_none")]
    pub vehicle_model: Option<String>,
}
