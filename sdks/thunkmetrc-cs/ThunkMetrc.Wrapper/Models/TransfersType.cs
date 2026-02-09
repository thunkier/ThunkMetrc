using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransfersType
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("BypassApproval")]
        public required bool BypassApproval { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExternalIncomingCanRecordExternalIdentifier")]
        public required bool ExternalIncomingCanRecordExternalIdentifier { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExternalIncomingExternalIdentifierRequired")]
        public required bool ExternalIncomingExternalIdentifierRequired { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExternalOutgoingCanRecordExternalIdentifier")]
        public required bool ExternalOutgoingCanRecordExternalIdentifier { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExternalOutgoingExternalIdentifierRequired")]
        public required bool ExternalOutgoingExternalIdentifierRequired { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ForExternalIncomingShipments")]
        public required bool ForExternalIncomingShipments { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ForExternalOutgoingShipments")]
        public required bool ForExternalOutgoingShipments { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ForLicensedShipments")]
        public required bool ForLicensedShipments { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresDestinationGrossWeight")]
        public required bool RequiresDestinationGrossWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresInvoiceNumber")]
        public required bool RequiresInvoiceNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresPDFDocument")]
        public required bool RequiresPDFDocument { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiresPackagesGrossWeight")]
        public required bool RequiresPackagesGrossWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TransactionType")]
        public required string TransactionType { get; set; }
    }
}
