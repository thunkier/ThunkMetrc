package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransferDeliveryPackageWholesale(
    @JsonProperty("PackageId")
    val packageId: Int? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("ReceiverWholesalePrice")
    val receiverWholesalePrice: String? = null,
    @JsonProperty("ShipperWholesalePrice")
    val shipperWholesalePrice: String? = null
)
