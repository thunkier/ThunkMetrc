use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreatePlantBatchesWasteRequest {
    #[serde(rename = "MixedMaterial", skip_serializing_if = "Option::is_none")]
    pub mixed_material: Option<String>,
    #[serde(rename = "Note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "PlantBatchName", skip_serializing_if = "Option::is_none")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "ReasonName", skip_serializing_if = "Option::is_none")]
    pub reason_name: Option<String>,
    #[serde(rename = "UnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteDate", skip_serializing_if = "Option::is_none")]
    pub waste_date: Option<String>,
    #[serde(rename = "WasteMethodName", skip_serializing_if = "Option::is_none")]
    pub waste_method_name: Option<String>,
    #[serde(rename = "WasteWeight", skip_serializing_if = "Option::is_none")]
    pub waste_weight: Option<f64>,
}
