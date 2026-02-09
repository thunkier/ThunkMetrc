package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class LabTestsCreateRecordRequestResultsItem {
    @JsonProperty("LabTestTypeName")
    public String labTestTypeName;
    @JsonProperty("Notes")
    public String notes;
    @JsonProperty("Passed")
    public Boolean passed;
    @JsonProperty("Quantity")
    public Double quantity;
}
