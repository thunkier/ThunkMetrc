using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PackagesPackage
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDepartureDateTime")]
        public required string ActualDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExternalId")]
        public int? ExternalId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GrossUnitOfWeightAbbreviation")]
        public required string GrossUnitOfWeightAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("GrossWeight")]
        public required double GrossWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ItemStrainName")]
        public required string ItemStrainName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestingStateName")]
        public required string LabTestingStateName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ManifestNumber")]
        public required string ManifestNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageId")]
        public required int PackageId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
        public required string PackageLabel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProcessingJobTypeName")]
        public required string ProcessingJobTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductCategoryName")]
        public required string ProductCategoryName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductName")]
        public required string ProductName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedDateTime")]
        public required string ReceivedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedQuantity")]
        public required double ReceivedQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceivedUnitOfMeasureAbbreviation")]
        public required string ReceivedUnitOfMeasureAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReceiverWholesalePrice")]
        public required double ReceiverWholesalePrice { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecipientFacilityLicenseNumber")]
        public required string RecipientFacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecipientFacilityName")]
        public required string RecipientFacilityName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipmentPackageStateName")]
        public required string ShipmentPackageStateName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShippedQuantity")]
        public required double ShippedQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShippedUnitOfMeasureAbbreviation")]
        public required string ShippedUnitOfMeasureAbbreviation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ShipperWholesalePrice")]
        public required double ShipperWholesalePrice { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourceHarvestNames")]
        public required string SourceHarvestNames { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SourcePackageLabels")]
        public required string SourcePackageLabels { get; set; }
    }
}
