use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdatePatientsRequest {
    #[serde(rename = "ActualDate", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "ConcentrateOuncesAllowed", skip_serializing_if = "Option::is_none")]
    pub concentrate_ounces_allowed: Option<i64>,
    #[serde(rename = "FlowerOuncesAllowed", skip_serializing_if = "Option::is_none")]
    pub flower_ounces_allowed: Option<i64>,
    #[serde(rename = "HasSalesLimitExemption", skip_serializing_if = "Option::is_none")]
    pub has_sales_limit_exemption: Option<bool>,
    #[serde(rename = "InfusedOuncesAllowed", skip_serializing_if = "Option::is_none")]
    pub infused_ounces_allowed: Option<i64>,
    #[serde(rename = "LicenseEffectiveEndDate", skip_serializing_if = "Option::is_none")]
    pub license_effective_end_date: Option<String>,
    #[serde(rename = "LicenseEffectiveStartDate", skip_serializing_if = "Option::is_none")]
    pub license_effective_start_date: Option<String>,
    #[serde(rename = "LicenseNumber", skip_serializing_if = "Option::is_none")]
    pub license_number: Option<String>,
    #[serde(rename = "MaxConcentrateThcPercentAllowed", skip_serializing_if = "Option::is_none")]
    pub max_concentrate_thc_percent_allowed: Option<i64>,
    #[serde(rename = "MaxFlowerThcPercentAllowed", skip_serializing_if = "Option::is_none")]
    pub max_flower_thc_percent_allowed: Option<i64>,
    #[serde(rename = "NewLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub new_license_number: Option<String>,
    #[serde(rename = "RecommendedPlants", skip_serializing_if = "Option::is_none")]
    pub recommended_plants: Option<i64>,
    #[serde(rename = "RecommendedSmokableQuantity", skip_serializing_if = "Option::is_none")]
    pub recommended_smokable_quantity: Option<i64>,
    #[serde(rename = "ThcOuncesAllowed", skip_serializing_if = "Option::is_none")]
    pub thc_ounces_allowed: Option<i64>,
}
