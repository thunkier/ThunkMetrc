package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantBatch {
    @JsonProperty("DestroyedCount")
    public Integer destroyedCount;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IsOnHold")
    public Boolean isOnHold;
    @JsonProperty("IsOnInvestigation")
    public Boolean isOnInvestigation;
    @JsonProperty("IsOnInvestigationHold")
    public Boolean isOnInvestigationHold;
    @JsonProperty("IsOnInvestigationRecall")
    public Boolean isOnInvestigationRecall;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("LocationId")
    public Integer locationId;
    @JsonProperty("LocationName")
    public String locationName;
    @JsonProperty("LocationTypeName")
    public String locationTypeName;
    @JsonProperty("MultiPlantBatch")
    public Boolean multiPlantBatch;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("PackagedCount")
    public Integer packagedCount;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PlantBatchTypeId")
    public Integer plantBatchTypeId;
    @JsonProperty("PlantBatchTypeName")
    public String plantBatchTypeName;
    @JsonProperty("PlantedDate")
    public String plantedDate;
    @JsonProperty("SourcePackageId")
    public Integer sourcePackageId;
    @JsonProperty("SourcePackageLabel")
    public String sourcePackageLabel;
    @JsonProperty("SourcePlantBatchIds")
    public List<Object> sourcePlantBatchIds;
    @JsonProperty("SourcePlantBatchNames")
    public String sourcePlantBatchNames;
    @JsonProperty("SourcePlantId")
    public Integer sourcePlantId;
    @JsonProperty("SourcePlantLabel")
    public String sourcePlantLabel;
    @JsonProperty("StrainId")
    public Integer strainId;
    @JsonProperty("StrainName")
    public String strainName;
    @JsonProperty("SublocationId")
    public Integer sublocationId;
    @JsonProperty("SublocationName")
    public String sublocationName;
    @JsonProperty("TrackedCount")
    public Integer trackedCount;
    @JsonProperty("UntrackedCount")
    public Integer untrackedCount;
}
