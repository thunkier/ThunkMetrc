using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TemplateDeliveryTransporterDetail
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDriverStartDateTime")]
        public string? ActualDriverStartDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverLayoverLeg")]
        public string? DriverLayoverLeg { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
        public required string DriverName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverOccupationalLicenseNumber")]
        public required string DriverOccupationalLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverVehicleLicenseNumber")]
        public required string DriverVehicleLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipmentDeliveryId")]
        public required int ShipmentDeliveryId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
        public required string VehicleLicensePlateNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
        public required string VehicleMake { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
        public required string VehicleModel { get; set; }
    }
}
