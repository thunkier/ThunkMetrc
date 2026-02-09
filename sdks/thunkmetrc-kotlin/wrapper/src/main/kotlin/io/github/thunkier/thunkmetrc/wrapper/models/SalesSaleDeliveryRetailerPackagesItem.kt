package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesSaleDeliveryRetailerPackagesItem(
    @JsonProperty("ArchivedDate")
    val archivedDate: String? = null,
    @JsonProperty("CompletedDateTime")
    val completedDateTime: String? = null,
    @JsonProperty("IsOnInvestigationHold")
    val isOnInvestigationHold: Boolean? = null,
    @JsonProperty("IsOnInvestigationRecall")
    val isOnInvestigationRecall: Boolean? = null,
    @JsonProperty("IsOnRecall")
    val isOnRecall: Boolean? = null,
    @JsonProperty("IsOnRecallCombined")
    val isOnRecallCombined: Boolean? = null,
    @JsonProperty("ItemServingSize")
    val itemServingSize: String? = null,
    @JsonProperty("ItemStrainName")
    val itemStrainName: String? = null,
    @JsonProperty("ItemSupplyDurationDays")
    val itemSupplyDurationDays: Int? = null,
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
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("PackageId")
    val packageId: Int? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("ProductCategoryName")
    val productCategoryName: String? = null,
    @JsonProperty("ProductName")
    val productName: String? = null,
    @JsonProperty("Quantity")
    val quantity: Double? = null,
    @JsonProperty("RecordedByUserName")
    val recordedByUserName: String? = null,
    @JsonProperty("RecordedDateTime")
    val recordedDateTime: String? = null,
    @JsonProperty("RetailerDeliveryState")
    val retailerDeliveryState: String? = null,
    @JsonProperty("TotalPrice")
    val totalPrice: Double? = null,
    @JsonProperty("UnitOfMeasureAbbreviation")
    val unitOfMeasureAbbreviation: String? = null,
    @JsonProperty("UnitOfMeasureName")
    val unitOfMeasureName: String? = null
)
