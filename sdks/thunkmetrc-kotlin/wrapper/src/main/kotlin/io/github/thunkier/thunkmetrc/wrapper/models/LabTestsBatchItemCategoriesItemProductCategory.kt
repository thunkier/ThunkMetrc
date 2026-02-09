package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class LabTestsBatchItemCategoriesItemProductCategory(
    @JsonProperty("CanBeDecontaminated")
    val canBeDecontaminated: Boolean? = null,
    @JsonProperty("CanBeDestroyed")
    val canBeDestroyed: Boolean? = null,
    @JsonProperty("CanBeRemediated")
    val canBeRemediated: Boolean? = null,
    @JsonProperty("CanContainSeeds")
    val canContainSeeds: Boolean? = null,
    @JsonProperty("LabTestBatchNames")
    val labTestBatchNames: List<Any>? = null,
    @JsonProperty("MaxDaysForRecoveryAfterFailure")
    val maxDaysForRecoveryAfterFailure: String? = null,
    @JsonProperty("MaxDaysToMaintainRTA")
    val maxDaysToMaintainRTA: String? = null,
    @JsonProperty("MaxDaysWithoutTestedSource")
    val maxDaysWithoutTestedSource: String? = null,
    @JsonProperty("MaxDaystoReachMinSources")
    val maxDaystoReachMinSources: String? = null,
    @JsonProperty("MaxRepackageDistanceFromSource")
    val maxRepackageDistanceFromSource: String? = null,
    @JsonProperty("MinConsecutiveSourcesForRTA")
    val minConsecutiveSourcesForRTA: String? = null,
    @JsonProperty("MinDaysForRTA")
    val minDaysForRTA: String? = null,
    @JsonProperty("MinNewPassesInRecoveryWindow")
    val minNewPassesInRecoveryWindow: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("ProductCategoryType")
    val productCategoryType: Int? = null,
    @JsonProperty("QuantityType")
    val quantityType: Int? = null,
    @JsonProperty("RequiresAdministrationMethod")
    val requiresAdministrationMethod: Boolean? = null,
    @JsonProperty("RequiresAllergens")
    val requiresAllergens: Boolean? = null,
    @JsonProperty("RequiresDescription")
    val requiresDescription: Boolean? = null,
    @JsonProperty("RequiresItemBrand")
    val requiresItemBrand: Boolean? = null,
    @JsonProperty("RequiresLabelPhotoDescription")
    val requiresLabelPhotoDescription: Boolean? = null,
    @JsonProperty("RequiresLabelPhotos")
    val requiresLabelPhotos: Int? = null,
    @JsonProperty("RequiresNumberOfDoses")
    val requiresNumberOfDoses: Boolean? = null,
    @JsonProperty("RequiresPackagingPhotoDescription")
    val requiresPackagingPhotoDescription: Boolean? = null,
    @JsonProperty("RequiresPackagingPhotos")
    val requiresPackagingPhotos: Int? = null,
    @JsonProperty("RequiresProductPDFDocuments")
    val requiresProductPDFDocuments: Int? = null,
    @JsonProperty("RequiresProductPhotoDescription")
    val requiresProductPhotoDescription: Boolean? = null,
    @JsonProperty("RequiresProductPhotos")
    val requiresProductPhotos: Int? = null,
    @JsonProperty("RequiresPublicIngredients")
    val requiresPublicIngredients: Boolean? = null,
    @JsonProperty("RequiresServingSize")
    val requiresServingSize: Boolean? = null,
    @JsonProperty("RequiresStrain")
    val requiresStrain: Boolean? = null,
    @JsonProperty("RequiresSupplyDurationDays")
    val requiresSupplyDurationDays: Boolean? = null,
    @JsonProperty("RequiresUnitCbdAContent")
    val requiresUnitCbdAContent: Boolean? = null,
    @JsonProperty("RequiresUnitCbdAContentDose")
    val requiresUnitCbdAContentDose: Boolean? = null,
    @JsonProperty("RequiresUnitCbdAPercent")
    val requiresUnitCbdAPercent: Boolean? = null,
    @JsonProperty("RequiresUnitCbdContent")
    val requiresUnitCbdContent: Boolean? = null,
    @JsonProperty("RequiresUnitCbdContentDose")
    val requiresUnitCbdContentDose: Boolean? = null,
    @JsonProperty("RequiresUnitCbdPercent")
    val requiresUnitCbdPercent: Boolean? = null,
    @JsonProperty("RequiresUnitThcAContent")
    val requiresUnitThcAContent: Boolean? = null,
    @JsonProperty("RequiresUnitThcAContentDose")
    val requiresUnitThcAContentDose: Boolean? = null,
    @JsonProperty("RequiresUnitThcAPercent")
    val requiresUnitThcAPercent: Boolean? = null,
    @JsonProperty("RequiresUnitThcContent")
    val requiresUnitThcContent: Boolean? = null,
    @JsonProperty("RequiresUnitThcContentDose")
    val requiresUnitThcContentDose: Boolean? = null,
    @JsonProperty("RequiresUnitThcPercent")
    val requiresUnitThcPercent: Boolean? = null,
    @JsonProperty("RequiresUnitVolume")
    val requiresUnitVolume: Boolean? = null,
    @JsonProperty("RequiresUnitWeight")
    val requiresUnitWeight: Boolean? = null
)
