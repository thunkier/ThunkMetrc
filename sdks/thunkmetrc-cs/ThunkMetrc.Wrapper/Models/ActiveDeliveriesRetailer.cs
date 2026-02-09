using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class ActiveDeliveriesRetailer
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("AcceptedDateTime")]
        public string? AcceptedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDepartureDateTime")]
        public string? ActualDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AllowFullEdit")]
        public required bool AllowFullEdit { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CompletedDateTime")]
        public string? CompletedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DateTime")]
        public required string DateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Destinations")]
        public required List<object> Destinations { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Direction")]
        public required string Direction { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverEmployeeId")]
        public required string DriverEmployeeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
        public required string DriverName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriversLicenseNumber")]
        public required string DriversLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
        public required string EstimatedDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityLicenseNumber")]
        public required string FacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityName")]
        public required string FacilityName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Leg")]
        public required int Leg { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
        public required List<object> Packages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecordedByUserName")]
        public string? RecordedByUserName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecordedDateTime")]
        public required string RecordedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RestockDateTime")]
        public string? RestockDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryNumber")]
        public required string RetailerDeliveryNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryState")]
        public required string RetailerDeliveryState { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalPackages")]
        public required int TotalPackages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalPrice")]
        public required int TotalPrice { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalPriceSold")]
        public required int TotalPriceSold { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleInfo")]
        public required string VehicleInfo { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
        public required string VehicleLicensePlateNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
        public required string VehicleMake { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
        public required string VehicleModel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VoidedDate")]
        public string? VoidedDate { get; set; }
    }
}
