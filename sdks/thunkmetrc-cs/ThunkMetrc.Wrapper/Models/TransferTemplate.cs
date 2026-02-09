using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransferTemplate
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualArrivalDateTime")]
        public string? ActualArrivalDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDepartureDateTime")]
        public string? ActualDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualReturnArrivalDateTime")]
        public string? ActualReturnArrivalDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualReturnDepartureDateTime")]
        public string? ActualReturnDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ContainsDonation")]
        public required bool ContainsDonation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ContainsPlantPackage")]
        public required bool ContainsPlantPackage { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ContainsProductPackage")]
        public required bool ContainsProductPackage { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ContainsProductRequiresRemediation")]
        public required bool ContainsProductRequiresRemediation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ContainsRemediatedProductPackage")]
        public required bool ContainsRemediatedProductPackage { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ContainsTestingSample")]
        public required bool ContainsTestingSample { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ContainsTradeSample")]
        public required bool ContainsTradeSample { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CreatedByUserName")]
        public string? CreatedByUserName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CreatedDateTime")]
        public required string CreatedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DeliveryCount")]
        public required int DeliveryCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DeliveryId")]
        public required int DeliveryId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DeliveryPackageCount")]
        public required int DeliveryPackageCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DeliveryReceivedPackageCount")]
        public required int DeliveryReceivedPackageCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
        public required string DriverName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverOccupationalLicenseNumber")]
        public required string DriverOccupationalLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverVehicleLicenseNumber")]
        public required string DriverVehicleLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedArrivalDateTime")]
        public required string EstimatedArrivalDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
        public required string EstimatedDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedReturnArrivalDateTime")]
        public string? EstimatedReturnArrivalDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedReturnDepartureDateTime")]
        public string? EstimatedReturnDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("InvoiceNumber")]
        public string? InvoiceNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsVoided")]
        public required bool IsVoided { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ManifestNumber")]
        public string? ManifestNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public required string Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("OriginatingTemplateId")]
        public int? OriginatingTemplateId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageCount")]
        public required int PackageCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedDateTime")]
        public string? ReceivedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedDeliveryCount")]
        public required int ReceivedDeliveryCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedPackageCount")]
        public required int ReceivedPackageCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecipientFacilityLicenseNumber")]
        public string? RecipientFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecipientFacilityName")]
        public string? RecipientFacilityName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipmentLicenseType")]
        public required int ShipmentLicenseType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipmentTransactionType")]
        public string? ShipmentTransactionType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipmentTypeName")]
        public string? ShipmentTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperFacilityLicenseNumber")]
        public required string ShipperFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperFacilityName")]
        public required string ShipperFacilityName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TransporterFacilityLicenseNumber")]
        public required string TransporterFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TransporterFacilityName")]
        public required string TransporterFacilityName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
        public required string VehicleLicensePlateNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
        public required string VehicleMake { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
        public required string VehicleModel { get; set; }
    }
}
