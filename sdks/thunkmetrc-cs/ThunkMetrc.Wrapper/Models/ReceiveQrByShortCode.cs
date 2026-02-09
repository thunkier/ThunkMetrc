using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class ReceiveQrByShortCode
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ChildTag")]
        public required string ChildTag { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Eaches")]
        public required List<string> Eaches { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabelSource")]
        public required string LabelSource { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("QrCount")]
        public required int QrCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Ranges")]
        public required List<List<int>> Ranges { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresVerification")]
        public required bool RequiresVerification { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SiblingTags")]
        public required List<string> SiblingTags { get; set; }
    }
}
