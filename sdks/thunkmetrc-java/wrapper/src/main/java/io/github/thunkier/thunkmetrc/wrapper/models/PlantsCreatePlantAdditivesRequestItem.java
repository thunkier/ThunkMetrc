package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsCreatePlantAdditivesRequestItem {
    @JsonProperty("ActiveIngredients")
    public List<PlantsCreatePlantAdditivesRequestItemActiveIngredients> activeIngredients;
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("AdditiveType")
    public String additiveType;
    @JsonProperty("ApplicationDevice")
    public String applicationDevice;
    @JsonProperty("EpaRegistrationNumber")
    public String epaRegistrationNumber;
    @JsonProperty("PlantLabels")
    public List<String> plantLabels;
    @JsonProperty("ProductSupplier")
    public String productSupplier;
    @JsonProperty("ProductTradeName")
    public String productTradeName;
    @JsonProperty("TotalAmountApplied")
    public Integer totalAmountApplied;
    @JsonProperty("TotalAmountUnitOfMeasure")
    public String totalAmountUnitOfMeasure;
    public static class PlantsCreatePlantAdditivesRequestItemActiveIngredients {
    @JsonProperty("Name")
    public String name;
    @JsonProperty("Percentage")
    public Double percentage;
}
}
