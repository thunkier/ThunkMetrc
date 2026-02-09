use serde::{Deserialize, Serialize};
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SaleReceipt {
    #[serde(rename = "ArchivedDate", skip_serializing_if = "Option::is_none")]
    pub archived_date: Option<String>,
    #[serde(rename = "CaregiverLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub caregiver_license_number: Option<String>,
    #[serde(rename = "ExternalReceiptNumber", skip_serializing_if = "Option::is_none")]
    pub external_receipt_number: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "IdentificationMethod", skip_serializing_if = "Option::is_none")]
    pub identification_method: Option<String>,
    #[serde(rename = "IsFinal", skip_serializing_if = "Option::is_none")]
    pub is_final: Option<bool>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PatientRegistrationLocationId", skip_serializing_if = "Option::is_none")]
    pub patient_registration_location_id: Option<i64>,
    #[serde(rename = "ReceiptNumber", skip_serializing_if = "Option::is_none")]
    pub receipt_number: Option<String>,
    #[serde(rename = "RecordedByUserName", skip_serializing_if = "Option::is_none")]
    pub recorded_by_user_name: Option<String>,
    #[serde(rename = "RecordedDateTime", skip_serializing_if = "Option::is_none")]
    pub recorded_date_time: Option<String>,
    #[serde(rename = "SalesCustomerType", skip_serializing_if = "Option::is_none")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime", skip_serializing_if = "Option::is_none")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "TotalPackages", skip_serializing_if = "Option::is_none")]
    pub total_packages: Option<i64>,
    #[serde(rename = "TotalPrice", skip_serializing_if = "Option::is_none")]
    pub total_price: Option<i64>,
    #[serde(rename = "Transactions", skip_serializing_if = "Option::is_none")]
    pub transactions: Option<Vec<serde_json::Value>>,
}
