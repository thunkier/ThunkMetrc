package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreatePlantsWasteRequest {
    @JsonProperty("LocationName")
    public String locationName;
    @JsonProperty("MixedMaterial")
    public String mixedMaterial;
    @JsonProperty("Note")
    public String note;
    @JsonProperty("PlantLabels")
    public List<Object> plantLabels;
    @JsonProperty("ReasonName")
    public String reasonName;
    @JsonProperty("SublocationName")
    public String sublocationName;
    @JsonProperty("UnitOfMeasureName")
    public String unitOfMeasureName;
    @JsonProperty("WasteDate")
    public String wasteDate;
    @JsonProperty("WasteMethodName")
    public String wasteMethodName;
    @JsonProperty("WasteWeight")
    public Double wasteWeight;
}
