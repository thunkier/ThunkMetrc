package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantBatchesCreatePackagesFromMotherPlantRequestItem {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("Count")
    public Integer count;
    @JsonProperty("ExpirationDate")
    public String expirationDate;
    @JsonProperty("Id")
    public Integer id;
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
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PlantBatch")
    public String plantBatch;
    @JsonProperty("SellByDate")
    public String sellByDate;
    @JsonProperty("Sublocation")
    public String sublocation;
    @JsonProperty("Tag")
    public String tag;
    @JsonProperty("UseByDate")
    public String useByDate;
}
