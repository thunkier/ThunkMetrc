using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransfersCreateHubArriveRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipmentDeliveryId")]
        public int? ShipmentDeliveryId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TransporterDirection")]
        public string? TransporterDirection { get; set; }
    }
}
