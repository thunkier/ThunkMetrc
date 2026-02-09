use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsUpdateAdjustPlantsRequestItem {
    #[serde(rename = "adjust_count", skip_serializing_if = "Option::is_none")]
    pub adjust_count: Option<i64>,
    #[serde(rename = "adjust_reason", skip_serializing_if = "Option::is_none")]
    pub adjust_reason: Option<String>,
    #[serde(rename = "adjustment_date", skip_serializing_if = "Option::is_none")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "reason_note", skip_serializing_if = "Option::is_none")]
    pub reason_note: Option<String>,
}
