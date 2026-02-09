package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreateGrowthPhaseRequest {
    @JsonProperty("Count")
    public Integer count;
    @JsonProperty("GrowthDate")
    public String growthDate;
    @JsonProperty("GrowthPhase")
    public String growthPhase;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("NewLocation")
    public String newLocation;
    @JsonProperty("NewSublocation")
    public String newSublocation;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("StartingTag")
    public String startingTag;
}
