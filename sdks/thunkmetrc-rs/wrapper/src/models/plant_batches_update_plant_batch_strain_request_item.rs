use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantBatchesUpdatePlantBatchStrainRequestItem {
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "strain_id", skip_serializing_if = "Option::is_none")]
    pub strain_id: Option<i64>,
    #[serde(rename = "strain_name", skip_serializing_if = "Option::is_none")]
    pub strain_name: Option<String>,
}
