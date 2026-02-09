package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class Item {
    @JsonProperty("AdministrationMethod")
    public String administrationMethod;
    @JsonProperty("Allergens")
    public String allergens;
    @JsonProperty("ApprovalStatus")
    public String approvalStatus;
    @JsonProperty("ApprovalStatusDateTime")
    public String approvalStatusDateTime;
    @JsonProperty("DefaultLabTestingState")
    public String defaultLabTestingState;
    @JsonProperty("Description")
    public String description;
    @JsonProperty("GlobalProductName")
    public String globalProductName;
    @JsonProperty("GlobalProductNumber")
    public String globalProductNumber;
    @JsonProperty("HasExpirationDate")
    public Boolean hasExpirationDate;
    @JsonProperty("HasSellByDate")
    public Boolean hasSellByDate;
    @JsonProperty("HasUseByDate")
    public Boolean hasUseByDate;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IsExpirationDateRequired")
    public Boolean isExpirationDateRequired;
    @JsonProperty("IsSellByDateRequired")
    public Boolean isSellByDateRequired;
    @JsonProperty("IsUseByDateRequired")
    public Boolean isUseByDateRequired;
    @JsonProperty("IsUsed")
    public Boolean isUsed;
    @JsonProperty("ItemBrandId")
    public Integer itemBrandId;
    @JsonProperty("ItemBrandName")
    public String itemBrandName;
    @JsonProperty("LabTestBatchNames")
    public List<Object> labTestBatchNames;
    @JsonProperty("LabelImages")
    public List<Object> labelImages;
    @JsonProperty("LabelPhotoDescription")
    public String labelPhotoDescription;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("NumberOfDoses")
    public String numberOfDoses;
    @JsonProperty("PackagingImages")
    public List<Object> packagingImages;
    @JsonProperty("PackagingPhotoDescription")
    public String packagingPhotoDescription;
    @JsonProperty("ProcessingJobCategoryName")
    public String processingJobCategoryName;
    @JsonProperty("ProcessingJobTypeName")
    public String processingJobTypeName;
    @JsonProperty("ProductBrandName")
    public String productBrandName;
    @JsonProperty("ProductCategoryName")
    public String productCategoryName;
    @JsonProperty("ProductCategoryType")
    public String productCategoryType;
    @JsonProperty("ProductImages")
    public List<Object> productImages;
    @JsonProperty("ProductPDFDocuments")
    public List<Object> productPDFDocuments;
    @JsonProperty("ProductPhotoDescription")
    public String productPhotoDescription;
    @JsonProperty("PublicIngredients")
    public String publicIngredients;
    @JsonProperty("QuantityType")
    public String quantityType;
    @JsonProperty("ServingSize")
    public String servingSize;
    @JsonProperty("StrainId")
    public Integer strainId;
    @JsonProperty("StrainName")
    public String strainName;
    @JsonProperty("UnitCbdAContent")
    public Double unitCbdAContent;
    @JsonProperty("UnitCbdAContentDose")
    public Double unitCbdAContentDose;
    @JsonProperty("UnitCbdAContentDoseUnitOfMeasureName")
    public String unitCbdAContentDoseUnitOfMeasureName;
    @JsonProperty("UnitCbdAContentUnitOfMeasureName")
    public String unitCbdAContentUnitOfMeasureName;
    @JsonProperty("UnitCbdAPercent")
    public Double unitCbdAPercent;
    @JsonProperty("UnitCbdContent")
    public Double unitCbdContent;
    @JsonProperty("UnitCbdContentDose")
    public Double unitCbdContentDose;
    @JsonProperty("UnitCbdContentDoseUnitOfMeasureName")
    public String unitCbdContentDoseUnitOfMeasureName;
    @JsonProperty("UnitCbdContentUnitOfMeasureName")
    public String unitCbdContentUnitOfMeasureName;
    @JsonProperty("UnitCbdPercent")
    public Double unitCbdPercent;
    @JsonProperty("UnitOfMeasureName")
    public String unitOfMeasureName;
    @JsonProperty("UnitQuantity")
    public Double unitQuantity;
    @JsonProperty("UnitQuantityUnitOfMeasureName")
    public String unitQuantityUnitOfMeasureName;
    @JsonProperty("UnitThcAContent")
    public Double unitThcAContent;
    @JsonProperty("UnitThcAContentDose")
    public Double unitThcAContentDose;
    @JsonProperty("UnitThcAContentDoseUnitOfMeasureId")
    public Integer unitThcAContentDoseUnitOfMeasureId;
    @JsonProperty("UnitThcAContentDoseUnitOfMeasureName")
    public String unitThcAContentDoseUnitOfMeasureName;
    @JsonProperty("UnitThcAContentUnitOfMeasureName")
    public String unitThcAContentUnitOfMeasureName;
    @JsonProperty("UnitThcAPercent")
    public Double unitThcAPercent;
    @JsonProperty("UnitThcContent")
    public Double unitThcContent;
    @JsonProperty("UnitThcContentDose")
    public Double unitThcContentDose;
    @JsonProperty("UnitThcContentDoseUnitOfMeasureId")
    public Integer unitThcContentDoseUnitOfMeasureId;
    @JsonProperty("UnitThcContentDoseUnitOfMeasureName")
    public String unitThcContentDoseUnitOfMeasureName;
    @JsonProperty("UnitThcContentUnitOfMeasureName")
    public String unitThcContentUnitOfMeasureName;
    @JsonProperty("UnitThcPercent")
    public Double unitThcPercent;
    @JsonProperty("UnitVolume")
    public Double unitVolume;
    @JsonProperty("UnitVolumeUnitOfMeasureName")
    public String unitVolumeUnitOfMeasureName;
    @JsonProperty("UnitWeight")
    public Double unitWeight;
    @JsonProperty("UnitWeightUnitOfMeasureName")
    public String unitWeightUnitOfMeasureName;
}
