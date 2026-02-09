package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransfersUpdateOutgoingTemplatesRequestItem {
    @JsonProperty("Destinations")
    public List<Object> destinations;
    @JsonProperty("DriverLicenseNumber")
    public String driverLicenseNumber;
    @JsonProperty("DriverName")
    public String driverName;
    @JsonProperty("DriverOccupationalLicenseNumber")
    public String driverOccupationalLicenseNumber;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("PhoneNumberForQuestions")
    public String phoneNumberForQuestions;
    @JsonProperty("TransferTemplateId")
    public Integer transferTemplateId;
    @JsonProperty("TransporterFacilityLicenseNumber")
    public String transporterFacilityLicenseNumber;
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
    public static class TransfersUpdateOutgoingTemplatesRequestItemDestinationsItemPackagesItem {
    @JsonProperty("GrossUnitOfWeightName")
    public String grossUnitOfWeightName;
    @JsonProperty("GrossWeight")
    public Double grossWeight;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("WholesalePrice")
    public String wholesalePrice;
}public static class TransfersUpdateOutgoingTemplatesRequestItemDestinationsItemTransportersItemTransporterDetailsItem {
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
}public static class TransfersUpdateOutgoingTemplatesRequestItemDestinationsItemTransportersItem {
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
}public static class TransfersUpdateOutgoingTemplatesRequestItemDestinationsItem {
    @JsonProperty("EstimatedArrivalDateTime")
    public String estimatedArrivalDateTime;
    @JsonProperty("EstimatedDepartureDateTime")
    public String estimatedDepartureDateTime;
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
