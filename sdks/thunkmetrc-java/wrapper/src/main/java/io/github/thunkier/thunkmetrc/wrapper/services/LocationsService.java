package io.github.thunkier.thunkmetrc.wrapper.services;

import io.github.thunkier.thunkmetrc.client.MetrcClient;
import io.github.thunkier.thunkmetrc.wrapper.MetrcRateLimiter;import io.github.thunkier.thunkmetrc.wrapper.MetrcExtensions;import io.github.thunkier.thunkmetrc.wrapper.models.*;import java.util.List;

public class LocationsService {
    private final MetrcClient client;
    private final MetrcRateLimiter rateLimiter;

    public LocationsService(MetrcClient client, MetrcRateLimiter rateLimiter) {
        this.client = client;
        this.rateLimiter = rateLimiter;
    }

    /**
     * Creates new locations for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * POST /locations/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse createLocations(
        String licenseNumber, List<LocationsCreateLocationsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.locations.createLocations(
                licenseNumber, body
            )
        );
    }

    /**
     * Archives a specified Location, identified by its Id, for a Facility.
     * Permissions Required:
     * - Manage Locations
     * DELETE /locations/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public Object deleteLocationsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (Object) client.locations.deleteLocationsById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of active locations for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * GET /locations/v2/active
     * @param lastModifiedEnd Query parameter
     * @param lastModifiedStart Query parameter
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<LocationsLocation> getActiveLocations(
        String lastModifiedEnd, String lastModifiedStart, String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<LocationsLocation>) client.locations.getActiveLocations(
                lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a list of inactive locations for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * GET /locations/v2/inactive
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<LocationsLocation> getInactiveLocations(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<LocationsLocation>) client.locations.getInactiveLocations(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Retrieves a Location by its Id.
     * Permissions Required:
     * - Manage Locations
     * GET /locations/v2/{id}
     * @param id Path parameter
     * @param licenseNumber Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    public LocationsLocation getLocationsById(
        String id, String licenseNumber
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (LocationsLocation) client.locations.getLocationsById(
                id, licenseNumber
            )
        );
    }

    /**
     * Retrieves a list of active location types for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * GET /locations/v2/types
     * @param licenseNumber Query parameter
     * @param pageNumber Query parameter
     * @param pageSize Query parameter
     * @throws Exception If request fails
     * @return Response object
     */
    @SuppressWarnings("unchecked")
    public PaginatedResponse<LocationsType> getLocationsTypes(
        String licenseNumber, String pageNumber, String pageSize
    ) throws Exception {
        return rateLimiter.execute(null,true,
            () -> (PaginatedResponse<LocationsType>) client.locations.getLocationsTypes(
                licenseNumber, pageNumber, pageSize
            )
        );
    }

    /**
     * Updates existing locations for a specified Facility.
     * Permissions Required:
     * - Manage Locations
     * PUT /locations/v2
     * @param licenseNumber Query parameter
     * @param body Request body
     * @throws Exception If request fails
     * @return Response object
     */
    public WriteResponse updateLocations(
        String licenseNumber, List<LocationsUpdateLocationsRequestItem> body
    ) throws Exception {
        return rateLimiter.execute(null,false,
            () -> (WriteResponse) client.locations.updateLocations(
                licenseNumber, body
            )
        );
    }

}

