package io.github.thunkier.thunkmetrc;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;
import io.github.thunkier.thunkmetrc.wrapper.models.*;
import io.github.thunkier.thunkmetrc.wrapper.services.ItemsService;

import java.util.*;
import java.util.concurrent.*;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.time.temporal.ChronoUnit;

public class ItemsExtensions {

    /**
     * Iterator for getActiveItems.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Item> iterateGetActiveItems(
        ItemsService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getActiveItems(
                    lastModifiedEnd, lastModifiedStart, licenseNumber, String.valueOf(page), pageSize
                );

                if (res instanceof PaginatedResponse) {
                    return (PaginatedResponse) res;
                } else if (res instanceof List) {
                    PaginatedResponse pr = new PaginatedResponse();
                    pr.Data = (List) res;
                    return pr;
                }
                return null;
            } catch (Exception e) {
                throw new RuntimeException(e);
            }
        });
    }

    /**
     * Sync items for getActiveItems.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<Item> syncGetActiveItems(
        ItemsService service,
        OffsetDateTime lastKnownSync,
        int bufferMinutes
        , String licenseNumber, String pageSize
    ) {
        OffsetDateTime end = OffsetDateTime.now();
        OffsetDateTime start = end.minusDays(1);
        if (lastKnownSync != null) {
            start = lastKnownSync.minusMinutes(bufferMinutes);
        }
        
        String startStr = start.format(DateTimeFormatter.ISO_INSTANT);
        String endStr = end.format(DateTimeFormatter.ISO_INSTANT);
        
        List<Item> results = new ArrayList<>();

        for (Item item : iterateGetActiveItems(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getActiveItems.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<Item>> syncGetActiveItemsParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            ItemsService svc = new ItemsService(client, limiter);
            return syncGetActiveItems(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getItemsBrands.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Brand> iterateGetItemsBrands(
        ItemsService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getItemsBrands(
                    licenseNumber, String.valueOf(page), pageSize
                );

                if (res instanceof PaginatedResponse) {
                    return (PaginatedResponse) res;
                } else if (res instanceof List) {
                    PaginatedResponse pr = new PaginatedResponse();
                    pr.Data = (List) res;
                    return pr;
                }
                return null;
            } catch (Exception e) {
                throw new RuntimeException(e);
            }
        });
    }

    /**
     * Iterator for getItemsCategories.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Category> iterateGetItemsCategories(
        ItemsService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getItemsCategories(
                    licenseNumber, String.valueOf(page), pageSize
                );

                if (res instanceof PaginatedResponse) {
                    return (PaginatedResponse) res;
                } else if (res instanceof List) {
                    PaginatedResponse pr = new PaginatedResponse();
                    pr.Data = (List) res;
                    return pr;
                }
                return null;
            } catch (Exception e) {
                throw new RuntimeException(e);
            }
        });
    }

    /**
     * Iterator for getInactiveItems.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Item> iterateGetInactiveItems(
        ItemsService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getInactiveItems(
                    licenseNumber, String.valueOf(page), pageSize
                );

                if (res instanceof PaginatedResponse) {
                    return (PaginatedResponse) res;
                } else if (res instanceof List) {
                    PaginatedResponse pr = new PaginatedResponse();
                    pr.Data = (List) res;
                    return pr;
                }
                return null;
            } catch (Exception e) {
                throw new RuntimeException(e);
            }
        });
    }
}

