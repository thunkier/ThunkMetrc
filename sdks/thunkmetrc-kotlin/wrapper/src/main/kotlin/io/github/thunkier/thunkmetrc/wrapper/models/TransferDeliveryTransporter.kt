package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransferDeliveryTransporter(
    @JsonProperty("TransporterDirection")
    val transporterDirection: String? = null,
    @JsonProperty("TransporterFacilityLicenseNumber")
    val transporterFacilityLicenseNumber: String? = null,
    @JsonProperty("TransporterFacilityName")
    val transporterFacilityName: String? = null
)
