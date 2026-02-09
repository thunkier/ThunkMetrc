package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsWastePackageProductLabel {
    @JsonProperty("IsActive")
    public Boolean isActive;
    @JsonProperty("IsChildFromParentWithLabel")
    public Boolean isChildFromParentWithLabel;
    @JsonProperty("LabelSource")
    public String labelSource;
    @JsonProperty("LastLabelGenerationDate")
    public String lastLabelGenerationDate;
    @JsonProperty("OriginalSourcePackageId")
    public Integer originalSourcePackageId;
    @JsonProperty("OriginalSourcePackageLabel")
    public String originalSourcePackageLabel;
    @JsonProperty("QrCodeRanges")
    public String qrCodeRanges;
    @JsonProperty("QrCount")
    public Integer qrCount;
}
