package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransferDeliveryPackage(
    @JsonProperty("ContainsRemediatedProduct")
    val containsRemediatedProduct: Boolean? = null,
    @JsonProperty("ExpirationDate")
    val expirationDate: String? = null,
    @JsonProperty("ExternalId")
    val externalId: Int? = null,
    @JsonProperty("GrossUnitOfWeightName")
    val grossUnitOfWeightName: String? = null,
    @JsonProperty("IsDonation")
    val isDonation: Boolean? = null,
    @JsonProperty("IsTestingSample")
    val isTestingSample: Boolean? = null,
    @JsonProperty("IsTradeSample")
    val isTradeSample: Boolean? = null,
    @JsonProperty("IsTradeSamplePersistent")
    val isTradeSamplePersistent: Boolean? = null,
    @JsonProperty("ItemBrandName")
    val itemBrandName: String? = null,
    @JsonProperty("ItemCategoryName")
    val itemCategoryName: String? = null,
    @JsonProperty("ItemId")
    val itemId: Int? = null,
    @JsonProperty("ItemName")
    val itemName: String? = null,
    @JsonProperty("ItemServingSize")
    val itemServingSize: String? = null,
    @JsonProperty("ItemStrainName")
    val itemStrainName: String? = null,
    @JsonProperty("ItemSupplyDurationDays")
    val itemSupplyDurationDays: Int? = null,
    @JsonProperty("ItemUnitCbdAContent")
    val itemUnitCbdAContent: Double? = null,
    @JsonProperty("ItemUnitCbdAContentDose")
    val itemUnitCbdAContentDose: Double? = null,
    @JsonProperty("ItemUnitCbdAContentDoseUnitOfMeasureName")
    val itemUnitCbdAContentDoseUnitOfMeasureName: String? = null,
    @JsonProperty("ItemUnitCbdAContentUnitOfMeasureName")
    val itemUnitCbdAContentUnitOfMeasureName: String? = null,
    @JsonProperty("ItemUnitCbdAPercent")
    val itemUnitCbdAPercent: Double? = null,
    @JsonProperty("ItemUnitCbdContent")
    val itemUnitCbdContent: Double? = null,
    @JsonProperty("ItemUnitCbdContentDose")
    val itemUnitCbdContentDose: Double? = null,
    @JsonProperty("ItemUnitCbdContentDoseUnitOfMeasureName")
    val itemUnitCbdContentDoseUnitOfMeasureName: String? = null,
    @JsonProperty("ItemUnitCbdContentUnitOfMeasureName")
    val itemUnitCbdContentUnitOfMeasureName: String? = null,
    @JsonProperty("ItemUnitCbdPercent")
    val itemUnitCbdPercent: Double? = null,
    @JsonProperty("ItemUnitQuantity")
    val itemUnitQuantity: Double? = null,
    @JsonProperty("ItemUnitQuantityUnitOfMeasureName")
    val itemUnitQuantityUnitOfMeasureName: String? = null,
    @JsonProperty("ItemUnitThcAContent")
    val itemUnitThcAContent: Double? = null,
    @JsonProperty("ItemUnitThcAContentDose")
    val itemUnitThcAContentDose: Double? = null,
    @JsonProperty("ItemUnitThcAContentDoseUnitOfMeasureName")
    val itemUnitThcAContentDoseUnitOfMeasureName: String? = null,
    @JsonProperty("ItemUnitThcAContentUnitOfMeasureName")
    val itemUnitThcAContentUnitOfMeasureName: String? = null,
    @JsonProperty("ItemUnitThcAPercent")
    val itemUnitThcAPercent: Double? = null,
    @JsonProperty("ItemUnitThcContent")
    val itemUnitThcContent: Double? = null,
    @JsonProperty("ItemUnitThcContentDose")
    val itemUnitThcContentDose: Double? = null,
    @JsonProperty("ItemUnitThcContentDoseUnitOfMeasureName")
    val itemUnitThcContentDoseUnitOfMeasureName: String? = null,
    @JsonProperty("ItemUnitThcContentUnitOfMeasureName")
    val itemUnitThcContentUnitOfMeasureName: String? = null,
    @JsonProperty("ItemUnitThcPercent")
    val itemUnitThcPercent: Double? = null,
    @JsonProperty("ItemUnitVolume")
    val itemUnitVolume: Double? = null,
    @JsonProperty("ItemUnitVolumeUnitOfMeasureName")
    val itemUnitVolumeUnitOfMeasureName: String? = null,
    @JsonProperty("ItemUnitWeight")
    val itemUnitWeight: Double? = null,
    @JsonProperty("ItemUnitWeightUnitOfMeasureName")
    val itemUnitWeightUnitOfMeasureName: String? = null,
    @JsonProperty("LabTestingState")
    val labTestingState: String? = null,
    @JsonProperty("PackageId")
    val packageId: Int? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("PackageType")
    val packageType: String? = null,
    @JsonProperty("PackagedDate")
    val packagedDate: String? = null,
    @JsonProperty("ProductLabel")
    val productLabel: TransfersTransferDeliveryPackageProductLabel? = null,
    @JsonProperty("ProductRequiresRemediation")
    val productRequiresRemediation: Boolean? = null,
    @JsonProperty("ProductionBatchNumber")
    val productionBatchNumber: String? = null,
    @JsonProperty("ReceivedQuantity")
    val receivedQuantity: Double? = null,
    @JsonProperty("ReceivedUnitOfMeasureName")
    val receivedUnitOfMeasureName: String? = null,
    @JsonProperty("RemediationDate")
    val remediationDate: String? = null,
    @JsonProperty("SellByDate")
    val sellByDate: String? = null,
    @JsonProperty("ShipmentPackageState")
    val shipmentPackageState: String? = null,
    @JsonProperty("ShippedQuantity")
    val shippedQuantity: Double? = null,
    @JsonProperty("ShippedUnitOfMeasureName")
    val shippedUnitOfMeasureName: String? = null,
    @JsonProperty("SourceHarvestNames")
    val sourceHarvestNames: String? = null,
    @JsonProperty("SourcePackageIsDonation")
    val sourcePackageIsDonation: Boolean? = null,
    @JsonProperty("SourcePackageIsTradeSample")
    val sourcePackageIsTradeSample: Boolean? = null,
    @JsonProperty("SourcePackageLabels")
    val sourcePackageLabels: String? = null,
    @JsonProperty("SourceProductionBatchNumbers")
    val sourceProductionBatchNumbers: String? = null,
    @JsonProperty("UseByDate")
    val useByDate: String? = null
)
