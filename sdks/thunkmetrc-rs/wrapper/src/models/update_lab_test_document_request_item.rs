use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateLabTestDocumentRequestItem {
    #[serde(rename = "DocumentFileBase64", skip_serializing_if = "Option::is_none")]
    pub document_file_base64: Option<String>,
    #[serde(rename = "DocumentFileName", skip_serializing_if = "Option::is_none")]
    pub document_file_name: Option<String>,
    #[serde(rename = "LabTestResultId", skip_serializing_if = "Option::is_none")]
    pub lab_test_result_id: Option<i64>,
}
