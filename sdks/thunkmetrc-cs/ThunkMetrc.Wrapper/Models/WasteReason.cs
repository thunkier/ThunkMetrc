using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class WasteReason
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresImmatureWasteWeight")]
        public required bool RequiresImmatureWasteWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresMatureWasteWeight")]
        public required bool RequiresMatureWasteWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresNote")]
        public required bool RequiresNote { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresWasteWeight")]
        public required bool RequiresWasteWeight { get; set; }
    }
}
