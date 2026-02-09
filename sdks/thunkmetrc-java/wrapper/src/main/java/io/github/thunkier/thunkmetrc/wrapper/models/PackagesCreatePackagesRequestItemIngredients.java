package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackagesCreatePackagesRequestItemIngredients {
    @JsonProperty("Package")
    public String package_;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
