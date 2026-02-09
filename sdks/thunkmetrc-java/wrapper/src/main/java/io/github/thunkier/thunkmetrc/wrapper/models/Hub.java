package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class Hub {
    @JsonProperty("ActualArrivalDateTime")
    public String actualArrivalDateTime;
    @JsonProperty("ActualDepartureDateTime")
    public String actualDepartureDateTime;
    @JsonProperty("ActualReturnArrivalDateTime")
    public String actualReturnArrivalDateTime;
    @JsonProperty("ActualReturnDepartureDateTime")
    public String actualReturnDepartureDateTime;
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
    @JsonProperty("IsLayover")
    public Boolean isLayover;
    @JsonProperty("IsVoided")
    public Boolean isVoided;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("ManifestNumber")
    public String manifestNumber;
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
    @JsonProperty("RejectedPackagesReturned")
    public Boolean rejectedPackagesReturned;
    @JsonProperty("ShipmentTransactionType")
    public Integer shipmentTransactionType;
    @JsonProperty("ShipmentTransporterDetails")
    public List<TransfersHubShipmentTransporterDetailsItem> shipmentTransporterDetails;
    @JsonProperty("ShipmentTypeName")
    public String shipmentTypeName;
    @JsonProperty("ShipperFacilityLicenseNumber")
    public String shipperFacilityLicenseNumber;
    @JsonProperty("ShipperFacilityName")
    public String shipperFacilityName;
    @JsonProperty("TransporterAcceptedDateTime")
    public String transporterAcceptedDateTime;
    @JsonProperty("TransporterActualArrivalDateTime")
    public String transporterActualArrivalDateTime;
    @JsonProperty("TransporterActualDepartureDateTime")
    public String transporterActualDepartureDateTime;
    @JsonProperty("TransporterEstimatedArrivalDateTime")
    public String transporterEstimatedArrivalDateTime;
    @JsonProperty("TransporterEstimatedDepartureDateTime")
    public String transporterEstimatedDepartureDateTime;
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
    public static class TransfersHubShipmentTransporterDetailsItem {
    @JsonProperty("ActualDriverStartDateTime")
    public String actualDriverStartDateTime;
    @JsonProperty("DriverLayoverLeg")
    public String driverLayoverLeg;
    @JsonProperty("DriverName")
    public String driverName;
    @JsonProperty("DriverOccupationalLicenseNumber")
    public String driverOccupationalLicenseNumber;
    @JsonProperty("DriverVehicleLicenseNumber")
    public String driverVehicleLicenseNumber;
    @JsonProperty("ShipmentDeliveryId")
    public Integer shipmentDeliveryId;
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
}
}
