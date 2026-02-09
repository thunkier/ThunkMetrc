using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransfersCreateIncomingExternalRequest
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
        [global::System.Text.Json.Serialization.JsonPropertyName("TransporterFacilityLicenseNumber")]
        public string? TransporterFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
        public string? VehicleLicensePlateNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
        public string? VehicleMake { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
        public string? VehicleModel { get; set; }
        public class TransfersCreateIncomingExternalRequestDestinationsItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedArrivalDateTime")]
            public string? EstimatedArrivalDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
            public string? EstimatedDepartureDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("GrossUnitOfWeightId")]
            public int? GrossUnitOfWeightId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("GrossWeight")]
            public double? GrossWeight { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("InvoiceNumber")]
            public string? InvoiceNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
            public List<object>? Packages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PlannedRoute")]
            public string? PlannedRoute { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientLicenseNumber")]
            public string? RecipientLicenseNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TransferTypeName")]
            public string? TransferTypeName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Transporters")]
            public List<object>? Transporters { get; set; }
            public class TransfersCreateIncomingExternalRequestDestinationsItemPackagesItem
            {
                [global::System.Text.Json.Serialization.JsonPropertyName("ExpirationDate")]
                public string? ExpirationDate { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("ExternalId")]
                public string? ExternalId { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("GrossUnitOfWeightName")]
                public string? GrossUnitOfWeightName { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("GrossWeight")]
                public double? GrossWeight { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("ItemName")]
                public string? ItemName { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("PackagedDate")]
                public string? PackagedDate { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
                public int? Quantity { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("SellByDate")]
                public string? SellByDate { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasureName")]
                public string? UnitOfMeasureName { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("UseByDate")]
                public string? UseByDate { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("WholesalePrice")]
                public string? WholesalePrice { get; set; }
            }
            public class TransfersCreateIncomingExternalRequestDestinationsItemTransportersItem
            {
                [global::System.Text.Json.Serialization.JsonPropertyName("DriverLayoverLeg")]
                public string? DriverLayoverLeg { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("DriverLicenseNumber")]
                public string? DriverLicenseNumber { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
                public string? DriverName { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("DriverOccupationalLicenseNumber")]
                public string? DriverOccupationalLicenseNumber { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedArrivalDateTime")]
                public string? EstimatedArrivalDateTime { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
                public string? EstimatedDepartureDateTime { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("IsLayover")]
                public bool? IsLayover { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("PhoneNumberForQuestions")]
                public string? PhoneNumberForQuestions { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("TransporterDetails")]
                public List<object>? TransporterDetails { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("TransporterFacilityLicenseNumber")]
                public string? TransporterFacilityLicenseNumber { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
                public string? VehicleLicensePlateNumber { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
                public string? VehicleMake { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
                public string? VehicleModel { get; set; }
                public class TransfersCreateIncomingExternalRequestDestinationsItemTransportersItemTransporterDetailsItem
                {
                    [global::System.Text.Json.Serialization.JsonPropertyName("DriverLayoverLeg")]
                    public string? DriverLayoverLeg { get; set; }
                    [global::System.Text.Json.Serialization.JsonPropertyName("DriverLicenseNumber")]
                    public string? DriverLicenseNumber { get; set; }
                    [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
                    public string? DriverName { get; set; }
                    [global::System.Text.Json.Serialization.JsonPropertyName("DriverOccupationalLicenseNumber")]
                    public string? DriverOccupationalLicenseNumber { get; set; }
                    [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
                    public string? VehicleLicensePlateNumber { get; set; }
                    [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
                    public string? VehicleMake { get; set; }
                    [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
                    public string? VehicleModel { get; set; }
                }
            }
        }
    }
}
