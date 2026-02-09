package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class AvailableTagPlant {
    @JsonProperty("FacilityId")
    public Integer facilityId;
    @JsonProperty("GroupTagTypeId")
    public String groupTagTypeId;
    @JsonProperty("GroupTagTypeName")
    public String groupTagTypeName;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("MaxGroupSize")
    public Integer maxGroupSize;
    @JsonProperty("TagInventoryTypeName")
    public String tagInventoryTypeName;
    @JsonProperty("TagTypeId")
    public String tagTypeId;
    @JsonProperty("TagTypeName")
    public String tagTypeName;
}
