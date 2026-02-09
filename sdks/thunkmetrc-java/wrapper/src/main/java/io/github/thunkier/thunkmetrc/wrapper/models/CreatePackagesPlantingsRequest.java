package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreatePackagesPlantingsRequest {
    @JsonProperty("LocationName")
    public String locationName;
    @JsonProperty("PackageAdjustmentAmount")
    public Integer packageAdjustmentAmount;
    @JsonProperty("PackageAdjustmentUnitOfMeasureName")
    public String packageAdjustmentUnitOfMeasureName;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PlantBatchName")
    public String plantBatchName;
    @JsonProperty("PlantBatchType")
    public String plantBatchType;
    @JsonProperty("PlantCount")
    public Integer plantCount;
    @JsonProperty("PlantedDate")
    public String plantedDate;
    @JsonProperty("StrainName")
    public String strainName;
    @JsonProperty("SublocationName")
    public String sublocationName;
    @JsonProperty("UnpackagedDate")
    public String unpackagedDate;
}
