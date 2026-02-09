package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackagesCreateTestingRequestItemIngredientsItem {
    @JsonProperty("Package")
    public String package_;
    @JsonProperty("Quantity")
    public Double quantity;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
