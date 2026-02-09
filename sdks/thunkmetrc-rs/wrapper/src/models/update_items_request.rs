use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateItemsRequest {
    #[serde(rename = "AdministrationMethod", skip_serializing_if = "Option::is_none")]
    pub administration_method: Option<String>,
    #[serde(rename = "Allergens", skip_serializing_if = "Option::is_none")]
    pub allergens: Option<String>,
    #[serde(rename = "Description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    #[serde(rename = "GlobalProductName", skip_serializing_if = "Option::is_none")]
    pub global_product_name: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "ItemBrand", skip_serializing_if = "Option::is_none")]
    pub item_brand: Option<String>,
    #[serde(rename = "ItemCategory", skip_serializing_if = "Option::is_none")]
    pub item_category: Option<String>,
    #[serde(rename = "ItemIngredients", skip_serializing_if = "Option::is_none")]
    pub item_ingredients: Option<String>,
    #[serde(rename = "LabelImageFileSystemIds", skip_serializing_if = "Option::is_none")]
    pub label_image_file_system_ids: Option<String>,
    #[serde(rename = "LabelPhotoDescription", skip_serializing_if = "Option::is_none")]
    pub label_photo_description: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "NumberOfDoses", skip_serializing_if = "Option::is_none")]
    pub number_of_doses: Option<String>,
    #[serde(rename = "PackagingImageFileSystemIds", skip_serializing_if = "Option::is_none")]
    pub packaging_image_file_system_ids: Option<String>,
    #[serde(rename = "PackagingPhotoDescription", skip_serializing_if = "Option::is_none")]
    pub packaging_photo_description: Option<String>,
    #[serde(rename = "ProcessingJobCategoryName", skip_serializing_if = "Option::is_none")]
    pub processing_job_category_name: Option<String>,
    #[serde(rename = "ProcessingJobTypeName", skip_serializing_if = "Option::is_none")]
    pub processing_job_type_name: Option<String>,
    #[serde(rename = "ProductImageFileSystemIds", skip_serializing_if = "Option::is_none")]
    pub product_image_file_system_ids: Option<String>,
    #[serde(rename = "ProductPDFFileSystemIds", skip_serializing_if = "Option::is_none")]
    pub product_pdf_file_system_ids: Option<String>,
    #[serde(rename = "ProductPhotoDescription", skip_serializing_if = "Option::is_none")]
    pub product_photo_description: Option<String>,
    #[serde(rename = "PublicIngredients", skip_serializing_if = "Option::is_none")]
    pub public_ingredients: Option<String>,
    #[serde(rename = "ServingSize", skip_serializing_if = "Option::is_none")]
    pub serving_size: Option<String>,
    #[serde(rename = "Strain", skip_serializing_if = "Option::is_none")]
    pub strain: Option<String>,
    #[serde(rename = "SupplyDurationDays", skip_serializing_if = "Option::is_none")]
    pub supply_duration_days: Option<i64>,
    #[serde(rename = "UnitCbdAContent", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content: Option<f64>,
    #[serde(rename = "UnitCbdAContentDose", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdAContentDoseUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdAContentUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdAPercent", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_percent: Option<f64>,
    #[serde(rename = "UnitCbdContent", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content: Option<f64>,
    #[serde(rename = "UnitCbdContentDose", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdContentDoseUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdContentUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdPercent", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_percent: Option<f64>,
    #[serde(rename = "UnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAContent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content: Option<f64>,
    #[serde(rename = "UnitThcAContentDose", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content_dose: Option<f64>,
    #[serde(rename = "UnitThcAContentDoseUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAContentUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAPercent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_percent: Option<f64>,
    #[serde(rename = "UnitThcContent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentDose", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_dose: Option<f64>,
    #[serde(rename = "UnitThcContentDoseUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContentUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitVolume", skip_serializing_if = "Option::is_none")]
    pub unit_volume: Option<f64>,
    #[serde(rename = "UnitVolumeUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_volume_unit_of_measure: Option<String>,
    #[serde(rename = "UnitWeight", skip_serializing_if = "Option::is_none")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub unit_weight_unit_of_measure: Option<String>,
}
