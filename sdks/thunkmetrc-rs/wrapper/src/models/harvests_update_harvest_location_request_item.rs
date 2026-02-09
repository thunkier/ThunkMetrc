use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct HarvestsUpdateHarvestLocationRequestItem {
    #[serde(rename = "actual_date", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "drying_location", skip_serializing_if = "Option::is_none")]
    pub drying_location: Option<String>,
    #[serde(rename = "drying_sublocation", skip_serializing_if = "Option::is_none")]
    pub drying_sublocation: Option<String>,
    #[serde(rename = "harvest_name", skip_serializing_if = "Option::is_none")]
    pub harvest_name: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
}
