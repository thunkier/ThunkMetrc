use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UnitOfMeasure {
    #[serde(rename = "Abbreviation", skip_serializing_if = "Option::is_none")]
    pub abbreviation: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "QuantityType", skip_serializing_if = "Option::is_none")]
    pub quantity_type: Option<String>,
}
