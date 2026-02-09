use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PackagesCreatePackagesRequestItem {
    #[serde(rename = "actual_date", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "expiration_date", skip_serializing_if = "Option::is_none")]
    pub expiration_date: Option<String>,
    #[serde(rename = "ingredients", skip_serializing_if = "Option::is_none")]
    pub ingredients: Option<Vec<PackagesCreatePackagesRequestItemIngredientsItem>>,
    #[serde(rename = "is_donation", skip_serializing_if = "Option::is_none")]
    pub is_donation: Option<bool>,
    #[serde(rename = "is_production_batch", skip_serializing_if = "Option::is_none")]
    pub is_production_batch: Option<bool>,
    #[serde(rename = "is_trade_sample", skip_serializing_if = "Option::is_none")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "item", skip_serializing_if = "Option::is_none")]
    pub item: Option<String>,
    #[serde(rename = "lab_test_stage_id", skip_serializing_if = "Option::is_none")]
    pub lab_test_stage_id: Option<i64>,
    #[serde(rename = "location", skip_serializing_if = "Option::is_none")]
    pub location: Option<String>,
    #[serde(rename = "note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "patient_license_number", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "processing_job_type_id", skip_serializing_if = "Option::is_none")]
    pub processing_job_type_id: Option<i64>,
    #[serde(rename = "product_requires_remediation", skip_serializing_if = "Option::is_none")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "production_batch_number", skip_serializing_if = "Option::is_none")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<i64>,
    #[serde(rename = "required_lab_test_batches", skip_serializing_if = "Option::is_none")]
    pub required_lab_test_batches: Option<bool>,
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
    #[serde(rename = "use_same_item", skip_serializing_if = "Option::is_none")]
    pub use_same_item: Option<bool>,
}
