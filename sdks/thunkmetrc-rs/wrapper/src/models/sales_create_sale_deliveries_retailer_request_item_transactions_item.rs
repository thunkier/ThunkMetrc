use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesCreateSaleDeliveriesRetailerRequestItemTransactionsItem {
    #[serde(rename = "city_tax", skip_serializing_if = "Option::is_none")]
    pub city_tax: Option<String>,
    #[serde(rename = "county_tax", skip_serializing_if = "Option::is_none")]
    pub county_tax: Option<String>,
    #[serde(rename = "discount_amount", skip_serializing_if = "Option::is_none")]
    pub discount_amount: Option<String>,
    #[serde(rename = "excise_tax", skip_serializing_if = "Option::is_none")]
    pub excise_tax: Option<String>,
    #[serde(rename = "invoice_number", skip_serializing_if = "Option::is_none")]
    pub invoice_number: Option<String>,
    #[serde(rename = "municipal_tax", skip_serializing_if = "Option::is_none")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "package_label", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "price", skip_serializing_if = "Option::is_none")]
    pub price: Option<String>,
    #[serde(rename = "qr_codes", skip_serializing_if = "Option::is_none")]
    pub qr_codes: Option<String>,
    #[serde(rename = "quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<i64>,
    #[serde(rename = "sales_tax", skip_serializing_if = "Option::is_none")]
    pub sales_tax: Option<String>,
    #[serde(rename = "sub_total", skip_serializing_if = "Option::is_none")]
    pub sub_total: Option<String>,
    #[serde(rename = "total_amount", skip_serializing_if = "Option::is_none")]
    pub total_amount: Option<f64>,
    #[serde(rename = "unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "unit_thc_content", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "unit_thc_content_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "unit_thc_percent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "unit_weight", skip_serializing_if = "Option::is_none")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "unit_weight_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_weight_unit_of_measure: Option<String>,
}
