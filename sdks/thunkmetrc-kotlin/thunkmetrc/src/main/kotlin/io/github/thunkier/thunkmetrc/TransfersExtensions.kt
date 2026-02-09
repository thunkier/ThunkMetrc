package io.github.thunkier.thunkmetrc

import io.github.thunkier.thunkmetrc.client.MetrcClient
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions
import io.github.thunkier.thunkmetrc.wrapper.models.*
import io.github.thunkier.thunkmetrc.wrapper.services.TransfersService
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.toList
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope
import java.time.format.DateTimeFormatter
import java.time.OffsetDateTime

/**
 * Lazy pagination iterator for getDeliveryPackageById.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetDeliveryPackageById(id: String, pageSize: String? = null): Flow<DeliveryPackage> {
    return MetrcExtensions.iteratePages { page ->
        this.getDeliveryPackageById(id, 
                    page.toString(),
                    pageSize,) as? PaginatedResponse<DeliveryPackage>
    }
}

/**
 * Lazy pagination iterator for getDeliveryPackageRequiredlabtestbatchById.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetDeliveryPackageRequiredlabtestbatchById(id: String, pageSize: String? = null): Flow<DeliveryPackageRequiredlabtestbatch> {
    return MetrcExtensions.iteratePages { page ->
        this.getDeliveryPackageRequiredlabtestbatchById(id, 
                    page.toString(),
                    pageSize,) as? PaginatedResponse<DeliveryPackageRequiredlabtestbatch>
    }
}

/**
 * Lazy pagination iterator for getDeliveryPackageWholesaleById.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetDeliveryPackageWholesaleById(id: String, pageSize: String? = null): Flow<DeliveryPackageWholesale> {
    return MetrcExtensions.iteratePages { page ->
        this.getDeliveryPackageWholesaleById(id, 
                    page.toString(),
                    pageSize,) as? PaginatedResponse<DeliveryPackageWholesale>
    }
}

/**
 * Lazy pagination iterator for getDeliveryTransporterById.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetDeliveryTransporterById(id: String, pageSize: String? = null): Flow<DeliveryTransporter> {
    return MetrcExtensions.iteratePages { page ->
        this.getDeliveryTransporterById(id, 
                    page.toString(),
                    pageSize,) as? PaginatedResponse<DeliveryTransporter>
    }
}

/**
 * Lazy pagination iterator for getDeliveryTransporterDetailById.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetDeliveryTransporterDetailById(id: String, pageSize: String? = null): Flow<DeliveryTransporterDetail> {
    return MetrcExtensions.iteratePages { page ->
        this.getDeliveryTransporterDetailById(id, 
                    page.toString(),
                    pageSize,) as? PaginatedResponse<DeliveryTransporterDetail>
    }
}

/**
 * Lazy pagination iterator for getHub.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetHub(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<Hub> {
    return MetrcExtensions.iteratePages { page ->
        this.getHub(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Hub>
    }
}
/**
 * Sync items for getHub.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun TransfersService.syncGetHub(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<Hub> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetHub(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getHub.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetHubParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<Hub>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = TransfersService(client, limiter)
    
    svc.syncGetHub(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getIncomingTransfers.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetIncomingTransfers(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<Transfer> {
    return MetrcExtensions.iteratePages { page ->
        this.getIncomingTransfers(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Transfer>
    }
}
/**
 * Sync items for getIncomingTransfers.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun TransfersService.syncGetIncomingTransfers(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<Transfer> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetIncomingTransfers(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getIncomingTransfers.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetIncomingTransfersParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<Transfer>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = TransfersService(client, limiter)
    
    svc.syncGetIncomingTransfers(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getOutgoingTemplateDeliveryById.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetOutgoingTemplateDeliveryById(id: String, pageSize: String? = null): Flow<TemplateDelivery> {
    return MetrcExtensions.iteratePages { page ->
        this.getOutgoingTemplateDeliveryById(id, 
                    page.toString(),
                    pageSize,) as? PaginatedResponse<TemplateDelivery>
    }
}

/**
 * Lazy pagination iterator for getOutgoingTemplateDeliveryPackageById.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetOutgoingTemplateDeliveryPackageById(id: String, pageSize: String? = null): Flow<TemplateDeliveryPackage> {
    return MetrcExtensions.iteratePages { page ->
        this.getOutgoingTemplateDeliveryPackageById(id, 
                    page.toString(),
                    pageSize,) as? PaginatedResponse<TemplateDeliveryPackage>
    }
}

/**
 * Lazy pagination iterator for getOutgoingTemplateDeliveryTransporterById.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetOutgoingTemplateDeliveryTransporterById(id: String, pageSize: String? = null): Flow<TemplateDeliveryTransporter> {
    return MetrcExtensions.iteratePages { page ->
        this.getOutgoingTemplateDeliveryTransporterById(id, 
                    page.toString(),
                    pageSize,) as? PaginatedResponse<TemplateDeliveryTransporter>
    }
}

/**
 * Lazy pagination iterator for getOutgoingTemplates.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetOutgoingTemplates(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<Template> {
    return MetrcExtensions.iteratePages { page ->
        this.getOutgoingTemplates(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Template>
    }
}
/**
 * Sync items for getOutgoingTemplates.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun TransfersService.syncGetOutgoingTemplates(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<Template> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetOutgoingTemplates(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getOutgoingTemplates.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetOutgoingTemplatesParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<Template>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = TransfersService(client, limiter)
    
    svc.syncGetOutgoingTemplates(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getOutgoingTransfers.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetOutgoingTransfers(lastModifiedEnd: String? = null, lastModifiedStart: String? = null, licenseNumber: String? = null, pageSize: String? = null): Flow<Transfer> {
    return MetrcExtensions.iteratePages { page ->
        this.getOutgoingTransfers(
                    lastModifiedEnd,
                    lastModifiedStart,
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Transfer>
    }
}
/**
 * Sync items for getOutgoingTransfers.
 * Retrieves all items updated within the specified time window.
 * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
 * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
 * @return List of items updated since the sync window start
 */
suspend fun TransfersService.syncGetOutgoingTransfers(
    lastKnownSync: OffsetDateTime? = null, 
    bufferMinutes: Int = 15
    , licenseNumber: String? = null, pageSize: String? = null
): List<Transfer> {
    val end = OffsetDateTime.now()
    var start = end.minusDays(1)
    if (lastKnownSync != null) {
        start = lastKnownSync.minusMinutes(bufferMinutes.toLong())
    }

    val startStr = start.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)
    val endStr = end.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME)

    return this.iterateGetOutgoingTransfers(endStr, startStr, licenseNumber, pageSize).toList()
}

/**
 * Sync items in parallel across multiple facilities for getOutgoingTransfers.
 * @param targets List of client/license pairs to sync
 * @param lastKnownSync Last known sync timestamp
 * @param bufferMinutes Buffer to subtract from lastKnownSync
 * @return Map of license number to list of synced items
 */
suspend fun syncGetOutgoingTransfersParallel(
    targets: List<Pair<MetrcClient, String>>,
    lastKnownSync: OffsetDateTime? = null,
    bufferMinutes: Int = 15
    , pageSize: String? = null
): Map<String, List<Transfer>> = MetrcExtensions.syncParallel(targets) { client, limiter, license ->
    val svc = TransfersService(client, limiter)
    
    svc.syncGetOutgoingTransfers(
        lastKnownSync, bufferMinutes,
        
        license, pageSize
    )
}

/**
 * Lazy pagination iterator for getRejectedTransfers.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetRejectedTransfers(licenseNumber: String? = null, pageSize: String? = null): Flow<Transfer> {
    return MetrcExtensions.iteratePages { page ->
        this.getRejectedTransfers(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<Transfer>
    }
}

/**
 * Lazy pagination iterator for getTransfersDeliveryById.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetTransfersDeliveryById(id: String, pageSize: String? = null): Flow<TransfersDelivery> {
    return MetrcExtensions.iteratePages { page ->
        this.getTransfersDeliveryById(id, 
                    page.toString(),
                    pageSize,) as? PaginatedResponse<TransfersDelivery>
    }
}

/**
 * Lazy pagination iterator for getTransfersTypes.
 * @return A Flow that yields items lazily across all pages
 */
@Suppress("UNCHECKED_CAST")
fun TransfersService.iterateGetTransfersTypes(licenseNumber: String? = null, pageSize: String? = null): Flow<TransfersType> {
    return MetrcExtensions.iteratePages { page ->
        this.getTransfersTypes(
                    licenseNumber,
                    page.toString(),
                    pageSize,) as? PaginatedResponse<TransfersType>
    }
}

