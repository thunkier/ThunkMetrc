using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class SalesCreateSaleDeliveriesRetailerRestockRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("DateTime")]
        public string? DateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Destinations")]
        public string? Destinations { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
        public string? EstimatedDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
        public List<object>? Packages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryId")]
        public int? RetailerDeliveryId { get; set; }
    }
}
