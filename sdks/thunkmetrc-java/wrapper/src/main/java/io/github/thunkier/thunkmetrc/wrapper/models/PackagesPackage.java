package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackagesPackage {
    @JsonProperty("ActualDepartureDateTime")
    public String actualDepartureDateTime;
    @JsonProperty("ExternalId")
    public Integer externalId;
    @JsonProperty("GrossUnitOfWeightAbbreviation")
    public String grossUnitOfWeightAbbreviation;
    @JsonProperty("GrossWeight")
    public Double grossWeight;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("ItemStrainName")
    public String itemStrainName;
    @JsonProperty("LabTestingStateName")
    public String labTestingStateName;
    @JsonProperty("ManifestNumber")
    public String manifestNumber;
    @JsonProperty("PackageId")
    public Integer packageId;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("ProcessingJobTypeName")
    public String processingJobTypeName;
    @JsonProperty("ProductCategoryName")
    public String productCategoryName;
    @JsonProperty("ProductName")
    public String productName;
    @JsonProperty("ReceivedDateTime")
    public String receivedDateTime;
    @JsonProperty("ReceivedQuantity")
    public Double receivedQuantity;
    @JsonProperty("ReceivedUnitOfMeasureAbbreviation")
    public String receivedUnitOfMeasureAbbreviation;
    @JsonProperty("ReceiverWholesalePrice")
    public Double receiverWholesalePrice;
    @JsonProperty("RecipientFacilityLicenseNumber")
    public String recipientFacilityLicenseNumber;
    @JsonProperty("RecipientFacilityName")
    public String recipientFacilityName;
    @JsonProperty("ShipmentPackageStateName")
    public String shipmentPackageStateName;
    @JsonProperty("ShippedQuantity")
    public Double shippedQuantity;
    @JsonProperty("ShippedUnitOfMeasureAbbreviation")
    public String shippedUnitOfMeasureAbbreviation;
    @JsonProperty("ShipperWholesalePrice")
    public Double shipperWholesalePrice;
    @JsonProperty("SourceHarvestNames")
    public String sourceHarvestNames;
    @JsonProperty("SourcePackageLabels")
    public String sourcePackageLabels;
}
