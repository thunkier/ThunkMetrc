package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class ProcessingJobCreateAdjustProcessingJobRequestItemPackages {
    @JsonProperty("Label")
    public String label;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
