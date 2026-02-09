use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct LabTestsBatchItemCategoriesItem {
    #[serde(rename = "ProductCategory", skip_serializing_if = "Option::is_none")]
    pub product_category: Option<LabTestsBatchItemCategoriesItemProductCategory>,
    #[serde(rename = "ProductCategoryId", skip_serializing_if = "Option::is_none")]
    pub product_category_id: Option<String>,
    #[serde(rename = "RequireLabTestBatch", skip_serializing_if = "Option::is_none")]
    pub require_lab_test_batch: Option<bool>,
}
