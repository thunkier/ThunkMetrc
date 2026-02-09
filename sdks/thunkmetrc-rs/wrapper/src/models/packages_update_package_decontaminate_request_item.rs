use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PackagesUpdatePackageDecontaminateRequestItem {
    #[serde(rename = "decontamination_date", skip_serializing_if = "Option::is_none")]
    pub decontamination_date: Option<String>,
    #[serde(rename = "decontamination_method_name", skip_serializing_if = "Option::is_none")]
    pub decontamination_method_name: Option<String>,
    #[serde(rename = "decontamination_steps", skip_serializing_if = "Option::is_none")]
    pub decontamination_steps: Option<String>,
    #[serde(rename = "package_label", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
}
