package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class UpdateItemRequest {
    @JsonProperty("Item")
    public String item;
    @JsonProperty("Label")
    public String label;
}
