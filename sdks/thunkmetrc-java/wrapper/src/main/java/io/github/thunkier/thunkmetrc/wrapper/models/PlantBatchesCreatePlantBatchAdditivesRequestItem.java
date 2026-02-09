package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantBatchesCreatePlantBatchAdditivesRequestItem {
    @JsonProperty("ActiveIngredients")
    public List<PlantBatchesCreatePlantBatchAdditivesRequestItemActiveIngredients> activeIngredients;
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("AdditiveType")
    public String additiveType;
    @JsonProperty("ApplicationDevice")
    public String applicationDevice;
    @JsonProperty("EpaRegistrationNumber")
    public String epaRegistrationNumber;
    @JsonProperty("PlantBatchName")
    public String plantBatchName;
    @JsonProperty("ProductSupplier")
    public String productSupplier;
    @JsonProperty("ProductTradeName")
    public String productTradeName;
    @JsonProperty("TotalAmountApplied")
    public Integer totalAmountApplied;
    @JsonProperty("TotalAmountUnitOfMeasure")
    public String totalAmountUnitOfMeasure;
    public static class PlantBatchesCreatePlantBatchAdditivesRequestItemActiveIngredients {
    @JsonProperty("Name")
    public String name;
    @JsonProperty("Percentage")
    public Double percentage;
}
}
