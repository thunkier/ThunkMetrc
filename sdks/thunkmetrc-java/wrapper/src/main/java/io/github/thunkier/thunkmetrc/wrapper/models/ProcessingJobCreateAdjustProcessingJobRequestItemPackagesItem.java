package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class ProcessingJobCreateAdjustProcessingJobRequestItemPackagesItem {
    @JsonProperty("Label")
    public String label;
    @JsonProperty("Quantity")
    public Double quantity;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
