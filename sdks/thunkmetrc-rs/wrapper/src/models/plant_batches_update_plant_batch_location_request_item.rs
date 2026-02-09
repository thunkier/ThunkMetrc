use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantBatchesUpdatePlantBatchLocationRequestItem {
    #[serde(rename = "location", skip_serializing_if = "Option::is_none")]
    pub location: Option<String>,
    #[serde(rename = "move_date", skip_serializing_if = "Option::is_none")]
    pub move_date: Option<String>,
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "sublocation", skip_serializing_if = "Option::is_none")]
    pub sublocation: Option<String>,
}
