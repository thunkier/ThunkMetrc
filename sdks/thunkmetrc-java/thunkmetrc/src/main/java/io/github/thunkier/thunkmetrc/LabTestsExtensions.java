package io.github.thunkier.thunkmetrc;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;
import io.github.thunkier.thunkmetrc.wrapper.models.*;
import io.github.thunkier.thunkmetrc.wrapper.services.LabTestsService;

import java.util.*;
import java.util.concurrent.*;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.time.temporal.ChronoUnit;

public class LabTestsExtensions {

    /**
     * Iterator for getLabTestsBatches.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Batch> iterateGetLabTestsBatches(
        LabTestsService service,
        String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getLabTestsBatches(
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
     * Iterator for getLabTestsTypes.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<LabTestsType> iterateGetLabTestsTypes(
        LabTestsService service,
        String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getLabTestsTypes(
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
     * Iterator for getLabTestsResults.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<Result> iterateGetLabTestsResults(
        LabTestsService service,
        String licenseNumber, String packageId, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getLabTestsResults(
                    licenseNumber, packageId, String.valueOf(page), pageSize
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

