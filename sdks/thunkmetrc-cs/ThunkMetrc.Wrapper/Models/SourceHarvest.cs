using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class SourceHarvest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestStartDate")]
        public required string HarvestStartDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestedByFacilityLicenseNumber")]
        public required string HarvestedByFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("HarvestedByFacilityName")]
        public required string HarvestedByFacilityName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
    }
}
