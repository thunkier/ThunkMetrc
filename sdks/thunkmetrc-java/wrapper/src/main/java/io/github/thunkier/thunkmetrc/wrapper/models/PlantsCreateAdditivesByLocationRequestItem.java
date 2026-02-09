package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsCreateAdditivesByLocationRequestItem {
    @JsonProperty("ActiveIngredients")
    public List<Object> activeIngredients;
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("AdditiveType")
    public String additiveType;
    @JsonProperty("ApplicationDevice")
    public String applicationDevice;
    @JsonProperty("EpaRegistrationNumber")
    public String epaRegistrationNumber;
    @JsonProperty("LocationName")
    public String locationName;
    @JsonProperty("ProductSupplier")
    public String productSupplier;
    @JsonProperty("ProductTradeName")
    public String productTradeName;
    @JsonProperty("SublocationName")
    public String sublocationName;
    @JsonProperty("TotalAmountApplied")
    public Integer totalAmountApplied;
    @JsonProperty("TotalAmountUnitOfMeasure")
    public String totalAmountUnitOfMeasure;
    public static class PlantsCreateAdditivesByLocationRequestItemActiveIngredientsItem {
    @JsonProperty("Name")
    public String name;
    @JsonProperty("Percentage")
    public Integer percentage;
}
}
