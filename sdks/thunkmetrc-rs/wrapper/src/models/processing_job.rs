use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ProcessingJob {
    #[serde(rename = "CountUnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub count_unit_of_measure_abbreviation: Option<String>,
    #[serde(rename = "CountUnitOfMeasureId", skip_serializing_if = "Option::is_none")]
    pub count_unit_of_measure_id: Option<i64>,
    #[serde(rename = "CountUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub count_unit_of_measure_name: Option<String>,
    #[serde(rename = "FinishNote", skip_serializing_if = "Option::is_none")]
    pub finish_note: Option<String>,
    #[serde(rename = "FinishedDate", skip_serializing_if = "Option::is_none")]
    pub finished_date: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "IsFinished", skip_serializing_if = "Option::is_none")]
    pub is_finished: Option<bool>,
    #[serde(rename = "JobTypeId", skip_serializing_if = "Option::is_none")]
    pub job_type_id: Option<i64>,
    #[serde(rename = "JobTypeName", skip_serializing_if = "Option::is_none")]
    pub job_type_name: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "Number", skip_serializing_if = "Option::is_none")]
    pub number: Option<String>,
    #[serde(rename = "Packages", skip_serializing_if = "Option::is_none")]
    pub packages: Option<Vec<serde_json::Value>>,
    #[serde(rename = "StartDate", skip_serializing_if = "Option::is_none")]
    pub start_date: Option<String>,
    #[serde(rename = "TotalCount", skip_serializing_if = "Option::is_none")]
    pub total_count: Option<i64>,
    #[serde(rename = "TotalCountWaste", skip_serializing_if = "Option::is_none")]
    pub total_count_waste: Option<String>,
    #[serde(rename = "TotalQuantity", skip_serializing_if = "Option::is_none")]
    pub total_quantity: Option<f64>,
    #[serde(rename = "TotalUnitOfMeasureId", skip_serializing_if = "Option::is_none")]
    pub total_unit_of_measure_id: Option<i64>,
    #[serde(rename = "TotalVolume", skip_serializing_if = "Option::is_none")]
    pub total_volume: Option<f64>,
    #[serde(rename = "TotalVolumeWaste", skip_serializing_if = "Option::is_none")]
    pub total_volume_waste: Option<String>,
    #[serde(rename = "TotalWeight", skip_serializing_if = "Option::is_none")]
    pub total_weight: Option<f64>,
    #[serde(rename = "TotalWeightWaste", skip_serializing_if = "Option::is_none")]
    pub total_weight_waste: Option<String>,
    #[serde(rename = "VolumeUnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub volume_unit_of_measure_abbreviation: Option<String>,
    #[serde(rename = "VolumeUnitOfMeasureId", skip_serializing_if = "Option::is_none")]
    pub volume_unit_of_measure_id: Option<i64>,
    #[serde(rename = "VolumeUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteCountUnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub waste_count_unit_of_measure_abbreviation: Option<String>,
    #[serde(rename = "WasteCountUnitOfMeasureId", skip_serializing_if = "Option::is_none")]
    pub waste_count_unit_of_measure_id: Option<i64>,
    #[serde(rename = "WasteCountUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub waste_count_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteVolumeUnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub waste_volume_unit_of_measure_abbreviation: Option<String>,
    #[serde(rename = "WasteVolumeUnitOfMeasureId", skip_serializing_if = "Option::is_none")]
    pub waste_volume_unit_of_measure_id: Option<i64>,
    #[serde(rename = "WasteVolumeUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub waste_volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteWeightUnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub waste_weight_unit_of_measure_abbreviation: Option<String>,
    #[serde(rename = "WasteWeightUnitOfMeasureId", skip_serializing_if = "Option::is_none")]
    pub waste_weight_unit_of_measure_id: Option<i64>,
    #[serde(rename = "WasteWeightUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub waste_weight_unit_of_measure_name: Option<String>,
    #[serde(rename = "WeightUnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub weight_unit_of_measure_abbreviation: Option<String>,
    #[serde(rename = "WeightUnitOfMeasureId", skip_serializing_if = "Option::is_none")]
    pub weight_unit_of_measure_id: Option<i64>,
    #[serde(rename = "WeightUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub weight_unit_of_measure_name: Option<String>,
}
