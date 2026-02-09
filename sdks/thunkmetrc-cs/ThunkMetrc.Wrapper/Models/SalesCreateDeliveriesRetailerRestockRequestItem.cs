using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class SalesCreateDeliveriesRetailerRestockRequestItem
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
        public class SalesCreateDeliveriesRetailerRestockRequestItemPackagesItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
            public string? PackageLabel { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
            public double? Quantity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RemoveCurrentPackage")]
            public bool? RemoveCurrentPackage { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TotalPrice")]
            public double? TotalPrice { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasure")]
            public string? UnitOfMeasure { get; set; }
        }
    }
}
