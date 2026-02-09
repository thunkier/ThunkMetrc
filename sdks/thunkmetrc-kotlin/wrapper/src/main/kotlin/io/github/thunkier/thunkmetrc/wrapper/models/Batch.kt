package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class Batch(
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("ItemCategories")
    val itemCategories: List<LabTestsBatchItemCategoriesItem>? = null,
    @JsonProperty("ItemCategoryCount")
    val itemCategoryCount: Int? = null,
    @JsonProperty("LabTestTypeCount")
    val labTestTypeCount: Int? = null,
    @JsonProperty("LabTestTypes")
    val labTestTypes: List<LabTestsBatchLabTestTypesItem>? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("RequiresAllFromLabTestBatch")
    val requiresAllFromLabTestBatch: Boolean? = null
)
