using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class UnitOfMeasure
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Abbreviation")]
        public required string Abbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("QuantityType")]
        public required string QuantityType { get; set; }
    }
}
