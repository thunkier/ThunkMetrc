package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class WriteResponse {
    @JsonProperty("Ids")
    public List<Integer> ids;
    @JsonProperty("Warnings")
    public String warnings;
}
