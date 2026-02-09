package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class LocationsUpdateLocationsRequest {
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("LocationTypeName")
    public String locationTypeName;
    @JsonProperty("Name")
    public String name;
}
