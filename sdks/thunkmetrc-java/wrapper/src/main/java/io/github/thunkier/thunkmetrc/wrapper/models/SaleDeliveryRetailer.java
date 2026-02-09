package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class SaleDeliveryRetailer {
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
    public List<SalesSaleDeliveryRetailerDestinationsItem> destinations;
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
    public List<SalesSaleDeliveryRetailerPackagesItem> packages;
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
    public static class SalesSaleDeliveryRetailerDestinationsItem {
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
    public static class SalesSaleDeliveryRetailerPackagesItem {
    @JsonProperty("ArchivedDate")
    public String archivedDate;
    @JsonProperty("CompletedDateTime")
    public String completedDateTime;
    @JsonProperty("IsOnInvestigationHold")
    public Boolean isOnInvestigationHold;
    @JsonProperty("IsOnInvestigationRecall")
    public Boolean isOnInvestigationRecall;
    @JsonProperty("IsOnRecall")
    public Boolean isOnRecall;
    @JsonProperty("IsOnRecallCombined")
    public Boolean isOnRecallCombined;
    @JsonProperty("ItemServingSize")
    public String itemServingSize;
    @JsonProperty("ItemStrainName")
    public String itemStrainName;
    @JsonProperty("ItemSupplyDurationDays")
    public Integer itemSupplyDurationDays;
    @JsonProperty("ItemUnitCbdContent")
    public Double itemUnitCbdContent;
    @JsonProperty("ItemUnitCbdContentDose")
    public Double itemUnitCbdContentDose;
    @JsonProperty("ItemUnitCbdContentDoseUnitOfMeasureName")
    public String itemUnitCbdContentDoseUnitOfMeasureName;
    @JsonProperty("ItemUnitCbdContentUnitOfMeasureName")
    public String itemUnitCbdContentUnitOfMeasureName;
    @JsonProperty("ItemUnitCbdPercent")
    public Double itemUnitCbdPercent;
    @JsonProperty("ItemUnitQuantity")
    public Double itemUnitQuantity;
    @JsonProperty("ItemUnitQuantityUnitOfMeasureName")
    public String itemUnitQuantityUnitOfMeasureName;
    @JsonProperty("ItemUnitThcContent")
    public Double itemUnitThcContent;
    @JsonProperty("ItemUnitThcContentDose")
    public Double itemUnitThcContentDose;
    @JsonProperty("ItemUnitThcContentDoseUnitOfMeasureName")
    public String itemUnitThcContentDoseUnitOfMeasureName;
    @JsonProperty("ItemUnitThcContentUnitOfMeasureName")
    public String itemUnitThcContentUnitOfMeasureName;
    @JsonProperty("ItemUnitThcPercent")
    public Double itemUnitThcPercent;
    @JsonProperty("ItemUnitVolume")
    public Double itemUnitVolume;
    @JsonProperty("ItemUnitVolumeUnitOfMeasureName")
    public String itemUnitVolumeUnitOfMeasureName;
    @JsonProperty("ItemUnitWeight")
    public Double itemUnitWeight;
    @JsonProperty("ItemUnitWeightUnitOfMeasureName")
    public String itemUnitWeightUnitOfMeasureName;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("PackageId")
    public Integer packageId;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("ProductCategoryName")
    public String productCategoryName;
    @JsonProperty("ProductName")
    public String productName;
    @JsonProperty("Quantity")
    public Double quantity;
    @JsonProperty("RecordedByUserName")
    public String recordedByUserName;
    @JsonProperty("RecordedDateTime")
    public String recordedDateTime;
    @JsonProperty("RetailerDeliveryState")
    public String retailerDeliveryState;
    @JsonProperty("TotalPrice")
    public Double totalPrice;
    @JsonProperty("UnitOfMeasureAbbreviation")
    public String unitOfMeasureAbbreviation;
    @JsonProperty("UnitOfMeasureName")
    public String unitOfMeasureName;
}
}
