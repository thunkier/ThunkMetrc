package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsCreatePlantsAdditivesUsingTemplateRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("AdditivesTemplateName")
    public String additivesTemplateName;
    @JsonProperty("PlantLabels")
    public List<Object> plantLabels;
    @JsonProperty("Rate")
    public String rate;
    @JsonProperty("TotalAmountApplied")
    public Integer totalAmountApplied;
    @JsonProperty("TotalAmountUnitOfMeasure")
    public String totalAmountUnitOfMeasure;
    @JsonProperty("Volume")
    public String volume;
}
