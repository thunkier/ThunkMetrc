package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class TransportersService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public TransportersService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Creates new driver records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * POST /transporters/v2/drivers
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createTransportersDrivers(
        String licenseNumber, List<TransportersCreateDriversRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transporters.createTransportersDrivers(
                licenseNumber, body
            )
        );
    }

    /**
     * Creates new vehicle records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * POST /transporters/v2/vehicles
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createTransportersVehicles(
        String licenseNumber, List<TransportersCreateVehiclesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transporters.createTransportersVehicles(
                licenseNumber, body
            )
        );
    }

    /**
     * Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
     * Permissions Required:
     * - Manage Transporters
     * DELETE /transporters/v2/drivers/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deleteTransportersDriverById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.transporters.deleteTransportersDriverById(
                id, licenseNumber
            )
        );
    }

    /**
     * Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
     * Permissions Required:
     * - Manage Transporters
     * DELETE /transporters/v2/vehicles/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deleteTransportersVehicleById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.transporters.deleteTransportersVehicleById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
     * Permissions Required:
     * - Transporters
     * GET /transporters/v2/drivers/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public TransportersDriver getTransportersDriverById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (TransportersDriver) client.transporters.getTransportersDriverById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of drivers for a Facility.
     * Permissions Required:
     * - Transporters
     * GET /transporters/v2/drivers
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<TransportersDriver> getTransportersDrivers(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<TransportersDriver>) client.transporters.getTransportersDrivers(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
     * Permissions Required:
     * - Transporters
     * GET /transporters/v2/vehicles/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public TransportersVehicle getTransportersVehicleById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (TransportersVehicle) client.transporters.getTransportersVehicleById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of vehicles for a Facility.
     * Permissions Required:
     * - Transporters
     * GET /transporters/v2/vehicles
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<TransportersVehicle> getTransportersVehicles(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<TransportersVehicle>) client.transporters.getTransportersVehicles(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Updates existing driver records for a Facility.
     * Permissions Required:
     * - Manage Transporters
     * PUT /transporters/v2/drivers
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updateTransportersDrivers(
        String licenseNumber, List<TransportersUpdateDriversRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transporters.updateTransportersDrivers(
                licenseNumber, body
            )
        );
    }

    /**
     * Updates existing vehicle records for a facility.
     * Permissions Required:
     * - Manage Transporters
     * PUT /transporters/v2/vehicles
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updateTransportersVehicles(
        String licenseNumber, List<TransportersUpdateVehiclesRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.transporters.updateTransportersVehicles(
                licenseNumber, body
            )
        );
    }

}

