package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class StartProcessingJobRequest(
    @JsonProperty("CountUnitOfMeasure")
    val countUnitOfMeasure: String? = null,
    @JsonProperty("JobName")
    val jobName: String? = null,
    @JsonProperty("JobType")
    val jobType: String? = null,
    @JsonProperty("Packages")
    val packages: List<Any>? = null,
    @JsonProperty("StartDate")
    val startDate: String? = null,
    @JsonProperty("VolumeUnitOfMeasure")
    val volumeUnitOfMeasure: String? = null,
    @JsonProperty("WeightUnitOfMeasure")
    val weightUnitOfMeasure: String? = null
)
