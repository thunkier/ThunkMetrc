use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ItemsUpdateItemsRequestItem {
    #[serde(rename = "administration_method", skip_serializing_if = "Option::is_none")]
    pub administration_method: Option<String>,
    #[serde(rename = "allergens", skip_serializing_if = "Option::is_none")]
    pub allergens: Option<String>,
    #[serde(rename = "description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    #[serde(rename = "global_product_name", skip_serializing_if = "Option::is_none")]
    pub global_product_name: Option<String>,
    #[serde(rename = "id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "item_brand", skip_serializing_if = "Option::is_none")]
    pub item_brand: Option<String>,
    #[serde(rename = "item_category", skip_serializing_if = "Option::is_none")]
    pub item_category: Option<String>,
    #[serde(rename = "item_ingredients", skip_serializing_if = "Option::is_none")]
    pub item_ingredients: Option<String>,
    #[serde(rename = "label_image_file_system_ids", skip_serializing_if = "Option::is_none")]
    pub label_image_file_system_ids: Option<String>,
    #[serde(rename = "label_photo_description", skip_serializing_if = "Option::is_none")]
    pub label_photo_description: Option<String>,
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "number_of_doses", skip_serializing_if = "Option::is_none")]
    pub number_of_doses: Option<String>,
    #[serde(rename = "packaging_image_file_system_ids", skip_serializing_if = "Option::is_none")]
    pub packaging_image_file_system_ids: Option<String>,
    #[serde(rename = "packaging_photo_description", skip_serializing_if = "Option::is_none")]
    pub packaging_photo_description: Option<String>,
    #[serde(rename = "processing_job_category_name", skip_serializing_if = "Option::is_none")]
    pub processing_job_category_name: Option<String>,
    #[serde(rename = "processing_job_type_name", skip_serializing_if = "Option::is_none")]
    pub processing_job_type_name: Option<String>,
    #[serde(rename = "product_image_file_system_ids", skip_serializing_if = "Option::is_none")]
    pub product_image_file_system_ids: Option<String>,
    #[serde(rename = "product_pdf_file_system_ids", skip_serializing_if = "Option::is_none")]
    pub product_pdf_file_system_ids: Option<String>,
    #[serde(rename = "product_photo_description", skip_serializing_if = "Option::is_none")]
    pub product_photo_description: Option<String>,
    #[serde(rename = "public_ingredients", skip_serializing_if = "Option::is_none")]
    pub public_ingredients: Option<String>,
    #[serde(rename = "serving_size", skip_serializing_if = "Option::is_none")]
    pub serving_size: Option<String>,
    #[serde(rename = "strain", skip_serializing_if = "Option::is_none")]
    pub strain: Option<String>,
    #[serde(rename = "supply_duration_days", skip_serializing_if = "Option::is_none")]
    pub supply_duration_days: Option<i64>,
    #[serde(rename = "unit_cbd_a_content", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content: Option<f64>,
    #[serde(rename = "unit_cbd_a_content_dose", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content_dose: Option<f64>,
    #[serde(rename = "unit_cbd_a_content_dose_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "unit_cbd_a_content_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "unit_cbd_a_percent", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_percent: Option<f64>,
    #[serde(rename = "unit_cbd_content", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content: Option<f64>,
    #[serde(rename = "unit_cbd_content_dose", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content_dose: Option<f64>,
    #[serde(rename = "unit_cbd_content_dose_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "unit_cbd_content_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content_unit_of_measure: Option<String>,
    #[serde(rename = "unit_cbd_percent", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_percent: Option<f64>,
    #[serde(rename = "unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "unit_thc_a_content", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content: Option<f64>,
    #[serde(rename = "unit_thc_a_content_dose", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content_dose: Option<f64>,
    #[serde(rename = "unit_thc_a_content_dose_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "unit_thc_a_content_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "unit_thc_a_percent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_percent: Option<f64>,
    #[serde(rename = "unit_thc_content", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "unit_thc_content_dose", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_dose: Option<f64>,
    #[serde(rename = "unit_thc_content_dose_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "unit_thc_content_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "unit_thc_percent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "unit_volume", skip_serializing_if = "Option::is_none")]
    pub unit_volume: Option<f64>,
    #[serde(rename = "unit_volume_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_volume_unit_of_measure: Option<String>,
    #[serde(rename = "unit_weight", skip_serializing_if = "Option::is_none")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "unit_weight_unit_of_measure", skip_serializing_if = "Option::is_none")]
    pub unit_weight_unit_of_measure: Option<String>,
}
