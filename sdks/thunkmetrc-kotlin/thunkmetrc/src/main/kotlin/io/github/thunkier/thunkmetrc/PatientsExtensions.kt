package io.github.thunkier.thunkmetrc

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import io.github.thunkier.thunkmetrc.wrapper.services.PatientsService
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope
import java.time.format.DateTimeFormatter
import java.time.OffsetDateTime

/**
 * Lazy pagination iterator for getActivePatients.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun PatientsService.iterateGetActivePatients(licenseNumber: String? = null, pageSize: String? = null): Flow<Patient> {
    return MetrcExtensions.iteratePages { page ->
        this.getActivePatients(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Patient>
    }
}

