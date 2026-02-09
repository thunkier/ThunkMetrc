package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackagesPackage(
    @JsonProperty("ActualDepartureDateTime")
    val actualDepartureDateTime: String? = null,
    @JsonProperty("ExternalId")
    val externalId: Int? = null,
    @JsonProperty("GrossUnitOfWeightAbbreviation")
    val grossUnitOfWeightAbbreviation: String? = null,
    @JsonProperty("GrossWeight")
    val grossWeight: Double? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("ItemStrainName")
    val itemStrainName: String? = null,
    @JsonProperty("LabTestingStateName")
    val labTestingStateName: String? = null,
    @JsonProperty("ManifestNumber")
    val manifestNumber: String? = null,
    @JsonProperty("PackageId")
    val packageId: Int? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("ProcessingJobTypeName")
    val processingJobTypeName: String? = null,
    @JsonProperty("ProductCategoryName")
    val productCategoryName: String? = null,
    @JsonProperty("ProductName")
    val productName: String? = null,
    @JsonProperty("ReceivedDateTime")
    val receivedDateTime: String? = null,
    @JsonProperty("ReceivedQuantity")
    val receivedQuantity: Double? = null,
    @JsonProperty("ReceivedUnitOfMeasureAbbreviation")
    val receivedUnitOfMeasureAbbreviation: String? = null,
    @JsonProperty("ReceiverWholesalePrice")
    val receiverWholesalePrice: Double? = null,
    @JsonProperty("RecipientFacilityLicenseNumber")
    val recipientFacilityLicenseNumber: String? = null,
    @JsonProperty("RecipientFacilityName")
    val recipientFacilityName: String? = null,
    @JsonProperty("ShipmentPackageStateName")
    val shipmentPackageStateName: String? = null,
    @JsonProperty("ShippedQuantity")
    val shippedQuantity: Double? = null,
    @JsonProperty("ShippedUnitOfMeasureAbbreviation")
    val shippedUnitOfMeasureAbbreviation: String? = null,
    @JsonProperty("ShipperWholesalePrice")
    val shipperWholesalePrice: Double? = null,
    @JsonProperty("SourceHarvestNames")
    val sourceHarvestNames: String? = null,
    @JsonProperty("SourcePackageLabels")
    val sourcePackageLabels: String? = null
)
