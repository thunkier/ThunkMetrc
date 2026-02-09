package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class DeliveryPackageRequiredlabtestbatch(
    @JsonProperty("LabTestBatchId")
    val labTestBatchId: Int? = null,
    @JsonProperty("LabTestBatchName")
    val labTestBatchName: String? = null,
    @JsonProperty("PackageId")
    val packageId: Int? = null
)
