package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class WebhooksUpdateWebhooksRequest {
    @JsonProperty("errorResponseJsonTemplate")
    public String errorResponseJsonTemplate;
    @JsonProperty("facilityLicenseNumbers")
    public String facilityLicenseNumbers;
    @JsonProperty("objectType")
    public String objectType;
    @JsonProperty("serverPublicKeyFingerprint")
    public String serverPublicKeyFingerprint;
    @JsonProperty("status")
    public String status;
    @JsonProperty("template")
    public String template;
    @JsonProperty("tpiApiKey")
    public String tpiApiKey;
    @JsonProperty("url")
    public String url;
    @JsonProperty("userApiKey")
    public String userApiKey;
    @JsonProperty("verb")
    public String verb;
}
