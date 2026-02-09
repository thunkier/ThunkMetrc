using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PatientCheckInsUpdatePatientCheckInsRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("CheckInDate")]
        public string? CheckInDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CheckInLocationId")]
        public int? CheckInLocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RegistrationExpiryDate")]
        public string? RegistrationExpiryDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RegistrationStartDate")]
        public string? RegistrationStartDate { get; set; }
    }
}
