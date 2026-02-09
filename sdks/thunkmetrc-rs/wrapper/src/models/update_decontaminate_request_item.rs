use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateDecontaminateRequestItem {
    #[serde(rename = "DecontaminationDate", skip_serializing_if = "Option::is_none")]
    pub decontamination_date: Option<String>,
    #[serde(rename = "DecontaminationMethodName", skip_serializing_if = "Option::is_none")]
    pub decontamination_method_name: Option<String>,
    #[serde(rename = "DecontaminationSteps", skip_serializing_if = "Option::is_none")]
    pub decontamination_steps: Option<String>,
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
}
