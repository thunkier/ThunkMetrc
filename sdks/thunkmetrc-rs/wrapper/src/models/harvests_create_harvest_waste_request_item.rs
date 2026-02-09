use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct HarvestsCreateHarvestWasteRequestItem {
    #[serde(rename = "actual_date", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "unit_of_weight", skip_serializing_if = "Option::is_none")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "waste_type", skip_serializing_if = "Option::is_none")]
    pub waste_type: Option<String>,
    #[serde(rename = "waste_weight", skip_serializing_if = "Option::is_none")]
    pub waste_weight: Option<f64>,
}
