use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreatePackagesPackagesRequestIngredientsItem {
    #[serde(rename = "Package", skip_serializing_if = "Option::is_none")]
    pub package: Option<String>,
    #[serde(rename = "Quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<i64>,
    #[serde(rename = "UnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure: Option<String>,
}
