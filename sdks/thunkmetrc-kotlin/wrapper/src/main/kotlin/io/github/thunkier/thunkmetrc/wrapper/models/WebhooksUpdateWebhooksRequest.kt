package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class WebhooksUpdateWebhooksRequest(
    @JsonProperty("errorResponseJsonTemplate")
    val errorResponseJsonTemplate: String? = null,
    @JsonProperty("facilityLicenseNumbers")
    val facilityLicenseNumbers: String? = null,
    @JsonProperty("objectType")
    val objectType: String? = null,
    @JsonProperty("serverPublicKeyFingerprint")
    val serverPublicKeyFingerprint: String? = null,
    @JsonProperty("status")
    val status: String? = null,
    @JsonProperty("template")
    val template: String? = null,
    @JsonProperty("tpiApiKey")
    val tpiApiKey: String? = null,
    @JsonProperty("url")
    val url: String? = null,
    @JsonProperty("userApiKey")
    val userApiKey: String? = null,
    @JsonProperty("verb")
    val verb: String? = null
)
