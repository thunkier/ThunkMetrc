using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class RetailIdPackagesInfo
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
        public required List<RetailIdRetailIdPackagesInfoPackagesItem> Packages { get; set; }
        public class RetailIdRetailIdPackagesInfoPackagesItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedBalance")]
            public required int EstimatedBalance { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("HasQrs")]
            public required bool HasQrs { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IssuanceId")]
            public required string IssuanceId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Issuances")]
            public required List<RetailIdRetailIdPackagesInfoPackagesItemIssuancesItem> Issuances { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("QrCount")]
            public required int QrCount { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RequiresVerification")]
            public required bool RequiresVerification { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SiblingCount")]
            public required int SiblingCount { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Tag")]
            public required string Tag { get; set; }
            public class RetailIdRetailIdPackagesInfoPackagesItemIssuancesItem
            {
                [global::System.Text.Json.Serialization.JsonPropertyName("CreatedAt")]
                public required string CreatedAt { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("LabelQuantity")]
                public required double LabelQuantity { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("LabelSet")]
                public required int LabelSet { get; set; }
            }
        }
    }
}
