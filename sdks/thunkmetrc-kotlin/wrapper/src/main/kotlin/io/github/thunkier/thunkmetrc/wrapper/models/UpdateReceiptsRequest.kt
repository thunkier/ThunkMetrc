package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdateReceiptsRequest(
    @JsonProperty("CaregiverLicenseNumber")
    val caregiverLicenseNumber: String? = null,
    @JsonProperty("ExternalReceiptNumber")
    val externalReceiptNumber: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IdentificationMethod")
    val identificationMethod: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PatientRegistrationLocationId")
    val patientRegistrationLocationId: Int? = null,
    @JsonProperty("SalesCustomerType")
    val salesCustomerType: String? = null,
    @JsonProperty("SalesDateTime")
    val salesDateTime: String? = null,
    @JsonProperty("Transactions")
    val transactions: List<Any>? = null
)
