package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesUpdateReceiptsRequest {
    @JsonProperty("CaregiverLicenseNumber")
    public String caregiverLicenseNumber;
    @JsonProperty("ExternalReceiptNumber")
    public String externalReceiptNumber;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IdentificationMethod")
    public String identificationMethod;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PatientRegistrationLocationId")
    public Integer patientRegistrationLocationId;
    @JsonProperty("SalesCustomerType")
    public String salesCustomerType;
    @JsonProperty("SalesDateTime")
    public String salesDateTime;
    @JsonProperty("Transactions")
    public List<Object> transactions;
    public static class SalesUpdateReceiptsRequestTransactionsItem {
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
