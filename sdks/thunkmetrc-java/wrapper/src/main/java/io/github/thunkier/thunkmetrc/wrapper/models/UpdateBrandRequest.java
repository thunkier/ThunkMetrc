package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class UpdateBrandRequest {
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Name")
    public String name;
}
