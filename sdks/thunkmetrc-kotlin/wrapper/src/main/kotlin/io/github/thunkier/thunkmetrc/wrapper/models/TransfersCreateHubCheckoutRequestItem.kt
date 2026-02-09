package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersCreateHubCheckoutRequestItem(
    @JsonProperty("ShipmentDeliveryId")
    val shipmentDeliveryId: Int? = null,
    @JsonProperty("TransporterDirection")
    val transporterDirection: String? = null
)
