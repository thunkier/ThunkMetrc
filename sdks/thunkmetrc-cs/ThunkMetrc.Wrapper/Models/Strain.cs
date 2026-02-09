using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class Strain
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("CbdLevel")]
        public string? CbdLevel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Genetics")]
        public required string Genetics { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IndicaPercentage")]
        public required int IndicaPercentage { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsUsed")]
        public required bool IsUsed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SativaPercentage")]
        public required int SativaPercentage { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TestingStatus")]
        public required string TestingStatus { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ThcLevel")]
        public string? ThcLevel { get; set; }
    }
}
