package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class Additive {
    @JsonProperty("AdditiveTypeName")
    public String additiveTypeName;
    @JsonProperty("AmountUnitOfMeasure")
    public String amountUnitOfMeasure;
    @JsonProperty("ApplicationDevice")
    public String applicationDevice;
    @JsonProperty("EpaRegistrationNumber")
    public String epaRegistrationNumber;
    @JsonProperty("Note")
    public String note;
    @JsonProperty("PlantBatchId")
    public Integer plantBatchId;
    @JsonProperty("PlantBatchName")
    public String plantBatchName;
    @JsonProperty("PlantCount")
    public Integer plantCount;
    @JsonProperty("ProductSupplier")
    public String productSupplier;
    @JsonProperty("ProductTradeName")
    public String productTradeName;
    @JsonProperty("Rate")
    public String rate;
    @JsonProperty("RestrictiveEntryIntervalQuantityDescription")
    public String restrictiveEntryIntervalQuantityDescription;
    @JsonProperty("RestrictiveEntryIntervalTimeDescription")
    public String restrictiveEntryIntervalTimeDescription;
    @JsonProperty("TotalAmountApplied")
    public Integer totalAmountApplied;
    @JsonProperty("Volume")
    public Double volume;
}
