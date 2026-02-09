use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TransfersDeliveryPackageProductLabel {
    #[serde(rename = "IsActive", skip_serializing_if = "Option::is_none")]
    pub is_active: Option<bool>,
    #[serde(rename = "IsChildFromParentWithLabel", skip_serializing_if = "Option::is_none")]
    pub is_child_from_parent_with_label: Option<bool>,
    #[serde(rename = "LabelSource", skip_serializing_if = "Option::is_none")]
    pub label_source: Option<String>,
    #[serde(rename = "LastLabelGenerationDate", skip_serializing_if = "Option::is_none")]
    pub last_label_generation_date: Option<String>,
    #[serde(rename = "OriginalSourcePackageId", skip_serializing_if = "Option::is_none")]
    pub original_source_package_id: Option<i64>,
    #[serde(rename = "OriginalSourcePackageLabel", skip_serializing_if = "Option::is_none")]
    pub original_source_package_label: Option<String>,
    #[serde(rename = "QrCodeRanges", skip_serializing_if = "Option::is_none")]
    pub qr_code_ranges: Option<String>,
    #[serde(rename = "QrCount", skip_serializing_if = "Option::is_none")]
    pub qr_count: Option<i64>,
}
