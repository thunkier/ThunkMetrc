use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct StartProcessingJobRequestItemPackagesItem {
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "Quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<f64>,
    #[serde(rename = "UnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure: Option<String>,
}
