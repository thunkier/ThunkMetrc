package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class UpdateExternalidRequest {
    @JsonProperty("ExternalId")
    public String externalId;
    @JsonProperty("PackageLabel")
    public String packageLabel;
}
