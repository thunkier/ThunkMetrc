package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class UpdateAdjustPlantsRequest {
    @JsonProperty("AdjustCount")
    public Integer adjustCount;
    @JsonProperty("AdjustReason")
    public String adjustReason;
    @JsonProperty("AdjustmentDate")
    public String adjustmentDate;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("ReasonNote")
    public String reasonNote;
}
