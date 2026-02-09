use serde::{Deserialize, Serialize};
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PackagesPackageSummaryProductLabel {
    #[serde(rename = "IsActive", skip_serializing_if = "Option::is_none")]
    pub is_active: Option<bool>,
    #[serde(rename = "IsChildFromParentWithLabel", skip_serializing_if = "Option::is_none")]
    pub is_child_from_parent_with_label: Option<bool>,
    #[serde(rename = "IsDiscontinued", skip_serializing_if = "Option::is_none")]
    pub is_discontinued: Option<bool>,
    #[serde(rename = "LastLabelGenerationDate", skip_serializing_if = "Option::is_none")]
    pub last_label_generation_date: Option<String>,
    #[serde(rename = "PackageId", skip_serializing_if = "Option::is_none")]
    pub package_id: Option<i64>,
    #[serde(rename = "QrCount", skip_serializing_if = "Option::is_none")]
    pub qr_count: Option<i64>,
    #[serde(rename = "TagId", skip_serializing_if = "Option::is_none")]
    pub tag_id: Option<i64>,
}
