package io.github.thunkier.thunkmetrc;

import io.github.thunkier.thunkmetrc.wrapper.MetrcWrapper;
import io.github.thunkier.thunkmetrc.wrapper.models.PackagesPackage;
import io.github.thunkier.thunkmetrc.wrapper.models.PaginatedResponse;

import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.util.ArrayList;
import java.util.List;

public class ThunkMetrc {
    private final MetrcWrapper wrapper;
    private static final DateTimeFormatter DATE_FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd'T'HH:mm:ssXXX");

    public ThunkMetrc(MetrcWrapper wrapper) {
        this.wrapper = wrapper;
    }

    public List<PackagesPackage> activePackagesInventorySync(
            String licenseNumber,
            OffsetDateTime lastModifiedStart,
            Integer bufferMinutes
    ) throws Exception {
        List<PackagesPackage> allPackages = new ArrayList<>();
        int pageNumber = 1;
        int pageSize = 20;

        String startStr = null;
        if (lastModifiedStart != null) {
            startStr = lastModifiedStart.format(DATE_FORMATTER);
        }

        String endStr = null;
        if (bufferMinutes != null && bufferMinutes > 0) {
             endStr = OffsetDateTime.now().minusMinutes(bufferMinutes).format(DATE_FORMATTER);
        }

        while (true) {
            PaginatedResponse<PackagesPackage> result = (PaginatedResponse<PackagesPackage>) wrapper.packages().getActivePackages(
                endStr,
                startStr,
                licenseNumber,
                String.valueOf(pageNumber),
                String.valueOf(pageSize)
            );

            if (result == null || result.Data == null || result.Data.isEmpty()) break;
            
            allPackages.addAll(result.Data);
            pageNumber++;
        }
        return allPackages;
    }
}

