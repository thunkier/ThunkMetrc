use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Staged {
    #[serde(rename = "CommissionedDateTime", skip_serializing_if = "Option::is_none")]
    pub commissioned_date_time: Option<String>,
    #[serde(rename = "DetachedDateTime", skip_serializing_if = "Option::is_none")]
    pub detached_date_time: Option<String>,
    #[serde(rename = "FacilityId", skip_serializing_if = "Option::is_none")]
    pub facility_id: Option<i64>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "IsArchived", skip_serializing_if = "Option::is_none")]
    pub is_archived: Option<bool>,
    #[serde(rename = "IsStaged", skip_serializing_if = "Option::is_none")]
    pub is_staged: Option<bool>,
    #[serde(rename = "IsUsed", skip_serializing_if = "Option::is_none")]
    pub is_used: Option<bool>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "MaxGroupSize", skip_serializing_if = "Option::is_none")]
    pub max_group_size: Option<i64>,
    #[serde(rename = "ProductLabel", skip_serializing_if = "Option::is_none")]
    pub product_label: Option<String>,
    #[serde(rename = "QrCount", skip_serializing_if = "Option::is_none")]
    pub qr_count: Option<i64>,
    #[serde(rename = "StatusName", skip_serializing_if = "Option::is_none")]
    pub status_name: Option<String>,
    #[serde(rename = "TagInventoryTypeName", skip_serializing_if = "Option::is_none")]
    pub tag_inventory_type_name: Option<String>,
    #[serde(rename = "TagTypeId", skip_serializing_if = "Option::is_none")]
    pub tag_type_id: Option<i64>,
    #[serde(rename = "TagTypeName", skip_serializing_if = "Option::is_none")]
    pub tag_type_name: Option<String>,
    #[serde(rename = "UsedDateTime", skip_serializing_if = "Option::is_none")]
    pub used_date_time: Option<String>,
}
