use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalesSaleDeliveryRetailerPackagesItem {
    #[serde(rename = "ArchivedDate", skip_serializing_if = "Option::is_none")]
    pub archived_date: Option<String>,
    #[serde(rename = "CompletedDateTime", skip_serializing_if = "Option::is_none")]
    pub completed_date_time: Option<String>,
    #[serde(rename = "IsOnInvestigationHold", skip_serializing_if = "Option::is_none")]
    pub is_on_investigation_hold: Option<bool>,
    #[serde(rename = "IsOnInvestigationRecall", skip_serializing_if = "Option::is_none")]
    pub is_on_investigation_recall: Option<bool>,
    #[serde(rename = "IsOnRecall", skip_serializing_if = "Option::is_none")]
    pub is_on_recall: Option<bool>,
    #[serde(rename = "IsOnRecallCombined", skip_serializing_if = "Option::is_none")]
    pub is_on_recall_combined: Option<bool>,
    #[serde(rename = "ItemServingSize", skip_serializing_if = "Option::is_none")]
    pub item_serving_size: Option<String>,
    #[serde(rename = "ItemStrainName", skip_serializing_if = "Option::is_none")]
    pub item_strain_name: Option<String>,
    #[serde(rename = "ItemSupplyDurationDays", skip_serializing_if = "Option::is_none")]
    pub item_supply_duration_days: Option<i64>,
    #[serde(rename = "ItemUnitCbdContent", skip_serializing_if = "Option::is_none")]
    pub item_unit_cbd_content: Option<f64>,
    #[serde(rename = "ItemUnitCbdContentDose", skip_serializing_if = "Option::is_none")]
    pub item_unit_cbd_content_dose: Option<f64>,
    #[serde(rename = "ItemUnitCbdContentDoseUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub item_unit_cbd_content_dose_unit_of_measure_name: Option<String>,
    #[serde(rename = "ItemUnitCbdContentUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub item_unit_cbd_content_unit_of_measure_name: Option<String>,
    #[serde(rename = "ItemUnitCbdPercent", skip_serializing_if = "Option::is_none")]
    pub item_unit_cbd_percent: Option<f64>,
    #[serde(rename = "ItemUnitQuantity", skip_serializing_if = "Option::is_none")]
    pub item_unit_quantity: Option<f64>,
    #[serde(rename = "ItemUnitQuantityUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub item_unit_quantity_unit_of_measure_name: Option<String>,
    #[serde(rename = "ItemUnitThcContent", skip_serializing_if = "Option::is_none")]
    pub item_unit_thc_content: Option<f64>,
    #[serde(rename = "ItemUnitThcContentDose", skip_serializing_if = "Option::is_none")]
    pub item_unit_thc_content_dose: Option<f64>,
    #[serde(rename = "ItemUnitThcContentDoseUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub item_unit_thc_content_dose_unit_of_measure_name: Option<String>,
    #[serde(rename = "ItemUnitThcContentUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub item_unit_thc_content_unit_of_measure_name: Option<String>,
    #[serde(rename = "ItemUnitThcPercent", skip_serializing_if = "Option::is_none")]
    pub item_unit_thc_percent: Option<f64>,
    #[serde(rename = "ItemUnitVolume", skip_serializing_if = "Option::is_none")]
    pub item_unit_volume: Option<f64>,
    #[serde(rename = "ItemUnitVolumeUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub item_unit_volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "ItemUnitWeight", skip_serializing_if = "Option::is_none")]
    pub item_unit_weight: Option<f64>,
    #[serde(rename = "ItemUnitWeightUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub item_unit_weight_unit_of_measure_name: Option<String>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "PackageId", skip_serializing_if = "Option::is_none")]
    pub package_id: Option<i64>,
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "ProductCategoryName", skip_serializing_if = "Option::is_none")]
    pub product_category_name: Option<String>,
    #[serde(rename = "ProductName", skip_serializing_if = "Option::is_none")]
    pub product_name: Option<String>,
    #[serde(rename = "Quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<f64>,
    #[serde(rename = "RecordedByUserName", skip_serializing_if = "Option::is_none")]
    pub recorded_by_user_name: Option<String>,
    #[serde(rename = "RecordedDateTime", skip_serializing_if = "Option::is_none")]
    pub recorded_date_time: Option<String>,
    #[serde(rename = "RetailerDeliveryState", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_state: Option<String>,
    #[serde(rename = "TotalPrice", skip_serializing_if = "Option::is_none")]
    pub total_price: Option<f64>,
    #[serde(rename = "UnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure_abbreviation: Option<String>,
    #[serde(rename = "UnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure_name: Option<String>,
}
