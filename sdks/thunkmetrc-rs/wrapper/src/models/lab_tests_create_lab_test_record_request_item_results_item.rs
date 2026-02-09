use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct LabTestsCreateLabTestRecordRequestItemResultsItem {
    #[serde(rename = "lab_test_type_name", skip_serializing_if = "Option::is_none")]
    pub lab_test_type_name: Option<String>,
    #[serde(rename = "notes", skip_serializing_if = "Option::is_none")]
    pub notes: Option<String>,
    #[serde(rename = "passed", skip_serializing_if = "Option::is_none")]
    pub passed: Option<bool>,
    #[serde(rename = "quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<f64>,
}
