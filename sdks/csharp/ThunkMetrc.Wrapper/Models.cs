using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper
{

    public class PlantBatchesCreateAdditivesV1RequestItem
    {
        public List<PlantBatchesCreateAdditivesV1RequestItemActiveIngredients>? ActiveIngredients { get; set; }
        public string? ActualDate { get; set; }
        public string? AdditiveType { get; set; }
        public string? ApplicationDevice { get; set; }
        public string? EpaRegistrationNumber { get; set; }
        public string? PlantBatchName { get; set; }
        public string? ProductSupplier { get; set; }
        public string? ProductTradeName { get; set; }
        public int TotalAmountApplied { get; set; }
        public string? TotalAmountUnitOfMeasure { get; set; }
    }

    public class PlantBatchesCreateAdditivesV1RequestItemActiveIngredients
    {
        public string? Name { get; set; }
        public double Percentage { get; set; }
    }

    public class PlantBatchesCreateAdditivesV2RequestItem
    {
        public List<PlantBatchesCreateAdditivesV2RequestItemActiveIngredients>? ActiveIngredients { get; set; }
        public string? ActualDate { get; set; }
        public string? AdditiveType { get; set; }
        public string? ApplicationDevice { get; set; }
        public string? EpaRegistrationNumber { get; set; }
        public string? PlantBatchName { get; set; }
        public string? ProductSupplier { get; set; }
        public string? ProductTradeName { get; set; }
        public int TotalAmountApplied { get; set; }
        public string? TotalAmountUnitOfMeasure { get; set; }
    }

    public class PlantBatchesCreateAdditivesV2RequestItemActiveIngredients
    {
        public string? Name { get; set; }
        public double Percentage { get; set; }
    }

    public class PlantBatchesCreateAdditivesUsingtemplateV2RequestItem
    {
        public string? ActualDate { get; set; }
        public string? AdditivesTemplateName { get; set; }
        public string? PlantBatchName { get; set; }
        public string? Rate { get; set; }
        public int TotalAmountApplied { get; set; }
        public string? TotalAmountUnitOfMeasure { get; set; }
        public string? Volume { get; set; }
    }

    public class PlantBatchesCreateAdjustV1RequestItem
    {
        public string? AdjustmentDate { get; set; }
        public string? AdjustmentReason { get; set; }
        public string? PlantBatchName { get; set; }
        public int Quantity { get; set; }
        public string? ReasonNote { get; set; }
    }

    public class PlantBatchesCreateAdjustV2RequestItem
    {
        public string? AdjustmentDate { get; set; }
        public string? AdjustmentReason { get; set; }
        public string? PlantBatchName { get; set; }
        public int Quantity { get; set; }
        public string? ReasonNote { get; set; }
    }

    public class PlantBatchesCreateChangegrowthphaseV1RequestItem
    {
        public int Count { get; set; }
        public string? CountPerPlant { get; set; }
        public string? GrowthDate { get; set; }
        public string? GrowthPhase { get; set; }
        public string? Name { get; set; }
        public string? NewLocation { get; set; }
        public string? NewSublocation { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? StartingTag { get; set; }
    }

    public class PlantBatchesCreateGrowthphaseV2RequestItem
    {
        public int Count { get; set; }
        public string? CountPerPlant { get; set; }
        public string? GrowthDate { get; set; }
        public string? GrowthPhase { get; set; }
        public string? Name { get; set; }
        public string? NewLocation { get; set; }
        public string? NewSublocation { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? StartingTag { get; set; }
    }

    public class PlantBatchesCreatePackageV2RequestItem
    {
        public string? ActualDate { get; set; }
        public int Count { get; set; }
        public string? ExpirationDate { get; set; }
        public int? Id { get; set; }
        public bool IsDonation { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatch { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UseByDate { get; set; }
    }

    public class PlantBatchesCreatePackageFrommotherplantV1RequestItem
    {
        public string? ActualDate { get; set; }
        public int Count { get; set; }
        public string? ExpirationDate { get; set; }
        public int? Id { get; set; }
        public bool IsDonation { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatch { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UseByDate { get; set; }
    }

    public class PlantBatchesCreatePackageFrommotherplantV2RequestItem
    {
        public string? ActualDate { get; set; }
        public int Count { get; set; }
        public string? ExpirationDate { get; set; }
        public int? Id { get; set; }
        public bool IsDonation { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatch { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UseByDate { get; set; }
    }

    public class PlantBatchesCreatePlantingsV2RequestItem
    {
        public string? ActualDate { get; set; }
        public int Count { get; set; }
        public string? Location { get; set; }
        public string? Name { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? SourcePlantBatches { get; set; }
        public string? Strain { get; set; }
        public string? Sublocation { get; set; }
        public string? Type { get; set; }
    }

    public class PlantBatchesCreateSplitV1RequestItem
    {
        public string? ActualDate { get; set; }
        public int Count { get; set; }
        public string? GroupName { get; set; }
        public string? Location { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatch { get; set; }
        public string? Strain { get; set; }
        public string? Sublocation { get; set; }
    }

    public class PlantBatchesCreateSplitV2RequestItem
    {
        public string? ActualDate { get; set; }
        public int Count { get; set; }
        public string? GroupName { get; set; }
        public string? Location { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatch { get; set; }
        public string? Strain { get; set; }
        public string? Sublocation { get; set; }
    }

    public class PlantBatchesCreateWasteV1RequestItem
    {
        public string? MixedMaterial { get; set; }
        public string? Note { get; set; }
        public string? PlantBatchName { get; set; }
        public string? ReasonName { get; set; }
        public string? UnitOfMeasureName { get; set; }
        public string? WasteDate { get; set; }
        public string? WasteMethodName { get; set; }
        public double WasteWeight { get; set; }
    }

    public class PlantBatchesCreateWasteV2RequestItem
    {
        public string? MixedMaterial { get; set; }
        public string? Note { get; set; }
        public string? PlantBatchName { get; set; }
        public string? ReasonName { get; set; }
        public string? UnitOfMeasureName { get; set; }
        public string? WasteDate { get; set; }
        public string? WasteMethodName { get; set; }
        public double WasteWeight { get; set; }
    }

    public class PlantBatchesCreatepackagesV1RequestItem
    {
        public string? ActualDate { get; set; }
        public int Count { get; set; }
        public string? ExpirationDate { get; set; }
        public int? Id { get; set; }
        public bool IsDonation { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatch { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UseByDate { get; set; }
    }

    public class PlantBatchesCreateplantingsV1RequestItem
    {
        public string? ActualDate { get; set; }
        public int Count { get; set; }
        public string? Location { get; set; }
        public string? Name { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? SourcePlantBatches { get; set; }
        public string? Strain { get; set; }
        public string? Sublocation { get; set; }
        public string? Type { get; set; }
    }

    public class PlantBatchesUpdateLocationV2RequestItem
    {
        public string? Location { get; set; }
        public string? MoveDate { get; set; }
        public string? Name { get; set; }
        public string? Sublocation { get; set; }
    }

    public class PlantBatchesUpdateMoveplantbatchesV1RequestItem
    {
        public string? Location { get; set; }
        public string? MoveDate { get; set; }
        public string? Name { get; set; }
        public string? Sublocation { get; set; }
    }

    public class PlantBatchesUpdateNameV2RequestItem
    {
        public string? Group { get; set; }
        public int? Id { get; set; }
        public string? NewGroup { get; set; }
    }

    public class PlantBatchesUpdateStrainV2RequestItem
    {
        public int Id { get; set; }
        public string? Name { get; set; }
        public int StrainId { get; set; }
        public string? StrainName { get; set; }
    }

    public class PlantBatchesUpdateTagV2RequestItem
    {
        public string? Group { get; set; }
        public int? Id { get; set; }
        public string? NewTag { get; set; }
        public string? ReplaceDate { get; set; }
        public int? TagId { get; set; }
    }

    public class ProcessingJobsCreateAdjustV1RequestItem
    {
        public string? AdjustmentDate { get; set; }
        public string? AdjustmentNote { get; set; }
        public string? AdjustmentReason { get; set; }
        public string? CountUnitOfMeasureName { get; set; }
        public int Id { get; set; }
        public List<ProcessingJobsCreateAdjustV1RequestItemPackages>? Packages { get; set; }
        public string? VolumeUnitOfMeasureName { get; set; }
        public string? WeightUnitOfMeasureName { get; set; }
    }

    public class ProcessingJobsCreateAdjustV1RequestItemPackages
    {
        public string? Label { get; set; }
        public int Quantity { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class ProcessingJobsCreateAdjustV2RequestItem
    {
        public string? AdjustmentDate { get; set; }
        public string? AdjustmentNote { get; set; }
        public string? AdjustmentReason { get; set; }
        public string? CountUnitOfMeasureName { get; set; }
        public int Id { get; set; }
        public List<ProcessingJobsCreateAdjustV2RequestItemPackages>? Packages { get; set; }
        public string? VolumeUnitOfMeasureName { get; set; }
        public string? WeightUnitOfMeasureName { get; set; }
    }

    public class ProcessingJobsCreateAdjustV2RequestItemPackages
    {
        public string? Label { get; set; }
        public int Quantity { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class ProcessingJobsCreateJobtypesV1RequestItem
    {
        public List<string?>? Attributes { get; set; }
        public string? Category { get; set; }
        public string? Description { get; set; }
        public string? Name { get; set; }
        public string? ProcessingSteps { get; set; }
    }

    public class ProcessingJobsCreateJobtypesV2RequestItem
    {
        public List<string?>? Attributes { get; set; }
        public string? Category { get; set; }
        public string? Description { get; set; }
        public string? Name { get; set; }
        public string? ProcessingSteps { get; set; }
    }

    public class ProcessingJobsCreateStartV1RequestItem
    {
        public string? CountUnitOfMeasure { get; set; }
        public string? JobName { get; set; }
        public string? JobType { get; set; }
        public List<ProcessingJobsCreateStartV1RequestItemPackages>? Packages { get; set; }
        public string? StartDate { get; set; }
        public string? VolumeUnitOfMeasure { get; set; }
        public string? WeightUnitOfMeasure { get; set; }
    }

    public class ProcessingJobsCreateStartV1RequestItemPackages
    {
        public string? Label { get; set; }
        public int Quantity { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class ProcessingJobsCreateStartV2RequestItem
    {
        public string? CountUnitOfMeasure { get; set; }
        public string? JobName { get; set; }
        public string? JobType { get; set; }
        public List<ProcessingJobsCreateStartV2RequestItemPackages>? Packages { get; set; }
        public string? StartDate { get; set; }
        public string? VolumeUnitOfMeasure { get; set; }
        public string? WeightUnitOfMeasure { get; set; }
    }

    public class ProcessingJobsCreateStartV2RequestItemPackages
    {
        public string? Label { get; set; }
        public int Quantity { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class ProcessingJobsCreatepackagesV1RequestItem
    {
        public string? ExpirationDate { get; set; }
        public string? FinishDate { get; set; }
        public string? FinishNote { get; set; }
        public bool FinishProcessingJob { get; set; }
        public string? Item { get; set; }
        public string? JobName { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PackageDate { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? ProductionBatchNumber { get; set; }
        public int Quantity { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UnitOfMeasure { get; set; }
        public string? UseByDate { get; set; }
        public string? WasteCountQuantity { get; set; }
        public string? WasteCountUnitOfMeasureName { get; set; }
        public string? WasteVolumeQuantity { get; set; }
        public string? WasteVolumeUnitOfMeasureName { get; set; }
        public string? WasteWeightQuantity { get; set; }
        public string? WasteWeightUnitOfMeasureName { get; set; }
    }

    public class ProcessingJobsCreatepackagesV2RequestItem
    {
        public string? ExpirationDate { get; set; }
        public string? FinishDate { get; set; }
        public string? FinishNote { get; set; }
        public bool FinishProcessingJob { get; set; }
        public string? Item { get; set; }
        public string? JobName { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PackageDate { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? ProductionBatchNumber { get; set; }
        public int Quantity { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UnitOfMeasure { get; set; }
        public string? UseByDate { get; set; }
        public string? WasteCountQuantity { get; set; }
        public string? WasteCountUnitOfMeasureName { get; set; }
        public string? WasteVolumeQuantity { get; set; }
        public string? WasteVolumeUnitOfMeasureName { get; set; }
        public string? WasteWeightQuantity { get; set; }
        public string? WasteWeightUnitOfMeasureName { get; set; }
    }

    public class ProcessingJobsUpdateFinishV1RequestItem
    {
        public string? FinishDate { get; set; }
        public string? FinishNote { get; set; }
        public int Id { get; set; }
        public string? TotalCountWaste { get; set; }
        public string? TotalVolumeWaste { get; set; }
        public int TotalWeightWaste { get; set; }
        public string? WasteCountUnitOfMeasureName { get; set; }
        public string? WasteVolumeUnitOfMeasureName { get; set; }
        public string? WasteWeightUnitOfMeasureName { get; set; }
    }

    public class ProcessingJobsUpdateFinishV2RequestItem
    {
        public string? FinishDate { get; set; }
        public string? FinishNote { get; set; }
        public int Id { get; set; }
        public string? TotalCountWaste { get; set; }
        public string? TotalVolumeWaste { get; set; }
        public int TotalWeightWaste { get; set; }
        public string? WasteCountUnitOfMeasureName { get; set; }
        public string? WasteVolumeUnitOfMeasureName { get; set; }
        public string? WasteWeightUnitOfMeasureName { get; set; }
    }

    public class ProcessingJobsUpdateJobtypesV1RequestItem
    {
        public List<string?>? Attributes { get; set; }
        public string? CategoryName { get; set; }
        public string? Description { get; set; }
        public int Id { get; set; }
        public string? Name { get; set; }
        public string? ProcessingSteps { get; set; }
    }

    public class ProcessingJobsUpdateJobtypesV2RequestItem
    {
        public List<string?>? Attributes { get; set; }
        public string? CategoryName { get; set; }
        public string? Description { get; set; }
        public int Id { get; set; }
        public string? Name { get; set; }
        public string? ProcessingSteps { get; set; }
    }

    public class ProcessingJobsUpdateUnfinishV1RequestItem
    {
        public int Id { get; set; }
    }

    public class ProcessingJobsUpdateUnfinishV2RequestItem
    {
        public int Id { get; set; }
    }

    public class SublocationsCreateV2RequestItem
    {
        public string? Name { get; set; }
    }

    public class SublocationsUpdateV2RequestItem
    {
        public int Id { get; set; }
        public string? Name { get; set; }
    }

    public class TransfersCreateExternalIncomingV1RequestItem
    {
        public List<TransfersCreateExternalIncomingV1RequestItemDestinations>? Destinations { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? ShipperAddress1 { get; set; }
        public string? ShipperAddress2 { get; set; }
        public string? ShipperAddressCity { get; set; }
        public string? ShipperAddressPostalCode { get; set; }
        public string? ShipperAddressState { get; set; }
        public string? ShipperLicenseNumber { get; set; }
        public string? ShipperMainPhoneNumber { get; set; }
        public string? ShipperName { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersCreateExternalIncomingV1RequestItemDestinations
    {
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int? GrossUnitOfWeightId { get; set; }
        public double? GrossWeight { get; set; }
        public string? InvoiceNumber { get; set; }
        public List<TransfersCreateExternalIncomingV1RequestItemDestinationsPackages>? Packages { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientLicenseNumber { get; set; }
        public string? TransferTypeName { get; set; }
        public List<TransfersCreateExternalIncomingV1RequestItemDestinationsTransporters>? Transporters { get; set; }
    }

    public class TransfersCreateExternalIncomingV1RequestItemDestinationsPackages
    {
        public string? ExpirationDate { get; set; }
        public string? ExternalId { get; set; }
        public string? GrossUnitOfWeightName { get; set; }
        public double GrossWeight { get; set; }
        public string? ItemName { get; set; }
        public string? PackagedDate { get; set; }
        public int Quantity { get; set; }
        public string? SellByDate { get; set; }
        public string? UnitOfMeasureName { get; set; }
        public string? UseByDate { get; set; }
        public string? WholesalePrice { get; set; }
    }

    public class TransfersCreateExternalIncomingV1RequestItemDestinationsTransporters
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public bool IsLayover { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? TransporterDetails { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersCreateExternalIncomingV2RequestItem
    {
        public List<TransfersCreateExternalIncomingV2RequestItemDestinations>? Destinations { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? ShipperAddress1 { get; set; }
        public string? ShipperAddress2 { get; set; }
        public string? ShipperAddressCity { get; set; }
        public string? ShipperAddressPostalCode { get; set; }
        public string? ShipperAddressState { get; set; }
        public string? ShipperLicenseNumber { get; set; }
        public string? ShipperMainPhoneNumber { get; set; }
        public string? ShipperName { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersCreateExternalIncomingV2RequestItemDestinations
    {
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int? GrossUnitOfWeightId { get; set; }
        public double? GrossWeight { get; set; }
        public string? InvoiceNumber { get; set; }
        public List<TransfersCreateExternalIncomingV2RequestItemDestinationsPackages>? Packages { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientLicenseNumber { get; set; }
        public string? TransferTypeName { get; set; }
        public List<TransfersCreateExternalIncomingV2RequestItemDestinationsTransporters>? Transporters { get; set; }
    }

    public class TransfersCreateExternalIncomingV2RequestItemDestinationsPackages
    {
        public string? ExpirationDate { get; set; }
        public string? ExternalId { get; set; }
        public string? GrossUnitOfWeightName { get; set; }
        public double GrossWeight { get; set; }
        public string? ItemName { get; set; }
        public string? PackagedDate { get; set; }
        public int Quantity { get; set; }
        public string? SellByDate { get; set; }
        public string? UnitOfMeasureName { get; set; }
        public string? UseByDate { get; set; }
        public string? WholesalePrice { get; set; }
    }

    public class TransfersCreateExternalIncomingV2RequestItemDestinationsTransporters
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public bool IsLayover { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public List<TransfersCreateExternalIncomingV2RequestItemDestinationsTransportersTransporterDetails>? TransporterDetails { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersCreateExternalIncomingV2RequestItemDestinationsTransportersTransporterDetails
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersCreateTemplatesV1RequestItem
    {
        public List<TransfersCreateTemplatesV1RequestItemDestinations>? Destinations { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? Name { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersCreateTemplatesV1RequestItemDestinations
    {
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public string? InvoiceNumber { get; set; }
        public List<TransfersCreateTemplatesV1RequestItemDestinationsPackages>? Packages { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientLicenseNumber { get; set; }
        public string? TransferTypeName { get; set; }
        public List<TransfersCreateTemplatesV1RequestItemDestinationsTransporters>? Transporters { get; set; }
    }

    public class TransfersCreateTemplatesV1RequestItemDestinationsPackages
    {
        public string? GrossUnitOfWeightName { get; set; }
        public double GrossWeight { get; set; }
        public string? PackageLabel { get; set; }
        public string? WholesalePrice { get; set; }
    }

    public class TransfersCreateTemplatesV1RequestItemDestinationsTransporters
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public bool IsLayover { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? TransporterDetails { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersCreateTemplatesOutgoingV2RequestItem
    {
        public List<TransfersCreateTemplatesOutgoingV2RequestItemDestinations>? Destinations { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? Name { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersCreateTemplatesOutgoingV2RequestItemDestinations
    {
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public string? InvoiceNumber { get; set; }
        public List<TransfersCreateTemplatesOutgoingV2RequestItemDestinationsPackages>? Packages { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientLicenseNumber { get; set; }
        public string? TransferTypeName { get; set; }
        public List<TransfersCreateTemplatesOutgoingV2RequestItemDestinationsTransporters>? Transporters { get; set; }
    }

    public class TransfersCreateTemplatesOutgoingV2RequestItemDestinationsPackages
    {
        public string? GrossUnitOfWeightName { get; set; }
        public double GrossWeight { get; set; }
        public string? PackageLabel { get; set; }
        public string? WholesalePrice { get; set; }
    }

    public class TransfersCreateTemplatesOutgoingV2RequestItemDestinationsTransporters
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public bool IsLayover { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public List<TransfersCreateTemplatesOutgoingV2RequestItemDestinationsTransportersTransporterDetails>? TransporterDetails { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersCreateTemplatesOutgoingV2RequestItemDestinationsTransportersTransporterDetails
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersUpdateExternalIncomingV1RequestItem
    {
        public List<TransfersUpdateExternalIncomingV1RequestItemDestinations>? Destinations { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? ShipperAddress1 { get; set; }
        public string? ShipperAddress2 { get; set; }
        public string? ShipperAddressCity { get; set; }
        public string? ShipperAddressPostalCode { get; set; }
        public string? ShipperAddressState { get; set; }
        public string? ShipperLicenseNumber { get; set; }
        public string? ShipperMainPhoneNumber { get; set; }
        public string? ShipperName { get; set; }
        public int TransferId { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersUpdateExternalIncomingV1RequestItemDestinations
    {
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int? GrossUnitOfWeightId { get; set; }
        public double? GrossWeight { get; set; }
        public string? InvoiceNumber { get; set; }
        public List<TransfersUpdateExternalIncomingV1RequestItemDestinationsPackages>? Packages { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientLicenseNumber { get; set; }
        public int? TransferDestinationId { get; set; }
        public string? TransferTypeName { get; set; }
        public List<TransfersUpdateExternalIncomingV1RequestItemDestinationsTransporters>? Transporters { get; set; }
    }

    public class TransfersUpdateExternalIncomingV1RequestItemDestinationsPackages
    {
        public string? ExpirationDate { get; set; }
        public string? ExternalId { get; set; }
        public string? GrossUnitOfWeightName { get; set; }
        public double GrossWeight { get; set; }
        public string? ItemName { get; set; }
        public string? PackagedDate { get; set; }
        public int Quantity { get; set; }
        public string? SellByDate { get; set; }
        public string? UnitOfMeasureName { get; set; }
        public string? UseByDate { get; set; }
        public string? WholesalePrice { get; set; }
    }

    public class TransfersUpdateExternalIncomingV1RequestItemDestinationsTransporters
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public bool IsLayover { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? TransporterDetails { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersUpdateExternalIncomingV2RequestItem
    {
        public List<TransfersUpdateExternalIncomingV2RequestItemDestinations>? Destinations { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? ShipperAddress1 { get; set; }
        public string? ShipperAddress2 { get; set; }
        public string? ShipperAddressCity { get; set; }
        public string? ShipperAddressPostalCode { get; set; }
        public string? ShipperAddressState { get; set; }
        public string? ShipperLicenseNumber { get; set; }
        public string? ShipperMainPhoneNumber { get; set; }
        public string? ShipperName { get; set; }
        public int TransferId { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersUpdateExternalIncomingV2RequestItemDestinations
    {
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int? GrossUnitOfWeightId { get; set; }
        public double? GrossWeight { get; set; }
        public string? InvoiceNumber { get; set; }
        public List<TransfersUpdateExternalIncomingV2RequestItemDestinationsPackages>? Packages { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientLicenseNumber { get; set; }
        public int? TransferDestinationId { get; set; }
        public string? TransferTypeName { get; set; }
        public List<TransfersUpdateExternalIncomingV2RequestItemDestinationsTransporters>? Transporters { get; set; }
    }

    public class TransfersUpdateExternalIncomingV2RequestItemDestinationsPackages
    {
        public string? ExpirationDate { get; set; }
        public string? ExternalId { get; set; }
        public string? GrossUnitOfWeightName { get; set; }
        public double GrossWeight { get; set; }
        public string? ItemName { get; set; }
        public string? PackagedDate { get; set; }
        public int Quantity { get; set; }
        public string? SellByDate { get; set; }
        public string? UnitOfMeasureName { get; set; }
        public string? UseByDate { get; set; }
        public string? WholesalePrice { get; set; }
    }

    public class TransfersUpdateExternalIncomingV2RequestItemDestinationsTransporters
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public bool IsLayover { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public List<TransfersUpdateExternalIncomingV2RequestItemDestinationsTransportersTransporterDetails>? TransporterDetails { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersUpdateExternalIncomingV2RequestItemDestinationsTransportersTransporterDetails
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersUpdateTemplatesV1RequestItem
    {
        public List<TransfersUpdateTemplatesV1RequestItemDestinations>? Destinations { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? Name { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public int TransferTemplateId { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersUpdateTemplatesV1RequestItemDestinations
    {
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public string? InvoiceNumber { get; set; }
        public List<TransfersUpdateTemplatesV1RequestItemDestinationsPackages>? Packages { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientLicenseNumber { get; set; }
        public int TransferDestinationId { get; set; }
        public string? TransferTypeName { get; set; }
        public List<TransfersUpdateTemplatesV1RequestItemDestinationsTransporters>? Transporters { get; set; }
    }

    public class TransfersUpdateTemplatesV1RequestItemDestinationsPackages
    {
        public string? GrossUnitOfWeightName { get; set; }
        public double GrossWeight { get; set; }
        public string? PackageLabel { get; set; }
        public string? WholesalePrice { get; set; }
    }

    public class TransfersUpdateTemplatesV1RequestItemDestinationsTransporters
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public bool IsLayover { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? TransporterDetails { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersUpdateTemplatesOutgoingV2RequestItem
    {
        public List<TransfersUpdateTemplatesOutgoingV2RequestItemDestinations>? Destinations { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? Name { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public int TransferTemplateId { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersUpdateTemplatesOutgoingV2RequestItemDestinations
    {
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public string? InvoiceNumber { get; set; }
        public List<TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsPackages>? Packages { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientLicenseNumber { get; set; }
        public int TransferDestinationId { get; set; }
        public string? TransferTypeName { get; set; }
        public List<TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsTransporters>? Transporters { get; set; }
    }

    public class TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsPackages
    {
        public string? GrossUnitOfWeightName { get; set; }
        public double GrossWeight { get; set; }
        public string? PackageLabel { get; set; }
        public string? WholesalePrice { get; set; }
    }

    public class TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsTransporters
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public bool IsLayover { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public List<TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsTransportersTransporterDetails>? TransporterDetails { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsTransportersTransporterDetails
    {
        public string? DriverLayoverLeg { get; set; }
        public string? DriverLicenseNumber { get; set; }
        public string? DriverName { get; set; }
        public string? DriverOccupationalLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class PlantsCreateAdditivesV1RequestItem
    {
        public List<PlantsCreateAdditivesV1RequestItemActiveIngredients>? ActiveIngredients { get; set; }
        public string? ActualDate { get; set; }
        public string? AdditiveType { get; set; }
        public string? ApplicationDevice { get; set; }
        public string? EpaRegistrationNumber { get; set; }
        public List<string?>? PlantLabels { get; set; }
        public string? ProductSupplier { get; set; }
        public string? ProductTradeName { get; set; }
        public int TotalAmountApplied { get; set; }
        public string? TotalAmountUnitOfMeasure { get; set; }
    }

    public class PlantsCreateAdditivesV1RequestItemActiveIngredients
    {
        public string? Name { get; set; }
        public double Percentage { get; set; }
    }

    public class PlantsCreateAdditivesV2RequestItem
    {
        public List<PlantsCreateAdditivesV2RequestItemActiveIngredients>? ActiveIngredients { get; set; }
        public string? ActualDate { get; set; }
        public string? AdditiveType { get; set; }
        public string? ApplicationDevice { get; set; }
        public string? EpaRegistrationNumber { get; set; }
        public List<string?>? PlantLabels { get; set; }
        public string? ProductSupplier { get; set; }
        public string? ProductTradeName { get; set; }
        public int TotalAmountApplied { get; set; }
        public string? TotalAmountUnitOfMeasure { get; set; }
    }

    public class PlantsCreateAdditivesV2RequestItemActiveIngredients
    {
        public string? Name { get; set; }
        public double Percentage { get; set; }
    }

    public class PlantsCreateAdditivesBylocationV1RequestItem
    {
        public List<PlantsCreateAdditivesBylocationV1RequestItemActiveIngredients>? ActiveIngredients { get; set; }
        public string? ActualDate { get; set; }
        public string? AdditiveType { get; set; }
        public string? ApplicationDevice { get; set; }
        public string? EpaRegistrationNumber { get; set; }
        public string? LocationName { get; set; }
        public string? ProductSupplier { get; set; }
        public string? ProductTradeName { get; set; }
        public string? SublocationName { get; set; }
        public int TotalAmountApplied { get; set; }
        public string? TotalAmountUnitOfMeasure { get; set; }
    }

    public class PlantsCreateAdditivesBylocationV1RequestItemActiveIngredients
    {
        public string? Name { get; set; }
        public double Percentage { get; set; }
    }

    public class PlantsCreateAdditivesBylocationV2RequestItem
    {
        public List<PlantsCreateAdditivesBylocationV2RequestItemActiveIngredients>? ActiveIngredients { get; set; }
        public string? ActualDate { get; set; }
        public string? AdditiveType { get; set; }
        public string? ApplicationDevice { get; set; }
        public string? EpaRegistrationNumber { get; set; }
        public string? LocationName { get; set; }
        public string? ProductSupplier { get; set; }
        public string? ProductTradeName { get; set; }
        public string? SublocationName { get; set; }
        public int TotalAmountApplied { get; set; }
        public string? TotalAmountUnitOfMeasure { get; set; }
    }

    public class PlantsCreateAdditivesBylocationV2RequestItemActiveIngredients
    {
        public string? Name { get; set; }
        public double Percentage { get; set; }
    }

    public class PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem
    {
        public string? ActualDate { get; set; }
        public string? AdditivesTemplateName { get; set; }
        public string? LocationName { get; set; }
        public string? Rate { get; set; }
        public string? SublocationName { get; set; }
        public int TotalAmountApplied { get; set; }
        public string? TotalAmountUnitOfMeasure { get; set; }
        public string? Volume { get; set; }
    }

    public class PlantsCreateAdditivesUsingtemplateV2RequestItem
    {
        public string? ActualDate { get; set; }
        public string? AdditivesTemplateName { get; set; }
        public List<string?>? PlantLabels { get; set; }
        public string? Rate { get; set; }
        public int TotalAmountApplied { get; set; }
        public string? TotalAmountUnitOfMeasure { get; set; }
        public string? Volume { get; set; }
    }

    public class PlantsCreateChangegrowthphasesV1RequestItem
    {
        public string? GrowthDate { get; set; }
        public string? GrowthPhase { get; set; }
        public int? Id { get; set; }
        public string? Label { get; set; }
        public string? NewLocation { get; set; }
        public string? NewSublocation { get; set; }
        public string? NewTag { get; set; }
    }

    public class PlantsCreateHarvestplantsV1RequestItem
    {
        public string? ActualDate { get; set; }
        public string? DryingLocation { get; set; }
        public string? DryingSublocation { get; set; }
        public string? HarvestName { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? Plant { get; set; }
        public string? UnitOfWeight { get; set; }
        public double Weight { get; set; }
    }

    public class PlantsCreateManicureV2RequestItem
    {
        public string? ActualDate { get; set; }
        public string? DryingLocation { get; set; }
        public string? DryingSublocation { get; set; }
        public string? HarvestName { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? Plant { get; set; }
        public int PlantCount { get; set; }
        public string? UnitOfWeight { get; set; }
        public double Weight { get; set; }
    }

    public class PlantsCreateManicureplantsV1RequestItem
    {
        public string? ActualDate { get; set; }
        public string? DryingLocation { get; set; }
        public string? DryingSublocation { get; set; }
        public string? HarvestName { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? Plant { get; set; }
        public int PlantCount { get; set; }
        public string? UnitOfWeight { get; set; }
        public double Weight { get; set; }
    }

    public class PlantsCreateMoveplantsV1RequestItem
    {
        public string? ActualDate { get; set; }
        public int? Id { get; set; }
        public string? Label { get; set; }
        public string? Location { get; set; }
        public string? Sublocation { get; set; }
    }

    public class PlantsCreatePlantbatchPackageV1RequestItem
    {
        public string? ActualDate { get; set; }
        public int Count { get; set; }
        public bool IsDonation { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PackageTag { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatchType { get; set; }
        public string? PlantLabel { get; set; }
        public string? Sublocation { get; set; }
    }

    public class PlantsCreatePlantbatchPackageV2RequestItem
    {
        public string? ActualDate { get; set; }
        public int Count { get; set; }
        public bool IsDonation { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PackageTag { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatchType { get; set; }
        public string? PlantLabel { get; set; }
        public string? Sublocation { get; set; }
    }

    public class PlantsCreatePlantingsV1RequestItem
    {
        public string? ActualDate { get; set; }
        public string? LocationName { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatchName { get; set; }
        public string? PlantBatchType { get; set; }
        public int PlantCount { get; set; }
        public string? PlantLabel { get; set; }
        public string? StrainName { get; set; }
        public string? SublocationName { get; set; }
    }

    public class PlantsCreatePlantingsV2RequestItem
    {
        public string? ActualDate { get; set; }
        public string? LocationName { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatchName { get; set; }
        public string? PlantBatchType { get; set; }
        public int PlantCount { get; set; }
        public string? PlantLabel { get; set; }
        public string? StrainName { get; set; }
        public string? SublocationName { get; set; }
    }

    public class PlantsCreateWasteV1RequestItem
    {
        public string? LocationName { get; set; }
        public string? MixedMaterial { get; set; }
        public string? Note { get; set; }
        public List<object?>? PlantLabels { get; set; }
        public string? ReasonName { get; set; }
        public string? SublocationName { get; set; }
        public string? UnitOfMeasureName { get; set; }
        public string? WasteDate { get; set; }
        public string? WasteMethodName { get; set; }
        public double WasteWeight { get; set; }
    }

    public class PlantsCreateWasteV2RequestItem
    {
        public string? LocationName { get; set; }
        public string? MixedMaterial { get; set; }
        public string? Note { get; set; }
        public List<object?>? PlantLabels { get; set; }
        public string? ReasonName { get; set; }
        public string? SublocationName { get; set; }
        public string? UnitOfMeasureName { get; set; }
        public string? WasteDate { get; set; }
        public string? WasteMethodName { get; set; }
        public double WasteWeight { get; set; }
    }

    public class PlantsUpdateAdjustV2RequestItem
    {
        public int AdjustCount { get; set; }
        public string? AdjustReason { get; set; }
        public string? AdjustmentDate { get; set; }
        public int Id { get; set; }
        public string? Label { get; set; }
        public string? ReasonNote { get; set; }
    }

    public class PlantsUpdateGrowthphaseV2RequestItem
    {
        public string? GrowthDate { get; set; }
        public string? GrowthPhase { get; set; }
        public int? Id { get; set; }
        public string? Label { get; set; }
        public string? NewLocation { get; set; }
        public string? NewSublocation { get; set; }
        public string? NewTag { get; set; }
    }

    public class PlantsUpdateHarvestV2RequestItem
    {
        public string? ActualDate { get; set; }
        public string? DryingLocation { get; set; }
        public string? DryingSublocation { get; set; }
        public string? HarvestName { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? Plant { get; set; }
        public string? UnitOfWeight { get; set; }
        public double Weight { get; set; }
    }

    public class PlantsUpdateLocationV2RequestItem
    {
        public string? ActualDate { get; set; }
        public int? Id { get; set; }
        public string? Label { get; set; }
        public string? Location { get; set; }
        public string? Sublocation { get; set; }
    }

    public class PlantsUpdateMergeV2RequestItem
    {
        public string? MergeDate { get; set; }
        public string? SourcePlantGroupLabel { get; set; }
        public string? TargetPlantGroupLabel { get; set; }
    }

    public class PlantsUpdateSplitV2RequestItem
    {
        public int PlantCount { get; set; }
        public string? SourcePlantLabel { get; set; }
        public string? SplitDate { get; set; }
        public string? StrainLabel { get; set; }
        public string? TagLabel { get; set; }
    }

    public class PlantsUpdateStrainV2RequestItem
    {
        public int Id { get; set; }
        public string? Label { get; set; }
        public int StrainId { get; set; }
        public string? StrainName { get; set; }
    }

    public class PlantsUpdateTagV2RequestItem
    {
        public int? Id { get; set; }
        public string? Label { get; set; }
        public string? NewTag { get; set; }
        public string? ReplaceDate { get; set; }
        public int? TagId { get; set; }
    }

    public class AdditivesTemplatesCreateV2RequestItem
    {
        public List<AdditivesTemplatesCreateV2RequestItemActiveIngredients>? ActiveIngredients { get; set; }
        public string? AdditiveType { get; set; }
        public string? ApplicationDevice { get; set; }
        public string? EpaRegistrationNumber { get; set; }
        public string? Name { get; set; }
        public string? Note { get; set; }
        public string? ProductSupplier { get; set; }
        public string? ProductTradeName { get; set; }
        public string? RestrictiveEntryIntervalQuantityDescription { get; set; }
        public string? RestrictiveEntryIntervalTimeDescription { get; set; }
    }

    public class AdditivesTemplatesCreateV2RequestItemActiveIngredients
    {
        public string? Name { get; set; }
        public double Percentage { get; set; }
    }

    public class AdditivesTemplatesUpdateV2RequestItem
    {
        public List<AdditivesTemplatesUpdateV2RequestItemActiveIngredients>? ActiveIngredients { get; set; }
        public string? AdditiveType { get; set; }
        public string? ApplicationDevice { get; set; }
        public string? EpaRegistrationNumber { get; set; }
        public int Id { get; set; }
        public string? Name { get; set; }
        public string? Note { get; set; }
        public string? ProductSupplier { get; set; }
        public string? ProductTradeName { get; set; }
        public string? RestrictiveEntryIntervalQuantityDescription { get; set; }
        public string? RestrictiveEntryIntervalTimeDescription { get; set; }
    }

    public class AdditivesTemplatesUpdateV2RequestItemActiveIngredients
    {
        public string? Name { get; set; }
        public double Percentage { get; set; }
    }

    public class HarvestsCreateFinishV1RequestItem
    {
        public string? ActualDate { get; set; }
        public int Id { get; set; }
    }

    public class HarvestsCreatePackageV1RequestItem
    {
        public string? ActualDate { get; set; }
        public bool DecontaminateProduct { get; set; }
        public string? DecontaminationDate { get; set; }
        public string? DecontaminationSteps { get; set; }
        public string? ExpirationDate { get; set; }
        public List<HarvestsCreatePackageV1RequestItemIngredients>? Ingredients { get; set; }
        public bool IsDonation { get; set; }
        public bool IsProductionBatch { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public int LabTestStageId { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int ProcessingJobTypeId { get; set; }
        public bool ProductRequiresDecontamination { get; set; }
        public bool ProductRequiresRemediation { get; set; }
        public string? ProductionBatchNumber { get; set; }
        public bool RemediateProduct { get; set; }
        public string? RemediationDate { get; set; }
        public int? RemediationMethodId { get; set; }
        public string? RemediationSteps { get; set; }
        public List<object?>? RequiredLabTestBatches { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UnitOfWeight { get; set; }
        public string? UseByDate { get; set; }
    }

    public class HarvestsCreatePackageV1RequestItemIngredients
    {
        public int HarvestId { get; set; }
        public string? HarvestName { get; set; }
        public string? UnitOfWeight { get; set; }
        public double Weight { get; set; }
    }

    public class HarvestsCreatePackageV2RequestItem
    {
        public string? ActualDate { get; set; }
        public bool DecontaminateProduct { get; set; }
        public string? DecontaminationDate { get; set; }
        public string? DecontaminationSteps { get; set; }
        public string? ExpirationDate { get; set; }
        public List<HarvestsCreatePackageV2RequestItemIngredients>? Ingredients { get; set; }
        public bool IsDonation { get; set; }
        public bool IsProductionBatch { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public int LabTestStageId { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int ProcessingJobTypeId { get; set; }
        public bool ProductRequiresDecontamination { get; set; }
        public bool ProductRequiresRemediation { get; set; }
        public string? ProductionBatchNumber { get; set; }
        public bool RemediateProduct { get; set; }
        public string? RemediationDate { get; set; }
        public int? RemediationMethodId { get; set; }
        public string? RemediationSteps { get; set; }
        public List<object?>? RequiredLabTestBatches { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UnitOfWeight { get; set; }
        public string? UseByDate { get; set; }
    }

    public class HarvestsCreatePackageV2RequestItemIngredients
    {
        public int HarvestId { get; set; }
        public string? HarvestName { get; set; }
        public string? UnitOfWeight { get; set; }
        public double Weight { get; set; }
    }

    public class HarvestsCreatePackageTestingV1RequestItem
    {
        public string? ActualDate { get; set; }
        public bool DecontaminateProduct { get; set; }
        public string? DecontaminationDate { get; set; }
        public string? DecontaminationSteps { get; set; }
        public string? ExpirationDate { get; set; }
        public List<HarvestsCreatePackageTestingV1RequestItemIngredients>? Ingredients { get; set; }
        public bool IsDonation { get; set; }
        public bool IsProductionBatch { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public int LabTestStageId { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int ProcessingJobTypeId { get; set; }
        public bool ProductRequiresDecontamination { get; set; }
        public bool ProductRequiresRemediation { get; set; }
        public string? ProductionBatchNumber { get; set; }
        public bool RemediateProduct { get; set; }
        public string? RemediationDate { get; set; }
        public int? RemediationMethodId { get; set; }
        public string? RemediationSteps { get; set; }
        public List<object?>? RequiredLabTestBatches { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UnitOfWeight { get; set; }
        public string? UseByDate { get; set; }
    }

    public class HarvestsCreatePackageTestingV1RequestItemIngredients
    {
        public int HarvestId { get; set; }
        public string? HarvestName { get; set; }
        public string? UnitOfWeight { get; set; }
        public double Weight { get; set; }
    }

    public class HarvestsCreatePackageTestingV2RequestItem
    {
        public string? ActualDate { get; set; }
        public bool DecontaminateProduct { get; set; }
        public string? DecontaminationDate { get; set; }
        public string? DecontaminationSteps { get; set; }
        public string? ExpirationDate { get; set; }
        public List<HarvestsCreatePackageTestingV2RequestItemIngredients>? Ingredients { get; set; }
        public bool IsDonation { get; set; }
        public bool IsProductionBatch { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public int LabTestStageId { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int ProcessingJobTypeId { get; set; }
        public bool ProductRequiresDecontamination { get; set; }
        public bool ProductRequiresRemediation { get; set; }
        public string? ProductionBatchNumber { get; set; }
        public bool RemediateProduct { get; set; }
        public string? RemediationDate { get; set; }
        public int? RemediationMethodId { get; set; }
        public string? RemediationSteps { get; set; }
        public List<object?>? RequiredLabTestBatches { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UnitOfWeight { get; set; }
        public string? UseByDate { get; set; }
    }

    public class HarvestsCreatePackageTestingV2RequestItemIngredients
    {
        public int HarvestId { get; set; }
        public string? HarvestName { get; set; }
        public string? UnitOfWeight { get; set; }
        public double Weight { get; set; }
    }

    public class HarvestsCreateRemoveWasteV1RequestItem
    {
        public string? ActualDate { get; set; }
        public int Id { get; set; }
        public string? UnitOfWeight { get; set; }
        public string? WasteType { get; set; }
        public double WasteWeight { get; set; }
    }

    public class HarvestsCreateUnfinishV1RequestItem
    {
        public int Id { get; set; }
    }

    public class HarvestsCreateWasteV2RequestItem
    {
        public string? ActualDate { get; set; }
        public int Id { get; set; }
        public string? UnitOfWeight { get; set; }
        public string? WasteType { get; set; }
        public double WasteWeight { get; set; }
    }

    public class HarvestsUpdateFinishV2RequestItem
    {
        public string? ActualDate { get; set; }
        public int Id { get; set; }
    }

    public class HarvestsUpdateLocationV2RequestItem
    {
        public string? ActualDate { get; set; }
        public string? DryingLocation { get; set; }
        public string? DryingSublocation { get; set; }
        public string? HarvestName { get; set; }
        public int Id { get; set; }
    }

    public class HarvestsUpdateMoveV1RequestItem
    {
        public string? ActualDate { get; set; }
        public string? DryingLocation { get; set; }
        public string? DryingSublocation { get; set; }
        public string? HarvestName { get; set; }
        public int Id { get; set; }
    }

    public class HarvestsUpdateRenameV1RequestItem
    {
        public int Id { get; set; }
        public string? NewName { get; set; }
        public string? OldName { get; set; }
    }

    public class HarvestsUpdateRenameV2RequestItem
    {
        public int Id { get; set; }
        public string? NewName { get; set; }
        public string? OldName { get; set; }
    }

    public class HarvestsUpdateRestoreHarvestedPlantsV2RequestItem
    {
        public int HarvestId { get; set; }
        public List<string?>? PlantTags { get; set; }
    }

    public class HarvestsUpdateUnfinishV2RequestItem
    {
        public int Id { get; set; }
    }

    public class LabTestsCreateRecordV1RequestItem
    {
        public string? DocumentFileBase64 { get; set; }
        public string? DocumentFileName { get; set; }
        public string? Label { get; set; }
        public string? ResultDate { get; set; }
        public List<LabTestsCreateRecordV1RequestItemResults>? Results { get; set; }
    }

    public class LabTestsCreateRecordV1RequestItemResults
    {
        public string? LabTestTypeName { get; set; }
        public string? Notes { get; set; }
        public bool Passed { get; set; }
        public double Quantity { get; set; }
    }

    public class LabTestsCreateRecordV2RequestItem
    {
        public string? DocumentFileBase64 { get; set; }
        public string? DocumentFileName { get; set; }
        public string? Label { get; set; }
        public string? ResultDate { get; set; }
        public List<LabTestsCreateRecordV2RequestItemResults>? Results { get; set; }
    }

    public class LabTestsCreateRecordV2RequestItemResults
    {
        public string? LabTestTypeName { get; set; }
        public string? Notes { get; set; }
        public bool Passed { get; set; }
        public double Quantity { get; set; }
    }

    public class LabTestsUpdateLabtestdocumentV1RequestItem
    {
        public string? DocumentFileBase64 { get; set; }
        public string? DocumentFileName { get; set; }
        public int LabTestResultId { get; set; }
    }

    public class LabTestsUpdateLabtestdocumentV2RequestItem
    {
        public string? DocumentFileBase64 { get; set; }
        public string? DocumentFileName { get; set; }
        public int LabTestResultId { get; set; }
    }

    public class LabTestsUpdateResultReleaseV1RequestItem
    {
        public string? PackageLabel { get; set; }
    }

    public class LabTestsUpdateResultReleaseV2RequestItem
    {
        public string? PackageLabel { get; set; }
    }

    public class SalesCreateDeliveryV1RequestItem
    {
        public int? ConsumerId { get; set; }
        public string? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientAddressCity { get; set; }
        public string? RecipientAddressCounty { get; set; }
        public string? RecipientAddressPostalCode { get; set; }
        public string? RecipientAddressState { get; set; }
        public string? RecipientAddressStreet1 { get; set; }
        public string? RecipientAddressStreet2 { get; set; }
        public string? RecipientName { get; set; }
        public int? RecipientZoneId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesCreateDeliveryV1RequestItemTransactions>? Transactions { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesCreateDeliveryV1RequestItemTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesCreateDeliveryV2RequestItem
    {
        public int? ConsumerId { get; set; }
        public string? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientAddressCity { get; set; }
        public string? RecipientAddressCounty { get; set; }
        public string? RecipientAddressPostalCode { get; set; }
        public string? RecipientAddressState { get; set; }
        public string? RecipientAddressStreet1 { get; set; }
        public string? RecipientAddressStreet2 { get; set; }
        public string? RecipientName { get; set; }
        public int? RecipientZoneId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesCreateDeliveryV2RequestItemTransactions>? Transactions { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesCreateDeliveryV2RequestItemTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesCreateDeliveryRetailerV1RequestItem
    {
        public string? DateTime { get; set; }
        public List<SalesCreateDeliveryRetailerV1RequestItemDestinations>? Destinations { get; set; }
        public string? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public List<SalesCreateDeliveryRetailerV1RequestItemPackages>? Packages { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesCreateDeliveryRetailerV1RequestItemDestinations
    {
        public string? ConsumerId { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? RecipientAddressCity { get; set; }
        public string? RecipientAddressCounty { get; set; }
        public string? RecipientAddressPostalCode { get; set; }
        public string? RecipientAddressState { get; set; }
        public string? RecipientAddressStreet1 { get; set; }
        public string? RecipientAddressStreet2 { get; set; }
        public string? RecipientName { get; set; }
        public string? RecipientZoneId { get; set; }
        public string? SalesCustomerType { get; set; }
        public List<SalesCreateDeliveryRetailerV1RequestItemDestinationsTransactions>? Transactions { get; set; }
    }

    public class SalesCreateDeliveryRetailerV1RequestItemDestinationsTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesCreateDeliveryRetailerV1RequestItemPackages
    {
        public string? DateTime { get; set; }
        public string? PackageLabel { get; set; }
        public int Quantity { get; set; }
        public double TotalPrice { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class SalesCreateDeliveryRetailerV2RequestItem
    {
        public string? DateTime { get; set; }
        public List<SalesCreateDeliveryRetailerV2RequestItemDestinations>? Destinations { get; set; }
        public string? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public List<SalesCreateDeliveryRetailerV2RequestItemPackages>? Packages { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesCreateDeliveryRetailerV2RequestItemDestinations
    {
        public string? ConsumerId { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? RecipientAddressCity { get; set; }
        public string? RecipientAddressCounty { get; set; }
        public string? RecipientAddressPostalCode { get; set; }
        public string? RecipientAddressState { get; set; }
        public string? RecipientAddressStreet1 { get; set; }
        public string? RecipientAddressStreet2 { get; set; }
        public string? RecipientName { get; set; }
        public string? RecipientZoneId { get; set; }
        public string? SalesCustomerType { get; set; }
        public List<SalesCreateDeliveryRetailerV2RequestItemDestinationsTransactions>? Transactions { get; set; }
    }

    public class SalesCreateDeliveryRetailerV2RequestItemDestinationsTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesCreateDeliveryRetailerV2RequestItemPackages
    {
        public string? DateTime { get; set; }
        public string? PackageLabel { get; set; }
        public int Quantity { get; set; }
        public double TotalPrice { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class SalesCreateDeliveryRetailerDepartV1RequestItem
    {
        public int RetailerDeliveryId { get; set; }
    }

    public class SalesCreateDeliveryRetailerDepartV2RequestItem
    {
        public int RetailerDeliveryId { get; set; }
    }

    public class SalesCreateDeliveryRetailerEndV1RequestItem
    {
        public string? ActualArrivalDateTime { get; set; }
        public List<SalesCreateDeliveryRetailerEndV1RequestItemPackages>? Packages { get; set; }
        public int RetailerDeliveryId { get; set; }
    }

    public class SalesCreateDeliveryRetailerEndV1RequestItemPackages
    {
        public int EndQuantity { get; set; }
        public string? EndUnitOfMeasure { get; set; }
        public string? Label { get; set; }
    }

    public class SalesCreateDeliveryRetailerEndV2RequestItem
    {
        public string? ActualArrivalDateTime { get; set; }
        public List<SalesCreateDeliveryRetailerEndV2RequestItemPackages>? Packages { get; set; }
        public int RetailerDeliveryId { get; set; }
    }

    public class SalesCreateDeliveryRetailerEndV2RequestItemPackages
    {
        public int EndQuantity { get; set; }
        public string? EndUnitOfMeasure { get; set; }
        public string? Label { get; set; }
    }

    public class SalesCreateDeliveryRetailerRestockV1RequestItem
    {
        public string? DateTime { get; set; }
        public string? Destinations { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public List<SalesCreateDeliveryRetailerRestockV1RequestItemPackages>? Packages { get; set; }
        public int RetailerDeliveryId { get; set; }
    }

    public class SalesCreateDeliveryRetailerRestockV1RequestItemPackages
    {
        public string? PackageLabel { get; set; }
        public int Quantity { get; set; }
        public bool RemoveCurrentPackage { get; set; }
        public double TotalPrice { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class SalesCreateDeliveryRetailerRestockV2RequestItem
    {
        public string? DateTime { get; set; }
        public string? Destinations { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public List<SalesCreateDeliveryRetailerRestockV2RequestItemPackages>? Packages { get; set; }
        public int RetailerDeliveryId { get; set; }
    }

    public class SalesCreateDeliveryRetailerRestockV2RequestItemPackages
    {
        public string? PackageLabel { get; set; }
        public int Quantity { get; set; }
        public bool RemoveCurrentPackage { get; set; }
        public double TotalPrice { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class SalesCreateDeliveryRetailerSaleV1RequestItem
    {
        public int? ConsumerId { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientAddressCity { get; set; }
        public string? RecipientAddressCounty { get; set; }
        public string? RecipientAddressPostalCode { get; set; }
        public string? RecipientAddressState { get; set; }
        public string? RecipientAddressStreet1 { get; set; }
        public string? RecipientAddressStreet2 { get; set; }
        public string? RecipientName { get; set; }
        public int? RecipientZoneId { get; set; }
        public int RetailerDeliveryId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesCreateDeliveryRetailerSaleV1RequestItemTransactions>? Transactions { get; set; }
    }

    public class SalesCreateDeliveryRetailerSaleV1RequestItemTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesCreateDeliveryRetailerSaleV2RequestItem
    {
        public int? ConsumerId { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientAddressCity { get; set; }
        public string? RecipientAddressCounty { get; set; }
        public string? RecipientAddressPostalCode { get; set; }
        public string? RecipientAddressState { get; set; }
        public string? RecipientAddressStreet1 { get; set; }
        public string? RecipientAddressStreet2 { get; set; }
        public string? RecipientName { get; set; }
        public int? RecipientZoneId { get; set; }
        public int RetailerDeliveryId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesCreateDeliveryRetailerSaleV2RequestItemTransactions>? Transactions { get; set; }
    }

    public class SalesCreateDeliveryRetailerSaleV2RequestItemTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesCreateReceiptV1RequestItem
    {
        public string? CaregiverLicenseNumber { get; set; }
        public string? ExternalReceiptNumber { get; set; }
        public string? IdentificationMethod { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int? PatientRegistrationLocationId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesCreateReceiptV1RequestItemTransactions>? Transactions { get; set; }
    }

    public class SalesCreateReceiptV1RequestItemTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesCreateReceiptV2RequestItem
    {
        public string? CaregiverLicenseNumber { get; set; }
        public string? ExternalReceiptNumber { get; set; }
        public string? IdentificationMethod { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int? PatientRegistrationLocationId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesCreateReceiptV2RequestItemTransactions>? Transactions { get; set; }
    }

    public class SalesCreateReceiptV2RequestItemTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesCreateTransactionByDateV1RequestItem
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesUpdateDeliveryV1RequestItem
    {
        public int? ConsumerId { get; set; }
        public string? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int Id { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientAddressCity { get; set; }
        public string? RecipientAddressCounty { get; set; }
        public string? RecipientAddressPostalCode { get; set; }
        public string? RecipientAddressState { get; set; }
        public string? RecipientAddressStreet1 { get; set; }
        public string? RecipientAddressStreet2 { get; set; }
        public string? RecipientName { get; set; }
        public string? RecipientZoneId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesUpdateDeliveryV1RequestItemTransactions>? Transactions { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesUpdateDeliveryV1RequestItemTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesUpdateDeliveryV2RequestItem
    {
        public int? ConsumerId { get; set; }
        public string? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int Id { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientAddressCity { get; set; }
        public string? RecipientAddressCounty { get; set; }
        public string? RecipientAddressPostalCode { get; set; }
        public string? RecipientAddressState { get; set; }
        public string? RecipientAddressStreet1 { get; set; }
        public string? RecipientAddressStreet2 { get; set; }
        public string? RecipientName { get; set; }
        public string? RecipientZoneId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesUpdateDeliveryV2RequestItemTransactions>? Transactions { get; set; }
        public string? TransporterFacilityLicenseNumber { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesUpdateDeliveryV2RequestItemTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesUpdateDeliveryCompleteV1RequestItem
    {
        public List<string?>? AcceptedPackages { get; set; }
        public string? ActualArrivalDateTime { get; set; }
        public int Id { get; set; }
        public string? PaymentType { get; set; }
        public List<SalesUpdateDeliveryCompleteV1RequestItemReturnedPackages>? ReturnedPackages { get; set; }
    }

    public class SalesUpdateDeliveryCompleteV1RequestItemReturnedPackages
    {
        public string? Label { get; set; }
        public int ReturnQuantityVerified { get; set; }
        public string? ReturnReason { get; set; }
        public string? ReturnReasonNote { get; set; }
        public string? ReturnUnitOfMeasure { get; set; }
    }

    public class SalesUpdateDeliveryCompleteV2RequestItem
    {
        public List<string?>? AcceptedPackages { get; set; }
        public string? ActualArrivalDateTime { get; set; }
        public int Id { get; set; }
        public string? PaymentType { get; set; }
        public List<SalesUpdateDeliveryCompleteV2RequestItemReturnedPackages>? ReturnedPackages { get; set; }
    }

    public class SalesUpdateDeliveryCompleteV2RequestItemReturnedPackages
    {
        public string? Label { get; set; }
        public int ReturnQuantityVerified { get; set; }
        public string? ReturnReason { get; set; }
        public string? ReturnReasonNote { get; set; }
        public string? ReturnUnitOfMeasure { get; set; }
    }

    public class SalesUpdateDeliveryHubV1RequestItem
    {
        public string? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int Id { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? PlannedRoute { get; set; }
        public string? TransporterFacilityId { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesUpdateDeliveryHubV2RequestItem
    {
        public string? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int Id { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? PlannedRoute { get; set; }
        public string? TransporterFacilityId { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesUpdateDeliveryHubAcceptV1RequestItem
    {
        public int Id { get; set; }
    }

    public class SalesUpdateDeliveryHubAcceptV2RequestItem
    {
        public int Id { get; set; }
    }

    public class SalesUpdateDeliveryHubDepartV1RequestItem
    {
        public int Id { get; set; }
    }

    public class SalesUpdateDeliveryHubDepartV2RequestItem
    {
        public int Id { get; set; }
    }

    public class SalesUpdateDeliveryHubVerifyIdV1RequestItem
    {
        public int Id { get; set; }
        public string? PaymentType { get; set; }
    }

    public class SalesUpdateDeliveryHubVerifyIdV2RequestItem
    {
        public int Id { get; set; }
        public string? PaymentType { get; set; }
    }

    public class SalesUpdateDeliveryRetailerV1RequestItem
    {
        public string? DateTime { get; set; }
        public List<SalesUpdateDeliveryRetailerV1RequestItemDestinations>? Destinations { get; set; }
        public string? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int Id { get; set; }
        public List<SalesUpdateDeliveryRetailerV1RequestItemPackages>? Packages { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesUpdateDeliveryRetailerV1RequestItemDestinations
    {
        public string? ConsumerId { get; set; }
        public int? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int Id { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientAddressCity { get; set; }
        public string? RecipientAddressCounty { get; set; }
        public string? RecipientAddressPostalCode { get; set; }
        public string? RecipientAddressState { get; set; }
        public string? RecipientAddressStreet1 { get; set; }
        public string? RecipientAddressStreet2 { get; set; }
        public string? RecipientName { get; set; }
        public string? RecipientZoneId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesUpdateDeliveryRetailerV1RequestItemDestinationsTransactions>? Transactions { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesUpdateDeliveryRetailerV1RequestItemDestinationsTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesUpdateDeliveryRetailerV1RequestItemPackages
    {
        public string? DateTime { get; set; }
        public string? PackageLabel { get; set; }
        public int Quantity { get; set; }
        public double TotalPrice { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class SalesUpdateDeliveryRetailerV2RequestItem
    {
        public string? DateTime { get; set; }
        public List<SalesUpdateDeliveryRetailerV2RequestItemDestinations>? Destinations { get; set; }
        public string? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int Id { get; set; }
        public List<SalesUpdateDeliveryRetailerV2RequestItemPackages>? Packages { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesUpdateDeliveryRetailerV2RequestItemDestinations
    {
        public string? ConsumerId { get; set; }
        public int? DriverEmployeeId { get; set; }
        public string? DriverName { get; set; }
        public string? DriversLicenseNumber { get; set; }
        public string? EstimatedArrivalDateTime { get; set; }
        public string? EstimatedDepartureDateTime { get; set; }
        public int Id { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PhoneNumberForQuestions { get; set; }
        public string? PlannedRoute { get; set; }
        public string? RecipientAddressCity { get; set; }
        public string? RecipientAddressCounty { get; set; }
        public string? RecipientAddressPostalCode { get; set; }
        public string? RecipientAddressState { get; set; }
        public string? RecipientAddressStreet1 { get; set; }
        public string? RecipientAddressStreet2 { get; set; }
        public string? RecipientName { get; set; }
        public string? RecipientZoneId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesUpdateDeliveryRetailerV2RequestItemDestinationsTransactions>? Transactions { get; set; }
        public string? VehicleLicensePlateNumber { get; set; }
        public string? VehicleMake { get; set; }
        public string? VehicleModel { get; set; }
    }

    public class SalesUpdateDeliveryRetailerV2RequestItemDestinationsTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesUpdateDeliveryRetailerV2RequestItemPackages
    {
        public string? DateTime { get; set; }
        public string? PackageLabel { get; set; }
        public int Quantity { get; set; }
        public double TotalPrice { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class SalesUpdateReceiptV1RequestItem
    {
        public string? CaregiverLicenseNumber { get; set; }
        public string? ExternalReceiptNumber { get; set; }
        public int Id { get; set; }
        public string? IdentificationMethod { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int? PatientRegistrationLocationId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesUpdateReceiptV1RequestItemTransactions>? Transactions { get; set; }
    }

    public class SalesUpdateReceiptV1RequestItemTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesUpdateReceiptV2RequestItem
    {
        public string? CaregiverLicenseNumber { get; set; }
        public string? ExternalReceiptNumber { get; set; }
        public int Id { get; set; }
        public string? IdentificationMethod { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int? PatientRegistrationLocationId { get; set; }
        public string? SalesCustomerType { get; set; }
        public string? SalesDateTime { get; set; }
        public List<SalesUpdateReceiptV2RequestItemTransactions>? Transactions { get; set; }
    }

    public class SalesUpdateReceiptV2RequestItemTransactions
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class SalesUpdateReceiptFinalizeV2RequestItem
    {
        public int Id { get; set; }
    }

    public class SalesUpdateReceiptUnfinalizeV2RequestItem
    {
        public int Id { get; set; }
    }

    public class SalesUpdateTransactionByDateV1RequestItem
    {
        public string? CityTax { get; set; }
        public string? CountyTax { get; set; }
        public string? DiscountAmount { get; set; }
        public string? ExciseTax { get; set; }
        public string? InvoiceNumber { get; set; }
        public string? MunicipalTax { get; set; }
        public string? PackageLabel { get; set; }
        public string? Price { get; set; }
        public string? QrCodes { get; set; }
        public int Quantity { get; set; }
        public string? SalesTax { get; set; }
        public string? SubTotal { get; set; }
        public double TotalAmount { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcContent { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class TransportersCreateDriverV2RequestItem
    {
        public string? DriversLicenseNumber { get; set; }
        public string? EmployeeId { get; set; }
        public string? Name { get; set; }
    }

    public class TransportersCreateVehicleV2RequestItem
    {
        public string? LicensePlateNumber { get; set; }
        public string? Make { get; set; }
        public string? Model { get; set; }
    }

    public class TransportersUpdateDriverV2RequestItem
    {
        public string? DriversLicenseNumber { get; set; }
        public string? EmployeeId { get; set; }
        public string? Id { get; set; }
        public string? Name { get; set; }
    }

    public class TransportersUpdateVehicleV2RequestItem
    {
        public string? Id { get; set; }
        public string? LicensePlateNumber { get; set; }
        public string? Make { get; set; }
        public string? Model { get; set; }
    }

    public class PackagesCreateV1RequestItem
    {
        public string? ActualDate { get; set; }
        public string? ExpirationDate { get; set; }
        public List<PackagesCreateV1RequestItemIngredients>? Ingredients { get; set; }
        public bool IsDonation { get; set; }
        public bool IsProductionBatch { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public int? LabTestStageId { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int? ProcessingJobTypeId { get; set; }
        public bool ProductRequiresRemediation { get; set; }
        public string? ProductionBatchNumber { get; set; }
        public int Quantity { get; set; }
        public bool? RequiredLabTestBatches { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UnitOfMeasure { get; set; }
        public string? UseByDate { get; set; }
        public bool UseSameItem { get; set; }
    }

    public class PackagesCreateV1RequestItemIngredients
    {
        public string? Package { get; set; }
        public int Quantity { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class PackagesCreateV2RequestItem
    {
        public string? ActualDate { get; set; }
        public string? ExpirationDate { get; set; }
        public List<PackagesCreateV2RequestItemIngredients>? Ingredients { get; set; }
        public bool IsDonation { get; set; }
        public bool IsProductionBatch { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public int? LabTestStageId { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int? ProcessingJobTypeId { get; set; }
        public bool ProductRequiresRemediation { get; set; }
        public string? ProductionBatchNumber { get; set; }
        public int Quantity { get; set; }
        public bool? RequiredLabTestBatches { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UnitOfMeasure { get; set; }
        public string? UseByDate { get; set; }
        public bool UseSameItem { get; set; }
    }

    public class PackagesCreateV2RequestItemIngredients
    {
        public string? Package { get; set; }
        public int Quantity { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class PackagesCreateAdjustV1RequestItem
    {
        public string? AdjustmentDate { get; set; }
        public string? AdjustmentReason { get; set; }
        public string? Label { get; set; }
        public int Quantity { get; set; }
        public string? ReasonNote { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class PackagesCreateAdjustV2RequestItem
    {
        public string? AdjustmentDate { get; set; }
        public string? AdjustmentReason { get; set; }
        public string? Label { get; set; }
        public int Quantity { get; set; }
        public string? ReasonNote { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class PackagesCreateChangeItemV1RequestItem
    {
        public string? Item { get; set; }
        public string? Label { get; set; }
    }

    public class PackagesCreateChangeLocationV1RequestItem
    {
        public string? Label { get; set; }
        public string? Location { get; set; }
        public string? MoveDate { get; set; }
        public string? Sublocation { get; set; }
    }

    public class PackagesCreateFinishV1RequestItem
    {
        public string? ActualDate { get; set; }
        public string? Label { get; set; }
    }

    public class PackagesCreatePlantingsV1RequestItem
    {
        public string? LocationName { get; set; }
        public int PackageAdjustmentAmount { get; set; }
        public string? PackageAdjustmentUnitOfMeasureName { get; set; }
        public string? PackageLabel { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatchName { get; set; }
        public string? PlantBatchType { get; set; }
        public int PlantCount { get; set; }
        public string? PlantedDate { get; set; }
        public string? StrainName { get; set; }
        public string? SublocationName { get; set; }
        public string? UnpackagedDate { get; set; }
    }

    public class PackagesCreatePlantingsV2RequestItem
    {
        public string? LocationName { get; set; }
        public int PackageAdjustmentAmount { get; set; }
        public string? PackageAdjustmentUnitOfMeasureName { get; set; }
        public string? PackageLabel { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? PlantBatchName { get; set; }
        public string? PlantBatchType { get; set; }
        public int PlantCount { get; set; }
        public string? PlantedDate { get; set; }
        public string? StrainName { get; set; }
        public string? SublocationName { get; set; }
        public string? UnpackagedDate { get; set; }
    }

    public class PackagesCreateRemediateV1RequestItem
    {
        public string? PackageLabel { get; set; }
        public string? RemediationDate { get; set; }
        public string? RemediationMethodName { get; set; }
        public string? RemediationSteps { get; set; }
    }

    public class PackagesCreateTestingV1RequestItem
    {
        public string? ActualDate { get; set; }
        public string? ExpirationDate { get; set; }
        public List<PackagesCreateTestingV1RequestItemIngredients>? Ingredients { get; set; }
        public bool IsDonation { get; set; }
        public bool IsProductionBatch { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public int? LabTestStageId { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int? ProcessingJobTypeId { get; set; }
        public bool ProductRequiresRemediation { get; set; }
        public string? ProductionBatchNumber { get; set; }
        public int Quantity { get; set; }
        public bool? RequiredLabTestBatches { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UnitOfMeasure { get; set; }
        public string? UseByDate { get; set; }
        public bool UseSameItem { get; set; }
    }

    public class PackagesCreateTestingV1RequestItemIngredients
    {
        public string? Package { get; set; }
        public int Quantity { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class PackagesCreateTestingV2RequestItem
    {
        public string? ActualDate { get; set; }
        public string? ExpirationDate { get; set; }
        public List<PackagesCreateTestingV2RequestItemIngredients>? Ingredients { get; set; }
        public bool IsDonation { get; set; }
        public bool IsProductionBatch { get; set; }
        public bool IsTradeSample { get; set; }
        public string? Item { get; set; }
        public int? LabTestStageId { get; set; }
        public string? Location { get; set; }
        public string? Note { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public int? ProcessingJobTypeId { get; set; }
        public bool ProductRequiresRemediation { get; set; }
        public string? ProductionBatchNumber { get; set; }
        public int Quantity { get; set; }
        public List<string?>? RequiredLabTestBatches { get; set; }
        public string? SellByDate { get; set; }
        public string? Sublocation { get; set; }
        public string? Tag { get; set; }
        public string? UnitOfMeasure { get; set; }
        public string? UseByDate { get; set; }
        public bool UseSameItem { get; set; }
    }

    public class PackagesCreateTestingV2RequestItemIngredients
    {
        public string? Package { get; set; }
        public int Quantity { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class PackagesCreateUnfinishV1RequestItem
    {
        public string? Label { get; set; }
    }

    public class PackagesUpdateAdjustV2RequestItem
    {
        public string? AdjustmentDate { get; set; }
        public string? AdjustmentReason { get; set; }
        public string? Label { get; set; }
        public int Quantity { get; set; }
        public string? ReasonNote { get; set; }
        public string? UnitOfMeasure { get; set; }
    }

    public class PackagesUpdateChangeNoteV1RequestItem
    {
        public string? Note { get; set; }
        public string? PackageLabel { get; set; }
    }

    public class PackagesUpdateDecontaminateV2RequestItem
    {
        public string? DecontaminationDate { get; set; }
        public string? DecontaminationMethodName { get; set; }
        public string? DecontaminationSteps { get; set; }
        public string? PackageLabel { get; set; }
    }

    public class PackagesUpdateDonationFlagV2RequestItem
    {
        public string? Label { get; set; }
    }

    public class PackagesUpdateDonationUnflagV2RequestItem
    {
        public string? Label { get; set; }
    }

    public class PackagesUpdateExternalidV2RequestItem
    {
        public string? ExternalId { get; set; }
        public string? PackageLabel { get; set; }
    }

    public class PackagesUpdateFinishV2RequestItem
    {
        public string? ActualDate { get; set; }
        public string? Label { get; set; }
    }

    public class PackagesUpdateItemV2RequestItem
    {
        public string? Item { get; set; }
        public string? Label { get; set; }
    }

    public class PackagesUpdateLabTestRequiredV2RequestItem
    {
        public string? Label { get; set; }
        public List<string?>? RequiredLabTestBatches { get; set; }
    }

    public class PackagesUpdateLocationV2RequestItem
    {
        public string? Label { get; set; }
        public string? Location { get; set; }
        public string? MoveDate { get; set; }
        public string? Sublocation { get; set; }
    }

    public class PackagesUpdateNoteV2RequestItem
    {
        public string? Note { get; set; }
        public string? PackageLabel { get; set; }
    }

    public class PackagesUpdateRemediateV2RequestItem
    {
        public string? PackageLabel { get; set; }
        public string? RemediationDate { get; set; }
        public string? RemediationMethodName { get; set; }
        public string? RemediationSteps { get; set; }
    }

    public class PackagesUpdateTradesampleFlagV2RequestItem
    {
        public string? Label { get; set; }
    }

    public class PackagesUpdateTradesampleUnflagV2RequestItem
    {
        public string? Label { get; set; }
    }

    public class PackagesUpdateUnfinishV2RequestItem
    {
        public string? Label { get; set; }
    }

    public class PackagesUpdateUsebydateV2RequestItem
    {
        public string? ExpirationDate { get; set; }
        public string? Label { get; set; }
        public string? SellByDate { get; set; }
        public string? UseByDate { get; set; }
    }

    public class PatientCheckInsCreateV1RequestItem
    {
        public string? CheckInDate { get; set; }
        public int CheckInLocationId { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? RegistrationExpiryDate { get; set; }
        public string? RegistrationStartDate { get; set; }
    }

    public class PatientCheckInsCreateV2RequestItem
    {
        public string? CheckInDate { get; set; }
        public int CheckInLocationId { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? RegistrationExpiryDate { get; set; }
        public string? RegistrationStartDate { get; set; }
    }

    public class PatientCheckInsUpdateV1RequestItem
    {
        public string? CheckInDate { get; set; }
        public int CheckInLocationId { get; set; }
        public int Id { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? RegistrationExpiryDate { get; set; }
        public string? RegistrationStartDate { get; set; }
    }

    public class PatientCheckInsUpdateV2RequestItem
    {
        public string? CheckInDate { get; set; }
        public int CheckInLocationId { get; set; }
        public int Id { get; set; }
        public string? PatientLicenseNumber { get; set; }
        public string? RegistrationExpiryDate { get; set; }
        public string? RegistrationStartDate { get; set; }
    }

    public class ItemsCreateV1RequestItem
    {
        public string? AdministrationMethod { get; set; }
        public string? Allergens { get; set; }
        public string? Description { get; set; }
        public string? GlobalProductName { get; set; }
        public string? ItemBrand { get; set; }
        public string? ItemCategory { get; set; }
        public string? ItemIngredients { get; set; }
        public string? LabelImageFileSystemIds { get; set; }
        public string? LabelPhotoDescription { get; set; }
        public string? Name { get; set; }
        public string? NumberOfDoses { get; set; }
        public string? PackagingImageFileSystemIds { get; set; }
        public string? PackagingPhotoDescription { get; set; }
        public string? ProcessingJobCategoryName { get; set; }
        public string? ProcessingJobTypeName { get; set; }
        public string? ProductImageFileSystemIds { get; set; }
        public string? ProductPdffileSystemIds { get; set; }
        public string? ProductPhotoDescription { get; set; }
        public string? PublicIngredients { get; set; }
        public string? ServingSize { get; set; }
        public string? Strain { get; set; }
        public int? SupplyDurationDays { get; set; }
        public double? UnitCbdAcontent { get; set; }
        public double? UnitCbdAcontentDose { get; set; }
        public string? UnitCbdAcontentDoseUnitOfMeasure { get; set; }
        public string? UnitCbdAcontentUnitOfMeasure { get; set; }
        public double? UnitCbdApercent { get; set; }
        public double? UnitCbdContent { get; set; }
        public double? UnitCbdContentDose { get; set; }
        public string? UnitCbdContentDoseUnitOfMeasure { get; set; }
        public string? UnitCbdContentUnitOfMeasure { get; set; }
        public double? UnitCbdPercent { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcAcontent { get; set; }
        public double? UnitThcAcontentDose { get; set; }
        public string? UnitThcAcontentDoseUnitOfMeasure { get; set; }
        public string? UnitThcAcontentUnitOfMeasure { get; set; }
        public double? UnitThcApercent { get; set; }
        public double? UnitThcContent { get; set; }
        public double? UnitThcContentDose { get; set; }
        public string? UnitThcContentDoseUnitOfMeasure { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitVolume { get; set; }
        public string? UnitVolumeUnitOfMeasure { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class ItemsCreateV2RequestItem
    {
        public string? AdministrationMethod { get; set; }
        public string? Allergens { get; set; }
        public string? Description { get; set; }
        public string? GlobalProductName { get; set; }
        public string? ItemBrand { get; set; }
        public string? ItemCategory { get; set; }
        public string? ItemIngredients { get; set; }
        public string? LabelImageFileSystemIds { get; set; }
        public string? LabelPhotoDescription { get; set; }
        public string? Name { get; set; }
        public string? NumberOfDoses { get; set; }
        public string? PackagingImageFileSystemIds { get; set; }
        public string? PackagingPhotoDescription { get; set; }
        public string? ProcessingJobCategoryName { get; set; }
        public string? ProcessingJobTypeName { get; set; }
        public string? ProductImageFileSystemIds { get; set; }
        public string? ProductPdffileSystemIds { get; set; }
        public string? ProductPhotoDescription { get; set; }
        public string? PublicIngredients { get; set; }
        public string? ServingSize { get; set; }
        public string? Strain { get; set; }
        public int? SupplyDurationDays { get; set; }
        public double? UnitCbdAcontent { get; set; }
        public double? UnitCbdAcontentDose { get; set; }
        public string? UnitCbdAcontentDoseUnitOfMeasure { get; set; }
        public string? UnitCbdAcontentUnitOfMeasure { get; set; }
        public double? UnitCbdApercent { get; set; }
        public double? UnitCbdContent { get; set; }
        public double? UnitCbdContentDose { get; set; }
        public string? UnitCbdContentDoseUnitOfMeasure { get; set; }
        public string? UnitCbdContentUnitOfMeasure { get; set; }
        public double? UnitCbdPercent { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcAcontent { get; set; }
        public double? UnitThcAcontentDose { get; set; }
        public string? UnitThcAcontentDoseUnitOfMeasure { get; set; }
        public string? UnitThcAcontentUnitOfMeasure { get; set; }
        public double? UnitThcApercent { get; set; }
        public double? UnitThcContent { get; set; }
        public double? UnitThcContentDose { get; set; }
        public string? UnitThcContentDoseUnitOfMeasure { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitVolume { get; set; }
        public string? UnitVolumeUnitOfMeasure { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class ItemsCreateBrandV2RequestItem
    {
        public string? Name { get; set; }
    }

    public class ItemsCreateFileV2RequestItem
    {
        public string? EncodedImageBase64 { get; set; }
        public string? FileName { get; set; }
    }

    public class ItemsCreatePhotoV1RequestItem
    {
        public string? EncodedImageBase64 { get; set; }
        public string? FileName { get; set; }
    }

    public class ItemsCreatePhotoV2RequestItem
    {
        public string? EncodedImageBase64 { get; set; }
        public string? FileName { get; set; }
    }

    public class ItemsCreateUpdateV1RequestItem
    {
        public string? AdministrationMethod { get; set; }
        public string? Allergens { get; set; }
        public string? Description { get; set; }
        public string? GlobalProductName { get; set; }
        public int Id { get; set; }
        public string? ItemBrand { get; set; }
        public string? ItemCategory { get; set; }
        public string? ItemIngredients { get; set; }
        public string? LabelImageFileSystemIds { get; set; }
        public string? LabelPhotoDescription { get; set; }
        public string? Name { get; set; }
        public string? NumberOfDoses { get; set; }
        public string? PackagingImageFileSystemIds { get; set; }
        public string? PackagingPhotoDescription { get; set; }
        public string? ProcessingJobCategoryName { get; set; }
        public string? ProcessingJobTypeName { get; set; }
        public string? ProductImageFileSystemIds { get; set; }
        public string? ProductPdffileSystemIds { get; set; }
        public string? ProductPhotoDescription { get; set; }
        public string? PublicIngredients { get; set; }
        public string? ServingSize { get; set; }
        public string? Strain { get; set; }
        public int? SupplyDurationDays { get; set; }
        public double? UnitCbdAcontent { get; set; }
        public double? UnitCbdAcontentDose { get; set; }
        public string? UnitCbdAcontentDoseUnitOfMeasure { get; set; }
        public string? UnitCbdAcontentUnitOfMeasure { get; set; }
        public double? UnitCbdApercent { get; set; }
        public double? UnitCbdContent { get; set; }
        public double? UnitCbdContentDose { get; set; }
        public string? UnitCbdContentDoseUnitOfMeasure { get; set; }
        public string? UnitCbdContentUnitOfMeasure { get; set; }
        public double? UnitCbdPercent { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcAcontent { get; set; }
        public double? UnitThcAcontentDose { get; set; }
        public string? UnitThcAcontentDoseUnitOfMeasure { get; set; }
        public string? UnitThcAcontentUnitOfMeasure { get; set; }
        public double? UnitThcApercent { get; set; }
        public double? UnitThcContent { get; set; }
        public double? UnitThcContentDose { get; set; }
        public string? UnitThcContentDoseUnitOfMeasure { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitVolume { get; set; }
        public string? UnitVolumeUnitOfMeasure { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class ItemsUpdateV2RequestItem
    {
        public string? AdministrationMethod { get; set; }
        public string? Allergens { get; set; }
        public string? Description { get; set; }
        public string? GlobalProductName { get; set; }
        public int Id { get; set; }
        public string? ItemBrand { get; set; }
        public string? ItemCategory { get; set; }
        public string? ItemIngredients { get; set; }
        public string? LabelImageFileSystemIds { get; set; }
        public string? LabelPhotoDescription { get; set; }
        public string? Name { get; set; }
        public string? NumberOfDoses { get; set; }
        public string? PackagingImageFileSystemIds { get; set; }
        public string? PackagingPhotoDescription { get; set; }
        public string? ProcessingJobCategoryName { get; set; }
        public string? ProcessingJobTypeName { get; set; }
        public string? ProductImageFileSystemIds { get; set; }
        public string? ProductPdffileSystemIds { get; set; }
        public string? ProductPhotoDescription { get; set; }
        public string? PublicIngredients { get; set; }
        public string? ServingSize { get; set; }
        public string? Strain { get; set; }
        public int? SupplyDurationDays { get; set; }
        public double? UnitCbdAcontent { get; set; }
        public double? UnitCbdAcontentDose { get; set; }
        public string? UnitCbdAcontentDoseUnitOfMeasure { get; set; }
        public string? UnitCbdAcontentUnitOfMeasure { get; set; }
        public double? UnitCbdApercent { get; set; }
        public double? UnitCbdContent { get; set; }
        public double? UnitCbdContentDose { get; set; }
        public string? UnitCbdContentDoseUnitOfMeasure { get; set; }
        public string? UnitCbdContentUnitOfMeasure { get; set; }
        public double? UnitCbdPercent { get; set; }
        public string? UnitOfMeasure { get; set; }
        public double? UnitThcAcontent { get; set; }
        public double? UnitThcAcontentDose { get; set; }
        public string? UnitThcAcontentDoseUnitOfMeasure { get; set; }
        public string? UnitThcAcontentUnitOfMeasure { get; set; }
        public double? UnitThcApercent { get; set; }
        public double? UnitThcContent { get; set; }
        public double? UnitThcContentDose { get; set; }
        public string? UnitThcContentDoseUnitOfMeasure { get; set; }
        public string? UnitThcContentUnitOfMeasure { get; set; }
        public double? UnitThcPercent { get; set; }
        public double? UnitVolume { get; set; }
        public string? UnitVolumeUnitOfMeasure { get; set; }
        public double? UnitWeight { get; set; }
        public string? UnitWeightUnitOfMeasure { get; set; }
    }

    public class ItemsUpdateBrandV2RequestItem
    {
        public int Id { get; set; }
        public string? Name { get; set; }
    }

    public class LocationsCreateV1RequestItem
    {
        public string? LocationTypeName { get; set; }
        public string? Name { get; set; }
    }

    public class LocationsCreateV2RequestItem
    {
        public string? LocationTypeName { get; set; }
        public string? Name { get; set; }
    }

    public class LocationsCreateUpdateV1RequestItem
    {
        public int Id { get; set; }
        public string? LocationTypeName { get; set; }
        public string? Name { get; set; }
    }

    public class LocationsUpdateV2RequestItem
    {
        public int Id { get; set; }
        public string? LocationTypeName { get; set; }
        public string? Name { get; set; }
    }

    public class PatientsCreateV2RequestItem
    {
        public string? ActualDate { get; set; }
        public int? ConcentrateOuncesAllowed { get; set; }
        public int? FlowerOuncesAllowed { get; set; }
        public bool HasSalesLimitExemption { get; set; }
        public int? InfusedOuncesAllowed { get; set; }
        public string? LicenseEffectiveEndDate { get; set; }
        public string? LicenseEffectiveStartDate { get; set; }
        public string? LicenseNumber { get; set; }
        public int? MaxConcentrateThcPercentAllowed { get; set; }
        public int? MaxFlowerThcPercentAllowed { get; set; }
        public int RecommendedPlants { get; set; }
        public int RecommendedSmokableQuantity { get; set; }
        public int? ThcOuncesAllowed { get; set; }
    }

    public class PatientsCreateAddV1RequestItem
    {
        public string? ActualDate { get; set; }
        public int? ConcentrateOuncesAllowed { get; set; }
        public int? FlowerOuncesAllowed { get; set; }
        public bool HasSalesLimitExemption { get; set; }
        public int? InfusedOuncesAllowed { get; set; }
        public string? LicenseEffectiveEndDate { get; set; }
        public string? LicenseEffectiveStartDate { get; set; }
        public string? LicenseNumber { get; set; }
        public int? MaxConcentrateThcPercentAllowed { get; set; }
        public int? MaxFlowerThcPercentAllowed { get; set; }
        public int RecommendedPlants { get; set; }
        public int RecommendedSmokableQuantity { get; set; }
        public int? ThcOuncesAllowed { get; set; }
    }

    public class PatientsCreateUpdateV1RequestItem
    {
        public string? ActualDate { get; set; }
        public int? ConcentrateOuncesAllowed { get; set; }
        public int? FlowerOuncesAllowed { get; set; }
        public bool HasSalesLimitExemption { get; set; }
        public int? InfusedOuncesAllowed { get; set; }
        public string? LicenseEffectiveEndDate { get; set; }
        public string? LicenseEffectiveStartDate { get; set; }
        public string? LicenseNumber { get; set; }
        public int? MaxConcentrateThcPercentAllowed { get; set; }
        public int? MaxFlowerThcPercentAllowed { get; set; }
        public string? NewLicenseNumber { get; set; }
        public int RecommendedPlants { get; set; }
        public int RecommendedSmokableQuantity { get; set; }
        public int? ThcOuncesAllowed { get; set; }
    }

    public class PatientsUpdateV2RequestItem
    {
        public string? ActualDate { get; set; }
        public int? ConcentrateOuncesAllowed { get; set; }
        public int? FlowerOuncesAllowed { get; set; }
        public bool HasSalesLimitExemption { get; set; }
        public int? InfusedOuncesAllowed { get; set; }
        public string? LicenseEffectiveEndDate { get; set; }
        public string? LicenseEffectiveStartDate { get; set; }
        public string? LicenseNumber { get; set; }
        public int? MaxConcentrateThcPercentAllowed { get; set; }
        public int? MaxFlowerThcPercentAllowed { get; set; }
        public string? NewLicenseNumber { get; set; }
        public int RecommendedPlants { get; set; }
        public int RecommendedSmokableQuantity { get; set; }
        public int? ThcOuncesAllowed { get; set; }
    }

    public class RetailIdCreateAssociateV2RequestItem
    {
        public string? PackageLabel { get; set; }
        public List<string?>? QrUrls { get; set; }
    }

    public class RetailIdCreateGenerateV2Request
    {
        public string? PackageLabel { get; set; }
        public int Quantity { get; set; }
    }

    public class RetailIdCreateMergeV2Request
    {
        public List<string?>? PackageLabels { get; set; }
    }

    public class RetailIdCreatePackageInfoV2Request
    {
        public List<string?>? PackageLabels { get; set; }
    }

    public class StrainsCreateV1RequestItem
    {
        public double CbdLevel { get; set; }
        public double IndicaPercentage { get; set; }
        public string? Name { get; set; }
        public double SativaPercentage { get; set; }
        public string? TestingStatus { get; set; }
        public double ThcLevel { get; set; }
    }

    public class StrainsCreateV2RequestItem
    {
        public double CbdLevel { get; set; }
        public double IndicaPercentage { get; set; }
        public string? Name { get; set; }
        public double SativaPercentage { get; set; }
        public string? TestingStatus { get; set; }
        public double ThcLevel { get; set; }
    }

    public class StrainsCreateUpdateV1RequestItem
    {
        public double CbdLevel { get; set; }
        public int Id { get; set; }
        public double IndicaPercentage { get; set; }
        public string? Name { get; set; }
        public double SativaPercentage { get; set; }
        public string? TestingStatus { get; set; }
        public double ThcLevel { get; set; }
    }

    public class StrainsUpdateV2RequestItem
    {
        public double CbdLevel { get; set; }
        public int Id { get; set; }
        public double IndicaPercentage { get; set; }
        public string? Name { get; set; }
        public double SativaPercentage { get; set; }
        public string? TestingStatus { get; set; }
        public double ThcLevel { get; set; }
    }
}
