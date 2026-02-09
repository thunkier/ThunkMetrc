using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class ProcessingJobTypesCategory
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ForItems")]
        public required bool ForItems { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ForProcessingJobs")]
        public required bool ForProcessingJobs { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresProcessingJobAttributes")]
        public required bool RequiresProcessingJobAttributes { get; set; }
    }
}
