use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantAdditive {
    #[serde(rename = "AdditiveTypeName", skip_serializing_if = "Option::is_none")]
    pub additive_type_name: Option<String>,
    #[serde(rename = "AmountUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub amount_unit_of_measure: Option<String>,
    #[serde(rename = "ApplicationDevice", skip_serializing_if = "Option::is_none")]
    pub application_device: Option<String>,
    #[serde(rename = "EpaRegistrationNumber", skip_serializing_if = "Option::is_none")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "Note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "PlantBatchId", skip_serializing_if = "Option::is_none")]
    pub plant_batch_id: Option<i64>,
    #[serde(rename = "PlantBatchName", skip_serializing_if = "Option::is_none")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "PlantCount", skip_serializing_if = "Option::is_none")]
    pub plant_count: Option<i64>,
    #[serde(rename = "ProductSupplier", skip_serializing_if = "Option::is_none")]
    pub product_supplier: Option<String>,
    #[serde(rename = "ProductTradeName", skip_serializing_if = "Option::is_none")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "Rate", skip_serializing_if = "Option::is_none")]
    pub rate: Option<String>,
    #[serde(rename = "RestrictiveEntryIntervalQuantityDescription", skip_serializing_if = "Option::is_none")]
    pub restrictive_entry_interval_quantity_description: Option<String>,
    #[serde(rename = "RestrictiveEntryIntervalTimeDescription", skip_serializing_if = "Option::is_none")]
    pub restrictive_entry_interval_time_description: Option<String>,
    #[serde(rename = "TotalAmountApplied", skip_serializing_if = "Option::is_none")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "Volume", skip_serializing_if = "Option::is_none")]
    pub volume: Option<f64>,
}
