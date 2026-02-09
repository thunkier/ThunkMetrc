package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class DeletePlantBatchesRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("Count")
    public Integer count;
    @JsonProperty("PlantBatch")
    public String plantBatch;
    @JsonProperty("ReasonNote")
    public String reasonNote;
    @JsonProperty("WasteMaterialMixed")
    public String wasteMaterialMixed;
    @JsonProperty("WasteMethodName")
    public String wasteMethodName;
    @JsonProperty("WasteReasonName")
    public String wasteReasonName;
    @JsonProperty("WasteUnitOfMeasure")
    public String wasteUnitOfMeasure;
    @JsonProperty("WasteWeight")
    public Double wasteWeight;
}
