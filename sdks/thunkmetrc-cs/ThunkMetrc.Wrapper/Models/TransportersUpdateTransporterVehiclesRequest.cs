using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransportersUpdateTransporterVehiclesRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public string? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LicensePlateNumber")]
        public string? LicensePlateNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Make")]
        public string? Make { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Model")]
        public string? Model { get; set; }
    }
}
