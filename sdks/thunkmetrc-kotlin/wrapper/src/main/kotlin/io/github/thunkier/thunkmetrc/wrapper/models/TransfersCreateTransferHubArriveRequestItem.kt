package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class TransfersCreateTransferHubArriveRequestItem(
    @JsonProperty("shipmentDeliveryId")
    val shipmentDeliveryId: Int? = null,
    @JsonProperty("transporterDirection")
    val transporterDirection: String? = null
)
