use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PackagesPackage {
    #[serde(rename = "ActualDepartureDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_departure_date_time: Option<String>,
    #[serde(rename = "ExternalId", skip_serializing_if = "Option::is_none")]
    pub external_id: Option<i64>,
    #[serde(rename = "GrossUnitOfWeightAbbreviation", skip_serializing_if = "Option::is_none")]
    pub gross_unit_of_weight_abbreviation: Option<String>,
    #[serde(rename = "GrossWeight", skip_serializing_if = "Option::is_none")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "ItemStrainName", skip_serializing_if = "Option::is_none")]
    pub item_strain_name: Option<String>,
    #[serde(rename = "LabTestingStateName", skip_serializing_if = "Option::is_none")]
    pub lab_testing_state_name: Option<String>,
    #[serde(rename = "ManifestNumber", skip_serializing_if = "Option::is_none")]
    pub manifest_number: Option<String>,
    #[serde(rename = "PackageId", skip_serializing_if = "Option::is_none")]
    pub package_id: Option<i64>,
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "ProcessingJobTypeName", skip_serializing_if = "Option::is_none")]
    pub processing_job_type_name: Option<String>,
    #[serde(rename = "ProductCategoryName", skip_serializing_if = "Option::is_none")]
    pub product_category_name: Option<String>,
    #[serde(rename = "ProductName", skip_serializing_if = "Option::is_none")]
    pub product_name: Option<String>,
    #[serde(rename = "ReceivedDateTime", skip_serializing_if = "Option::is_none")]
    pub received_date_time: Option<String>,
    #[serde(rename = "ReceivedQuantity", skip_serializing_if = "Option::is_none")]
    pub received_quantity: Option<f64>,
    #[serde(rename = "ReceivedUnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub received_unit_of_measure_abbreviation: Option<String>,
    #[serde(rename = "ReceiverWholesalePrice", skip_serializing_if = "Option::is_none")]
    pub receiver_wholesale_price: Option<f64>,
    #[serde(rename = "RecipientFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub recipient_facility_license_number: Option<String>,
    #[serde(rename = "RecipientFacilityName", skip_serializing_if = "Option::is_none")]
    pub recipient_facility_name: Option<String>,
    #[serde(rename = "ShipmentPackageStateName", skip_serializing_if = "Option::is_none")]
    pub shipment_package_state_name: Option<String>,
    #[serde(rename = "ShippedQuantity", skip_serializing_if = "Option::is_none")]
    pub shipped_quantity: Option<f64>,
    #[serde(rename = "ShippedUnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub shipped_unit_of_measure_abbreviation: Option<String>,
    #[serde(rename = "ShipperWholesalePrice", skip_serializing_if = "Option::is_none")]
    pub shipper_wholesale_price: Option<f64>,
    #[serde(rename = "SourceHarvestNames", skip_serializing_if = "Option::is_none")]
    pub source_harvest_names: Option<String>,
    #[serde(rename = "SourcePackageLabels", skip_serializing_if = "Option::is_none")]
    pub source_package_labels: Option<String>,
}
