package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesUpdateSaleDeliveriesRetailerRequestItemDestinations(
    @JsonProperty("consumerId")
    val consumerId: String? = null,
    @JsonProperty("driverEmployeeId")
    val driverEmployeeId: Int? = null,
    @JsonProperty("driverName")
    val driverName: String? = null,
    @JsonProperty("driversLicenseNumber")
    val driversLicenseNumber: String? = null,
    @JsonProperty("estimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("estimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("id")
    val id: Int? = null,
    @JsonProperty("patientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("phoneNumberForQuestions")
    val phoneNumberForQuestions: String? = null,
    @JsonProperty("plannedRoute")
    val plannedRoute: String? = null,
    @JsonProperty("recipientAddressCity")
    val recipientAddressCity: String? = null,
    @JsonProperty("recipientAddressCounty")
    val recipientAddressCounty: String? = null,
    @JsonProperty("recipientAddressPostalCode")
    val recipientAddressPostalCode: String? = null,
    @JsonProperty("recipientAddressState")
    val recipientAddressState: String? = null,
    @JsonProperty("recipientAddressStreet1")
    val recipientAddressStreet1: String? = null,
    @JsonProperty("recipientAddressStreet2")
    val recipientAddressStreet2: String? = null,
    @JsonProperty("recipientName")
    val recipientName: String? = null,
    @JsonProperty("recipientZoneId")
    val recipientZoneId: String? = null,
    @JsonProperty("salesCustomerType")
    val salesCustomerType: String? = null,
    @JsonProperty("salesDateTime")
    val salesDateTime: String? = null,
    @JsonProperty("transactions")
    val transactions: List<SalesUpdateSaleDeliveriesRetailerRequestItemDestinationsTransactions>? = null,
    @JsonProperty("vehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("vehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("vehicleModel")
    val vehicleModel: String? = null
)
