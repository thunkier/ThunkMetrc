package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class ProcessingJobFinishProcessingJobRequest {
    @JsonProperty("FinishDate")
    public String finishDate;
    @JsonProperty("FinishNote")
    public String finishNote;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("TotalCountWaste")
    public String totalCountWaste;
    @JsonProperty("TotalVolumeWaste")
    public String totalVolumeWaste;
    @JsonProperty("TotalWeightWaste")
    public Integer totalWeightWaste;
    @JsonProperty("WasteCountUnitOfMeasureName")
    public String wasteCountUnitOfMeasureName;
    @JsonProperty("WasteVolumeUnitOfMeasureName")
    public String wasteVolumeUnitOfMeasureName;
    @JsonProperty("WasteWeightUnitOfMeasureName")
    public String wasteWeightUnitOfMeasureName;
}
