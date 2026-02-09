package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class Adjustment {
    @JsonProperty("AdjustedByUserName")
    public String adjustedByUserName;
    @JsonProperty("AdjustmentDate")
    public String adjustmentDate;
    @JsonProperty("AdjustmentNote")
    public String adjustmentNote;
    @JsonProperty("AdjustmentQuantity")
    public Double adjustmentQuantity;
    @JsonProperty("AdjustmentReasonName")
    public String adjustmentReasonName;
    @JsonProperty("AdjustmentUnitOfMeasureAbbreviation")
    public String adjustmentUnitOfMeasureAbbreviation;
    @JsonProperty("AdjustmentUnitOfMeasureName")
    public String adjustmentUnitOfMeasureName;
    @JsonProperty("ItemCategoryName")
    public String itemCategoryName;
    @JsonProperty("ItemName")
    public String itemName;
    @JsonProperty("PackageId")
    public Integer packageId;
    @JsonProperty("PackageLabTestResultExpirationDateTime")
    public String packageLabTestResultExpirationDateTime;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("PackagedDate")
    public String packagedDate;
}
