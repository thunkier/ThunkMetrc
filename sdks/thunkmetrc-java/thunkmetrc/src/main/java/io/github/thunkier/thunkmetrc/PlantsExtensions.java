package io.github.thunkier.thunkmetrc;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;
import io.github.thunkier.thunkmetrc.wrapper.models.*;
import io.github.thunkier.thunkmetrc.wrapper.services.PlantsService;

import java.util.*;
import java.util.concurrent.*;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.time.temporal.ChronoUnit;

public class PlantsExtensions {

    /**
     * Iterator for getPlantsAdditives.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Additive> iterateGetPlantsAdditives(
        PlantsService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getPlantsAdditives(
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
     * Sync items for getPlantsAdditives.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<Additive> syncGetPlantsAdditives(
        PlantsService service,
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
        
        List<Additive> results = new ArrayList<>();

        for (Additive item : iterateGetPlantsAdditives(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getPlantsAdditives.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<Additive>> syncGetPlantsAdditivesParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PlantsService svc = new PlantsService(client, limiter);
            return syncGetPlantsAdditives(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getFloweringPlants.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Plant> iterateGetFloweringPlants(
        PlantsService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getFloweringPlants(
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
     * Sync items for getFloweringPlants.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<Plant> syncGetFloweringPlants(
        PlantsService service,
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
        
        List<Plant> results = new ArrayList<>();

        for (Plant item : iterateGetFloweringPlants(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getFloweringPlants.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<Plant>> syncGetFloweringPlantsParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PlantsService svc = new PlantsService(client, limiter);
            return syncGetFloweringPlants(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getInactivePlants.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Plant> iterateGetInactivePlants(
        PlantsService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getInactivePlants(
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
     * Sync items for getInactivePlants.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<Plant> syncGetInactivePlants(
        PlantsService service,
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
        
        List<Plant> results = new ArrayList<>();

        for (Plant item : iterateGetInactivePlants(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getInactivePlants.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<Plant>> syncGetInactivePlantsParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PlantsService svc = new PlantsService(client, limiter);
            return syncGetInactivePlants(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getMotherInactivePlants.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Mother> iterateGetMotherInactivePlants(
        PlantsService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getMotherInactivePlants(
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
     * Sync items for getMotherInactivePlants.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<Mother> syncGetMotherInactivePlants(
        PlantsService service,
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
        
        List<Mother> results = new ArrayList<>();

        for (Mother item : iterateGetMotherInactivePlants(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getMotherInactivePlants.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<Mother>> syncGetMotherInactivePlantsParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PlantsService svc = new PlantsService(client, limiter);
            return syncGetMotherInactivePlants(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getMotherOnHoldPlants.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Mother> iterateGetMotherOnHoldPlants(
        PlantsService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getMotherOnHoldPlants(
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
     * Sync items for getMotherOnHoldPlants.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<Mother> syncGetMotherOnHoldPlants(
        PlantsService service,
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
        
        List<Mother> results = new ArrayList<>();

        for (Mother item : iterateGetMotherOnHoldPlants(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getMotherOnHoldPlants.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<Mother>> syncGetMotherOnHoldPlantsParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PlantsService svc = new PlantsService(client, limiter);
            return syncGetMotherOnHoldPlants(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getMotherPlants.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Mother> iterateGetMotherPlants(
        PlantsService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getMotherPlants(
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
     * Sync items for getMotherPlants.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<Mother> syncGetMotherPlants(
        PlantsService service,
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
        
        List<Mother> results = new ArrayList<>();

        for (Mother item : iterateGetMotherPlants(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getMotherPlants.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<Mother>> syncGetMotherPlantsParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PlantsService svc = new PlantsService(client, limiter);
            return syncGetMotherPlants(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getOnHoldPlants.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Plant> iterateGetOnHoldPlants(
        PlantsService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getOnHoldPlants(
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
     * Sync items for getOnHoldPlants.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<Plant> syncGetOnHoldPlants(
        PlantsService service,
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
        
        List<Plant> results = new ArrayList<>();

        for (Plant item : iterateGetOnHoldPlants(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getOnHoldPlants.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<Plant>> syncGetOnHoldPlantsParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PlantsService svc = new PlantsService(client, limiter);
            return syncGetOnHoldPlants(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getPlantsWaste.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<PlantsWaste> iterateGetPlantsWaste(
        PlantsService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getPlantsWaste(
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
     * Iterator for getPlantsWasteMethods.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<WasteMethod> iterateGetPlantsWasteMethods(
        PlantsService service,
        String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getPlantsWasteMethods(
                    String.valueOf(page), pageSize
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
     * Iterator for getPlantsWasteReasons.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<WasteReason> iterateGetPlantsWasteReasons(
        PlantsService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getPlantsWasteReasons(
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
     * Iterator for getVegetativePlants.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Plant> iterateGetVegetativePlants(
        PlantsService service,
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getVegetativePlants(
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
     * Sync items for getVegetativePlants.
     * Retrieves all items updated within the specified time window.
     * @param service The service instance
     * @param lastKnownSync Last known sync timestamp (uses 1 day ago if null)
     * @param bufferMinutes Buffer to subtract from lastKnownSync to catch overlaps
     * @return List of items updated since the sync window start
     */
    public static List<Plant> syncGetVegetativePlants(
        PlantsService service,
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
        
        List<Plant> results = new ArrayList<>();

        for (Plant item : iterateGetVegetativePlants(service, endStr, startStr, licenseNumber, pageSize)) {
            results.add(item);
        }
        return results;
    }

    /**
     * Sync items in parallel across multiple facilities for getVegetativePlants.
     * @param targets List of client/license pairs to sync
     * @param lastKnownSync Last known sync timestamp
     * @param bufferMinutes Buffer to subtract from lastKnownSync
     * @param concurrencyLimit Maximum number of concurrent sync operations
     * @return Map of license number to list of synced items
     */
    public static Map<String, List<Plant>> syncGetVegetativePlantsParallel(
        List<MetrcExtensions.SyncTarget<MetrcClient>> targets,
        OffsetDateTime lastKnownSync,
        int bufferMinutes,
        int concurrencyLimit
        , String pageSize
    ) {
        return MetrcExtensions.syncParallel(targets, concurrencyLimit, (client, limiter, license) -> {
            PlantsService svc = new PlantsService(client, limiter);
            return syncGetVegetativePlants(
                svc,
                lastKnownSync, bufferMinutes,
                
                license, pageSize
            );
        });
    }

    /**
     * Iterator for getPlantsWasteById.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<PlantsWaste> iterateGetPlantsWasteById(
        PlantsService service,
        String id, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getPlantsWasteById(
                    id, licenseNumber, String.valueOf(page), pageSize
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
     * Iterator for getPlantsWastePackageById.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<WastePackage> iterateGetPlantsWastePackageById(
        PlantsService service,
        String id, String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getPlantsWastePackageById(
                    id, licenseNumber, String.valueOf(page), pageSize
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

