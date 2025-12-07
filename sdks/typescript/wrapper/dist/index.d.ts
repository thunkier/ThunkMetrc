export * from '@thunkmetrc/client';
import { MetrcClient as InternalMetrcClient } from '@thunkmetrc/client';
import { RateLimiterConfig } from './RateLimiter';
export declare class MetrcWrapper {
    client: InternalMetrcClient;
    private rateLimiter;
    constructor(client: InternalMetrcClient, config?: RateLimiterConfig);
    tagsGetPackageAvailableV2(body?: any, licenseNumber?: string): Promise<any>;
    tagsGetPlantAvailableV2(body?: any, licenseNumber?: string): Promise<any>;
    tagsGetStagedV2(body?: any, licenseNumber?: string): Promise<any>;
    transfersCreateExternalIncomingV1(body: TransfersCreateExternalIncomingV1RequestItem[], licenseNumber?: string): Promise<any>;
    transfersCreateExternalIncomingV2(body: TransfersCreateExternalIncomingV2RequestItem[], licenseNumber?: string): Promise<any>;
    transfersCreateTemplatesV1(body: TransfersCreateTemplatesV1RequestItem[], licenseNumber?: string): Promise<any>;
    transfersCreateTemplatesOutgoingV2(body: TransfersCreateTemplatesOutgoingV2RequestItem[], licenseNumber?: string): Promise<any>;
    transfersDeleteExternalIncomingV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    transfersDeleteExternalIncomingV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    transfersDeleteTemplatesV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    transfersDeleteTemplatesOutgoingV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    transfersGetDeliveriesPackagesStatesV1(body?: any, No?: string): Promise<any>;
    transfersGetDeliveriesPackagesStatesV2(body?: any, No?: string): Promise<any>;
    transfersGetDeliveryV1(id: string, body?: any, No?: string): Promise<any>;
    transfersGetDeliveryV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetDeliveryPackageV1(id: string, body?: any, No?: string): Promise<any>;
    transfersGetDeliveryPackageV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetDeliveryPackageRequiredlabtestbatchesV1(id: string, body?: any, No?: string): Promise<any>;
    transfersGetDeliveryPackageRequiredlabtestbatchesV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetDeliveryPackageWholesaleV1(id: string, body?: any, No?: string): Promise<any>;
    transfersGetDeliveryPackageWholesaleV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetDeliveryTransportersV1(id: string, body?: any, No?: string): Promise<any>;
    transfersGetDeliveryTransportersV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetDeliveryTransportersDetailsV1(id: string, body?: any, No?: string): Promise<any>;
    transfersGetDeliveryTransportersDetailsV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetHubV2(body?: any, estimatedArrivalEnd?: string, estimatedArrivalStart?: string, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetIncomingV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    transfersGetIncomingV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetOutgoingV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    transfersGetOutgoingV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetRejectedV1(body?: any, licenseNumber?: string): Promise<any>;
    transfersGetRejectedV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetTemplatesV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    transfersGetTemplatesDeliveryV1(id: string, body?: any, No?: string): Promise<any>;
    transfersGetTemplatesDeliveryPackageV1(id: string, body?: any, No?: string): Promise<any>;
    transfersGetTemplatesDeliveryTransportersV1(id: string, body?: any, No?: string): Promise<any>;
    transfersGetTemplatesDeliveryTransportersDetailsV1(id: string, body?: any, No?: string): Promise<any>;
    transfersGetTemplatesOutgoingV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetTemplatesOutgoingDeliveryV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetTemplatesOutgoingDeliveryPackageV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetTemplatesOutgoingDeliveryTransportersV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id: string, body?: any, No?: string): Promise<any>;
    transfersGetTypesV1(body?: any, licenseNumber?: string): Promise<any>;
    transfersGetTypesV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    transfersUpdateExternalIncomingV1(body: TransfersUpdateExternalIncomingV1RequestItem[], licenseNumber?: string): Promise<any>;
    transfersUpdateExternalIncomingV2(body: TransfersUpdateExternalIncomingV2RequestItem[], licenseNumber?: string): Promise<any>;
    transfersUpdateTemplatesV1(body: TransfersUpdateTemplatesV1RequestItem[], licenseNumber?: string): Promise<any>;
    transfersUpdateTemplatesOutgoingV2(body: TransfersUpdateTemplatesOutgoingV2RequestItem[], licenseNumber?: string): Promise<any>;
    wasteMethodsGetAllV2(body?: any, No?: string): Promise<any>;
    additivesTemplatesCreateV2(body: AdditivestemplatesCreateV2RequestItem[], licenseNumber?: string): Promise<any>;
    additivesTemplatesGetV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    additivesTemplatesGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    additivesTemplatesGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    additivesTemplatesUpdateV2(body: AdditivestemplatesUpdateV2RequestItem[], licenseNumber?: string): Promise<any>;
    facilitiesGetAllV1(body?: any, No?: string): Promise<any>;
    facilitiesGetAllV2(body?: any, No?: string): Promise<any>;
    harvestsCreateFinishV1(body: HarvestsCreateFinishV1RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsCreatePackageV1(body: HarvestsCreatePackageV1RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsCreatePackageV2(body: HarvestsCreatePackageV2RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsCreatePackageTestingV1(body: HarvestsCreatePackageTestingV1RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsCreatePackageTestingV2(body: HarvestsCreatePackageTestingV2RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsCreateRemoveWasteV1(body: HarvestsCreateRemoveWasteV1RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsCreateUnfinishV1(body: HarvestsCreateUnfinishV1RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsCreateWasteV2(body: HarvestsCreateWasteV2RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsDeleteWasteV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    harvestsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    harvestsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    harvestsGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    harvestsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    harvestsGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    harvestsGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    harvestsGetOnholdV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    harvestsGetOnholdV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    harvestsGetWasteV2(body?: any, harvestId?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    harvestsGetWasteTypesV1(body?: any, No?: string): Promise<any>;
    harvestsGetWasteTypesV2(body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    harvestsUpdateFinishV2(body: HarvestsUpdateFinishV2RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsUpdateLocationV2(body: HarvestsUpdateLocationV2RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsUpdateMoveV1(body: HarvestsUpdateMoveV1RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsUpdateRenameV1(body: HarvestsUpdateRenameV1RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsUpdateRenameV2(body: HarvestsUpdateRenameV2RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsUpdateRestoreHarvestedPlantsV2(body: HarvestsUpdateRestoreHarvestedPlantsV2RequestItem[], licenseNumber?: string): Promise<any>;
    harvestsUpdateUnfinishV2(body: HarvestsUpdateUnfinishV2RequestItem[], licenseNumber?: string): Promise<any>;
    itemsCreateV1(body: ItemsCreateV1RequestItem[], licenseNumber?: string): Promise<any>;
    itemsCreateV2(body: ItemsCreateV2RequestItem[], licenseNumber?: string): Promise<any>;
    itemsCreateBrandV2(body: ItemsCreateBrandV2RequestItem[], licenseNumber?: string): Promise<any>;
    itemsCreateFileV2(body: ItemsCreateFileV2RequestItem[], licenseNumber?: string): Promise<any>;
    itemsCreatePhotoV1(body: ItemsCreatePhotoV1RequestItem[], licenseNumber?: string): Promise<any>;
    itemsCreatePhotoV2(body: ItemsCreatePhotoV2RequestItem[], licenseNumber?: string): Promise<any>;
    itemsCreateUpdateV1(body: ItemsCreateUpdateV1RequestItem[], licenseNumber?: string): Promise<any>;
    itemsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    itemsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    itemsDeleteBrandV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    itemsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    itemsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    itemsGetActiveV1(body?: any, licenseNumber?: string): Promise<any>;
    itemsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    itemsGetBrandsV1(body?: any, licenseNumber?: string): Promise<any>;
    itemsGetBrandsV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    itemsGetCategoriesV1(body?: any, licenseNumber?: string): Promise<any>;
    itemsGetCategoriesV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    itemsGetFileV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    itemsGetInactiveV1(body?: any, licenseNumber?: string): Promise<any>;
    itemsGetInactiveV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    itemsGetPhotoV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    itemsGetPhotoV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    itemsUpdateV2(body: ItemsUpdateV2RequestItem[], licenseNumber?: string): Promise<any>;
    itemsUpdateBrandV2(body: ItemsUpdateBrandV2RequestItem[], licenseNumber?: string): Promise<any>;
    labTestsCreateRecordV1(body: LabtestsCreateRecordV1RequestItem[], licenseNumber?: string): Promise<any>;
    labTestsCreateRecordV2(body: LabtestsCreateRecordV2RequestItem[], licenseNumber?: string): Promise<any>;
    labTestsGetBatchesV2(body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    labTestsGetLabtestdocumentV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    labTestsGetLabtestdocumentV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    labTestsGetResultsV1(body?: any, licenseNumber?: string, packageId?: string): Promise<any>;
    labTestsGetResultsV2(body?: any, licenseNumber?: string, packageId?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    labTestsGetStatesV1(body?: any, No?: string): Promise<any>;
    labTestsGetStatesV2(body?: any, No?: string): Promise<any>;
    labTestsGetTypesV1(body?: any, No?: string): Promise<any>;
    labTestsGetTypesV2(body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    labTestsUpdateLabtestdocumentV1(body: LabtestsUpdateLabtestdocumentV1RequestItem[], licenseNumber?: string): Promise<any>;
    labTestsUpdateLabtestdocumentV2(body: LabtestsUpdateLabtestdocumentV2RequestItem[], licenseNumber?: string): Promise<any>;
    labTestsUpdateResultReleaseV1(body: LabtestsUpdateResultReleaseV1RequestItem[], licenseNumber?: string): Promise<any>;
    labTestsUpdateResultReleaseV2(body: LabtestsUpdateResultReleaseV2RequestItem[], licenseNumber?: string): Promise<any>;
    strainsCreateV1(body: StrainsCreateV1RequestItem[], licenseNumber?: string): Promise<any>;
    strainsCreateV2(body: StrainsCreateV2RequestItem[], licenseNumber?: string): Promise<any>;
    strainsCreateUpdateV1(body: StrainsCreateUpdateV1RequestItem[], licenseNumber?: string): Promise<any>;
    strainsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    strainsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    strainsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    strainsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    strainsGetActiveV1(body?: any, licenseNumber?: string): Promise<any>;
    strainsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    strainsGetInactiveV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    strainsUpdateV2(body: StrainsUpdateV2RequestItem[], licenseNumber?: string): Promise<any>;
    transportersCreateDriverV2(body: TransportersCreateDriverV2RequestItem[], licenseNumber?: string): Promise<any>;
    transportersCreateVehicleV2(body: TransportersCreateVehicleV2RequestItem[], licenseNumber?: string): Promise<any>;
    transportersDeleteDriverV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    transportersDeleteVehicleV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    transportersGetDriverV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    transportersGetDriversV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    transportersGetVehicleV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    transportersGetVehiclesV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    transportersUpdateDriverV2(body: TransportersUpdateDriverV2RequestItem[], licenseNumber?: string): Promise<any>;
    transportersUpdateVehicleV2(body: TransportersUpdateVehicleV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreateV1(body: PackagesCreateV1RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreateV2(body: PackagesCreateV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreateAdjustV1(body: PackagesCreateAdjustV1RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreateAdjustV2(body: PackagesCreateAdjustV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreateChangeItemV1(body: PackagesCreateChangeItemV1RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreateChangeLocationV1(body: PackagesCreateChangeLocationV1RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreateFinishV1(body: PackagesCreateFinishV1RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreatePlantingsV1(body: PackagesCreatePlantingsV1RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreatePlantingsV2(body: PackagesCreatePlantingsV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreateRemediateV1(body: PackagesCreateRemediateV1RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreateTestingV1(body: PackagesCreateTestingV1RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreateTestingV2(body: PackagesCreateTestingV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesCreateUnfinishV1(body: PackagesCreateUnfinishV1RequestItem[], licenseNumber?: string): Promise<any>;
    packagesDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    packagesGetV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    packagesGetV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    packagesGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    packagesGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    packagesGetAdjustReasonsV1(body?: any, licenseNumber?: string): Promise<any>;
    packagesGetAdjustReasonsV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    packagesGetByLabelV1(label: string, body?: any, licenseNumber?: string): Promise<any>;
    packagesGetByLabelV2(label: string, body?: any, licenseNumber?: string): Promise<any>;
    packagesGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    packagesGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    packagesGetIntransitV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    packagesGetLabsamplesV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    packagesGetOnholdV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    packagesGetOnholdV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    packagesGetSourceHarvestV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    packagesGetTransferredV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    packagesGetTypesV1(body?: any, No?: string): Promise<any>;
    packagesGetTypesV2(body?: any, No?: string): Promise<any>;
    packagesUpdateAdjustV2(body: PackagesUpdateAdjustV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateChangeNoteV1(body: PackagesUpdateChangeNoteV1RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateDecontaminateV2(body: PackagesUpdateDecontaminateV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateDonationFlagV2(body: PackagesUpdateDonationFlagV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateDonationUnflagV2(body: PackagesUpdateDonationUnflagV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateExternalidV2(body: PackagesUpdateExternalidV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateFinishV2(body: PackagesUpdateFinishV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateItemV2(body: PackagesUpdateItemV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateLabTestRequiredV2(body: PackagesUpdateLabTestRequiredV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateLocationV2(body: PackagesUpdateLocationV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateNoteV2(body: PackagesUpdateNoteV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateRemediateV2(body: PackagesUpdateRemediateV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateTradesampleFlagV2(body: PackagesUpdateTradesampleFlagV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateTradesampleUnflagV2(body: PackagesUpdateTradesampleUnflagV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateUnfinishV2(body: PackagesUpdateUnfinishV2RequestItem[], licenseNumber?: string): Promise<any>;
    packagesUpdateUsebydateV2(body: PackagesUpdateUsebydateV2RequestItem[], licenseNumber?: string): Promise<any>;
    sublocationsCreateV2(body: SublocationsCreateV2RequestItem[], licenseNumber?: string): Promise<any>;
    sublocationsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    sublocationsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    sublocationsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    sublocationsGetInactiveV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    sublocationsUpdateV2(body: SublocationsUpdateV2RequestItem[], licenseNumber?: string): Promise<any>;
    unitsOfMeasureGetActiveV1(body?: any, No?: string): Promise<any>;
    unitsOfMeasureGetActiveV2(body?: any, No?: string): Promise<any>;
    unitsOfMeasureGetInactiveV2(body?: any, No?: string): Promise<any>;
    plantBatchesCreateAdditivesV1(body: PlantbatchesCreateAdditivesV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreateAdditivesV2(body: PlantbatchesCreateAdditivesV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreateAdditivesUsingtemplateV2(body: PlantbatchesCreateAdditivesUsingtemplateV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreateAdjustV1(body: PlantbatchesCreateAdjustV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreateAdjustV2(body: PlantbatchesCreateAdjustV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreateChangegrowthphaseV1(body: PlantbatchesCreateChangegrowthphaseV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreateGrowthphaseV2(body: PlantbatchesCreateGrowthphaseV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreatePackageV2(body: PlantbatchesCreatePackageV2RequestItem[], isFromMotherPlant?: string, licenseNumber?: string): Promise<any>;
    plantBatchesCreatePackageFrommotherplantV1(body: PlantbatchesCreatePackageFrommotherplantV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreatePackageFrommotherplantV2(body: PlantbatchesCreatePackageFrommotherplantV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreatePlantingsV2(body: PlantbatchesCreatePlantingsV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreateSplitV1(body: PlantbatchesCreateSplitV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreateSplitV2(body: PlantbatchesCreateSplitV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreateWasteV1(body: PlantbatchesCreateWasteV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreateWasteV2(body: PlantbatchesCreateWasteV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesCreatepackagesV1(body: PlantbatchesCreatepackagesV1RequestItem[], isFromMotherPlant?: string, licenseNumber?: string): Promise<any>;
    plantBatchesCreateplantingsV1(body: PlantbatchesCreateplantingsV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesDeleteV1(body?: any, licenseNumber?: string): Promise<any>;
    plantBatchesDeleteV2(body?: any, licenseNumber?: string): Promise<any>;
    plantBatchesGetV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    plantBatchesGetV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    plantBatchesGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    plantBatchesGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantBatchesGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    plantBatchesGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantBatchesGetTypesV1(body?: any, No?: string): Promise<any>;
    plantBatchesGetTypesV2(body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    plantBatchesGetWasteV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantBatchesGetWasteReasonsV1(body?: any, licenseNumber?: string): Promise<any>;
    plantBatchesGetWasteReasonsV2(body?: any, licenseNumber?: string): Promise<any>;
    plantBatchesUpdateLocationV2(body: PlantbatchesUpdateLocationV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesUpdateMoveplantbatchesV1(body: PlantbatchesUpdateMoveplantbatchesV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesUpdateNameV2(body: PlantbatchesUpdateNameV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesUpdateStrainV2(body: PlantbatchesUpdateStrainV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantBatchesUpdateTagV2(body: PlantbatchesUpdateTagV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateAdditivesV1(body: PlantsCreateAdditivesV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateAdditivesV2(body: PlantsCreateAdditivesV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateAdditivesBylocationV1(body: PlantsCreateAdditivesBylocationV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateAdditivesBylocationV2(body: PlantsCreateAdditivesBylocationV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateAdditivesBylocationUsingtemplateV2(body: PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateAdditivesUsingtemplateV2(body: PlantsCreateAdditivesUsingtemplateV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateChangegrowthphasesV1(body: PlantsCreateChangegrowthphasesV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateHarvestplantsV1(body: PlantsCreateHarvestplantsV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateManicureV2(body: PlantsCreateManicureV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateManicureplantsV1(body: PlantsCreateManicureplantsV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateMoveplantsV1(body: PlantsCreateMoveplantsV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreatePlantbatchPackageV1(body: PlantsCreatePlantbatchPackageV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreatePlantbatchPackageV2(body: PlantsCreatePlantbatchPackageV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreatePlantingsV1(body: PlantsCreatePlantingsV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreatePlantingsV2(body: PlantsCreatePlantingsV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateWasteV1(body: PlantsCreateWasteV1RequestItem[], licenseNumber?: string): Promise<any>;
    plantsCreateWasteV2(body: PlantsCreateWasteV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsDeleteV1(body?: any, licenseNumber?: string): Promise<any>;
    plantsDeleteV2(body?: any, licenseNumber?: string): Promise<any>;
    plantsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    plantsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    plantsGetAdditivesV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    plantsGetAdditivesV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetAdditivesTypesV1(body?: any, No?: string): Promise<any>;
    plantsGetAdditivesTypesV2(body?: any, No?: string): Promise<any>;
    plantsGetByLabelV1(label: string, body?: any, licenseNumber?: string): Promise<any>;
    plantsGetByLabelV2(label: string, body?: any, licenseNumber?: string): Promise<any>;
    plantsGetFloweringV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    plantsGetFloweringV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetGrowthPhasesV1(body?: any, licenseNumber?: string): Promise<any>;
    plantsGetGrowthPhasesV2(body?: any, licenseNumber?: string): Promise<any>;
    plantsGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    plantsGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetMotherV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetMotherInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetMotherOnholdV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetOnholdV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    plantsGetOnholdV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetVegetativeV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    plantsGetVegetativeV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetWasteV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetWasteMethodsAllV1(body?: any, No?: string): Promise<any>;
    plantsGetWasteMethodsAllV2(body?: any, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetWastePackageV2(id: string, body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetWastePlantV2(id: string, body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsGetWasteReasonsV1(body?: any, licenseNumber?: string): Promise<any>;
    plantsGetWasteReasonsV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    plantsUpdateAdjustV2(body: PlantsUpdateAdjustV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsUpdateGrowthphaseV2(body: PlantsUpdateGrowthphaseV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsUpdateHarvestV2(body: PlantsUpdateHarvestV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsUpdateLocationV2(body: PlantsUpdateLocationV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsUpdateMergeV2(body: PlantsUpdateMergeV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsUpdateSplitV2(body: PlantsUpdateSplitV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsUpdateStrainV2(body: PlantsUpdateStrainV2RequestItem[], licenseNumber?: string): Promise<any>;
    plantsUpdateTagV2(body: PlantsUpdateTagV2RequestItem[], licenseNumber?: string): Promise<any>;
    employeesGetAllV1(body?: any, licenseNumber?: string): Promise<any>;
    employeesGetAllV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    employeesGetPermissionsV2(body?: any, employeeLicenseNumber?: string, licenseNumber?: string): Promise<any>;
    patientCheckInsCreateV1(body: PatientcheckinsCreateV1RequestItem[], licenseNumber?: string): Promise<any>;
    patientCheckInsCreateV2(body: PatientcheckinsCreateV2RequestItem[], licenseNumber?: string): Promise<any>;
    patientCheckInsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    patientCheckInsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    patientCheckInsGetAllV1(body?: any, checkinDateEnd?: string, checkinDateStart?: string, licenseNumber?: string): Promise<any>;
    patientCheckInsGetAllV2(body?: any, checkinDateEnd?: string, checkinDateStart?: string, licenseNumber?: string): Promise<any>;
    patientCheckInsGetLocationsV1(body?: any, No?: string): Promise<any>;
    patientCheckInsGetLocationsV2(body?: any, No?: string): Promise<any>;
    patientCheckInsUpdateV1(body: PatientcheckinsUpdateV1RequestItem[], licenseNumber?: string): Promise<any>;
    patientCheckInsUpdateV2(body: PatientcheckinsUpdateV2RequestItem[], licenseNumber?: string): Promise<any>;
    patientsStatusGetStatusesByPatientLicenseNumberV1(patientLicenseNumber: string, body?: any, licenseNumber?: string): Promise<any>;
    patientsStatusGetStatusesByPatientLicenseNumberV2(patientLicenseNumber: string, body?: any, licenseNumber?: string): Promise<any>;
    processingJobsCreateAdjustV1(body: ProcessingjobsCreateAdjustV1RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsCreateAdjustV2(body: ProcessingjobsCreateAdjustV2RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsCreateJobtypesV1(body: ProcessingjobsCreateJobtypesV1RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsCreateJobtypesV2(body: ProcessingjobsCreateJobtypesV2RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsCreateStartV1(body: ProcessingjobsCreateStartV1RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsCreateStartV2(body: ProcessingjobsCreateStartV2RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsCreatepackagesV1(body: ProcessingjobsCreatepackagesV1RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsCreatepackagesV2(body: ProcessingjobsCreatepackagesV2RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    processingJobsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    processingJobsDeleteJobtypesV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    processingJobsDeleteJobtypesV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    processingJobsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    processingJobsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    processingJobsGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    processingJobsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    processingJobsGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    processingJobsGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    processingJobsGetJobtypesActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    processingJobsGetJobtypesActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    processingJobsGetJobtypesAttributesV1(body?: any, licenseNumber?: string): Promise<any>;
    processingJobsGetJobtypesAttributesV2(body?: any, licenseNumber?: string): Promise<any>;
    processingJobsGetJobtypesCategoriesV1(body?: any, licenseNumber?: string): Promise<any>;
    processingJobsGetJobtypesCategoriesV2(body?: any, licenseNumber?: string): Promise<any>;
    processingJobsGetJobtypesInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    processingJobsGetJobtypesInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    processingJobsUpdateFinishV1(body: ProcessingjobsUpdateFinishV1RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsUpdateFinishV2(body: ProcessingjobsUpdateFinishV2RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsUpdateJobtypesV1(body: ProcessingjobsUpdateJobtypesV1RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsUpdateJobtypesV2(body: ProcessingjobsUpdateJobtypesV2RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsUpdateUnfinishV1(body: ProcessingjobsUpdateUnfinishV1RequestItem[], licenseNumber?: string): Promise<any>;
    processingJobsUpdateUnfinishV2(body: ProcessingjobsUpdateUnfinishV2RequestItem[], licenseNumber?: string): Promise<any>;
    retailIdCreateAssociateV2(body: RetailidCreateAssociateV2RequestItem[], licenseNumber?: string): Promise<any>;
    retailIdCreateGenerateV2(body: RetailidCreateGenerateV2Request, licenseNumber?: string): Promise<any>;
    retailIdCreateMergeV2(body: RetailidCreateMergeV2Request, licenseNumber?: string): Promise<any>;
    retailIdCreatePackageInfoV2(body: RetailidCreatePackageInfoV2Request, licenseNumber?: string): Promise<any>;
    retailIdGetReceiveByLabelV2(label: string, body?: any, licenseNumber?: string): Promise<any>;
    retailIdGetReceiveQrByShortCodeV2(shortCode: string, body?: any, licenseNumber?: string): Promise<any>;
    sandboxCreateIntegratorSetupV2(body?: any, userKey?: string): Promise<any>;
    caregiversStatusGetByCaregiverLicenseNumberV1(caregiverLicenseNumber: string, body?: any, licenseNumber?: string): Promise<any>;
    caregiversStatusGetByCaregiverLicenseNumberV2(caregiverLicenseNumber: string, body?: any, licenseNumber?: string): Promise<any>;
    locationsCreateV1(body: LocationsCreateV1RequestItem[], licenseNumber?: string): Promise<any>;
    locationsCreateV2(body: LocationsCreateV2RequestItem[], licenseNumber?: string): Promise<any>;
    locationsCreateUpdateV1(body: LocationsCreateUpdateV1RequestItem[], licenseNumber?: string): Promise<any>;
    locationsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    locationsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    locationsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    locationsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    locationsGetActiveV1(body?: any, licenseNumber?: string): Promise<any>;
    locationsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    locationsGetInactiveV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    locationsGetTypesV1(body?: any, licenseNumber?: string): Promise<any>;
    locationsGetTypesV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    locationsUpdateV2(body: LocationsUpdateV2RequestItem[], licenseNumber?: string): Promise<any>;
    patientsCreateV2(body: PatientsCreateV2RequestItem[], licenseNumber?: string): Promise<any>;
    patientsCreateAddV1(body: PatientsCreateAddV1RequestItem[], licenseNumber?: string): Promise<any>;
    patientsCreateUpdateV1(body: PatientsCreateUpdateV1RequestItem[], licenseNumber?: string): Promise<any>;
    patientsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    patientsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    patientsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    patientsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    patientsGetActiveV1(body?: any, licenseNumber?: string): Promise<any>;
    patientsGetActiveV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    patientsUpdateV2(body: PatientsUpdateV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryV1(body: SalesCreateDeliveryV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryV2(body: SalesCreateDeliveryV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryRetailerV1(body: SalesCreateDeliveryRetailerV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryRetailerV2(body: SalesCreateDeliveryRetailerV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryRetailerDepartV1(body: SalesCreateDeliveryRetailerDepartV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryRetailerDepartV2(body: SalesCreateDeliveryRetailerDepartV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryRetailerEndV1(body: SalesCreateDeliveryRetailerEndV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryRetailerEndV2(body: SalesCreateDeliveryRetailerEndV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryRetailerRestockV1(body: SalesCreateDeliveryRetailerRestockV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryRetailerRestockV2(body: SalesCreateDeliveryRetailerRestockV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryRetailerSaleV1(body: SalesCreateDeliveryRetailerSaleV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateDeliveryRetailerSaleV2(body: SalesCreateDeliveryRetailerSaleV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateReceiptV1(body: SalesCreateReceiptV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateReceiptV2(body: SalesCreateReceiptV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesCreateTransactionByDateV1(date: string, body: SalesCreateTransactionByDateV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesDeleteDeliveryV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesDeleteDeliveryV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesDeleteDeliveryRetailerV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesDeleteDeliveryRetailerV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesDeleteReceiptV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesDeleteReceiptV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesGetCountiesV1(body?: any, No?: string): Promise<any>;
    salesGetCountiesV2(body?: any, No?: string): Promise<any>;
    salesGetCustomertypesV1(body?: any, No?: string): Promise<any>;
    salesGetCustomertypesV2(body?: any, No?: string): Promise<any>;
    salesGetDeliveriesActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any>;
    salesGetDeliveriesActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any>;
    salesGetDeliveriesInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any>;
    salesGetDeliveriesInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any>;
    salesGetDeliveriesRetailerActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    salesGetDeliveriesRetailerActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    salesGetDeliveriesRetailerInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any>;
    salesGetDeliveriesRetailerInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    salesGetDeliveriesReturnreasonsV1(body?: any, licenseNumber?: string): Promise<any>;
    salesGetDeliveriesReturnreasonsV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any>;
    salesGetDeliveryV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesGetDeliveryV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesGetDeliveryRetailerV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesGetDeliveryRetailerV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesGetPatientRegistrationsLocationsV1(body?: any, No?: string): Promise<any>;
    salesGetPatientRegistrationsLocationsV2(body?: any, No?: string): Promise<any>;
    salesGetPaymenttypesV1(body?: any, licenseNumber?: string): Promise<any>;
    salesGetPaymenttypesV2(body?: any, licenseNumber?: string): Promise<any>;
    salesGetReceiptV1(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesGetReceiptV2(id: string, body?: any, licenseNumber?: string): Promise<any>;
    salesGetReceiptsActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any>;
    salesGetReceiptsActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any>;
    salesGetReceiptsExternalByExternalNumberV2(externalNumber: string, body?: any, licenseNumber?: string): Promise<any>;
    salesGetReceiptsInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any>;
    salesGetReceiptsInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any>;
    salesGetTransactionsV1(body?: any, licenseNumber?: string): Promise<any>;
    salesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesDateStart: string, salesDateEnd: string, body?: any, licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryV1(body: SalesUpdateDeliveryV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryV2(body: SalesUpdateDeliveryV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryCompleteV1(body: SalesUpdateDeliveryCompleteV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryCompleteV2(body: SalesUpdateDeliveryCompleteV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryHubV1(body: SalesUpdateDeliveryHubV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryHubV2(body: SalesUpdateDeliveryHubV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryHubAcceptV1(body: SalesUpdateDeliveryHubAcceptV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryHubAcceptV2(body: SalesUpdateDeliveryHubAcceptV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryHubDepartV1(body: SalesUpdateDeliveryHubDepartV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryHubDepartV2(body: SalesUpdateDeliveryHubDepartV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryHubVerifyIdV1(body: SalesUpdateDeliveryHubVerifyIdV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryHubVerifyIdV2(body: SalesUpdateDeliveryHubVerifyIdV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryRetailerV1(body: SalesUpdateDeliveryRetailerV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateDeliveryRetailerV2(body: SalesUpdateDeliveryRetailerV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateReceiptV1(body: SalesUpdateReceiptV1RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateReceiptV2(body: SalesUpdateReceiptV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateReceiptFinalizeV2(body: SalesUpdateReceiptFinalizeV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateReceiptUnfinalizeV2(body: SalesUpdateReceiptUnfinalizeV2RequestItem[], licenseNumber?: string): Promise<any>;
    salesUpdateTransactionByDateV1(date: string, body: SalesUpdateTransactionByDateV1RequestItem[], licenseNumber?: string): Promise<any>;
}
export interface TransfersCreateExternalIncomingV1RequestItem {
    'Destinations'?: TransfersCreateExternalIncomingV1RequestItem_Destinations[];
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'ShipperAddress1'?: string;
    'ShipperAddress2'?: string;
    'ShipperAddressCity'?: string;
    'ShipperAddressPostalCode'?: string;
    'ShipperAddressState'?: string;
    'ShipperLicenseNumber'?: string;
    'ShipperMainPhoneNumber'?: string;
    'ShipperName'?: string;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersCreateExternalIncomingV1RequestItem_Destinations {
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'GrossUnitOfWeightId'?: number;
    'GrossWeight'?: number;
    'InvoiceNumber'?: string;
    'Packages'?: TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages[];
    'PlannedRoute'?: string;
    'RecipientLicenseNumber'?: string;
    'TransferTypeName'?: string;
    'Transporters'?: TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters[];
}
export interface TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages {
    'ExpirationDate'?: string;
    'ExternalId'?: string;
    'GrossUnitOfWeightName'?: string;
    'GrossWeight'?: number;
    'ItemName'?: string;
    'PackagedDate'?: string;
    'Quantity'?: number;
    'SellByDate'?: string;
    'UnitOfMeasureName'?: string;
    'UseByDate'?: string;
    'WholesalePrice'?: string;
}
export interface TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'IsLayover'?: boolean;
    'PhoneNumberForQuestions'?: string;
    'TransporterDetails'?: string;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersCreateExternalIncomingV2RequestItem {
    'Destinations'?: TransfersCreateExternalIncomingV2RequestItem_Destinations[];
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'ShipperAddress1'?: string;
    'ShipperAddress2'?: string;
    'ShipperAddressCity'?: string;
    'ShipperAddressPostalCode'?: string;
    'ShipperAddressState'?: string;
    'ShipperLicenseNumber'?: string;
    'ShipperMainPhoneNumber'?: string;
    'ShipperName'?: string;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersCreateExternalIncomingV2RequestItem_Destinations {
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'GrossUnitOfWeightId'?: number;
    'GrossWeight'?: number;
    'InvoiceNumber'?: string;
    'Packages'?: TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages[];
    'PlannedRoute'?: string;
    'RecipientLicenseNumber'?: string;
    'TransferTypeName'?: string;
    'Transporters'?: TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters[];
}
export interface TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages {
    'ExpirationDate'?: string;
    'ExternalId'?: string;
    'GrossUnitOfWeightName'?: string;
    'GrossWeight'?: number;
    'ItemName'?: string;
    'PackagedDate'?: string;
    'Quantity'?: number;
    'SellByDate'?: string;
    'UnitOfMeasureName'?: string;
    'UseByDate'?: string;
    'WholesalePrice'?: string;
}
export interface TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'IsLayover'?: boolean;
    'PhoneNumberForQuestions'?: string;
    'TransporterDetails'?: TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails[];
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersCreateTemplatesV1RequestItem {
    'Destinations'?: TransfersCreateTemplatesV1RequestItem_Destinations[];
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'Name'?: string;
    'PhoneNumberForQuestions'?: string;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersCreateTemplatesV1RequestItem_Destinations {
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'InvoiceNumber'?: string;
    'Packages'?: TransfersCreateTemplatesV1RequestItem_Destinations_Packages[];
    'PlannedRoute'?: string;
    'RecipientLicenseNumber'?: string;
    'TransferTypeName'?: string;
    'Transporters'?: TransfersCreateTemplatesV1RequestItem_Destinations_Transporters[];
}
export interface TransfersCreateTemplatesV1RequestItem_Destinations_Packages {
    'GrossUnitOfWeightName'?: string;
    'GrossWeight'?: number;
    'PackageLabel'?: string;
    'WholesalePrice'?: string;
}
export interface TransfersCreateTemplatesV1RequestItem_Destinations_Transporters {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'IsLayover'?: boolean;
    'PhoneNumberForQuestions'?: string;
    'TransporterDetails'?: string;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersCreateTemplatesOutgoingV2RequestItem {
    'Destinations'?: TransfersCreateTemplatesOutgoingV2RequestItem_Destinations[];
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'Name'?: string;
    'PhoneNumberForQuestions'?: string;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersCreateTemplatesOutgoingV2RequestItem_Destinations {
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'InvoiceNumber'?: string;
    'Packages'?: TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages[];
    'PlannedRoute'?: string;
    'RecipientLicenseNumber'?: string;
    'TransferTypeName'?: string;
    'Transporters'?: TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters[];
}
export interface TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages {
    'GrossUnitOfWeightName'?: string;
    'GrossWeight'?: number;
    'PackageLabel'?: string;
    'WholesalePrice'?: string;
}
export interface TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'IsLayover'?: boolean;
    'PhoneNumberForQuestions'?: string;
    'TransporterDetails'?: TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails[];
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersUpdateExternalIncomingV1RequestItem {
    'Destinations'?: TransfersUpdateExternalIncomingV1RequestItem_Destinations[];
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'ShipperAddress1'?: string;
    'ShipperAddress2'?: string;
    'ShipperAddressCity'?: string;
    'ShipperAddressPostalCode'?: string;
    'ShipperAddressState'?: string;
    'ShipperLicenseNumber'?: string;
    'ShipperMainPhoneNumber'?: string;
    'ShipperName'?: string;
    'TransferId'?: number;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersUpdateExternalIncomingV1RequestItem_Destinations {
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'GrossUnitOfWeightId'?: number;
    'GrossWeight'?: number;
    'InvoiceNumber'?: string;
    'Packages'?: TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages[];
    'PlannedRoute'?: string;
    'RecipientLicenseNumber'?: string;
    'TransferDestinationId'?: number;
    'TransferTypeName'?: string;
    'Transporters'?: TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters[];
}
export interface TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages {
    'ExpirationDate'?: string;
    'ExternalId'?: string;
    'GrossUnitOfWeightName'?: string;
    'GrossWeight'?: number;
    'ItemName'?: string;
    'PackagedDate'?: string;
    'Quantity'?: number;
    'SellByDate'?: string;
    'UnitOfMeasureName'?: string;
    'UseByDate'?: string;
    'WholesalePrice'?: string;
}
export interface TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'IsLayover'?: boolean;
    'PhoneNumberForQuestions'?: string;
    'TransporterDetails'?: string;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersUpdateExternalIncomingV2RequestItem {
    'Destinations'?: TransfersUpdateExternalIncomingV2RequestItem_Destinations[];
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'ShipperAddress1'?: string;
    'ShipperAddress2'?: string;
    'ShipperAddressCity'?: string;
    'ShipperAddressPostalCode'?: string;
    'ShipperAddressState'?: string;
    'ShipperLicenseNumber'?: string;
    'ShipperMainPhoneNumber'?: string;
    'ShipperName'?: string;
    'TransferId'?: number;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersUpdateExternalIncomingV2RequestItem_Destinations {
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'GrossUnitOfWeightId'?: number;
    'GrossWeight'?: number;
    'InvoiceNumber'?: string;
    'Packages'?: TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages[];
    'PlannedRoute'?: string;
    'RecipientLicenseNumber'?: string;
    'TransferDestinationId'?: number;
    'TransferTypeName'?: string;
    'Transporters'?: TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters[];
}
export interface TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages {
    'ExpirationDate'?: string;
    'ExternalId'?: string;
    'GrossUnitOfWeightName'?: string;
    'GrossWeight'?: number;
    'ItemName'?: string;
    'PackagedDate'?: string;
    'Quantity'?: number;
    'SellByDate'?: string;
    'UnitOfMeasureName'?: string;
    'UseByDate'?: string;
    'WholesalePrice'?: string;
}
export interface TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'IsLayover'?: boolean;
    'PhoneNumberForQuestions'?: string;
    'TransporterDetails'?: TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails[];
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersUpdateTemplatesV1RequestItem {
    'Destinations'?: TransfersUpdateTemplatesV1RequestItem_Destinations[];
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'Name'?: string;
    'PhoneNumberForQuestions'?: string;
    'TransferTemplateId'?: number;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersUpdateTemplatesV1RequestItem_Destinations {
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'InvoiceNumber'?: string;
    'Packages'?: TransfersUpdateTemplatesV1RequestItem_Destinations_Packages[];
    'PlannedRoute'?: string;
    'RecipientLicenseNumber'?: string;
    'TransferDestinationId'?: number;
    'TransferTypeName'?: string;
    'Transporters'?: TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters[];
}
export interface TransfersUpdateTemplatesV1RequestItem_Destinations_Packages {
    'GrossUnitOfWeightName'?: string;
    'GrossWeight'?: number;
    'PackageLabel'?: string;
    'WholesalePrice'?: string;
}
export interface TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'IsLayover'?: boolean;
    'PhoneNumberForQuestions'?: string;
    'TransporterDetails'?: string;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersUpdateTemplatesOutgoingV2RequestItem {
    'Destinations'?: TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations[];
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'Name'?: string;
    'PhoneNumberForQuestions'?: string;
    'TransferTemplateId'?: number;
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations {
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'InvoiceNumber'?: string;
    'Packages'?: TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages[];
    'PlannedRoute'?: string;
    'RecipientLicenseNumber'?: string;
    'TransferDestinationId'?: number;
    'TransferTypeName'?: string;
    'Transporters'?: TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters[];
}
export interface TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages {
    'GrossUnitOfWeightName'?: string;
    'GrossWeight'?: number;
    'PackageLabel'?: string;
    'WholesalePrice'?: string;
}
export interface TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'IsLayover'?: boolean;
    'PhoneNumberForQuestions'?: string;
    'TransporterDetails'?: TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails[];
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails {
    'DriverLayoverLeg'?: string;
    'DriverLicenseNumber'?: string;
    'DriverName'?: string;
    'DriverOccupationalLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface AdditivestemplatesCreateV2RequestItem {
    'ActiveIngredients'?: AdditivestemplatesCreateV2RequestItem_ActiveIngredients[];
    'AdditiveType'?: string;
    'ApplicationDevice'?: string;
    'EpaRegistrationNumber'?: string;
    'Name'?: string;
    'Note'?: string;
    'ProductSupplier'?: string;
    'ProductTradeName'?: string;
    'RestrictiveEntryIntervalQuantityDescription'?: string;
    'RestrictiveEntryIntervalTimeDescription'?: string;
}
export interface AdditivestemplatesCreateV2RequestItem_ActiveIngredients {
    'Name'?: string;
    'Percentage'?: number;
}
export interface AdditivestemplatesUpdateV2RequestItem {
    'ActiveIngredients'?: AdditivestemplatesUpdateV2RequestItem_ActiveIngredients[];
    'AdditiveType'?: string;
    'ApplicationDevice'?: string;
    'EpaRegistrationNumber'?: string;
    'Id'?: number;
    'Name'?: string;
    'Note'?: string;
    'ProductSupplier'?: string;
    'ProductTradeName'?: string;
    'RestrictiveEntryIntervalQuantityDescription'?: string;
    'RestrictiveEntryIntervalTimeDescription'?: string;
}
export interface AdditivestemplatesUpdateV2RequestItem_ActiveIngredients {
    'Name'?: string;
    'Percentage'?: number;
}
export interface HarvestsCreateFinishV1RequestItem {
    'ActualDate'?: string;
    'Id'?: number;
}
export interface HarvestsCreatePackageV1RequestItem {
    'ActualDate'?: string;
    'DecontaminateProduct'?: boolean;
    'DecontaminationDate'?: string;
    'DecontaminationSteps'?: string;
    'ExpirationDate'?: string;
    'Ingredients'?: HarvestsCreatePackageV1RequestItem_Ingredients[];
    'IsDonation'?: boolean;
    'IsProductionBatch'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'LabTestStageId'?: number;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'ProcessingJobTypeId'?: number;
    'ProductRequiresDecontamination'?: boolean;
    'ProductRequiresRemediation'?: boolean;
    'ProductionBatchNumber'?: string;
    'RemediateProduct'?: boolean;
    'RemediationDate'?: string;
    'RemediationMethodId'?: number;
    'RemediationSteps'?: string;
    'RequiredLabTestBatches'?: any[];
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UnitOfWeight'?: string;
    'UseByDate'?: string;
}
export interface HarvestsCreatePackageV1RequestItem_Ingredients {
    'HarvestId'?: number;
    'HarvestName'?: string;
    'UnitOfWeight'?: string;
    'Weight'?: number;
}
export interface HarvestsCreatePackageV2RequestItem {
    'ActualDate'?: string;
    'DecontaminateProduct'?: boolean;
    'DecontaminationDate'?: string;
    'DecontaminationSteps'?: string;
    'ExpirationDate'?: string;
    'Ingredients'?: HarvestsCreatePackageV2RequestItem_Ingredients[];
    'IsDonation'?: boolean;
    'IsProductionBatch'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'LabTestStageId'?: number;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'ProcessingJobTypeId'?: number;
    'ProductRequiresDecontamination'?: boolean;
    'ProductRequiresRemediation'?: boolean;
    'ProductionBatchNumber'?: string;
    'RemediateProduct'?: boolean;
    'RemediationDate'?: string;
    'RemediationMethodId'?: number;
    'RemediationSteps'?: string;
    'RequiredLabTestBatches'?: any[];
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UnitOfWeight'?: string;
    'UseByDate'?: string;
}
export interface HarvestsCreatePackageV2RequestItem_Ingredients {
    'HarvestId'?: number;
    'HarvestName'?: string;
    'UnitOfWeight'?: string;
    'Weight'?: number;
}
export interface HarvestsCreatePackageTestingV1RequestItem {
    'ActualDate'?: string;
    'DecontaminateProduct'?: boolean;
    'DecontaminationDate'?: string;
    'DecontaminationSteps'?: string;
    'ExpirationDate'?: string;
    'Ingredients'?: HarvestsCreatePackageTestingV1RequestItem_Ingredients[];
    'IsDonation'?: boolean;
    'IsProductionBatch'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'LabTestStageId'?: number;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'ProcessingJobTypeId'?: number;
    'ProductRequiresDecontamination'?: boolean;
    'ProductRequiresRemediation'?: boolean;
    'ProductionBatchNumber'?: string;
    'RemediateProduct'?: boolean;
    'RemediationDate'?: string;
    'RemediationMethodId'?: number;
    'RemediationSteps'?: string;
    'RequiredLabTestBatches'?: any[];
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UnitOfWeight'?: string;
    'UseByDate'?: string;
}
export interface HarvestsCreatePackageTestingV1RequestItem_Ingredients {
    'HarvestId'?: number;
    'HarvestName'?: string;
    'UnitOfWeight'?: string;
    'Weight'?: number;
}
export interface HarvestsCreatePackageTestingV2RequestItem {
    'ActualDate'?: string;
    'DecontaminateProduct'?: boolean;
    'DecontaminationDate'?: string;
    'DecontaminationSteps'?: string;
    'ExpirationDate'?: string;
    'Ingredients'?: HarvestsCreatePackageTestingV2RequestItem_Ingredients[];
    'IsDonation'?: boolean;
    'IsProductionBatch'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'LabTestStageId'?: number;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'ProcessingJobTypeId'?: number;
    'ProductRequiresDecontamination'?: boolean;
    'ProductRequiresRemediation'?: boolean;
    'ProductionBatchNumber'?: string;
    'RemediateProduct'?: boolean;
    'RemediationDate'?: string;
    'RemediationMethodId'?: number;
    'RemediationSteps'?: string;
    'RequiredLabTestBatches'?: any[];
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UnitOfWeight'?: string;
    'UseByDate'?: string;
}
export interface HarvestsCreatePackageTestingV2RequestItem_Ingredients {
    'HarvestId'?: number;
    'HarvestName'?: string;
    'UnitOfWeight'?: string;
    'Weight'?: number;
}
export interface HarvestsCreateRemoveWasteV1RequestItem {
    'ActualDate'?: string;
    'Id'?: number;
    'UnitOfWeight'?: string;
    'WasteType'?: string;
    'WasteWeight'?: number;
}
export interface HarvestsCreateUnfinishV1RequestItem {
    'Id'?: number;
}
export interface HarvestsCreateWasteV2RequestItem {
    'ActualDate'?: string;
    'Id'?: number;
    'UnitOfWeight'?: string;
    'WasteType'?: string;
    'WasteWeight'?: number;
}
export interface HarvestsUpdateFinishV2RequestItem {
    'ActualDate'?: string;
    'Id'?: number;
}
export interface HarvestsUpdateLocationV2RequestItem {
    'ActualDate'?: string;
    'DryingLocation'?: string;
    'DryingSublocation'?: string;
    'HarvestName'?: string;
    'Id'?: number;
}
export interface HarvestsUpdateMoveV1RequestItem {
    'ActualDate'?: string;
    'DryingLocation'?: string;
    'DryingSublocation'?: string;
    'HarvestName'?: string;
    'Id'?: number;
}
export interface HarvestsUpdateRenameV1RequestItem {
    'Id'?: number;
    'NewName'?: string;
    'OldName'?: string;
}
export interface HarvestsUpdateRenameV2RequestItem {
    'Id'?: number;
    'NewName'?: string;
    'OldName'?: string;
}
export interface HarvestsUpdateRestoreHarvestedPlantsV2RequestItem {
    'HarvestId'?: number;
    'PlantTags'?: string[];
}
export interface HarvestsUpdateUnfinishV2RequestItem {
    'Id'?: number;
}
export interface ItemsCreateV1RequestItem {
    'AdministrationMethod'?: string;
    'Allergens'?: string;
    'Description'?: string;
    'GlobalProductName'?: string;
    'ItemBrand'?: string;
    'ItemCategory'?: string;
    'ItemIngredients'?: string;
    'LabelImageFileSystemIds'?: string;
    'LabelPhotoDescription'?: string;
    'Name'?: string;
    'NumberOfDoses'?: string;
    'PackagingImageFileSystemIds'?: string;
    'PackagingPhotoDescription'?: string;
    'ProcessingJobCategoryName'?: string;
    'ProcessingJobTypeName'?: string;
    'ProductImageFileSystemIds'?: string;
    'ProductPDFFileSystemIds'?: string;
    'ProductPhotoDescription'?: string;
    'PublicIngredients'?: string;
    'ServingSize'?: string;
    'Strain'?: string;
    'SupplyDurationDays'?: number;
    'UnitCbdAContent'?: number;
    'UnitCbdAContentDose'?: number;
    'UnitCbdAContentDoseUnitOfMeasure'?: string;
    'UnitCbdAContentUnitOfMeasure'?: string;
    'UnitCbdAPercent'?: number;
    'UnitCbdContent'?: number;
    'UnitCbdContentDose'?: number;
    'UnitCbdContentDoseUnitOfMeasure'?: string;
    'UnitCbdContentUnitOfMeasure'?: string;
    'UnitCbdPercent'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcAContent'?: number;
    'UnitThcAContentDose'?: number;
    'UnitThcAContentDoseUnitOfMeasure'?: string;
    'UnitThcAContentUnitOfMeasure'?: string;
    'UnitThcAPercent'?: number;
    'UnitThcContent'?: number;
    'UnitThcContentDose'?: number;
    'UnitThcContentDoseUnitOfMeasure'?: string;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitVolume'?: number;
    'UnitVolumeUnitOfMeasure'?: string;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface ItemsCreateV2RequestItem {
    'AdministrationMethod'?: string;
    'Allergens'?: string;
    'Description'?: string;
    'GlobalProductName'?: string;
    'ItemBrand'?: string;
    'ItemCategory'?: string;
    'ItemIngredients'?: string;
    'LabelImageFileSystemIds'?: string;
    'LabelPhotoDescription'?: string;
    'Name'?: string;
    'NumberOfDoses'?: string;
    'PackagingImageFileSystemIds'?: string;
    'PackagingPhotoDescription'?: string;
    'ProcessingJobCategoryName'?: string;
    'ProcessingJobTypeName'?: string;
    'ProductImageFileSystemIds'?: string;
    'ProductPDFFileSystemIds'?: string;
    'ProductPhotoDescription'?: string;
    'PublicIngredients'?: string;
    'ServingSize'?: string;
    'Strain'?: string;
    'SupplyDurationDays'?: number;
    'UnitCbdAContent'?: number;
    'UnitCbdAContentDose'?: number;
    'UnitCbdAContentDoseUnitOfMeasure'?: string;
    'UnitCbdAContentUnitOfMeasure'?: string;
    'UnitCbdAPercent'?: number;
    'UnitCbdContent'?: number;
    'UnitCbdContentDose'?: number;
    'UnitCbdContentDoseUnitOfMeasure'?: string;
    'UnitCbdContentUnitOfMeasure'?: string;
    'UnitCbdPercent'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcAContent'?: number;
    'UnitThcAContentDose'?: number;
    'UnitThcAContentDoseUnitOfMeasure'?: string;
    'UnitThcAContentUnitOfMeasure'?: string;
    'UnitThcAPercent'?: number;
    'UnitThcContent'?: number;
    'UnitThcContentDose'?: number;
    'UnitThcContentDoseUnitOfMeasure'?: string;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitVolume'?: number;
    'UnitVolumeUnitOfMeasure'?: string;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface ItemsCreateBrandV2RequestItem {
    'Name'?: string;
}
export interface ItemsCreateFileV2RequestItem {
    'EncodedImageBase64'?: string;
    'FileName'?: string;
}
export interface ItemsCreatePhotoV1RequestItem {
    'EncodedImageBase64'?: string;
    'FileName'?: string;
}
export interface ItemsCreatePhotoV2RequestItem {
    'EncodedImageBase64'?: string;
    'FileName'?: string;
}
export interface ItemsCreateUpdateV1RequestItem {
    'AdministrationMethod'?: string;
    'Allergens'?: string;
    'Description'?: string;
    'GlobalProductName'?: string;
    'Id'?: number;
    'ItemBrand'?: string;
    'ItemCategory'?: string;
    'ItemIngredients'?: string;
    'LabelImageFileSystemIds'?: string;
    'LabelPhotoDescription'?: string;
    'Name'?: string;
    'NumberOfDoses'?: string;
    'PackagingImageFileSystemIds'?: string;
    'PackagingPhotoDescription'?: string;
    'ProcessingJobCategoryName'?: string;
    'ProcessingJobTypeName'?: string;
    'ProductImageFileSystemIds'?: string;
    'ProductPDFFileSystemIds'?: string;
    'ProductPhotoDescription'?: string;
    'PublicIngredients'?: string;
    'ServingSize'?: string;
    'Strain'?: string;
    'SupplyDurationDays'?: number;
    'UnitCbdAContent'?: number;
    'UnitCbdAContentDose'?: number;
    'UnitCbdAContentDoseUnitOfMeasure'?: string;
    'UnitCbdAContentUnitOfMeasure'?: string;
    'UnitCbdAPercent'?: number;
    'UnitCbdContent'?: number;
    'UnitCbdContentDose'?: number;
    'UnitCbdContentDoseUnitOfMeasure'?: string;
    'UnitCbdContentUnitOfMeasure'?: string;
    'UnitCbdPercent'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcAContent'?: number;
    'UnitThcAContentDose'?: number;
    'UnitThcAContentDoseUnitOfMeasure'?: string;
    'UnitThcAContentUnitOfMeasure'?: string;
    'UnitThcAPercent'?: number;
    'UnitThcContent'?: number;
    'UnitThcContentDose'?: number;
    'UnitThcContentDoseUnitOfMeasure'?: string;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitVolume'?: number;
    'UnitVolumeUnitOfMeasure'?: string;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface ItemsUpdateV2RequestItem {
    'AdministrationMethod'?: string;
    'Allergens'?: string;
    'Description'?: string;
    'GlobalProductName'?: string;
    'Id'?: number;
    'ItemBrand'?: string;
    'ItemCategory'?: string;
    'ItemIngredients'?: string;
    'LabelImageFileSystemIds'?: string;
    'LabelPhotoDescription'?: string;
    'Name'?: string;
    'NumberOfDoses'?: string;
    'PackagingImageFileSystemIds'?: string;
    'PackagingPhotoDescription'?: string;
    'ProcessingJobCategoryName'?: string;
    'ProcessingJobTypeName'?: string;
    'ProductImageFileSystemIds'?: string;
    'ProductPDFFileSystemIds'?: string;
    'ProductPhotoDescription'?: string;
    'PublicIngredients'?: string;
    'ServingSize'?: string;
    'Strain'?: string;
    'SupplyDurationDays'?: number;
    'UnitCbdAContent'?: number;
    'UnitCbdAContentDose'?: number;
    'UnitCbdAContentDoseUnitOfMeasure'?: string;
    'UnitCbdAContentUnitOfMeasure'?: string;
    'UnitCbdAPercent'?: number;
    'UnitCbdContent'?: number;
    'UnitCbdContentDose'?: number;
    'UnitCbdContentDoseUnitOfMeasure'?: string;
    'UnitCbdContentUnitOfMeasure'?: string;
    'UnitCbdPercent'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcAContent'?: number;
    'UnitThcAContentDose'?: number;
    'UnitThcAContentDoseUnitOfMeasure'?: string;
    'UnitThcAContentUnitOfMeasure'?: string;
    'UnitThcAPercent'?: number;
    'UnitThcContent'?: number;
    'UnitThcContentDose'?: number;
    'UnitThcContentDoseUnitOfMeasure'?: string;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitVolume'?: number;
    'UnitVolumeUnitOfMeasure'?: string;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface ItemsUpdateBrandV2RequestItem {
    'Id'?: number;
    'Name'?: string;
}
export interface LabtestsCreateRecordV1RequestItem {
    'DocumentFileBase64'?: string;
    'DocumentFileName'?: string;
    'Label'?: string;
    'ResultDate'?: string;
    'Results'?: LabtestsCreateRecordV1RequestItem_Results[];
}
export interface LabtestsCreateRecordV1RequestItem_Results {
    'LabTestTypeName'?: string;
    'Notes'?: string;
    'Passed'?: boolean;
    'Quantity'?: number;
}
export interface LabtestsCreateRecordV2RequestItem {
    'DocumentFileBase64'?: string;
    'DocumentFileName'?: string;
    'Label'?: string;
    'ResultDate'?: string;
    'Results'?: LabtestsCreateRecordV2RequestItem_Results[];
}
export interface LabtestsCreateRecordV2RequestItem_Results {
    'LabTestTypeName'?: string;
    'Notes'?: string;
    'Passed'?: boolean;
    'Quantity'?: number;
}
export interface LabtestsUpdateLabtestdocumentV1RequestItem {
    'DocumentFileBase64'?: string;
    'DocumentFileName'?: string;
    'LabTestResultId'?: number;
}
export interface LabtestsUpdateLabtestdocumentV2RequestItem {
    'DocumentFileBase64'?: string;
    'DocumentFileName'?: string;
    'LabTestResultId'?: number;
}
export interface LabtestsUpdateResultReleaseV1RequestItem {
    'PackageLabel'?: string;
}
export interface LabtestsUpdateResultReleaseV2RequestItem {
    'PackageLabel'?: string;
}
export interface StrainsCreateV1RequestItem {
    'CbdLevel'?: number;
    'IndicaPercentage'?: number;
    'Name'?: string;
    'SativaPercentage'?: number;
    'TestingStatus'?: string;
    'ThcLevel'?: number;
}
export interface StrainsCreateV2RequestItem {
    'CbdLevel'?: number;
    'IndicaPercentage'?: number;
    'Name'?: string;
    'SativaPercentage'?: number;
    'TestingStatus'?: string;
    'ThcLevel'?: number;
}
export interface StrainsCreateUpdateV1RequestItem {
    'CbdLevel'?: number;
    'Id'?: number;
    'IndicaPercentage'?: number;
    'Name'?: string;
    'SativaPercentage'?: number;
    'TestingStatus'?: string;
    'ThcLevel'?: number;
}
export interface StrainsUpdateV2RequestItem {
    'CbdLevel'?: number;
    'Id'?: number;
    'IndicaPercentage'?: number;
    'Name'?: string;
    'SativaPercentage'?: number;
    'TestingStatus'?: string;
    'ThcLevel'?: number;
}
export interface TransportersCreateDriverV2RequestItem {
    'DriversLicenseNumber'?: string;
    'EmployeeId'?: string;
    'Name'?: string;
}
export interface TransportersCreateVehicleV2RequestItem {
    'LicensePlateNumber'?: string;
    'Make'?: string;
    'Model'?: string;
}
export interface TransportersUpdateDriverV2RequestItem {
    'DriversLicenseNumber'?: string;
    'EmployeeId'?: string;
    'Id'?: string;
    'Name'?: string;
}
export interface TransportersUpdateVehicleV2RequestItem {
    'Id'?: string;
    'LicensePlateNumber'?: string;
    'Make'?: string;
    'Model'?: string;
}
export interface PackagesCreateV1RequestItem {
    'ActualDate'?: string;
    'ExpirationDate'?: string;
    'Ingredients'?: PackagesCreateV1RequestItem_Ingredients[];
    'IsDonation'?: boolean;
    'IsProductionBatch'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'LabTestStageId'?: number;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'ProcessingJobTypeId'?: number;
    'ProductRequiresRemediation'?: boolean;
    'ProductionBatchNumber'?: string;
    'Quantity'?: number;
    'RequiredLabTestBatches'?: boolean;
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UnitOfMeasure'?: string;
    'UseByDate'?: string;
    'UseSameItem'?: boolean;
}
export interface PackagesCreateV1RequestItem_Ingredients {
    'Package'?: string;
    'Quantity'?: number;
    'UnitOfMeasure'?: string;
}
export interface PackagesCreateV2RequestItem {
    'ActualDate'?: string;
    'ExpirationDate'?: string;
    'Ingredients'?: PackagesCreateV2RequestItem_Ingredients[];
    'IsDonation'?: boolean;
    'IsProductionBatch'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'LabTestStageId'?: number;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'ProcessingJobTypeId'?: number;
    'ProductRequiresRemediation'?: boolean;
    'ProductionBatchNumber'?: string;
    'Quantity'?: number;
    'RequiredLabTestBatches'?: boolean;
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UnitOfMeasure'?: string;
    'UseByDate'?: string;
    'UseSameItem'?: boolean;
}
export interface PackagesCreateV2RequestItem_Ingredients {
    'Package'?: string;
    'Quantity'?: number;
    'UnitOfMeasure'?: string;
}
export interface PackagesCreateAdjustV1RequestItem {
    'AdjustmentDate'?: string;
    'AdjustmentReason'?: string;
    'Label'?: string;
    'Quantity'?: number;
    'ReasonNote'?: string;
    'UnitOfMeasure'?: string;
}
export interface PackagesCreateAdjustV2RequestItem {
    'AdjustmentDate'?: string;
    'AdjustmentReason'?: string;
    'Label'?: string;
    'Quantity'?: number;
    'ReasonNote'?: string;
    'UnitOfMeasure'?: string;
}
export interface PackagesCreateChangeItemV1RequestItem {
    'Item'?: string;
    'Label'?: string;
}
export interface PackagesCreateChangeLocationV1RequestItem {
    'Label'?: string;
    'Location'?: string;
    'MoveDate'?: string;
    'Sublocation'?: string;
}
export interface PackagesCreateFinishV1RequestItem {
    'ActualDate'?: string;
    'Label'?: string;
}
export interface PackagesCreatePlantingsV1RequestItem {
    'LocationName'?: string;
    'PackageAdjustmentAmount'?: number;
    'PackageAdjustmentUnitOfMeasureName'?: string;
    'PackageLabel'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatchName'?: string;
    'PlantBatchType'?: string;
    'PlantCount'?: number;
    'PlantedDate'?: string;
    'StrainName'?: string;
    'SublocationName'?: string;
    'UnpackagedDate'?: string;
}
export interface PackagesCreatePlantingsV2RequestItem {
    'LocationName'?: string;
    'PackageAdjustmentAmount'?: number;
    'PackageAdjustmentUnitOfMeasureName'?: string;
    'PackageLabel'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatchName'?: string;
    'PlantBatchType'?: string;
    'PlantCount'?: number;
    'PlantedDate'?: string;
    'StrainName'?: string;
    'SublocationName'?: string;
    'UnpackagedDate'?: string;
}
export interface PackagesCreateRemediateV1RequestItem {
    'PackageLabel'?: string;
    'RemediationDate'?: string;
    'RemediationMethodName'?: string;
    'RemediationSteps'?: string;
}
export interface PackagesCreateTestingV1RequestItem {
    'ActualDate'?: string;
    'ExpirationDate'?: string;
    'Ingredients'?: PackagesCreateTestingV1RequestItem_Ingredients[];
    'IsDonation'?: boolean;
    'IsProductionBatch'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'LabTestStageId'?: number;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'ProcessingJobTypeId'?: number;
    'ProductRequiresRemediation'?: boolean;
    'ProductionBatchNumber'?: string;
    'Quantity'?: number;
    'RequiredLabTestBatches'?: boolean;
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UnitOfMeasure'?: string;
    'UseByDate'?: string;
    'UseSameItem'?: boolean;
}
export interface PackagesCreateTestingV1RequestItem_Ingredients {
    'Package'?: string;
    'Quantity'?: number;
    'UnitOfMeasure'?: string;
}
export interface PackagesCreateTestingV2RequestItem {
    'ActualDate'?: string;
    'ExpirationDate'?: string;
    'Ingredients'?: PackagesCreateTestingV2RequestItem_Ingredients[];
    'IsDonation'?: boolean;
    'IsProductionBatch'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'LabTestStageId'?: number;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'ProcessingJobTypeId'?: number;
    'ProductRequiresRemediation'?: boolean;
    'ProductionBatchNumber'?: string;
    'Quantity'?: number;
    'RequiredLabTestBatches'?: string[];
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UnitOfMeasure'?: string;
    'UseByDate'?: string;
    'UseSameItem'?: boolean;
}
export interface PackagesCreateTestingV2RequestItem_Ingredients {
    'Package'?: string;
    'Quantity'?: number;
    'UnitOfMeasure'?: string;
}
export interface PackagesCreateUnfinishV1RequestItem {
    'Label'?: string;
}
export interface PackagesUpdateAdjustV2RequestItem {
    'AdjustmentDate'?: string;
    'AdjustmentReason'?: string;
    'Label'?: string;
    'Quantity'?: number;
    'ReasonNote'?: string;
    'UnitOfMeasure'?: string;
}
export interface PackagesUpdateChangeNoteV1RequestItem {
    'Note'?: string;
    'PackageLabel'?: string;
}
export interface PackagesUpdateDecontaminateV2RequestItem {
    'DecontaminationDate'?: string;
    'DecontaminationMethodName'?: string;
    'DecontaminationSteps'?: string;
    'PackageLabel'?: string;
}
export interface PackagesUpdateDonationFlagV2RequestItem {
    'Label'?: string;
}
export interface PackagesUpdateDonationUnflagV2RequestItem {
    'Label'?: string;
}
export interface PackagesUpdateExternalidV2RequestItem {
    'ExternalId'?: string;
    'PackageLabel'?: string;
}
export interface PackagesUpdateFinishV2RequestItem {
    'ActualDate'?: string;
    'Label'?: string;
}
export interface PackagesUpdateItemV2RequestItem {
    'Item'?: string;
    'Label'?: string;
}
export interface PackagesUpdateLabTestRequiredV2RequestItem {
    'Label'?: string;
    'RequiredLabTestBatches'?: string[];
}
export interface PackagesUpdateLocationV2RequestItem {
    'Label'?: string;
    'Location'?: string;
    'MoveDate'?: string;
    'Sublocation'?: string;
}
export interface PackagesUpdateNoteV2RequestItem {
    'Note'?: string;
    'PackageLabel'?: string;
}
export interface PackagesUpdateRemediateV2RequestItem {
    'PackageLabel'?: string;
    'RemediationDate'?: string;
    'RemediationMethodName'?: string;
    'RemediationSteps'?: string;
}
export interface PackagesUpdateTradesampleFlagV2RequestItem {
    'Label'?: string;
}
export interface PackagesUpdateTradesampleUnflagV2RequestItem {
    'Label'?: string;
}
export interface PackagesUpdateUnfinishV2RequestItem {
    'Label'?: string;
}
export interface PackagesUpdateUsebydateV2RequestItem {
    'ExpirationDate'?: string;
    'Label'?: string;
    'SellByDate'?: string;
    'UseByDate'?: string;
}
export interface SublocationsCreateV2RequestItem {
    'Name'?: string;
}
export interface SublocationsUpdateV2RequestItem {
    'Id'?: number;
    'Name'?: string;
}
export interface PlantbatchesCreateAdditivesV1RequestItem {
    'ActiveIngredients'?: PlantbatchesCreateAdditivesV1RequestItem_ActiveIngredients[];
    'ActualDate'?: string;
    'AdditiveType'?: string;
    'ApplicationDevice'?: string;
    'EpaRegistrationNumber'?: string;
    'PlantBatchName'?: string;
    'ProductSupplier'?: string;
    'ProductTradeName'?: string;
    'TotalAmountApplied'?: number;
    'TotalAmountUnitOfMeasure'?: string;
}
export interface PlantbatchesCreateAdditivesV1RequestItem_ActiveIngredients {
    'Name'?: string;
    'Percentage'?: number;
}
export interface PlantbatchesCreateAdditivesV2RequestItem {
    'ActiveIngredients'?: PlantbatchesCreateAdditivesV2RequestItem_ActiveIngredients[];
    'ActualDate'?: string;
    'AdditiveType'?: string;
    'ApplicationDevice'?: string;
    'EpaRegistrationNumber'?: string;
    'PlantBatchName'?: string;
    'ProductSupplier'?: string;
    'ProductTradeName'?: string;
    'TotalAmountApplied'?: number;
    'TotalAmountUnitOfMeasure'?: string;
}
export interface PlantbatchesCreateAdditivesV2RequestItem_ActiveIngredients {
    'Name'?: string;
    'Percentage'?: number;
}
export interface PlantbatchesCreateAdditivesUsingtemplateV2RequestItem {
    'ActualDate'?: string;
    'AdditivesTemplateName'?: string;
    'PlantBatchName'?: string;
    'Rate'?: string;
    'TotalAmountApplied'?: number;
    'TotalAmountUnitOfMeasure'?: string;
    'Volume'?: string;
}
export interface PlantbatchesCreateAdjustV1RequestItem {
    'AdjustmentDate'?: string;
    'AdjustmentReason'?: string;
    'PlantBatchName'?: string;
    'Quantity'?: number;
    'ReasonNote'?: string;
}
export interface PlantbatchesCreateAdjustV2RequestItem {
    'AdjustmentDate'?: string;
    'AdjustmentReason'?: string;
    'PlantBatchName'?: string;
    'Quantity'?: number;
    'ReasonNote'?: string;
}
export interface PlantbatchesCreateChangegrowthphaseV1RequestItem {
    'Count'?: number;
    'CountPerPlant'?: string;
    'GrowthDate'?: string;
    'GrowthPhase'?: string;
    'Name'?: string;
    'NewLocation'?: string;
    'NewSublocation'?: string;
    'PatientLicenseNumber'?: string;
    'StartingTag'?: string;
}
export interface PlantbatchesCreateGrowthphaseV2RequestItem {
    'Count'?: number;
    'CountPerPlant'?: string;
    'GrowthDate'?: string;
    'GrowthPhase'?: string;
    'Name'?: string;
    'NewLocation'?: string;
    'NewSublocation'?: string;
    'PatientLicenseNumber'?: string;
    'StartingTag'?: string;
}
export interface PlantbatchesCreatePackageV2RequestItem {
    'ActualDate'?: string;
    'Count'?: number;
    'ExpirationDate'?: string;
    'Id'?: number;
    'IsDonation'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatch'?: string;
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UseByDate'?: string;
}
export interface PlantbatchesCreatePackageFrommotherplantV1RequestItem {
    'ActualDate'?: string;
    'Count'?: number;
    'ExpirationDate'?: string;
    'Id'?: number;
    'IsDonation'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatch'?: string;
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UseByDate'?: string;
}
export interface PlantbatchesCreatePackageFrommotherplantV2RequestItem {
    'ActualDate'?: string;
    'Count'?: number;
    'ExpirationDate'?: string;
    'Id'?: number;
    'IsDonation'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatch'?: string;
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UseByDate'?: string;
}
export interface PlantbatchesCreatePlantingsV2RequestItem {
    'ActualDate'?: string;
    'Count'?: number;
    'Location'?: string;
    'Name'?: string;
    'PatientLicenseNumber'?: string;
    'SourcePlantBatches'?: string;
    'Strain'?: string;
    'Sublocation'?: string;
    'Type'?: string;
}
export interface PlantbatchesCreateSplitV1RequestItem {
    'ActualDate'?: string;
    'Count'?: number;
    'GroupName'?: string;
    'Location'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatch'?: string;
    'Strain'?: string;
    'Sublocation'?: string;
}
export interface PlantbatchesCreateSplitV2RequestItem {
    'ActualDate'?: string;
    'Count'?: number;
    'GroupName'?: string;
    'Location'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatch'?: string;
    'Strain'?: string;
    'Sublocation'?: string;
}
export interface PlantbatchesCreateWasteV1RequestItem {
    'MixedMaterial'?: string;
    'Note'?: string;
    'PlantBatchName'?: string;
    'ReasonName'?: string;
    'UnitOfMeasureName'?: string;
    'WasteDate'?: string;
    'WasteMethodName'?: string;
    'WasteWeight'?: number;
}
export interface PlantbatchesCreateWasteV2RequestItem {
    'MixedMaterial'?: string;
    'Note'?: string;
    'PlantBatchName'?: string;
    'ReasonName'?: string;
    'UnitOfMeasureName'?: string;
    'WasteDate'?: string;
    'WasteMethodName'?: string;
    'WasteWeight'?: number;
}
export interface PlantbatchesCreatepackagesV1RequestItem {
    'ActualDate'?: string;
    'Count'?: number;
    'ExpirationDate'?: string;
    'Id'?: number;
    'IsDonation'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'Location'?: string;
    'Note'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatch'?: string;
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UseByDate'?: string;
}
export interface PlantbatchesCreateplantingsV1RequestItem {
    'ActualDate'?: string;
    'Count'?: number;
    'Location'?: string;
    'Name'?: string;
    'PatientLicenseNumber'?: string;
    'SourcePlantBatches'?: string;
    'Strain'?: string;
    'Sublocation'?: string;
    'Type'?: string;
}
export interface PlantbatchesUpdateLocationV2RequestItem {
    'Location'?: string;
    'MoveDate'?: string;
    'Name'?: string;
    'Sublocation'?: string;
}
export interface PlantbatchesUpdateMoveplantbatchesV1RequestItem {
    'Location'?: string;
    'MoveDate'?: string;
    'Name'?: string;
    'Sublocation'?: string;
}
export interface PlantbatchesUpdateNameV2RequestItem {
    'Group'?: string;
    'Id'?: number;
    'NewGroup'?: string;
}
export interface PlantbatchesUpdateStrainV2RequestItem {
    'Id'?: number;
    'Name'?: string;
    'StrainId'?: number;
    'StrainName'?: string;
}
export interface PlantbatchesUpdateTagV2RequestItem {
    'Group'?: string;
    'Id'?: number;
    'NewTag'?: string;
    'ReplaceDate'?: string;
    'TagId'?: number;
}
export interface PlantsCreateAdditivesV1RequestItem {
    'ActiveIngredients'?: PlantsCreateAdditivesV1RequestItem_ActiveIngredients[];
    'ActualDate'?: string;
    'AdditiveType'?: string;
    'ApplicationDevice'?: string;
    'EpaRegistrationNumber'?: string;
    'PlantLabels'?: string[];
    'ProductSupplier'?: string;
    'ProductTradeName'?: string;
    'TotalAmountApplied'?: number;
    'TotalAmountUnitOfMeasure'?: string;
}
export interface PlantsCreateAdditivesV1RequestItem_ActiveIngredients {
    'Name'?: string;
    'Percentage'?: number;
}
export interface PlantsCreateAdditivesV2RequestItem {
    'ActiveIngredients'?: PlantsCreateAdditivesV2RequestItem_ActiveIngredients[];
    'ActualDate'?: string;
    'AdditiveType'?: string;
    'ApplicationDevice'?: string;
    'EpaRegistrationNumber'?: string;
    'PlantLabels'?: string[];
    'ProductSupplier'?: string;
    'ProductTradeName'?: string;
    'TotalAmountApplied'?: number;
    'TotalAmountUnitOfMeasure'?: string;
}
export interface PlantsCreateAdditivesV2RequestItem_ActiveIngredients {
    'Name'?: string;
    'Percentage'?: number;
}
export interface PlantsCreateAdditivesBylocationV1RequestItem {
    'ActiveIngredients'?: PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients[];
    'ActualDate'?: string;
    'AdditiveType'?: string;
    'ApplicationDevice'?: string;
    'EpaRegistrationNumber'?: string;
    'LocationName'?: string;
    'ProductSupplier'?: string;
    'ProductTradeName'?: string;
    'SublocationName'?: string;
    'TotalAmountApplied'?: number;
    'TotalAmountUnitOfMeasure'?: string;
}
export interface PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients {
    'Name'?: string;
    'Percentage'?: number;
}
export interface PlantsCreateAdditivesBylocationV2RequestItem {
    'ActiveIngredients'?: PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients[];
    'ActualDate'?: string;
    'AdditiveType'?: string;
    'ApplicationDevice'?: string;
    'EpaRegistrationNumber'?: string;
    'LocationName'?: string;
    'ProductSupplier'?: string;
    'ProductTradeName'?: string;
    'SublocationName'?: string;
    'TotalAmountApplied'?: number;
    'TotalAmountUnitOfMeasure'?: string;
}
export interface PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients {
    'Name'?: string;
    'Percentage'?: number;
}
export interface PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem {
    'ActualDate'?: string;
    'AdditivesTemplateName'?: string;
    'LocationName'?: string;
    'Rate'?: string;
    'SublocationName'?: string;
    'TotalAmountApplied'?: number;
    'TotalAmountUnitOfMeasure'?: string;
    'Volume'?: string;
}
export interface PlantsCreateAdditivesUsingtemplateV2RequestItem {
    'ActualDate'?: string;
    'AdditivesTemplateName'?: string;
    'PlantLabels'?: string[];
    'Rate'?: string;
    'TotalAmountApplied'?: number;
    'TotalAmountUnitOfMeasure'?: string;
    'Volume'?: string;
}
export interface PlantsCreateChangegrowthphasesV1RequestItem {
    'GrowthDate'?: string;
    'GrowthPhase'?: string;
    'Id'?: number;
    'Label'?: string;
    'NewLocation'?: string;
    'NewSublocation'?: string;
    'NewTag'?: string;
}
export interface PlantsCreateHarvestplantsV1RequestItem {
    'ActualDate'?: string;
    'DryingLocation'?: string;
    'DryingSublocation'?: string;
    'HarvestName'?: string;
    'PatientLicenseNumber'?: string;
    'Plant'?: string;
    'UnitOfWeight'?: string;
    'Weight'?: number;
}
export interface PlantsCreateManicureV2RequestItem {
    'ActualDate'?: string;
    'DryingLocation'?: string;
    'DryingSublocation'?: string;
    'HarvestName'?: string;
    'PatientLicenseNumber'?: string;
    'Plant'?: string;
    'PlantCount'?: number;
    'UnitOfWeight'?: string;
    'Weight'?: number;
}
export interface PlantsCreateManicureplantsV1RequestItem {
    'ActualDate'?: string;
    'DryingLocation'?: string;
    'DryingSublocation'?: string;
    'HarvestName'?: string;
    'PatientLicenseNumber'?: string;
    'Plant'?: string;
    'PlantCount'?: number;
    'UnitOfWeight'?: string;
    'Weight'?: number;
}
export interface PlantsCreateMoveplantsV1RequestItem {
    'ActualDate'?: string;
    'Id'?: number;
    'Label'?: string;
    'Location'?: string;
    'Sublocation'?: string;
}
export interface PlantsCreatePlantbatchPackageV1RequestItem {
    'ActualDate'?: string;
    'Count'?: number;
    'IsDonation'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'Location'?: string;
    'Note'?: string;
    'PackageTag'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatchType'?: string;
    'PlantLabel'?: string;
    'Sublocation'?: string;
}
export interface PlantsCreatePlantbatchPackageV2RequestItem {
    'ActualDate'?: string;
    'Count'?: number;
    'IsDonation'?: boolean;
    'IsTradeSample'?: boolean;
    'Item'?: string;
    'Location'?: string;
    'Note'?: string;
    'PackageTag'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatchType'?: string;
    'PlantLabel'?: string;
    'Sublocation'?: string;
}
export interface PlantsCreatePlantingsV1RequestItem {
    'ActualDate'?: string;
    'LocationName'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatchName'?: string;
    'PlantBatchType'?: string;
    'PlantCount'?: number;
    'PlantLabel'?: string;
    'StrainName'?: string;
    'SublocationName'?: string;
}
export interface PlantsCreatePlantingsV2RequestItem {
    'ActualDate'?: string;
    'LocationName'?: string;
    'PatientLicenseNumber'?: string;
    'PlantBatchName'?: string;
    'PlantBatchType'?: string;
    'PlantCount'?: number;
    'PlantLabel'?: string;
    'StrainName'?: string;
    'SublocationName'?: string;
}
export interface PlantsCreateWasteV1RequestItem {
    'LocationName'?: string;
    'MixedMaterial'?: string;
    'Note'?: string;
    'PlantLabels'?: any[];
    'ReasonName'?: string;
    'SublocationName'?: string;
    'UnitOfMeasureName'?: string;
    'WasteDate'?: string;
    'WasteMethodName'?: string;
    'WasteWeight'?: number;
}
export interface PlantsCreateWasteV2RequestItem {
    'LocationName'?: string;
    'MixedMaterial'?: string;
    'Note'?: string;
    'PlantLabels'?: any[];
    'ReasonName'?: string;
    'SublocationName'?: string;
    'UnitOfMeasureName'?: string;
    'WasteDate'?: string;
    'WasteMethodName'?: string;
    'WasteWeight'?: number;
}
export interface PlantsUpdateAdjustV2RequestItem {
    'AdjustCount'?: number;
    'AdjustReason'?: string;
    'AdjustmentDate'?: string;
    'Id'?: number;
    'Label'?: string;
    'ReasonNote'?: string;
}
export interface PlantsUpdateGrowthphaseV2RequestItem {
    'GrowthDate'?: string;
    'GrowthPhase'?: string;
    'Id'?: number;
    'Label'?: string;
    'NewLocation'?: string;
    'NewSublocation'?: string;
    'NewTag'?: string;
}
export interface PlantsUpdateHarvestV2RequestItem {
    'ActualDate'?: string;
    'DryingLocation'?: string;
    'DryingSublocation'?: string;
    'HarvestName'?: string;
    'PatientLicenseNumber'?: string;
    'Plant'?: string;
    'UnitOfWeight'?: string;
    'Weight'?: number;
}
export interface PlantsUpdateLocationV2RequestItem {
    'ActualDate'?: string;
    'Id'?: number;
    'Label'?: string;
    'Location'?: string;
    'Sublocation'?: string;
}
export interface PlantsUpdateMergeV2RequestItem {
    'MergeDate'?: string;
    'SourcePlantGroupLabel'?: string;
    'TargetPlantGroupLabel'?: string;
}
export interface PlantsUpdateSplitV2RequestItem {
    'PlantCount'?: number;
    'SourcePlantLabel'?: string;
    'SplitDate'?: string;
    'StrainLabel'?: string;
    'TagLabel'?: string;
}
export interface PlantsUpdateStrainV2RequestItem {
    'Id'?: number;
    'Label'?: string;
    'StrainId'?: number;
    'StrainName'?: string;
}
export interface PlantsUpdateTagV2RequestItem {
    'Id'?: number;
    'Label'?: string;
    'NewTag'?: string;
    'ReplaceDate'?: string;
    'TagId'?: number;
}
export interface PatientcheckinsCreateV1RequestItem {
    'CheckInDate'?: string;
    'CheckInLocationId'?: number;
    'PatientLicenseNumber'?: string;
    'RegistrationExpiryDate'?: string;
    'RegistrationStartDate'?: string;
}
export interface PatientcheckinsCreateV2RequestItem {
    'CheckInDate'?: string;
    'CheckInLocationId'?: number;
    'PatientLicenseNumber'?: string;
    'RegistrationExpiryDate'?: string;
    'RegistrationStartDate'?: string;
}
export interface PatientcheckinsUpdateV1RequestItem {
    'CheckInDate'?: string;
    'CheckInLocationId'?: number;
    'Id'?: number;
    'PatientLicenseNumber'?: string;
    'RegistrationExpiryDate'?: string;
    'RegistrationStartDate'?: string;
}
export interface PatientcheckinsUpdateV2RequestItem {
    'CheckInDate'?: string;
    'CheckInLocationId'?: number;
    'Id'?: number;
    'PatientLicenseNumber'?: string;
    'RegistrationExpiryDate'?: string;
    'RegistrationStartDate'?: string;
}
export interface ProcessingjobsCreateAdjustV1RequestItem {
    'AdjustmentDate'?: string;
    'AdjustmentNote'?: string;
    'AdjustmentReason'?: string;
    'CountUnitOfMeasureName'?: string;
    'Id'?: number;
    'Packages'?: ProcessingjobsCreateAdjustV1RequestItem_Packages[];
    'VolumeUnitOfMeasureName'?: string;
    'WeightUnitOfMeasureName'?: string;
}
export interface ProcessingjobsCreateAdjustV1RequestItem_Packages {
    'Label'?: string;
    'Quantity'?: number;
    'UnitOfMeasure'?: string;
}
export interface ProcessingjobsCreateAdjustV2RequestItem {
    'AdjustmentDate'?: string;
    'AdjustmentNote'?: string;
    'AdjustmentReason'?: string;
    'CountUnitOfMeasureName'?: string;
    'Id'?: number;
    'Packages'?: ProcessingjobsCreateAdjustV2RequestItem_Packages[];
    'VolumeUnitOfMeasureName'?: string;
    'WeightUnitOfMeasureName'?: string;
}
export interface ProcessingjobsCreateAdjustV2RequestItem_Packages {
    'Label'?: string;
    'Quantity'?: number;
    'UnitOfMeasure'?: string;
}
export interface ProcessingjobsCreateJobtypesV1RequestItem {
    'Attributes'?: string[];
    'Category'?: string;
    'Description'?: string;
    'Name'?: string;
    'ProcessingSteps'?: string;
}
export interface ProcessingjobsCreateJobtypesV2RequestItem {
    'Attributes'?: string[];
    'Category'?: string;
    'Description'?: string;
    'Name'?: string;
    'ProcessingSteps'?: string;
}
export interface ProcessingjobsCreateStartV1RequestItem {
    'CountUnitOfMeasure'?: string;
    'JobName'?: string;
    'JobType'?: string;
    'Packages'?: ProcessingjobsCreateStartV1RequestItem_Packages[];
    'StartDate'?: string;
    'VolumeUnitOfMeasure'?: string;
    'WeightUnitOfMeasure'?: string;
}
export interface ProcessingjobsCreateStartV1RequestItem_Packages {
    'Label'?: string;
    'Quantity'?: number;
    'UnitOfMeasure'?: string;
}
export interface ProcessingjobsCreateStartV2RequestItem {
    'CountUnitOfMeasure'?: string;
    'JobName'?: string;
    'JobType'?: string;
    'Packages'?: ProcessingjobsCreateStartV2RequestItem_Packages[];
    'StartDate'?: string;
    'VolumeUnitOfMeasure'?: string;
    'WeightUnitOfMeasure'?: string;
}
export interface ProcessingjobsCreateStartV2RequestItem_Packages {
    'Label'?: string;
    'Quantity'?: number;
    'UnitOfMeasure'?: string;
}
export interface ProcessingjobsCreatepackagesV1RequestItem {
    'ExpirationDate'?: string;
    'FinishDate'?: string;
    'FinishNote'?: string;
    'FinishProcessingJob'?: boolean;
    'Item'?: string;
    'JobName'?: string;
    'Location'?: string;
    'Note'?: string;
    'PackageDate'?: string;
    'PatientLicenseNumber'?: string;
    'ProductionBatchNumber'?: string;
    'Quantity'?: number;
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UnitOfMeasure'?: string;
    'UseByDate'?: string;
    'WasteCountQuantity'?: string;
    'WasteCountUnitOfMeasureName'?: string;
    'WasteVolumeQuantity'?: string;
    'WasteVolumeUnitOfMeasureName'?: string;
    'WasteWeightQuantity'?: string;
    'WasteWeightUnitOfMeasureName'?: string;
}
export interface ProcessingjobsCreatepackagesV2RequestItem {
    'ExpirationDate'?: string;
    'FinishDate'?: string;
    'FinishNote'?: string;
    'FinishProcessingJob'?: boolean;
    'Item'?: string;
    'JobName'?: string;
    'Location'?: string;
    'Note'?: string;
    'PackageDate'?: string;
    'PatientLicenseNumber'?: string;
    'ProductionBatchNumber'?: string;
    'Quantity'?: number;
    'SellByDate'?: string;
    'Sublocation'?: string;
    'Tag'?: string;
    'UnitOfMeasure'?: string;
    'UseByDate'?: string;
    'WasteCountQuantity'?: string;
    'WasteCountUnitOfMeasureName'?: string;
    'WasteVolumeQuantity'?: string;
    'WasteVolumeUnitOfMeasureName'?: string;
    'WasteWeightQuantity'?: string;
    'WasteWeightUnitOfMeasureName'?: string;
}
export interface ProcessingjobsUpdateFinishV1RequestItem {
    'FinishDate'?: string;
    'FinishNote'?: string;
    'Id'?: number;
    'TotalCountWaste'?: string;
    'TotalVolumeWaste'?: string;
    'TotalWeightWaste'?: number;
    'WasteCountUnitOfMeasureName'?: string;
    'WasteVolumeUnitOfMeasureName'?: string;
    'WasteWeightUnitOfMeasureName'?: string;
}
export interface ProcessingjobsUpdateFinishV2RequestItem {
    'FinishDate'?: string;
    'FinishNote'?: string;
    'Id'?: number;
    'TotalCountWaste'?: string;
    'TotalVolumeWaste'?: string;
    'TotalWeightWaste'?: number;
    'WasteCountUnitOfMeasureName'?: string;
    'WasteVolumeUnitOfMeasureName'?: string;
    'WasteWeightUnitOfMeasureName'?: string;
}
export interface ProcessingjobsUpdateJobtypesV1RequestItem {
    'Attributes'?: string[];
    'CategoryName'?: string;
    'Description'?: string;
    'Id'?: number;
    'Name'?: string;
    'ProcessingSteps'?: string;
}
export interface ProcessingjobsUpdateJobtypesV2RequestItem {
    'Attributes'?: string[];
    'CategoryName'?: string;
    'Description'?: string;
    'Id'?: number;
    'Name'?: string;
    'ProcessingSteps'?: string;
}
export interface ProcessingjobsUpdateUnfinishV1RequestItem {
    'Id'?: number;
}
export interface ProcessingjobsUpdateUnfinishV2RequestItem {
    'Id'?: number;
}
export interface RetailidCreateAssociateV2RequestItem {
    'PackageLabel'?: string;
    'QrUrls'?: string[];
}
export interface RetailidCreateGenerateV2Request {
    'PackageLabel'?: string;
    'Quantity'?: number;
}
export interface RetailidCreateMergeV2Request {
    'packageLabels'?: string[];
}
export interface RetailidCreatePackageInfoV2Request {
    'packageLabels'?: string[];
}
export interface LocationsCreateV1RequestItem {
    'LocationTypeName'?: string;
    'Name'?: string;
}
export interface LocationsCreateV2RequestItem {
    'LocationTypeName'?: string;
    'Name'?: string;
}
export interface LocationsCreateUpdateV1RequestItem {
    'Id'?: number;
    'LocationTypeName'?: string;
    'Name'?: string;
}
export interface LocationsUpdateV2RequestItem {
    'Id'?: number;
    'LocationTypeName'?: string;
    'Name'?: string;
}
export interface PatientsCreateV2RequestItem {
    'ActualDate'?: string;
    'ConcentrateOuncesAllowed'?: number;
    'FlowerOuncesAllowed'?: number;
    'HasSalesLimitExemption'?: boolean;
    'InfusedOuncesAllowed'?: number;
    'LicenseEffectiveEndDate'?: string;
    'LicenseEffectiveStartDate'?: string;
    'LicenseNumber'?: string;
    'MaxConcentrateThcPercentAllowed'?: number;
    'MaxFlowerThcPercentAllowed'?: number;
    'RecommendedPlants'?: number;
    'RecommendedSmokableQuantity'?: number;
    'ThcOuncesAllowed'?: number;
}
export interface PatientsCreateAddV1RequestItem {
    'ActualDate'?: string;
    'ConcentrateOuncesAllowed'?: number;
    'FlowerOuncesAllowed'?: number;
    'HasSalesLimitExemption'?: boolean;
    'InfusedOuncesAllowed'?: number;
    'LicenseEffectiveEndDate'?: string;
    'LicenseEffectiveStartDate'?: string;
    'LicenseNumber'?: string;
    'MaxConcentrateThcPercentAllowed'?: number;
    'MaxFlowerThcPercentAllowed'?: number;
    'RecommendedPlants'?: number;
    'RecommendedSmokableQuantity'?: number;
    'ThcOuncesAllowed'?: number;
}
export interface PatientsCreateUpdateV1RequestItem {
    'ActualDate'?: string;
    'ConcentrateOuncesAllowed'?: number;
    'FlowerOuncesAllowed'?: number;
    'HasSalesLimitExemption'?: boolean;
    'InfusedOuncesAllowed'?: number;
    'LicenseEffectiveEndDate'?: string;
    'LicenseEffectiveStartDate'?: string;
    'LicenseNumber'?: string;
    'MaxConcentrateThcPercentAllowed'?: number;
    'MaxFlowerThcPercentAllowed'?: number;
    'NewLicenseNumber'?: string;
    'RecommendedPlants'?: number;
    'RecommendedSmokableQuantity'?: number;
    'ThcOuncesAllowed'?: number;
}
export interface PatientsUpdateV2RequestItem {
    'ActualDate'?: string;
    'ConcentrateOuncesAllowed'?: number;
    'FlowerOuncesAllowed'?: number;
    'HasSalesLimitExemption'?: boolean;
    'InfusedOuncesAllowed'?: number;
    'LicenseEffectiveEndDate'?: string;
    'LicenseEffectiveStartDate'?: string;
    'LicenseNumber'?: string;
    'MaxConcentrateThcPercentAllowed'?: number;
    'MaxFlowerThcPercentAllowed'?: number;
    'NewLicenseNumber'?: string;
    'RecommendedPlants'?: number;
    'RecommendedSmokableQuantity'?: number;
    'ThcOuncesAllowed'?: number;
}
export interface SalesCreateDeliveryV1RequestItem {
    'ConsumerId'?: number;
    'DriverEmployeeId'?: string;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'PatientLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'PlannedRoute'?: string;
    'RecipientAddressCity'?: string;
    'RecipientAddressCounty'?: string;
    'RecipientAddressPostalCode'?: string;
    'RecipientAddressState'?: string;
    'RecipientAddressStreet1'?: string;
    'RecipientAddressStreet2'?: string;
    'RecipientName'?: string;
    'RecipientZoneId'?: number;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesCreateDeliveryV1RequestItem_Transactions[];
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesCreateDeliveryV1RequestItem_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesCreateDeliveryV2RequestItem {
    'ConsumerId'?: number;
    'DriverEmployeeId'?: string;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'PatientLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'PlannedRoute'?: string;
    'RecipientAddressCity'?: string;
    'RecipientAddressCounty'?: string;
    'RecipientAddressPostalCode'?: string;
    'RecipientAddressState'?: string;
    'RecipientAddressStreet1'?: string;
    'RecipientAddressStreet2'?: string;
    'RecipientName'?: string;
    'RecipientZoneId'?: number;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesCreateDeliveryV2RequestItem_Transactions[];
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesCreateDeliveryV2RequestItem_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesCreateDeliveryRetailerV1RequestItem {
    'DateTime'?: string;
    'Destinations'?: SalesCreateDeliveryRetailerV1RequestItem_Destinations[];
    'DriverEmployeeId'?: string;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Packages'?: SalesCreateDeliveryRetailerV1RequestItem_Packages[];
    'PhoneNumberForQuestions'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesCreateDeliveryRetailerV1RequestItem_Destinations {
    'ConsumerId'?: string;
    'EstimatedArrivalDateTime'?: string;
    'PatientLicenseNumber'?: string;
    'RecipientAddressCity'?: string;
    'RecipientAddressCounty'?: string;
    'RecipientAddressPostalCode'?: string;
    'RecipientAddressState'?: string;
    'RecipientAddressStreet1'?: string;
    'RecipientAddressStreet2'?: string;
    'RecipientName'?: string;
    'RecipientZoneId'?: string;
    'SalesCustomerType'?: string;
    'Transactions'?: SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions[];
}
export interface SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesCreateDeliveryRetailerV1RequestItem_Packages {
    'DateTime'?: string;
    'PackageLabel'?: string;
    'Quantity'?: number;
    'TotalPrice'?: number;
    'UnitOfMeasure'?: string;
}
export interface SalesCreateDeliveryRetailerV2RequestItem {
    'DateTime'?: string;
    'Destinations'?: SalesCreateDeliveryRetailerV2RequestItem_Destinations[];
    'DriverEmployeeId'?: string;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Packages'?: SalesCreateDeliveryRetailerV2RequestItem_Packages[];
    'PhoneNumberForQuestions'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesCreateDeliveryRetailerV2RequestItem_Destinations {
    'ConsumerId'?: string;
    'EstimatedArrivalDateTime'?: string;
    'PatientLicenseNumber'?: string;
    'RecipientAddressCity'?: string;
    'RecipientAddressCounty'?: string;
    'RecipientAddressPostalCode'?: string;
    'RecipientAddressState'?: string;
    'RecipientAddressStreet1'?: string;
    'RecipientAddressStreet2'?: string;
    'RecipientName'?: string;
    'RecipientZoneId'?: string;
    'SalesCustomerType'?: string;
    'Transactions'?: SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions[];
}
export interface SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesCreateDeliveryRetailerV2RequestItem_Packages {
    'DateTime'?: string;
    'PackageLabel'?: string;
    'Quantity'?: number;
    'TotalPrice'?: number;
    'UnitOfMeasure'?: string;
}
export interface SalesCreateDeliveryRetailerDepartV1RequestItem {
    'RetailerDeliveryId'?: number;
}
export interface SalesCreateDeliveryRetailerDepartV2RequestItem {
    'RetailerDeliveryId'?: number;
}
export interface SalesCreateDeliveryRetailerEndV1RequestItem {
    'ActualArrivalDateTime'?: string;
    'Packages'?: SalesCreateDeliveryRetailerEndV1RequestItem_Packages[];
    'RetailerDeliveryId'?: number;
}
export interface SalesCreateDeliveryRetailerEndV1RequestItem_Packages {
    'EndQuantity'?: number;
    'EndUnitOfMeasure'?: string;
    'Label'?: string;
}
export interface SalesCreateDeliveryRetailerEndV2RequestItem {
    'ActualArrivalDateTime'?: string;
    'Packages'?: SalesCreateDeliveryRetailerEndV2RequestItem_Packages[];
    'RetailerDeliveryId'?: number;
}
export interface SalesCreateDeliveryRetailerEndV2RequestItem_Packages {
    'EndQuantity'?: number;
    'EndUnitOfMeasure'?: string;
    'Label'?: string;
}
export interface SalesCreateDeliveryRetailerRestockV1RequestItem {
    'DateTime'?: string;
    'Destinations'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Packages'?: SalesCreateDeliveryRetailerRestockV1RequestItem_Packages[];
    'RetailerDeliveryId'?: number;
}
export interface SalesCreateDeliveryRetailerRestockV1RequestItem_Packages {
    'PackageLabel'?: string;
    'Quantity'?: number;
    'RemoveCurrentPackage'?: boolean;
    'TotalPrice'?: number;
    'UnitOfMeasure'?: string;
}
export interface SalesCreateDeliveryRetailerRestockV2RequestItem {
    'DateTime'?: string;
    'Destinations'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Packages'?: SalesCreateDeliveryRetailerRestockV2RequestItem_Packages[];
    'RetailerDeliveryId'?: number;
}
export interface SalesCreateDeliveryRetailerRestockV2RequestItem_Packages {
    'PackageLabel'?: string;
    'Quantity'?: number;
    'RemoveCurrentPackage'?: boolean;
    'TotalPrice'?: number;
    'UnitOfMeasure'?: string;
}
export interface SalesCreateDeliveryRetailerSaleV1RequestItem {
    'ConsumerId'?: number;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'PatientLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'PlannedRoute'?: string;
    'RecipientAddressCity'?: string;
    'RecipientAddressCounty'?: string;
    'RecipientAddressPostalCode'?: string;
    'RecipientAddressState'?: string;
    'RecipientAddressStreet1'?: string;
    'RecipientAddressStreet2'?: string;
    'RecipientName'?: string;
    'RecipientZoneId'?: number;
    'RetailerDeliveryId'?: number;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions[];
}
export interface SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesCreateDeliveryRetailerSaleV2RequestItem {
    'ConsumerId'?: number;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'PatientLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'PlannedRoute'?: string;
    'RecipientAddressCity'?: string;
    'RecipientAddressCounty'?: string;
    'RecipientAddressPostalCode'?: string;
    'RecipientAddressState'?: string;
    'RecipientAddressStreet1'?: string;
    'RecipientAddressStreet2'?: string;
    'RecipientName'?: string;
    'RecipientZoneId'?: number;
    'RetailerDeliveryId'?: number;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions[];
}
export interface SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesCreateReceiptV1RequestItem {
    'CaregiverLicenseNumber'?: string;
    'ExternalReceiptNumber'?: string;
    'IdentificationMethod'?: string;
    'PatientLicenseNumber'?: string;
    'PatientRegistrationLocationId'?: number;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesCreateReceiptV1RequestItem_Transactions[];
}
export interface SalesCreateReceiptV1RequestItem_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesCreateReceiptV2RequestItem {
    'CaregiverLicenseNumber'?: string;
    'ExternalReceiptNumber'?: string;
    'IdentificationMethod'?: string;
    'PatientLicenseNumber'?: string;
    'PatientRegistrationLocationId'?: number;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesCreateReceiptV2RequestItem_Transactions[];
}
export interface SalesCreateReceiptV2RequestItem_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesCreateTransactionByDateV1RequestItem {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesUpdateDeliveryV1RequestItem {
    'ConsumerId'?: number;
    'DriverEmployeeId'?: string;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Id'?: number;
    'PatientLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'PlannedRoute'?: string;
    'RecipientAddressCity'?: string;
    'RecipientAddressCounty'?: string;
    'RecipientAddressPostalCode'?: string;
    'RecipientAddressState'?: string;
    'RecipientAddressStreet1'?: string;
    'RecipientAddressStreet2'?: string;
    'RecipientName'?: string;
    'RecipientZoneId'?: string;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesUpdateDeliveryV1RequestItem_Transactions[];
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesUpdateDeliveryV1RequestItem_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesUpdateDeliveryV2RequestItem {
    'ConsumerId'?: number;
    'DriverEmployeeId'?: string;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Id'?: number;
    'PatientLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'PlannedRoute'?: string;
    'RecipientAddressCity'?: string;
    'RecipientAddressCounty'?: string;
    'RecipientAddressPostalCode'?: string;
    'RecipientAddressState'?: string;
    'RecipientAddressStreet1'?: string;
    'RecipientAddressStreet2'?: string;
    'RecipientName'?: string;
    'RecipientZoneId'?: string;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesUpdateDeliveryV2RequestItem_Transactions[];
    'TransporterFacilityLicenseNumber'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesUpdateDeliveryV2RequestItem_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesUpdateDeliveryCompleteV1RequestItem {
    'AcceptedPackages'?: string[];
    'ActualArrivalDateTime'?: string;
    'Id'?: number;
    'PaymentType'?: string;
    'ReturnedPackages'?: SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages[];
}
export interface SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages {
    'Label'?: string;
    'ReturnQuantityVerified'?: number;
    'ReturnReason'?: string;
    'ReturnReasonNote'?: string;
    'ReturnUnitOfMeasure'?: string;
}
export interface SalesUpdateDeliveryCompleteV2RequestItem {
    'AcceptedPackages'?: string[];
    'ActualArrivalDateTime'?: string;
    'Id'?: number;
    'PaymentType'?: string;
    'ReturnedPackages'?: SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages[];
}
export interface SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages {
    'Label'?: string;
    'ReturnQuantityVerified'?: number;
    'ReturnReason'?: string;
    'ReturnReasonNote'?: string;
    'ReturnUnitOfMeasure'?: string;
}
export interface SalesUpdateDeliveryHubV1RequestItem {
    'DriverEmployeeId'?: string;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Id'?: number;
    'PhoneNumberForQuestions'?: string;
    'PlannedRoute'?: string;
    'TransporterFacilityId'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesUpdateDeliveryHubV2RequestItem {
    'DriverEmployeeId'?: string;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Id'?: number;
    'PhoneNumberForQuestions'?: string;
    'PlannedRoute'?: string;
    'TransporterFacilityId'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesUpdateDeliveryHubAcceptV1RequestItem {
    'Id'?: number;
}
export interface SalesUpdateDeliveryHubAcceptV2RequestItem {
    'Id'?: number;
}
export interface SalesUpdateDeliveryHubDepartV1RequestItem {
    'Id'?: number;
}
export interface SalesUpdateDeliveryHubDepartV2RequestItem {
    'Id'?: number;
}
export interface SalesUpdateDeliveryHubVerifyIdV1RequestItem {
    'Id'?: number;
    'PaymentType'?: string;
}
export interface SalesUpdateDeliveryHubVerifyIdV2RequestItem {
    'Id'?: number;
    'PaymentType'?: string;
}
export interface SalesUpdateDeliveryRetailerV1RequestItem {
    'DateTime'?: string;
    'Destinations'?: SalesUpdateDeliveryRetailerV1RequestItem_Destinations[];
    'DriverEmployeeId'?: string;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Id'?: number;
    'Packages'?: SalesUpdateDeliveryRetailerV1RequestItem_Packages[];
    'PhoneNumberForQuestions'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesUpdateDeliveryRetailerV1RequestItem_Destinations {
    'ConsumerId'?: string;
    'DriverEmployeeId'?: number;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Id'?: number;
    'PatientLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'PlannedRoute'?: string;
    'RecipientAddressCity'?: string;
    'RecipientAddressCounty'?: string;
    'RecipientAddressPostalCode'?: string;
    'RecipientAddressState'?: string;
    'RecipientAddressStreet1'?: string;
    'RecipientAddressStreet2'?: string;
    'RecipientName'?: string;
    'RecipientZoneId'?: string;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions[];
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesUpdateDeliveryRetailerV1RequestItem_Packages {
    'DateTime'?: string;
    'PackageLabel'?: string;
    'Quantity'?: number;
    'TotalPrice'?: number;
    'UnitOfMeasure'?: string;
}
export interface SalesUpdateDeliveryRetailerV2RequestItem {
    'DateTime'?: string;
    'Destinations'?: SalesUpdateDeliveryRetailerV2RequestItem_Destinations[];
    'DriverEmployeeId'?: string;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Id'?: number;
    'Packages'?: SalesUpdateDeliveryRetailerV2RequestItem_Packages[];
    'PhoneNumberForQuestions'?: string;
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesUpdateDeliveryRetailerV2RequestItem_Destinations {
    'ConsumerId'?: string;
    'DriverEmployeeId'?: number;
    'DriverName'?: string;
    'DriversLicenseNumber'?: string;
    'EstimatedArrivalDateTime'?: string;
    'EstimatedDepartureDateTime'?: string;
    'Id'?: number;
    'PatientLicenseNumber'?: string;
    'PhoneNumberForQuestions'?: string;
    'PlannedRoute'?: string;
    'RecipientAddressCity'?: string;
    'RecipientAddressCounty'?: string;
    'RecipientAddressPostalCode'?: string;
    'RecipientAddressState'?: string;
    'RecipientAddressStreet1'?: string;
    'RecipientAddressStreet2'?: string;
    'RecipientName'?: string;
    'RecipientZoneId'?: string;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions[];
    'VehicleLicensePlateNumber'?: string;
    'VehicleMake'?: string;
    'VehicleModel'?: string;
}
export interface SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesUpdateDeliveryRetailerV2RequestItem_Packages {
    'DateTime'?: string;
    'PackageLabel'?: string;
    'Quantity'?: number;
    'TotalPrice'?: number;
    'UnitOfMeasure'?: string;
}
export interface SalesUpdateReceiptV1RequestItem {
    'CaregiverLicenseNumber'?: string;
    'ExternalReceiptNumber'?: string;
    'Id'?: number;
    'IdentificationMethod'?: string;
    'PatientLicenseNumber'?: string;
    'PatientRegistrationLocationId'?: number;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesUpdateReceiptV1RequestItem_Transactions[];
}
export interface SalesUpdateReceiptV1RequestItem_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesUpdateReceiptV2RequestItem {
    'CaregiverLicenseNumber'?: string;
    'ExternalReceiptNumber'?: string;
    'Id'?: number;
    'IdentificationMethod'?: string;
    'PatientLicenseNumber'?: string;
    'PatientRegistrationLocationId'?: number;
    'SalesCustomerType'?: string;
    'SalesDateTime'?: string;
    'Transactions'?: SalesUpdateReceiptV2RequestItem_Transactions[];
}
export interface SalesUpdateReceiptV2RequestItem_Transactions {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
export interface SalesUpdateReceiptFinalizeV2RequestItem {
    'Id'?: number;
}
export interface SalesUpdateReceiptUnfinalizeV2RequestItem {
    'Id'?: number;
}
export interface SalesUpdateTransactionByDateV1RequestItem {
    'CityTax'?: string;
    'CountyTax'?: string;
    'DiscountAmount'?: string;
    'ExciseTax'?: string;
    'InvoiceNumber'?: string;
    'MunicipalTax'?: string;
    'PackageLabel'?: string;
    'Price'?: string;
    'QrCodes'?: string;
    'Quantity'?: number;
    'SalesTax'?: string;
    'SubTotal'?: string;
    'TotalAmount'?: number;
    'UnitOfMeasure'?: string;
    'UnitThcContent'?: number;
    'UnitThcContentUnitOfMeasure'?: string;
    'UnitThcPercent'?: number;
    'UnitWeight'?: number;
    'UnitWeightUnitOfMeasure'?: string;
}
