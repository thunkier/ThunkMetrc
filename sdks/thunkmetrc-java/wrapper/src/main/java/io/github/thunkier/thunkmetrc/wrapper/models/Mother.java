package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class Mother {
    @JsonProperty("ClonedCount")
    public Integer clonedCount;
    @JsonProperty("DescendedCount")
    public Integer descendedCount;
    @JsonProperty("DestroyedByUserName")
    public String destroyedByUserName;
    @JsonProperty("DestroyedDate")
    public String destroyedDate;
    @JsonProperty("DestroyedNote")
    public String destroyedNote;
    @JsonProperty("FloweringDate")
    public String floweringDate;
    @JsonProperty("GroupTagTypeMax")
    public Integer groupTagTypeMax;
    @JsonProperty("GrowthPhase")
    public String growthPhase;
    @JsonProperty("HarvestCount")
    public Integer harvestCount;
    @JsonProperty("HarvestId")
    public Integer harvestId;
    @JsonProperty("HarvestedDate")
    public String harvestedDate;
    @JsonProperty("HarvestedUnitOfWeightAbbreviation")
    public String harvestedUnitOfWeightAbbreviation;
    @JsonProperty("HarvestedUnitOfWeightName")
    public String harvestedUnitOfWeightName;
    @JsonProperty("HarvestedWetWeight")
    public Double harvestedWetWeight;
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
    @JsonProperty("Label")
    public String label;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("LocationId")
    public Integer locationId;
    @JsonProperty("LocationName")
    public String locationName;
    @JsonProperty("LocationTypeName")
    public String locationTypeName;
    @JsonProperty("MotherPlantDate")
    public String motherPlantDate;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PlantBatchId")
    public Integer plantBatchId;
    @JsonProperty("PlantBatchName")
    public String plantBatchName;
    @JsonProperty("PlantBatchTypeId")
    public Integer plantBatchTypeId;
    @JsonProperty("PlantBatchTypeName")
    public String plantBatchTypeName;
    @JsonProperty("PlantedDate")
    public String plantedDate;
    @JsonProperty("State")
    public String state;
    @JsonProperty("StrainId")
    public Integer strainId;
    @JsonProperty("StrainName")
    public String strainName;
    @JsonProperty("SublocationId")
    public Integer sublocationId;
    @JsonProperty("SublocationName")
    public String sublocationName;
    @JsonProperty("SurvivedCount")
    public Integer survivedCount;
    @JsonProperty("TagTypeMax")
    public Integer tagTypeMax;
    @JsonProperty("VegetativeDate")
    public String vegetativeDate;
}
