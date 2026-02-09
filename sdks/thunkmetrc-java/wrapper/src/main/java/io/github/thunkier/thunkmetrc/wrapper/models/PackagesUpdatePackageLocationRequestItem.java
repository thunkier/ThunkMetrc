package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackagesUpdatePackageLocationRequestItem {
    @JsonProperty("Label")
    public String label;
    @JsonProperty("Location")
    public String location;
    @JsonProperty("MoveDate")
    public String moveDate;
    @JsonProperty("Sublocation")
    public String sublocation;
}
