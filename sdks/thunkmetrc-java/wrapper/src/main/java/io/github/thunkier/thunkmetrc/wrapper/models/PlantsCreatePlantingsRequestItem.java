package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsCreatePlantingsRequestItem {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("LocationName")
    public String locationName;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PlantBatchName")
    public String plantBatchName;
    @JsonProperty("PlantBatchType")
    public String plantBatchType;
    @JsonProperty("PlantCount")
    public Integer plantCount;
    @JsonProperty("PlantLabel")
    public String plantLabel;
    @JsonProperty("StrainName")
    public String strainName;
    @JsonProperty("SublocationName")
    public String sublocationName;
}
