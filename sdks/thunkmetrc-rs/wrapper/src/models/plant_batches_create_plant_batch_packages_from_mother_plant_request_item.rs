use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantBatchesCreatePlantBatchPackagesFromMotherPlantRequestItem {
    #[serde(rename = "actual_date", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "count", skip_serializing_if = "Option::is_none")]
    pub count: Option<i64>,
    #[serde(rename = "expiration_date", skip_serializing_if = "Option::is_none")]
    pub expiration_date: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "is_donation", skip_serializing_if = "Option::is_none")]
    pub is_donation: Option<bool>,
    #[serde(rename = "is_trade_sample", skip_serializing_if = "Option::is_none")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "item", skip_serializing_if = "Option::is_none")]
    pub item: Option<String>,
    #[serde(rename = "location", skip_serializing_if = "Option::is_none")]
    pub location: Option<String>,
    #[serde(rename = "note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "patient_license_number", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "plant_batch", skip_serializing_if = "Option::is_none")]
    pub plant_batch: Option<String>,
    #[serde(rename = "sell_by_date", skip_serializing_if = "Option::is_none")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "sublocation", skip_serializing_if = "Option::is_none")]
    pub sublocation: Option<String>,
    #[serde(rename = "tag", skip_serializing_if = "Option::is_none")]
    pub tag: Option<String>,
    #[serde(rename = "use_by_date", skip_serializing_if = "Option::is_none")]
    pub use_by_date: Option<String>,
}
