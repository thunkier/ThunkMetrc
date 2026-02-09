use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateTestingRequestItem {
    #[serde(rename = "ActualDate", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "ExpirationDate", skip_serializing_if = "Option::is_none")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Ingredients", skip_serializing_if = "Option::is_none")]
    pub ingredients: Option<Vec<serde_json::Value>>,
    #[serde(rename = "IsDonation", skip_serializing_if = "Option::is_none")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsFinishedGood", skip_serializing_if = "Option::is_none")]
    pub is_finished_good: Option<bool>,
    #[serde(rename = "IsProductionBatch", skip_serializing_if = "Option::is_none")]
    pub is_production_batch: Option<bool>,
    #[serde(rename = "IsTradeSample", skip_serializing_if = "Option::is_none")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item", skip_serializing_if = "Option::is_none")]
    pub item: Option<String>,
    #[serde(rename = "LabTestStageId", skip_serializing_if = "Option::is_none")]
    pub lab_test_stage_id: Option<i64>,
    #[serde(rename = "Location", skip_serializing_if = "Option::is_none")]
    pub location: Option<String>,
    #[serde(rename = "Note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProcessingJobTypeId", skip_serializing_if = "Option::is_none")]
    pub processing_job_type_id: Option<i64>,
    #[serde(rename = "ProductRequiresRemediation", skip_serializing_if = "Option::is_none")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "ProductionBatchNumber", skip_serializing_if = "Option::is_none")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "Quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<f64>,
    #[serde(rename = "RequiredLabTestBatches", skip_serializing_if = "Option::is_none")]
    pub required_lab_test_batches: Option<Vec<serde_json::Value>>,
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
    #[serde(rename = "UseSameItem", skip_serializing_if = "Option::is_none")]
    pub use_same_item: Option<bool>,
}
