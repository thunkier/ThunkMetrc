use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateAdjustPackagesRequestItem {
    #[serde(rename = "AdjustmentDate", skip_serializing_if = "Option::is_none")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "AdjustmentReason", skip_serializing_if = "Option::is_none")]
    pub adjustment_reason: Option<String>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "Quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<f64>,
    #[serde(rename = "ReasonNote", skip_serializing_if = "Option::is_none")]
    pub reason_note: Option<String>,
    #[serde(rename = "UnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure: Option<String>,
}
