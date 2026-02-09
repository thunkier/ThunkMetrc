package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantWasteMethodsAll {
    @JsonProperty("ForPlants")
    public Boolean forPlants;
    @JsonProperty("ForProductDestruction")
    public Boolean forProductDestruction;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("Name")
    public String name;
}
