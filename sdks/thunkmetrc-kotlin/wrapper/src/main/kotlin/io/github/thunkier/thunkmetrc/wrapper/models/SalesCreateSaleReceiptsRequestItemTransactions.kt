package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesCreateSaleReceiptsRequestItemTransactions(
    @JsonProperty("cityTax")
    val cityTax: String? = null,
    @JsonProperty("countyTax")
    val countyTax: String? = null,
    @JsonProperty("discountAmount")
    val discountAmount: String? = null,
    @JsonProperty("exciseTax")
    val exciseTax: String? = null,
    @JsonProperty("invoiceNumber")
    val invoiceNumber: String? = null,
    @JsonProperty("municipalTax")
    val municipalTax: String? = null,
    @JsonProperty("packageLabel")
    val packageLabel: String? = null,
    @JsonProperty("price")
    val price: String? = null,
    @JsonProperty("qrCodes")
    val qrCodes: String? = null,
    @JsonProperty("quantity")
    val quantity: Int? = null,
    @JsonProperty("salesTax")
    val salesTax: String? = null,
    @JsonProperty("subTotal")
    val subTotal: String? = null,
    @JsonProperty("totalAmount")
    val totalAmount: Double? = null,
    @JsonProperty("unitOfMeasure")
    val unitOfMeasure: String? = null,
    @JsonProperty("unitThcContent")
    val unitThcContent: Double? = null,
    @JsonProperty("unitThcContentUnitOfMeasure")
    val unitThcContentUnitOfMeasure: String? = null,
    @JsonProperty("unitThcPercent")
    val unitThcPercent: Double? = null,
    @JsonProperty("unitWeight")
    val unitWeight: Double? = null,
    @JsonProperty("unitWeightUnitOfMeasure")
    val unitWeightUnitOfMeasure: String? = null
)
