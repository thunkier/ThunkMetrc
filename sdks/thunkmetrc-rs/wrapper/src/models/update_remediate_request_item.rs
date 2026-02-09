use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateRemediateRequestItem {
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "RemediationDate", skip_serializing_if = "Option::is_none")]
    pub remediation_date: Option<String>,
    #[serde(rename = "RemediationMethodName", skip_serializing_if = "Option::is_none")]
    pub remediation_method_name: Option<String>,
    #[serde(rename = "RemediationSteps", skip_serializing_if = "Option::is_none")]
    pub remediation_steps: Option<String>,
}
