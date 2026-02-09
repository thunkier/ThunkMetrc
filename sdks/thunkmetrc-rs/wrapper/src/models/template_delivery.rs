use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TemplateDelivery {
    #[serde(rename = "ActualArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_arrival_date_time: Option<String>,
    #[serde(rename = "ActualDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_departure_date_time: Option<String>,
    #[serde(rename = "ActualReturnArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_return_arrival_date_time: Option<String>,
    #[serde(rename = "ActualReturnDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_return_departure_date_time: Option<String>,
    #[serde(rename = "DeliveryNumber", skip_serializing_if = "Option::is_none")]
    pub delivery_number: Option<i64>,
    #[serde(rename = "DeliveryPackageCount", skip_serializing_if = "Option::is_none")]
    pub delivery_package_count: Option<i64>,
    #[serde(rename = "DeliveryReceivedPackageCount", skip_serializing_if = "Option::is_none")]
    pub delivery_received_package_count: Option<i64>,
    #[serde(rename = "EstimatedArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "EstimatedReturnArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_return_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedReturnDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_return_departure_date_time: Option<String>,
    #[serde(rename = "GrossUnitOfWeightId", skip_serializing_if = "Option::is_none")]
    pub gross_unit_of_weight_id: Option<i64>,
    #[serde(rename = "GrossUnitOfWeightName", skip_serializing_if = "Option::is_none")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "GrossWeight", skip_serializing_if = "Option::is_none")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "InvoiceNumber", skip_serializing_if = "Option::is_none")]
    pub invoice_number: Option<String>,
    #[serde(rename = "ManifestNumber", skip_serializing_if = "Option::is_none")]
    pub manifest_number: Option<String>,
    #[serde(rename = "PDFDocumentFileSystemId", skip_serializing_if = "Option::is_none")]
    pub pdf_document_file_system_id: Option<i64>,
    #[serde(rename = "PlannedRoute", skip_serializing_if = "Option::is_none")]
    pub planned_route: Option<String>,
    #[serde(rename = "ReceivedDateTime", skip_serializing_if = "Option::is_none")]
    pub received_date_time: Option<String>,
    #[serde(rename = "RecipientFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub recipient_facility_license_number: Option<String>,
    #[serde(rename = "RecipientFacilityName", skip_serializing_if = "Option::is_none")]
    pub recipient_facility_name: Option<String>,
    #[serde(rename = "RejectedPackagesReturned", skip_serializing_if = "Option::is_none")]
    pub rejected_packages_returned: Option<bool>,
    #[serde(rename = "ShipmentTransactionType", skip_serializing_if = "Option::is_none")]
    pub shipment_transaction_type: Option<String>,
    #[serde(rename = "ShipmentTypeName", skip_serializing_if = "Option::is_none")]
    pub shipment_type_name: Option<String>,
}
