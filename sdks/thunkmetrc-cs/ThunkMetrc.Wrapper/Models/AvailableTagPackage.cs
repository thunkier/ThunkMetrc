using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class AvailableTagPackage
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityId")]
        public required int FacilityId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GroupTagTypeId")]
        public required string GroupTagTypeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GroupTagTypeName")]
        public required string GroupTagTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public required string Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("MaxGroupSize")]
        public required int MaxGroupSize { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TagInventoryTypeName")]
        public required string TagInventoryTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TagTypeId")]
        public required string TagTypeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TagTypeName")]
        public required string TagTypeName { get; set; }
    }
}
