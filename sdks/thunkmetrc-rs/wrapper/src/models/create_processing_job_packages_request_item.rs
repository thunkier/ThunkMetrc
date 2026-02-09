use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateProcessingJobPackagesRequestItem {
    #[serde(rename = "ExpirationDate", skip_serializing_if = "Option::is_none")]
    pub expiration_date: Option<String>,
    #[serde(rename = "FinishDate", skip_serializing_if = "Option::is_none")]
    pub finish_date: Option<String>,
    #[serde(rename = "FinishNote", skip_serializing_if = "Option::is_none")]
    pub finish_note: Option<String>,
    #[serde(rename = "FinishProcessingJob", skip_serializing_if = "Option::is_none")]
    pub finish_processing_job: Option<bool>,
    #[serde(rename = "IsFinishedGood", skip_serializing_if = "Option::is_none")]
    pub is_finished_good: Option<bool>,
    #[serde(rename = "Item", skip_serializing_if = "Option::is_none")]
    pub item: Option<String>,
    #[serde(rename = "JobName", skip_serializing_if = "Option::is_none")]
    pub job_name: Option<String>,
    #[serde(rename = "Location", skip_serializing_if = "Option::is_none")]
    pub location: Option<String>,
    #[serde(rename = "Note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "PackageDate", skip_serializing_if = "Option::is_none")]
    pub package_date: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProductionBatchNumber", skip_serializing_if = "Option::is_none")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "Quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<f64>,
    #[serde(rename = "SellByDate", skip_serializing_if = "Option::is_none")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation", skip_serializing_if = "Option::is_none")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag", skip_serializing_if = "Option::is_none")]
    pub tag: Option<String>,
    #[serde(rename = "UnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UseByDate", skip_serializing_if = "Option::is_none")]
    pub use_by_date: Option<String>,
    #[serde(rename = "WasteCountQuantity", skip_serializing_if = "Option::is_none")]
    pub waste_count_quantity: Option<f64>,
    #[serde(rename = "WasteCountUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub waste_count_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteVolumeQuantity", skip_serializing_if = "Option::is_none")]
    pub waste_volume_quantity: Option<f64>,
    #[serde(rename = "WasteVolumeUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub waste_volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteWeightQuantity", skip_serializing_if = "Option::is_none")]
    pub waste_weight_quantity: Option<f64>,
    #[serde(rename = "WasteWeightUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub waste_weight_unit_of_measure_name: Option<String>,
}
