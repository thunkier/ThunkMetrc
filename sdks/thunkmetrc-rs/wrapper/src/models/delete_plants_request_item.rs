use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DeletePlantsRequestItem {
    #[serde(rename = "ActualDate", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count", skip_serializing_if = "Option::is_none")]
    pub count: Option<i64>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "ReasonNote", skip_serializing_if = "Option::is_none")]
    pub reason_note: Option<String>,
    #[serde(rename = "WasteMaterialMixed", skip_serializing_if = "Option::is_none")]
    pub waste_material_mixed: Option<String>,
    #[serde(rename = "WasteMethodName", skip_serializing_if = "Option::is_none")]
    pub waste_method_name: Option<String>,
    #[serde(rename = "WasteReasonName", skip_serializing_if = "Option::is_none")]
    pub waste_reason_name: Option<String>,
    #[serde(rename = "WasteUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub waste_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteWeight", skip_serializing_if = "Option::is_none")]
    pub waste_weight: Option<f64>,
}
