use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateHarvestsPackagesRequestIngredientsItem {
    #[serde(rename = "HarvestId", skip_serializing_if = "Option::is_none")]
    pub harvest_id: Option<i64>,
    #[serde(rename = "HarvestName", skip_serializing_if = "Option::is_none")]
    pub harvest_name: Option<String>,
    #[serde(rename = "UnitOfWeight", skip_serializing_if = "Option::is_none")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "Weight", skip_serializing_if = "Option::is_none")]
    pub weight: Option<f64>,
}
