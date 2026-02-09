use serde::{Deserialize, Serialize};
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SaleDelivery {
    #[serde(rename = "AcceptedDateTime", skip_serializing_if = "Option::is_none")]
    pub accepted_date_time: Option<String>,
    #[serde(rename = "ActualArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_arrival_date_time: Option<String>,
    #[serde(rename = "ActualDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_departure_date_time: Option<String>,
    #[serde(rename = "ActualReturnArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_return_arrival_date_time: Option<String>,
    #[serde(rename = "ActualReturnDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_return_departure_date_time: Option<String>,
    #[serde(rename = "CompletedDateTime", skip_serializing_if = "Option::is_none")]
    pub completed_date_time: Option<String>,
    #[serde(rename = "ConsumerId", skip_serializing_if = "Option::is_none")]
    pub consumer_id: Option<String>,
    #[serde(rename = "DeliveryNumber", skip_serializing_if = "Option::is_none")]
    pub delivery_number: Option<String>,
    #[serde(rename = "Direction", skip_serializing_if = "Option::is_none")]
    pub direction: Option<String>,
    #[serde(rename = "DriverEmployeeId", skip_serializing_if = "Option::is_none")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "EstimatedReturnArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_return_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedReturnDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_return_departure_date_time: Option<String>,
    #[serde(rename = "FacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub facility_license_number: Option<String>,
    #[serde(rename = "FacilityName", skip_serializing_if = "Option::is_none")]
    pub facility_name: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
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
    #[serde(rename = "RecipientDeliveryZoneId", skip_serializing_if = "Option::is_none")]
    pub recipient_delivery_zone_id: Option<i64>,
    #[serde(rename = "RecipientDeliveryZoneName", skip_serializing_if = "Option::is_none")]
    pub recipient_delivery_zone_name: Option<String>,
    #[serde(rename = "RecipientName", skip_serializing_if = "Option::is_none")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId", skip_serializing_if = "Option::is_none")]
    pub recipient_zone_id: Option<i64>,
    #[serde(rename = "RecordedByUserName", skip_serializing_if = "Option::is_none")]
    pub recorded_by_user_name: Option<String>,
    #[serde(rename = "RecordedDateTime", skip_serializing_if = "Option::is_none")]
    pub recorded_date_time: Option<String>,
    #[serde(rename = "SalesCustomerType", skip_serializing_if = "Option::is_none")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime", skip_serializing_if = "Option::is_none")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "SalesDeliveryState", skip_serializing_if = "Option::is_none")]
    pub sales_delivery_state: Option<String>,
    #[serde(rename = "ShipperFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub shipper_facility_license_number: Option<String>,
    #[serde(rename = "TotalPackages", skip_serializing_if = "Option::is_none")]
    pub total_packages: Option<i64>,
    #[serde(rename = "TotalPrice", skip_serializing_if = "Option::is_none")]
    pub total_price: Option<i64>,
    #[serde(rename = "Transactions", skip_serializing_if = "Option::is_none")]
    pub transactions: Option<Vec<serde_json::Value>>,
    #[serde(rename = "TransporterFacilityId", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_id: Option<i64>,
    #[serde(rename = "TransporterFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "TransporterFacilityName", skip_serializing_if = "Option::is_none")]
    pub transporter_facility_name: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber", skip_serializing_if = "Option::is_none")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake", skip_serializing_if = "Option::is_none")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel", skip_serializing_if = "Option::is_none")]
    pub vehicle_model: Option<String>,
    #[serde(rename = "VoidedDate", skip_serializing_if = "Option::is_none")]
    pub voided_date: Option<String>,
}
