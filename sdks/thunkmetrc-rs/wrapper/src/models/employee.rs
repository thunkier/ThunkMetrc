use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Employee {
    #[serde(rename = "FullName", skip_serializing_if = "Option::is_none")]
    pub full_name: Option<String>,
    #[serde(rename = "IsIndustryAdmin", skip_serializing_if = "Option::is_none")]
    pub is_industry_admin: Option<bool>,
    #[serde(rename = "IsManager", skip_serializing_if = "Option::is_none")]
    pub is_manager: Option<bool>,
    #[serde(rename = "IsOwner", skip_serializing_if = "Option::is_none")]
    pub is_owner: Option<bool>,
    #[serde(rename = "License", skip_serializing_if = "Option::is_none")]
    pub license: Option<EmployeeLicense>,
}
