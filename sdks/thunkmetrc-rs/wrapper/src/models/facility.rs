use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Facility {
    #[serde(rename = "Alias", skip_serializing_if = "Option::is_none")]
    pub alias: Option<String>,
    #[serde(rename = "AllowTransferOfOnHoldPackages", skip_serializing_if = "Option::is_none")]
    pub allow_transfer_of_on_hold_packages: Option<bool>,
    #[serde(rename = "AllowTransferOfOnRecallPackages", skip_serializing_if = "Option::is_none")]
    pub allow_transfer_of_on_recall_packages: Option<bool>,
    #[serde(rename = "CredentialedDate", skip_serializing_if = "Option::is_none")]
    pub credentialed_date: Option<String>,
    #[serde(rename = "DisplayName", skip_serializing_if = "Option::is_none")]
    pub display_name: Option<String>,
    #[serde(rename = "Email", skip_serializing_if = "Option::is_none")]
    pub email: Option<String>,
    #[serde(rename = "FacilityId", skip_serializing_if = "Option::is_none")]
    pub facility_id: Option<i64>,
    #[serde(rename = "FacilityType", skip_serializing_if = "Option::is_none")]
    pub facility_type: Option<FacilityFacilityType>,
    #[serde(rename = "HireDate", skip_serializing_if = "Option::is_none")]
    pub hire_date: Option<String>,
    #[serde(rename = "IsFinancialContact", skip_serializing_if = "Option::is_none")]
    pub is_financial_contact: Option<bool>,
    #[serde(rename = "IsManager", skip_serializing_if = "Option::is_none")]
    pub is_manager: Option<bool>,
    #[serde(rename = "IsOwner", skip_serializing_if = "Option::is_none")]
    pub is_owner: Option<bool>,
    #[serde(rename = "License", skip_serializing_if = "Option::is_none")]
    pub license: Option<FacilityLicense>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "Occupations", skip_serializing_if = "Option::is_none")]
    pub occupations: Option<Vec<serde_json::Value>>,
    #[serde(rename = "SupportActivationDate", skip_serializing_if = "Option::is_none")]
    pub support_activation_date: Option<String>,
    #[serde(rename = "SupportExpirationDate", skip_serializing_if = "Option::is_none")]
    pub support_expiration_date: Option<String>,
}
