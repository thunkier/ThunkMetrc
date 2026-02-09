package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class ProcessingJobCreateAdjustProcessingJobRequest {
    @JsonProperty("AdjustmentDate")
    public String adjustmentDate;
    @JsonProperty("AdjustmentNote")
    public String adjustmentNote;
    @JsonProperty("AdjustmentReason")
    public String adjustmentReason;
    @JsonProperty("CountUnitOfMeasureName")
    public String countUnitOfMeasureName;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Packages")
    public List<Object> packages;
    @JsonProperty("VolumeUnitOfMeasureName")
    public String volumeUnitOfMeasureName;
    @JsonProperty("WeightUnitOfMeasureName")
    public String weightUnitOfMeasureName;
    public static class ProcessingJobCreateAdjustProcessingJobRequestPackagesItem {
    @JsonProperty("Label")
    public String label;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
}
