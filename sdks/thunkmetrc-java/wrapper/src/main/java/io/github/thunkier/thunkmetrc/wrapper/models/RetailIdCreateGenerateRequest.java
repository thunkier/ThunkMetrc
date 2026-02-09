package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class RetailIdCreateGenerateRequest {
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("Quantity")
    public Double quantity;
}
