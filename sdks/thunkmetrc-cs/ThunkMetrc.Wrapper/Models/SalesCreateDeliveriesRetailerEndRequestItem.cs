using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class SalesCreateDeliveriesRetailerEndRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualArrivalDateTime")]
        public string? ActualArrivalDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
        public List<object>? Packages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryId")]
        public int? RetailerDeliveryId { get; set; }
        public class SalesCreateDeliveriesRetailerEndRequestItemPackagesItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("EndQuantity")]
            public double? EndQuantity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EndUnitOfMeasure")]
            public string? EndUnitOfMeasure { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
            public string? Label { get; set; }
        }
    }
}
