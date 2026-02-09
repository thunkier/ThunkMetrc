package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransferHub(
    @JsonProperty("ActualArrivalDateTime")
    val actualArrivalDateTime: String? = null,
    @JsonProperty("ActualDepartureDateTime")
    val actualDepartureDateTime: String? = null,
    @JsonProperty("ActualReturnArrivalDateTime")
    val actualReturnArrivalDateTime: String? = null,
    @JsonProperty("ActualReturnDepartureDateTime")
    val actualReturnDepartureDateTime: String? = null,
    @JsonProperty("CreatedByUserName")
    val createdByUserName: String? = null,
    @JsonProperty("CreatedDateTime")
    val createdDateTime: String? = null,
    @JsonProperty("DeliveryCount")
    val deliveryCount: Int? = null,
    @JsonProperty("DeliveryId")
    val deliveryId: Int? = null,
    @JsonProperty("DeliveryPackageCount")
    val deliveryPackageCount: Int? = null,
    @JsonProperty("DeliveryReceivedPackageCount")
    val deliveryReceivedPackageCount: Int? = null,
    @JsonProperty("DriverName")
    val driverName: String? = null,
    @JsonProperty("DriverOccupationalLicenseNumber")
    val driverOccupationalLicenseNumber: String? = null,
    @JsonProperty("DriverVehicleLicenseNumber")
    val driverVehicleLicenseNumber: String? = null,
    @JsonProperty("EstimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("EstimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("EstimatedReturnArrivalDateTime")
    val estimatedReturnArrivalDateTime: String? = null,
    @JsonProperty("EstimatedReturnDepartureDateTime")
    val estimatedReturnDepartureDateTime: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IsLayover")
    val isLayover: Boolean? = null,
    @JsonProperty("IsVoided")
    val isVoided: Boolean? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("ManifestNumber")
    val manifestNumber: String? = null,
    @JsonProperty("PackageCount")
    val packageCount: Int? = null,
    @JsonProperty("ReceivedDateTime")
    val receivedDateTime: String? = null,
    @JsonProperty("ReceivedDeliveryCount")
    val receivedDeliveryCount: Int? = null,
    @JsonProperty("ReceivedPackageCount")
    val receivedPackageCount: Int? = null,
    @JsonProperty("RecipientFacilityLicenseNumber")
    val recipientFacilityLicenseNumber: String? = null,
    @JsonProperty("RecipientFacilityName")
    val recipientFacilityName: String? = null,
    @JsonProperty("RejectedPackagesReturned")
    val rejectedPackagesReturned: Boolean? = null,
    @JsonProperty("ShipmentTransactionType")
    val shipmentTransactionType: Int? = null,
    @JsonProperty("ShipmentTypeName")
    val shipmentTypeName: String? = null,
    @JsonProperty("ShipperFacilityLicenseNumber")
    val shipperFacilityLicenseNumber: String? = null,
    @JsonProperty("ShipperFacilityName")
    val shipperFacilityName: String? = null,
    @JsonProperty("TransporterAcceptedDateTime")
    val transporterAcceptedDateTime: String? = null,
    @JsonProperty("TransporterActualArrivalDateTime")
    val transporterActualArrivalDateTime: String? = null,
    @JsonProperty("TransporterActualDepartureDateTime")
    val transporterActualDepartureDateTime: String? = null,
    @JsonProperty("TransporterEstimatedArrivalDateTime")
    val transporterEstimatedArrivalDateTime: String? = null,
    @JsonProperty("TransporterEstimatedDepartureDateTime")
    val transporterEstimatedDepartureDateTime: String? = null,
    @JsonProperty("TransporterFacilityLicenseNumber")
    val transporterFacilityLicenseNumber: String? = null,
    @JsonProperty("TransporterFacilityName")
    val transporterFacilityName: String? = null,
    @JsonProperty("VehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("VehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("VehicleModel")
    val vehicleModel: String? = null
)
