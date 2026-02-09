package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class Strain {
    @JsonProperty("CbdLevel")
    public String cbdLevel;
    @JsonProperty("Genetics")
    public String genetics;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IndicaPercentage")
    public Integer indicaPercentage;
    @JsonProperty("IsUsed")
    public Boolean isUsed;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("SativaPercentage")
    public Integer sativaPercentage;
    @JsonProperty("TestingStatus")
    public String testingStatus;
    @JsonProperty("ThcLevel")
    public String thcLevel;
}
