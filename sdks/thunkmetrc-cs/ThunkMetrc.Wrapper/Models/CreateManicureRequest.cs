using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class CreateManicureRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DryingLocation")]
        public string? DryingLocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DryingSublocation")]
        public string? DryingSublocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestName")]
        public string? HarvestName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Plant")]
        public string? Plant { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantCount")]
        public int? PlantCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfWeight")]
        public string? UnitOfWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Weight")]
        public double? Weight { get; set; }
    }
}
