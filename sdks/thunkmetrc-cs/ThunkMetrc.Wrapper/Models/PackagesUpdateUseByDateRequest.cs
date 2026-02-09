using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PackagesUpdateUseByDateRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ExpirationDate")]
        public string? ExpirationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public string? Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SellByDate")]
        public string? SellByDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UseByDate")]
        public string? UseByDate { get; set; }
    }
}
