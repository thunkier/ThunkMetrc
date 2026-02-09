using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class HarvestsCreateHarvestWasteRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfWeight")]
        public string? UnitOfWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteType")]
        public string? WasteType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteWeight")]
        public double? WasteWeight { get; set; }
    }
}
