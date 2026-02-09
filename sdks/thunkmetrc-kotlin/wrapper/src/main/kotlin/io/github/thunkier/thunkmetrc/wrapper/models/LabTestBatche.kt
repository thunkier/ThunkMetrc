package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class LabTestBatche(
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("ItemCategories")
    val itemCategories: List<LabTestsLabTestBatcheItemCategoriesItem>? = null,
    @JsonProperty("ItemCategoryCount")
    val itemCategoryCount: Int? = null,
    @JsonProperty("LabTestTypeCount")
    val labTestTypeCount: Int? = null,
    @JsonProperty("LabTestTypes")
    val labTestTypes: List<LabTestsLabTestBatcheLabTestTypesItem>? = null,
    @JsonProperty("Name")
    val name: String? = null,
    @JsonProperty("RequiresAllFromLabTestBatch")
    val requiresAllFromLabTestBatch: Boolean? = null
)
