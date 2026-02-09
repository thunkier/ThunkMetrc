package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantBatchesCreateAdjustPlantBatchesRequest {
    @JsonProperty("AdjustmentDate")
    public String adjustmentDate;
    @JsonProperty("AdjustmentReason")
    public String adjustmentReason;
    @JsonProperty("PlantBatchName")
    public String plantBatchName;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("ReasonNote")
    public String reasonNote;
}
