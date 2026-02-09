use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdatePlantBatchesStrainRequest {
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "StrainId", skip_serializing_if = "Option::is_none")]
    pub strain_id: Option<i64>,
    #[serde(rename = "StrainName", skip_serializing_if = "Option::is_none")]
    pub strain_name: Option<String>,
}
