using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransfersCreateOutgoingTemplatesRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Destinations")]
        public List<object>? Destinations { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverLicenseNumber")]
        public string? DriverLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
        public string? DriverName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverOccupationalLicenseNumber")]
        public string? DriverOccupationalLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PhoneNumberForQuestions")]
        public string? PhoneNumberForQuestions { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TransporterFacilityLicenseNumber")]
        public string? TransporterFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
        public string? VehicleLicensePlateNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
        public string? VehicleMake { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
        public string? VehicleModel { get; set; }
        public class TransfersCreateOutgoingTemplatesRequestDestinationsItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedArrivalDateTime")]
            public string? EstimatedArrivalDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
            public string? EstimatedDepartureDateTime { get; set; }
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
            public class TransfersCreateOutgoingTemplatesRequestDestinationsItemPackagesItem
            {
                [global::System.Text.Json.Serialization.JsonPropertyName("GrossUnitOfWeightName")]
                public string? GrossUnitOfWeightName { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("GrossWeight")]
                public double? GrossWeight { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
                public string? PackageLabel { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("WholesalePrice")]
                public string? WholesalePrice { get; set; }
            }
            public class TransfersCreateOutgoingTemplatesRequestDestinationsItemTransportersItem
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
                public class TransfersCreateOutgoingTemplatesRequestDestinationsItemTransportersItemTransporterDetailsItem
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
