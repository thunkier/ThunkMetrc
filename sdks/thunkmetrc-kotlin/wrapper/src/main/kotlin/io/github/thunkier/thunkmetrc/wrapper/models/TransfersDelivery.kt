package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersDelivery(
    @JsonProperty("ActualArrivalDateTime")
    val actualArrivalDateTime: String? = null,
    @JsonProperty("ActualDepartureDateTime")
    val actualDepartureDateTime: String? = null,
    @JsonProperty("ActualReturnArrivalDateTime")
    val actualReturnArrivalDateTime: String? = null,
    @JsonProperty("ActualReturnDepartureDateTime")
    val actualReturnDepartureDateTime: String? = null,
    @JsonProperty("DeliveryNumber")
    val deliveryNumber: Int? = null,
    @JsonProperty("DeliveryPackageCount")
    val deliveryPackageCount: Int? = null,
    @JsonProperty("DeliveryReceivedPackageCount")
    val deliveryReceivedPackageCount: Int? = null,
    @JsonProperty("EstimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("EstimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("EstimatedReturnArrivalDateTime")
    val estimatedReturnArrivalDateTime: String? = null,
    @JsonProperty("EstimatedReturnDepartureDateTime")
    val estimatedReturnDepartureDateTime: String? = null,
    @JsonProperty("GrossUnitOfWeightId")
    val grossUnitOfWeightId: Int? = null,
    @JsonProperty("GrossUnitOfWeightName")
    val grossUnitOfWeightName: String? = null,
    @JsonProperty("GrossWeight")
    val grossWeight: Double? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("InvoiceNumber")
    val invoiceNumber: String? = null,
    @JsonProperty("ManifestNumber")
    val manifestNumber: String? = null,
    @JsonProperty("PDFDocumentFileSystemId")
    val pDFDocumentFileSystemId: Int? = null,
    @JsonProperty("PlannedRoute")
    val plannedRoute: String? = null,
    @JsonProperty("ReceivedDateTime")
    val receivedDateTime: String? = null,
    @JsonProperty("RecipientFacilityLicenseNumber")
    val recipientFacilityLicenseNumber: String? = null,
    @JsonProperty("RecipientFacilityName")
    val recipientFacilityName: String? = null,
    @JsonProperty("RejectedPackagesReturned")
    val rejectedPackagesReturned: Boolean? = null,
    @JsonProperty("ShipmentTransactionType")
    val shipmentTransactionType: String? = null,
    @JsonProperty("ShipmentTypeName")
    val shipmentTypeName: String? = null
)
