using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class HarvestWaste
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public required string ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfWeightName")]
        public required string UnitOfWeightName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteTypeName")]
        public required string WasteTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteWeight")]
        public required double WasteWeight { get; set; }
    }
}
