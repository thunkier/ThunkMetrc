using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PackagesUpdateItemRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Item")]
        public string? Item { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public string? Label { get; set; }
    }
}
