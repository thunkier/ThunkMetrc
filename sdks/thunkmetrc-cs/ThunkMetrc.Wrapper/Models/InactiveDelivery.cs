using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class InactiveDelivery
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
}
