use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PackagesCreatePackagePlantingsRequestItem {
    #[serde(rename = "location_name", skip_serializing_if = "Option::is_none")]
    pub location_name: Option<String>,
    #[serde(rename = "package_adjustment_amount", skip_serializing_if = "Option::is_none")]
    pub package_adjustment_amount: Option<i64>,
    #[serde(rename = "package_adjustment_unit_of_measure_name", skip_serializing_if = "Option::is_none")]
    pub package_adjustment_unit_of_measure_name: Option<String>,
    #[serde(rename = "package_label", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "patient_license_number", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "plant_batch_name", skip_serializing_if = "Option::is_none")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "plant_batch_type", skip_serializing_if = "Option::is_none")]
    pub plant_batch_type: Option<String>,
    #[serde(rename = "plant_count", skip_serializing_if = "Option::is_none")]
    pub plant_count: Option<i64>,
    #[serde(rename = "planted_date", skip_serializing_if = "Option::is_none")]
    pub planted_date: Option<String>,
    #[serde(rename = "strain_name", skip_serializing_if = "Option::is_none")]
    pub strain_name: Option<String>,
    #[serde(rename = "sublocation_name", skip_serializing_if = "Option::is_none")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "unpackaged_date", skip_serializing_if = "Option::is_none")]
    pub unpackaged_date: Option<String>,
}
