package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class ItemCategory {
    @JsonProperty("CanBeDecontaminated")
    public Boolean canBeDecontaminated;
    @JsonProperty("CanBeDestroyed")
    public Boolean canBeDestroyed;
    @JsonProperty("CanBeRemediated")
    public Boolean canBeRemediated;
    @JsonProperty("CanContainSeeds")
    public Boolean canContainSeeds;
    @JsonProperty("LabTestBatchNames")
    public List<Object> labTestBatchNames;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("ProductCategoryType")
    public String productCategoryType;
    @JsonProperty("QuantityType")
    public String quantityType;
    @JsonProperty("RequiresAdministrationMethod")
    public Boolean requiresAdministrationMethod;
    @JsonProperty("RequiresAllergens")
    public Boolean requiresAllergens;
    @JsonProperty("RequiresDescription")
    public Boolean requiresDescription;
    @JsonProperty("RequiresItemBrand")
    public Boolean requiresItemBrand;
    @JsonProperty("RequiresLabelPhotoDescription")
    public Boolean requiresLabelPhotoDescription;
    @JsonProperty("RequiresLabelPhotos")
    public Integer requiresLabelPhotos;
    @JsonProperty("RequiresNumberOfDoses")
    public Boolean requiresNumberOfDoses;
    @JsonProperty("RequiresPackagingPhotoDescription")
    public Boolean requiresPackagingPhotoDescription;
    @JsonProperty("RequiresPackagingPhotos")
    public Integer requiresPackagingPhotos;
    @JsonProperty("RequiresProductPDFDocuments")
    public Integer requiresProductPDFDocuments;
    @JsonProperty("RequiresProductPhotoDescription")
    public Boolean requiresProductPhotoDescription;
    @JsonProperty("RequiresProductPhotos")
    public Integer requiresProductPhotos;
    @JsonProperty("RequiresPublicIngredients")
    public Boolean requiresPublicIngredients;
    @JsonProperty("RequiresServingSize")
    public Boolean requiresServingSize;
    @JsonProperty("RequiresStrain")
    public Boolean requiresStrain;
    @JsonProperty("RequiresSupplyDurationDays")
    public Boolean requiresSupplyDurationDays;
    @JsonProperty("RequiresUnitCbdAContent")
    public Boolean requiresUnitCbdAContent;
    @JsonProperty("RequiresUnitCbdAContentDose")
    public Boolean requiresUnitCbdAContentDose;
    @JsonProperty("RequiresUnitCbdAPercent")
    public Boolean requiresUnitCbdAPercent;
    @JsonProperty("RequiresUnitCbdContent")
    public Boolean requiresUnitCbdContent;
    @JsonProperty("RequiresUnitCbdContentDose")
    public Boolean requiresUnitCbdContentDose;
    @JsonProperty("RequiresUnitCbdPercent")
    public Boolean requiresUnitCbdPercent;
    @JsonProperty("RequiresUnitThcAContent")
    public Boolean requiresUnitThcAContent;
    @JsonProperty("RequiresUnitThcAContentDose")
    public Boolean requiresUnitThcAContentDose;
    @JsonProperty("RequiresUnitThcAPercent")
    public Boolean requiresUnitThcAPercent;
    @JsonProperty("RequiresUnitThcContent")
    public Boolean requiresUnitThcContent;
    @JsonProperty("RequiresUnitThcContentDose")
    public Boolean requiresUnitThcContentDose;
    @JsonProperty("RequiresUnitThcPercent")
    public Boolean requiresUnitThcPercent;
    @JsonProperty("RequiresUnitVolume")
    public Boolean requiresUnitVolume;
    @JsonProperty("RequiresUnitWeight")
    public Boolean requiresUnitWeight;
}
