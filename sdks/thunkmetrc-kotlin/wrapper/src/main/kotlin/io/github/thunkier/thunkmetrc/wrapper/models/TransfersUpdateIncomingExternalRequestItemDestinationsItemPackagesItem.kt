package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersUpdateIncomingExternalRequestItemDestinationsItemPackagesItem(
    @JsonProperty("ExpirationDate")
    val expirationDate: String? = null,
    @JsonProperty("ExternalId")
    val externalId: String? = null,
    @JsonProperty("GrossUnitOfWeightName")
    val grossUnitOfWeightName: String? = null,
    @JsonProperty("GrossWeight")
    val grossWeight: Double? = null,
    @JsonProperty("IsFinishedGood")
    val isFinishedGood: Boolean? = null,
    @JsonProperty("ItemName")
    val itemName: String? = null,
    @JsonProperty("PackagedDate")
    val packagedDate: String? = null,
    @JsonProperty("Quantity")
    val quantity: Double? = null,
    @JsonProperty("SellByDate")
    val sellByDate: String? = null,
    @JsonProperty("UnitOfMeasureName")
    val unitOfMeasureName: String? = null,
    @JsonProperty("UseByDate")
    val useByDate: String? = null,
    @JsonProperty("WholesalePrice")
    val wholesalePrice: String? = null
)
