package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesCreateDeliveriesRetailerRestockRequestItem {
    @JsonProperty("DateTime")
    public String dateTime;
    @JsonProperty("Destinations")
    public String destinations;
    @JsonProperty("EstimatedDepartureDateTime")
    public String estimatedDepartureDateTime;
    @JsonProperty("Packages")
    public List<Object> packages;
    @JsonProperty("RetailerDeliveryId")
    public Integer retailerDeliveryId;
    public static class SalesCreateDeliveriesRetailerRestockRequestItemPackagesItem {
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("Quantity")
    public Double quantity;
    @JsonProperty("RemoveCurrentPackage")
    public Boolean removeCurrentPackage;
    @JsonProperty("TotalPrice")
    public Double totalPrice;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
}
