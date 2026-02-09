using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class CreateStrainsRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("CbdLevel")]
        public double? CbdLevel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IndicaPercentage")]
        public double? IndicaPercentage { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SativaPercentage")]
        public double? SativaPercentage { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TestingStatus")]
        public string? TestingStatus { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ThcLevel")]
        public double? ThcLevel { get; set; }
    }
}
