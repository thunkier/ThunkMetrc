use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TransferTemplateDeliveryPackage {
    #[serde(rename = "ContainsRemediatedProduct", skip_serializing_if = "Option::is_none")]
    pub contains_remediated_product: Option<bool>,
    #[serde(rename = "ExpirationDate", skip_serializing_if = "Option::is_none")]
    pub expiration_date: Option<String>,
    #[serde(rename = "ExternalId", skip_serializing_if = "Option::is_none")]
    pub external_id: Option<i64>,
    #[serde(rename = "GrossUnitOfWeightName", skip_serializing_if = "Option::is_none")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "IsDonation", skip_serializing_if = "Option::is_none")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsTestingSample", skip_serializing_if = "Option::is_none")]
    pub is_testing_sample: Option<bool>,
    #[serde(rename = "IsTradeSample", skip_serializing_if = "Option::is_none")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "IsTradeSamplePersistent", skip_serializing_if = "Option::is_none")]
    pub is_trade_sample_persistent: Option<bool>,
    #[serde(rename = "ItemBrandName", skip_serializing_if = "Option::is_none")]
    pub item_brand_name: Option<String>,
    #[serde(rename = "ItemCategoryName", skip_serializing_if = "Option::is_none")]
    pub item_category_name: Option<String>,
    #[serde(rename = "ItemId", skip_serializing_if = "Option::is_none")]
    pub item_id: Option<i64>,
    #[serde(rename = "ItemName", skip_serializing_if = "Option::is_none")]
    pub item_name: Option<String>,
    #[serde(rename = "ItemServingSize", skip_serializing_if = "Option::is_none")]
    pub item_serving_size: Option<String>,
    #[serde(rename = "ItemStrainName", skip_serializing_if = "Option::is_none")]
    pub item_strain_name: Option<String>,
    #[serde(rename = "ItemSupplyDurationDays", skip_serializing_if = "Option::is_none")]
    pub item_supply_duration_days: Option<i64>,
    #[serde(rename = "ItemUnitCbdAContent", skip_serializing_if = "Option::is_none")]
    pub item_unit_cbd_a_content: Option<f64>,
    #[serde(rename = "ItemUnitCbdAContentDose", skip_serializing_if = "Option::is_none")]
    pub item_unit_cbd_a_content_dose: Option<f64>,
    #[serde(rename = "ItemUnitCbdAContentDoseUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub item_unit_cbd_a_content_dose_unit_of_measure_name: Option<String>,
    #[serde(rename = "ItemUnitCbdAContentUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub item_unit_cbd_a_content_unit_of_measure_name: Option<String>,
    #[serde(rename = "ItemUnitCbdAPercent", skip_serializing_if = "Option::is_none")]
    pub item_unit_cbd_a_percent: Option<f64>,
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
    #[serde(rename = "ItemUnitThcAContent", skip_serializing_if = "Option::is_none")]
    pub item_unit_thc_a_content: Option<f64>,
    #[serde(rename = "ItemUnitThcAContentDose", skip_serializing_if = "Option::is_none")]
    pub item_unit_thc_a_content_dose: Option<f64>,
    #[serde(rename = "ItemUnitThcAContentDoseUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub item_unit_thc_a_content_dose_unit_of_measure_name: Option<String>,
    #[serde(rename = "ItemUnitThcAContentUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub item_unit_thc_a_content_unit_of_measure_name: Option<String>,
    #[serde(rename = "ItemUnitThcAPercent", skip_serializing_if = "Option::is_none")]
    pub item_unit_thc_a_percent: Option<f64>,
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
    #[serde(rename = "LabTestingState", skip_serializing_if = "Option::is_none")]
    pub lab_testing_state: Option<String>,
    #[serde(rename = "PackageId", skip_serializing_if = "Option::is_none")]
    pub package_id: Option<i64>,
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "PackageType", skip_serializing_if = "Option::is_none")]
    pub package_type: Option<String>,
    #[serde(rename = "PackagedDate", skip_serializing_if = "Option::is_none")]
    pub packaged_date: Option<String>,
    #[serde(rename = "ProductLabel", skip_serializing_if = "Option::is_none")]
    pub product_label: Option<TransfersTransferTemplateDeliveryPackageProductLabel>,
    #[serde(rename = "ProductRequiresRemediation", skip_serializing_if = "Option::is_none")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "ProductionBatchNumber", skip_serializing_if = "Option::is_none")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "ReceivedQuantity", skip_serializing_if = "Option::is_none")]
    pub received_quantity: Option<f64>,
    #[serde(rename = "ReceivedUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub received_unit_of_measure_name: Option<String>,
    #[serde(rename = "RemediationDate", skip_serializing_if = "Option::is_none")]
    pub remediation_date: Option<String>,
    #[serde(rename = "SellByDate", skip_serializing_if = "Option::is_none")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "ShipmentPackageState", skip_serializing_if = "Option::is_none")]
    pub shipment_package_state: Option<String>,
    #[serde(rename = "ShippedQuantity", skip_serializing_if = "Option::is_none")]
    pub shipped_quantity: Option<f64>,
    #[serde(rename = "ShippedUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub shipped_unit_of_measure_name: Option<String>,
    #[serde(rename = "SourceHarvestNames", skip_serializing_if = "Option::is_none")]
    pub source_harvest_names: Option<String>,
    #[serde(rename = "SourcePackageIsDonation", skip_serializing_if = "Option::is_none")]
    pub source_package_is_donation: Option<bool>,
    #[serde(rename = "SourcePackageIsTradeSample", skip_serializing_if = "Option::is_none")]
    pub source_package_is_trade_sample: Option<bool>,
    #[serde(rename = "SourcePackageLabels", skip_serializing_if = "Option::is_none")]
    pub source_package_labels: Option<String>,
    #[serde(rename = "SourceProductionBatchNumbers", skip_serializing_if = "Option::is_none")]
    pub source_production_batch_numbers: Option<String>,
    #[serde(rename = "UseByDate", skip_serializing_if = "Option::is_none")]
    pub use_by_date: Option<String>,
}
