package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ItemsCreateItemsRequestItem(
    @JsonProperty("AdministrationMethod")
    val administrationMethod: String? = null,
    @JsonProperty("Allergens")
    val allergens: String? = null,
    @JsonProperty("Description")
    val description: String? = null,
    @JsonProperty("GlobalProductName")
    val globalProductName: String? = null,
    @JsonProperty("ItemBrand")
    val itemBrand: String? = null,
    @JsonProperty("ItemCategory")
    val itemCategory: String? = null,
    @JsonProperty("ItemIngredients")
    val itemIngredients: String? = null,
    @JsonProperty("LabelImageFileSystemIds")
    val labelImageFileSystemIds: String? = null,
    @JsonProperty("LabelPhotoDescription")
    val labelPhotoDescription: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("NumberOfDoses")
    val numberOfDoses: String? = null,
    @JsonProperty("PackagingImageFileSystemIds")
    val packagingImageFileSystemIds: String? = null,
    @JsonProperty("PackagingPhotoDescription")
    val packagingPhotoDescription: String? = null,
    @JsonProperty("ProcessingJobCategoryName")
    val processingJobCategoryName: String? = null,
    @JsonProperty("ProcessingJobTypeName")
    val processingJobTypeName: String? = null,
    @JsonProperty("ProductImageFileSystemIds")
    val productImageFileSystemIds: String? = null,
    @JsonProperty("ProductPDFFileSystemIds")
    val productPDFFileSystemIds: String? = null,
    @JsonProperty("ProductPhotoDescription")
    val productPhotoDescription: String? = null,
    @JsonProperty("PublicIngredients")
    val publicIngredients: String? = null,
    @JsonProperty("ServingSize")
    val servingSize: String? = null,
    @JsonProperty("Strain")
    val strain: String? = null,
    @JsonProperty("SupplyDurationDays")
    val supplyDurationDays: Int? = null,
    @JsonProperty("UnitCbdAContent")
    val unitCbdAContent: Double? = null,
    @JsonProperty("UnitCbdAContentDose")
    val unitCbdAContentDose: Double? = null,
    @JsonProperty("UnitCbdAContentDoseUnitOfMeasure")
    val unitCbdAContentDoseUnitOfMeasure: String? = null,
    @JsonProperty("UnitCbdAContentUnitOfMeasure")
    val unitCbdAContentUnitOfMeasure: String? = null,
    @JsonProperty("UnitCbdAPercent")
    val unitCbdAPercent: Double? = null,
    @JsonProperty("UnitCbdContent")
    val unitCbdContent: Double? = null,
    @JsonProperty("UnitCbdContentDose")
    val unitCbdContentDose: Double? = null,
    @JsonProperty("UnitCbdContentDoseUnitOfMeasure")
    val unitCbdContentDoseUnitOfMeasure: String? = null,
    @JsonProperty("UnitCbdContentUnitOfMeasure")
    val unitCbdContentUnitOfMeasure: String? = null,
    @JsonProperty("UnitCbdPercent")
    val unitCbdPercent: Double? = null,
    @JsonProperty("UnitOfMeasure")
    val unitOfMeasure: String? = null,
    @JsonProperty("UnitThcAContent")
    val unitThcAContent: Double? = null,
    @JsonProperty("UnitThcAContentDose")
    val unitThcAContentDose: Double? = null,
    @JsonProperty("UnitThcAContentDoseUnitOfMeasure")
    val unitThcAContentDoseUnitOfMeasure: String? = null,
    @JsonProperty("UnitThcAContentUnitOfMeasure")
    val unitThcAContentUnitOfMeasure: String? = null,
    @JsonProperty("UnitThcAPercent")
    val unitThcAPercent: Double? = null,
    @JsonProperty("UnitThcContent")
    val unitThcContent: Double? = null,
    @JsonProperty("UnitThcContentDose")
    val unitThcContentDose: Double? = null,
    @JsonProperty("UnitThcContentDoseUnitOfMeasure")
    val unitThcContentDoseUnitOfMeasure: String? = null,
    @JsonProperty("UnitThcContentUnitOfMeasure")
    val unitThcContentUnitOfMeasure: String? = null,
    @JsonProperty("UnitThcPercent")
    val unitThcPercent: Double? = null,
    @JsonProperty("UnitVolume")
    val unitVolume: Double? = null,
    @JsonProperty("UnitVolumeUnitOfMeasure")
    val unitVolumeUnitOfMeasure: String? = null,
    @JsonProperty("UnitWeight")
    val unitWeight: Double? = null,
    @JsonProperty("UnitWeightUnitOfMeasure")
    val unitWeightUnitOfMeasure: String? = null
)
