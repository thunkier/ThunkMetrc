"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __exportStar = (this && this.__exportStar) || function(m, exports) {
    for (var p in m) if (p !== "default" && !Object.prototype.hasOwnProperty.call(exports, p)) __createBinding(exports, m, p);
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.MetrcWrapper = void 0;
__exportStar(require("@thunkmetrc/client"), exports);
const RateLimiter_1 = require("./RateLimiter");
class MetrcWrapper {
    constructor(client, config) {
        this.client = client;
        this.rateLimiter = new RateLimiter_1.MetrcRateLimiter(config);
    }
    async tagsGetPackageAvailableV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.tagsGetPackageAvailableV2(body, licenseNumber));
    }
    async tagsGetPlantAvailableV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.tagsGetPlantAvailableV2(body, licenseNumber));
    }
    async tagsGetStagedV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.tagsGetStagedV2(body, licenseNumber));
    }
    async transfersCreateExternalIncomingV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersCreateExternalIncomingV1(body, licenseNumber));
    }
    async transfersCreateExternalIncomingV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersCreateExternalIncomingV2(body, licenseNumber));
    }
    async transfersCreateTemplatesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersCreateTemplatesV1(body, licenseNumber));
    }
    async transfersCreateTemplatesOutgoingV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersCreateTemplatesOutgoingV2(body, licenseNumber));
    }
    async transfersDeleteExternalIncomingV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersDeleteExternalIncomingV1(id, body, licenseNumber));
    }
    async transfersDeleteExternalIncomingV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersDeleteExternalIncomingV2(id, body, licenseNumber));
    }
    async transfersDeleteTemplatesV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersDeleteTemplatesV1(id, body, licenseNumber));
    }
    async transfersDeleteTemplatesOutgoingV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersDeleteTemplatesOutgoingV2(id, body, licenseNumber));
    }
    async transfersGetDeliveriesPackagesStatesV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveriesPackagesStatesV1(body, No));
    }
    async transfersGetDeliveriesPackagesStatesV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveriesPackagesStatesV2(body, No));
    }
    async transfersGetDeliveryV1(id, body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryV1(id, body, No));
    }
    async transfersGetDeliveryV2(id, body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryV2(id, body, pageNumber, pageSize));
    }
    async transfersGetDeliveryPackageV1(id, body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageV1(id, body, No));
    }
    async transfersGetDeliveryPackageV2(id, body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageV2(id, body, pageNumber, pageSize));
    }
    async transfersGetDeliveryPackageRequiredlabtestbatchesV1(id, body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageRequiredlabtestbatchesV1(id, body, No));
    }
    async transfersGetDeliveryPackageRequiredlabtestbatchesV2(id, body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageRequiredlabtestbatchesV2(id, body, pageNumber, pageSize));
    }
    async transfersGetDeliveryPackageWholesaleV1(id, body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageWholesaleV1(id, body, No));
    }
    async transfersGetDeliveryPackageWholesaleV2(id, body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageWholesaleV2(id, body, pageNumber, pageSize));
    }
    async transfersGetDeliveryTransportersV1(id, body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryTransportersV1(id, body, No));
    }
    async transfersGetDeliveryTransportersV2(id, body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryTransportersV2(id, body, pageNumber, pageSize));
    }
    async transfersGetDeliveryTransportersDetailsV1(id, body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryTransportersDetailsV1(id, body, No));
    }
    async transfersGetDeliveryTransportersDetailsV2(id, body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryTransportersDetailsV2(id, body, pageNumber, pageSize));
    }
    async transfersGetHubV2(body, estimatedArrivalEnd, estimatedArrivalStart, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetHubV2(body, estimatedArrivalEnd, estimatedArrivalStart, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async transfersGetIncomingV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetIncomingV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async transfersGetIncomingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetIncomingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async transfersGetOutgoingV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetOutgoingV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async transfersGetOutgoingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetOutgoingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async transfersGetRejectedV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetRejectedV1(body, licenseNumber));
    }
    async transfersGetRejectedV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetRejectedV2(body, licenseNumber, pageNumber, pageSize));
    }
    async transfersGetTemplatesV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async transfersGetTemplatesDeliveryV1(id, body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesDeliveryV1(id, body, No));
    }
    async transfersGetTemplatesDeliveryPackageV1(id, body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesDeliveryPackageV1(id, body, No));
    }
    async transfersGetTemplatesDeliveryTransportersV1(id, body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesDeliveryTransportersV1(id, body, No));
    }
    async transfersGetTemplatesDeliveryTransportersDetailsV1(id, body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesDeliveryTransportersDetailsV1(id, body, No));
    }
    async transfersGetTemplatesOutgoingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesOutgoingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async transfersGetTemplatesOutgoingDeliveryV2(id, body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesOutgoingDeliveryV2(id, body, pageNumber, pageSize));
    }
    async transfersGetTemplatesOutgoingDeliveryPackageV2(id, body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesOutgoingDeliveryPackageV2(id, body, pageNumber, pageSize));
    }
    async transfersGetTemplatesOutgoingDeliveryTransportersV2(id, body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesOutgoingDeliveryTransportersV2(id, body, pageNumber, pageSize));
    }
    async transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id, body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id, body, No));
    }
    async transfersGetTypesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTypesV1(body, licenseNumber));
    }
    async transfersGetTypesV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transfersGetTypesV2(body, licenseNumber, pageNumber, pageSize));
    }
    async transfersUpdateExternalIncomingV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersUpdateExternalIncomingV1(body, licenseNumber));
    }
    async transfersUpdateExternalIncomingV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersUpdateExternalIncomingV2(body, licenseNumber));
    }
    async transfersUpdateTemplatesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersUpdateTemplatesV1(body, licenseNumber));
    }
    async transfersUpdateTemplatesOutgoingV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transfersUpdateTemplatesOutgoingV2(body, licenseNumber));
    }
    async wasteMethodsGetAllV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.wasteMethodsGetAllV2(body, No));
    }
    async additivesTemplatesCreateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.additivesTemplatesCreateV2(body, licenseNumber));
    }
    async additivesTemplatesGetV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.additivesTemplatesGetV2(id, body, licenseNumber));
    }
    async additivesTemplatesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.additivesTemplatesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async additivesTemplatesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.additivesTemplatesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async additivesTemplatesUpdateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.additivesTemplatesUpdateV2(body, licenseNumber));
    }
    async facilitiesGetAllV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.facilitiesGetAllV1(body, No));
    }
    async facilitiesGetAllV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.facilitiesGetAllV2(body, No));
    }
    async harvestsCreateFinishV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsCreateFinishV1(body, licenseNumber));
    }
    async harvestsCreatePackageV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsCreatePackageV1(body, licenseNumber));
    }
    async harvestsCreatePackageV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsCreatePackageV2(body, licenseNumber));
    }
    async harvestsCreatePackageTestingV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsCreatePackageTestingV1(body, licenseNumber));
    }
    async harvestsCreatePackageTestingV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsCreatePackageTestingV2(body, licenseNumber));
    }
    async harvestsCreateRemoveWasteV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsCreateRemoveWasteV1(body, licenseNumber));
    }
    async harvestsCreateUnfinishV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsCreateUnfinishV1(body, licenseNumber));
    }
    async harvestsCreateWasteV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsCreateWasteV2(body, licenseNumber));
    }
    async harvestsDeleteWasteV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsDeleteWasteV2(id, body, licenseNumber));
    }
    async harvestsGetV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.harvestsGetV1(id, body, licenseNumber));
    }
    async harvestsGetV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.harvestsGetV2(id, body, licenseNumber));
    }
    async harvestsGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.harvestsGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async harvestsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.harvestsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async harvestsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.harvestsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async harvestsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.harvestsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async harvestsGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.harvestsGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async harvestsGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.harvestsGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async harvestsGetWasteV2(body, harvestId, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.harvestsGetWasteV2(body, harvestId, licenseNumber, pageNumber, pageSize));
    }
    async harvestsGetWasteTypesV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.harvestsGetWasteTypesV1(body, No));
    }
    async harvestsGetWasteTypesV2(body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.harvestsGetWasteTypesV2(body, pageNumber, pageSize));
    }
    async harvestsUpdateFinishV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateFinishV2(body, licenseNumber));
    }
    async harvestsUpdateLocationV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateLocationV2(body, licenseNumber));
    }
    async harvestsUpdateMoveV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateMoveV1(body, licenseNumber));
    }
    async harvestsUpdateRenameV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateRenameV1(body, licenseNumber));
    }
    async harvestsUpdateRenameV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateRenameV2(body, licenseNumber));
    }
    async harvestsUpdateRestoreHarvestedPlantsV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateRestoreHarvestedPlantsV2(body, licenseNumber));
    }
    async harvestsUpdateUnfinishV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateUnfinishV2(body, licenseNumber));
    }
    async itemsCreateV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsCreateV1(body, licenseNumber));
    }
    async itemsCreateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsCreateV2(body, licenseNumber));
    }
    async itemsCreateBrandV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsCreateBrandV2(body, licenseNumber));
    }
    async itemsCreateFileV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsCreateFileV2(body, licenseNumber));
    }
    async itemsCreatePhotoV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsCreatePhotoV1(body, licenseNumber));
    }
    async itemsCreatePhotoV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsCreatePhotoV2(body, licenseNumber));
    }
    async itemsCreateUpdateV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsCreateUpdateV1(body, licenseNumber));
    }
    async itemsDeleteV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsDeleteV1(id, body, licenseNumber));
    }
    async itemsDeleteV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsDeleteV2(id, body, licenseNumber));
    }
    async itemsDeleteBrandV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsDeleteBrandV2(id, body, licenseNumber));
    }
    async itemsGetV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetV1(id, body, licenseNumber));
    }
    async itemsGetV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetV2(id, body, licenseNumber));
    }
    async itemsGetActiveV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetActiveV1(body, licenseNumber));
    }
    async itemsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async itemsGetBrandsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetBrandsV1(body, licenseNumber));
    }
    async itemsGetBrandsV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetBrandsV2(body, licenseNumber, pageNumber, pageSize));
    }
    async itemsGetCategoriesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetCategoriesV1(body, licenseNumber));
    }
    async itemsGetCategoriesV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetCategoriesV2(body, licenseNumber, pageNumber, pageSize));
    }
    async itemsGetFileV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetFileV2(id, body, licenseNumber));
    }
    async itemsGetInactiveV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetInactiveV1(body, licenseNumber));
    }
    async itemsGetInactiveV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetInactiveV2(body, licenseNumber, pageNumber, pageSize));
    }
    async itemsGetPhotoV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetPhotoV1(id, body, licenseNumber));
    }
    async itemsGetPhotoV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.itemsGetPhotoV2(id, body, licenseNumber));
    }
    async itemsUpdateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsUpdateV2(body, licenseNumber));
    }
    async itemsUpdateBrandV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.itemsUpdateBrandV2(body, licenseNumber));
    }
    async labTestsCreateRecordV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.labTestsCreateRecordV1(body, licenseNumber));
    }
    async labTestsCreateRecordV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.labTestsCreateRecordV2(body, licenseNumber));
    }
    async labTestsGetBatchesV2(body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.labTestsGetBatchesV2(body, pageNumber, pageSize));
    }
    async labTestsGetLabtestdocumentV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.labTestsGetLabtestdocumentV1(id, body, licenseNumber));
    }
    async labTestsGetLabtestdocumentV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.labTestsGetLabtestdocumentV2(id, body, licenseNumber));
    }
    async labTestsGetResultsV1(body, licenseNumber, packageId) {
        return this.rateLimiter.execute(null, true, () => this.client.labTestsGetResultsV1(body, licenseNumber, packageId));
    }
    async labTestsGetResultsV2(body, licenseNumber, packageId, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.labTestsGetResultsV2(body, licenseNumber, packageId, pageNumber, pageSize));
    }
    async labTestsGetStatesV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.labTestsGetStatesV1(body, No));
    }
    async labTestsGetStatesV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.labTestsGetStatesV2(body, No));
    }
    async labTestsGetTypesV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.labTestsGetTypesV1(body, No));
    }
    async labTestsGetTypesV2(body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.labTestsGetTypesV2(body, pageNumber, pageSize));
    }
    async labTestsUpdateLabtestdocumentV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.labTestsUpdateLabtestdocumentV1(body, licenseNumber));
    }
    async labTestsUpdateLabtestdocumentV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.labTestsUpdateLabtestdocumentV2(body, licenseNumber));
    }
    async labTestsUpdateResultReleaseV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.labTestsUpdateResultReleaseV1(body, licenseNumber));
    }
    async labTestsUpdateResultReleaseV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.labTestsUpdateResultReleaseV2(body, licenseNumber));
    }
    async strainsCreateV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.strainsCreateV1(body, licenseNumber));
    }
    async strainsCreateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.strainsCreateV2(body, licenseNumber));
    }
    async strainsCreateUpdateV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.strainsCreateUpdateV1(body, licenseNumber));
    }
    async strainsDeleteV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.strainsDeleteV1(id, body, licenseNumber));
    }
    async strainsDeleteV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.strainsDeleteV2(id, body, licenseNumber));
    }
    async strainsGetV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.strainsGetV1(id, body, licenseNumber));
    }
    async strainsGetV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.strainsGetV2(id, body, licenseNumber));
    }
    async strainsGetActiveV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.strainsGetActiveV1(body, licenseNumber));
    }
    async strainsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.strainsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async strainsGetInactiveV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.strainsGetInactiveV2(body, licenseNumber, pageNumber, pageSize));
    }
    async strainsUpdateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.strainsUpdateV2(body, licenseNumber));
    }
    async transportersCreateDriverV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transportersCreateDriverV2(body, licenseNumber));
    }
    async transportersCreateVehicleV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transportersCreateVehicleV2(body, licenseNumber));
    }
    async transportersDeleteDriverV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transportersDeleteDriverV2(id, body, licenseNumber));
    }
    async transportersDeleteVehicleV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transportersDeleteVehicleV2(id, body, licenseNumber));
    }
    async transportersGetDriverV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.transportersGetDriverV2(id, body, licenseNumber));
    }
    async transportersGetDriversV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transportersGetDriversV2(body, licenseNumber, pageNumber, pageSize));
    }
    async transportersGetVehicleV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.transportersGetVehicleV2(id, body, licenseNumber));
    }
    async transportersGetVehiclesV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.transportersGetVehiclesV2(body, licenseNumber, pageNumber, pageSize));
    }
    async transportersUpdateDriverV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transportersUpdateDriverV2(body, licenseNumber));
    }
    async transportersUpdateVehicleV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.transportersUpdateVehicleV2(body, licenseNumber));
    }
    async packagesCreateV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreateV1(body, licenseNumber));
    }
    async packagesCreateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreateV2(body, licenseNumber));
    }
    async packagesCreateAdjustV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreateAdjustV1(body, licenseNumber));
    }
    async packagesCreateAdjustV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreateAdjustV2(body, licenseNumber));
    }
    async packagesCreateChangeItemV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreateChangeItemV1(body, licenseNumber));
    }
    async packagesCreateChangeLocationV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreateChangeLocationV1(body, licenseNumber));
    }
    async packagesCreateFinishV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreateFinishV1(body, licenseNumber));
    }
    async packagesCreatePlantingsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreatePlantingsV1(body, licenseNumber));
    }
    async packagesCreatePlantingsV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreatePlantingsV2(body, licenseNumber));
    }
    async packagesCreateRemediateV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreateRemediateV1(body, licenseNumber));
    }
    async packagesCreateTestingV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreateTestingV1(body, licenseNumber));
    }
    async packagesCreateTestingV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreateTestingV2(body, licenseNumber));
    }
    async packagesCreateUnfinishV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesCreateUnfinishV1(body, licenseNumber));
    }
    async packagesDeleteV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesDeleteV2(id, body, licenseNumber));
    }
    async packagesGetV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetV1(id, body, licenseNumber));
    }
    async packagesGetV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetV2(id, body, licenseNumber));
    }
    async packagesGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async packagesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async packagesGetAdjustReasonsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetAdjustReasonsV1(body, licenseNumber));
    }
    async packagesGetAdjustReasonsV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetAdjustReasonsV2(body, licenseNumber, pageNumber, pageSize));
    }
    async packagesGetByLabelV1(label, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetByLabelV1(label, body, licenseNumber));
    }
    async packagesGetByLabelV2(label, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetByLabelV2(label, body, licenseNumber));
    }
    async packagesGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async packagesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async packagesGetIntransitV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetIntransitV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async packagesGetLabsamplesV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetLabsamplesV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async packagesGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async packagesGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async packagesGetSourceHarvestV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetSourceHarvestV2(id, body, licenseNumber));
    }
    async packagesGetTransferredV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetTransferredV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async packagesGetTypesV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetTypesV1(body, No));
    }
    async packagesGetTypesV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.packagesGetTypesV2(body, No));
    }
    async packagesUpdateAdjustV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateAdjustV2(body, licenseNumber));
    }
    async packagesUpdateChangeNoteV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateChangeNoteV1(body, licenseNumber));
    }
    async packagesUpdateDecontaminateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateDecontaminateV2(body, licenseNumber));
    }
    async packagesUpdateDonationFlagV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateDonationFlagV2(body, licenseNumber));
    }
    async packagesUpdateDonationUnflagV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateDonationUnflagV2(body, licenseNumber));
    }
    async packagesUpdateExternalidV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateExternalidV2(body, licenseNumber));
    }
    async packagesUpdateFinishV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateFinishV2(body, licenseNumber));
    }
    async packagesUpdateItemV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateItemV2(body, licenseNumber));
    }
    async packagesUpdateLabTestRequiredV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateLabTestRequiredV2(body, licenseNumber));
    }
    async packagesUpdateLocationV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateLocationV2(body, licenseNumber));
    }
    async packagesUpdateNoteV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateNoteV2(body, licenseNumber));
    }
    async packagesUpdateRemediateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateRemediateV2(body, licenseNumber));
    }
    async packagesUpdateTradesampleFlagV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateTradesampleFlagV2(body, licenseNumber));
    }
    async packagesUpdateTradesampleUnflagV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateTradesampleUnflagV2(body, licenseNumber));
    }
    async packagesUpdateUnfinishV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateUnfinishV2(body, licenseNumber));
    }
    async packagesUpdateUsebydateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateUsebydateV2(body, licenseNumber));
    }
    async sublocationsCreateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.sublocationsCreateV2(body, licenseNumber));
    }
    async sublocationsDeleteV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.sublocationsDeleteV2(id, body, licenseNumber));
    }
    async sublocationsGetV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.sublocationsGetV2(id, body, licenseNumber));
    }
    async sublocationsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.sublocationsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async sublocationsGetInactiveV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.sublocationsGetInactiveV2(body, licenseNumber, pageNumber, pageSize));
    }
    async sublocationsUpdateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.sublocationsUpdateV2(body, licenseNumber));
    }
    async unitsOfMeasureGetActiveV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.unitsOfMeasureGetActiveV1(body, No));
    }
    async unitsOfMeasureGetActiveV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.unitsOfMeasureGetActiveV2(body, No));
    }
    async unitsOfMeasureGetInactiveV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.unitsOfMeasureGetInactiveV2(body, No));
    }
    async plantBatchesCreateAdditivesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateAdditivesV1(body, licenseNumber));
    }
    async plantBatchesCreateAdditivesV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateAdditivesV2(body, licenseNumber));
    }
    async plantBatchesCreateAdditivesUsingtemplateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateAdditivesUsingtemplateV2(body, licenseNumber));
    }
    async plantBatchesCreateAdjustV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateAdjustV1(body, licenseNumber));
    }
    async plantBatchesCreateAdjustV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateAdjustV2(body, licenseNumber));
    }
    async plantBatchesCreateChangegrowthphaseV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateChangegrowthphaseV1(body, licenseNumber));
    }
    async plantBatchesCreateGrowthphaseV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateGrowthphaseV2(body, licenseNumber));
    }
    async plantBatchesCreatePackageV2(body, isFromMotherPlant, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreatePackageV2(body, isFromMotherPlant, licenseNumber));
    }
    async plantBatchesCreatePackageFrommotherplantV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreatePackageFrommotherplantV1(body, licenseNumber));
    }
    async plantBatchesCreatePackageFrommotherplantV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreatePackageFrommotherplantV2(body, licenseNumber));
    }
    async plantBatchesCreatePlantingsV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreatePlantingsV2(body, licenseNumber));
    }
    async plantBatchesCreateSplitV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateSplitV1(body, licenseNumber));
    }
    async plantBatchesCreateSplitV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateSplitV2(body, licenseNumber));
    }
    async plantBatchesCreateWasteV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateWasteV1(body, licenseNumber));
    }
    async plantBatchesCreateWasteV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateWasteV2(body, licenseNumber));
    }
    async plantBatchesCreatepackagesV1(body, isFromMotherPlant, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreatepackagesV1(body, isFromMotherPlant, licenseNumber));
    }
    async plantBatchesCreateplantingsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateplantingsV1(body, licenseNumber));
    }
    async plantBatchesDeleteV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesDeleteV1(body, licenseNumber));
    }
    async plantBatchesDeleteV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesDeleteV2(body, licenseNumber));
    }
    async plantBatchesGetV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetV1(id, body, licenseNumber));
    }
    async plantBatchesGetV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetV2(id, body, licenseNumber));
    }
    async plantBatchesGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async plantBatchesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async plantBatchesGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async plantBatchesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async plantBatchesGetTypesV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetTypesV1(body, No));
    }
    async plantBatchesGetTypesV2(body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetTypesV2(body, pageNumber, pageSize));
    }
    async plantBatchesGetWasteV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetWasteV2(body, licenseNumber, pageNumber, pageSize));
    }
    async plantBatchesGetWasteReasonsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetWasteReasonsV1(body, licenseNumber));
    }
    async plantBatchesGetWasteReasonsV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetWasteReasonsV2(body, licenseNumber));
    }
    async plantBatchesUpdateLocationV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesUpdateLocationV2(body, licenseNumber));
    }
    async plantBatchesUpdateMoveplantbatchesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesUpdateMoveplantbatchesV1(body, licenseNumber));
    }
    async plantBatchesUpdateNameV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesUpdateNameV2(body, licenseNumber));
    }
    async plantBatchesUpdateStrainV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesUpdateStrainV2(body, licenseNumber));
    }
    async plantBatchesUpdateTagV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantBatchesUpdateTagV2(body, licenseNumber));
    }
    async plantsCreateAdditivesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesV1(body, licenseNumber));
    }
    async plantsCreateAdditivesV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesV2(body, licenseNumber));
    }
    async plantsCreateAdditivesBylocationV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesBylocationV1(body, licenseNumber));
    }
    async plantsCreateAdditivesBylocationV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesBylocationV2(body, licenseNumber));
    }
    async plantsCreateAdditivesBylocationUsingtemplateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesBylocationUsingtemplateV2(body, licenseNumber));
    }
    async plantsCreateAdditivesUsingtemplateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesUsingtemplateV2(body, licenseNumber));
    }
    async plantsCreateChangegrowthphasesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateChangegrowthphasesV1(body, licenseNumber));
    }
    async plantsCreateHarvestplantsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateHarvestplantsV1(body, licenseNumber));
    }
    async plantsCreateManicureV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateManicureV2(body, licenseNumber));
    }
    async plantsCreateManicureplantsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateManicureplantsV1(body, licenseNumber));
    }
    async plantsCreateMoveplantsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateMoveplantsV1(body, licenseNumber));
    }
    async plantsCreatePlantbatchPackageV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreatePlantbatchPackageV1(body, licenseNumber));
    }
    async plantsCreatePlantbatchPackageV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreatePlantbatchPackageV2(body, licenseNumber));
    }
    async plantsCreatePlantingsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreatePlantingsV1(body, licenseNumber));
    }
    async plantsCreatePlantingsV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreatePlantingsV2(body, licenseNumber));
    }
    async plantsCreateWasteV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateWasteV1(body, licenseNumber));
    }
    async plantsCreateWasteV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsCreateWasteV2(body, licenseNumber));
    }
    async plantsDeleteV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsDeleteV1(body, licenseNumber));
    }
    async plantsDeleteV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsDeleteV2(body, licenseNumber));
    }
    async plantsGetV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetV1(id, body, licenseNumber));
    }
    async plantsGetV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetV2(id, body, licenseNumber));
    }
    async plantsGetAdditivesV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetAdditivesV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async plantsGetAdditivesV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetAdditivesV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async plantsGetAdditivesTypesV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetAdditivesTypesV1(body, No));
    }
    async plantsGetAdditivesTypesV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetAdditivesTypesV2(body, No));
    }
    async plantsGetByLabelV1(label, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetByLabelV1(label, body, licenseNumber));
    }
    async plantsGetByLabelV2(label, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetByLabelV2(label, body, licenseNumber));
    }
    async plantsGetFloweringV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetFloweringV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async plantsGetFloweringV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetFloweringV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async plantsGetGrowthPhasesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetGrowthPhasesV1(body, licenseNumber));
    }
    async plantsGetGrowthPhasesV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetGrowthPhasesV2(body, licenseNumber));
    }
    async plantsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async plantsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async plantsGetMotherV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetMotherV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async plantsGetMotherInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetMotherInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async plantsGetMotherOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetMotherOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async plantsGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async plantsGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async plantsGetVegetativeV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetVegetativeV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async plantsGetVegetativeV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetVegetativeV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async plantsGetWasteV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetWasteV2(body, licenseNumber, pageNumber, pageSize));
    }
    async plantsGetWasteMethodsAllV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetWasteMethodsAllV1(body, No));
    }
    async plantsGetWasteMethodsAllV2(body, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetWasteMethodsAllV2(body, pageNumber, pageSize));
    }
    async plantsGetWastePackageV2(id, body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetWastePackageV2(id, body, licenseNumber, pageNumber, pageSize));
    }
    async plantsGetWastePlantV2(id, body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetWastePlantV2(id, body, licenseNumber, pageNumber, pageSize));
    }
    async plantsGetWasteReasonsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetWasteReasonsV1(body, licenseNumber));
    }
    async plantsGetWasteReasonsV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.plantsGetWasteReasonsV2(body, licenseNumber, pageNumber, pageSize));
    }
    async plantsUpdateAdjustV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateAdjustV2(body, licenseNumber));
    }
    async plantsUpdateGrowthphaseV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateGrowthphaseV2(body, licenseNumber));
    }
    async plantsUpdateHarvestV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateHarvestV2(body, licenseNumber));
    }
    async plantsUpdateLocationV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateLocationV2(body, licenseNumber));
    }
    async plantsUpdateMergeV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateMergeV2(body, licenseNumber));
    }
    async plantsUpdateSplitV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateSplitV2(body, licenseNumber));
    }
    async plantsUpdateStrainV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateStrainV2(body, licenseNumber));
    }
    async plantsUpdateTagV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateTagV2(body, licenseNumber));
    }
    async employeesGetAllV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.employeesGetAllV1(body, licenseNumber));
    }
    async employeesGetAllV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.employeesGetAllV2(body, licenseNumber, pageNumber, pageSize));
    }
    async employeesGetPermissionsV2(body, employeeLicenseNumber, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.employeesGetPermissionsV2(body, employeeLicenseNumber, licenseNumber));
    }
    async patientCheckInsCreateV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsCreateV1(body, licenseNumber));
    }
    async patientCheckInsCreateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsCreateV2(body, licenseNumber));
    }
    async patientCheckInsDeleteV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsDeleteV1(id, body, licenseNumber));
    }
    async patientCheckInsDeleteV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsDeleteV2(id, body, licenseNumber));
    }
    async patientCheckInsGetAllV1(body, checkinDateEnd, checkinDateStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.patientCheckInsGetAllV1(body, checkinDateEnd, checkinDateStart, licenseNumber));
    }
    async patientCheckInsGetAllV2(body, checkinDateEnd, checkinDateStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.patientCheckInsGetAllV2(body, checkinDateEnd, checkinDateStart, licenseNumber));
    }
    async patientCheckInsGetLocationsV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.patientCheckInsGetLocationsV1(body, No));
    }
    async patientCheckInsGetLocationsV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.patientCheckInsGetLocationsV2(body, No));
    }
    async patientCheckInsUpdateV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsUpdateV1(body, licenseNumber));
    }
    async patientCheckInsUpdateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsUpdateV2(body, licenseNumber));
    }
    async patientsStatusGetStatusesByPatientLicenseNumberV1(patientLicenseNumber, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.patientsStatusGetStatusesByPatientLicenseNumberV1(patientLicenseNumber, body, licenseNumber));
    }
    async patientsStatusGetStatusesByPatientLicenseNumberV2(patientLicenseNumber, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.patientsStatusGetStatusesByPatientLicenseNumberV2(patientLicenseNumber, body, licenseNumber));
    }
    async processingJobsCreateAdjustV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateAdjustV1(body, licenseNumber));
    }
    async processingJobsCreateAdjustV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateAdjustV2(body, licenseNumber));
    }
    async processingJobsCreateJobtypesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateJobtypesV1(body, licenseNumber));
    }
    async processingJobsCreateJobtypesV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateJobtypesV2(body, licenseNumber));
    }
    async processingJobsCreateStartV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateStartV1(body, licenseNumber));
    }
    async processingJobsCreateStartV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateStartV2(body, licenseNumber));
    }
    async processingJobsCreatepackagesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreatepackagesV1(body, licenseNumber));
    }
    async processingJobsCreatepackagesV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreatepackagesV2(body, licenseNumber));
    }
    async processingJobsDeleteV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsDeleteV1(id, body, licenseNumber));
    }
    async processingJobsDeleteV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsDeleteV2(id, body, licenseNumber));
    }
    async processingJobsDeleteJobtypesV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsDeleteJobtypesV1(id, body, licenseNumber));
    }
    async processingJobsDeleteJobtypesV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsDeleteJobtypesV2(id, body, licenseNumber));
    }
    async processingJobsGetV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetV1(id, body, licenseNumber));
    }
    async processingJobsGetV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetV2(id, body, licenseNumber));
    }
    async processingJobsGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async processingJobsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async processingJobsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async processingJobsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async processingJobsGetJobtypesActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async processingJobsGetJobtypesActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async processingJobsGetJobtypesAttributesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesAttributesV1(body, licenseNumber));
    }
    async processingJobsGetJobtypesAttributesV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesAttributesV2(body, licenseNumber));
    }
    async processingJobsGetJobtypesCategoriesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesCategoriesV1(body, licenseNumber));
    }
    async processingJobsGetJobtypesCategoriesV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesCategoriesV2(body, licenseNumber));
    }
    async processingJobsGetJobtypesInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async processingJobsGetJobtypesInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async processingJobsUpdateFinishV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateFinishV1(body, licenseNumber));
    }
    async processingJobsUpdateFinishV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateFinishV2(body, licenseNumber));
    }
    async processingJobsUpdateJobtypesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateJobtypesV1(body, licenseNumber));
    }
    async processingJobsUpdateJobtypesV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateJobtypesV2(body, licenseNumber));
    }
    async processingJobsUpdateUnfinishV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateUnfinishV1(body, licenseNumber));
    }
    async processingJobsUpdateUnfinishV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateUnfinishV2(body, licenseNumber));
    }
    async retailIdCreateAssociateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.retailIdCreateAssociateV2(body, licenseNumber));
    }
    async retailIdCreateGenerateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.retailIdCreateGenerateV2(body, licenseNumber));
    }
    async retailIdCreateMergeV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.retailIdCreateMergeV2(body, licenseNumber));
    }
    async retailIdCreatePackageInfoV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.retailIdCreatePackageInfoV2(body, licenseNumber));
    }
    async retailIdGetReceiveByLabelV2(label, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.retailIdGetReceiveByLabelV2(label, body, licenseNumber));
    }
    async retailIdGetReceiveQrByShortCodeV2(shortCode, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.retailIdGetReceiveQrByShortCodeV2(shortCode, body, licenseNumber));
    }
    async sandboxCreateIntegratorSetupV2(body, userKey) {
        return this.rateLimiter.execute(null, false, () => this.client.sandboxCreateIntegratorSetupV2(body, userKey));
    }
    async caregiversStatusGetByCaregiverLicenseNumberV1(caregiverLicenseNumber, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.caregiversStatusGetByCaregiverLicenseNumberV1(caregiverLicenseNumber, body, licenseNumber));
    }
    async caregiversStatusGetByCaregiverLicenseNumberV2(caregiverLicenseNumber, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.caregiversStatusGetByCaregiverLicenseNumberV2(caregiverLicenseNumber, body, licenseNumber));
    }
    async locationsCreateV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.locationsCreateV1(body, licenseNumber));
    }
    async locationsCreateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.locationsCreateV2(body, licenseNumber));
    }
    async locationsCreateUpdateV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.locationsCreateUpdateV1(body, licenseNumber));
    }
    async locationsDeleteV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.locationsDeleteV1(id, body, licenseNumber));
    }
    async locationsDeleteV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.locationsDeleteV2(id, body, licenseNumber));
    }
    async locationsGetV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.locationsGetV1(id, body, licenseNumber));
    }
    async locationsGetV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.locationsGetV2(id, body, licenseNumber));
    }
    async locationsGetActiveV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.locationsGetActiveV1(body, licenseNumber));
    }
    async locationsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.locationsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async locationsGetInactiveV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.locationsGetInactiveV2(body, licenseNumber, pageNumber, pageSize));
    }
    async locationsGetTypesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.locationsGetTypesV1(body, licenseNumber));
    }
    async locationsGetTypesV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.locationsGetTypesV2(body, licenseNumber, pageNumber, pageSize));
    }
    async locationsUpdateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.locationsUpdateV2(body, licenseNumber));
    }
    async patientsCreateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientsCreateV2(body, licenseNumber));
    }
    async patientsCreateAddV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientsCreateAddV1(body, licenseNumber));
    }
    async patientsCreateUpdateV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientsCreateUpdateV1(body, licenseNumber));
    }
    async patientsDeleteV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientsDeleteV1(id, body, licenseNumber));
    }
    async patientsDeleteV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientsDeleteV2(id, body, licenseNumber));
    }
    async patientsGetV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.patientsGetV1(id, body, licenseNumber));
    }
    async patientsGetV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.patientsGetV2(id, body, licenseNumber));
    }
    async patientsGetActiveV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.patientsGetActiveV1(body, licenseNumber));
    }
    async patientsGetActiveV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.patientsGetActiveV2(body, licenseNumber, pageNumber, pageSize));
    }
    async patientsUpdateV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.patientsUpdateV2(body, licenseNumber));
    }
    async salesCreateDeliveryV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryV1(body, licenseNumber));
    }
    async salesCreateDeliveryV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryV2(body, licenseNumber));
    }
    async salesCreateDeliveryRetailerV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerV1(body, licenseNumber));
    }
    async salesCreateDeliveryRetailerV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerV2(body, licenseNumber));
    }
    async salesCreateDeliveryRetailerDepartV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerDepartV1(body, licenseNumber));
    }
    async salesCreateDeliveryRetailerDepartV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerDepartV2(body, licenseNumber));
    }
    async salesCreateDeliveryRetailerEndV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerEndV1(body, licenseNumber));
    }
    async salesCreateDeliveryRetailerEndV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerEndV2(body, licenseNumber));
    }
    async salesCreateDeliveryRetailerRestockV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerRestockV1(body, licenseNumber));
    }
    async salesCreateDeliveryRetailerRestockV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerRestockV2(body, licenseNumber));
    }
    async salesCreateDeliveryRetailerSaleV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerSaleV1(body, licenseNumber));
    }
    async salesCreateDeliveryRetailerSaleV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerSaleV2(body, licenseNumber));
    }
    async salesCreateReceiptV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateReceiptV1(body, licenseNumber));
    }
    async salesCreateReceiptV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateReceiptV2(body, licenseNumber));
    }
    async salesCreateTransactionByDateV1(date, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesCreateTransactionByDateV1(date, body, licenseNumber));
    }
    async salesDeleteDeliveryV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesDeleteDeliveryV1(id, body, licenseNumber));
    }
    async salesDeleteDeliveryV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesDeleteDeliveryV2(id, body, licenseNumber));
    }
    async salesDeleteDeliveryRetailerV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesDeleteDeliveryRetailerV1(id, body, licenseNumber));
    }
    async salesDeleteDeliveryRetailerV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesDeleteDeliveryRetailerV2(id, body, licenseNumber));
    }
    async salesDeleteReceiptV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesDeleteReceiptV1(id, body, licenseNumber));
    }
    async salesDeleteReceiptV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesDeleteReceiptV2(id, body, licenseNumber));
    }
    async salesGetCountiesV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetCountiesV1(body, No));
    }
    async salesGetCountiesV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetCountiesV2(body, No));
    }
    async salesGetCustomertypesV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetCustomertypesV1(body, No));
    }
    async salesGetCustomertypesV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetCustomertypesV2(body, No));
    }
    async salesGetDeliveriesActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
    }
    async salesGetDeliveriesActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
    }
    async salesGetDeliveriesInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
    }
    async salesGetDeliveriesInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
    }
    async salesGetDeliveriesRetailerActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesRetailerActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async salesGetDeliveriesRetailerActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesRetailerActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async salesGetDeliveriesRetailerInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesRetailerInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
    }
    async salesGetDeliveriesRetailerInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesRetailerInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
    }
    async salesGetDeliveriesReturnreasonsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesReturnreasonsV1(body, licenseNumber));
    }
    async salesGetDeliveriesReturnreasonsV2(body, licenseNumber, pageNumber, pageSize) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesReturnreasonsV2(body, licenseNumber, pageNumber, pageSize));
    }
    async salesGetDeliveryV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveryV1(id, body, licenseNumber));
    }
    async salesGetDeliveryV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveryV2(id, body, licenseNumber));
    }
    async salesGetDeliveryRetailerV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveryRetailerV1(id, body, licenseNumber));
    }
    async salesGetDeliveryRetailerV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveryRetailerV2(id, body, licenseNumber));
    }
    async salesGetPatientRegistrationsLocationsV1(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetPatientRegistrationsLocationsV1(body, No));
    }
    async salesGetPatientRegistrationsLocationsV2(body, No) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetPatientRegistrationsLocationsV2(body, No));
    }
    async salesGetPaymenttypesV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetPaymenttypesV1(body, licenseNumber));
    }
    async salesGetPaymenttypesV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetPaymenttypesV2(body, licenseNumber));
    }
    async salesGetReceiptV1(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptV1(id, body, licenseNumber));
    }
    async salesGetReceiptV2(id, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptV2(id, body, licenseNumber));
    }
    async salesGetReceiptsActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptsActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
    }
    async salesGetReceiptsActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptsActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
    }
    async salesGetReceiptsExternalByExternalNumberV2(externalNumber, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptsExternalByExternalNumberV2(externalNumber, body, licenseNumber));
    }
    async salesGetReceiptsInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptsInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
    }
    async salesGetReceiptsInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptsInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
    }
    async salesGetTransactionsV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetTransactionsV1(body, licenseNumber));
    }
    async salesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesDateStart, salesDateEnd, body, licenseNumber) {
        return this.rateLimiter.execute(null, true, () => this.client.salesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesDateStart, salesDateEnd, body, licenseNumber));
    }
    async salesUpdateDeliveryV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryV1(body, licenseNumber));
    }
    async salesUpdateDeliveryV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryV2(body, licenseNumber));
    }
    async salesUpdateDeliveryCompleteV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryCompleteV1(body, licenseNumber));
    }
    async salesUpdateDeliveryCompleteV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryCompleteV2(body, licenseNumber));
    }
    async salesUpdateDeliveryHubV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubV1(body, licenseNumber));
    }
    async salesUpdateDeliveryHubV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubV2(body, licenseNumber));
    }
    async salesUpdateDeliveryHubAcceptV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubAcceptV1(body, licenseNumber));
    }
    async salesUpdateDeliveryHubAcceptV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubAcceptV2(body, licenseNumber));
    }
    async salesUpdateDeliveryHubDepartV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubDepartV1(body, licenseNumber));
    }
    async salesUpdateDeliveryHubDepartV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubDepartV2(body, licenseNumber));
    }
    async salesUpdateDeliveryHubVerifyIdV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubVerifyIdV1(body, licenseNumber));
    }
    async salesUpdateDeliveryHubVerifyIdV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubVerifyIdV2(body, licenseNumber));
    }
    async salesUpdateDeliveryRetailerV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryRetailerV1(body, licenseNumber));
    }
    async salesUpdateDeliveryRetailerV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryRetailerV2(body, licenseNumber));
    }
    async salesUpdateReceiptV1(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateReceiptV1(body, licenseNumber));
    }
    async salesUpdateReceiptV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateReceiptV2(body, licenseNumber));
    }
    async salesUpdateReceiptFinalizeV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateReceiptFinalizeV2(body, licenseNumber));
    }
    async salesUpdateReceiptUnfinalizeV2(body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateReceiptUnfinalizeV2(body, licenseNumber));
    }
    async salesUpdateTransactionByDateV1(date, body, licenseNumber) {
        return this.rateLimiter.execute(null, false, () => this.client.salesUpdateTransactionByDateV1(date, body, licenseNumber));
    }
}
exports.MetrcWrapper = MetrcWrapper;
