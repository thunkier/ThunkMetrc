package io.github.thunkier.thunkmetrc;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;
import io.github.thunkier.thunkmetrc.wrapper.models.*;
import io.github.thunkier.thunkmetrc.wrapper.services.TransportersService;

import java.util.*;
import java.util.concurrent.*;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.time.temporal.ChronoUnit;

public class TransportersExtensions {

    /**
     * Iterator for getTransportersDrivers.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<TransportersDriver> iterateGetTransportersDrivers(
        TransportersService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getTransportersDrivers(
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
     * Iterator for getTransportersVehicles.
     * Returns an Iterable that fetches pages lazily and yields items.
     */
    @SuppressWarnings("unchecked")
    public static Iterable<TransportersVehicle> iterateGetTransportersVehicles(
        TransportersService service,
        String licenseNumber, String pageSize
    ) {return MetrcExtensions.iteratePages(page -> {
            try {
                Object res = service.getTransportersVehicles(
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

