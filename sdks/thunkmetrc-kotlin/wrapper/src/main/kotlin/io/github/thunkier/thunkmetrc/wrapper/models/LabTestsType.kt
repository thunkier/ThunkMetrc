package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class LabTestsType(
    @JsonProperty("AlwaysPasses")
    val alwaysPasses: Boolean? = null,
    @JsonProperty("DependencyMode")
    val dependencyMode: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("InformationalOnly")
    val informationalOnly: Boolean? = null,
    @JsonProperty("LabTestResultExpirationDays")
    val labTestResultExpirationDays: Int? = null,
    @JsonProperty("LabTestResultMaximum")
    val labTestResultMaximum: Int? = null,
    @JsonProperty("LabTestResultMinimum")
    val labTestResultMinimum: Int? = null,
    @JsonProperty("LabTestResultMode")
    val labTestResultMode: String? = null,
    @JsonProperty("MaxAllowedFailureCount")
    val maxAllowedFailureCount: Int? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("RequiresTestResult")
    val requiresTestResult: Boolean? = null,
    @JsonProperty("ResearchAndDevelopment")
    val researchAndDevelopment: Boolean? = null
)
