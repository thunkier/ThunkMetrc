using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class Facility
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Alias")]
        public string? Alias { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AllowTransferOfOnHoldPackages")]
        public required bool AllowTransferOfOnHoldPackages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AllowTransferOfOnRecallPackages")]
        public required bool AllowTransferOfOnRecallPackages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CredentialedDate")]
        public required string CredentialedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DisplayName")]
        public required string DisplayName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Email")]
        public string? Email { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityId")]
        public required int FacilityId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityType")]
        public required FacilityFacilityType FacilityType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HireDate")]
        public required string HireDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsFinancialContact")]
        public required bool IsFinancialContact { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsManager")]
        public required bool IsManager { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOwner")]
        public required bool IsOwner { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("License")]
        public required FacilityLicense License { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Occupations")]
        public required List<object> Occupations { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SupportActivationDate")]
        public string? SupportActivationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SupportExpirationDate")]
        public string? SupportExpirationDate { get; set; }
        public class FacilityFacilityType
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("AdvancedSales")]
            public required bool AdvancedSales { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanAccessCatalog")]
            public required bool CanAccessCatalog { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanAdjustSourcePackagesWithPartials")]
            public required bool CanAdjustSourcePackagesWithPartials { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanAssignLocationsToPackages")]
            public required bool CanAssignLocationsToPackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanAssignLocationsToPlantBatches")]
            public required bool CanAssignLocationsToPlantBatches { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanClonePlantBatches")]
            public required bool CanClonePlantBatches { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanCreateDerivedPackages")]
            public required bool CanCreateDerivedPackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanCreateImmaturePlantPackagesFromPlants")]
            public required bool CanCreateImmaturePlantPackagesFromPlants { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanCreateOpeningBalancePackages")]
            public required bool CanCreateOpeningBalancePackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanCreateOpeningBalancePlantBatches")]
            public required bool CanCreateOpeningBalancePlantBatches { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanCreatePartialPackages")]
            public required bool CanCreatePartialPackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanCreateProcessValidationPackages")]
            public required bool CanCreateProcessValidationPackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanCreateTradeSamplePackages")]
            public required bool CanCreateTradeSamplePackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanDecontaminatePackagesWithFailedLabResults")]
            public required bool CanDecontaminatePackagesWithFailedLabResults { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanDeliverSalesToConsumers")]
            public required bool CanDeliverSalesToConsumers { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanDeliverSalesToPatients")]
            public required bool CanDeliverSalesToPatients { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanDestroyProduct")]
            public required bool CanDestroyProduct { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanDonatePackages")]
            public required bool CanDonatePackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanDownloadProductLabel")]
            public required bool CanDownloadProductLabel { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanGrowPlants")]
            public required bool CanGrowPlants { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanHaveMemberPatients")]
            public required bool CanHaveMemberPatients { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanInfuseProducts")]
            public required bool CanInfuseProducts { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanManageItemBrands")]
            public required bool CanManageItemBrands { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanPackageVegetativePlants")]
            public required bool CanPackageVegetativePlants { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanPackageWaste")]
            public required bool CanPackageWaste { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanReceiveAssociateProductLabel")]
            public required bool CanReceiveAssociateProductLabel { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanRecordProcessingJobs")]
            public required bool CanRecordProcessingJobs { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanRecordProductForDestruction")]
            public required bool CanRecordProductForDestruction { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanRemediatePackagesWithFailedLabResults")]
            public required bool CanRemediatePackagesWithFailedLabResults { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanReportAdulteration")]
            public required bool CanReportAdulteration { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanReportHarvestSchedules")]
            public required bool CanReportHarvestSchedules { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanReportOperationalExceptions")]
            public required bool CanReportOperationalExceptions { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanReportPatientCheckIns")]
            public required bool CanReportPatientCheckIns { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanReportPatientsAdverseResponses")]
            public required bool CanReportPatientsAdverseResponses { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanReportStrainProperties")]
            public required bool CanReportStrainProperties { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanRequestProductDecontamination")]
            public required bool CanRequestProductDecontamination { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanRequestProductRemediation")]
            public required bool CanRequestProductRemediation { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanRequireHarvestSampleLabTestBatches")]
            public required bool CanRequireHarvestSampleLabTestBatches { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanRequirePackageSampleLabTestBatches")]
            public required bool CanRequirePackageSampleLabTestBatches { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanReturnPackageQuantity")]
            public required bool CanReturnPackageQuantity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanSellToCaregivers")]
            public required bool CanSellToCaregivers { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanSellToConsumers")]
            public required bool CanSellToConsumers { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanSellToExternalPatients")]
            public required bool CanSellToExternalPatients { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanSellToPatients")]
            public required bool CanSellToPatients { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanSpecifyPatientSalesLimitExemption")]
            public required bool CanSpecifyPatientSalesLimitExemption { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanSubmitHarvestsForTesting")]
            public required bool CanSubmitHarvestsForTesting { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanSubmitPackagesForTesting")]
            public required bool CanSubmitPackagesForTesting { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanTagPlantBatches")]
            public required bool CanTagPlantBatches { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanTestPackages")]
            public required bool CanTestPackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanTrackMotherPlantsAsGrowthPhase")]
            public required bool CanTrackMotherPlantsAsGrowthPhase { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanTrackVegetativePlants")]
            public required bool CanTrackVegetativePlants { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanTransferFromExternalFacilities")]
            public required bool CanTransferFromExternalFacilities { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanUpdateLocationsOnPackages")]
            public required bool CanUpdateLocationsOnPackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanUpdatePlantStrains")]
            public required bool CanUpdatePlantStrains { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanUseComplianceLabel")]
            public required bool CanUseComplianceLabel { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CanViewSourcePackages")]
            public required bool CanViewSourcePackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("DeliverySalesToNonResidentConsumer")]
            public required bool DeliverySalesToNonResidentConsumer { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EnableExternalIdentifier")]
            public required bool EnableExternalIdentifier { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EnableSublocations")]
            public required bool EnableSublocations { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsHemp")]
            public required bool IsHemp { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsMedical")]
            public required bool IsMedical { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsRetail")]
            public required bool IsRetail { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsSalesDeliveryHub")]
            public required bool IsSalesDeliveryHub { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PackagesRequirePatientAffiliation")]
            public required bool PackagesRequirePatientAffiliation { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchesCanContainMotherPlants")]
            public required bool PlantBatchesCanContainMotherPlants { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PlantsRequirePatientAffiliation")]
            public required bool PlantsRequirePatientAffiliation { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RestrictHarvestPlantRestoreTimeHours")]
            public required int RestrictHarvestPlantRestoreTimeHours { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RestrictPlantBatchAdjustmentTimeHours")]
            public int? RestrictPlantBatchAdjustmentTimeHours { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RestrictWholesalePriceEditDays")]
            public int? RestrictWholesalePriceEditDays { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RetailSalesToNonResidentConsumer")]
            public required bool RetailSalesToNonResidentConsumer { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDelivery")]
            public required bool RetailerDelivery { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryAllowDonations")]
            public required bool RetailerDeliveryAllowDonations { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryAllowPartialPackages")]
            public required bool RetailerDeliveryAllowPartialPackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryAllowTradeSamples")]
            public required bool RetailerDeliveryAllowTradeSamples { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryRequirePrice")]
            public required bool RetailerDeliveryRequirePrice { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDeliveryAllowAddress")]
            public required bool SalesDeliveryAllowAddress { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDeliveryAllowCity")]
            public required bool SalesDeliveryAllowCity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDeliveryAllowCounty")]
            public required bool SalesDeliveryAllowCounty { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDeliveryAllowPlannedRoute")]
            public required bool SalesDeliveryAllowPlannedRoute { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDeliveryAllowState")]
            public required bool SalesDeliveryAllowState { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDeliveryAllowZip")]
            public required bool SalesDeliveryAllowZip { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDeliveryRequireConsumerId")]
            public required bool SalesDeliveryRequireConsumerId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDeliveryRequirePatientNumber")]
            public required bool SalesDeliveryRequirePatientNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDeliveryRequireRecipientName")]
            public required bool SalesDeliveryRequireRecipientName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDeliveryRequireZone")]
            public required bool SalesDeliveryRequireZone { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesRequireCaregiverNumber")]
            public required bool SalesRequireCaregiverNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesRequireCaregiverPatientNumber")]
            public required bool SalesRequireCaregiverPatientNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesRequireExternalPatientIdentificationMethod")]
            public required bool SalesRequireExternalPatientIdentificationMethod { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesRequireExternalPatientNumber")]
            public required bool SalesRequireExternalPatientNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesRequirePatientNumber")]
            public required bool SalesRequirePatientNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TaxExemptReportingFeesFacilityType")]
            public required bool TaxExemptReportingFeesFacilityType { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TaxExemptTagOrdersFacilityType")]
            public required bool TaxExemptTagOrdersFacilityType { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TestsRequireLabSample")]
            public required bool TestsRequireLabSample { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TotalMemberPatientsAllowed")]
            public int? TotalMemberPatientsAllowed { get; set; }
        }
        public class FacilityLicense
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("EmployeeLicenseNumber")]
            public required string EmployeeLicenseNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EndDate")]
            public required string EndDate { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LicenseType")]
            public required string LicenseType { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Number")]
            public required string Number { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("StartDate")]
            public required string StartDate { get; set; }
        }
    }
}
