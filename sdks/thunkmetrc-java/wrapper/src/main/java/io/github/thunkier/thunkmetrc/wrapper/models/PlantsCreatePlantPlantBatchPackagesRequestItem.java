package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsCreatePlantPlantBatchPackagesRequestItem {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("Count")
    public Integer count;
    @JsonProperty("IsDonation")
    public Boolean isDonation;
    @JsonProperty("IsTradeSample")
    public Boolean isTradeSample;
    @JsonProperty("Item")
    public String item;
    @JsonProperty("Location")
    public String location;
    @JsonProperty("Note")
    public String note;
    @JsonProperty("PackageTag")
    public String packageTag;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PlantBatchType")
    public String plantBatchType;
    @JsonProperty("PlantLabel")
    public String plantLabel;
    @JsonProperty("Sublocation")
    public String sublocation;
}
