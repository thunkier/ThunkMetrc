use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateHarvestsLocationRequestItem {
    #[serde(rename = "ActualDate", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "DryingLocation", skip_serializing_if = "Option::is_none")]
    pub drying_location: Option<String>,
    #[serde(rename = "DryingSublocation", skip_serializing_if = "Option::is_none")]
    pub drying_sublocation: Option<String>,
    #[serde(rename = "HarvestName", skip_serializing_if = "Option::is_none")]
    pub harvest_name: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
}
