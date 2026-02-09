use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Hub {
    #[serde(rename = "ActualArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_arrival_date_time: Option<String>,
    #[serde(rename = "ActualDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_departure_date_time: Option<String>,
    #[serde(rename = "ActualReturnArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_return_arrival_date_time: Option<String>,
    #[serde(rename = "ActualReturnDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_return_departure_date_time: Option<String>,
    #[serde(rename = "CreatedByUserName", skip_serializing_if = "Option::is_none")]
    pub created_by_user_name: Option<String>,
    #[serde(rename = "CreatedDateTime", skip_serializing_if = "Option::is_none")]
    pub created_date_time: Option<String>,
    #[serde(rename = "DeliveryCount", skip_serializing_if = "Option::is_none")]
    pub delivery_count: Option<i64>,
    #[serde(rename = "DeliveryId", skip_serializing_if = "Option::is_none")]
    pub delivery_id: Option<i64>,
    #[serde(rename = "DeliveryPackageCount", skip_serializing_if = "Option::is_none")]
    pub delivery_package_count: Option<i64>,
    #[serde(rename = "DeliveryReceivedPackageCount", skip_serializing_if = "Option::is_none")]
    pub delivery_received_package_count: Option<i64>,
    #[serde(rename = "DriverName", skip_serializing_if = "Option::is_none")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "DriverVehicleLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub driver_vehicle_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "EstimatedReturnArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_return_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedReturnDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_return_departure_date_time: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "IsLayover", skip_serializing_if = "Option::is_none")]
    pub is_layover: Option<bool>,
    #[serde(rename = "IsVoided", skip_serializing_if = "Option::is_none")]
    pub is_voided: Option<bool>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "ManifestNumber", skip_serializing_if = "Option::is_none")]
    pub manifest_number: Option<String>,
    #[serde(rename = "PackageCount", skip_serializing_if = "Option::is_none")]
    pub package_count: Option<i64>,
    #[serde(rename = "ReceivedDateTime", skip_serializing_if = "Option::is_none")]
    pub received_date_time: Option<String>,
    #[serde(rename = "ReceivedDeliveryCount", skip_serializing_if = "Option::is_none")]
    pub received_delivery_count: Option<i64>,
    #[serde(rename = "ReceivedPackageCount", skip_serializing_if = "Option::is_none")]
    pub received_package_count: Option<i64>,
    #[serde(rename = "RecipientFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub recipient_facility_license_number: Option<String>,
    #[serde(rename = "RecipientFacilityName", skip_serializing_if = "Option::is_none")]
    pub recipient_facility_name: Option<String>,
    #[serde(rename = "RejectedPackagesReturned", skip_serializing_if = "Option::is_none")]
    pub rejected_packages_returned: Option<bool>,
    #[serde(rename = "ShipmentTransactionType", skip_serializing_if = "Option::is_none")]
    pub shipment_transaction_type: Option<i64>,
    #[serde(rename = "ShipmentTransporterDetails", skip_serializing_if = "Option::is_none")]
    pub shipment_transporter_details: Option<Vec<TransfersHubShipmentTransporterDetailsItem>>,
    #[serde(rename = "ShipmentTypeName", skip_serializing_if = "Option::is_none")]
    pub shipment_type_name: Option<String>,
    #[serde(rename = "ShipperFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub shipper_facility_license_number: Option<String>,
    #[serde(rename = "ShipperFacilityName", skip_serializing_if = "Option::is_none")]
    pub shipper_facility_name: Option<String>,
    #[serde(rename = "TransporterAcceptedDateTime", skip_serializing_if = "Option::is_none")]
    pub transporter_accepted_date_time: Option<String>,
    #[serde(rename = "TransporterActualArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub transporter_actual_arrival_date_time: Option<String>,
    #[serde(rename = "TransporterActualDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub transporter_actual_departure_date_time: Option<String>,
    #[serde(rename = "TransporterEstimatedArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub transporter_estimated_arrival_date_time: Option<String>,
    #[serde(rename = "TransporterEstimatedDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub transporter_estimated_departure_date_time: Option<String>,
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
}
