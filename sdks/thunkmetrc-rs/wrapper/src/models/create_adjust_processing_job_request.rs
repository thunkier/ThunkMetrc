use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateAdjustProcessingJobRequest {
    #[serde(rename = "AdjustmentDate", skip_serializing_if = "Option::is_none")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "AdjustmentNote", skip_serializing_if = "Option::is_none")]
    pub adjustment_note: Option<String>,
    #[serde(rename = "AdjustmentReason", skip_serializing_if = "Option::is_none")]
    pub adjustment_reason: Option<String>,
    #[serde(rename = "CountUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub count_unit_of_measure_name: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "Packages", skip_serializing_if = "Option::is_none")]
    pub packages: Option<Vec<serde_json::Value>>,
    #[serde(rename = "VolumeUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "WeightUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub weight_unit_of_measure_name: Option<String>,
}
