using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class CaregiversStatus
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Active")]
        public required bool Active { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CaregiverLicenseNumber")]
        public required string CaregiverLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Patients")]
        public required List<string> Patients { get; set; }
    }
}
