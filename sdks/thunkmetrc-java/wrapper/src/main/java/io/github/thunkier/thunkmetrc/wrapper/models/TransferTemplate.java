package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransferTemplate {
    @JsonProperty("ActualArrivalDateTime")
    public String actualArrivalDateTime;
    @JsonProperty("ActualDepartureDateTime")
    public String actualDepartureDateTime;
    @JsonProperty("ActualReturnArrivalDateTime")
    public String actualReturnArrivalDateTime;
    @JsonProperty("ActualReturnDepartureDateTime")
    public String actualReturnDepartureDateTime;
    @JsonProperty("ContainsDonation")
    public Boolean containsDonation;
    @JsonProperty("ContainsPlantPackage")
    public Boolean containsPlantPackage;
    @JsonProperty("ContainsProductPackage")
    public Boolean containsProductPackage;
    @JsonProperty("ContainsProductRequiresRemediation")
    public Boolean containsProductRequiresRemediation;
    @JsonProperty("ContainsRemediatedProductPackage")
    public Boolean containsRemediatedProductPackage;
    @JsonProperty("ContainsTestingSample")
    public Boolean containsTestingSample;
    @JsonProperty("ContainsTradeSample")
    public Boolean containsTradeSample;
    @JsonProperty("CreatedByUserName")
    public String createdByUserName;
    @JsonProperty("CreatedDateTime")
    public String createdDateTime;
    @JsonProperty("DeliveryCount")
    public Integer deliveryCount;
    @JsonProperty("DeliveryId")
    public Integer deliveryId;
    @JsonProperty("DeliveryPackageCount")
    public Integer deliveryPackageCount;
    @JsonProperty("DeliveryReceivedPackageCount")
    public Integer deliveryReceivedPackageCount;
    @JsonProperty("DriverName")
    public String driverName;
    @JsonProperty("DriverOccupationalLicenseNumber")
    public String driverOccupationalLicenseNumber;
    @JsonProperty("DriverVehicleLicenseNumber")
    public String driverVehicleLicenseNumber;
    @JsonProperty("EstimatedArrivalDateTime")
    public String estimatedArrivalDateTime;
    @JsonProperty("EstimatedDepartureDateTime")
    public String estimatedDepartureDateTime;
    @JsonProperty("EstimatedReturnArrivalDateTime")
    public String estimatedReturnArrivalDateTime;
    @JsonProperty("EstimatedReturnDepartureDateTime")
    public String estimatedReturnDepartureDateTime;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("InvoiceNumber")
    public String invoiceNumber;
    @JsonProperty("IsVoided")
    public Boolean isVoided;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("ManifestNumber")
    public String manifestNumber;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("OriginatingTemplateId")
    public Integer originatingTemplateId;
    @JsonProperty("PackageCount")
    public Integer packageCount;
    @JsonProperty("ReceivedDateTime")
    public String receivedDateTime;
    @JsonProperty("ReceivedDeliveryCount")
    public Integer receivedDeliveryCount;
    @JsonProperty("ReceivedPackageCount")
    public Integer receivedPackageCount;
    @JsonProperty("RecipientFacilityLicenseNumber")
    public String recipientFacilityLicenseNumber;
    @JsonProperty("RecipientFacilityName")
    public String recipientFacilityName;
    @JsonProperty("ShipmentLicenseType")
    public Integer shipmentLicenseType;
    @JsonProperty("ShipmentTransactionType")
    public String shipmentTransactionType;
    @JsonProperty("ShipmentTypeName")
    public String shipmentTypeName;
    @JsonProperty("ShipperFacilityLicenseNumber")
    public String shipperFacilityLicenseNumber;
    @JsonProperty("ShipperFacilityName")
    public String shipperFacilityName;
    @JsonProperty("TransporterFacilityLicenseNumber")
    public String transporterFacilityLicenseNumber;
    @JsonProperty("TransporterFacilityName")
    public String transporterFacilityName;
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
}
