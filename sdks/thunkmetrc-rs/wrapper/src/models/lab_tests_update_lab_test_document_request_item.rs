use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct LabTestsUpdateLabTestDocumentRequestItem {
    #[serde(rename = "document_file_base64", skip_serializing_if = "Option::is_none")]
    pub document_file_base64: Option<String>,
    #[serde(rename = "document_file_name", skip_serializing_if = "Option::is_none")]
    pub document_file_name: Option<String>,
    #[serde(rename = "lab_test_result_id", skip_serializing_if = "Option::is_none")]
    pub lab_test_result_id: Option<i64>,
}
