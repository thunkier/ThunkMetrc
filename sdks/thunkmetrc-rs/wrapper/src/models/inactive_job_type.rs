use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct InactiveJobType {
    #[serde(rename = "Attributes", skip_serializing_if = "Option::is_none")]
    pub attributes: Option<Vec<ProcessingJobInactiveJobTypeAttributesItem>>,
    #[serde(rename = "CategoryName", skip_serializing_if = "Option::is_none")]
    pub category_name: Option<String>,
    #[serde(rename = "Description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    #[serde(rename = "ForItems", skip_serializing_if = "Option::is_none")]
    pub for_items: Option<bool>,
    #[serde(rename = "ForProcessingJobs", skip_serializing_if = "Option::is_none")]
    pub for_processing_jobs: Option<bool>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "ProcessingSteps", skip_serializing_if = "Option::is_none")]
    pub processing_steps: Option<String>,
    #[serde(rename = "RequiresProcessingJobAttributes", skip_serializing_if = "Option::is_none")]
    pub requires_processing_job_attributes: Option<bool>,
}
