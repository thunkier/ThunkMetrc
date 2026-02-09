use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsDeletePlantsRequestItem {
    #[serde(rename = "actual_date", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "count", skip_serializing_if = "Option::is_none")]
    pub count: Option<i64>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "reason_note", skip_serializing_if = "Option::is_none")]
    pub reason_note: Option<String>,
    #[serde(rename = "waste_material_mixed", skip_serializing_if = "Option::is_none")]
    pub waste_material_mixed: Option<String>,
    #[serde(rename = "waste_method_name", skip_serializing_if = "Option::is_none")]
    pub waste_method_name: Option<String>,
    #[serde(rename = "waste_reason_name", skip_serializing_if = "Option::is_none")]
    pub waste_reason_name: Option<String>,
    #[serde(rename = "waste_unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub waste_unit_of_measure_name: Option<String>,
    #[serde(rename = "waste_weight", skip_serializing_if = "Option::is_none")]
    pub waste_weight: Option<f64>,
}
