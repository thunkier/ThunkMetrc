package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackagesPackageSummaryProductLabel {
    @JsonProperty("IsActive")
    public Boolean isActive;
    @JsonProperty("IsChildFromParentWithLabel")
    public Boolean isChildFromParentWithLabel;
    @JsonProperty("IsDiscontinued")
    public Boolean isDiscontinued;
    @JsonProperty("LastLabelGenerationDate")
    public String lastLabelGenerationDate;
    @JsonProperty("PackageId")
    public Integer packageId;
    @JsonProperty("QrCount")
    public Integer qrCount;
    @JsonProperty("TagId")
    public Integer tagId;
}
