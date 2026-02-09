package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransferDeliveryTransporterDetail {
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
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
}
