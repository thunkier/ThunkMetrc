package io.github.thunkier.thunkmetrc.wrapper;

import io.github.thunkier.thunkmetrc.wrapper.models.PaginatedResponse;

import java.util.ArrayList;
import java.util.Collections;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.function.Function;

/**
 * Extension utilities for Metrc API operations.
 * Provides pagination iteration and parallel sync capabilities.
 */
public class MetrcExtensions {

    /**
     * Creates an iterable that automatically fetches pages from a paginated API endpoint.
     * @param fetcher Function that fetches a page given a page number (1-indexed)
     * @param <T> The type of items in the paginated response
     * @return An iterable that yields all items across all pages
     */
    public static <T> Iterable<T> iteratePages(Function<Integer, PaginatedResponse<T>> fetcher) {
        return () -> new PageIterator<>(fetcher);
    }

    /**
     * Represents a sync target with a client and license number.
     * @param <C> The client type
     */
    public static class SyncTarget<C> {
        public final C client;
        public final String licenseNumber;

        public SyncTarget(C client, String licenseNumber) {
            this.client = client;
            this.licenseNumber = licenseNumber;
        }
    }

    /**
     * Functional interface for sync operations.
     * @param <C> The client type
     * @param <T> The item type
     */
    @FunctionalInterface
    public interface SyncOperation<C, T> {
        List<T> execute(C client, MetrcRateLimiter limiter, String licenseNumber) throws Exception;
    }

    /**
     * Execute sync operations in parallel across multiple facilities.
     * @param targets List of client/license pairs to sync
     * @param concurrencyLimit Maximum number of concurrent operations
     * @param operation The sync operation to execute for each target
     * @param <C> The client type
     * @param <T> The item type
     * @return Map of license number to list of synced items
     */
    public static <C, T> Map<String, List<T>> syncParallel(
        List<SyncTarget<C>> targets,
        int concurrencyLimit,
        SyncOperation<C, T> operation
    ) {
        Map<String, List<T>> results = new ConcurrentHashMap<>();
        ExecutorService executor = Executors.newFixedThreadPool(concurrencyLimit);
        MetrcRateLimiter sharedLimiter = new MetrcRateLimiter();

        List<Future<?>> futures = new ArrayList<>();

        for (SyncTarget<C> target : targets) {
            futures.add(executor.submit(() -> {
                try {
                    List<T> items = operation.execute(target.client, sharedLimiter, target.licenseNumber);
                    results.put(target.licenseNumber, items);
                } catch (Exception e) {
                    throw new RuntimeException("Failed to sync license: " + target.licenseNumber, e);
                }
            }));
        }

        for (Future<?> f : futures) {
            try {
                f.get();
            } catch (Exception e) {
                throw new RuntimeException(e);
            }
        }
        executor.shutdown();

        return results;
    }

    private static class PageIterator<T> implements Iterator<T> {
        private final Function<Integer, PaginatedResponse<T>> fetcher;
        private int currentPage = 1;
        private Iterator<T> currentBuffer = Collections.emptyIterator();
        private boolean hasMorePages = true;

        PageIterator(Function<Integer, PaginatedResponse<T>> fetcher) {
            this.fetcher = fetcher;
        }

        @Override
        public boolean hasNext() {
            if (currentBuffer.hasNext()) {
                return true;
            }
            if (!hasMorePages) {
                return false;
            }

            PaginatedResponse<T> response = fetcher.apply(currentPage);
            if (response == null || response.Data == null || response.Data.isEmpty()) {
                hasMorePages = false;
                return false;
            }

            currentBuffer = response.Data.iterator();

            int totalPages = Integer.MAX_VALUE;
            if (response.Total != null && response.PageSize != null && response.PageSize > 0) {
                totalPages = (int) Math.ceil((double) response.Total / response.PageSize);
            }

            if (currentPage >= totalPages) {
                hasMorePages = false;
            }

            currentPage++;
            return currentBuffer.hasNext();
        }

        @Override
        public T next() {
            return currentBuffer.next();
        }
    }
}
