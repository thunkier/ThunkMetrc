using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class LocationsLocation
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ForHarvests")]
        public required bool ForHarvests { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ForPackages")]
        public required bool ForPackages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ForPlantBatches")]
        public required bool ForPlantBatches { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ForPlants")]
        public required bool ForPlants { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationTypeId")]
        public required int LocationTypeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationTypeName")]
        public required string LocationTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
    }
}
