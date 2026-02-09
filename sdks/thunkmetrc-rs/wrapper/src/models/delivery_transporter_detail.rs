use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DeliveryTransporterDetail {
    #[serde(rename = "ActualDriverStartDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_driver_start_date_time: Option<String>,
    #[serde(rename = "DriverLayoverLeg", skip_serializing_if = "Option::is_none")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverName", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "DriverVehicleLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub driver_vehicle_license_number: Option<String>,
    #[serde(rename = "ShipmentDeliveryId", skip_serializing_if = "Option::is_none")]
    pub shipment_delivery_id: Option<i64>,
    #[serde(rename = "VehicleLicensePlateNumber", skip_serializing_if = "Option::is_none")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake", skip_serializing_if = "Option::is_none")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel", skip_serializing_if = "Option::is_none")]
    pub vehicle_model: Option<String>,
}
