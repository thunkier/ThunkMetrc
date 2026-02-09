package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersCreateHubCheckinRequestItem(
    @JsonProperty("ShipmentDeliveryId")
    val shipmentDeliveryId: Int? = null,
    @JsonProperty("TransporterDirection")
    val transporterDirection: String? = null
)
