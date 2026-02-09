use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Patient {
    #[serde(rename = "HasSalesLimitExemption", skip_serializing_if = "Option::is_none")]
    pub has_sales_limit_exemption: Option<bool>,
    #[serde(rename = "LicenseEffectiveEndDate", skip_serializing_if = "Option::is_none")]
    pub license_effective_end_date: Option<String>,
    #[serde(rename = "LicenseEffectiveStartDate", skip_serializing_if = "Option::is_none")]
    pub license_effective_start_date: Option<String>,
    #[serde(rename = "LicenseNumber", skip_serializing_if = "Option::is_none")]
    pub license_number: Option<String>,
    #[serde(rename = "OtherFacilitiesCount", skip_serializing_if = "Option::is_none")]
    pub other_facilities_count: Option<i64>,
    #[serde(rename = "PatientId", skip_serializing_if = "Option::is_none")]
    pub patient_id: Option<i64>,
    #[serde(rename = "RecommendedPlants", skip_serializing_if = "Option::is_none")]
    pub recommended_plants: Option<i64>,
    #[serde(rename = "RecommendedSmokableQuantity", skip_serializing_if = "Option::is_none")]
    pub recommended_smokable_quantity: Option<f64>,
    #[serde(rename = "RegistrationDate", skip_serializing_if = "Option::is_none")]
    pub registration_date: Option<String>,
}
