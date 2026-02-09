package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreatePackagesPackagesRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("ExpirationDate")
    public String expirationDate;
    @JsonProperty("Ingredients")
    public List<Object> ingredients;
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
    @JsonProperty("ProductRequiresRemediation")
    public Boolean productRequiresRemediation;
    @JsonProperty("ProductionBatchNumber")
    public String productionBatchNumber;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("RequiredLabTestBatches")
    public Boolean requiredLabTestBatches;
    @JsonProperty("SellByDate")
    public String sellByDate;
    @JsonProperty("Sublocation")
    public String sublocation;
    @JsonProperty("Tag")
    public String tag;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
    @JsonProperty("UseByDate")
    public String useByDate;
    @JsonProperty("UseSameItem")
    public Boolean useSameItem;
    public static class CreatePackagesPackagesRequestIngredientsItem {
    @JsonProperty("Package")
    public String package_;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
}
