package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class Result {
    @JsonProperty("ExpirationDateTime")
    public String expirationDateTime;
    @JsonProperty("LabFacilityLicenseNumber")
    public String labFacilityLicenseNumber;
    @JsonProperty("LabFacilityName")
    public String labFacilityName;
    @JsonProperty("LabTestDetailRevokedDate")
    public String labTestDetailRevokedDate;
    @JsonProperty("LabTestResultDocumentFileId")
    public Integer labTestResultDocumentFileId;
    @JsonProperty("LabTestResultId")
    public Integer labTestResultId;
    @JsonProperty("OverallPassed")
    public Boolean overallPassed;
    @JsonProperty("PackageId")
    public Integer packageId;
    @JsonProperty("ProductCategoryName")
    public String productCategoryName;
    @JsonProperty("ProductName")
    public String productName;
    @JsonProperty("ResultReleaseDateTime")
    public String resultReleaseDateTime;
    @JsonProperty("ResultReleased")
    public Boolean resultReleased;
    @JsonProperty("RevokedDate")
    public String revokedDate;
    @JsonProperty("SourcePackageLabel")
    public String sourcePackageLabel;
    @JsonProperty("TestComment")
    public String testComment;
    @JsonProperty("TestInformationalOnly")
    public Boolean testInformationalOnly;
    @JsonProperty("TestPassed")
    public Boolean testPassed;
    @JsonProperty("TestPerformedDate")
    public String testPerformedDate;
    @JsonProperty("TestResultLevel")
    public Double testResultLevel;
    @JsonProperty("TestTypeName")
    public String testTypeName;
}
