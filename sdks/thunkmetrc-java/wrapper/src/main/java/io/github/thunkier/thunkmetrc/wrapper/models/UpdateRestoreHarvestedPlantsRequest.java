package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class UpdateRestoreHarvestedPlantsRequest {
    @JsonProperty("HarvestId")
    public Integer harvestId;
    @JsonProperty("PlantTags")
    public List<Object> plantTags;
}
