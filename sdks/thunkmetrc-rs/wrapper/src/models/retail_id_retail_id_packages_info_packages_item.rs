use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct RetailIdRetailIdPackagesInfoPackagesItem {
    #[serde(rename = "EstimatedBalance", skip_serializing_if = "Option::is_none")]
    pub estimated_balance: Option<i64>,
    #[serde(rename = "HasQrs", skip_serializing_if = "Option::is_none")]
    pub has_qrs: Option<bool>,
    #[serde(rename = "IssuanceId", skip_serializing_if = "Option::is_none")]
    pub issuance_id: Option<String>,
    #[serde(rename = "Issuances", skip_serializing_if = "Option::is_none")]
    pub issuances: Option<Vec<RetailIdRetailIdPackagesInfoPackagesItemIssuancesItem>>,
    #[serde(rename = "QrCount", skip_serializing_if = "Option::is_none")]
    pub qr_count: Option<i64>,
    #[serde(rename = "RequiresVerification", skip_serializing_if = "Option::is_none")]
    pub requires_verification: Option<bool>,
    #[serde(rename = "SiblingCount", skip_serializing_if = "Option::is_none")]
    pub sibling_count: Option<i64>,
    #[serde(rename = "Tag", skip_serializing_if = "Option::is_none")]
    pub tag: Option<String>,
}
