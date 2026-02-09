use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateHarvestsWasteRequest {
    #[serde(rename = "ActualDate", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "UnitOfWeight", skip_serializing_if = "Option::is_none")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "WasteType", skip_serializing_if = "Option::is_none")]
    pub waste_type: Option<String>,
    #[serde(rename = "WasteWeight", skip_serializing_if = "Option::is_none")]
    pub waste_weight: Option<f64>,
}
