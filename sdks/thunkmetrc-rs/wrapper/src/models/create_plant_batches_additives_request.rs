use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreatePlantBatchesAdditivesRequest {
    #[serde(rename = "ActiveIngredients", skip_serializing_if = "Option::is_none")]
    pub active_ingredients: Option<Vec<serde_json::Value>>,
    #[serde(rename = "ActualDate", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "AdditiveType", skip_serializing_if = "Option::is_none")]
    pub additive_type: Option<String>,
    #[serde(rename = "ApplicationDevice", skip_serializing_if = "Option::is_none")]
    pub application_device: Option<String>,
    #[serde(rename = "EpaRegistrationNumber", skip_serializing_if = "Option::is_none")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "PlantBatchName", skip_serializing_if = "Option::is_none")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "ProductSupplier", skip_serializing_if = "Option::is_none")]
    pub product_supplier: Option<String>,
    #[serde(rename = "ProductTradeName", skip_serializing_if = "Option::is_none")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "TotalAmountApplied", skip_serializing_if = "Option::is_none")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "TotalAmountUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub total_amount_unit_of_measure: Option<String>,
}
