use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantBatchesCreateAdjustPlantBatchesRequestItem {
    #[serde(rename = "adjustment_date", skip_serializing_if = "Option::is_none")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "adjustment_reason", skip_serializing_if = "Option::is_none")]
    pub adjustment_reason: Option<String>,
    #[serde(rename = "plant_batch_name", skip_serializing_if = "Option::is_none")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<i64>,
    #[serde(rename = "reason_note", skip_serializing_if = "Option::is_none")]
    pub reason_note: Option<String>,
}
