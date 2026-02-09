using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class CreateDriversRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("DriversLicenseNumber")]
        public string? DriversLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EmployeeId")]
        public string? EmployeeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
    }
}
