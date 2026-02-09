use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateNoteRequest {
    #[serde(rename = "Note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
}
