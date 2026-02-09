package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsUpdatePlantStrainRequestItem {
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("StrainId")
    public Integer strainId;
    @JsonProperty("StrainName")
    public String strainName;
}
