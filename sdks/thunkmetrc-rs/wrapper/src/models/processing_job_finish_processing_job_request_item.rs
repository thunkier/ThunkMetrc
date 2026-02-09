use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ProcessingJobFinishProcessingJobRequestItem {
    #[serde(rename = "finish_date", skip_serializing_if = "Option::is_none")]
    pub finish_date: Option<String>,
    #[serde(rename = "finish_note", skip_serializing_if = "Option::is_none")]
    pub finish_note: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "total_count_waste", skip_serializing_if = "Option::is_none")]
    pub total_count_waste: Option<String>,
    #[serde(rename = "total_volume_waste", skip_serializing_if = "Option::is_none")]
    pub total_volume_waste: Option<String>,
    #[serde(rename = "total_weight_waste", skip_serializing_if = "Option::is_none")]
    pub total_weight_waste: Option<i64>,
    #[serde(rename = "waste_count_unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub waste_count_unit_of_measure_name: Option<String>,
    #[serde(rename = "waste_volume_unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub waste_volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "waste_weight_unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub waste_weight_unit_of_measure_name: Option<String>,
}
