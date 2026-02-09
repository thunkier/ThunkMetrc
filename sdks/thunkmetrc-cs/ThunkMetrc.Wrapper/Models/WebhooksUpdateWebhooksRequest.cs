using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class WebhooksUpdateWebhooksRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("errorResponseJsonTemplate")]
        public string? errorResponseJsonTemplate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("facilityLicenseNumbers")]
        public string? facilityLicenseNumbers { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("objectType")]
        public string? objectType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("serverPublicKeyFingerprint")]
        public string? serverPublicKeyFingerprint { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("status")]
        public string? status { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("template")]
        public string? template { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("tpiApiKey")]
        public string? tpiApiKey { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("url")]
        public string? url { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("userApiKey")]
        public string? userApiKey { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("verb")]
        public string? verb { get; set; }
    }
}
