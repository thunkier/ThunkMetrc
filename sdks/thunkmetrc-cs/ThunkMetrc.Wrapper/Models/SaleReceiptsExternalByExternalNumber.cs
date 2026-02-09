using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class SaleReceiptsExternalByExternalNumber
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ArchivedDate")]
        public string? ArchivedDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CaregiverLicenseNumber")]
        public string? CaregiverLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExternalReceiptNumber")]
        public required string ExternalReceiptNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IdentificationMethod")]
        public string? IdentificationMethod { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsFinal")]
        public required bool IsFinal { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientRegistrationLocationId")]
        public int? PatientRegistrationLocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceiptNumber")]
        public string? ReceiptNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecordedByUserName")]
        public string? RecordedByUserName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecordedDateTime")]
        public required string RecordedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SalesCustomerType")]
        public required string SalesCustomerType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SalesDateTime")]
        public required string SalesDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalPackages")]
        public required int TotalPackages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalPrice")]
        public required int TotalPrice { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Transactions")]
        public required List<object> Transactions { get; set; }
    }
}
