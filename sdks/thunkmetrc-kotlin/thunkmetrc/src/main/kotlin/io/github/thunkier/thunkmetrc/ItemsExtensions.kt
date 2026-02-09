package io.github.thunkier.thunkmetrc

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import io.github.thunkier.thunkmetrc.wrapper.services.ItemsService
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope
import java.time.format.DateTimeFormatter
import java.time.OffsetDateTime

/**
 * Lazy pagination iterator for getActiveItems.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun ItemsService.iterateGetActiveItems(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<Item> {
    return MetrcExtensions.iteratePages { page ->
        this.getActiveItems(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Item>
    }
}
/**
 * Sync items for getActiveItems.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun ItemsService.syncGetActiveItems(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<Item> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetActiveItems(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getActiveItems.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetActiveItemsParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<Item>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = ItemsService(client, limiter)
    
    svc.syncGetActiveItems(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getBrands.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun ItemsService.iterateGetBrands(licenseNumber: String? = null, pageSize: String? = null): Flow<Brand> {
    return MetrcExtensions.iteratePages { page ->
        this.getBrands(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Brand>
    }
}

/**
 * Lazy pagination iterator for getCategories.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun ItemsService.iterateGetCategories(licenseNumber: String? = null, pageSize: String? = null): Flow<Category> {
    return MetrcExtensions.iteratePages { page ->
        this.getCategories(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Category>
    }
}

/**
 * Lazy pagination iterator for getInactiveItems.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun ItemsService.iterateGetInactiveItems(licenseNumber: String? = null, pageSize: String? = null): Flow<Item> {
    return MetrcExtensions.iteratePages { page ->
        this.getInactiveItems(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Item>
    }
}

