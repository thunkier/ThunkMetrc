using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class UpdateDeliveriesRetailerRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("DateTime")]
        public string? DateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Destinations")]
        public List<object>? Destinations { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverEmployeeId")]
        public string? DriverEmployeeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
        public string? DriverName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DriversLicenseNumber")]
        public string? DriversLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
        public string? EstimatedDepartureDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Packages")]
        public List<object>? Packages { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PhoneNumberForQuestions")]
        public string? PhoneNumberForQuestions { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
        public string? VehicleLicensePlateNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
        public string? VehicleMake { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
        public string? VehicleModel { get; set; }
        public class UpdateDeliveriesRetailerRequestDestinationsItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("ConsumerId")]
            public string? ConsumerId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("DriverEmployeeId")]
            public int? DriverEmployeeId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("DriverName")]
            public string? DriverName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("DriversLicenseNumber")]
            public string? DriversLicenseNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedArrivalDateTime")]
            public string? EstimatedArrivalDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("EstimatedDepartureDateTime")]
            public string? EstimatedDepartureDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
            public int? Id { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
            public string? PatientLicenseNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PhoneNumberForQuestions")]
            public string? PhoneNumberForQuestions { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PlannedRoute")]
            public string? PlannedRoute { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressCity")]
            public string? RecipientAddressCity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressCounty")]
            public string? RecipientAddressCounty { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressPostalCode")]
            public string? RecipientAddressPostalCode { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressState")]
            public string? RecipientAddressState { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressStreet1")]
            public string? RecipientAddressStreet1 { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientAddressStreet2")]
            public string? RecipientAddressStreet2 { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientName")]
            public string? RecipientName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("RecipientZoneId")]
            public string? RecipientZoneId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesCustomerType")]
            public string? SalesCustomerType { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("SalesDateTime")]
            public string? SalesDateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Transactions")]
            public List<object>? Transactions { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("VehicleLicensePlateNumber")]
            public string? VehicleLicensePlateNumber { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("VehicleMake")]
            public string? VehicleMake { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("VehicleModel")]
            public string? VehicleModel { get; set; }
            public class UpdateDeliveriesRetailerRequestDestinationsItemTransactionsItem
            {
                [global::System.Text.Json.Serialization.JsonPropertyName("CityTax")]
                public string? CityTax { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("CountyTax")]
                public string? CountyTax { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("DiscountAmount")]
                public string? DiscountAmount { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("ExciseTax")]
                public string? ExciseTax { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("InvoiceNumber")]
                public string? InvoiceNumber { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("MunicipalTax")]
                public string? MunicipalTax { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
                public string? PackageLabel { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("Price")]
                public string? Price { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("QrCodes")]
                public string? QrCodes { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
                public int? Quantity { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("SalesTax")]
                public string? SalesTax { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("SubTotal")]
                public string? SubTotal { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("TotalAmount")]
                public double? TotalAmount { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasure")]
                public string? UnitOfMeasure { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcContent")]
                public double? UnitThcContent { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcContentUnitOfMeasure")]
                public string? UnitThcContentUnitOfMeasure { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("UnitThcPercent")]
                public double? UnitThcPercent { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("UnitWeight")]
                public double? UnitWeight { get; set; }
                [global::System.Text.Json.Serialization.JsonPropertyName("UnitWeightUnitOfMeasure")]
                public string? UnitWeightUnitOfMeasure { get; set; }
            }
        }
        public class UpdateDeliveriesRetailerRequestPackagesItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("DateTime")]
            public string? DateTime { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("PackageLabel")]
            public string? PackageLabel { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
            public int? Quantity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("TotalPrice")]
            public double? TotalPrice { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasure")]
            public string? UnitOfMeasure { get; set; }
        }
    }
}
