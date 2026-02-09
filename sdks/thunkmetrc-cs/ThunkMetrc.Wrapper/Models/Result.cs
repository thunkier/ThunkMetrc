using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class Result
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ExpirationDateTime")]
        public string? ExpirationDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabFacilityLicenseNumber")]
        public required string LabFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabFacilityName")]
        public required string LabFacilityName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestDetailRevokedDate")]
        public string? LabTestDetailRevokedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultDocumentFileId")]
        public int? LabTestResultDocumentFileId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultId")]
        public required int LabTestResultId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("OverallPassed")]
        public required bool OverallPassed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageId")]
        public required int PackageId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductCategoryName")]
        public required string ProductCategoryName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductName")]
        public required string ProductName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ResultReleaseDateTime")]
        public required string ResultReleaseDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ResultReleased")]
        public required bool ResultReleased { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RevokedDate")]
        public string? RevokedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePackageLabel")]
        public required string SourcePackageLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TestComment")]
        public required string TestComment { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TestInformationalOnly")]
        public required bool TestInformationalOnly { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TestPassed")]
        public required bool TestPassed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TestPerformedDate")]
        public required string TestPerformedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TestResultLevel")]
        public required double TestResultLevel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TestTypeName")]
        public required string TestTypeName { get; set; }
    }
}
