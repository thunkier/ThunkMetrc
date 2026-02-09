package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantBatchesCreatePlantBatchesPlantingsRequestItem {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("Count")
    public Integer count;
    @JsonProperty("Location")
    public String location;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("SourcePlantBatches")
    public String sourcePlantBatches;
    @JsonProperty("Strain")
    public String strain;
    @JsonProperty("Sublocation")
    public String sublocation;
    @JsonProperty("Type")
    public String type;
}
