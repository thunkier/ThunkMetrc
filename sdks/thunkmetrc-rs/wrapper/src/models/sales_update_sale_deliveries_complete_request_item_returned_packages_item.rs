use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesUpdateSaleDeliveriesCompleteRequestItemReturnedPackagesItem {
    #[serde(rename = "label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "return_quantity_verified", skip_serializing_if = "Option::is_none")]
    pub return_quantity_verified: Option<i64>,
    #[serde(rename = "return_reason", skip_serializing_if = "Option::is_none")]
    pub return_reason: Option<String>,
    #[serde(rename = "return_reason_note", skip_serializing_if = "Option::is_none")]
    pub return_reason_note: Option<String>,
    #[serde(rename = "return_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub return_unit_of_measure: Option<String>,
}
