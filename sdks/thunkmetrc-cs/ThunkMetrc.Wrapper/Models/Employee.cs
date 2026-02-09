using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class Employee
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("FullName")]
        public required string FullName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsIndustryAdmin")]
        public required bool IsIndustryAdmin { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsManager")]
        public required bool IsManager { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsOwner")]
        public required bool IsOwner { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("License")]
        public required EmployeeLicense License { get; set; }
        public class EmployeeLicense
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("EmployeeLicenseNumber")]
            public required string EmployeeLicenseNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EndDate")]
            public required string EndDate { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LicenseType")]
            public required string LicenseType { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Number")]
            public required string Number { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("StartDate")]
            public required string StartDate { get; set; }
        }
    }
}
