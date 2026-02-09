package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsDeletePlantsRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("Count")
    public Integer count;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("ReasonNote")
    public String reasonNote;
    @JsonProperty("WasteMaterialMixed")
    public String wasteMaterialMixed;
    @JsonProperty("WasteMethodName")
    public String wasteMethodName;
    @JsonProperty("WasteReasonName")
    public String wasteReasonName;
    @JsonProperty("WasteUnitOfMeasureName")
    public String wasteUnitOfMeasureName;
    @JsonProperty("WasteWeight")
    public Double wasteWeight;
}
