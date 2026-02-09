use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Receive {
    #[serde(rename = "ChildTag", skip_serializing_if = "Option::is_none")]
    pub child_tag: Option<String>,
    #[serde(rename = "Eaches", skip_serializing_if = "Option::is_none")]
    pub eaches: Option<Vec<String>>,
    #[serde(rename = "LabelSource", skip_serializing_if = "Option::is_none")]
    pub label_source: Option<String>,
    #[serde(rename = "QrCount", skip_serializing_if = "Option::is_none")]
    pub qr_count: Option<i64>,
    #[serde(rename = "Ranges", skip_serializing_if = "Option::is_none")]
    pub ranges: Option<Vec<Vec<i64>>>,
    #[serde(rename = "RequiresVerification", skip_serializing_if = "Option::is_none")]
    pub requires_verification: Option<bool>,
    #[serde(rename = "SiblingTags", skip_serializing_if = "Option::is_none")]
    pub sibling_tags: Option<Vec<String>>,
}
