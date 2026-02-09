package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SaleReceipt(
    @JsonProperty("ArchivedDate")
    val archivedDate: String? = null,
    @JsonProperty("CaregiverLicenseNumber")
    val caregiverLicenseNumber: String? = null,
    @JsonProperty("ExternalReceiptNumber")
    val externalReceiptNumber: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IdentificationMethod")
    val identificationMethod: String? = null,
    @JsonProperty("IsFinal")
    val isFinal: Boolean? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PatientRegistrationLocationId")
    val patientRegistrationLocationId: Int? = null,
    @JsonProperty("ReceiptNumber")
    val receiptNumber: String? = null,
    @JsonProperty("RecordedByUserName")
    val recordedByUserName: String? = null,
    @JsonProperty("RecordedDateTime")
    val recordedDateTime: String? = null,
    @JsonProperty("SalesCustomerType")
    val salesCustomerType: String? = null,
    @JsonProperty("SalesDateTime")
    val salesDateTime: String? = null,
    @JsonProperty("TotalPackages")
    val totalPackages: Int? = null,
    @JsonProperty("TotalPrice")
    val totalPrice: Int? = null,
    @JsonProperty("Transactions")
    val transactions: List<Any>? = null
)
