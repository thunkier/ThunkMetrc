package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreateDeliveriesRetailerEndRequest {
    @JsonProperty("ActualArrivalDateTime")
    public String actualArrivalDateTime;
    @JsonProperty("Packages")
    public List<Object> packages;
    @JsonProperty("RetailerDeliveryId")
    public Integer retailerDeliveryId;
    public static class CreateDeliveriesRetailerEndRequestPackagesItem {
    @JsonProperty("EndQuantity")
    public Integer endQuantity;
    @JsonProperty("EndUnitOfMeasure")
    public String endUnitOfMeasure;
    @JsonProperty("Label")
    public String label;
}
}
