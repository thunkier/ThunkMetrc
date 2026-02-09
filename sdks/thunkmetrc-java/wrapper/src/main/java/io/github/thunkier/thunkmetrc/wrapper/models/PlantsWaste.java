package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PlantsWaste {
    @JsonProperty("GrowthPhase")
    public Integer growthPhase;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("LocationId")
    public Integer locationId;
    @JsonProperty("LocationName")
    public String locationName;
    @JsonProperty("PlantId")
    public Integer plantId;
    @JsonProperty("PlantWasteId")
    public Integer plantWasteId;
    @JsonProperty("StateName")
    public String stateName;
    @JsonProperty("SublocationId")
    public Integer sublocationId;
    @JsonProperty("SublocationName")
    public String sublocationName;
    @JsonProperty("WasteAmount")
    public Integer wasteAmount;
    @JsonProperty("WasteUnitOfMeasureAbbreviation")
    public String wasteUnitOfMeasureAbbreviation;
}
