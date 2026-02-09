using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class SaleDeliveryRetailer
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("AcceptedDateTime")]
        public string? AcceptedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDepartureDateTime")]
        public string? ActualDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AllowFullEdit")]
        public required bool AllowFullEdit { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("CompletedDateTime")]
        public string? CompletedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DateTime")]
        public required string DateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Destinations")]
        public required List<SalesSaleDeliveryRetailerDestinationsItem> Destinations { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Direction")]
        public required string Direction { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverEmployeeId")]
        public required string DriverEmployeeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
        public required string DriverName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriversLicenseNumber")]
        public required string DriversLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
        public required string EstimatedDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityLicenseNumber")]
        public required string FacilityLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityName")]
        public required string FacilityName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Leg")]
        public required int Leg { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
        public required List<SalesSaleDeliveryRetailerPackagesItem> Packages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecordedByUserName")]
        public string? RecordedByUserName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RecordedDateTime")]
        public required string RecordedDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RestockDateTime")]
        public string? RestockDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryNumber")]
        public required string RetailerDeliveryNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryState")]
        public required string RetailerDeliveryState { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalPackages")]
        public required int TotalPackages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalPrice")]
        public required int TotalPrice { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalPriceSold")]
        public required int TotalPriceSold { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleInfo")]
        public required string VehicleInfo { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
        public required string VehicleLicensePlateNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
        public required string VehicleMake { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
        public required string VehicleModel { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VoidedDate")]
        public string? VoidedDate { get; set; }
        public class SalesSaleDeliveryRetailerDestinationsItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("AcceptedDateTime")]
            public string? AcceptedDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ActualArrivalDateTime")]
            public string? ActualArrivalDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ActualDepartureDateTime")]
            public string? ActualDepartureDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ActualReturnArrivalDateTime")]
            public string? ActualReturnArrivalDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ActualReturnDepartureDateTime")]
            public string? ActualReturnDepartureDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CompletedDateTime")]
            public string? CompletedDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ConsumerId")]
            public required string ConsumerId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("DeliveryNumber")]
            public string? DeliveryNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Direction")]
            public required string Direction { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("DriverEmployeeId")]
            public required string DriverEmployeeId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
            public required string DriverName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("DriversLicenseNumber")]
            public required string DriversLicenseNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedArrivalDateTime")]
            public required string EstimatedArrivalDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
            public required string EstimatedDepartureDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedReturnArrivalDateTime")]
            public string? EstimatedReturnArrivalDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedReturnDepartureDateTime")]
            public string? EstimatedReturnDepartureDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("FacilityLicenseNumber")]
            public string? FacilityLicenseNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("FacilityName")]
            public string? FacilityName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
            public required int Id { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
            public required string LastModified { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
            public string? PatientLicenseNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PlannedRoute")]
            public required string PlannedRoute { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressCity")]
            public required string RecipientAddressCity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressCounty")]
            public string? RecipientAddressCounty { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressPostalCode")]
            public required string RecipientAddressPostalCode { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressState")]
            public required string RecipientAddressState { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressStreet1")]
            public required string RecipientAddressStreet1 { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressStreet2")]
            public required string RecipientAddressStreet2 { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientDeliveryZoneId")]
            public int? RecipientDeliveryZoneId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientDeliveryZoneName")]
            public string? RecipientDeliveryZoneName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientName")]
            public string? RecipientName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientZoneId")]
            public int? RecipientZoneId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecordedByUserName")]
            public string? RecordedByUserName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecordedDateTime")]
            public required string RecordedDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesCustomerType")]
            public required string SalesCustomerType { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDateTime")]
            public required string SalesDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDeliveryState")]
            public required string SalesDeliveryState { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ShipperFacilityLicenseNumber")]
            public string? ShipperFacilityLicenseNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TotalPackages")]
            public required int TotalPackages { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TotalPrice")]
            public required int TotalPrice { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Transactions")]
            public required List<object> Transactions { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TransporterFacilityId")]
            public int? TransporterFacilityId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TransporterFacilityLicenseNumber")]
            public string? TransporterFacilityLicenseNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TransporterFacilityName")]
            public string? TransporterFacilityName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
            public required string VehicleLicensePlateNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
            public required string VehicleMake { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
            public required string VehicleModel { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("VoidedDate")]
            public string? VoidedDate { get; set; }
        }
        public class SalesSaleDeliveryRetailerPackagesItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("ArchivedDate")]
            public string? ArchivedDate { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("CompletedDateTime")]
            public string? CompletedDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsOnInvestigationHold")]
            public required bool IsOnInvestigationHold { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsOnInvestigationRecall")]
            public required bool IsOnInvestigationRecall { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsOnRecall")]
            public bool? IsOnRecall { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("IsOnRecallCombined")]
            public required bool IsOnRecallCombined { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemServingSize")]
            public string? ItemServingSize { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemStrainName")]
            public string? ItemStrainName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemSupplyDurationDays")]
            public int? ItemSupplyDurationDays { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdContent")]
            public double? ItemUnitCbdContent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdContentDose")]
            public double? ItemUnitCbdContentDose { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdContentDoseUnitOfMeasureName")]
            public string? ItemUnitCbdContentDoseUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdContentUnitOfMeasureName")]
            public string? ItemUnitCbdContentUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitCbdPercent")]
            public double? ItemUnitCbdPercent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitQuantity")]
            public double? ItemUnitQuantity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitQuantityUnitOfMeasureName")]
            public string? ItemUnitQuantityUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcContent")]
            public double? ItemUnitThcContent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcContentDose")]
            public double? ItemUnitThcContentDose { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcContentDoseUnitOfMeasureName")]
            public string? ItemUnitThcContentDoseUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcContentUnitOfMeasureName")]
            public string? ItemUnitThcContentUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitThcPercent")]
            public double? ItemUnitThcPercent { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitVolume")]
            public double? ItemUnitVolume { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitVolumeUnitOfMeasureName")]
            public string? ItemUnitVolumeUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitWeight")]
            public double? ItemUnitWeight { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ItemUnitWeightUnitOfMeasureName")]
            public string? ItemUnitWeightUnitOfMeasureName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
            public required string LastModified { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PackageId")]
            public required int PackageId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
            public required string PackageLabel { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ProductCategoryName")]
            public string? ProductCategoryName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("ProductName")]
            public string? ProductName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
            public required double Quantity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecordedByUserName")]
            public string? RecordedByUserName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecordedDateTime")]
            public required string RecordedDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RetailerDeliveryState")]
            public string? RetailerDeliveryState { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TotalPrice")]
            public required double TotalPrice { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasureAbbreviation")]
            public string? UnitOfMeasureAbbreviation { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasureName")]
            public required string UnitOfMeasureName { get; set; }
        }
    }
}
