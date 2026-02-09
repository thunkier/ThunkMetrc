package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreateGenerateRequest {
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("Quantity")
    public Integer quantity;
}
