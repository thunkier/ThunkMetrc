using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransporterDriver
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("DriversLicenseNumber")]
        public required string DriversLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EmployeeId")]
        public required string EmployeeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityId")]
        public required int FacilityId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
    }
}
