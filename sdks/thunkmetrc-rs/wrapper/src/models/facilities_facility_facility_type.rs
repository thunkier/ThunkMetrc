use serde::{Deserialize, Serialize};
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct FacilitiesFacilityFacilityType {
    #[serde(rename = "AdvancedSales", skip_serializing_if = "Option::is_none")]
    pub advanced_sales: Option<bool>,
    #[serde(rename = "CanAccessCatalog", skip_serializing_if = "Option::is_none")]
    pub can_access_catalog: Option<bool>,
    #[serde(rename = "CanAdjustSourcePackagesWithPartials", skip_serializing_if = "Option::is_none")]
    pub can_adjust_source_packages_with_partials: Option<bool>,
    #[serde(rename = "CanAssignLocationsToPackages", skip_serializing_if = "Option::is_none")]
    pub can_assign_locations_to_packages: Option<bool>,
    #[serde(rename = "CanAssignLocationsToPlantBatches", skip_serializing_if = "Option::is_none")]
    pub can_assign_locations_to_plant_batches: Option<bool>,
    #[serde(rename = "CanClonePlantBatches", skip_serializing_if = "Option::is_none")]
    pub can_clone_plant_batches: Option<bool>,
    #[serde(rename = "CanCreateDerivedPackages", skip_serializing_if = "Option::is_none")]
    pub can_create_derived_packages: Option<bool>,
    #[serde(rename = "CanCreateImmaturePlantPackagesFromPlants", skip_serializing_if = "Option::is_none")]
    pub can_create_immature_plant_packages_from_plants: Option<bool>,
    #[serde(rename = "CanCreateOpeningBalancePackages", skip_serializing_if = "Option::is_none")]
    pub can_create_opening_balance_packages: Option<bool>,
    #[serde(rename = "CanCreateOpeningBalancePlantBatches", skip_serializing_if = "Option::is_none")]
    pub can_create_opening_balance_plant_batches: Option<bool>,
    #[serde(rename = "CanCreatePartialPackages", skip_serializing_if = "Option::is_none")]
    pub can_create_partial_packages: Option<bool>,
    #[serde(rename = "CanCreateProcessValidationPackages", skip_serializing_if = "Option::is_none")]
    pub can_create_process_validation_packages: Option<bool>,
    #[serde(rename = "CanCreateTradeSamplePackages", skip_serializing_if = "Option::is_none")]
    pub can_create_trade_sample_packages: Option<bool>,
    #[serde(rename = "CanDecontaminatePackagesWithFailedLabResults", skip_serializing_if = "Option::is_none")]
    pub can_decontaminate_packages_with_failed_lab_results: Option<bool>,
    #[serde(rename = "CanDeliverSalesToConsumers", skip_serializing_if = "Option::is_none")]
    pub can_deliver_sales_to_consumers: Option<bool>,
    #[serde(rename = "CanDeliverSalesToPatients", skip_serializing_if = "Option::is_none")]
    pub can_deliver_sales_to_patients: Option<bool>,
    #[serde(rename = "CanDestroyProduct", skip_serializing_if = "Option::is_none")]
    pub can_destroy_product: Option<bool>,
    #[serde(rename = "CanDonatePackages", skip_serializing_if = "Option::is_none")]
    pub can_donate_packages: Option<bool>,
    #[serde(rename = "CanDownloadProductLabel", skip_serializing_if = "Option::is_none")]
    pub can_download_product_label: Option<bool>,
    #[serde(rename = "CanGrowPlants", skip_serializing_if = "Option::is_none")]
    pub can_grow_plants: Option<bool>,
    #[serde(rename = "CanHaveMemberPatients", skip_serializing_if = "Option::is_none")]
    pub can_have_member_patients: Option<bool>,
    #[serde(rename = "CanInfuseProducts", skip_serializing_if = "Option::is_none")]
    pub can_infuse_products: Option<bool>,
    #[serde(rename = "CanManageItemBrands", skip_serializing_if = "Option::is_none")]
    pub can_manage_item_brands: Option<bool>,
    #[serde(rename = "CanPackageVegetativePlants", skip_serializing_if = "Option::is_none")]
    pub can_package_vegetative_plants: Option<bool>,
    #[serde(rename = "CanPackageWaste", skip_serializing_if = "Option::is_none")]
    pub can_package_waste: Option<bool>,
    #[serde(rename = "CanReceiveAssociateProductLabel", skip_serializing_if = "Option::is_none")]
    pub can_receive_associate_product_label: Option<bool>,
    #[serde(rename = "CanRecordProcessingJobs", skip_serializing_if = "Option::is_none")]
    pub can_record_processing_jobs: Option<bool>,
    #[serde(rename = "CanRecordProductForDestruction", skip_serializing_if = "Option::is_none")]
    pub can_record_product_for_destruction: Option<bool>,
    #[serde(rename = "CanRemediatePackagesWithFailedLabResults", skip_serializing_if = "Option::is_none")]
    pub can_remediate_packages_with_failed_lab_results: Option<bool>,
    #[serde(rename = "CanReportAdulteration", skip_serializing_if = "Option::is_none")]
    pub can_report_adulteration: Option<bool>,
    #[serde(rename = "CanReportHarvestSchedules", skip_serializing_if = "Option::is_none")]
    pub can_report_harvest_schedules: Option<bool>,
    #[serde(rename = "CanReportOperationalExceptions", skip_serializing_if = "Option::is_none")]
    pub can_report_operational_exceptions: Option<bool>,
    #[serde(rename = "CanReportPatientCheckIns", skip_serializing_if = "Option::is_none")]
    pub can_report_patient_check_ins: Option<bool>,
    #[serde(rename = "CanReportPatientsAdverseResponses", skip_serializing_if = "Option::is_none")]
    pub can_report_patients_adverse_responses: Option<bool>,
    #[serde(rename = "CanReportStrainProperties", skip_serializing_if = "Option::is_none")]
    pub can_report_strain_properties: Option<bool>,
    #[serde(rename = "CanRequestProductDecontamination", skip_serializing_if = "Option::is_none")]
    pub can_request_product_decontamination: Option<bool>,
    #[serde(rename = "CanRequestProductRemediation", skip_serializing_if = "Option::is_none")]
    pub can_request_product_remediation: Option<bool>,
    #[serde(rename = "CanRequireHarvestSampleLabTestBatches", skip_serializing_if = "Option::is_none")]
    pub can_require_harvest_sample_lab_test_batches: Option<bool>,
    #[serde(rename = "CanRequirePackageSampleLabTestBatches", skip_serializing_if = "Option::is_none")]
    pub can_require_package_sample_lab_test_batches: Option<bool>,
    #[serde(rename = "CanReturnPackageQuantity", skip_serializing_if = "Option::is_none")]
    pub can_return_package_quantity: Option<bool>,
    #[serde(rename = "CanSellToCaregivers", skip_serializing_if = "Option::is_none")]
    pub can_sell_to_caregivers: Option<bool>,
    #[serde(rename = "CanSellToConsumers", skip_serializing_if = "Option::is_none")]
    pub can_sell_to_consumers: Option<bool>,
    #[serde(rename = "CanSellToExternalPatients", skip_serializing_if = "Option::is_none")]
    pub can_sell_to_external_patients: Option<bool>,
    #[serde(rename = "CanSellToPatients", skip_serializing_if = "Option::is_none")]
    pub can_sell_to_patients: Option<bool>,
    #[serde(rename = "CanSpecifyPatientSalesLimitExemption", skip_serializing_if = "Option::is_none")]
    pub can_specify_patient_sales_limit_exemption: Option<bool>,
    #[serde(rename = "CanSubmitHarvestsForTesting", skip_serializing_if = "Option::is_none")]
    pub can_submit_harvests_for_testing: Option<bool>,
    #[serde(rename = "CanSubmitPackagesForTesting", skip_serializing_if = "Option::is_none")]
    pub can_submit_packages_for_testing: Option<bool>,
    #[serde(rename = "CanTagPlantBatches", skip_serializing_if = "Option::is_none")]
    pub can_tag_plant_batches: Option<bool>,
    #[serde(rename = "CanTestPackages", skip_serializing_if = "Option::is_none")]
    pub can_test_packages: Option<bool>,
    #[serde(rename = "CanTrackMotherPlantsAsGrowthPhase", skip_serializing_if = "Option::is_none")]
    pub can_track_mother_plants_as_growth_phase: Option<bool>,
    #[serde(rename = "CanTrackVegetativePlants", skip_serializing_if = "Option::is_none")]
    pub can_track_vegetative_plants: Option<bool>,
    #[serde(rename = "CanTransferFromExternalFacilities", skip_serializing_if = "Option::is_none")]
    pub can_transfer_from_external_facilities: Option<bool>,
    #[serde(rename = "CanUpdateLocationsOnPackages", skip_serializing_if = "Option::is_none")]
    pub can_update_locations_on_packages: Option<bool>,
    #[serde(rename = "CanUpdatePlantStrains", skip_serializing_if = "Option::is_none")]
    pub can_update_plant_strains: Option<bool>,
    #[serde(rename = "CanUseComplianceLabel", skip_serializing_if = "Option::is_none")]
    pub can_use_compliance_label: Option<bool>,
    #[serde(rename = "CanViewSourcePackages", skip_serializing_if = "Option::is_none")]
    pub can_view_source_packages: Option<bool>,
    #[serde(rename = "DeliverySalesToNonResidentConsumer", skip_serializing_if = "Option::is_none")]
    pub delivery_sales_to_non_resident_consumer: Option<bool>,
    #[serde(rename = "EnableExternalIdentifier", skip_serializing_if = "Option::is_none")]
    pub enable_external_identifier: Option<bool>,
    #[serde(rename = "EnableSublocations", skip_serializing_if = "Option::is_none")]
    pub enable_sublocations: Option<bool>,
    #[serde(rename = "IsHemp", skip_serializing_if = "Option::is_none")]
    pub is_hemp: Option<bool>,
    #[serde(rename = "IsMedical", skip_serializing_if = "Option::is_none")]
    pub is_medical: Option<bool>,
    #[serde(rename = "IsRetail", skip_serializing_if = "Option::is_none")]
    pub is_retail: Option<bool>,
    #[serde(rename = "IsSalesDeliveryHub", skip_serializing_if = "Option::is_none")]
    pub is_sales_delivery_hub: Option<bool>,
    #[serde(rename = "PackagesRequirePatientAffiliation", skip_serializing_if = "Option::is_none")]
    pub packages_require_patient_affiliation: Option<bool>,
    #[serde(rename = "PlantBatchesCanContainMotherPlants", skip_serializing_if = "Option::is_none")]
    pub plant_batches_can_contain_mother_plants: Option<bool>,
    #[serde(rename = "PlantsRequirePatientAffiliation", skip_serializing_if = "Option::is_none")]
    pub plants_require_patient_affiliation: Option<bool>,
    #[serde(rename = "RestrictHarvestPlantRestoreTimeHours", skip_serializing_if = "Option::is_none")]
    pub restrict_harvest_plant_restore_time_hours: Option<i64>,
    #[serde(rename = "RestrictPlantBatchAdjustmentTimeHours", skip_serializing_if = "Option::is_none")]
    pub restrict_plant_batch_adjustment_time_hours: Option<i64>,
    #[serde(rename = "RestrictWholesalePriceEditDays", skip_serializing_if = "Option::is_none")]
    pub restrict_wholesale_price_edit_days: Option<i64>,
    #[serde(rename = "RetailSalesToNonResidentConsumer", skip_serializing_if = "Option::is_none")]
    pub retail_sales_to_non_resident_consumer: Option<bool>,
    #[serde(rename = "RetailerDelivery", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery: Option<bool>,
    #[serde(rename = "RetailerDeliveryAllowDonations", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_allow_donations: Option<bool>,
    #[serde(rename = "RetailerDeliveryAllowPartialPackages", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_allow_partial_packages: Option<bool>,
    #[serde(rename = "RetailerDeliveryAllowTradeSamples", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_allow_trade_samples: Option<bool>,
    #[serde(rename = "RetailerDeliveryRequirePrice", skip_serializing_if = "Option::is_none")]
    pub retailer_delivery_require_price: Option<bool>,
    #[serde(rename = "SalesDeliveryAllowAddress", skip_serializing_if = "Option::is_none")]
    pub sales_delivery_allow_address: Option<bool>,
    #[serde(rename = "SalesDeliveryAllowCity", skip_serializing_if = "Option::is_none")]
    pub sales_delivery_allow_city: Option<bool>,
    #[serde(rename = "SalesDeliveryAllowCounty", skip_serializing_if = "Option::is_none")]
    pub sales_delivery_allow_county: Option<bool>,
    #[serde(rename = "SalesDeliveryAllowPlannedRoute", skip_serializing_if = "Option::is_none")]
    pub sales_delivery_allow_planned_route: Option<bool>,
    #[serde(rename = "SalesDeliveryAllowState", skip_serializing_if = "Option::is_none")]
    pub sales_delivery_allow_state: Option<bool>,
    #[serde(rename = "SalesDeliveryAllowZip", skip_serializing_if = "Option::is_none")]
    pub sales_delivery_allow_zip: Option<bool>,
    #[serde(rename = "SalesDeliveryRequireConsumerId", skip_serializing_if = "Option::is_none")]
    pub sales_delivery_require_consumer_id: Option<bool>,
    #[serde(rename = "SalesDeliveryRequirePatientNumber", skip_serializing_if = "Option::is_none")]
    pub sales_delivery_require_patient_number: Option<bool>,
    #[serde(rename = "SalesDeliveryRequireRecipientName", skip_serializing_if = "Option::is_none")]
    pub sales_delivery_require_recipient_name: Option<bool>,
    #[serde(rename = "SalesDeliveryRequireZone", skip_serializing_if = "Option::is_none")]
    pub sales_delivery_require_zone: Option<bool>,
    #[serde(rename = "SalesRequireCaregiverNumber", skip_serializing_if = "Option::is_none")]
    pub sales_require_caregiver_number: Option<bool>,
    #[serde(rename = "SalesRequireCaregiverPatientNumber", skip_serializing_if = "Option::is_none")]
    pub sales_require_caregiver_patient_number: Option<bool>,
    #[serde(rename = "SalesRequireExternalPatientIdentificationMethod", skip_serializing_if = "Option::is_none")]
    pub sales_require_external_patient_identification_method: Option<bool>,
    #[serde(rename = "SalesRequireExternalPatientNumber", skip_serializing_if = "Option::is_none")]
    pub sales_require_external_patient_number: Option<bool>,
    #[serde(rename = "SalesRequirePatientNumber", skip_serializing_if = "Option::is_none")]
    pub sales_require_patient_number: Option<bool>,
    #[serde(rename = "TaxExemptReportingFeesFacilityType", skip_serializing_if = "Option::is_none")]
    pub tax_exempt_reporting_fees_facility_type: Option<bool>,
    #[serde(rename = "TaxExemptTagOrdersFacilityType", skip_serializing_if = "Option::is_none")]
    pub tax_exempt_tag_orders_facility_type: Option<bool>,
    #[serde(rename = "TestsRequireLabSample", skip_serializing_if = "Option::is_none")]
    pub tests_require_lab_sample: Option<bool>,
    #[serde(rename = "TotalMemberPatientsAllowed", skip_serializing_if = "Option::is_none")]
    pub total_member_patients_allowed: Option<i64>,
}
