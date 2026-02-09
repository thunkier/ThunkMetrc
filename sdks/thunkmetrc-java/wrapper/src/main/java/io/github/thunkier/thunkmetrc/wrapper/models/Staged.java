package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class Staged {
    @JsonProperty("CommissionedDateTime")
    public String commissionedDateTime;
    @JsonProperty("DetachedDateTime")
    public String detachedDateTime;
    @JsonProperty("FacilityId")
    public Integer facilityId;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IsArchived")
    public Boolean isArchived;
    @JsonProperty("IsStaged")
    public Boolean isStaged;
    @JsonProperty("IsUsed")
    public Boolean isUsed;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("MaxGroupSize")
    public Integer maxGroupSize;
    @JsonProperty("ProductLabel")
    public String productLabel;
    @JsonProperty("QrCount")
    public Integer qrCount;
    @JsonProperty("StatusName")
    public String statusName;
    @JsonProperty("TagInventoryTypeName")
    public String tagInventoryTypeName;
    @JsonProperty("TagTypeId")
    public Integer tagTypeId;
    @JsonProperty("TagTypeName")
    public String tagTypeName;
    @JsonProperty("UsedDateTime")
    public String usedDateTime;
}
