use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct HarvestsUpdateRestoreHarvestedPlantsRequestItem {
    #[serde(rename = "harvest_id", skip_serializing_if = "Option::is_none")]
    pub harvest_id: Option<i64>,
    #[serde(rename = "plant_tags", skip_serializing_if = "Option::is_none")]
    pub plant_tags: Option<Vec<String>>,
}
