package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransfersUpdateIncomingExternalRequest {
    @JsonProperty("Destinations")
    public List<Object> destinations;
    @JsonProperty("DriverLicenseNumber")
    public String driverLicenseNumber;
    @JsonProperty("DriverName")
    public String driverName;
    @JsonProperty("DriverOccupationalLicenseNumber")
    public String driverOccupationalLicenseNumber;
    @JsonProperty("PhoneNumberForQuestions")
    public String phoneNumberForQuestions;
    @JsonProperty("ShipperAddress1")
    public String shipperAddress1;
    @JsonProperty("ShipperAddress2")
    public String shipperAddress2;
    @JsonProperty("ShipperAddressCity")
    public String shipperAddressCity;
    @JsonProperty("ShipperAddressPostalCode")
    public String shipperAddressPostalCode;
    @JsonProperty("ShipperAddressState")
    public String shipperAddressState;
    @JsonProperty("ShipperLicenseNumber")
    public String shipperLicenseNumber;
    @JsonProperty("ShipperMainPhoneNumber")
    public String shipperMainPhoneNumber;
    @JsonProperty("ShipperName")
    public String shipperName;
    @JsonProperty("TransferId")
    public Integer transferId;
    @JsonProperty("TransporterFacilityLicenseNumber")
    public String transporterFacilityLicenseNumber;
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
    public static class TransfersUpdateIncomingExternalRequestDestinationsItemPackagesItem {
    @JsonProperty("ExpirationDate")
    public String expirationDate;
    @JsonProperty("ExternalId")
    public String externalId;
    @JsonProperty("GrossUnitOfWeightName")
    public String grossUnitOfWeightName;
    @JsonProperty("GrossWeight")
    public Double grossWeight;
    @JsonProperty("ItemName")
    public String itemName;
    @JsonProperty("PackagedDate")
    public String packagedDate;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("SellByDate")
    public String sellByDate;
    @JsonProperty("UnitOfMeasureName")
    public String unitOfMeasureName;
    @JsonProperty("UseByDate")
    public String useByDate;
    @JsonProperty("WholesalePrice")
    public String wholesalePrice;
}public static class TransfersUpdateIncomingExternalRequestDestinationsItemTransportersItemTransporterDetailsItem {
    @JsonProperty("DriverLayoverLeg")
    public String driverLayoverLeg;
    @JsonProperty("DriverLicenseNumber")
    public String driverLicenseNumber;
    @JsonProperty("DriverName")
    public String driverName;
    @JsonProperty("DriverOccupationalLicenseNumber")
    public String driverOccupationalLicenseNumber;
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
}public static class TransfersUpdateIncomingExternalRequestDestinationsItemTransportersItem {
    @JsonProperty("DriverLayoverLeg")
    public String driverLayoverLeg;
    @JsonProperty("DriverLicenseNumber")
    public String driverLicenseNumber;
    @JsonProperty("DriverName")
    public String driverName;
    @JsonProperty("DriverOccupationalLicenseNumber")
    public String driverOccupationalLicenseNumber;
    @JsonProperty("EstimatedArrivalDateTime")
    public String estimatedArrivalDateTime;
    @JsonProperty("EstimatedDepartureDateTime")
    public String estimatedDepartureDateTime;
    @JsonProperty("IsLayover")
    public Boolean isLayover;
    @JsonProperty("PhoneNumberForQuestions")
    public String phoneNumberForQuestions;
    @JsonProperty("TransporterDetails")
    public List<Object> transporterDetails;
    @JsonProperty("TransporterFacilityLicenseNumber")
    public String transporterFacilityLicenseNumber;
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
}public static class TransfersUpdateIncomingExternalRequestDestinationsItem {
    @JsonProperty("EstimatedArrivalDateTime")
    public String estimatedArrivalDateTime;
    @JsonProperty("EstimatedDepartureDateTime")
    public String estimatedDepartureDateTime;
    @JsonProperty("GrossUnitOfWeightId")
    public Integer grossUnitOfWeightId;
    @JsonProperty("GrossWeight")
    public Double grossWeight;
    @JsonProperty("InvoiceNumber")
    public String invoiceNumber;
    @JsonProperty("Packages")
    public List<Object> packages;
    @JsonProperty("PlannedRoute")
    public String plannedRoute;
    @JsonProperty("RecipientLicenseNumber")
    public String recipientLicenseNumber;
    @JsonProperty("TransferDestinationId")
    public Integer transferDestinationId;
    @JsonProperty("TransferTypeName")
    public String transferTypeName;
    @JsonProperty("Transporters")
    public List<Object> transporters;
}
}
