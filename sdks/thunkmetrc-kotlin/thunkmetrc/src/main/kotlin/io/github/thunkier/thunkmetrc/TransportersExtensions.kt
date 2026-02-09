package io.github.thunkier.thunkmetrc

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import io.github.thunkier.thunkmetrc.wrapper.services.TransportersService
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope
import java.time.format.DateTimeFormatter
import java.time.OffsetDateTime

/**
 * Lazy pagination iterator for getDrivers.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransportersService.iterateGetDrivers(licenseNumber: String? = null, pageSize: String? = null): Flow<TransportersDriver> {
    return MetrcExtensions.iteratePages { page ->
        this.getDrivers(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<TransportersDriver>
    }
}

/**
 * Lazy pagination iterator for getVehicles.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransportersService.iterateGetVehicles(licenseNumber: String? = null, pageSize: String? = null): Flow<TransportersVehicle> {
    return MetrcExtensions.iteratePages { page ->
        this.getVehicles(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<TransportersVehicle>
    }
}

