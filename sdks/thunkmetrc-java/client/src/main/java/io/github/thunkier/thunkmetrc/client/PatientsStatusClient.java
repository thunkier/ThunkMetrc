package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class PatientsStatusClient {
    private final MetrcClient client;

    public PatientsStatusClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * Permissions Required:
     * - Lookup Patients
     * GET GetPatientsStatusesByPatientLicenseNumber
     * @param patientLicenseNumber Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPatientsStatusesByPatientLicenseNumber(
        String patientLicenseNumber, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/patients/v2/statuses/"+URLEncoder.encode(patientLicenseNumber, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

}

