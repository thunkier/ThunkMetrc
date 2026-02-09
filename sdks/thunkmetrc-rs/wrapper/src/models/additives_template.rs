use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AdditivesTemplate {
    #[serde(rename = "ActiveIngredients", skip_serializing_if = "Option::is_none")]
    pub active_ingredients: Option<Vec<AdditivesTemplateActiveIngredientsItem>>,
    #[serde(rename = "AdditiveType", skip_serializing_if = "Option::is_none")]
    pub additive_type: Option<String>,
    #[serde(rename = "AdditiveTypeName", skip_serializing_if = "Option::is_none")]
    pub additive_type_name: Option<String>,
    #[serde(rename = "ApplicationDevice", skip_serializing_if = "Option::is_none")]
    pub application_device: Option<String>,
    #[serde(rename = "EpaRegistrationNumber", skip_serializing_if = "Option::is_none")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "FacilityId", skip_serializing_if = "Option::is_none")]
    pub facility_id: Option<i64>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "IsActive", skip_serializing_if = "Option::is_none")]
    pub is_active: Option<bool>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "Note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "ProductSupplier", skip_serializing_if = "Option::is_none")]
    pub product_supplier: Option<String>,
    #[serde(rename = "ProductTradeName", skip_serializing_if = "Option::is_none")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "RestrictiveEntryIntervalQuantityDescription", skip_serializing_if = "Option::is_none")]
    pub restrictive_entry_interval_quantity_description: Option<String>,
    #[serde(rename = "RestrictiveEntryIntervalTimeDescription", skip_serializing_if = "Option::is_none")]
    pub restrictive_entry_interval_time_description: Option<String>,
}
