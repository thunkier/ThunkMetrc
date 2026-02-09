package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesSaleDeliveryRetailerDestinationsItem {
    @JsonProperty("AcceptedDateTime")
    public String acceptedDateTime;
    @JsonProperty("ActualArrivalDateTime")
    public String actualArrivalDateTime;
    @JsonProperty("ActualDepartureDateTime")
    public String actualDepartureDateTime;
    @JsonProperty("ActualReturnArrivalDateTime")
    public String actualReturnArrivalDateTime;
    @JsonProperty("ActualReturnDepartureDateTime")
    public String actualReturnDepartureDateTime;
    @JsonProperty("CompletedDateTime")
    public String completedDateTime;
    @JsonProperty("ConsumerId")
    public String consumerId;
    @JsonProperty("DeliveryNumber")
    public String deliveryNumber;
    @JsonProperty("Direction")
    public String direction;
    @JsonProperty("DriverEmployeeId")
    public String driverEmployeeId;
    @JsonProperty("DriverName")
    public String driverName;
    @JsonProperty("DriversLicenseNumber")
    public String driversLicenseNumber;
    @JsonProperty("EstimatedArrivalDateTime")
    public String estimatedArrivalDateTime;
    @JsonProperty("EstimatedDepartureDateTime")
    public String estimatedDepartureDateTime;
    @JsonProperty("EstimatedReturnArrivalDateTime")
    public String estimatedReturnArrivalDateTime;
    @JsonProperty("EstimatedReturnDepartureDateTime")
    public String estimatedReturnDepartureDateTime;
    @JsonProperty("FacilityLicenseNumber")
    public String facilityLicenseNumber;
    @JsonProperty("FacilityName")
    public String facilityName;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PlannedRoute")
    public String plannedRoute;
    @JsonProperty("RecipientAddressCity")
    public String recipientAddressCity;
    @JsonProperty("RecipientAddressCounty")
    public String recipientAddressCounty;
    @JsonProperty("RecipientAddressPostalCode")
    public String recipientAddressPostalCode;
    @JsonProperty("RecipientAddressState")
    public String recipientAddressState;
    @JsonProperty("RecipientAddressStreet1")
    public String recipientAddressStreet1;
    @JsonProperty("RecipientAddressStreet2")
    public String recipientAddressStreet2;
    @JsonProperty("RecipientDeliveryZoneId")
    public Integer recipientDeliveryZoneId;
    @JsonProperty("RecipientDeliveryZoneName")
    public String recipientDeliveryZoneName;
    @JsonProperty("RecipientName")
    public String recipientName;
    @JsonProperty("RecipientZoneId")
    public Integer recipientZoneId;
    @JsonProperty("RecordedByUserName")
    public String recordedByUserName;
    @JsonProperty("RecordedDateTime")
    public String recordedDateTime;
    @JsonProperty("SalesCustomerType")
    public String salesCustomerType;
    @JsonProperty("SalesDateTime")
    public String salesDateTime;
    @JsonProperty("SalesDeliveryState")
    public String salesDeliveryState;
    @JsonProperty("ShipperFacilityLicenseNumber")
    public String shipperFacilityLicenseNumber;
    @JsonProperty("TotalPackages")
    public Integer totalPackages;
    @JsonProperty("TotalPrice")
    public Integer totalPrice;
    @JsonProperty("Transactions")
    public List<Object> transactions;
    @JsonProperty("TransporterFacilityId")
    public Integer transporterFacilityId;
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
    @JsonProperty("VoidedDate")
    public String voidedDate;
}
