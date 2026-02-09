use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateRestoreHarvestedPlantsRequest {
    #[serde(rename = "HarvestId", skip_serializing_if = "Option::is_none")]
    pub harvest_id: Option<i64>,
    #[serde(rename = "PlantTags", skip_serializing_if = "Option::is_none")]
    pub plant_tags: Option<Vec<serde_json::Value>>,
}
