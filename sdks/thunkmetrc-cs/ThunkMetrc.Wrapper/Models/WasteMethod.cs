using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class WasteMethod
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ForPlants")]
        public required bool ForPlants { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ForProductDestruction")]
        public required bool ForProductDestruction { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
    }
}
