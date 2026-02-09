using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class SalesUpdateSaleDeliveriesCompleteRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("AcceptedPackages")]
        public List<object>? AcceptedPackages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualArrivalDateTime")]
        public string? ActualArrivalDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PaymentType")]
        public string? PaymentType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReturnedPackages")]
        public List<object>? ReturnedPackages { get; set; }
    }
}
