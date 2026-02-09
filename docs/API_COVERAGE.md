# 📊 Public SDK Coverage Report

> **Coverage Sources:**
> - 🟢 **Captured**: Verified against real API response
> - 🔵 **Example**: Derived from Bruno documentation
> - 🟣 **Inferred**: Auto-derived from sibling endpoints

## 📈 Overall Coverage

| Metric | Count | %
|--------|-------|---
| **Total Endpoints** | 262 | 100%
| **Covered** | 257 | 98.1%
| 🟢 Captured | 74 | 28.2%
| 🔵 Example | 183 | 69.8%
| 🟣 Inferred | 0 | 0.0%

`███████████████████░` **98.1%**

## 📋 Coverage by Group

| Group | Total | Covered | Gap | Status |
|-------|-------|---------|-----|--------|
| **AdditivesTemplates** | 6 | 6 | 0 | 🟩 |
| **CaregiversStatus** | 1 | 1 | 0 | 🟩 |
| **Employees** | 2 | 1 | 1 | 🟧 |
| **Facilities** | 1 | 1 | 0 | 🟩 |
| **Harvests** | 15 | 15 | 0 | 🟩 |
| **Items** | 14 | 14 | 0 | 🟩 |
| **LabTests** | 7 | 7 | 0 | 🟩 |
| **Locations** | 7 | 7 | 0 | 🟩 |
| **Packages** | 34 | 32 | 2 | 🟨 |
| **PatientCheckIns** | 4 | 4 | 0 | 🟩 |
| **Patients** | 5 | 5 | 0 | 🟩 |
| **PatientsStatus** | 1 | 1 | 0 | 🟩 |
| **PlantBatches** | 21 | 21 | 0 | 🟩 |
| **Plants** | 35 | 35 | 0 | 🟩 |
| **ProcessingJob** | 15 | 13 | 2 | 🟨 |
| **RetailId** | 7 | 7 | 0 | 🟩 |
| **Sales** | 33 | 33 | 0 | 🟩 |
| **Strains** | 6 | 6 | 0 | 🟩 |
| **Sublocations** | 6 | 6 | 0 | 🟩 |
| **Tags** | 3 | 3 | 0 | 🟩 |
| **Transfers** | 26 | 26 | 0 | 🟩 |
| **Transporters** | 8 | 8 | 0 | 🟩 |
| **UnitsOfMeasure** | 3 | 3 | 0 | 🟩 |
| **WasteMethods** | 1 | 1 | 0 | 🟩 |
| **Webhooks** | 1 | 1 | 0 | 🟩 |

## 🕵️ Detailed Breakdown

### AdditivesTemplates

<details>
<summary>Endpoints (6/6 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateAdditivesTemplates` | 🔵 Example | `WriteResponse` |
| `GetActiveAdditivesTemplates` | 🔵 Example | `AdditivesTemplate` |
| `GetAdditivesTemplateById` | 🔵 Example | `AdditivesTemplate` |
| `GetAdditivesTemplatesById` | 🔵 Example | `AdditivesTemplate` |
| `GetInactiveAdditivesTemplates` | 🔵 Example | `AdditivesTemplate` |
| `UpdateAdditivesTemplates` | 🔵 Example | `WriteResponse` |

</details>

### CaregiversStatus

<details>
<summary>Endpoints (1/1 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `GetCaregiversStatusByCaregiverLicenseNumber` | 🔵 Example | `CaregiversStatus` |

</details>

### Employees

<details>
<summary>Endpoints (1/2 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `GetEmployees` | 🔵 Example | `Employee` |
| `GetPermissions` | ❌ Missing | `-` |

</details>

### Facilities

<details>
<summary>Endpoints (1/1 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `GetFacilities` | 🟢 Captured | `Facility` |

</details>

### Harvests

<details>
<summary>Endpoints (15/15 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateHarvestsPackages` | 🔵 Example | `WriteResponse` |
| `CreateHarvestsWaste` | 🔵 Example | `WriteResponse` |
| `CreatePackagesTesting` | 🔵 Example | `WriteResponse` |
| `FinishHarvests` | 🔵 Example | `WriteResponse` |
| `GetActiveHarvests` | 🟢 Captured | `Harvest` |
| `GetHarvestById` | 🔵 Example | `Harvest` |
| `GetHarvestsById` | 🔵 Example | `Harvest` |
| `GetHarvestsWaste` | 🔵 Example | `HarvestsWaste` |
| `GetInactiveHarvests` | 🟢 Captured | `Harvest` |
| `GetOnHoldHarvests` | 🟢 Captured | `Harvest` |
| `GetWasteTypes` | 🟢 Captured | `WasteType` |
| `UnfinishHarvests` | 🔵 Example | `WriteResponse` |
| `UpdateHarvestsLocation` | 🔵 Example | `WriteResponse` |
| `UpdateRename` | 🔵 Example | `WriteResponse` |
| `UpdateRestoreHarvestedPlants` | 🔵 Example | `WriteResponse` |

</details>

### Items

<details>
<summary>Endpoints (14/14 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateBrand` | 🔵 Example | `WriteResponse` |
| `CreateFile` | 🔵 Example | `WriteResponse` |
| `CreateItems` | 🔵 Example | `WriteResponse` |
| `CreatePhoto` | 🔵 Example | `WriteResponse` |
| `GetActiveItems` | 🟢 Captured | `Item` |
| `GetBrands` | 🔵 Example | `Brand` |
| `GetCategories` | 🟢 Captured | `Category` |
| `GetFileById` | 🔵 Example | `File` |
| `GetInactiveItems` | 🟢 Captured | `Item` |
| `GetItemById` | 🔵 Example | `Item` |
| `GetItemsById` | 🔵 Example | `Item` |
| `GetPhotoById` | 🔵 Example | `Photo` |
| `UpdateBrand` | 🔵 Example | `WriteResponse` |
| `UpdateItems` | 🔵 Example | `WriteResponse` |

</details>

### LabTests

<details>
<summary>Endpoints (7/7 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateRecord` | 🔵 Example | `WriteResponse` |
| `GetBatches` | 🟢 Captured | `Batch` |
| `GetLabTestsTypes` | 🟢 Captured | `LabTestsType` |
| `GetResults` | 🔵 Example | `Result` |
| `GetStates` | 🟢 Captured | `-` |
| `UpdateLabTestDocument` | 🔵 Example | `WriteResponse` |
| `UpdateResultsRelease` | 🔵 Example | `WriteResponse` |

</details>

### Locations

<details>
<summary>Endpoints (7/7 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateLocations` | 🔵 Example | `WriteResponse` |
| `GetActiveLocations` | 🟢 Captured | `LocationsLocation` |
| `GetInactiveLocations` | 🟢 Captured | `LocationsLocation` |
| `GetLocationById` | 🔵 Example | `LocationsLocation` |
| `GetLocationsById` | 🔵 Example | `LocationsLocation` |
| `GetLocationsTypes` | 🟢 Captured | `LocationsType` |
| `UpdateLocations` | 🔵 Example | `WriteResponse` |

</details>

### Packages

<details>
<summary>Endpoints (32/34 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateAdjustPackages` | 🔵 Example | `WriteResponse` |
| `CreatePackagesPackages` | 🔵 Example | `WriteResponse` |
| `CreatePackagesPlantings` | 🔵 Example | `WriteResponse` |
| `CreateTesting` | 🔵 Example | `WriteResponse` |
| `FinishPackages` | ❌ Missing | `-` |
| `FinishedgoodFlag` | 🔵 Example | `WriteResponse` |
| `FinishedgoodUnflag` | 🔵 Example | `WriteResponse` |
| `GetActivePackages` | 🟢 Captured | `PackagesPackage` |
| `GetAdjustReasons` | 🟢 Captured | `AdjustReason` |
| `GetAdjustments` | 🔵 Example | `Adjustment` |
| `GetInTransitPackages` | 🟢 Captured | `InTransit` |
| `GetInactivePackages` | 🟢 Captured | `PackagesPackage` |
| `GetLabSamples` | 🟢 Captured | `PackagesPackage` |
| `GetOnHoldPackages` | 🟢 Captured | `PackagesPackage` |
| `GetPackageById` | 🔵 Example | `PackagesPackage` |
| `GetPackagesById` | 🔵 Example | `PackagesPackage` |
| `GetPackagesByLabel` | 🔵 Example | `PackagesPackage` |
| `GetPackagesTypes` | 🟢 Captured | `-` |
| `GetSourceHarvestById` | 🔵 Example | `SourceHarvest` |
| `GetTransferred` | 🟢 Captured | `PackagesPackage` |
| `UnfinishPackages` | ❌ Missing | `-` |
| `UpdateAdjustPackages` | 🔵 Example | `WriteResponse` |
| `UpdateDecontaminate` | 🔵 Example | `WriteResponse` |
| `UpdateDonationFlag` | 🔵 Example | `WriteResponse` |
| `UpdateDonationUnflag` | 🔵 Example | `WriteResponse` |
| `UpdateExternalid` | 🔵 Example | `WriteResponse` |
| `UpdateItem` | 🔵 Example | `WriteResponse` |
| `UpdateLabtestsRequired` | 🔵 Example | `WriteResponse` |
| `UpdateNote` | 🔵 Example | `WriteResponse` |
| `UpdatePackagesLocation` | 🔵 Example | `WriteResponse` |
| `UpdateRemediate` | 🔵 Example | `WriteResponse` |
| `UpdateTradeSampleFlag` | 🔵 Example | `WriteResponse` |
| `UpdateTradeSampleUnflag` | 🔵 Example | `WriteResponse` |
| `UpdateUseByDate` | 🔵 Example | `WriteResponse` |

</details>

### PatientCheckIns

<details>
<summary>Endpoints (4/4 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreatePatientCheckIns` | 🔵 Example | `WriteResponse` |
| `GetLocations` | 🟢 Captured | `PatientCheckInsLocation` |
| `GetPatientCheckIns` | 🔵 Example | `PatientCheckIn` |
| `UpdatePatientCheckIns` | 🔵 Example | `WriteResponse` |

</details>

### Patients

<details>
<summary>Endpoints (5/5 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreatePatients` | 🔵 Example | `WriteResponse` |
| `GetActivePatients` | 🔵 Example | `Patient` |
| `GetPatientById` | 🔵 Example | `Patient` |
| `GetPatientsById` | 🔵 Example | `Patient` |
| `UpdatePatients` | 🔵 Example | `WriteResponse` |

</details>

### PatientsStatus

<details>
<summary>Endpoints (1/1 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `GetPatientsStatusesByPatientLicenseNumber` | 🔵 Example | `PatientsStatus` |

</details>

### PlantBatches

<details>
<summary>Endpoints (21/21 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateAdjustPlantBatches` | 🔵 Example | `WriteResponse` |
| `CreateGrowthPhase` | 🔵 Example | `WriteResponse` |
| `CreatePackagesFromMotherPlant` | 🔵 Example | `WriteResponse` |
| `CreatePlantBatchesAdditives` | 🔵 Example | `WriteResponse` |
| `CreatePlantBatchesAdditivesUsingTemplate` | 🔵 Example | `WriteResponse` |
| `CreatePlantBatchesPackages` | 🔵 Example | `WriteResponse` |
| `CreatePlantBatchesPlantings` | 🔵 Example | `WriteResponse` |
| `CreatePlantBatchesWaste` | 🔵 Example | `WriteResponse` |
| `CreateSplit` | 🔵 Example | `WriteResponse` |
| `DeletePlantBatches` | 🔵 Example | `WriteResponse` |
| `GetActivePlantBatches` | 🟢 Captured | `PlantBatch` |
| `GetInactivePlantBatches` | 🟢 Captured | `PlantBatch` |
| `GetPlantBatchById` | 🔵 Example | `PlantBatch` |
| `GetPlantBatchesById` | 🔵 Example | `PlantBatch` |
| `GetPlantBatchesTypes` | 🟢 Captured | `PlantBatchesType` |
| `GetPlantBatchesWaste` | 🟢 Captured | `PlantBatchesWaste` |
| `GetPlantBatchesWasteReasons` | 🟢 Captured | `PlantBatchesWasteReason` |
| `UpdateName` | 🔵 Example | `WriteResponse` |
| `UpdatePlantBatchesLocation` | 🔵 Example | `WriteResponse` |
| `UpdatePlantBatchesStrain` | 🔵 Example | `WriteResponse` |
| `UpdatePlantBatchesTag` | 🔵 Example | `WriteResponse` |

</details>

### Plants

<details>
<summary>Endpoints (35/35 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateAdditivesByLocation` | 🔵 Example | `WriteResponse` |
| `CreateAdditivesByLocationUsingTemplate` | 🔵 Example | `WriteResponse` |
| `CreateManicure` | 🔵 Example | `WriteResponse` |
| `CreatePlantBatchPackages` | 🔵 Example | `WriteResponse` |
| `CreatePlantsAdditives` | 🔵 Example | `WriteResponse` |
| `CreatePlantsAdditivesUsingTemplate` | 🔵 Example | `WriteResponse` |
| `CreatePlantsPlantings` | 🔵 Example | `WriteResponse` |
| `CreatePlantsWaste` | 🔵 Example | `WriteResponse` |
| `DeletePlants` | 🔵 Example | `WriteResponse` |
| `GetAdditives` | 🟢 Captured | `Additive` |
| `GetAdditivesTypes` | 🟢 Captured | `-` |
| `GetFloweringPlants` | 🟢 Captured | `Plant` |
| `GetGrowthPhases` | 🟢 Captured | `-` |
| `GetInactivePlants` | 🟢 Captured | `Plant` |
| `GetMotherInactivePlants` | 🟢 Captured | `Mother` |
| `GetMotherOnHoldPlants` | 🟢 Captured | `Mother` |
| `GetMotherPlants` | 🟢 Captured | `Mother` |
| `GetOnHoldPlants` | 🟢 Captured | `Plant` |
| `GetPlantById` | 🔵 Example | `Plant` |
| `GetPlantsById` | 🔵 Example | `Plant` |
| `GetPlantsByLabel` | 🔵 Example | `Plant` |
| `GetPlantsWaste` | 🟢 Captured | `PlantsWaste` |
| `GetPlantsWasteMethods` | 🟢 Captured | `WasteMethod` |
| `GetPlantsWasteReasons` | 🟢 Captured | `WasteReason` |
| `GetVegetativePlants` | 🟢 Captured | `Plant` |
| `GetWasteById` | 🔵 Example | `PlantsWaste` |
| `GetWastePackageById` | 🔵 Example | `WastePackage` |
| `UpdateAdjustPlants` | 🔵 Example | `WriteResponse` |
| `UpdateGrowthPhase` | 🔵 Example | `WriteResponse` |
| `UpdateHarvest` | 🔵 Example | `WriteResponse` |
| `UpdateMerge` | 🔵 Example | `WriteResponse` |
| `UpdatePlantsLocation` | 🔵 Example | `WriteResponse` |
| `UpdatePlantsStrain` | 🔵 Example | `WriteResponse` |
| `UpdatePlantsTag` | 🔵 Example | `WriteResponse` |
| `UpdateSplit` | 🔵 Example | `WriteResponse` |

</details>

### ProcessingJob

<details>
<summary>Endpoints (13/15 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateAdjustProcessingJob` | 🔵 Example | `WriteResponse` |
| `CreateJobTypes` | 🔵 Example | `WriteResponse` |
| `CreateProcessingJobPackages` | 🔵 Example | `WriteResponse` |
| `FinishProcessingJob` | ❌ Missing | `-` |
| `GetActiveJobTypes` | 🟢 Captured | `ActiveJobType` |
| `GetActiveProcessingJob` | 🟢 Captured | `ProcessingJob` |
| `GetInactiveJobTypes` | 🟢 Captured | `InactiveJobType` |
| `GetInactiveProcessingJob` | 🟢 Captured | `ProcessingJob` |
| `GetJobTypeById` | 🔵 Example | `InactiveJobType` |
| `GetJobTypesAttributes` | 🟢 Captured | `JobTypesAttribute` |
| `GetJobTypesCategories` | 🟢 Captured | `JobTypesCategory` |
| `GetProcessingJobById` | 🔵 Example | `ProcessingJob` |
| `StartProcessingJob` | 🔵 Example | `WriteResponse` |
| `UnfinishProcessingJob` | ❌ Missing | `-` |
| `UpdateJobTypes` | 🔵 Example | `WriteResponse` |

</details>

### RetailId

<details>
<summary>Endpoints (7/7 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateAssociate` | 🔵 Example | `WriteResponse` |
| `CreateGenerate` | 🔵 Example | `WriteResponse` |
| `CreateMerge` | 🔵 Example | `WriteResponse` |
| `CreatePackagesInfo` | 🔵 Example | `WriteResponse` |
| `GetAllotment` | 🔵 Example | `Allotment` |
| `GetReceiveByLabel` | 🔵 Example | `Receive` |
| `GetReceiveQrByShortCode` | 🔵 Example | `ReceiveQrByShortCode` |

</details>

### Sales

<details>
<summary>Endpoints (33/33 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateDeliveries` | 🔵 Example | `WriteResponse` |
| `CreateDeliveriesRetailerDepart` | 🔵 Example | `WriteResponse` |
| `CreateDeliveriesRetailerEnd` | 🔵 Example | `WriteResponse` |
| `CreateDeliveriesRetailerRestock` | 🔵 Example | `WriteResponse` |
| `CreateReceipts` | 🔵 Example | `WriteResponse` |
| `CreateSalesDeliveriesRetailer` | 🔵 Example | `WriteResponse` |
| `GetActiveDeliveries` | 🟢 Captured | `ActiveDelivery` |
| `GetActiveDeliveriesRetailer` | 🔵 Example | `ActiveDeliveriesRetailer` |
| `GetActiveReceipts` | 🟢 Captured | `ActiveReceipt` |
| `GetCounties` | 🟢 Captured | `County` |
| `GetCustomerTypes` | 🟢 Captured | `-` |
| `GetDeliveriesRetailerById` | 🔵 Example | `InactiveDeliveriesRetailer` |
| `GetDeliveriesReturnReasons` | 🟢 Captured | `DeliveriesReturnReason` |
| `GetDeliveryById` | 🔵 Example | `ActiveDelivery` |
| `GetDeliveryRetailerById` | 🔵 Example | `DeliveryRetailer` |
| `GetInactiveDeliveries` | 🟢 Captured | `InactiveDelivery` |
| `GetInactiveDeliveriesRetailer` | 🔵 Example | `InactiveDeliveriesRetailer` |
| `GetInactiveReceipts` | 🟢 Captured | `InactiveReceipt` |
| `GetPatientRegistrationLocations` | 🟢 Captured | `PatientRegistrationLocation` |
| `GetPaymentTypes` | 🟢 Captured | `-` |
| `GetReceiptById` | 🔵 Example | `ActiveReceipt` |
| `GetReceiptsExternalByExternalNumber` | 🔵 Example | `ReceiptsExternalByExternalNumber` |
| `GetSalesDeliveryById` | 🔵 Example | `SalesDelivery` |
| `UpdateDeliveries` | 🔵 Example | `WriteResponse` |
| `UpdateDeliveriesComplete` | 🔵 Example | `WriteResponse` |
| `UpdateDeliveriesHub` | 🔵 Example | `WriteResponse` |
| `UpdateDeliveriesHubAccept` | 🔵 Example | `WriteResponse` |
| `UpdateDeliveriesHubDepart` | 🔵 Example | `WriteResponse` |
| `UpdateDeliveriesHubVerifyID` | 🔵 Example | `WriteResponse` |
| `UpdateDeliveriesRetailer` | 🔵 Example | `WriteResponse` |
| `UpdateReceipts` | 🔵 Example | `WriteResponse` |
| `UpdateReceiptsFinalize` | 🔵 Example | `WriteResponse` |
| `UpdateReceiptsUnfinalize` | 🔵 Example | `WriteResponse` |

</details>

### Strains

<details>
<summary>Endpoints (6/6 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateStrains` | 🔵 Example | `WriteResponse` |
| `GetActiveStrains` | 🟢 Captured | `Strain` |
| `GetInactiveStrains` | 🟢 Captured | `Strain` |
| `GetStrainById` | 🔵 Example | `Strain` |
| `GetStrainsById` | 🔵 Example | `Strain` |
| `UpdateStrains` | 🔵 Example | `WriteResponse` |

</details>

### Sublocations

<details>
<summary>Endpoints (6/6 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateSublocations` | 🔵 Example | `WriteResponse` |
| `GetActiveSublocations` | 🟢 Captured | `Sublocation` |
| `GetInactiveSublocations` | 🟢 Captured | `Sublocation` |
| `GetSublocationById` | 🔵 Example | `Sublocation` |
| `GetSublocationsById` | 🔵 Example | `Sublocation` |
| `UpdateSublocations` | 🔵 Example | `WriteResponse` |

</details>

### Tags

<details>
<summary>Endpoints (3/3 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `GetAvailablePackage` | 🟢 Captured | `Tag` |
| `GetAvailablePlant` | 🟢 Captured | `Tag` |
| `GetStagedTags` | 🔵 Example | `Staged` |

</details>

### Transfers

<details>
<summary>Endpoints (26/26 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateHubArrive` | 🔵 Example | `WriteResponse` |
| `CreateHubCheckin` | 🔵 Example | `WriteResponse` |
| `CreateHubCheckout` | 🔵 Example | `WriteResponse` |
| `CreateHubDepart` | 🔵 Example | `WriteResponse` |
| `CreateIncomingExternal` | 🔵 Example | `WriteResponse` |
| `CreateOutgoingTemplates` | 🔵 Example | `WriteResponse` |
| `GetDeliveriesPackagesStates` | 🟢 Captured | `-` |
| `GetDeliveryPackageById` | 🔵 Example | `DeliveryPackage` |
| `GetDeliveryPackageRequiredlabtestbatchById` | 🔵 Example | `DeliveryPackageRequiredlabtestbatch` |
| `GetDeliveryPackageWholesaleById` | 🔵 Example | `DeliveryPackageWholesale` |
| `GetDeliveryTransporterById` | 🔵 Example | `DeliveryTransporter` |
| `GetDeliveryTransporterDetailById` | 🔵 Example | `DeliveryTransporterDetail` |
| `GetHub` | 🟢 Captured | `Hub` |
| `GetIncomingTransfers` | 🟢 Captured | `Transfer` |
| `GetManifestPdfById` | 🔵 Example | `ManifestPdf` |
| `GetOutgoingTemplateDeliveryById` | 🔵 Example | `TemplateDelivery` |
| `GetOutgoingTemplateDeliveryPackageById` | 🔵 Example | `TemplateDeliveryPackage` |
| `GetOutgoingTemplateDeliveryTransporterById` | 🔵 Example | `TemplateDeliveryTransporter` |
| `GetOutgoingTemplateDeliveryTransporterDetailById` | 🔵 Example | `TemplateDeliveryTransporterDetail` |
| `GetOutgoingTemplates` | 🟢 Captured | `Template` |
| `GetOutgoingTransfers` | 🟢 Captured | `Transfer` |
| `GetRejectedTransfers` | 🟢 Captured | `Transfer` |
| `GetTransfersDeliveryById` | 🔵 Example | `TransfersDelivery` |
| `GetTransfersTypes` | 🟢 Captured | `TransfersType` |
| `UpdateIncomingExternal` | 🔵 Example | `WriteResponse` |
| `UpdateOutgoingTemplates` | 🔵 Example | `WriteResponse` |

</details>

### Transporters

<details>
<summary>Endpoints (8/8 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `CreateDrivers` | 🔵 Example | `WriteResponse` |
| `CreateVehicles` | 🔵 Example | `WriteResponse` |
| `GetDriverById` | 🔵 Example | `TransportersDriver` |
| `GetDrivers` | 🟢 Captured | `TransportersDriver` |
| `GetVehicleById` | 🔵 Example | `TransportersVehicle` |
| `GetVehicles` | 🟢 Captured | `TransportersVehicle` |
| `UpdateDrivers` | 🔵 Example | `WriteResponse` |
| `UpdateVehicles` | 🔵 Example | `WriteResponse` |

</details>

### UnitsOfMeasure

<details>
<summary>Endpoints (3/3 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `GetActiveUnitsOfMeasure` | 🟢 Captured | `UnitOfMeasure` |
| `GetInactiveUnitsOfMeasure` | 🟢 Captured | `UnitOfMeasure` |
| `GetUnitOfMeasureById` | 🔵 Example | `UnitOfMeasure` |

</details>

### WasteMethods

<details>
<summary>Endpoints (1/1 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `GetWasteMethodsWasteMethods` | 🟢 Captured | `WasteMethod` |

</details>

### Webhooks

<details>
<summary>Endpoints (1/1 covered)</summary>

| Operation | Source | Model |
|-----------|--------|-------|
| `UpdateWebhooks` | 🔵 Example | `WriteResponse` |

</details>

---
*Generated from manifest.json + response scanning*
