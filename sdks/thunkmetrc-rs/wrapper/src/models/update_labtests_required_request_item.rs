use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateLabtestsRequiredRequestItem {
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "RequiredLabTestBatches", skip_serializing_if = "Option::is_none")]
    pub required_lab_test_batches: Option<Vec<serde_json::Value>>,
}
