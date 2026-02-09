using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class JobTypesAttribute
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
    }
}
