use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsCreatePlantAdditivesByLocationRequestItem {
    #[serde(rename = "active_ingredients", skip_serializing_if = "Option::is_none")]
    pub active_ingredients: Option<Vec<PlantsCreatePlantAdditivesByLocationRequestItemActiveIngredientsItem>>,
    #[serde(rename = "actual_date", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "additive_type", skip_serializing_if = "Option::is_none")]
    pub additive_type: Option<String>,
    #[serde(rename = "application_device", skip_serializing_if = "Option::is_none")]
    pub application_device: Option<String>,
    #[serde(rename = "epa_registration_number", skip_serializing_if = "Option::is_none")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "location_name", skip_serializing_if = "Option::is_none")]
    pub location_name: Option<String>,
    #[serde(rename = "product_supplier", skip_serializing_if = "Option::is_none")]
    pub product_supplier: Option<String>,
    #[serde(rename = "product_trade_name", skip_serializing_if = "Option::is_none")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "sublocation_name", skip_serializing_if = "Option::is_none")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "total_amount_applied", skip_serializing_if = "Option::is_none")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "total_amount_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub total_amount_unit_of_measure: Option<String>,
}
