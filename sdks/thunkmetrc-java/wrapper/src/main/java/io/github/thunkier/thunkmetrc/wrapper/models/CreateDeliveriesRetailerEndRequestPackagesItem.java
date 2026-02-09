package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreateDeliveriesRetailerEndRequestPackagesItem {
    @JsonProperty("EndQuantity")
    public Integer endQuantity;
    @JsonProperty("EndUnitOfMeasure")
    public String endUnitOfMeasure;
    @JsonProperty("Label")
    public String label;
}
