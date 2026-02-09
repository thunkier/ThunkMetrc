using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransporterVehicle
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityId")]
        public required int FacilityId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LicensePlateNumber")]
        public required string LicensePlateNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Make")]
        public required string Make { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Model")]
        public required string Model { get; set; }
    }
}
