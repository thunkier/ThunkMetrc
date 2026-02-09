use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AdditivesTemplatesUpdateAdditivesTemplatesRequestItem {
    #[serde(rename = "active_ingredients", skip_serializing_if = "Option::is_none")]
    pub active_ingredients: Option<Vec<AdditivesTemplatesUpdateAdditivesTemplatesRequestItemActiveIngredientsItem>>,
    #[serde(rename = "additive_type", skip_serializing_if = "Option::is_none")]
    pub additive_type: Option<String>,
    #[serde(rename = "application_device", skip_serializing_if = "Option::is_none")]
    pub application_device: Option<String>,
    #[serde(rename = "epa_registration_number", skip_serializing_if = "Option::is_none")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "product_supplier", skip_serializing_if = "Option::is_none")]
    pub product_supplier: Option<String>,
    #[serde(rename = "product_trade_name", skip_serializing_if = "Option::is_none")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "restrictive_entry_interval_quantity_description", skip_serializing_if = "Option::is_none")]
    pub restrictive_entry_interval_quantity_description: Option<String>,
    #[serde(rename = "restrictive_entry_interval_time_description", skip_serializing_if = "Option::is_none")]
    pub restrictive_entry_interval_time_description: Option<String>,
}
