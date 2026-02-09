package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsUpdatePlantsLocationRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("Location")
    public String location;
    @JsonProperty("Sublocation")
    public String sublocation;
}
