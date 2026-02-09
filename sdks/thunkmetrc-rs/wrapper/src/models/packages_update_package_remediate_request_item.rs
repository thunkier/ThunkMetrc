use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PackagesUpdatePackageRemediateRequestItem {
    #[serde(rename = "package_label", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "remediation_date", skip_serializing_if = "Option::is_none")]
    pub remediation_date: Option<String>,
    #[serde(rename = "remediation_method_name", skip_serializing_if = "Option::is_none")]
    pub remediation_method_name: Option<String>,
    #[serde(rename = "remediation_steps", skip_serializing_if = "Option::is_none")]
    pub remediation_steps: Option<String>,
}
