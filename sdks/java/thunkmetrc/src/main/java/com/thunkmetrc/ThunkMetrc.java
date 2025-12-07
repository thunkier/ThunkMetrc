package com.thunkmetrc;

import com.thunkmetrc.wrapper.MetrcWrapper;
import java.time.OffsetDateTime;
import java.time.format.DateTimeFormatter;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

public class ThunkMetrc {
    private final MetrcWrapper wrapper;
    // Format: yyyy-MM-ddTHH:mm:ssK (no fractional seconds)
    private static final DateTimeFormatter DATE_FORMATTER = DateTimeFormatter.ofPattern("yyyy-MM-dd'T'HH:mm:ssXXX");

    public ThunkMetrc(MetrcWrapper wrapper) {
        this.wrapper = wrapper;
    }

    public List<Object> activePackagesInventorySync(
            String licenseNumber,
            OffsetDateTime lastModifiedStart,
            Integer bufferMinutes
    ) throws Exception {
        List<Object> allPackages = new ArrayList<>();
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
            // Arguments: lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, body (null)
            // Order from generated wrapper:
            // Path Params (none for this call?) -> Wait, Packages/v2/active has NO path params.
            // URL: /packages/v2/active
            // Args: licenseNumber (query), etc.
            // Client/Wrapper generated Query Params are sorted alphabetically.
            // Params: lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize.
            // All strings.
            
            // Signature: activePackagesInventorySync(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize, Object body)
            
            Object result = wrapper.packagesGetActiveV2(
                endStr,
                startStr,
                licenseNumber,
                String.valueOf(pageNumber),
                String.valueOf(pageSize)
            );

            if (result == null) break;

            if (result instanceof List) {
                List<?> list = (List<?>) result;
                if (list.isEmpty()) break;
                allPackages.addAll(list);
                pageNumber++;
            } else {
                // Unexpected response type (maybe error object if not thrown?)
                // Assuming client throws on error.
                break;
            }
        }
        return allPackages;
    }
}
