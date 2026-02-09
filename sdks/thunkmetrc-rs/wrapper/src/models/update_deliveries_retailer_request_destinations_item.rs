use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateDeliveriesRetailerRequestDestinationsItem {
    #[serde(rename = "ConsumerId", skip_serializing_if = "Option::is_none")]
    pub consumer_id: Option<String>,
    #[serde(rename = "DriverEmployeeId", skip_serializing_if = "Option::is_none")]
    pub driver_employee_id: Option<i64>,
    #[serde(rename = "DriverName", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions", skip_serializing_if = "Option::is_none")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "PlannedRoute", skip_serializing_if = "Option::is_none")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientAddressCity", skip_serializing_if = "Option::is_none")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty", skip_serializing_if = "Option::is_none")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode", skip_serializing_if = "Option::is_none")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState", skip_serializing_if = "Option::is_none")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1", skip_serializing_if = "Option::is_none")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2", skip_serializing_if = "Option::is_none")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName", skip_serializing_if = "Option::is_none")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId", skip_serializing_if = "Option::is_none")]
    pub recipient_zone_id: Option<String>,
    #[serde(rename = "SalesCustomerType", skip_serializing_if = "Option::is_none")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime", skip_serializing_if = "Option::is_none")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions", skip_serializing_if = "Option::is_none")]
    pub transactions: Option<Vec<serde_json::Value>>,
    #[serde(rename = "VehicleLicensePlateNumber", skip_serializing_if = "Option::is_none")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake", skip_serializing_if = "Option::is_none")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel", skip_serializing_if = "Option::is_none")]
    pub vehicle_model: Option<String>,
}
