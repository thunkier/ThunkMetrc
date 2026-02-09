package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreatePlantBatchesAdditivesUsingTemplateRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("AdditivesTemplateName")
    public String additivesTemplateName;
    @JsonProperty("PlantBatchName")
    public String plantBatchName;
    @JsonProperty("Rate")
    public String rate;
    @JsonProperty("TotalAmountApplied")
    public Integer totalAmountApplied;
    @JsonProperty("TotalAmountUnitOfMeasure")
    public String totalAmountUnitOfMeasure;
    @JsonProperty("Volume")
    public String volume;
}
