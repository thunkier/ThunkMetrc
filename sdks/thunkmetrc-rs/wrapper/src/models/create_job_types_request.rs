use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateJobTypesRequest {
    #[serde(rename = "Attributes", skip_serializing_if = "Option::is_none")]
    pub attributes: Option<Vec<serde_json::Value>>,
    #[serde(rename = "Category", skip_serializing_if = "Option::is_none")]
    pub category: Option<String>,
    #[serde(rename = "Description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "ProcessingSteps", skip_serializing_if = "Option::is_none")]
    pub processing_steps: Option<String>,
}
