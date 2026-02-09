package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateSaleReceiptsRequestItem(
    @JsonProperty("caregiverLicenseNumber")
    val caregiverLicenseNumber: String? = null,
    @JsonProperty("externalReceiptNumber")
    val externalReceiptNumber: String? = null,
    @JsonProperty("identificationMethod")
    val identificationMethod: String? = null,
    @JsonProperty("patientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("patientRegistrationLocationId")
    val patientRegistrationLocationId: Int? = null,
    @JsonProperty("salesCustomerType")
    val salesCustomerType: String? = null,
    @JsonProperty("salesDateTime")
    val salesDateTime: String? = null,
    @JsonProperty("transactions")
    val transactions: List<SalesCreateSaleReceiptsRequestItemTransactions>? = null
)
