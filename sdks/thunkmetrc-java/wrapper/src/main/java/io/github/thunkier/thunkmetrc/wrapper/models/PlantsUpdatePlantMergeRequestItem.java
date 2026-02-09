package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsUpdatePlantMergeRequestItem {
    @JsonProperty("MergeDate")
    public String mergeDate;
    @JsonProperty("SourcePlantGroupLabel")
    public String sourcePlantGroupLabel;
    @JsonProperty("TargetPlantGroupLabel")
    public String targetPlantGroupLabel;
}
