use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Adjustment {
    #[serde(rename = "AdjustedByUserName", skip_serializing_if = "Option::is_none")]
    pub adjusted_by_user_name: Option<String>,
    #[serde(rename = "AdjustmentDate", skip_serializing_if = "Option::is_none")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "AdjustmentNote", skip_serializing_if = "Option::is_none")]
    pub adjustment_note: Option<String>,
    #[serde(rename = "AdjustmentQuantity", skip_serializing_if = "Option::is_none")]
    pub adjustment_quantity: Option<f64>,
    #[serde(rename = "AdjustmentReasonName", skip_serializing_if = "Option::is_none")]
    pub adjustment_reason_name: Option<String>,
    #[serde(rename = "AdjustmentUnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub adjustment_unit_of_measure_abbreviation: Option<String>,
    #[serde(rename = "AdjustmentUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub adjustment_unit_of_measure_name: Option<String>,
    #[serde(rename = "ItemCategoryName", skip_serializing_if = "Option::is_none")]
    pub item_category_name: Option<String>,
    #[serde(rename = "ItemName", skip_serializing_if = "Option::is_none")]
    pub item_name: Option<String>,
    #[serde(rename = "PackageId", skip_serializing_if = "Option::is_none")]
    pub package_id: Option<i64>,
    #[serde(rename = "PackageLabTestResultExpirationDateTime", skip_serializing_if = "Option::is_none")]
    pub package_lab_test_result_expiration_date_time: Option<String>,
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "PackagedDate", skip_serializing_if = "Option::is_none")]
    pub packaged_date: Option<String>,
}
