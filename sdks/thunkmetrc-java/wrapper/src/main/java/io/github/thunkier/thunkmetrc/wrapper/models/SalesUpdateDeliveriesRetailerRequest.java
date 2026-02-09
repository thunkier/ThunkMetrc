package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesUpdateDeliveriesRetailerRequest {
    @JsonProperty("DateTime")
    public String dateTime;
    @JsonProperty("Destinations")
    public List<Object> destinations;
    @JsonProperty("DriverEmployeeId")
    public String driverEmployeeId;
    @JsonProperty("DriverName")
    public String driverName;
    @JsonProperty("DriversLicenseNumber")
    public String driversLicenseNumber;
    @JsonProperty("EstimatedDepartureDateTime")
    public String estimatedDepartureDateTime;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Packages")
    public List<Object> packages;
    @JsonProperty("PhoneNumberForQuestions")
    public String phoneNumberForQuestions;
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
    public static class SalesUpdateDeliveriesRetailerRequestDestinationsItemTransactionsItem {
    @JsonProperty("CityTax")
    public String cityTax;
    @JsonProperty("CountyTax")
    public String countyTax;
    @JsonProperty("DiscountAmount")
    public String discountAmount;
    @JsonProperty("ExciseTax")
    public String exciseTax;
    @JsonProperty("InvoiceNumber")
    public String invoiceNumber;
    @JsonProperty("MunicipalTax")
    public String municipalTax;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("Price")
    public String price;
    @JsonProperty("QrCodes")
    public String qrCodes;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("SalesTax")
    public String salesTax;
    @JsonProperty("SubTotal")
    public String subTotal;
    @JsonProperty("TotalAmount")
    public Double totalAmount;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
    @JsonProperty("UnitThcContent")
    public Double unitThcContent;
    @JsonProperty("UnitThcContentUnitOfMeasure")
    public String unitThcContentUnitOfMeasure;
    @JsonProperty("UnitThcPercent")
    public Double unitThcPercent;
    @JsonProperty("UnitWeight")
    public Double unitWeight;
    @JsonProperty("UnitWeightUnitOfMeasure")
    public String unitWeightUnitOfMeasure;
}public static class SalesUpdateDeliveriesRetailerRequestDestinationsItem {
    @JsonProperty("ConsumerId")
    public String consumerId;
    @JsonProperty("DriverEmployeeId")
    public Integer driverEmployeeId;
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
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PhoneNumberForQuestions")
    public String phoneNumberForQuestions;
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
    @JsonProperty("RecipientName")
    public String recipientName;
    @JsonProperty("RecipientZoneId")
    public String recipientZoneId;
    @JsonProperty("SalesCustomerType")
    public String salesCustomerType;
    @JsonProperty("SalesDateTime")
    public String salesDateTime;
    @JsonProperty("Transactions")
    public List<Object> transactions;
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
}
    public static class SalesUpdateDeliveriesRetailerRequestPackagesItem {
    @JsonProperty("DateTime")
    public String dateTime;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("Quantity")
    public Integer quantity;
    @JsonProperty("TotalPrice")
    public Double totalPrice;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
}
