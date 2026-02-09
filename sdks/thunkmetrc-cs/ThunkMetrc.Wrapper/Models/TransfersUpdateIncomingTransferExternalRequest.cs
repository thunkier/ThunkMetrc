using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransfersUpdateIncomingTransferExternalRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Destinations")]
        public List<object>? Destinations { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverLicenseNumber")]
        public string? DriverLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
        public string? DriverName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverOccupationalLicenseNumber")]
        public string? DriverOccupationalLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PhoneNumberForQuestions")]
        public string? PhoneNumberForQuestions { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperAddress1")]
        public string? ShipperAddress1 { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperAddress2")]
        public string? ShipperAddress2 { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperAddressCity")]
        public string? ShipperAddressCity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperAddressPostalCode")]
        public string? ShipperAddressPostalCode { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperAddressState")]
        public string? ShipperAddressState { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperLicenseNumber")]
        public string? ShipperLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperMainPhoneNumber")]
        public string? ShipperMainPhoneNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperName")]
        public string? ShipperName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TransferId")]
        public int? TransferId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TransporterFacilityLicenseNumber")]
        public string? TransporterFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
        public string? VehicleLicensePlateNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
        public string? VehicleMake { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
        public string? VehicleModel { get; set; }
    }
}
