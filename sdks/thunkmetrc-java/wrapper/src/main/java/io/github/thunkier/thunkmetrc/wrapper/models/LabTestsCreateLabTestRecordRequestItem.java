package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class LabTestsCreateLabTestRecordRequestItem {
    @JsonProperty("DocumentFileBase64")
    public String documentFileBase64;
    @JsonProperty("DocumentFileName")
    public String documentFileName;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("ResultDate")
    public String resultDate;
    @JsonProperty("Results")
    public List<LabTestsCreateLabTestRecordRequestItemResults> results;
    public static class LabTestsCreateLabTestRecordRequestItemResults {
    @JsonProperty("LabTestTypeName")
    public String labTestTypeName;
    @JsonProperty("Notes")
    public String notes;
    @JsonProperty("Passed")
    public Boolean passed;
    @JsonProperty("Quantity")
    public Double quantity;
}
}
