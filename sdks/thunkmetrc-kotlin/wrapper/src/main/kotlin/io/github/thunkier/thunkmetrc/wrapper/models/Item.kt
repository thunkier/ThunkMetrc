package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Item(
    @JsonProperty("AdministrationMethod")
    val administrationMethod: String? = null,
    @JsonProperty("Allergens")
    val allergens: String? = null,
    @JsonProperty("ApprovalStatus")
    val approvalStatus: String? = null,
    @JsonProperty("ApprovalStatusDateTime")
    val approvalStatusDateTime: String? = null,
    @JsonProperty("DefaultLabTestingState")
    val defaultLabTestingState: String? = null,
    @JsonProperty("Description")
    val description: String? = null,
    @JsonProperty("GlobalProductName")
    val globalProductName: String? = null,
    @JsonProperty("GlobalProductNumber")
    val globalProductNumber: String? = null,
    @JsonProperty("HasExpirationDate")
    val hasExpirationDate: Boolean? = null,
    @JsonProperty("HasSellByDate")
    val hasSellByDate: Boolean? = null,
    @JsonProperty("HasUseByDate")
    val hasUseByDate: Boolean? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IsExpirationDateRequired")
    val isExpirationDateRequired: Boolean? = null,
    @JsonProperty("IsSellByDateRequired")
    val isSellByDateRequired: Boolean? = null,
    @JsonProperty("IsUseByDateRequired")
    val isUseByDateRequired: Boolean? = null,
    @JsonProperty("IsUsed")
    val isUsed: Boolean? = null,
    @JsonProperty("ItemBrandId")
    val itemBrandId: Int? = null,
    @JsonProperty("ItemBrandName")
    val itemBrandName: String? = null,
    @JsonProperty("LabTestBatchNames")
    val labTestBatchNames: List<Any>? = null,
    @JsonProperty("LabelImages")
    val labelImages: List<Any>? = null,
    @JsonProperty("LabelPhotoDescription")
    val labelPhotoDescription: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("NumberOfDoses")
    val numberOfDoses: String? = null,
    @JsonProperty("PackagingImages")
    val packagingImages: List<Any>? = null,
    @JsonProperty("PackagingPhotoDescription")
    val packagingPhotoDescription: String? = null,
    @JsonProperty("ProcessingJobCategoryName")
    val processingJobCategoryName: String? = null,
    @JsonProperty("ProcessingJobTypeName")
    val processingJobTypeName: String? = null,
    @JsonProperty("ProductBrandName")
    val productBrandName: String? = null,
    @JsonProperty("ProductCategoryName")
    val productCategoryName: String? = null,
    @JsonProperty("ProductCategoryType")
    val productCategoryType: String? = null,
    @JsonProperty("ProductImages")
    val productImages: List<Any>? = null,
    @JsonProperty("ProductPDFDocuments")
    val productPDFDocuments: List<Any>? = null,
    @JsonProperty("ProductPhotoDescription")
    val productPhotoDescription: String? = null,
    @JsonProperty("PublicIngredients")
    val publicIngredients: String? = null,
    @JsonProperty("QuantityType")
    val quantityType: String? = null,
    @JsonProperty("ServingSize")
    val servingSize: String? = null,
    @JsonProperty("StrainId")
    val strainId: Int? = null,
    @JsonProperty("StrainName")
    val strainName: String? = null,
    @JsonProperty("UnitCbdAContent")
    val unitCbdAContent: Double? = null,
    @JsonProperty("UnitCbdAContentDose")
    val unitCbdAContentDose: Double? = null,
    @JsonProperty("UnitCbdAContentDoseUnitOfMeasureName")
    val unitCbdAContentDoseUnitOfMeasureName: String? = null,
    @JsonProperty("UnitCbdAContentUnitOfMeasureName")
    val unitCbdAContentUnitOfMeasureName: String? = null,
    @JsonProperty("UnitCbdAPercent")
    val unitCbdAPercent: Double? = null,
    @JsonProperty("UnitCbdContent")
    val unitCbdContent: Double? = null,
    @JsonProperty("UnitCbdContentDose")
    val unitCbdContentDose: Double? = null,
    @JsonProperty("UnitCbdContentDoseUnitOfMeasureName")
    val unitCbdContentDoseUnitOfMeasureName: String? = null,
    @JsonProperty("UnitCbdContentUnitOfMeasureName")
    val unitCbdContentUnitOfMeasureName: String? = null,
    @JsonProperty("UnitCbdPercent")
    val unitCbdPercent: Double? = null,
    @JsonProperty("UnitOfMeasureName")
    val unitOfMeasureName: String? = null,
    @JsonProperty("UnitQuantity")
    val unitQuantity: Double? = null,
    @JsonProperty("UnitQuantityUnitOfMeasureName")
    val unitQuantityUnitOfMeasureName: String? = null,
    @JsonProperty("UnitThcAContent")
    val unitThcAContent: Double? = null,
    @JsonProperty("UnitThcAContentDose")
    val unitThcAContentDose: Double? = null,
    @JsonProperty("UnitThcAContentDoseUnitOfMeasureId")
    val unitThcAContentDoseUnitOfMeasureId: Int? = null,
    @JsonProperty("UnitThcAContentDoseUnitOfMeasureName")
    val unitThcAContentDoseUnitOfMeasureName: String? = null,
    @JsonProperty("UnitThcAContentUnitOfMeasureName")
    val unitThcAContentUnitOfMeasureName: String? = null,
    @JsonProperty("UnitThcAPercent")
    val unitThcAPercent: Double? = null,
    @JsonProperty("UnitThcContent")
    val unitThcContent: Double? = null,
    @JsonProperty("UnitThcContentDose")
    val unitThcContentDose: Double? = null,
    @JsonProperty("UnitThcContentDoseUnitOfMeasureId")
    val unitThcContentDoseUnitOfMeasureId: Int? = null,
    @JsonProperty("UnitThcContentDoseUnitOfMeasureName")
    val unitThcContentDoseUnitOfMeasureName: String? = null,
    @JsonProperty("UnitThcContentUnitOfMeasureName")
    val unitThcContentUnitOfMeasureName: String? = null,
    @JsonProperty("UnitThcPercent")
    val unitThcPercent: Double? = null,
    @JsonProperty("UnitVolume")
    val unitVolume: Double? = null,
    @JsonProperty("UnitVolumeUnitOfMeasureName")
    val unitVolumeUnitOfMeasureName: String? = null,
    @JsonProperty("UnitWeight")
    val unitWeight: Double? = null,
    @JsonProperty("UnitWeightUnitOfMeasureName")
    val unitWeightUnitOfMeasureName: String? = null
)
