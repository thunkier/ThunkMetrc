package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class HarvestsCreateHarvestPackagesRequestItem {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("DecontaminateProduct")
    public Boolean decontaminateProduct;
    @JsonProperty("DecontaminationDate")
    public String decontaminationDate;
    @JsonProperty("DecontaminationSteps")
    public String decontaminationSteps;
    @JsonProperty("ExpirationDate")
    public String expirationDate;
    @JsonProperty("Ingredients")
    public List<HarvestsCreateHarvestPackagesRequestItemIngredients> ingredients;
    @JsonProperty("IsDonation")
    public Boolean isDonation;
    @JsonProperty("IsProductionBatch")
    public Boolean isProductionBatch;
    @JsonProperty("IsTradeSample")
    public Boolean isTradeSample;
    @JsonProperty("Item")
    public String item;
    @JsonProperty("LabTestStageId")
    public Integer labTestStageId;
    @JsonProperty("Location")
    public String location;
    @JsonProperty("Note")
    public String note;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("ProcessingJobTypeId")
    public Integer processingJobTypeId;
    @JsonProperty("ProductRequiresDecontamination")
    public Boolean productRequiresDecontamination;
    @JsonProperty("ProductRequiresRemediation")
    public Boolean productRequiresRemediation;
    @JsonProperty("ProductionBatchNumber")
    public String productionBatchNumber;
    @JsonProperty("RemediateProduct")
    public Boolean remediateProduct;
    @JsonProperty("RemediationDate")
    public String remediationDate;
    @JsonProperty("RemediationMethodId")
    public Integer remediationMethodId;
    @JsonProperty("RemediationSteps")
    public String remediationSteps;
    @JsonProperty("RequiredLabTestBatches")
    public List<Object> requiredLabTestBatches;
    @JsonProperty("SellByDate")
    public String sellByDate;
    @JsonProperty("Sublocation")
    public String sublocation;
    @JsonProperty("Tag")
    public String tag;
    @JsonProperty("UnitOfWeight")
    public String unitOfWeight;
    @JsonProperty("UseByDate")
    public String useByDate;
    public static class HarvestsCreateHarvestPackagesRequestItemIngredients {
    @JsonProperty("HarvestId")
    public Integer harvestId;
    @JsonProperty("HarvestName")
    public String harvestName;
    @JsonProperty("UnitOfWeight")
    public String unitOfWeight;
    @JsonProperty("Weight")
    public Double weight;
}
}
