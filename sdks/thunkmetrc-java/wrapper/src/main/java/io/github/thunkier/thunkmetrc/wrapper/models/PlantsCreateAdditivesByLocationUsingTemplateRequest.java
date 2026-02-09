package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsCreateAdditivesByLocationUsingTemplateRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("AdditivesTemplateName")
    public String additivesTemplateName;
    @JsonProperty("LocationName")
    public String locationName;
    @JsonProperty("Rate")
    public String rate;
    @JsonProperty("SublocationName")
    public String sublocationName;
    @JsonProperty("TotalAmountApplied")
    public Integer totalAmountApplied;
    @JsonProperty("TotalAmountUnitOfMeasure")
    public String totalAmountUnitOfMeasure;
    @JsonProperty("Volume")
    public String volume;
}
