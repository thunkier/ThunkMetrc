package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class LabTestsBatchItemCategoriesItem(
    @JsonProperty("ProductCategory")
    val productCategory: LabTestsBatchItemCategoriesItemProductCategory? = null,
    @JsonProperty("ProductCategoryId")
    val productCategoryId: String? = null,
    @JsonProperty("RequireLabTestBatch")
    val requireLabTestBatch: Boolean? = null
)
