package io.github.thunkier.thunkmetrc;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;
import io.github.thunkier.thunkmetrc.wrapper.models.*;
import io.github.thunkier.thunkmetrc.wrapper.services.PackagesService;

import java.util.*;
import java.util.concurrent.*;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.time.temporal.ChronoUnit;

public class PackagesExtensions {

    /**
     * Iterator for getActivePackages.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<PackagesPackage> iterateGetActivePackages(
        PackagesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getActivePackages(
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
     * Sync items for getActivePackages.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<PackagesPackage> syncGetActivePackages(
        PackagesService service,
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
        
        List<PackagesPackage> results = new ArrayList<>();

        for (PackagesPackage item : iterateGetActivePackages(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getActivePackages.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<PackagesPackage>> syncGetActivePackagesParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PackagesService svc = new PackagesService(client, limiter);
            return syncGetActivePackages(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getPackagesAdjustReasons.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<AdjustReason> iterateGetPackagesAdjustReasons(
        PackagesService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getPackagesAdjustReasons(
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
     * Iterator for getPackagesAdjustments.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Adjustment> iterateGetPackagesAdjustments(
        PackagesService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getPackagesAdjustments(
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
     * Iterator for getInTransitPackages.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<InTransit> iterateGetInTransitPackages(
        PackagesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getInTransitPackages(
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
     * Sync items for getInTransitPackages.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<InTransit> syncGetInTransitPackages(
        PackagesService service,
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
        
        List<InTransit> results = new ArrayList<>();

        for (InTransit item : iterateGetInTransitPackages(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getInTransitPackages.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<InTransit>> syncGetInTransitPackagesParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PackagesService svc = new PackagesService(client, limiter);
            return syncGetInTransitPackages(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getInactivePackages.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<PackagesPackage> iterateGetInactivePackages(
        PackagesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getInactivePackages(
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
     * Sync items for getInactivePackages.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<PackagesPackage> syncGetInactivePackages(
        PackagesService service,
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
        
        List<PackagesPackage> results = new ArrayList<>();

        for (PackagesPackage item : iterateGetInactivePackages(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getInactivePackages.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<PackagesPackage>> syncGetInactivePackagesParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PackagesService svc = new PackagesService(client, limiter);
            return syncGetInactivePackages(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getPackagesLabSamples.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<PackagesPackage> iterateGetPackagesLabSamples(
        PackagesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getPackagesLabSamples(
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
     * Sync items for getPackagesLabSamples.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<PackagesPackage> syncGetPackagesLabSamples(
        PackagesService service,
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
        
        List<PackagesPackage> results = new ArrayList<>();

        for (PackagesPackage item : iterateGetPackagesLabSamples(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getPackagesLabSamples.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<PackagesPackage>> syncGetPackagesLabSamplesParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PackagesService svc = new PackagesService(client, limiter);
            return syncGetPackagesLabSamples(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getOnHoldPackages.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<PackagesPackage> iterateGetOnHoldPackages(
        PackagesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getOnHoldPackages(
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
     * Sync items for getOnHoldPackages.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<PackagesPackage> syncGetOnHoldPackages(
        PackagesService service,
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
        
        List<PackagesPackage> results = new ArrayList<>();

        for (PackagesPackage item : iterateGetOnHoldPackages(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getOnHoldPackages.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<PackagesPackage>> syncGetOnHoldPackagesParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PackagesService svc = new PackagesService(client, limiter);
            return syncGetOnHoldPackages(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getPackagesTransferred.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<PackagesPackage> iterateGetPackagesTransferred(
        PackagesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getPackagesTransferred(
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
     * Sync items for getPackagesTransferred.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<PackagesPackage> syncGetPackagesTransferred(
        PackagesService service,
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
        
        List<PackagesPackage> results = new ArrayList<>();

        for (PackagesPackage item : iterateGetPackagesTransferred(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getPackagesTransferred.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<PackagesPackage>> syncGetPackagesTransferredParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PackagesService svc = new PackagesService(client, limiter);
            return syncGetPackagesTransferred(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }
}

