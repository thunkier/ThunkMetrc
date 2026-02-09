package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackagesUpdateRemediateRequestItem {
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("RemediationDate")
    public String remediationDate;
    @JsonProperty("RemediationMethodName")
    public String remediationMethodName;
    @JsonProperty("RemediationSteps")
    public String remediationSteps;
}
