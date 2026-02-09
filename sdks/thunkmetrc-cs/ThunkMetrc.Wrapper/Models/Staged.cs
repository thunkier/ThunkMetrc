using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class Staged
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("CommissionedDateTime")]
        public string? CommissionedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DetachedDateTime")]
        public string? DetachedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityId")]
        public required int FacilityId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsArchived")]
        public required bool IsArchived { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsStaged")]
        public required bool IsStaged { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsUsed")]
        public required bool IsUsed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public required string Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("MaxGroupSize")]
        public required int MaxGroupSize { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductLabel")]
        public string? ProductLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("QrCount")]
        public int? QrCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StatusName")]
        public string? StatusName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TagInventoryTypeName")]
        public required string TagInventoryTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TagTypeId")]
        public int? TagTypeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TagTypeName")]
        public string? TagTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UsedDateTime")]
        public string? UsedDateTime { get; set; }
    }
}
