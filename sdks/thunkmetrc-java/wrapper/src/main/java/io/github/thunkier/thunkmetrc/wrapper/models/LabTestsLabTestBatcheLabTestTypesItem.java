package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class LabTestsLabTestBatcheLabTestTypesItem {
    @JsonProperty("AlwaysPasses")
    public Boolean alwaysPasses;
    @JsonProperty("DependencyMode")
    public Integer dependencyMode;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("InformationalOnly")
    public Boolean informationalOnly;
    @JsonProperty("LabTestResultExpirationDays")
    public Integer labTestResultExpirationDays;
    @JsonProperty("LabTestResultMaximum")
    public String labTestResultMaximum;
    @JsonProperty("LabTestResultMinimum")
    public String labTestResultMinimum;
    @JsonProperty("LabTestResultMode")
    public Integer labTestResultMode;
    @JsonProperty("MaxAllowedFailureCount")
    public Integer maxAllowedFailureCount;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("RequiresTestResult")
    public Boolean requiresTestResult;
    @JsonProperty("ResearchAndDevelopment")
    public Boolean researchAndDevelopment;
}
