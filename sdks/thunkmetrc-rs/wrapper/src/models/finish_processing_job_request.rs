use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct FinishProcessingJobRequest {
    #[serde(rename = "FinishDate", skip_serializing_if = "Option::is_none")]
    pub finish_date: Option<String>,
    #[serde(rename = "FinishNote", skip_serializing_if = "Option::is_none")]
    pub finish_note: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "TotalCountWaste", skip_serializing_if = "Option::is_none")]
    pub total_count_waste: Option<String>,
    #[serde(rename = "TotalVolumeWaste", skip_serializing_if = "Option::is_none")]
    pub total_volume_waste: Option<String>,
    #[serde(rename = "TotalWeightWaste", skip_serializing_if = "Option::is_none")]
    pub total_weight_waste: Option<i64>,
    #[serde(rename = "WasteCountUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub waste_count_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteVolumeUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub waste_volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteWeightUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub waste_weight_unit_of_measure_name: Option<String>,
}
