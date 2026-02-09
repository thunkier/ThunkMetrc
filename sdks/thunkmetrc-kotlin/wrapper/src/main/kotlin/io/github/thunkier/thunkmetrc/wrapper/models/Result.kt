package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Result(
    @JsonProperty("ExpirationDateTime")
    val expirationDateTime: String? = null,
    @JsonProperty("LabFacilityLicenseNumber")
    val labFacilityLicenseNumber: String? = null,
    @JsonProperty("LabFacilityName")
    val labFacilityName: String? = null,
    @JsonProperty("LabTestDetailRevokedDate")
    val labTestDetailRevokedDate: String? = null,
    @JsonProperty("LabTestResultDocumentFileId")
    val labTestResultDocumentFileId: Int? = null,
    @JsonProperty("LabTestResultId")
    val labTestResultId: Int? = null,
    @JsonProperty("OverallPassed")
    val overallPassed: Boolean? = null,
    @JsonProperty("PackageId")
    val packageId: Int? = null,
    @JsonProperty("ProductCategoryName")
    val productCategoryName: String? = null,
    @JsonProperty("ProductName")
    val productName: String? = null,
    @JsonProperty("ResultReleaseDateTime")
    val resultReleaseDateTime: String? = null,
    @JsonProperty("ResultReleased")
    val resultReleased: Boolean? = null,
    @JsonProperty("RevokedDate")
    val revokedDate: String? = null,
    @JsonProperty("SourcePackageLabel")
    val sourcePackageLabel: String? = null,
    @JsonProperty("TestComment")
    val testComment: String? = null,
    @JsonProperty("TestInformationalOnly")
    val testInformationalOnly: Boolean? = null,
    @JsonProperty("TestPassed")
    val testPassed: Boolean? = null,
    @JsonProperty("TestPerformedDate")
    val testPerformedDate: String? = null,
    @JsonProperty("TestResultLevel")
    val testResultLevel: Double? = null,
    @JsonProperty("TestTypeName")
    val testTypeName: String? = null
)
