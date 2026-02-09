package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersCreateTransferHubCheckinRequestItem(
    @JsonProperty("shipmentDeliveryId")
    val shipmentDeliveryId: Int? = null,
    @JsonProperty("transporterDirection")
    val transporterDirection: String? = null
)
