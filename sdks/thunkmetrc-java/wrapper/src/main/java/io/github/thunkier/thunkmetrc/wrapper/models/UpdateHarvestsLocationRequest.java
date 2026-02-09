package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class UpdateHarvestsLocationRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("DryingLocation")
    public String dryingLocation;
    @JsonProperty("DryingSublocation")
    public String dryingSublocation;
    @JsonProperty("HarvestName")
    public String harvestName;
    @JsonProperty("Id")
    public Integer id;
}
