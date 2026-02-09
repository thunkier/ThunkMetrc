use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ProcessingJobCreateProcessingJobPackagesRequestItem {
    #[serde(rename = "expiration_date", skip_serializing_if = "Option::is_none")]
    pub expiration_date: Option<String>,
    #[serde(rename = "finish_date", skip_serializing_if = "Option::is_none")]
    pub finish_date: Option<String>,
    #[serde(rename = "finish_note", skip_serializing_if = "Option::is_none")]
    pub finish_note: Option<String>,
    #[serde(rename = "finish_processing_job", skip_serializing_if = "Option::is_none")]
    pub finish_processing_job: Option<bool>,
    #[serde(rename = "item", skip_serializing_if = "Option::is_none")]
    pub item: Option<String>,
    #[serde(rename = "job_name", skip_serializing_if = "Option::is_none")]
    pub job_name: Option<String>,
    #[serde(rename = "location", skip_serializing_if = "Option::is_none")]
    pub location: Option<String>,
    #[serde(rename = "note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "package_date", skip_serializing_if = "Option::is_none")]
    pub package_date: Option<String>,
    #[serde(rename = "patient_license_number", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "production_batch_number", skip_serializing_if = "Option::is_none")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<i64>,
    #[serde(rename = "sell_by_date", skip_serializing_if = "Option::is_none")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "sublocation", skip_serializing_if = "Option::is_none")]
    pub sublocation: Option<String>,
    #[serde(rename = "tag", skip_serializing_if = "Option::is_none")]
    pub tag: Option<String>,
    #[serde(rename = "unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "use_by_date", skip_serializing_if = "Option::is_none")]
    pub use_by_date: Option<String>,
    #[serde(rename = "waste_count_quantity", skip_serializing_if = "Option::is_none")]
    pub waste_count_quantity: Option<String>,
    #[serde(rename = "waste_count_unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub waste_count_unit_of_measure_name: Option<String>,
    #[serde(rename = "waste_volume_quantity", skip_serializing_if = "Option::is_none")]
    pub waste_volume_quantity: Option<String>,
    #[serde(rename = "waste_volume_unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub waste_volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "waste_weight_quantity", skip_serializing_if = "Option::is_none")]
    pub waste_weight_quantity: Option<String>,
    #[serde(rename = "waste_weight_unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub waste_weight_unit_of_measure_name: Option<String>,
}
