package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantBatchesUpdatePlantBatchesLocationRequestItem {
    @JsonProperty("Location")
    public String location;
    @JsonProperty("MoveDate")
    public String moveDate;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("Sublocation")
    public String sublocation;
}
