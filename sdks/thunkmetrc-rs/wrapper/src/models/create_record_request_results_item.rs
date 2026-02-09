use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateRecordRequestResultsItem {
    #[serde(rename = "LabTestTypeName", skip_serializing_if = "Option::is_none")]
    pub lab_test_type_name: Option<String>,
    #[serde(rename = "Notes", skip_serializing_if = "Option::is_none")]
    pub notes: Option<String>,
    #[serde(rename = "Passed", skip_serializing_if = "Option::is_none")]
    pub passed: Option<bool>,
    #[serde(rename = "Quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<f64>,
}
