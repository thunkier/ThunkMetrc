using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class LocationsUpdateLocationsRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationTypeName")]
        public string? LocationTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
    }
}
