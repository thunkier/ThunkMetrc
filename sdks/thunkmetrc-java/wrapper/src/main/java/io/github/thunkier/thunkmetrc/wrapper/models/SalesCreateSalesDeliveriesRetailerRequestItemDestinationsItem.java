package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesCreateSalesDeliveriesRetailerRequestItemDestinationsItem {
    @JsonProperty("ConsumerId")
    public String consumerId;
    @JsonProperty("EstimatedArrivalDateTime")
    public String estimatedArrivalDateTime;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
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
    @JsonProperty("Transactions")
    public List<Object> transactions;
    public static class SalesCreateSalesDeliveriesRetailerRequestItemDestinationsItemTransactionsItem {
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
    public Double quantity;
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
