package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantBatchType {
    @JsonProperty("CanBeCloned")
    public Boolean canBeCloned;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("Name")
    public String name;
}
