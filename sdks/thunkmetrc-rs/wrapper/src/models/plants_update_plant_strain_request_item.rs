use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsUpdatePlantStrainRequestItem {
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "strain_id", skip_serializing_if = "Option::is_none")]
    pub strain_id: Option<i64>,
    #[serde(rename = "strain_name", skip_serializing_if = "Option::is_none")]
    pub strain_name: Option<String>,
}
