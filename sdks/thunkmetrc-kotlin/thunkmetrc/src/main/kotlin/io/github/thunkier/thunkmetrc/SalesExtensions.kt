package io.github.thunkier.thunkmetrc

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import io.github.thunkier.thunkmetrc.wrapper.services.SalesService
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope
import java.time.format.DateTimeFormatter
import java.time.OffsetDateTime

/**
 * Lazy pagination iterator for getActiveDeliveries.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun SalesService.iterateGetActiveDeliveries(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<ActiveDelivery> {
    return MetrcExtensions.iteratePages { page ->
        this.getActiveDeliveries(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<ActiveDelivery>
    }
}
/**
 * Sync items for getActiveDeliveries.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun SalesService.syncGetActiveDeliveries(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<ActiveDelivery> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetActiveDeliveries(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getActiveDeliveries.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetActiveDeliveriesParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<ActiveDelivery>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = SalesService(client, limiter)
    
    svc.syncGetActiveDeliveries(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getActiveDeliveriesRetailer.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun SalesService.iterateGetActiveDeliveriesRetailer(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<ActiveDeliveriesRetailer> {
    return MetrcExtensions.iteratePages { page ->
        this.getActiveDeliveriesRetailer(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<ActiveDeliveriesRetailer>
    }
}
/**
 * Sync items for getActiveDeliveriesRetailer.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun SalesService.syncGetActiveDeliveriesRetailer(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<ActiveDeliveriesRetailer> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetActiveDeliveriesRetailer(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getActiveDeliveriesRetailer.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetActiveDeliveriesRetailerParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<ActiveDeliveriesRetailer>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = SalesService(client, limiter)
    
    svc.syncGetActiveDeliveriesRetailer(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getActiveReceipts.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun SalesService.iterateGetActiveReceipts(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<ActiveReceipt> {
    return MetrcExtensions.iteratePages { page ->
        this.getActiveReceipts(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<ActiveReceipt>
    }
}
/**
 * Sync items for getActiveReceipts.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun SalesService.syncGetActiveReceipts(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<ActiveReceipt> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetActiveReceipts(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getActiveReceipts.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetActiveReceiptsParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<ActiveReceipt>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = SalesService(client, limiter)
    
    svc.syncGetActiveReceipts(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getDeliveriesReturnReasons.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun SalesService.iterateGetDeliveriesReturnReasons(licenseNumber: String? = null, pageSize: String? = null): Flow<DeliveriesReturnReason> {
    return MetrcExtensions.iteratePages { page ->
        this.getDeliveriesReturnReasons(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<DeliveriesReturnReason>
    }
}

/**
 * Lazy pagination iterator for getInactiveDeliveries.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun SalesService.iterateGetInactiveDeliveries(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<InactiveDelivery> {
    return MetrcExtensions.iteratePages { page ->
        this.getInactiveDeliveries(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<InactiveDelivery>
    }
}
/**
 * Sync items for getInactiveDeliveries.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun SalesService.syncGetInactiveDeliveries(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<InactiveDelivery> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetInactiveDeliveries(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getInactiveDeliveries.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetInactiveDeliveriesParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<InactiveDelivery>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = SalesService(client, limiter)
    
    svc.syncGetInactiveDeliveries(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getInactiveDeliveriesRetailer.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun SalesService.iterateGetInactiveDeliveriesRetailer(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<InactiveDeliveriesRetailer> {
    return MetrcExtensions.iteratePages { page ->
        this.getInactiveDeliveriesRetailer(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<InactiveDeliveriesRetailer>
    }
}
/**
 * Sync items for getInactiveDeliveriesRetailer.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun SalesService.syncGetInactiveDeliveriesRetailer(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<InactiveDeliveriesRetailer> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetInactiveDeliveriesRetailer(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getInactiveDeliveriesRetailer.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetInactiveDeliveriesRetailerParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<InactiveDeliveriesRetailer>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = SalesService(client, limiter)
    
    svc.syncGetInactiveDeliveriesRetailer(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getInactiveReceipts.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun SalesService.iterateGetInactiveReceipts(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<InactiveReceipt> {
    return MetrcExtensions.iteratePages { page ->
        this.getInactiveReceipts(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<InactiveReceipt>
    }
}
/**
 * Sync items for getInactiveReceipts.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun SalesService.syncGetInactiveReceipts(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<InactiveReceipt> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetInactiveReceipts(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getInactiveReceipts.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetInactiveReceiptsParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<InactiveReceipt>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = SalesService(client, limiter)
    
    svc.syncGetInactiveReceipts(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

