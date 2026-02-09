using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransportersUpdateDriversRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("DriversLicenseNumber")]
        public string? DriversLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EmployeeId")]
        public string? EmployeeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public string? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
    }
}
