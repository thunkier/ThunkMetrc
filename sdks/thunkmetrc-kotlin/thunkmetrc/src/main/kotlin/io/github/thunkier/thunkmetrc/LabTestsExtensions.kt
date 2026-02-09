package io.github.thunkier.thunkmetrc

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import io.github.thunkier.thunkmetrc.wrapper.services.LabTestsService
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope
import java.time.format.DateTimeFormatter
import java.time.OffsetDateTime

/**
 * Lazy pagination iterator for getBatches.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun LabTestsService.iterateGetBatches(pageSize: String? = null): Flow<Batch> {
    return MetrcExtensions.iteratePages { page ->
        this.getBatches(
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Batch>
    }
}

/**
 * Lazy pagination iterator for getLabTestsTypes.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun LabTestsService.iterateGetLabTestsTypes(pageSize: String? = null): Flow<LabTestsType> {
    return MetrcExtensions.iteratePages { page ->
        this.getLabTestsTypes(
                    page.toString(),
                    pageSize,) as? PaginatedResponse<LabTestsType>
    }
}

/**
 * Lazy pagination iterator for getResults.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun LabTestsService.iterateGetResults(licenseNumber: String? = null, packageId: String? = null, pageSize: String? = null): Flow<Result> {
    return MetrcExtensions.iteratePages { page ->
        this.getResults(
                    licenseNumber,
                    packageId,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Result>
    }
}

