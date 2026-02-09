use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ProcessingJobStartProcessingJobRequestItem {
    #[serde(rename = "count_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub count_unit_of_measure: Option<String>,
    #[serde(rename = "job_name", skip_serializing_if = "Option::is_none")]
    pub job_name: Option<String>,
    #[serde(rename = "job_type", skip_serializing_if = "Option::is_none")]
    pub job_type: Option<String>,
    #[serde(rename = "packages", skip_serializing_if = "Option::is_none")]
    pub packages: Option<Vec<ProcessingJobStartProcessingJobRequestItemPackagesItem>>,
    #[serde(rename = "start_date", skip_serializing_if = "Option::is_none")]
    pub start_date: Option<String>,
    #[serde(rename = "volume_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub volume_unit_of_measure: Option<String>,
    #[serde(rename = "weight_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub weight_unit_of_measure: Option<String>,
}
