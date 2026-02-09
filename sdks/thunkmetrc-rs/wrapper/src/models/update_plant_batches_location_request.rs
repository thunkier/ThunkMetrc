use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdatePlantBatchesLocationRequest {
    #[serde(rename = "Location", skip_serializing_if = "Option::is_none")]
    pub location: Option<String>,
    #[serde(rename = "MoveDate", skip_serializing_if = "Option::is_none")]
    pub move_date: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "Sublocation", skip_serializing_if = "Option::is_none")]
    pub sublocation: Option<String>,
}
