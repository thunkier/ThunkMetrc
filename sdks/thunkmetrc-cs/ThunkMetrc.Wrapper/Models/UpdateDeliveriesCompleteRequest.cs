using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class UpdateDeliveriesCompleteRequest
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
        public class UpdateDeliveriesCompleteRequestReturnedPackagesItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
            public string? Label { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ReturnQuantityVerified")]
            public int? ReturnQuantityVerified { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ReturnReason")]
            public string? ReturnReason { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ReturnReasonNote")]
            public string? ReturnReasonNote { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ReturnUnitOfMeasure")]
            public string? ReturnUnitOfMeasure { get; set; }
        }
    }
}
