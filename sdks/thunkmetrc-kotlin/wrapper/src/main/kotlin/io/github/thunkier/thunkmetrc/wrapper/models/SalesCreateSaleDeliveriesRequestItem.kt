package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateSaleDeliveriesRequestItem(
    @JsonProperty("consumerId")
    val consumerId: Int? = null,
    @JsonProperty("driverEmployeeId")
    val driverEmployeeId: String? = null,
    @JsonProperty("driverName")
    val driverName: String? = null,
    @JsonProperty("driversLicenseNumber")
    val driversLicenseNumber: String? = null,
    @JsonProperty("estimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("estimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
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
    val recipientZoneId: Int? = null,
    @JsonProperty("salesCustomerType")
    val salesCustomerType: String? = null,
    @JsonProperty("salesDateTime")
    val salesDateTime: String? = null,
    @JsonProperty("transactions")
    val transactions: List<SalesCreateSaleDeliveriesRequestItemTransactions>? = null,
    @JsonProperty("transporterFacilityLicenseNumber")
    val transporterFacilityLicenseNumber: String? = null,
    @JsonProperty("vehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("vehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("vehicleModel")
    val vehicleModel: String? = null
)
