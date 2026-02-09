package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class AdditivesTemplate {
    @JsonProperty("ActiveIngredients")
    public List<AdditivesTemplateActiveIngredientsItem> activeIngredients;
    @JsonProperty("AdditiveType")
    public String additiveType;
    @JsonProperty("AdditiveTypeName")
    public String additiveTypeName;
    @JsonProperty("ApplicationDevice")
    public String applicationDevice;
    @JsonProperty("EpaRegistrationNumber")
    public String epaRegistrationNumber;
    @JsonProperty("FacilityId")
    public Integer facilityId;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IsActive")
    public Boolean isActive;
    @JsonProperty("LastModified")
    public String lastModified;
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
    public static class AdditivesTemplateActiveIngredientsItem {
    @JsonProperty("Name")
    public String name;
    @JsonProperty("Percentage")
    public Double percentage;
}
}
