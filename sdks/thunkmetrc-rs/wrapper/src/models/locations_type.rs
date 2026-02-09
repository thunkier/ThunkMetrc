use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct LocationsType {
    #[serde(rename = "ForHarvests", skip_serializing_if = "Option::is_none")]
    pub for_harvests: Option<bool>,
    #[serde(rename = "ForPackages", skip_serializing_if = "Option::is_none")]
    pub for_packages: Option<bool>,
    #[serde(rename = "ForPlantBatches", skip_serializing_if = "Option::is_none")]
    pub for_plant_batches: Option<bool>,
    #[serde(rename = "ForPlants", skip_serializing_if = "Option::is_none")]
    pub for_plants: Option<bool>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
}
