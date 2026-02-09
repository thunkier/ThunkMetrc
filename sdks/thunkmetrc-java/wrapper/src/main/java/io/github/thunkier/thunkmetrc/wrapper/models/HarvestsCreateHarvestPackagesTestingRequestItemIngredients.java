package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class HarvestsCreateHarvestPackagesTestingRequestItemIngredients {
    @JsonProperty("HarvestId")
    public Integer harvestId;
    @JsonProperty("HarvestName")
    public String harvestName;
    @JsonProperty("UnitOfWeight")
    public String unitOfWeight;
    @JsonProperty("Weight")
    public Double weight;
}
