package io.github.thunkier.thunkmetrc;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;
import io.github.thunkier.thunkmetrc.wrapper.models.*;
import io.github.thunkier.thunkmetrc.wrapper.services.LocationsService;

import java.util.*;
import java.util.concurrent.*;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.time.temporal.ChronoUnit;

public class LocationsExtensions {

    /**
     * Iterator for getActiveLocations.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<LocationsLocation> iterateGetActiveLocations(
        LocationsService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getActiveLocations(
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
     * Sync items for getActiveLocations.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<LocationsLocation> syncGetActiveLocations(
        LocationsService service,
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
        
        List<LocationsLocation> results = new ArrayList<>();

        for (LocationsLocation item : iterateGetActiveLocations(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getActiveLocations.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<LocationsLocation>> syncGetActiveLocationsParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            LocationsService svc = new LocationsService(client, limiter);
            return syncGetActiveLocations(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getInactiveLocations.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<LocationsLocation> iterateGetInactiveLocations(
        LocationsService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getInactiveLocations(
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
     * Iterator for getLocationsTypes.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<LocationsType> iterateGetLocationsTypes(
        LocationsService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getLocationsTypes(
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

