package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesUpdateSaleDeliveriesRetailerRequestItem(
    @JsonProperty("dateTime")
    val dateTime: String? = null,
    @JsonProperty("destinations")
    val destinations: List<SalesUpdateSaleDeliveriesRetailerRequestItemDestinations>? = null,
    @JsonProperty("driverEmployeeId")
    val driverEmployeeId: String? = null,
    @JsonProperty("driverName")
    val driverName: String? = null,
    @JsonProperty("driversLicenseNumber")
    val driversLicenseNumber: String? = null,
    @JsonProperty("estimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("packages")
    val packages: List<SalesUpdateSaleDeliveriesRetailerRequestItemPackages>? = null,
    @JsonProperty("phoneNumberForQuestions")
    val phoneNumberForQuestions: String? = null,
    @JsonProperty("vehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("vehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("vehicleModel")
    val vehicleModel: String? = null
)
