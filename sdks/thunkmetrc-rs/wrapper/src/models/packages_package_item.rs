use serde::{Deserialize, Serialize};
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PackagesPackageItem {
    #[serde(rename = "AdministrationMethod", skip_serializing_if = "Option::is_none")]
    pub administration_method: Option<String>,
    #[serde(rename = "Allergens", skip_serializing_if = "Option::is_none")]
    pub allergens: Option<String>,
    #[serde(rename = "ApprovalStatus", skip_serializing_if = "Option::is_none")]
    pub approval_status: Option<i64>,
    #[serde(rename = "ApprovalStatusDateTime", skip_serializing_if = "Option::is_none")]
    pub approval_status_date_time: Option<String>,
    #[serde(rename = "DefaultLabTestingState", skip_serializing_if = "Option::is_none")]
    pub default_lab_testing_state: Option<i64>,
    #[serde(rename = "Description", skip_serializing_if = "Option::is_none")]
    pub description: Option<String>,
    #[serde(rename = "GlobalProductName", skip_serializing_if = "Option::is_none")]
    pub global_product_name: Option<String>,
    #[serde(rename = "GlobalProductNumber", skip_serializing_if = "Option::is_none")]
    pub global_product_number: Option<String>,
    #[serde(rename = "HasExpirationDate", skip_serializing_if = "Option::is_none")]
    pub has_expiration_date: Option<bool>,
    #[serde(rename = "HasSellByDate", skip_serializing_if = "Option::is_none")]
    pub has_sell_by_date: Option<bool>,
    #[serde(rename = "HasUseByDate", skip_serializing_if = "Option::is_none")]
    pub has_use_by_date: Option<bool>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "IsExpirationDateRequired", skip_serializing_if = "Option::is_none")]
    pub is_expiration_date_required: Option<bool>,
    #[serde(rename = "IsSellByDateRequired", skip_serializing_if = "Option::is_none")]
    pub is_sell_by_date_required: Option<bool>,
    #[serde(rename = "IsUseByDateRequired", skip_serializing_if = "Option::is_none")]
    pub is_use_by_date_required: Option<bool>,
    #[serde(rename = "IsUsed", skip_serializing_if = "Option::is_none")]
    pub is_used: Option<bool>,
    #[serde(rename = "ItemBrandId", skip_serializing_if = "Option::is_none")]
    pub item_brand_id: Option<i64>,
    #[serde(rename = "ItemBrandName", skip_serializing_if = "Option::is_none")]
    pub item_brand_name: Option<String>,
    #[serde(rename = "LabTestBatchNames", skip_serializing_if = "Option::is_none")]
    pub lab_test_batch_names: Option<Vec<serde_json::Value>>,
    #[serde(rename = "LabelImages", skip_serializing_if = "Option::is_none")]
    pub label_images: Option<Vec<serde_json::Value>>,
    #[serde(rename = "LabelPhotoDescription", skip_serializing_if = "Option::is_none")]
    pub label_photo_description: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "NumberOfDoses", skip_serializing_if = "Option::is_none")]
    pub number_of_doses: Option<String>,
    #[serde(rename = "PackagingImages", skip_serializing_if = "Option::is_none")]
    pub packaging_images: Option<Vec<serde_json::Value>>,
    #[serde(rename = "PackagingPhotoDescription", skip_serializing_if = "Option::is_none")]
    pub packaging_photo_description: Option<String>,
    #[serde(rename = "ProcessingJobCategoryName", skip_serializing_if = "Option::is_none")]
    pub processing_job_category_name: Option<String>,
    #[serde(rename = "ProcessingJobTypeName", skip_serializing_if = "Option::is_none")]
    pub processing_job_type_name: Option<String>,
    #[serde(rename = "ProductBrandName", skip_serializing_if = "Option::is_none")]
    pub product_brand_name: Option<String>,
    #[serde(rename = "ProductCategoryName", skip_serializing_if = "Option::is_none")]
    pub product_category_name: Option<String>,
    #[serde(rename = "ProductCategoryType", skip_serializing_if = "Option::is_none")]
    pub product_category_type: Option<i64>,
    #[serde(rename = "ProductImages", skip_serializing_if = "Option::is_none")]
    pub product_images: Option<Vec<serde_json::Value>>,
    #[serde(rename = "ProductPDFDocuments", skip_serializing_if = "Option::is_none")]
    pub product_pdf_documents: Option<Vec<serde_json::Value>>,
    #[serde(rename = "ProductPhotoDescription", skip_serializing_if = "Option::is_none")]
    pub product_photo_description: Option<String>,
    #[serde(rename = "PublicIngredients", skip_serializing_if = "Option::is_none")]
    pub public_ingredients: Option<String>,
    #[serde(rename = "QuantityType", skip_serializing_if = "Option::is_none")]
    pub quantity_type: Option<i64>,
    #[serde(rename = "ServingSize", skip_serializing_if = "Option::is_none")]
    pub serving_size: Option<String>,
    #[serde(rename = "StrainId", skip_serializing_if = "Option::is_none")]
    pub strain_id: Option<i64>,
    #[serde(rename = "StrainName", skip_serializing_if = "Option::is_none")]
    pub strain_name: Option<String>,
    #[serde(rename = "UnitCbdAContent", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content: Option<f64>,
    #[serde(rename = "UnitCbdAContentDose", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdAContentDoseUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content_dose_unit_of_measure_name: Option<String>,
    #[serde(rename = "UnitCbdAContentUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_content_unit_of_measure_name: Option<String>,
    #[serde(rename = "UnitCbdAPercent", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_a_percent: Option<f64>,
    #[serde(rename = "UnitCbdContent", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content: Option<f64>,
    #[serde(rename = "UnitCbdContentDose", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdContentDoseUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content_dose_unit_of_measure_name: Option<String>,
    #[serde(rename = "UnitCbdContentUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_content_unit_of_measure_name: Option<String>,
    #[serde(rename = "UnitCbdPercent", skip_serializing_if = "Option::is_none")]
    pub unit_cbd_percent: Option<f64>,
    #[serde(rename = "UnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "UnitQuantity", skip_serializing_if = "Option::is_none")]
    pub unit_quantity: Option<f64>,
    #[serde(rename = "UnitQuantityUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_quantity_unit_of_measure_name: Option<String>,
    #[serde(rename = "UnitThcAContent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content: Option<f64>,
    #[serde(rename = "UnitThcAContentDose", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content_dose: Option<f64>,
    #[serde(rename = "UnitThcAContentDoseUnitOfMeasureId", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content_dose_unit_of_measure_id: Option<i64>,
    #[serde(rename = "UnitThcAContentDoseUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content_dose_unit_of_measure_name: Option<String>,
    #[serde(rename = "UnitThcAContentUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_content_unit_of_measure_name: Option<String>,
    #[serde(rename = "UnitThcAPercent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_a_percent: Option<f64>,
    #[serde(rename = "UnitThcContent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentDose", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_dose: Option<f64>,
    #[serde(rename = "UnitThcContentDoseUnitOfMeasureId", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_dose_unit_of_measure_id: Option<i64>,
    #[serde(rename = "UnitThcContentDoseUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_dose_unit_of_measure_name: Option<String>,
    #[serde(rename = "UnitThcContentUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_thc_content_unit_of_measure_name: Option<String>,
    #[serde(rename = "UnitThcPercent", skip_serializing_if = "Option::is_none")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitVolume", skip_serializing_if = "Option::is_none")]
    pub unit_volume: Option<f64>,
    #[serde(rename = "UnitVolumeUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "UnitWeight", skip_serializing_if = "Option::is_none")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_weight_unit_of_measure_name: Option<String>,
}
