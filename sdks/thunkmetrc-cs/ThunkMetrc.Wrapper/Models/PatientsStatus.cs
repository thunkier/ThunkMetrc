using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PatientsStatus
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Active")]
        public required bool Active { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ConcentrateOuncesAllowed")]
        public required int ConcentrateOuncesAllowed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ConcentrateOuncesAvailable")]
        public required int ConcentrateOuncesAvailable { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ConcentrateOuncesPurchased")]
        public required int ConcentrateOuncesPurchased { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FlowerOuncesAllowed")]
        public required int FlowerOuncesAllowed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FlowerOuncesAvailable")]
        public required int FlowerOuncesAvailable { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FlowerOuncesPurchased")]
        public required int FlowerOuncesPurchased { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IdentificationMethod")]
        public string? IdentificationMethod { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("InfusedOuncesAllowed")]
        public required int InfusedOuncesAllowed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("InfusedOuncesAvailable")]
        public required int InfusedOuncesAvailable { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("InfusedOuncesPurchased")]
        public required int InfusedOuncesPurchased { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseExpirationDate")]
        public required string PatientLicenseExpirationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public required string PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientRegistrationStartDate")]
        public required string PatientRegistrationStartDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PurchaseAmountDays")]
        public required int PurchaseAmountDays { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RegistrationNumber")]
        public required string RegistrationNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ThcOuncesAllowed")]
        public required int ThcOuncesAllowed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ThcOuncesAvailable")]
        public required int ThcOuncesAvailable { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ThcOuncesPurchased")]
        public required int ThcOuncesPurchased { get; set; }
    }
}
