use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Transfer {
    #[serde(rename = "ActualArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_arrival_date_time: Option<String>,
    #[serde(rename = "ActualDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_departure_date_time: Option<String>,
    #[serde(rename = "ActualReturnArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_return_arrival_date_time: Option<String>,
    #[serde(rename = "ActualReturnDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_return_departure_date_time: Option<String>,
    #[serde(rename = "ContainsDonation", skip_serializing_if = "Option::is_none")]
    pub contains_donation: Option<bool>,
    #[serde(rename = "ContainsPlantPackage", skip_serializing_if = "Option::is_none")]
    pub contains_plant_package: Option<bool>,
    #[serde(rename = "ContainsProductPackage", skip_serializing_if = "Option::is_none")]
    pub contains_product_package: Option<bool>,
    #[serde(rename = "ContainsProductRequiresRemediation", skip_serializing_if = "Option::is_none")]
    pub contains_product_requires_remediation: Option<bool>,
    #[serde(rename = "ContainsRemediatedProductPackage", skip_serializing_if = "Option::is_none")]
    pub contains_remediated_product_package: Option<bool>,
    #[serde(rename = "ContainsTestingSample", skip_serializing_if = "Option::is_none")]
    pub contains_testing_sample: Option<bool>,
    #[serde(rename = "ContainsTradeSample", skip_serializing_if = "Option::is_none")]
    pub contains_trade_sample: Option<bool>,
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
    #[serde(rename = "InvoiceNumber", skip_serializing_if = "Option::is_none")]
    pub invoice_number: Option<String>,
    #[serde(rename = "IsVoided", skip_serializing_if = "Option::is_none")]
    pub is_voided: Option<bool>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "ManifestNumber", skip_serializing_if = "Option::is_none")]
    pub manifest_number: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "OriginatingTemplateId", skip_serializing_if = "Option::is_none")]
    pub originating_template_id: Option<i64>,
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
    #[serde(rename = "ShipmentLicenseType", skip_serializing_if = "Option::is_none")]
    pub shipment_license_type: Option<i64>,
    #[serde(rename = "ShipmentTransactionType", skip_serializing_if = "Option::is_none")]
    pub shipment_transaction_type: Option<String>,
    #[serde(rename = "ShipmentTypeName", skip_serializing_if = "Option::is_none")]
    pub shipment_type_name: Option<String>,
    #[serde(rename = "ShipperFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub shipper_facility_license_number: Option<String>,
    #[serde(rename = "ShipperFacilityName", skip_serializing_if = "Option::is_none")]
    pub shipper_facility_name: Option<String>,
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
