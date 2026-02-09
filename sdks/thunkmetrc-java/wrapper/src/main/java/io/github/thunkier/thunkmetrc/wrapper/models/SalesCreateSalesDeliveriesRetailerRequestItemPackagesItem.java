package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesCreateSalesDeliveriesRetailerRequestItemPackagesItem {
    @JsonProperty("DateTime")
    public String dateTime;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("Quantity")
    public Double quantity;
    @JsonProperty("TotalPrice")
    public Double totalPrice;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
