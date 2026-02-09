package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreateProcessingJobPackagesRequest {
    @JsonProperty("ExpirationDate")
    public String expirationDate;
    @JsonProperty("FinishDate")
    public String finishDate;
    @JsonProperty("FinishNote")
    public String finishNote;
    @JsonProperty("FinishProcessingJob")
    public Boolean finishProcessingJob;
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
    public Integer quantity;
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
    public String wasteCountQuantity;
    @JsonProperty("WasteCountUnitOfMeasureName")
    public String wasteCountUnitOfMeasureName;
    @JsonProperty("WasteVolumeQuantity")
    public String wasteVolumeQuantity;
    @JsonProperty("WasteVolumeUnitOfMeasureName")
    public String wasteVolumeUnitOfMeasureName;
    @JsonProperty("WasteWeightQuantity")
    public String wasteWeightQuantity;
    @JsonProperty("WasteWeightUnitOfMeasureName")
    public String wasteWeightUnitOfMeasureName;
}
