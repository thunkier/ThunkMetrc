package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class CreateTestingRequestIngredientsItem(
    @JsonProperty("Package")
    val package_: String? = null,
    @JsonProperty("Quantity")
    val quantity: Int? = null,
    @JsonProperty("UnitOfMeasure")
    val unitOfMeasure: String? = null
)
