package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ActiveSaleDeliveriesRetailer(
    @JsonProperty("AcceptedDateTime")
    val acceptedDateTime: String? = null,
    @JsonProperty("ActualDepartureDateTime")
    val actualDepartureDateTime: String? = null,
    @JsonProperty("AllowFullEdit")
    val allowFullEdit: Boolean? = null,
    @JsonProperty("CompletedDateTime")
    val completedDateTime: String? = null,
    @JsonProperty("DateTime")
    val dateTime: String? = null,
    @JsonProperty("Destinations")
    val destinations: List<Any>? = null,
    @JsonProperty("Direction")
    val direction: String? = null,
    @JsonProperty("DriverEmployeeId")
    val driverEmployeeId: String? = null,
    @JsonProperty("DriverName")
    val driverName: String? = null,
    @JsonProperty("DriversLicenseNumber")
    val driversLicenseNumber: String? = null,
    @JsonProperty("EstimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("FacilityLicenseNumber")
    val facilityLicenseNumber: String? = null,
    @JsonProperty("FacilityName")
    val facilityName: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("Leg")
    val leg: Int? = null,
    @JsonProperty("Packages")
    val packages: List<Any>? = null,
    @JsonProperty("RecordedByUserName")
    val recordedByUserName: String? = null,
    @JsonProperty("RecordedDateTime")
    val recordedDateTime: String? = null,
    @JsonProperty("RestockDateTime")
    val restockDateTime: String? = null,
    @JsonProperty("RetailerDeliveryNumber")
    val retailerDeliveryNumber: String? = null,
    @JsonProperty("RetailerDeliveryState")
    val retailerDeliveryState: String? = null,
    @JsonProperty("TotalPackages")
    val totalPackages: Int? = null,
    @JsonProperty("TotalPrice")
    val totalPrice: Int? = null,
    @JsonProperty("TotalPriceSold")
    val totalPriceSold: Int? = null,
    @JsonProperty("VehicleInfo")
    val vehicleInfo: String? = null,
    @JsonProperty("VehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("VehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("VehicleModel")
    val vehicleModel: String? = null,
    @JsonProperty("VoidedDate")
    val voidedDate: String? = null
)
