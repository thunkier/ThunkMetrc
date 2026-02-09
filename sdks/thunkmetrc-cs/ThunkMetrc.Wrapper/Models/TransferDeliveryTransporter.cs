using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransferDeliveryTransporter
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("TransporterDirection")]
        public required string TransporterDirection { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TransporterFacilityLicenseNumber")]
        public required string TransporterFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TransporterFacilityName")]
        public required string TransporterFacilityName { get; set; }
    }
}
