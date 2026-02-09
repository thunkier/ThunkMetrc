package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PatientsStatus(
    @JsonProperty("Active")
    val active: Boolean? = null,
    @JsonProperty("ConcentrateOuncesAllowed")
    val concentrateOuncesAllowed: Int? = null,
    @JsonProperty("ConcentrateOuncesAvailable")
    val concentrateOuncesAvailable: Int? = null,
    @JsonProperty("ConcentrateOuncesPurchased")
    val concentrateOuncesPurchased: Int? = null,
    @JsonProperty("FlowerOuncesAllowed")
    val flowerOuncesAllowed: Int? = null,
    @JsonProperty("FlowerOuncesAvailable")
    val flowerOuncesAvailable: Int? = null,
    @JsonProperty("FlowerOuncesPurchased")
    val flowerOuncesPurchased: Int? = null,
    @JsonProperty("IdentificationMethod")
    val identificationMethod: String? = null,
    @JsonProperty("InfusedOuncesAllowed")
    val infusedOuncesAllowed: Int? = null,
    @JsonProperty("InfusedOuncesAvailable")
    val infusedOuncesAvailable: Int? = null,
    @JsonProperty("InfusedOuncesPurchased")
    val infusedOuncesPurchased: Int? = null,
    @JsonProperty("PatientLicenseExpirationDate")
    val patientLicenseExpirationDate: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PatientRegistrationStartDate")
    val patientRegistrationStartDate: String? = null,
    @JsonProperty("PurchaseAmountDays")
    val purchaseAmountDays: Int? = null,
    @JsonProperty("RegistrationNumber")
    val registrationNumber: String? = null,
    @JsonProperty("ThcOuncesAllowed")
    val thcOuncesAllowed: Int? = null,
    @JsonProperty("ThcOuncesAvailable")
    val thcOuncesAvailable: Int? = null,
    @JsonProperty("ThcOuncesPurchased")
    val thcOuncesPurchased: Int? = null
)
