package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesUpdateDeliveriesHubRequest {
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
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("PhoneNumberForQuestions")
    public String phoneNumberForQuestions;
    @JsonProperty("PlannedRoute")
    public String plannedRoute;
    @JsonProperty("TransporterFacilityId")
    public String transporterFacilityId;
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
}
