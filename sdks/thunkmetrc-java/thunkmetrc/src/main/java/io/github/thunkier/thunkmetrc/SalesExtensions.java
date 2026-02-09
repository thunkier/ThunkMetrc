package io.github.thunkier.thunkmetrc;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;
import io.github.thunkier.thunkmetrc.wrapper.models.*;
import io.github.thunkier.thunkmetrc.wrapper.services.SalesService;

import java.util.*;
import java.util.concurrent.*;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.time.temporal.ChronoUnit;

public class SalesExtensions {

    /**
     * Iterator for getSalesActiveDeliveries.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<ActiveDelivery> iterateGetSalesActiveDeliveries(
        SalesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getSalesActiveDeliveries(
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
     * Sync items for getSalesActiveDeliveries.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<ActiveDelivery> syncGetSalesActiveDeliveries(
        SalesService service,
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
        
        List<ActiveDelivery> results = new ArrayList<>();

        for (ActiveDelivery item : iterateGetSalesActiveDeliveries(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getSalesActiveDeliveries.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<ActiveDelivery>> syncGetSalesActiveDeliveriesParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            SalesService svc = new SalesService(client, limiter);
            return syncGetSalesActiveDeliveries(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getSalesActiveDeliveriesRetailer.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<ActiveDeliveriesRetailer> iterateGetSalesActiveDeliveriesRetailer(
        SalesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getSalesActiveDeliveriesRetailer(
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
     * Sync items for getSalesActiveDeliveriesRetailer.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<ActiveDeliveriesRetailer> syncGetSalesActiveDeliveriesRetailer(
        SalesService service,
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
        
        List<ActiveDeliveriesRetailer> results = new ArrayList<>();

        for (ActiveDeliveriesRetailer item : iterateGetSalesActiveDeliveriesRetailer(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getSalesActiveDeliveriesRetailer.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<ActiveDeliveriesRetailer>> syncGetSalesActiveDeliveriesRetailerParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            SalesService svc = new SalesService(client, limiter);
            return syncGetSalesActiveDeliveriesRetailer(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getSalesActiveReceipts.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<ActiveReceipt> iterateGetSalesActiveReceipts(
        SalesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getSalesActiveReceipts(
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
     * Sync items for getSalesActiveReceipts.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<ActiveReceipt> syncGetSalesActiveReceipts(
        SalesService service,
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
        
        List<ActiveReceipt> results = new ArrayList<>();

        for (ActiveReceipt item : iterateGetSalesActiveReceipts(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getSalesActiveReceipts.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<ActiveReceipt>> syncGetSalesActiveReceiptsParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            SalesService svc = new SalesService(client, limiter);
            return syncGetSalesActiveReceipts(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getSalesDeliveriesReturnReasons.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<DeliveriesReturnReason> iterateGetSalesDeliveriesReturnReasons(
        SalesService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getSalesDeliveriesReturnReasons(
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
     * Iterator for getSalesInactiveDeliveries.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<InactiveDelivery> iterateGetSalesInactiveDeliveries(
        SalesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getSalesInactiveDeliveries(
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
     * Sync items for getSalesInactiveDeliveries.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<InactiveDelivery> syncGetSalesInactiveDeliveries(
        SalesService service,
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
        
        List<InactiveDelivery> results = new ArrayList<>();

        for (InactiveDelivery item : iterateGetSalesInactiveDeliveries(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getSalesInactiveDeliveries.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<InactiveDelivery>> syncGetSalesInactiveDeliveriesParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            SalesService svc = new SalesService(client, limiter);
            return syncGetSalesInactiveDeliveries(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getSalesInactiveDeliveriesRetailer.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<InactiveDeliveriesRetailer> iterateGetSalesInactiveDeliveriesRetailer(
        SalesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getSalesInactiveDeliveriesRetailer(
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
     * Sync items for getSalesInactiveDeliveriesRetailer.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<InactiveDeliveriesRetailer> syncGetSalesInactiveDeliveriesRetailer(
        SalesService service,
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
        
        List<InactiveDeliveriesRetailer> results = new ArrayList<>();

        for (InactiveDeliveriesRetailer item : iterateGetSalesInactiveDeliveriesRetailer(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getSalesInactiveDeliveriesRetailer.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<InactiveDeliveriesRetailer>> syncGetSalesInactiveDeliveriesRetailerParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            SalesService svc = new SalesService(client, limiter);
            return syncGetSalesInactiveDeliveriesRetailer(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getSalesInactiveReceipts.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<InactiveReceipt> iterateGetSalesInactiveReceipts(
        SalesService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getSalesInactiveReceipts(
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
     * Sync items for getSalesInactiveReceipts.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<InactiveReceipt> syncGetSalesInactiveReceipts(
        SalesService service,
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
        
        List<InactiveReceipt> results = new ArrayList<>();

        for (InactiveReceipt item : iterateGetSalesInactiveReceipts(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getSalesInactiveReceipts.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<InactiveReceipt>> syncGetSalesInactiveReceiptsParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            SalesService svc = new SalesService(client, limiter);
            return syncGetSalesInactiveReceipts(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }
}

