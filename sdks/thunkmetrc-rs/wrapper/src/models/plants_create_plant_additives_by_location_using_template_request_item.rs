use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsCreatePlantAdditivesByLocationUsingTemplateRequestItem {
    #[serde(rename = "actual_date", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "additives_template_name", skip_serializing_if = "Option::is_none")]
    pub additives_template_name: Option<String>,
    #[serde(rename = "location_name", skip_serializing_if = "Option::is_none")]
    pub location_name: Option<String>,
    #[serde(rename = "rate", skip_serializing_if = "Option::is_none")]
    pub rate: Option<String>,
    #[serde(rename = "sublocation_name", skip_serializing_if = "Option::is_none")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "total_amount_applied", skip_serializing_if = "Option::is_none")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "total_amount_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub total_amount_unit_of_measure: Option<String>,
    #[serde(rename = "volume", skip_serializing_if = "Option::is_none")]
    pub volume: Option<String>,
}
