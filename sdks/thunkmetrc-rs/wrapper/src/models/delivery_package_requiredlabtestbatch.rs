use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct DeliveryPackageRequiredlabtestbatch {
    #[serde(rename = "LabTestBatchId", skip_serializing_if = "Option::is_none")]
    pub lab_test_batch_id: Option<i64>,
    #[serde(rename = "LabTestBatchName", skip_serializing_if = "Option::is_none")]
    pub lab_test_batch_name: Option<String>,
    #[serde(rename = "PackageId", skip_serializing_if = "Option::is_none")]
    pub package_id: Option<i64>,
}
