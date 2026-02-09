package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreatePackagesInfoRequest {
    @JsonProperty("packageLabels")
    public List<Object> packageLabels;
}
