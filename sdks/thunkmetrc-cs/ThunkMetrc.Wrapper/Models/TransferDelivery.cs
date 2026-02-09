using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransferDelivery
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualArrivalDateTime")]
        public string? ActualArrivalDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDepartureDateTime")]
        public string? ActualDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualReturnArrivalDateTime")]
        public string? ActualReturnArrivalDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualReturnDepartureDateTime")]
        public string? ActualReturnDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DeliveryNumber")]
        public required int DeliveryNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DeliveryPackageCount")]
        public required int DeliveryPackageCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DeliveryReceivedPackageCount")]
        public required int DeliveryReceivedPackageCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedArrivalDateTime")]
        public required string EstimatedArrivalDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
        public required string EstimatedDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedReturnArrivalDateTime")]
        public string? EstimatedReturnArrivalDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedReturnDepartureDateTime")]
        public string? EstimatedReturnDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GrossUnitOfWeightId")]
        public int? GrossUnitOfWeightId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GrossUnitOfWeightName")]
        public string? GrossUnitOfWeightName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GrossWeight")]
        public double? GrossWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("InvoiceNumber")]
        public string? InvoiceNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ManifestNumber")]
        public string? ManifestNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PDFDocumentFileSystemId")]
        public int? PDFDocumentFileSystemId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlannedRoute")]
        public required string PlannedRoute { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedDateTime")]
        public required string ReceivedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecipientFacilityLicenseNumber")]
        public required string RecipientFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecipientFacilityName")]
        public required string RecipientFacilityName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RejectedPackagesReturned")]
        public required bool RejectedPackagesReturned { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipmentTransactionType")]
        public required string ShipmentTransactionType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipmentTypeName")]
        public required string ShipmentTypeName { get; set; }
    }
}
