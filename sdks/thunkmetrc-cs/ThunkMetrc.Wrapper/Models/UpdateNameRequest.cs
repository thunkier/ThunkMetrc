using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class UpdateNameRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Group")]
        public string? Group { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("NewGroup")]
        public string? NewGroup { get; set; }
    }
}
