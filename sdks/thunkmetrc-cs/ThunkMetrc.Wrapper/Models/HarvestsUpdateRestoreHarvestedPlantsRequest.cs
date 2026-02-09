using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class HarvestsUpdateRestoreHarvestedPlantsRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestId")]
        public int? HarvestId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantTags")]
        public List<object>? PlantTags { get; set; }
    }
}
