use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SaleDeliveryRetailer {
    #[serde(rename = "AcceptedDateTime", skip_serializing_if = "Option::is_none")]
    pub accepted_date_time: Option<String>,
    #[serde(rename = "ActualDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_departure_date_time: Option<String>,
    #[serde(rename = "AllowFullEdit", skip_serializing_if = "Option::is_none")]
    pub allow_full_edit: Option<bool>,
    #[serde(rename = "CompletedDateTime", skip_serializing_if = "Option::is_none")]
    pub completed_date_time: Option<String>,
    #[serde(rename = "DateTime", skip_serializing_if = "Option::is_none")]
    pub date_time: Option<String>,
    #[serde(rename = "Destinations", skip_serializing_if = "Option::is_none")]
    pub destinations: Option<Vec<SalesSaleDeliveryRetailerDestinationsItem>>,
    #[serde(rename = "Direction", skip_serializing_if = "Option::is_none")]
    pub direction: Option<String>,
    #[serde(rename = "DriverEmployeeId", skip_serializing_if = "Option::is_none")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "FacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub facility_license_number: Option<String>,
    #[serde(rename = "FacilityName", skip_serializing_if = "Option::is_none")]
    pub facility_name: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "Leg", skip_serializing_if = "Option::is_none")]
    pub leg: Option<i64>,
    #[serde(rename = "Packages", skip_serializing_if = "Option::is_none")]
    pub packages: Option<Vec<SalesSaleDeliveryRetailerPackagesItem>>,
    #[serde(rename = "RecordedByUserName", skip_serializing_if = "Option::is_none")]
    pub recorded_by_user_name: Option<String>,
    #[serde(rename = "RecordedDateTime", skip_serializing_if = "Option::is_none")]
    pub recorded_date_time: Option<String>,
    #[serde(rename = "RestockDateTime", skip_serializing_if = "Option::is_none")]
    pub restock_date_time: Option<String>,
    #[serde(rename = "RetailerDeliveryNumber", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_number: Option<String>,
    #[serde(rename = "RetailerDeliveryState", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_state: Option<String>,
    #[serde(rename = "TotalPackages", skip_serializing_if = "Option::is_none")]
    pub total_packages: Option<i64>,
    #[serde(rename = "TotalPrice", skip_serializing_if = "Option::is_none")]
    pub total_price: Option<i64>,
    #[serde(rename = "TotalPriceSold", skip_serializing_if = "Option::is_none")]
    pub total_price_sold: Option<i64>,
    #[serde(rename = "VehicleInfo", skip_serializing_if = "Option::is_none")]
    pub vehicle_info: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber", skip_serializing_if = "Option::is_none")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake", skip_serializing_if = "Option::is_none")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel", skip_serializing_if = "Option::is_none")]
    pub vehicle_model: Option<String>,
    #[serde(rename = "VoidedDate", skip_serializing_if = "Option::is_none")]
    pub voided_date: Option<String>,
}
