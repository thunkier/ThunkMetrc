package io.github.thunkier.thunkmetrc.client;

import java.io.IOException;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.List;

public class TransportersClient {
    private final MetrcClient client;

    public TransportersClient(MetrcClient client) {
        this.client = client;
    }
    /**
     * Creates new driver records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * POST CreateDrivers
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createTransportersDrivers(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/drivers");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Creates new vehicle records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * POST CreateVehicles
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object createTransportersVehicles(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/vehicles");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("POST", path.toString(), body);
    }

    /**
     * Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
     * Permissions Required:
     * - Manage Transporters
     * DELETE DeleteDriverById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deleteTransportersDriverById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/drivers/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
     * Permissions Required:
     * - Manage Transporters
     * DELETE DeleteVehicleById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object deleteTransportersVehicleById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/vehicles/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("DELETE", path.toString(), null);
    }

    /**
     * Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
     * Permissions Required:
     * - Transporters
     * GET GetDriverById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransportersDriverById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/drivers/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of drivers for a Facility.
     * Permissions Required:
     * - Transporters
     * GET GetDrivers
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransportersDrivers(
        String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/drivers");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
     * Permissions Required:
     * - Transporters
     * GET GetVehicleById
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransportersVehicleById(
        String id, String licenseNumber
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/vehicles/"+URLEncoder.encode(id, StandardCharsets.UTF_8)+"");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Retrieves a list of vehicles for a Facility.
     * Permissions Required:
     * - Transporters
     * GET GetVehicles
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws IOException If request fails
     * @return Response object
     */
    public Object getTransportersVehicles(
        String licenseNumber, String pageNumber, String pageSize
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/vehicles");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (pageNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageNumber=").append(URLEncoder.encode(pageNumber, StandardCharsets.UTF_8));
        }
        if (pageSize != null) {
            if (query.length() > 0) query.append("&");
            query.append("pageSize=").append(URLEncoder.encode(pageSize, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("GET", path.toString(), null);
    }

    /**
     * Updates existing driver records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * PUT UpdateDrivers
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateTransportersDrivers(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/drivers");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

    /**
     * Updates existing vehicle records for a facility.
     * Permissions Required:
     * - Manage Transporters
     * PUT UpdateVehicles
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws IOException If request fails
     * @return Response object
     */
    public Object updateTransportersVehicles(
        String licenseNumber, Object body
    ) throws IOException {
        StringBuilder path = new StringBuilder("/transporters/v2/vehicles");
        StringBuilder query = new StringBuilder();
        if (licenseNumber != null) {
            if (query.length() > 0) query.append("&");
            query.append("licenseNumber=").append(URLEncoder.encode(licenseNumber, StandardCharsets.UTF_8));
        }
        if (query.length() > 0) path.append("?").append(query);

        return client.send("PUT", path.toString(), body);
    }

}

