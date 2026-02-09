package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class StrainsUpdateStrainsRequest {
    @JsonProperty("CbdLevel")
    public Double cbdLevel;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IndicaPercentage")
    public Double indicaPercentage;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("SativaPercentage")
    public Double sativaPercentage;
    @JsonProperty("TestingStatus")
    public String testingStatus;
    @JsonProperty("ThcLevel")
    public Double thcLevel;
}
