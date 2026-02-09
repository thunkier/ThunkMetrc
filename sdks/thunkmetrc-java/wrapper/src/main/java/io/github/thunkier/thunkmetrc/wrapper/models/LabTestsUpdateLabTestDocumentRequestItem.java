package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class LabTestsUpdateLabTestDocumentRequestItem {
    @JsonProperty("DocumentFileBase64")
    public String documentFileBase64;
    @JsonProperty("DocumentFileName")
    public String documentFileName;
    @JsonProperty("LabTestResultId")
    public Integer labTestResultId;
}
