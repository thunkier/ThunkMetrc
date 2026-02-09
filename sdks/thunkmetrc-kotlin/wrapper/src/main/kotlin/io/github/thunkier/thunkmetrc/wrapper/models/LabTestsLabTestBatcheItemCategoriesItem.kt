package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class LabTestsLabTestBatcheItemCategoriesItem(
    @JsonProperty("ProductCategory")
    val productCategory: LabTestsLabTestBatcheItemCategoriesItemProductCategory? = null,
    @JsonProperty("ProductCategoryId")
    val productCategoryId: String? = null,
    @JsonProperty("RequireLabTestBatch")
    val requireLabTestBatch: Boolean? = null
)
