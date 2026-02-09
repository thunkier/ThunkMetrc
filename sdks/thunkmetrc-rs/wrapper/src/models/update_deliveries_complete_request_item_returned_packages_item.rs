use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateDeliveriesCompleteRequestItemReturnedPackagesItem {
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "ReturnQuantityVerified", skip_serializing_if = "Option::is_none")]
    pub return_quantity_verified: Option<i64>,
    #[serde(rename = "ReturnReason", skip_serializing_if = "Option::is_none")]
    pub return_reason: Option<String>,
    #[serde(rename = "ReturnReasonNote", skip_serializing_if = "Option::is_none")]
    pub return_reason_note: Option<String>,
    #[serde(rename = "ReturnUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub return_unit_of_measure: Option<String>,
}
