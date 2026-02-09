package io.github.thunkier.thunkmetrc

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import io.github.thunkier.thunkmetrc.wrapper.services.PackagesService
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope
import java.time.format.DateTimeFormatter
import java.time.OffsetDateTime

/**
 * Lazy pagination iterator for getActivePackages.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun PackagesService.iterateGetActivePackages(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<PackagesPackage> {
    return MetrcExtensions.iteratePages { page ->
        this.getActivePackages(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<PackagesPackage>
    }
}
/**
 * Sync items for getActivePackages.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun PackagesService.syncGetActivePackages(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<PackagesPackage> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetActivePackages(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getActivePackages.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetActivePackagesParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<PackagesPackage>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = PackagesService(client, limiter)
    
    svc.syncGetActivePackages(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getAdjustReasons.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun PackagesService.iterateGetAdjustReasons(licenseNumber: String? = null, pageSize: String? = null): Flow<AdjustReason> {
    return MetrcExtensions.iteratePages { page ->
        this.getAdjustReasons(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<AdjustReason>
    }
}

/**
 * Lazy pagination iterator for getAdjustments.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun PackagesService.iterateGetAdjustments(licenseNumber: String? = null, pageSize: String? = null): Flow<Adjustment> {
    return MetrcExtensions.iteratePages { page ->
        this.getAdjustments(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Adjustment>
    }
}

/**
 * Lazy pagination iterator for getInTransitPackages.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun PackagesService.iterateGetInTransitPackages(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<InTransit> {
    return MetrcExtensions.iteratePages { page ->
        this.getInTransitPackages(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<InTransit>
    }
}
/**
 * Sync items for getInTransitPackages.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun PackagesService.syncGetInTransitPackages(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<InTransit> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetInTransitPackages(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getInTransitPackages.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetInTransitPackagesParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<InTransit>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = PackagesService(client, limiter)
    
    svc.syncGetInTransitPackages(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getInactivePackages.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun PackagesService.iterateGetInactivePackages(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<PackagesPackage> {
    return MetrcExtensions.iteratePages { page ->
        this.getInactivePackages(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<PackagesPackage>
    }
}
/**
 * Sync items for getInactivePackages.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun PackagesService.syncGetInactivePackages(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<PackagesPackage> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetInactivePackages(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getInactivePackages.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetInactivePackagesParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<PackagesPackage>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = PackagesService(client, limiter)
    
    svc.syncGetInactivePackages(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getLabSamples.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun PackagesService.iterateGetLabSamples(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<PackagesPackage> {
    return MetrcExtensions.iteratePages { page ->
        this.getLabSamples(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<PackagesPackage>
    }
}
/**
 * Sync items for getLabSamples.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun PackagesService.syncGetLabSamples(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<PackagesPackage> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetLabSamples(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getLabSamples.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetLabSamplesParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<PackagesPackage>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = PackagesService(client, limiter)
    
    svc.syncGetLabSamples(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getOnHoldPackages.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun PackagesService.iterateGetOnHoldPackages(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<PackagesPackage> {
    return MetrcExtensions.iteratePages { page ->
        this.getOnHoldPackages(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<PackagesPackage>
    }
}
/**
 * Sync items for getOnHoldPackages.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun PackagesService.syncGetOnHoldPackages(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<PackagesPackage> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetOnHoldPackages(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getOnHoldPackages.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetOnHoldPackagesParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<PackagesPackage>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = PackagesService(client, limiter)
    
    svc.syncGetOnHoldPackages(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getTransferred.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun PackagesService.iterateGetTransferred(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<PackagesPackage> {
    return MetrcExtensions.iteratePages { page ->
        this.getTransferred(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<PackagesPackage>
    }
}
/**
 * Sync items for getTransferred.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun PackagesService.syncGetTransferred(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<PackagesPackage> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetTransferred(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getTransferred.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetTransferredParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<PackagesPackage>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = PackagesService(client, limiter)
    
    svc.syncGetTransferred(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

