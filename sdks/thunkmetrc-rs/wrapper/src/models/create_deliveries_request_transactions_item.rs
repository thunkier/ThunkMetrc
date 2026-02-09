use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateDeliveriesRequestTransactionsItem {
    #[serde(rename = "CityTax", skip_serializing_if = "Option::is_none")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax", skip_serializing_if = "Option::is_none")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount", skip_serializing_if = "Option::is_none")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax", skip_serializing_if = "Option::is_none")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber", skip_serializing_if = "Option::is_none")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax", skip_serializing_if = "Option::is_none")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "Price", skip_serializing_if = "Option::is_none")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes", skip_serializing_if = "Option::is_none")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax", skip_serializing_if = "Option::is_none")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal", skip_serializing_if = "Option::is_none")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount", skip_serializing_if = "Option::is_none")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight", skip_serializing_if = "Option::is_none")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_weight_unit_of_measure: Option<String>,
}
