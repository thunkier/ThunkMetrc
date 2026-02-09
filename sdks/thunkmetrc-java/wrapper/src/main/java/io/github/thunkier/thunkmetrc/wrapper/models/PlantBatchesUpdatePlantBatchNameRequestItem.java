package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantBatchesUpdatePlantBatchNameRequestItem {
    @JsonProperty("Group")
    public String group;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("NewGroup")
    public String newGroup;
}
