package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class HarvestWaste {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("UnitOfWeightName")
    public String unitOfWeightName;
    @JsonProperty("WasteTypeName")
    public String wasteTypeName;
    @JsonProperty("WasteWeight")
    public Double wasteWeight;
}
