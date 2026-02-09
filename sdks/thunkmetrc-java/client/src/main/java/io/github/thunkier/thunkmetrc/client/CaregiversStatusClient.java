package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class CaregiversStatusClient {
    private final MetrcClient client;

    public CaregiversStatusClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * Permissions Required:
     * - Lookup Caregivers
     * GET GetCaregiversStatusByCaregiverLicenseNumber
     * @param caregiverLicenseNumber Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getCaregiversStatusByCaregiverLicenseNumber(
        String caregiverLicenseNumber, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/caregivers/v2/status/"+URLEncoder.encode(caregiverLicenseNumber, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

}

