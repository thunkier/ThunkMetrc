package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class LocationsLocation {
    @JsonProperty("ForHarvests")
    public Boolean forHarvests;
    @JsonProperty("ForPackages")
    public Boolean forPackages;
    @JsonProperty("ForPlantBatches")
    public Boolean forPlantBatches;
    @JsonProperty("ForPlants")
    public Boolean forPlants;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("LocationTypeId")
    public Integer locationTypeId;
    @JsonProperty("LocationTypeName")
    public String locationTypeName;
    @JsonProperty("Name")
    public String name;
}
