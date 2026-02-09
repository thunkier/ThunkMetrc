use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct StrainsUpdateStrainsRequestItem {
    #[serde(rename = "cbd_level", skip_serializing_if = "Option::is_none")]
    pub cbd_level: Option<f64>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "indica_percentage", skip_serializing_if = "Option::is_none")]
    pub indica_percentage: Option<f64>,
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "sativa_percentage", skip_serializing_if = "Option::is_none")]
    pub sativa_percentage: Option<f64>,
    #[serde(rename = "testing_status", skip_serializing_if = "Option::is_none")]
    pub testing_status: Option<String>,
    #[serde(rename = "thc_level", skip_serializing_if = "Option::is_none")]
    pub thc_level: Option<f64>,
}
