use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateDeliveriesCompleteRequestItem {
    #[serde(rename = "AcceptedPackages", skip_serializing_if = "Option::is_none")]
    pub accepted_packages: Option<Vec<serde_json::Value>>,
    #[serde(rename = "ActualArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub actual_arrival_date_time: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "PaymentType", skip_serializing_if = "Option::is_none")]
    pub payment_type: Option<String>,
    #[serde(rename = "ReturnedPackages", skip_serializing_if = "Option::is_none")]
    pub returned_packages: Option<Vec<serde_json::Value>>,
}
