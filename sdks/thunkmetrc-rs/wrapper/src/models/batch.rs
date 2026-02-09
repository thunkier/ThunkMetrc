use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Batch {
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "ItemCategories", skip_serializing_if = "Option::is_none")]
    pub item_categories: Option<Vec<LabTestsBatchItemCategoriesItem>>,
    #[serde(rename = "ItemCategoryCount", skip_serializing_if = "Option::is_none")]
    pub item_category_count: Option<i64>,
    #[serde(rename = "LabTestTypeCount", skip_serializing_if = "Option::is_none")]
    pub lab_test_type_count: Option<i64>,
    #[serde(rename = "LabTestTypes", skip_serializing_if = "Option::is_none")]
    pub lab_test_types: Option<Vec<LabTestsBatchLabTestTypesItem>>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "RequiresAllFromLabTestBatch", skip_serializing_if = "Option::is_none")]
    pub requires_all_from_lab_test_batch: Option<bool>,
}
