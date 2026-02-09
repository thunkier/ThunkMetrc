package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class FinishHarvestsRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("Id")
    public Integer id;
}
