using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class SalesUpdateReceiptsRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("CaregiverLicenseNumber")]
        public string? CaregiverLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExternalReceiptNumber")]
        public string? ExternalReceiptNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IdentificationMethod")]
        public string? IdentificationMethod { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientRegistrationLocationId")]
        public int? PatientRegistrationLocationId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SalesCustomerType")]
        public string? SalesCustomerType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SalesDateTime")]
        public string? SalesDateTime { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Transactions")]
        public List<object>? Transactions { get; set; }
        public class SalesUpdateReceiptsRequestItemTransactionsItem
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
            public double? Quantity { get; set; }
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
}
