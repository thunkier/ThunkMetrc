package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantBatchesWaste {
    @JsonProperty("PlantBatchId")
    public Integer plantBatchId;
    @JsonProperty("PlantBatchName")
    public String plantBatchName;
    @JsonProperty("PlantCount")
    public Integer plantCount;
    @JsonProperty("PlantWasteNumber")
    public String plantWasteNumber;
    @JsonProperty("WasteDate")
    public String wasteDate;
    @JsonProperty("WasteMethodName")
    public String wasteMethodName;
    @JsonProperty("WasteReasonName")
    public String wasteReasonName;
    @JsonProperty("WasteUnitOfMeasureName")
    public String wasteUnitOfMeasureName;
    @JsonProperty("WasteWeight")
    public Double wasteWeight;
}
