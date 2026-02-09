use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AvailableTagPlant {
    #[serde(rename = "FacilityId", skip_serializing_if = "Option::is_none")]
    pub facility_id: Option<i64>,
    #[serde(rename = "GroupTagTypeId", skip_serializing_if = "Option::is_none")]
    pub group_tag_type_id: Option<String>,
    #[serde(rename = "GroupTagTypeName", skip_serializing_if = "Option::is_none")]
    pub group_tag_type_name: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "MaxGroupSize", skip_serializing_if = "Option::is_none")]
    pub max_group_size: Option<i64>,
    #[serde(rename = "TagInventoryTypeName", skip_serializing_if = "Option::is_none")]
    pub tag_inventory_type_name: Option<String>,
    #[serde(rename = "TagTypeId", skip_serializing_if = "Option::is_none")]
    pub tag_type_id: Option<String>,
    #[serde(rename = "TagTypeName", skip_serializing_if = "Option::is_none")]
    pub tag_type_name: Option<String>,
}
