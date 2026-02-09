using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PatientCheckIn
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("CheckInDate")]
        public required string CheckInDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CheckInLocationId")]
        public required int CheckInLocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CheckInLocationName")]
        public required string CheckInLocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public required string PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RegistrationExpiryDate")]
        public required string RegistrationExpiryDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RegistrationStartDate")]
        public required string RegistrationStartDate { get; set; }
    }
}
