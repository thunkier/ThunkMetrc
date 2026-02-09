using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class ActiveJobType
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Attributes")]
        public required List<ProcessingJobActiveJobTypeAttributesItem> Attributes { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CategoryName")]
        public required string CategoryName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Description")]
        public required string Description { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ForItems")]
        public required bool ForItems { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ForProcessingJobs")]
        public required bool ForProcessingJobs { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProcessingSteps")]
        public required string ProcessingSteps { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresProcessingJobAttributes")]
        public required bool RequiresProcessingJobAttributes { get; set; }
        public class ProcessingJobActiveJobTypeAttributesItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
            public required int Id { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
            public required string LastModified { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
            public required string Name { get; set; }
        }
    }
}
