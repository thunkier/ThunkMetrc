use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct LabTestsCreateLabTestRecordRequestItem {
    #[serde(rename = "document_file_base64", skip_serializing_if = "Option::is_none")]
    pub document_file_base64: Option<String>,
    #[serde(rename = "document_file_name", skip_serializing_if = "Option::is_none")]
    pub document_file_name: Option<String>,
    #[serde(rename = "label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "result_date", skip_serializing_if = "Option::is_none")]
    pub result_date: Option<String>,
    #[serde(rename = "results", skip_serializing_if = "Option::is_none")]
    pub results: Option<Vec<LabTestsCreateLabTestRecordRequestItemResultsItem>>,
}
