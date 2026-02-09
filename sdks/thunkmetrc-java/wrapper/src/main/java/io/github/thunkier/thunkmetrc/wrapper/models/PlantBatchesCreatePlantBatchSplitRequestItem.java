package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantBatchesCreatePlantBatchSplitRequestItem {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("Count")
    public Integer count;
    @JsonProperty("GroupName")
    public String groupName;
    @JsonProperty("Location")
    public String location;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PlantBatch")
    public String plantBatch;
    @JsonProperty("Strain")
    public String strain;
    @JsonProperty("Sublocation")
    public String sublocation;
}
