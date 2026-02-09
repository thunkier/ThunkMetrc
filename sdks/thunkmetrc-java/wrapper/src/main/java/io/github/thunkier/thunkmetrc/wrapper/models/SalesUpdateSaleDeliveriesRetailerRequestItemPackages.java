package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesUpdateSaleDeliveriesRetailerRequestItemPackages {
    @JsonProperty("DateTime")
    public String dateTime;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("TotalPrice")
    public Double totalPrice;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
