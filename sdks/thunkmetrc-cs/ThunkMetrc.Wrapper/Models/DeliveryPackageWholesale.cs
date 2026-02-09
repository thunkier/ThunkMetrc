using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class DeliveryPackageWholesale
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageId")]
        public required int PackageId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
        public required string PackageLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceiverWholesalePrice")]
        public string? ReceiverWholesalePrice { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperWholesalePrice")]
        public string? ShipperWholesalePrice { get; set; }
    }
}
