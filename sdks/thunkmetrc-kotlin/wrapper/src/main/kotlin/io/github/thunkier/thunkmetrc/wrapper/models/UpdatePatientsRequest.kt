package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class UpdatePatientsRequest(
    @JsonProperty("ActualDate")
    val actualDate: String? = null,
    @JsonProperty("ConcentrateOuncesAllowed")
    val concentrateOuncesAllowed: Int? = null,
    @JsonProperty("FlowerOuncesAllowed")
    val flowerOuncesAllowed: Int? = null,
    @JsonProperty("HasSalesLimitExemption")
    val hasSalesLimitExemption: Boolean? = null,
    @JsonProperty("InfusedOuncesAllowed")
    val infusedOuncesAllowed: Int? = null,
    @JsonProperty("LicenseEffectiveEndDate")
    val licenseEffectiveEndDate: String? = null,
    @JsonProperty("LicenseEffectiveStartDate")
    val licenseEffectiveStartDate: String? = null,
    @JsonProperty("LicenseNumber")
    val licenseNumber: String? = null,
    @JsonProperty("MaxConcentrateThcPercentAllowed")
    val maxConcentrateThcPercentAllowed: Int? = null,
    @JsonProperty("MaxFlowerThcPercentAllowed")
    val maxFlowerThcPercentAllowed: Int? = null,
    @JsonProperty("NewLicenseNumber")
    val newLicenseNumber: String? = null,
    @JsonProperty("RecommendedPlants")
    val recommendedPlants: Int? = null,
    @JsonProperty("RecommendedSmokableQuantity")
    val recommendedSmokableQuantity: Int? = null,
    @JsonProperty("ThcOuncesAllowed")
    val thcOuncesAllowed: Int? = null
)
