using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class LabTestsType
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("AlwaysPasses")]
        public required bool AlwaysPasses { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DependencyMode")]
        public required string DependencyMode { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("InformationalOnly")]
        public required bool InformationalOnly { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultExpirationDays")]
        public int? LabTestResultExpirationDays { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultMaximum")]
        public required int LabTestResultMaximum { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultMinimum")]
        public required int LabTestResultMinimum { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestResultMode")]
        public required string LabTestResultMode { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("MaxAllowedFailureCount")]
        public required int MaxAllowedFailureCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresTestResult")]
        public required bool RequiresTestResult { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ResearchAndDevelopment")]
        public required bool ResearchAndDevelopment { get; set; }
    }
}
