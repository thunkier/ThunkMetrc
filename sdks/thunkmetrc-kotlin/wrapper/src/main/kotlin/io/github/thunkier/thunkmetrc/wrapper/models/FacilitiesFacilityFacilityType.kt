package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class FacilitiesFacilityFacilityType(
    @JsonProperty("AdvancedSales")
    val advancedSales: Boolean? = null,
    @JsonProperty("CanAccessCatalog")
    val canAccessCatalog: Boolean? = null,
    @JsonProperty("CanAdjustSourcePackagesWithPartials")
    val canAdjustSourcePackagesWithPartials: Boolean? = null,
    @JsonProperty("CanAssignLocationsToPackages")
    val canAssignLocationsToPackages: Boolean? = null,
    @JsonProperty("CanAssignLocationsToPlantBatches")
    val canAssignLocationsToPlantBatches: Boolean? = null,
    @JsonProperty("CanClonePlantBatches")
    val canClonePlantBatches: Boolean? = null,
    @JsonProperty("CanCreateDerivedPackages")
    val canCreateDerivedPackages: Boolean? = null,
    @JsonProperty("CanCreateImmaturePlantPackagesFromPlants")
    val canCreateImmaturePlantPackagesFromPlants: Boolean? = null,
    @JsonProperty("CanCreateOpeningBalancePackages")
    val canCreateOpeningBalancePackages: Boolean? = null,
    @JsonProperty("CanCreateOpeningBalancePlantBatches")
    val canCreateOpeningBalancePlantBatches: Boolean? = null,
    @JsonProperty("CanCreatePartialPackages")
    val canCreatePartialPackages: Boolean? = null,
    @JsonProperty("CanCreateProcessValidationPackages")
    val canCreateProcessValidationPackages: Boolean? = null,
    @JsonProperty("CanCreateTradeSamplePackages")
    val canCreateTradeSamplePackages: Boolean? = null,
    @JsonProperty("CanDecontaminatePackagesWithFailedLabResults")
    val canDecontaminatePackagesWithFailedLabResults: Boolean? = null,
    @JsonProperty("CanDeliverSalesToConsumers")
    val canDeliverSalesToConsumers: Boolean? = null,
    @JsonProperty("CanDeliverSalesToPatients")
    val canDeliverSalesToPatients: Boolean? = null,
    @JsonProperty("CanDestroyProduct")
    val canDestroyProduct: Boolean? = null,
    @JsonProperty("CanDonatePackages")
    val canDonatePackages: Boolean? = null,
    @JsonProperty("CanDownloadProductLabel")
    val canDownloadProductLabel: Boolean? = null,
    @JsonProperty("CanGrowPlants")
    val canGrowPlants: Boolean? = null,
    @JsonProperty("CanHaveMemberPatients")
    val canHaveMemberPatients: Boolean? = null,
    @JsonProperty("CanInfuseProducts")
    val canInfuseProducts: Boolean? = null,
    @JsonProperty("CanManageItemBrands")
    val canManageItemBrands: Boolean? = null,
    @JsonProperty("CanPackageVegetativePlants")
    val canPackageVegetativePlants: Boolean? = null,
    @JsonProperty("CanPackageWaste")
    val canPackageWaste: Boolean? = null,
    @JsonProperty("CanReceiveAssociateProductLabel")
    val canReceiveAssociateProductLabel: Boolean? = null,
    @JsonProperty("CanRecordProcessingJobs")
    val canRecordProcessingJobs: Boolean? = null,
    @JsonProperty("CanRecordProductForDestruction")
    val canRecordProductForDestruction: Boolean? = null,
    @JsonProperty("CanRemediatePackagesWithFailedLabResults")
    val canRemediatePackagesWithFailedLabResults: Boolean? = null,
    @JsonProperty("CanReportAdulteration")
    val canReportAdulteration: Boolean? = null,
    @JsonProperty("CanReportHarvestSchedules")
    val canReportHarvestSchedules: Boolean? = null,
    @JsonProperty("CanReportOperationalExceptions")
    val canReportOperationalExceptions: Boolean? = null,
    @JsonProperty("CanReportPatientCheckIns")
    val canReportPatientCheckIns: Boolean? = null,
    @JsonProperty("CanReportPatientsAdverseResponses")
    val canReportPatientsAdverseResponses: Boolean? = null,
    @JsonProperty("CanReportStrainProperties")
    val canReportStrainProperties: Boolean? = null,
    @JsonProperty("CanRequestProductDecontamination")
    val canRequestProductDecontamination: Boolean? = null,
    @JsonProperty("CanRequestProductRemediation")
    val canRequestProductRemediation: Boolean? = null,
    @JsonProperty("CanRequireHarvestSampleLabTestBatches")
    val canRequireHarvestSampleLabTestBatches: Boolean? = null,
    @JsonProperty("CanRequirePackageSampleLabTestBatches")
    val canRequirePackageSampleLabTestBatches: Boolean? = null,
    @JsonProperty("CanReturnPackageQuantity")
    val canReturnPackageQuantity: Boolean? = null,
    @JsonProperty("CanSellToCaregivers")
    val canSellToCaregivers: Boolean? = null,
    @JsonProperty("CanSellToConsumers")
    val canSellToConsumers: Boolean? = null,
    @JsonProperty("CanSellToExternalPatients")
    val canSellToExternalPatients: Boolean? = null,
    @JsonProperty("CanSellToPatients")
    val canSellToPatients: Boolean? = null,
    @JsonProperty("CanSpecifyPatientSalesLimitExemption")
    val canSpecifyPatientSalesLimitExemption: Boolean? = null,
    @JsonProperty("CanSubmitHarvestsForTesting")
    val canSubmitHarvestsForTesting: Boolean? = null,
    @JsonProperty("CanSubmitPackagesForTesting")
    val canSubmitPackagesForTesting: Boolean? = null,
    @JsonProperty("CanTagPlantBatches")
    val canTagPlantBatches: Boolean? = null,
    @JsonProperty("CanTestPackages")
    val canTestPackages: Boolean? = null,
    @JsonProperty("CanTrackMotherPlantsAsGrowthPhase")
    val canTrackMotherPlantsAsGrowthPhase: Boolean? = null,
    @JsonProperty("CanTrackVegetativePlants")
    val canTrackVegetativePlants: Boolean? = null,
    @JsonProperty("CanTransferFromExternalFacilities")
    val canTransferFromExternalFacilities: Boolean? = null,
    @JsonProperty("CanUpdateLocationsOnPackages")
    val canUpdateLocationsOnPackages: Boolean? = null,
    @JsonProperty("CanUpdatePlantStrains")
    val canUpdatePlantStrains: Boolean? = null,
    @JsonProperty("CanUseComplianceLabel")
    val canUseComplianceLabel: Boolean? = null,
    @JsonProperty("CanViewSourcePackages")
    val canViewSourcePackages: Boolean? = null,
    @JsonProperty("DeliverySalesToNonResidentConsumer")
    val deliverySalesToNonResidentConsumer: Boolean? = null,
    @JsonProperty("EnableExternalIdentifier")
    val enableExternalIdentifier: Boolean? = null,
    @JsonProperty("EnableSublocations")
    val enableSublocations: Boolean? = null,
    @JsonProperty("IsHemp")
    val isHemp: Boolean? = null,
    @JsonProperty("IsMedical")
    val isMedical: Boolean? = null,
    @JsonProperty("IsRetail")
    val isRetail: Boolean? = null,
    @JsonProperty("IsSalesDeliveryHub")
    val isSalesDeliveryHub: Boolean? = null,
    @JsonProperty("PackagesRequirePatientAffiliation")
    val packagesRequirePatientAffiliation: Boolean? = null,
    @JsonProperty("PlantBatchesCanContainMotherPlants")
    val plantBatchesCanContainMotherPlants: Boolean? = null,
    @JsonProperty("PlantsRequirePatientAffiliation")
    val plantsRequirePatientAffiliation: Boolean? = null,
    @JsonProperty("RestrictHarvestPlantRestoreTimeHours")
    val restrictHarvestPlantRestoreTimeHours: Int? = null,
    @JsonProperty("RestrictPlantBatchAdjustmentTimeHours")
    val restrictPlantBatchAdjustmentTimeHours: Int? = null,
    @JsonProperty("RestrictWholesalePriceEditDays")
    val restrictWholesalePriceEditDays: Int? = null,
    @JsonProperty("RetailSalesToNonResidentConsumer")
    val retailSalesToNonResidentConsumer: Boolean? = null,
    @JsonProperty("RetailerDelivery")
    val retailerDelivery: Boolean? = null,
    @JsonProperty("RetailerDeliveryAllowDonations")
    val retailerDeliveryAllowDonations: Boolean? = null,
    @JsonProperty("RetailerDeliveryAllowPartialPackages")
    val retailerDeliveryAllowPartialPackages: Boolean? = null,
    @JsonProperty("RetailerDeliveryAllowTradeSamples")
    val retailerDeliveryAllowTradeSamples: Boolean? = null,
    @JsonProperty("RetailerDeliveryRequirePrice")
    val retailerDeliveryRequirePrice: Boolean? = null,
    @JsonProperty("SalesDeliveryAllowAddress")
    val salesDeliveryAllowAddress: Boolean? = null,
    @JsonProperty("SalesDeliveryAllowCity")
    val salesDeliveryAllowCity: Boolean? = null,
    @JsonProperty("SalesDeliveryAllowCounty")
    val salesDeliveryAllowCounty: Boolean? = null,
    @JsonProperty("SalesDeliveryAllowPlannedRoute")
    val salesDeliveryAllowPlannedRoute: Boolean? = null,
    @JsonProperty("SalesDeliveryAllowState")
    val salesDeliveryAllowState: Boolean? = null,
    @JsonProperty("SalesDeliveryAllowZip")
    val salesDeliveryAllowZip: Boolean? = null,
    @JsonProperty("SalesDeliveryRequireConsumerId")
    val salesDeliveryRequireConsumerId: Boolean? = null,
    @JsonProperty("SalesDeliveryRequirePatientNumber")
    val salesDeliveryRequirePatientNumber: Boolean? = null,
    @JsonProperty("SalesDeliveryRequireRecipientName")
    val salesDeliveryRequireRecipientName: Boolean? = null,
    @JsonProperty("SalesDeliveryRequireZone")
    val salesDeliveryRequireZone: Boolean? = null,
    @JsonProperty("SalesRequireCaregiverNumber")
    val salesRequireCaregiverNumber: Boolean? = null,
    @JsonProperty("SalesRequireCaregiverPatientNumber")
    val salesRequireCaregiverPatientNumber: Boolean? = null,
    @JsonProperty("SalesRequireExternalPatientIdentificationMethod")
    val salesRequireExternalPatientIdentificationMethod: Boolean? = null,
    @JsonProperty("SalesRequireExternalPatientNumber")
    val salesRequireExternalPatientNumber: Boolean? = null,
    @JsonProperty("SalesRequirePatientNumber")
    val salesRequirePatientNumber: Boolean? = null,
    @JsonProperty("TaxExemptReportingFeesFacilityType")
    val taxExemptReportingFeesFacilityType: Boolean? = null,
    @JsonProperty("TaxExemptTagOrdersFacilityType")
    val taxExemptTagOrdersFacilityType: Boolean? = null,
    @JsonProperty("TestsRequireLabSample")
    val testsRequireLabSample: Boolean? = null,
    @JsonProperty("TotalMemberPatientsAllowed")
    val totalMemberPatientsAllowed: Int? = null
)
