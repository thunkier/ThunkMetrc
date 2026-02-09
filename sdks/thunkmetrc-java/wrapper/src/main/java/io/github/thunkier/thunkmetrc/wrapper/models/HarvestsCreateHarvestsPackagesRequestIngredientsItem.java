package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class HarvestsCreateHarvestsPackagesRequestIngredientsItem {
    @JsonProperty("HarvestId")
    public Integer harvestId;
    @JsonProperty("HarvestName")
    public String harvestName;
    @JsonProperty("UnitOfWeight")
    public String unitOfWeight;
    @JsonProperty("Weight")
    public Double weight;
}
