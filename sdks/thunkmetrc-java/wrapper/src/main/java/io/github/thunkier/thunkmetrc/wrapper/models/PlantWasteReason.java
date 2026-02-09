package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantWasteReason {
    @JsonProperty("Name")
    public String name;
    @JsonProperty("RequiresImmatureWasteWeight")
    public Boolean requiresImmatureWasteWeight;
    @JsonProperty("RequiresMatureWasteWeight")
    public Boolean requiresMatureWasteWeight;
    @JsonProperty("RequiresNote")
    public Boolean requiresNote;
    @JsonProperty("RequiresWasteWeight")
    public Boolean requiresWasteWeight;
}
