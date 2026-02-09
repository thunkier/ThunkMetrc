package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class AdditivesTemplatesUpdateAdditivesTemplatesRequestItem {
    @JsonProperty("ActiveIngredients")
    public List<Object> activeIngredients;
    @JsonProperty("AdditiveType")
    public String additiveType;
    @JsonProperty("ApplicationDevice")
    public String applicationDevice;
    @JsonProperty("EpaRegistrationNumber")
    public String epaRegistrationNumber;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("Note")
    public String note;
    @JsonProperty("ProductSupplier")
    public String productSupplier;
    @JsonProperty("ProductTradeName")
    public String productTradeName;
    @JsonProperty("RestrictiveEntryIntervalQuantityDescription")
    public String restrictiveEntryIntervalQuantityDescription;
    @JsonProperty("RestrictiveEntryIntervalTimeDescription")
    public String restrictiveEntryIntervalTimeDescription;
    public static class AdditivesTemplatesUpdateAdditivesTemplatesRequestItemActiveIngredientsItem {
    @JsonProperty("Name")
    public String name;
    @JsonProperty("Percentage")
    public Double percentage;
}
}
