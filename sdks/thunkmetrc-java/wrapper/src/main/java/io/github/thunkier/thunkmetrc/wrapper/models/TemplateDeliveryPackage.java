package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TemplateDeliveryPackage {
    @JsonProperty("ContainsRemediatedProduct")
    public Boolean containsRemediatedProduct;
    @JsonProperty("ExpirationDate")
    public String expirationDate;
    @JsonProperty("ExternalId")
    public Integer externalId;
    @JsonProperty("GrossUnitOfWeightName")
    public String grossUnitOfWeightName;
    @JsonProperty("IsDonation")
    public Boolean isDonation;
    @JsonProperty("IsFinishedGood")
    public Boolean isFinishedGood;
    @JsonProperty("IsTestingSample")
    public Boolean isTestingSample;
    @JsonProperty("IsTradeSample")
    public Boolean isTradeSample;
    @JsonProperty("IsTradeSamplePersistent")
    public Boolean isTradeSamplePersistent;
    @JsonProperty("ItemBrandName")
    public String itemBrandName;
    @JsonProperty("ItemCategoryName")
    public String itemCategoryName;
    @JsonProperty("ItemId")
    public Integer itemId;
    @JsonProperty("ItemName")
    public String itemName;
    @JsonProperty("ItemServingSize")
    public String itemServingSize;
    @JsonProperty("ItemStrainName")
    public String itemStrainName;
    @JsonProperty("ItemSupplyDurationDays")
    public Integer itemSupplyDurationDays;
    @JsonProperty("ItemUnitCbdAContent")
    public Double itemUnitCbdAContent;
    @JsonProperty("ItemUnitCbdAContentDose")
    public Double itemUnitCbdAContentDose;
    @JsonProperty("ItemUnitCbdAContentDoseUnitOfMeasureName")
    public String itemUnitCbdAContentDoseUnitOfMeasureName;
    @JsonProperty("ItemUnitCbdAContentUnitOfMeasureName")
    public String itemUnitCbdAContentUnitOfMeasureName;
    @JsonProperty("ItemUnitCbdAPercent")
    public Double itemUnitCbdAPercent;
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
    @JsonProperty("ItemUnitThcAContent")
    public Double itemUnitThcAContent;
    @JsonProperty("ItemUnitThcAContentDose")
    public Double itemUnitThcAContentDose;
    @JsonProperty("ItemUnitThcAContentDoseUnitOfMeasureName")
    public String itemUnitThcAContentDoseUnitOfMeasureName;
    @JsonProperty("ItemUnitThcAContentUnitOfMeasureName")
    public String itemUnitThcAContentUnitOfMeasureName;
    @JsonProperty("ItemUnitThcAPercent")
    public Double itemUnitThcAPercent;
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
    @JsonProperty("LabTestingState")
    public String labTestingState;
    @JsonProperty("PackageId")
    public Integer packageId;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("PackageType")
    public String packageType;
    @JsonProperty("PackagedDate")
    public String packagedDate;
    @JsonProperty("ProductLabel")
    public TransfersTemplateDeliveryPackageProductLabel productLabel;
    @JsonProperty("ProductRequiresRemediation")
    public Boolean productRequiresRemediation;
    @JsonProperty("ProductionBatchNumber")
    public String productionBatchNumber;
    @JsonProperty("ReceivedQuantity")
    public Double receivedQuantity;
    @JsonProperty("ReceivedUnitOfMeasureName")
    public String receivedUnitOfMeasureName;
    @JsonProperty("RemediationDate")
    public String remediationDate;
    @JsonProperty("SellByDate")
    public String sellByDate;
    @JsonProperty("ShipmentPackageState")
    public String shipmentPackageState;
    @JsonProperty("ShippedQuantity")
    public Double shippedQuantity;
    @JsonProperty("ShippedUnitOfMeasureName")
    public String shippedUnitOfMeasureName;
    @JsonProperty("SourceHarvestNames")
    public String sourceHarvestNames;
    @JsonProperty("SourcePackageIsDonation")
    public Boolean sourcePackageIsDonation;
    @JsonProperty("SourcePackageIsTradeSample")
    public Boolean sourcePackageIsTradeSample;
    @JsonProperty("SourcePackageLabels")
    public String sourcePackageLabels;
    @JsonProperty("SourceProductionBatchNumbers")
    public String sourceProductionBatchNumbers;
    @JsonProperty("UseByDate")
    public String useByDate;
    public static class TransfersTemplateDeliveryPackageProductLabel {
    @JsonProperty("IsActive")
    public Boolean isActive;
    @JsonProperty("IsChildFromParentWithLabel")
    public Boolean isChildFromParentWithLabel;
    @JsonProperty("LabelSource")
    public String labelSource;
    @JsonProperty("LastLabelGenerationDate")
    public String lastLabelGenerationDate;
    @JsonProperty("OriginalSourcePackageId")
    public Integer originalSourcePackageId;
    @JsonProperty("OriginalSourcePackageLabel")
    public String originalSourcePackageLabel;
    @JsonProperty("QrCodeRanges")
    public String qrCodeRanges;
    @JsonProperty("QrCount")
    public Integer qrCount;
}
}
