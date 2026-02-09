package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class PackageSourceHarvest(
    @JsonProperty("HarvestStartDate")
    val harvestStartDate: String? = null,
    @JsonProperty("HarvestedByFacilityLicenseNumber")
    val harvestedByFacilityLicenseNumber: String? = null,
    @JsonProperty("HarvestedByFacilityName")
    val harvestedByFacilityName: String? = null,
    @JsonProperty("Name")
    val name: String? = null
)
