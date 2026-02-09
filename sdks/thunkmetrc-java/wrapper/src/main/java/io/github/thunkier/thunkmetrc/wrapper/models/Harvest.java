package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class Harvest {
    @JsonProperty("ArchivedDate")
    public String archivedDate;
    @JsonProperty("CurrentWeight")
    public Double currentWeight;
    @JsonProperty("DryingLocationId")
    public Integer dryingLocationId;
    @JsonProperty("DryingLocationName")
    public String dryingLocationName;
    @JsonProperty("DryingLocationTypeName")
    public String dryingLocationTypeName;
    @JsonProperty("DryingSublocationId")
    public Integer dryingSublocationId;
    @JsonProperty("DryingSublocationName")
    public String dryingSublocationName;
    @JsonProperty("FinishedDate")
    public String finishedDate;
    @JsonProperty("HarvestStartDate")
    public String harvestStartDate;
    @JsonProperty("HarvestType")
    public String harvestType;
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
    @JsonProperty("LabTestingState")
    public String labTestingState;
    @JsonProperty("LabTestingStateDate")
    public String labTestingStateDate;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("PackageCount")
    public Integer packageCount;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PlantCount")
    public Integer plantCount;
    @JsonProperty("SourceStrainCount")
    public Integer sourceStrainCount;
    @JsonProperty("SourceStrainNames")
    public String sourceStrainNames;
    @JsonProperty("Strains")
    public List<Object> strains;
    @JsonProperty("TotalPackagedWeight")
    public Double totalPackagedWeight;
    @JsonProperty("TotalRestoredWeight")
    public Double totalRestoredWeight;
    @JsonProperty("TotalWasteWeight")
    public Double totalWasteWeight;
    @JsonProperty("TotalWetWeight")
    public Double totalWetWeight;
    @JsonProperty("UnitOfWeightName")
    public String unitOfWeightName;
}
