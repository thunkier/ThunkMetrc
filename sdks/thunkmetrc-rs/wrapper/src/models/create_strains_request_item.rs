use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateStrainsRequestItem {
    #[serde(rename = "CbdLevel", skip_serializing_if = "Option::is_none")]
    pub cbd_level: Option<f64>,
    #[serde(rename = "IndicaPercentage", skip_serializing_if = "Option::is_none")]
    pub indica_percentage: Option<i64>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "SativaPercentage", skip_serializing_if = "Option::is_none")]
    pub sativa_percentage: Option<i64>,
    #[serde(rename = "TestingStatus", skip_serializing_if = "Option::is_none")]
    pub testing_status: Option<String>,
    #[serde(rename = "ThcLevel", skip_serializing_if = "Option::is_none")]
    pub thc_level: Option<f64>,
}
