package io.github.thunkier.thunkmetrc;

import io.github.thunkier.thunkmetrc.wrapper.MetrcWrapper;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.util.ArrayList;

import java.util.List;
import java.util.Map;
import java.util.concurrent.*;


public class InventorySync {

    public static class SynchronizationTarget {
        public String licenseNumber;
        public MetrcWrapper wrapper;

        public SynchronizationTarget(String licenseNumber, MetrcWrapper wrapper) {
            this.licenseNumber = licenseNumber;
            this.wrapper = wrapper;
        }
    }

    /**
     * Sync active packages for multiple targets concurrently.
     * @param targets List of targets
     * @param lastKnownSync Optional last sync time
     * @param bufferMinutes Buffer minutes if last sync time is provided
     * @return Map of LicenseNumber to List of Packages (Maps)
     * @throws Exception If sync fails
     */
    public static Map<String, List<Object>> activePackagesInventorySync(
            List<SynchronizationTarget> targets,
            OffsetDateTime lastKnownSync,
            int bufferMinutes
    ) throws Exception {
        
        OffsetDateTime endTime = OffsetDateTime.now();
        OffsetDateTime startTime;
        if (lastKnownSync != null) {
            startTime = lastKnownSync.minusMinutes(bufferMinutes);
        } else {
            startTime = endTime.minusDays(1);
        }

        String startStr = startTime.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME);
        String endStr = endTime.format(DateTimeFormatter.ISO_OFFSET_DATE_TIME);

        ExecutorService executor = Executors.newFixedThreadPool(20);
        Map<String, List<Object>> results = new ConcurrentHashMap<>();

        List<Callable<Void>> tasks = new ArrayList<>();

        for (SynchronizationTarget target : targets) {
            tasks.add(() -> {
                List<Object> packages = new ArrayList<>();
                // Use the generated iterator
                // Method: iteratePackagesGetActiveV2
                // Params: body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
                // We pass null for body, pageNumber, pageSize
                // Order from template:
                // PathParams (none)
                // QueryParams (LastModifiedEnd, LastModifiedStart, LicenseNumber, PageNumber, PageSize)
                // Body (Object) - Wait, wrapper template puts body at end or based on method.
                // Let's assume standard order. Standard generated order puts body if POST/PUT, then QueryParams.
                // GET method: no body arg usually, but helper might put it.
                // Let's check generated wrapper signature via file view if unsure, or rely on template logic.
                // Template logic:
                // PathParams...
                // QueryParams (minus pageNumber)...
                // Body (if POST/PUT)...
                
                // packagesGetActiveV2 is GET. Check template again.
                // Template: 
                // iterate{{MethodName}}( PathParams..., QueryParams(param != pageNumber)..., Body(if POST/PUT) )
                
                // iterateGetActivePackages params: service, lastModifiedEnd, lastModifiedStart, licenseNumber, pageSize
                
                Iterable<?> pages = PackagesExtensions.iterateGetActivePackages(
                        target.wrapper.packages(), endStr, startStr, target.licenseNumber, null
                );

                for (Object page : pages) {
                    if (page instanceof Map) {
                        Map<?, ?> m = (Map<?, ?>) page;
                        Object data = m.get("Data");
                        if (data instanceof List) {
                            packages.addAll((List<?>) data);
                        }
                    } else if (page instanceof List) {
                        packages.addAll((List<?>) page);
                    }
                }
                
                results.put(target.licenseNumber, packages);
                return null;
            });
        }

        try {
            List<Future<Void>> futures = executor.invokeAll(tasks);
            for (Future<Void> f : futures) {
                f.get(); // Check for exceptions
            }
        } finally {
            executor.shutdown();
        }

        return results;
    }
}
