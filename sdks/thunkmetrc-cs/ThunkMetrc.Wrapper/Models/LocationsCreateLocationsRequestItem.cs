using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class LocationsCreateLocationsRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationTypeName")]
        public string? LocationTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
    }
}
