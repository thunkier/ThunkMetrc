package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class UnitsOfMeasure {
    @JsonProperty("Abbreviation")
    public String abbreviation;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("QuantityType")
    public String quantityType;
}
