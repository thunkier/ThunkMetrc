use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateDeliveriesRetailerEndRequestPackagesItem {
    #[serde(rename = "EndQuantity", skip_serializing_if = "Option::is_none")]
    pub end_quantity: Option<i64>,
    #[serde(rename = "EndUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub end_unit_of_measure: Option<String>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
}
