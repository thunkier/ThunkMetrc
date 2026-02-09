using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantBatchesUpdatePlantBatchesStrainRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainId")]
        public int? StrainId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("StrainName")]
        public string? StrainName { get; set; }
    }
}
