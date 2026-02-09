use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesUpdateSaleReceiptsRequestItem {
    #[serde(rename = "caregiver_license_number", skip_serializing_if = "Option::is_none")]
    pub caregiver_license_number: Option<String>,
    #[serde(rename = "external_receipt_number", skip_serializing_if = "Option::is_none")]
    pub external_receipt_number: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "identification_method", skip_serializing_if = "Option::is_none")]
    pub identification_method: Option<String>,
    #[serde(rename = "patient_license_number", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "patient_registration_location_id", skip_serializing_if = "Option::is_none")]
    pub patient_registration_location_id: Option<i64>,
    #[serde(rename = "sales_customer_type", skip_serializing_if = "Option::is_none")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "sales_date_time", skip_serializing_if = "Option::is_none")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "transactions", skip_serializing_if = "Option::is_none")]
    pub transactions: Option<Vec<SalesUpdateSaleReceiptsRequestItemTransactionsItem>>,
}
