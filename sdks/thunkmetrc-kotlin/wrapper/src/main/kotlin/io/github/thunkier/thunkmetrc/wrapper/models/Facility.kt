package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Facility(
    @JsonProperty("Alias")
    val alias: String? = null,
    @JsonProperty("AllowTransferOfOnHoldPackages")
    val allowTransferOfOnHoldPackages: Boolean? = null,
    @JsonProperty("AllowTransferOfOnRecallPackages")
    val allowTransferOfOnRecallPackages: Boolean? = null,
    @JsonProperty("CredentialedDate")
    val credentialedDate: String? = null,
    @JsonProperty("DisplayName")
    val displayName: String? = null,
    @JsonProperty("Email")
    val email: String? = null,
    @JsonProperty("FacilityId")
    val facilityId: Int? = null,
    @JsonProperty("FacilityType")
    val facilityType: FacilityFacilityType? = null,
    @JsonProperty("HireDate")
    val hireDate: String? = null,
    @JsonProperty("IsFinancialContact")
    val isFinancialContact: Boolean? = null,
    @JsonProperty("IsManager")
    val isManager: Boolean? = null,
    @JsonProperty("IsOwner")
    val isOwner: Boolean? = null,
    @JsonProperty("License")
    val license: FacilityLicense? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("Occupations")
    val occupations: List<Any>? = null,
    @JsonProperty("SupportActivationDate")
    val supportActivationDate: String? = null,
    @JsonProperty("SupportExpirationDate")
    val supportExpirationDate: String? = null
)
