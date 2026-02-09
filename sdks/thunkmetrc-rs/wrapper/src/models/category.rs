use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Category {
    #[serde(rename = "CanBeDecontaminated", skip_serializing_if = "Option::is_none")]
    pub can_be_decontaminated: Option<bool>,
    #[serde(rename = "CanBeDestroyed", skip_serializing_if = "Option::is_none")]
    pub can_be_destroyed: Option<bool>,
    #[serde(rename = "CanBeRemediated", skip_serializing_if = "Option::is_none")]
    pub can_be_remediated: Option<bool>,
    #[serde(rename = "CanContainSeeds", skip_serializing_if = "Option::is_none")]
    pub can_contain_seeds: Option<bool>,
    #[serde(rename = "LabTestBatchNames", skip_serializing_if = "Option::is_none")]
    pub lab_test_batch_names: Option<Vec<serde_json::Value>>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "ProductCategoryType", skip_serializing_if = "Option::is_none")]
    pub product_category_type: Option<String>,
    #[serde(rename = "QuantityType", skip_serializing_if = "Option::is_none")]
    pub quantity_type: Option<String>,
    #[serde(rename = "RequiresAdministrationMethod", skip_serializing_if = "Option::is_none")]
    pub requires_administration_method: Option<bool>,
    #[serde(rename = "RequiresAllergens", skip_serializing_if = "Option::is_none")]
    pub requires_allergens: Option<bool>,
    #[serde(rename = "RequiresDescription", skip_serializing_if = "Option::is_none")]
    pub requires_description: Option<bool>,
    #[serde(rename = "RequiresItemBrand", skip_serializing_if = "Option::is_none")]
    pub requires_item_brand: Option<bool>,
    #[serde(rename = "RequiresLabelPhotoDescription", skip_serializing_if = "Option::is_none")]
    pub requires_label_photo_description: Option<bool>,
    #[serde(rename = "RequiresLabelPhotos", skip_serializing_if = "Option::is_none")]
    pub requires_label_photos: Option<i64>,
    #[serde(rename = "RequiresNumberOfDoses", skip_serializing_if = "Option::is_none")]
    pub requires_number_of_doses: Option<bool>,
    #[serde(rename = "RequiresPackagingPhotoDescription", skip_serializing_if = "Option::is_none")]
    pub requires_packaging_photo_description: Option<bool>,
    #[serde(rename = "RequiresPackagingPhotos", skip_serializing_if = "Option::is_none")]
    pub requires_packaging_photos: Option<i64>,
    #[serde(rename = "RequiresProductPDFDocuments", skip_serializing_if = "Option::is_none")]
    pub requires_product_pdf_documents: Option<i64>,
    #[serde(rename = "RequiresProductPhotoDescription", skip_serializing_if = "Option::is_none")]
    pub requires_product_photo_description: Option<bool>,
    #[serde(rename = "RequiresProductPhotos", skip_serializing_if = "Option::is_none")]
    pub requires_product_photos: Option<i64>,
    #[serde(rename = "RequiresPublicIngredients", skip_serializing_if = "Option::is_none")]
    pub requires_public_ingredients: Option<bool>,
    #[serde(rename = "RequiresServingSize", skip_serializing_if = "Option::is_none")]
    pub requires_serving_size: Option<bool>,
    #[serde(rename = "RequiresStrain", skip_serializing_if = "Option::is_none")]
    pub requires_strain: Option<bool>,
    #[serde(rename = "RequiresSupplyDurationDays", skip_serializing_if = "Option::is_none")]
    pub requires_supply_duration_days: Option<bool>,
    #[serde(rename = "RequiresUnitCbdAContent", skip_serializing_if = "Option::is_none")]
    pub requires_unit_cbd_a_content: Option<bool>,
    #[serde(rename = "RequiresUnitCbdAContentDose", skip_serializing_if = "Option::is_none")]
    pub requires_unit_cbd_a_content_dose: Option<bool>,
    #[serde(rename = "RequiresUnitCbdAPercent", skip_serializing_if = "Option::is_none")]
    pub requires_unit_cbd_a_percent: Option<bool>,
    #[serde(rename = "RequiresUnitCbdContent", skip_serializing_if = "Option::is_none")]
    pub requires_unit_cbd_content: Option<bool>,
    #[serde(rename = "RequiresUnitCbdContentDose", skip_serializing_if = "Option::is_none")]
    pub requires_unit_cbd_content_dose: Option<bool>,
    #[serde(rename = "RequiresUnitCbdPercent", skip_serializing_if = "Option::is_none")]
    pub requires_unit_cbd_percent: Option<bool>,
    #[serde(rename = "RequiresUnitThcAContent", skip_serializing_if = "Option::is_none")]
    pub requires_unit_thc_a_content: Option<bool>,
    #[serde(rename = "RequiresUnitThcAContentDose", skip_serializing_if = "Option::is_none")]
    pub requires_unit_thc_a_content_dose: Option<bool>,
    #[serde(rename = "RequiresUnitThcAPercent", skip_serializing_if = "Option::is_none")]
    pub requires_unit_thc_a_percent: Option<bool>,
    #[serde(rename = "RequiresUnitThcContent", skip_serializing_if = "Option::is_none")]
    pub requires_unit_thc_content: Option<bool>,
    #[serde(rename = "RequiresUnitThcContentDose", skip_serializing_if = "Option::is_none")]
    pub requires_unit_thc_content_dose: Option<bool>,
    #[serde(rename = "RequiresUnitThcPercent", skip_serializing_if = "Option::is_none")]
    pub requires_unit_thc_percent: Option<bool>,
    #[serde(rename = "RequiresUnitVolume", skip_serializing_if = "Option::is_none")]
    pub requires_unit_volume: Option<bool>,
    #[serde(rename = "RequiresUnitWeight", skip_serializing_if = "Option::is_none")]
    pub requires_unit_weight: Option<bool>,
}
