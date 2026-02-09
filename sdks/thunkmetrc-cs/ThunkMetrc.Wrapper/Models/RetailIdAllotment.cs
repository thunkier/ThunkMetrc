using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class RetailIdAllotment
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Allotment")]
        public required int Allotment { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityId")]
        public required int FacilityId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityLicenseNumber")]
        public required string FacilityLicenseNumber { get; set; }
    }
}
