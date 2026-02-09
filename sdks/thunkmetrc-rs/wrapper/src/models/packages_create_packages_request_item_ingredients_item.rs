use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PackagesCreatePackagesRequestItemIngredientsItem {
    #[serde(rename = "package", skip_serializing_if = "Option::is_none")]
    pub package: Option<String>,
    #[serde(rename = "quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<i64>,
    #[serde(rename = "unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure: Option<String>,
}
