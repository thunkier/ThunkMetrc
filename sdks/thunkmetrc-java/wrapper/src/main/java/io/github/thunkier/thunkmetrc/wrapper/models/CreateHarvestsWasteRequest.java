package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreateHarvestsWasteRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("UnitOfWeight")
    public String unitOfWeight;
    @JsonProperty("WasteType")
    public String wasteType;
    @JsonProperty("WasteWeight")
    public Double wasteWeight;
}
