use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantWasteMethodsAll {
    #[serde(rename = "ForPlants", skip_serializing_if = "Option::is_none")]
    pub for_plants: Option<bool>,
    #[serde(rename = "ForProductDestruction", skip_serializing_if = "Option::is_none")]
    pub for_product_destruction: Option<bool>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
}
