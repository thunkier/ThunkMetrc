use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct RetailIdRetailIdPackagesInfoPackagesItemIssuancesItem {
    #[serde(rename = "CreatedAt", skip_serializing_if = "Option::is_none")]
    pub created_at: Option<String>,
    #[serde(rename = "LabelQuantity", skip_serializing_if = "Option::is_none")]
    pub label_quantity: Option<f64>,
    #[serde(rename = "LabelSet", skip_serializing_if = "Option::is_none")]
    pub label_set: Option<i64>,
}
