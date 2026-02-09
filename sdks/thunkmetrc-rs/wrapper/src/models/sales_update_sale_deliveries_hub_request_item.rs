use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesUpdateSaleDeliveriesHubRequestItem {
    #[serde(rename = "driver_employee_id", skip_serializing_if = "Option::is_none")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "driver_name", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "drivers_license_number", skip_serializing_if = "Option::is_none")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "estimated_arrival_date_time", skip_serializing_if = "Option::is_none")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "estimated_departure_date_time", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "phone_number_for_questions", skip_serializing_if = "Option::is_none")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "planned_route", skip_serializing_if = "Option::is_none")]
    pub planned_route: Option<String>,
    #[serde(rename = "transporter_facility_id", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_id: Option<String>,
    #[serde(rename = "vehicle_license_plate_number", skip_serializing_if = "Option::is_none")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "vehicle_make", skip_serializing_if = "Option::is_none")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "vehicle_model", skip_serializing_if = "Option::is_none")]
    pub vehicle_model: Option<String>,
}
