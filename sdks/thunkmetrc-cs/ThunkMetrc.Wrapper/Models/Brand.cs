using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class Brand
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Status")]
        public required string Status { get; set; }
    }
}
