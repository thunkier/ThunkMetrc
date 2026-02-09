using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class SalesUpdateSaleDeliveriesRetailerRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("DateTime")]
        public string? DateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Destinations")]
        public List<object>? Destinations { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverEmployeeId")]
        public string? DriverEmployeeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
        public string? DriverName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriversLicenseNumber")]
        public string? DriversLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
        public string? EstimatedDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
        public List<object>? Packages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PhoneNumberForQuestions")]
        public string? PhoneNumberForQuestions { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
        public string? VehicleLicensePlateNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
        public string? VehicleMake { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
        public string? VehicleModel { get; set; }
    }
}
