using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class SalesUpdateSaleReceiptsRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("CaregiverLicenseNumber")]
        public string? CaregiverLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExternalReceiptNumber")]
        public string? ExternalReceiptNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IdentificationMethod")]
        public string? IdentificationMethod { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientRegistrationLocationId")]
        public int? PatientRegistrationLocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SalesCustomerType")]
        public string? SalesCustomerType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SalesDateTime")]
        public string? SalesDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Transactions")]
        public List<object>? Transactions { get; set; }
    }
}
