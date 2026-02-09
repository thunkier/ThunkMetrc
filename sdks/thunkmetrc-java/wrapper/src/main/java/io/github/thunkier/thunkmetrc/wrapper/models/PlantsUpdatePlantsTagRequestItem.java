package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsUpdatePlantsTagRequestItem {
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("NewTag")
    public String newTag;
    @JsonProperty("ReplaceDate")
    public String replaceDate;
    @JsonProperty("TagId")
    public Integer tagId;
}
