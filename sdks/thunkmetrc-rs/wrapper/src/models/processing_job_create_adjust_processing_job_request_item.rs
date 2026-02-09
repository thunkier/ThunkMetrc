use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ProcessingJobCreateAdjustProcessingJobRequestItem {
    #[serde(rename = "adjustment_date", skip_serializing_if = "Option::is_none")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "adjustment_note", skip_serializing_if = "Option::is_none")]
    pub adjustment_note: Option<String>,
    #[serde(rename = "adjustment_reason", skip_serializing_if = "Option::is_none")]
    pub adjustment_reason: Option<String>,
    #[serde(rename = "count_unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub count_unit_of_measure_name: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "packages", skip_serializing_if = "Option::is_none")]
    pub packages: Option<Vec<ProcessingJobCreateAdjustProcessingJobRequestItemPackagesItem>>,
    #[serde(rename = "volume_unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "weight_unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub weight_unit_of_measure_name: Option<String>,
}
