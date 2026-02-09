package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class SalesUpdateDeliveriesRequestItemTransactionsItem(
    @JsonProperty("CityTax")
    val cityTax: String? = null,
    @JsonProperty("CountyTax")
    val countyTax: String? = null,
    @JsonProperty("DiscountAmount")
    val discountAmount: String? = null,
    @JsonProperty("ExciseTax")
    val exciseTax: String? = null,
    @JsonProperty("InvoiceNumber")
    val invoiceNumber: String? = null,
    @JsonProperty("MunicipalTax")
    val municipalTax: String? = null,
    @JsonProperty("PackageLabel")
    val packageLabel: String? = null,
    @JsonProperty("Price")
    val price: String? = null,
    @JsonProperty("QrCodes")
    val qrCodes: String? = null,
    @JsonProperty("Quantity")
    val quantity: Double? = null,
    @JsonProperty("SalesTax")
    val salesTax: String? = null,
    @JsonProperty("SubTotal")
    val subTotal: String? = null,
    @JsonProperty("TotalAmount")
    val totalAmount: Double? = null,
    @JsonProperty("UnitOfMeasure")
    val unitOfMeasure: String? = null,
    @JsonProperty("UnitThcContent")
    val unitThcContent: Double? = null,
    @JsonProperty("UnitThcContentUnitOfMeasure")
    val unitThcContentUnitOfMeasure: String? = null,
    @JsonProperty("UnitThcPercent")
    val unitThcPercent: Double? = null,
    @JsonProperty("UnitWeight")
    val unitWeight: Double? = null,
    @JsonProperty("UnitWeightUnitOfMeasure")
    val unitWeightUnitOfMeasure: String? = null
)
