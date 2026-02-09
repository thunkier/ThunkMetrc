package io.github.thunkier.thunkmetrc

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import io.github.thunkier.thunkmetrc.wrapper.services.HarvestsService
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope
import java.time.format.DateTimeFormatter
import java.time.OffsetDateTime

/**
 * Lazy pagination iterator for getActiveHarvests.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun HarvestsService.iterateGetActiveHarvests(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<Harvest> {
    return MetrcExtensions.iteratePages { page ->
        this.getActiveHarvests(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Harvest>
    }
}
/**
 * Sync items for getActiveHarvests.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun HarvestsService.syncGetActiveHarvests(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<Harvest> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetActiveHarvests(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getActiveHarvests.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetActiveHarvestsParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<Harvest>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = HarvestsService(client, limiter)
    
    svc.syncGetActiveHarvests(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getHarvestsWaste.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun HarvestsService.iterateGetHarvestsWaste(harvestId: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<HarvestsWaste> {
    return MetrcExtensions.iteratePages { page ->
        this.getHarvestsWaste(
                    harvestId,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<HarvestsWaste>
    }
}

/**
 * Lazy pagination iterator for getInactiveHarvests.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun HarvestsService.iterateGetInactiveHarvests(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<Harvest> {
    return MetrcExtensions.iteratePages { page ->
        this.getInactiveHarvests(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Harvest>
    }
}
/**
 * Sync items for getInactiveHarvests.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun HarvestsService.syncGetInactiveHarvests(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<Harvest> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetInactiveHarvests(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getInactiveHarvests.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetInactiveHarvestsParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<Harvest>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = HarvestsService(client, limiter)
    
    svc.syncGetInactiveHarvests(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getOnHoldHarvests.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun HarvestsService.iterateGetOnHoldHarvests(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<Harvest> {
    return MetrcExtensions.iteratePages { page ->
        this.getOnHoldHarvests(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Harvest>
    }
}
/**
 * Sync items for getOnHoldHarvests.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun HarvestsService.syncGetOnHoldHarvests(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<Harvest> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetOnHoldHarvests(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getOnHoldHarvests.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetOnHoldHarvestsParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<Harvest>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = HarvestsService(client, limiter)
    
    svc.syncGetOnHoldHarvests(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getWasteTypes.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun HarvestsService.iterateGetWasteTypes(pageSize: String? = null): Flow<WasteType> {
    return MetrcExtensions.iteratePages { page ->
        this.getWasteTypes(
                    page.toString(),
                    pageSize,) as? PaginatedResponse<WasteType>
    }
}

