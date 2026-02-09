package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class UpdateAdjustPackagesRequest {
    @JsonProperty("AdjustmentDate")
    public String adjustmentDate;
    @JsonProperty("AdjustmentReason")
    public String adjustmentReason;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("ReasonNote")
    public String reasonNote;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
