using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class ProcessingJobUpdateJobTypesRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Attributes")]
        public List<object>? Attributes { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CategoryName")]
        public string? CategoryName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Description")]
        public string? Description { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProcessingSteps")]
        public string? ProcessingSteps { get; set; }
    }
}
