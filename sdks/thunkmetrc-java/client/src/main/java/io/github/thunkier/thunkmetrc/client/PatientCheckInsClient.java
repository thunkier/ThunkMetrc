package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class PatientCheckInsClient {
    private final MetrcClient client;

    public PatientCheckInsClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Records patient check-ins for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * POST CreatePatientCheckIns
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createPatientCheckIns(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v2");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Archives a Patient Check-In, identified by its Id, for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * DELETE DeletePatientCheckInsById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deletePatientCheckInsById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v2/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Retrieves a list of Patient Check-In locations.
     * GET GetLocations
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPatientCheckInsLocations(
        
    ) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v2/locations");
        StringBuilder query = new StringBuilder();
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of patient check-ins for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * GET GetPatientCheckIns
     * @param checkinDateEnd Query parameter
     * @param checkinDateStart Query parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getPatientCheckIns(
        String checkinDateEnd, String checkinDateStart, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v2");
        StringBuilder query = new StringBuilder();
        if (checkinDateEnd != null) {
            if (query.length() > 0) query.append("&");
            query.append("checkinDateEnd=").append(URLEncoder.encode(checkinDateEnd, StandardCharsets.UTF_8));
        }
        if (checkinDateStart != null) {
            if (query.length() > 0) query.append("&");
            query.append("checkinDateStart=").append(URLEncoder.encode(checkinDateStart, StandardCharsets.UTF_8));
        }
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Updates patient check-ins for a specified Facility.
     * Permissions Required:
     * - ManagePatientsCheckIns
     * PUT UpdatePatientCheckIns
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updatePatientCheckIns(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/patient-checkins/v2");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

}

