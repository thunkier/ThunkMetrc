use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct HarvestsCreateHarvestPackagesTestingRequestItemIngredientsItem {
    #[serde(rename = "harvest_id", skip_serializing_if = "Option::is_none")]
    pub harvest_id: Option<i64>,
    #[serde(rename = "harvest_name", skip_serializing_if = "Option::is_none")]
    pub harvest_name: Option<String>,
    #[serde(rename = "unit_of_weight", skip_serializing_if = "Option::is_none")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "weight", skip_serializing_if = "Option::is_none")]
    pub weight: Option<f64>,
}
