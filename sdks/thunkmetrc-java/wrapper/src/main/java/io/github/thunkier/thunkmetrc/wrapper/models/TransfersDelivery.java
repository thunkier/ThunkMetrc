package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransfersDelivery {
    @JsonProperty("ActualArrivalDateTime")
    public String actualArrivalDateTime;
    @JsonProperty("ActualDepartureDateTime")
    public String actualDepartureDateTime;
    @JsonProperty("ActualReturnArrivalDateTime")
    public String actualReturnArrivalDateTime;
    @JsonProperty("ActualReturnDepartureDateTime")
    public String actualReturnDepartureDateTime;
    @JsonProperty("DeliveryNumber")
    public Integer deliveryNumber;
    @JsonProperty("DeliveryPackageCount")
    public Integer deliveryPackageCount;
    @JsonProperty("DeliveryReceivedPackageCount")
    public Integer deliveryReceivedPackageCount;
    @JsonProperty("EstimatedArrivalDateTime")
    public String estimatedArrivalDateTime;
    @JsonProperty("EstimatedDepartureDateTime")
    public String estimatedDepartureDateTime;
    @JsonProperty("EstimatedReturnArrivalDateTime")
    public String estimatedReturnArrivalDateTime;
    @JsonProperty("EstimatedReturnDepartureDateTime")
    public String estimatedReturnDepartureDateTime;
    @JsonProperty("GrossUnitOfWeightId")
    public Integer grossUnitOfWeightId;
    @JsonProperty("GrossUnitOfWeightName")
    public String grossUnitOfWeightName;
    @JsonProperty("GrossWeight")
    public Double grossWeight;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("InvoiceNumber")
    public String invoiceNumber;
    @JsonProperty("ManifestNumber")
    public String manifestNumber;
    @JsonProperty("PDFDocumentFileSystemId")
    public Integer pDFDocumentFileSystemId;
    @JsonProperty("PlannedRoute")
    public String plannedRoute;
    @JsonProperty("ReceivedDateTime")
    public String receivedDateTime;
    @JsonProperty("RecipientFacilityLicenseNumber")
    public String recipientFacilityLicenseNumber;
    @JsonProperty("RecipientFacilityName")
    public String recipientFacilityName;
    @JsonProperty("RejectedPackagesReturned")
    public Boolean rejectedPackagesReturned;
    @JsonProperty("ShipmentTransactionType")
    public String shipmentTransactionType;
    @JsonProperty("ShipmentTypeName")
    public String shipmentTypeName;
}
