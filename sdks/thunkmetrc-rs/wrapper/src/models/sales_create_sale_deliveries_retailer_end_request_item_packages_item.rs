use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesCreateSaleDeliveriesRetailerEndRequestItemPackagesItem {
    #[serde(rename = "end_quantity", skip_serializing_if = "Option::is_none")]
    pub end_quantity: Option<i64>,
    #[serde(rename = "end_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub end_unit_of_measure: Option<String>,
    #[serde(rename = "label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
}
