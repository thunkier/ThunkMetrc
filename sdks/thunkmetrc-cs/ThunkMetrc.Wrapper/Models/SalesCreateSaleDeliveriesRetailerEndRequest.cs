using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class SalesCreateSaleDeliveriesRetailerEndRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualArrivalDateTime")]
        public string? ActualArrivalDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
        public List<object>? Packages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryId")]
        public int? RetailerDeliveryId { get; set; }
    }
}
