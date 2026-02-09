package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsUpdatePlantSplitRequestItem {
    @JsonProperty("PlantCount")
    public Integer plantCount;
    @JsonProperty("SourcePlantLabel")
    public String sourcePlantLabel;
    @JsonProperty("SplitDate")
    public String splitDate;
    @JsonProperty("StrainLabel")
    public String strainLabel;
    @JsonProperty("TagLabel")
    public String tagLabel;
}
