package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class StrainsCreateStrainsRequestItem {
    @JsonProperty("CbdLevel")
    public Double cbdLevel;
    @JsonProperty("IndicaPercentage")
    public Integer indicaPercentage;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("SativaPercentage")
    public Integer sativaPercentage;
    @JsonProperty("TestingStatus")
    public String testingStatus;
    @JsonProperty("ThcLevel")
    public Double thcLevel;
}
