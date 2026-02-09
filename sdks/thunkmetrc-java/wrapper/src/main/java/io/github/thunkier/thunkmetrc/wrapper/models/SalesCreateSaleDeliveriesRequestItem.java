package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesCreateSaleDeliveriesRequestItem {
    @JsonProperty("ConsumerId")
    public Integer consumerId;
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
    public Integer recipientZoneId;
    @JsonProperty("SalesCustomerType")
    public String salesCustomerType;
    @JsonProperty("SalesDateTime")
    public String salesDateTime;
    @JsonProperty("Transactions")
    public List<SalesCreateSaleDeliveriesRequestItemTransactions> transactions;
    @JsonProperty("TransporterFacilityLicenseNumber")
    public String transporterFacilityLicenseNumber;
    @JsonProperty("VehicleLicensePlateNumber")
    public String vehicleLicensePlateNumber;
    @JsonProperty("VehicleMake")
    public String vehicleMake;
    @JsonProperty("VehicleModel")
    public String vehicleModel;
    public static class SalesCreateSaleDeliveriesRequestItemTransactions {
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
}
}
