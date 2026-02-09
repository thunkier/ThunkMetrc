use serde::{Deserialize, Serialize};
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantWasteSummary {
    #[serde(rename = "PlantBatchId", skip_serializing_if = "Option::is_none")]
    pub plant_batch_id: Option<i64>,
    #[serde(rename = "PlantBatchName", skip_serializing_if = "Option::is_none")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "PlantCount", skip_serializing_if = "Option::is_none")]
    pub plant_count: Option<i64>,
    #[serde(rename = "PlantWasteNumber", skip_serializing_if = "Option::is_none")]
    pub plant_waste_number: Option<String>,
    #[serde(rename = "WasteDate", skip_serializing_if = "Option::is_none")]
    pub waste_date: Option<String>,
    #[serde(rename = "WasteMethodName", skip_serializing_if = "Option::is_none")]
    pub waste_method_name: Option<String>,
    #[serde(rename = "WasteReasonName", skip_serializing_if = "Option::is_none")]
    pub waste_reason_name: Option<String>,
    #[serde(rename = "WasteUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub waste_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteWeight", skip_serializing_if = "Option::is_none")]
    pub waste_weight: Option<f64>,
}
