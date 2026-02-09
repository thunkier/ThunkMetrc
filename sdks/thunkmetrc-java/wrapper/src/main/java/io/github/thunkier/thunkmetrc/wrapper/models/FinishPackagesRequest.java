package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class FinishPackagesRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("Label")
    public String label;
}
