package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesCreateDeliveriesRetailerRestockRequestPackagesItem {
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("RemoveCurrentPackage")
    public Boolean removeCurrentPackage;
    @JsonProperty("TotalPrice")
    public Double totalPrice;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
