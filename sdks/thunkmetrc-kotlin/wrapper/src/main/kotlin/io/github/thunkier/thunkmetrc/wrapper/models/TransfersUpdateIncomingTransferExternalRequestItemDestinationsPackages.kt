package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersUpdateIncomingTransferExternalRequestItemDestinationsPackages(
    @JsonProperty("expirationDate")
    val expirationDate: String? = null,
    @JsonProperty("externalId")
    val externalId: String? = null,
    @JsonProperty("grossUnitOfWeightName")
    val grossUnitOfWeightName: String? = null,
    @JsonProperty("grossWeight")
    val grossWeight: Double? = null,
    @JsonProperty("itemName")
    val itemName: String? = null,
    @JsonProperty("packagedDate")
    val packagedDate: String? = null,
    @JsonProperty("quantity")
    val quantity: Int? = null,
    @JsonProperty("sellByDate")
    val sellByDate: String? = null,
    @JsonProperty("unitOfMeasureName")
    val unitOfMeasureName: String? = null,
    @JsonProperty("useByDate")
    val useByDate: String? = null,
    @JsonProperty("wholesalePrice")
    val wholesalePrice: String? = null
)
