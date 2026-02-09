use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PatientsStatus {
    #[serde(rename = "Active", skip_serializing_if = "Option::is_none")]
    pub active: Option<bool>,
    #[serde(rename = "ConcentrateOuncesAllowed", skip_serializing_if = "Option::is_none")]
    pub concentrate_ounces_allowed: Option<i64>,
    #[serde(rename = "ConcentrateOuncesAvailable", skip_serializing_if = "Option::is_none")]
    pub concentrate_ounces_available: Option<i64>,
    #[serde(rename = "ConcentrateOuncesPurchased", skip_serializing_if = "Option::is_none")]
    pub concentrate_ounces_purchased: Option<i64>,
    #[serde(rename = "FlowerOuncesAllowed", skip_serializing_if = "Option::is_none")]
    pub flower_ounces_allowed: Option<i64>,
    #[serde(rename = "FlowerOuncesAvailable", skip_serializing_if = "Option::is_none")]
    pub flower_ounces_available: Option<i64>,
    #[serde(rename = "FlowerOuncesPurchased", skip_serializing_if = "Option::is_none")]
    pub flower_ounces_purchased: Option<i64>,
    #[serde(rename = "IdentificationMethod", skip_serializing_if = "Option::is_none")]
    pub identification_method: Option<String>,
    #[serde(rename = "InfusedOuncesAllowed", skip_serializing_if = "Option::is_none")]
    pub infused_ounces_allowed: Option<i64>,
    #[serde(rename = "InfusedOuncesAvailable", skip_serializing_if = "Option::is_none")]
    pub infused_ounces_available: Option<i64>,
    #[serde(rename = "InfusedOuncesPurchased", skip_serializing_if = "Option::is_none")]
    pub infused_ounces_purchased: Option<i64>,
    #[serde(rename = "PatientLicenseExpirationDate", skip_serializing_if = "Option::is_none")]
    pub patient_license_expiration_date: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PatientRegistrationStartDate", skip_serializing_if = "Option::is_none")]
    pub patient_registration_start_date: Option<String>,
    #[serde(rename = "PurchaseAmountDays", skip_serializing_if = "Option::is_none")]
    pub purchase_amount_days: Option<i64>,
    #[serde(rename = "RegistrationNumber", skip_serializing_if = "Option::is_none")]
    pub registration_number: Option<String>,
    #[serde(rename = "ThcOuncesAllowed", skip_serializing_if = "Option::is_none")]
    pub thc_ounces_allowed: Option<i64>,
    #[serde(rename = "ThcOuncesAvailable", skip_serializing_if = "Option::is_none")]
    pub thc_ounces_available: Option<i64>,
    #[serde(rename = "ThcOuncesPurchased", skip_serializing_if = "Option::is_none")]
    pub thc_ounces_purchased: Option<i64>,
}
