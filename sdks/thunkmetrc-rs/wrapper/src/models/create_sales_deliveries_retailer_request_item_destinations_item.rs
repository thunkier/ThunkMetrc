use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateSalesDeliveriesRetailerRequestItemDestinationsItem {
    #[serde(rename = "ConsumerId", skip_serializing_if = "Option::is_none")]
    pub consumer_id: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime", skip_serializing_if = "Option::is_none")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "RecipientAddressCity", skip_serializing_if = "Option::is_none")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty", skip_serializing_if = "Option::is_none")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode", skip_serializing_if = "Option::is_none")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState", skip_serializing_if = "Option::is_none")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1", skip_serializing_if = "Option::is_none")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2", skip_serializing_if = "Option::is_none")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName", skip_serializing_if = "Option::is_none")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId", skip_serializing_if = "Option::is_none")]
    pub recipient_zone_id: Option<String>,
    #[serde(rename = "SalesCustomerType", skip_serializing_if = "Option::is_none")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "Transactions", skip_serializing_if = "Option::is_none")]
    pub transactions: Option<Vec<serde_json::Value>>,
}
