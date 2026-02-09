package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class UpdatePlantBatchesTagRequest {
    @JsonProperty("Group")
    public String group;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("NewTag")
    public String newTag;
    @JsonProperty("ReplaceDate")
    public String replaceDate;
    @JsonProperty("TagId")
    public Integer tagId;
}
