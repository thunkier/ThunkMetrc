use serde::{Deserialize, Serialize};
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PackageSummary {
    #[serde(rename = "ArchivedDate", skip_serializing_if = "Option::is_none")]
    pub archived_date: Option<String>,
    #[serde(rename = "ContainsDecontaminatedProduct", skip_serializing_if = "Option::is_none")]
    pub contains_decontaminated_product: Option<bool>,
    #[serde(rename = "ContainsRemediatedProduct", skip_serializing_if = "Option::is_none")]
    pub contains_remediated_product: Option<bool>,
    #[serde(rename = "DecontaminationDate", skip_serializing_if = "Option::is_none")]
    pub decontamination_date: Option<String>,
    #[serde(rename = "ExpirationDate", skip_serializing_if = "Option::is_none")]
    pub expiration_date: Option<String>,
    #[serde(rename = "ExternalId", skip_serializing_if = "Option::is_none")]
    pub external_id: Option<i64>,
    #[serde(rename = "FinishedDate", skip_serializing_if = "Option::is_none")]
    pub finished_date: Option<String>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "InitialLabTestingState", skip_serializing_if = "Option::is_none")]
    pub initial_lab_testing_state: Option<String>,
    #[serde(rename = "IsDonation", skip_serializing_if = "Option::is_none")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsDonationPersistent", skip_serializing_if = "Option::is_none")]
    pub is_donation_persistent: Option<bool>,
    #[serde(rename = "IsFinished", skip_serializing_if = "Option::is_none")]
    pub is_finished: Option<bool>,
    #[serde(rename = "IsOnHold", skip_serializing_if = "Option::is_none")]
    pub is_on_hold: Option<bool>,
    #[serde(rename = "IsOnHoldCombined", skip_serializing_if = "Option::is_none")]
    pub is_on_hold_combined: Option<bool>,
    #[serde(rename = "IsOnInvestigation", skip_serializing_if = "Option::is_none")]
    pub is_on_investigation: Option<bool>,
    #[serde(rename = "IsOnInvestigationHold", skip_serializing_if = "Option::is_none")]
    pub is_on_investigation_hold: Option<bool>,
    #[serde(rename = "IsOnInvestigationRecall", skip_serializing_if = "Option::is_none")]
    pub is_on_investigation_recall: Option<bool>,
    #[serde(rename = "IsOnRecall", skip_serializing_if = "Option::is_none")]
    pub is_on_recall: Option<bool>,
    #[serde(rename = "IsOnRecallCombined", skip_serializing_if = "Option::is_none")]
    pub is_on_recall_combined: Option<bool>,
    #[serde(rename = "IsOnRetailerDelivery", skip_serializing_if = "Option::is_none")]
    pub is_on_retailer_delivery: Option<bool>,
    #[serde(rename = "IsProcessValidationTestingSample", skip_serializing_if = "Option::is_none")]
    pub is_process_validation_testing_sample: Option<bool>,
    #[serde(rename = "IsProductionBatch", skip_serializing_if = "Option::is_none")]
    pub is_production_batch: Option<bool>,
    #[serde(rename = "IsTestingSample", skip_serializing_if = "Option::is_none")]
    pub is_testing_sample: Option<bool>,
    #[serde(rename = "IsTradeSample", skip_serializing_if = "Option::is_none")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "IsTradeSamplePersistent", skip_serializing_if = "Option::is_none")]
    pub is_trade_sample_persistent: Option<bool>,
    #[serde(rename = "Item", skip_serializing_if = "Option::is_none")]
    pub item: Option<PackagesPackageSummaryItem>,
    #[serde(rename = "ItemFromFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub item_from_facility_license_number: Option<String>,
    #[serde(rename = "ItemFromFacilityName", skip_serializing_if = "Option::is_none")]
    pub item_from_facility_name: Option<String>,
    #[serde(rename = "LabTestResultExpirationDateTime", skip_serializing_if = "Option::is_none")]
    pub lab_test_result_expiration_date_time: Option<String>,
    #[serde(rename = "LabTestStage", skip_serializing_if = "Option::is_none")]
    pub lab_test_stage: Option<String>,
    #[serde(rename = "LabTestStageId", skip_serializing_if = "Option::is_none")]
    pub lab_test_stage_id: Option<i64>,
    #[serde(rename = "LabTestingPerformedDate", skip_serializing_if = "Option::is_none")]
    pub lab_testing_performed_date: Option<String>,
    #[serde(rename = "LabTestingRecordedDate", skip_serializing_if = "Option::is_none")]
    pub lab_testing_recorded_date: Option<String>,
    #[serde(rename = "LabTestingState", skip_serializing_if = "Option::is_none")]
    pub lab_testing_state: Option<String>,
    #[serde(rename = "LabTestingStateDate", skip_serializing_if = "Option::is_none")]
    pub lab_testing_state_date: Option<String>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "LabelsLastGeneratedDateTime", skip_serializing_if = "Option::is_none")]
    pub labels_last_generated_date_time: Option<String>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "LocationId", skip_serializing_if = "Option::is_none")]
    pub location_id: Option<i64>,
    #[serde(rename = "LocationName", skip_serializing_if = "Option::is_none")]
    pub location_name: Option<String>,
    #[serde(rename = "LocationTypeName", skip_serializing_if = "Option::is_none")]
    pub location_type_name: Option<String>,
    #[serde(rename = "Note", skip_serializing_if = "Option::is_none")]
    pub note: Option<String>,
    #[serde(rename = "OriginalPackageQuantity", skip_serializing_if = "Option::is_none")]
    pub original_package_quantity: Option<f64>,
    #[serde(rename = "PackageForProductDestruction", skip_serializing_if = "Option::is_none")]
    pub package_for_product_destruction: Option<String>,
    #[serde(rename = "PackageType", skip_serializing_if = "Option::is_none")]
    pub package_type: Option<String>,
    #[serde(rename = "PackagedDate", skip_serializing_if = "Option::is_none")]
    pub packaged_date: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProductLabel", skip_serializing_if = "Option::is_none")]
    pub product_label: Option<PackagesPackageSummaryProductLabel>,
    #[serde(rename = "ProductRequiresDecontamination", skip_serializing_if = "Option::is_none")]
    pub product_requires_decontamination: Option<bool>,
    #[serde(rename = "ProductRequiresRemediation", skip_serializing_if = "Option::is_none")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "ProductionBatchNumber", skip_serializing_if = "Option::is_none")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "Quantity", skip_serializing_if = "Option::is_none")]
    pub quantity: Option<f64>,
    #[serde(rename = "ReceivedDateTime", skip_serializing_if = "Option::is_none")]
    pub received_date_time: Option<String>,
    #[serde(rename = "ReceivedFromFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub received_from_facility_license_number: Option<String>,
    #[serde(rename = "ReceivedFromFacilityName", skip_serializing_if = "Option::is_none")]
    pub received_from_facility_name: Option<String>,
    #[serde(rename = "ReceivedFromManifestNumber", skip_serializing_if = "Option::is_none")]
    pub received_from_manifest_number: Option<String>,
    #[serde(rename = "RemediationDate", skip_serializing_if = "Option::is_none")]
    pub remediation_date: Option<String>,
    #[serde(rename = "SellByDate", skip_serializing_if = "Option::is_none")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "SourceHarvestCount", skip_serializing_if = "Option::is_none")]
    pub source_harvest_count: Option<i64>,
    #[serde(rename = "SourceHarvestNames", skip_serializing_if = "Option::is_none")]
    pub source_harvest_names: Option<String>,
    #[serde(rename = "SourcePackageCount", skip_serializing_if = "Option::is_none")]
    pub source_package_count: Option<i64>,
    #[serde(rename = "SourcePackageIsDonation", skip_serializing_if = "Option::is_none")]
    pub source_package_is_donation: Option<bool>,
    #[serde(rename = "SourcePackageIsTradeSample", skip_serializing_if = "Option::is_none")]
    pub source_package_is_trade_sample: Option<bool>,
    #[serde(rename = "SourcePackageLabels", skip_serializing_if = "Option::is_none")]
    pub source_package_labels: Option<String>,
    #[serde(rename = "SourceProcessingJobCount", skip_serializing_if = "Option::is_none")]
    pub source_processing_job_count: Option<i64>,
    #[serde(rename = "SourceProductionBatchNumbers", skip_serializing_if = "Option::is_none")]
    pub source_production_batch_numbers: Option<String>,
    #[serde(rename = "SublocationId", skip_serializing_if = "Option::is_none")]
    pub sublocation_id: Option<i64>,
    #[serde(rename = "SublocationName", skip_serializing_if = "Option::is_none")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "UnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure_abbreviation: Option<String>,
    #[serde(rename = "UnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "UseByDate", skip_serializing_if = "Option::is_none")]
    pub use_by_date: Option<String>,
}
