package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class ActiveSaleDeliveriesRetailer {
    @JsonProperty("AcceptedDateTime")
    public String acceptedDateTime;
    @JsonProperty("ActualDepartureDateTime")
    public String actualDepartureDateTime;
    @JsonProperty("AllowFullEdit")
    public Boolean allowFullEdit;
    @JsonProperty("CompletedDateTime")
    public String completedDateTime;
    @JsonProperty("DateTime")
    public String dateTime;
    @JsonProperty("Destinations")
    public List<Object> destinations;
    @JsonProperty("Direction")
    public String direction;
    @JsonProperty("DriverEmployeeId")
    public String driverEmployeeId;
    @JsonProperty("DriverName")
    public String driverName;
    @JsonProperty("DriversLicenseNumber")
    public String driversLicenseNumber;
    @JsonProperty("EstimatedDepartureDateTime")
    public String estimatedDepartureDateTime;
    @JsonProperty("FacilityLicenseNumber")
    public String facilityLicenseNumber;
    @JsonProperty("FacilityName")
    public String facilityName;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("Leg")
    public Integer leg;
    @JsonProperty("Packages")
    public List<Object> packages;
    @JsonProperty("RecordedByUserName")
    public String recordedByUserName;
    @JsonProperty("RecordedDateTime")
    public String recordedDateTime;
    @JsonProperty("RestockDateTime")
    public String restockDateTime;
    @JsonProperty("RetailerDeliveryNumber")
    public String retailerDeliveryNumber;
    @JsonProperty("RetailerDeliveryState")
    public String retailerDeliveryState;
    @JsonProperty("TotalPackages")
    public Integer totalPackages;
    @JsonProperty("TotalPrice")
    public Integer totalPrice;
    @JsonProperty("TotalPriceSold")
    public Integer totalPriceSold;
    @JsonProperty("VehicleInfo")
    public String vehicleInfo;
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
    @JsonProperty("VoidedDate")
    public String voidedDate;
}
