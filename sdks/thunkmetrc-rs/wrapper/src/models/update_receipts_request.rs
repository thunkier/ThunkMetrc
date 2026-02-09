use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateReceiptsRequest {
    #[serde(rename = "CaregiverLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub caregiver_license_number: Option<String>,
    #[serde(rename = "ExternalReceiptNumber", skip_serializing_if = "Option::is_none")]
    pub external_receipt_number: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "IdentificationMethod", skip_serializing_if = "Option::is_none")]
    pub identification_method: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PatientRegistrationLocationId", skip_serializing_if = "Option::is_none")]
    pub patient_registration_location_id: Option<i64>,
    #[serde(rename = "SalesCustomerType", skip_serializing_if = "Option::is_none")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime", skip_serializing_if = "Option::is_none")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions", skip_serializing_if = "Option::is_none")]
    pub transactions: Option<Vec<serde_json::Value>>,
}
