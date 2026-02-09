use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesCreateSaleDeliveriesRetailerRequestItem {
    #[serde(rename = "consumer_id", skip_serializing_if = "Option::is_none")]
    pub consumer_id: Option<i64>,
    #[serde(rename = "estimated_arrival_date_time", skip_serializing_if = "Option::is_none")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "estimated_departure_date_time", skip_serializing_if = "Option::is_none")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "patient_license_number", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "phone_number_for_questions", skip_serializing_if = "Option::is_none")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "planned_route", skip_serializing_if = "Option::is_none")]
    pub planned_route: Option<String>,
    #[serde(rename = "recipient_address_city", skip_serializing_if = "Option::is_none")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "recipient_address_county", skip_serializing_if = "Option::is_none")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "recipient_address_postal_code", skip_serializing_if = "Option::is_none")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "recipient_address_state", skip_serializing_if = "Option::is_none")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "recipient_address_street1", skip_serializing_if = "Option::is_none")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "recipient_address_street2", skip_serializing_if = "Option::is_none")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "recipient_name", skip_serializing_if = "Option::is_none")]
    pub recipient_name: Option<String>,
    #[serde(rename = "recipient_zone_id", skip_serializing_if = "Option::is_none")]
    pub recipient_zone_id: Option<i64>,
    #[serde(rename = "retailer_delivery_id", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_id: Option<i64>,
    #[serde(rename = "sales_customer_type", skip_serializing_if = "Option::is_none")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "sales_date_time", skip_serializing_if = "Option::is_none")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "transactions", skip_serializing_if = "Option::is_none")]
    pub transactions: Option<Vec<SalesCreateSaleDeliveriesRetailerRequestItemTransactionsItem>>,
}
