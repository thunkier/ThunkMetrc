using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantBatchesUpdatePlantBatchesLocationRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Location")]
        public string? Location { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("MoveDate")]
        public string? MoveDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Sublocation")]
        public string? Sublocation { get; set; }
    }
}
