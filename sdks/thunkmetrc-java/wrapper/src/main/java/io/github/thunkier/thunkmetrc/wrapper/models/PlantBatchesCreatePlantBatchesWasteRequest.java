package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantBatchesCreatePlantBatchesWasteRequest {
    @JsonProperty("MixedMaterial")
    public String mixedMaterial;
    @JsonProperty("Note")
    public String note;
    @JsonProperty("PlantBatchName")
    public String plantBatchName;
    @JsonProperty("ReasonName")
    public String reasonName;
    @JsonProperty("UnitOfMeasureName")
    public String unitOfMeasureName;
    @JsonProperty("WasteDate")
    public String wasteDate;
    @JsonProperty("WasteMethodName")
    public String wasteMethodName;
    @JsonProperty("WasteWeight")
    public Double wasteWeight;
}
