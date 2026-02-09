use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateAdjustPlantsRequestItem {
    #[serde(rename = "AdjustCount", skip_serializing_if = "Option::is_none")]
    pub adjust_count: Option<i64>,
    #[serde(rename = "AdjustReason", skip_serializing_if = "Option::is_none")]
    pub adjust_reason: Option<String>,
    #[serde(rename = "AdjustmentDate", skip_serializing_if = "Option::is_none")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "ReasonNote", skip_serializing_if = "Option::is_none")]
    pub reason_note: Option<String>,
}
