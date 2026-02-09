package io.github.thunkier.thunkmetrc

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import io.github.thunkier.thunkmetrc.wrapper.services.LocationsService
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope
import java.time.format.DateTimeFormatter
import java.time.OffsetDateTime

/**
 * Lazy pagination iterator for getActiveLocations.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun LocationsService.iterateGetActiveLocations(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<LocationsLocation> {
    return MetrcExtensions.iteratePages { page ->
        this.getActiveLocations(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<LocationsLocation>
    }
}
/**
 * Sync items for getActiveLocations.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun LocationsService.syncGetActiveLocations(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<LocationsLocation> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetActiveLocations(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getActiveLocations.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetActiveLocationsParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<LocationsLocation>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = LocationsService(client, limiter)
    
    svc.syncGetActiveLocations(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getInactiveLocations.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun LocationsService.iterateGetInactiveLocations(licenseNumber: String? = null, pageSize: String? = null): Flow<LocationsLocation> {
    return MetrcExtensions.iteratePages { page ->
        this.getInactiveLocations(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<LocationsLocation>
    }
}

/**
 * Lazy pagination iterator for getLocationsTypes.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun LocationsService.iterateGetLocationsTypes(licenseNumber: String? = null, pageSize: String? = null): Flow<LocationsType> {
    return MetrcExtensions.iteratePages { page ->
        this.getLocationsTypes(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<LocationsType>
    }
}

