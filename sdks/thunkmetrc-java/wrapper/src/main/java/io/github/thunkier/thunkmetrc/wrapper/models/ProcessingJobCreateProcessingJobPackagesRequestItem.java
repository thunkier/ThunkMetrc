package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class ProcessingJobCreateProcessingJobPackagesRequestItem {
    @JsonProperty("ExpirationDate")
    public String expirationDate;
    @JsonProperty("FinishDate")
    public String finishDate;
    @JsonProperty("FinishNote")
    public String finishNote;
    @JsonProperty("FinishProcessingJob")
    public Boolean finishProcessingJob;
    @JsonProperty("IsFinishedGood")
    public Boolean isFinishedGood;
    @JsonProperty("Item")
    public String item;
    @JsonProperty("JobName")
    public String jobName;
    @JsonProperty("Location")
    public String location;
    @JsonProperty("Note")
    public String note;
    @JsonProperty("PackageDate")
    public String packageDate;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("ProductionBatchNumber")
    public String productionBatchNumber;
    @JsonProperty("Quantity")
    public Double quantity;
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
    @JsonProperty("WasteCountQuantity")
    public Double wasteCountQuantity;
    @JsonProperty("WasteCountUnitOfMeasureName")
    public String wasteCountUnitOfMeasureName;
    @JsonProperty("WasteVolumeQuantity")
    public Double wasteVolumeQuantity;
    @JsonProperty("WasteVolumeUnitOfMeasureName")
    public String wasteVolumeUnitOfMeasureName;
    @JsonProperty("WasteWeightQuantity")
    public Double wasteWeightQuantity;
    @JsonProperty("WasteWeightUnitOfMeasureName")
    public String wasteWeightUnitOfMeasureName;
}
