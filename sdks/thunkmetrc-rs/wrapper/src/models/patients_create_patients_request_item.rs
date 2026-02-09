use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PatientsCreatePatientsRequestItem {
    #[serde(rename = "actual_date", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "concentrate_ounces_allowed", skip_serializing_if = "Option::is_none")]
    pub concentrate_ounces_allowed: Option<i64>,
    #[serde(rename = "flower_ounces_allowed", skip_serializing_if = "Option::is_none")]
    pub flower_ounces_allowed: Option<i64>,
    #[serde(rename = "has_sales_limit_exemption", skip_serializing_if = "Option::is_none")]
    pub has_sales_limit_exemption: Option<bool>,
    #[serde(rename = "infused_ounces_allowed", skip_serializing_if = "Option::is_none")]
    pub infused_ounces_allowed: Option<i64>,
    #[serde(rename = "license_effective_end_date", skip_serializing_if = "Option::is_none")]
    pub license_effective_end_date: Option<String>,
    #[serde(rename = "license_effective_start_date", skip_serializing_if = "Option::is_none")]
    pub license_effective_start_date: Option<String>,
    #[serde(rename = "license_number", skip_serializing_if = "Option::is_none")]
    pub license_number: Option<String>,
    #[serde(rename = "max_concentrate_thc_percent_allowed", skip_serializing_if = "Option::is_none")]
    pub max_concentrate_thc_percent_allowed: Option<i64>,
    #[serde(rename = "max_flower_thc_percent_allowed", skip_serializing_if = "Option::is_none")]
    pub max_flower_thc_percent_allowed: Option<i64>,
    #[serde(rename = "recommended_plants", skip_serializing_if = "Option::is_none")]
    pub recommended_plants: Option<i64>,
    #[serde(rename = "recommended_smokable_quantity", skip_serializing_if = "Option::is_none")]
    pub recommended_smokable_quantity: Option<i64>,
    #[serde(rename = "thc_ounces_allowed", skip_serializing_if = "Option::is_none")]
    pub thc_ounces_allowed: Option<i64>,
}
