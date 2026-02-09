package io.github.thunkier.thunkmetrc.wrapper

import io.github.thunkier.thunkmetrc.wrapper.models.PaginatedResponse
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.coroutineScope
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.asFlow
import kotlinx.coroutines.flow.emitAll
import kotlinx.coroutines.flow.flow
import kotlin.math.ceil

/**
 * Extension utilities for Metrc API operations.
 * Provides pagination iteration and parallel sync capabilities.
 */
object MetrcExtensions {

    /**
     * Creates a Flow that automatically fetches pages from a paginated API endpoint.
     * @param fetcher Suspend function that fetches a page given a page number (1-indexed)
     * @return A Flow that emits all items across all pages
     */
    fun <T> iteratePages(fetcher: suspend (Int) -> PaginatedResponse<T>?): Flow<T> = flow {
        var page = 1
        while (true) {
            val res = fetcher(page) ?: break

            val items = res.Data ?: emptyList()
            if (items.isEmpty()) break

            emitAll(items.asFlow())

            val total = res.Total ?: 0
            val pageSize = res.PageSize ?: 0
            if (pageSize > 0) {
                val totalPages = ceil(total.toDouble() / pageSize).toInt()
                if (page >= totalPages) break
            }
            page++
        }
    }

    /**
     * Execute sync operations in parallel across multiple facilities.
     * @param targets List of client/license pairs to sync
     * @param operation The suspend sync operation to execute for each target
     * @return Map of license number to list of synced items
     */
    suspend fun <C, T> syncParallel(
        targets: List<Pair<C, String>>,
        operation: suspend (C, MetrcRateLimiter, String) -> List<T>
    ): Map<String, List<T>> = coroutineScope {
        val sharedLimiter = MetrcRateLimiter()

        targets.map { (client, license) ->
            async {
                val items = operation(client, sharedLimiter, license)
                license to items
            }
        }.awaitAll().toMap()
    }
}
