package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class AdditivesTemplate(
    @JsonProperty("ActiveIngredients")
    val activeIngredients: List<AdditivesTemplateActiveIngredientsItem>? = null,
    @JsonProperty("AdditiveType")
    val additiveType: String? = null,
    @JsonProperty("AdditiveTypeName")
    val additiveTypeName: String? = null,
    @JsonProperty("ApplicationDevice")
    val applicationDevice: String? = null,
    @JsonProperty("EpaRegistrationNumber")
    val epaRegistrationNumber: String? = null,
    @JsonProperty("FacilityId")
    val facilityId: Int? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("IsActive")
    val isActive: Boolean? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("Note")
    val note: String? = null,
    @JsonProperty("ProductSupplier")
    val productSupplier: String? = null,
    @JsonProperty("ProductTradeName")
    val productTradeName: String? = null,
    @JsonProperty("RestrictiveEntryIntervalQuantityDescription")
    val restrictiveEntryIntervalQuantityDescription: String? = null,
    @JsonProperty("RestrictiveEntryIntervalTimeDescription")
    val restrictiveEntryIntervalTimeDescription: String? = null
)
