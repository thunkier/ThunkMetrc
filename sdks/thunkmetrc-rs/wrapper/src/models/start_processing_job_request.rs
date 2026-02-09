use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct StartProcessingJobRequest {
    #[serde(rename = "CountUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub count_unit_of_measure: Option<String>,
    #[serde(rename = "JobName", skip_serializing_if = "Option::is_none")]
    pub job_name: Option<String>,
    #[serde(rename = "JobType", skip_serializing_if = "Option::is_none")]
    pub job_type: Option<String>,
    #[serde(rename = "Packages", skip_serializing_if = "Option::is_none")]
    pub packages: Option<Vec<serde_json::Value>>,
    #[serde(rename = "StartDate", skip_serializing_if = "Option::is_none")]
    pub start_date: Option<String>,
    #[serde(rename = "VolumeUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub volume_unit_of_measure: Option<String>,
    #[serde(rename = "WeightUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub weight_unit_of_measure: Option<String>,
}
