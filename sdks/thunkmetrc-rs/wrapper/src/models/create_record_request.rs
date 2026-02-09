use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateRecordRequest {
    #[serde(rename = "DocumentFileBase64", skip_serializing_if = "Option::is_none")]
    pub document_file_base64: Option<String>,
    #[serde(rename = "DocumentFileName", skip_serializing_if = "Option::is_none")]
    pub document_file_name: Option<String>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "ResultDate", skip_serializing_if = "Option::is_none")]
    pub result_date: Option<String>,
    #[serde(rename = "Results", skip_serializing_if = "Option::is_none")]
    pub results: Option<Vec<serde_json::Value>>,
}
