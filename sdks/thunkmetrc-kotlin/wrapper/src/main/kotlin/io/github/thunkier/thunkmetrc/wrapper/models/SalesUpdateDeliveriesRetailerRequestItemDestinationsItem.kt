package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesUpdateDeliveriesRetailerRequestItemDestinationsItem(
    @JsonProperty("ConsumerId")
    val consumerId: String? = null,
    @JsonProperty("DriverEmployeeId")
    val driverEmployeeId: Int? = null,
    @JsonProperty("DriverName")
    val driverName: String? = null,
    @JsonProperty("DriversLicenseNumber")
    val driversLicenseNumber: String? = null,
    @JsonProperty("EstimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("EstimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PhoneNumberForQuestions")
    val phoneNumberForQuestions: String? = null,
    @JsonProperty("PlannedRoute")
    val plannedRoute: String? = null,
    @JsonProperty("RecipientAddressCity")
    val recipientAddressCity: String? = null,
    @JsonProperty("RecipientAddressCounty")
    val recipientAddressCounty: String? = null,
    @JsonProperty("RecipientAddressPostalCode")
    val recipientAddressPostalCode: String? = null,
    @JsonProperty("RecipientAddressState")
    val recipientAddressState: String? = null,
    @JsonProperty("RecipientAddressStreet1")
    val recipientAddressStreet1: String? = null,
    @JsonProperty("RecipientAddressStreet2")
    val recipientAddressStreet2: String? = null,
    @JsonProperty("RecipientName")
    val recipientName: String? = null,
    @JsonProperty("RecipientZoneId")
    val recipientZoneId: String? = null,
    @JsonProperty("SalesCustomerType")
    val salesCustomerType: String? = null,
    @JsonProperty("SalesDateTime")
    val salesDateTime: String? = null,
    @JsonProperty("Transactions")
    val transactions: List<Any>? = null,
    @JsonProperty("VehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("VehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("VehicleModel")
    val vehicleModel: String? = null
)
