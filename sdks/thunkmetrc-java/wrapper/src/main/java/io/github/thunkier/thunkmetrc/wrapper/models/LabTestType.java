package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class LabTestType {
    @JsonProperty("AlwaysPasses")
    public Boolean alwaysPasses;
    @JsonProperty("DependencyMode")
    public String dependencyMode;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("InformationalOnly")
    public Boolean informationalOnly;
    @JsonProperty("LabTestResultExpirationDays")
    public Integer labTestResultExpirationDays;
    @JsonProperty("LabTestResultMaximum")
    public Integer labTestResultMaximum;
    @JsonProperty("LabTestResultMinimum")
    public Integer labTestResultMinimum;
    @JsonProperty("LabTestResultMode")
    public String labTestResultMode;
    @JsonProperty("MaxAllowedFailureCount")
    public Integer maxAllowedFailureCount;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("RequiresTestResult")
    public Boolean requiresTestResult;
    @JsonProperty("ResearchAndDevelopment")
    public Boolean researchAndDevelopment;
}
