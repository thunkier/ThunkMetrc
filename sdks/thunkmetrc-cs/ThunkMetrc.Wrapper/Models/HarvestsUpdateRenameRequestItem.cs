using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class HarvestsUpdateRenameRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("NewName")]
        public string? NewName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("OldName")]
        public string? OldName { get; set; }
    }
}
