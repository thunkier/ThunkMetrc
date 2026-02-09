package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class UpdateGrowthPhaseRequest {
    @JsonProperty("GrowthDate")
    public String growthDate;
    @JsonProperty("GrowthPhase")
    public String growthPhase;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("NewLocation")
    public String newLocation;
    @JsonProperty("NewSublocation")
    public String newSublocation;
    @JsonProperty("NewTag")
    public String newTag;
}
