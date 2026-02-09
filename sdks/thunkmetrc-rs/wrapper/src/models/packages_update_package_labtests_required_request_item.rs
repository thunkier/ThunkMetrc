use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PackagesUpdatePackageLabtestsRequiredRequestItem {
    #[serde(rename = "label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "required_lab_test_batches", skip_serializing_if = "Option::is_none")]
    pub required_lab_test_batches: Option<Vec<String>>,
}
